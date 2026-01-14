package vendor_callback

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

const (
	DefaultBaseURL = "https://huoying.shangancheng.com"
)

type Client struct {
	AppID      string
	AppSecret  string
	BaseURL    string
	HTTPClient *http.Client
}

type Option func(*Client)

func WithBaseURL(url string) Option {
	return func(c *Client) {
		c.BaseURL = url
	}
}

func WithHTTPClient(client *http.Client) Option {
	return func(c *Client) {
		c.HTTPClient = client
	}
}

func NewClient(appID, appSecret string, opts ...Option) *Client {
	c := &Client{
		AppID:      appID,
		AppSecret:  appSecret,
		BaseURL:    DefaultBaseURL,
		HTTPClient: http.DefaultClient,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// generateSignature 生成签名
// 算法: MD5(X-App-Id + X-Timestamp + AppSecret)
func (c *Client) generateSignature(timestamp int64) string {
	raw := c.AppID + strconv.FormatInt(timestamp, 10) + c.AppSecret
	hasher := md5.New()
	hasher.Write([]byte(raw))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (c *Client) doRequest(ctx context.Context, method, path string, reqBody interface{}, respData interface{}) error {
	var bodyReader io.Reader
	if reqBody != nil {
		jsonBytes, err := json.Marshal(reqBody)
		if err != nil {
			return fmt.Errorf("marshal request body failed: %w", err)
		}
		bodyReader = bytes.NewReader(jsonBytes)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.BaseURL+path, bodyReader)
	if err != nil {
		return fmt.Errorf("create request failed: %w", err)
	}

	// Set Headers
	timestamp := time.Now().Unix()
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-App-Id", c.AppID)
	req.Header.Set("X-Timestamp", strconv.FormatInt(timestamp, 10))
	req.Header.Set("X-Sign", c.generateSignature(timestamp))

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("do request failed: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response body failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(bodyBytes))
	}

	// Decode generic ApiResponse first to check code
	var genericResp ApiResponse[json.RawMessage]
	if err := json.Unmarshal(bodyBytes, &genericResp); err != nil {
		return fmt.Errorf("unmarshal response failed: %w, body: %s", err, string(bodyBytes))
	}

	if genericResp.Code != 0 {
		return fmt.Errorf("api error: code=%d, msg=%s", genericResp.Code, genericResp.Msg)
	}

	// Unmarshal data if needed
	if respData != nil && genericResp.Data != nil {
		if err := json.Unmarshal(genericResp.Data, respData); err != nil {
			return fmt.Errorf("unmarshal data failed: %w", err)
		}
	}

	return nil
}

// CheckCode 校验码信息接口
func (c *Client) CheckCode(ctx context.Context, req *CheckCodeRequest) (*CheckCodeResponseData, error) {
	var data CheckCodeResponseData
	err := c.doRequest(ctx, http.MethodPost, "/api/v1/callback/check-code", req, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// ProcessCode 统一码处理接口
func (c *Client) ProcessCode(ctx context.Context, req *ProcessCodeRequest) (*ProcessCodeResponseData, error) {
	var data ProcessCodeResponseData
	err := c.doRequest(ctx, http.MethodPost, "/api/v1/callback/process-code", req, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// PaymentCallback 支付上报接口
func (c *Client) PaymentCallback(ctx context.Context, req *PaymentCallbackRequest) (*PaymentCallbackResponseData, error) {
	var data PaymentCallbackResponseData
	err := c.doRequest(ctx, http.MethodPost, "/api/v1/callback/payment", req, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
