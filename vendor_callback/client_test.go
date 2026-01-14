package vendor_callback

import (
	"context"
	"testing"
)

func TestClient_ProcessCode(t *testing.T) {
	appID := "app_ayAkIFfeDazdfsDY"
	appSecret := "lClISPPOv0wbiIZEQMlHWDk5wnpdIIES"

	baseUrl := "http://192.168.2.81:8046"

	client := NewClient(appID, appSecret, WithBaseURL(baseUrl))

	req := &ProcessCodeRequest{
		Code:   "1146e66d87",
		UserID: "793",
	}

	resp, err := client.ProcessCode(context.Background(), req)
	if err != nil {
		t.Fatalf("ProcessCode failed: %v", err)
	}

	if !resp.Success {
		t.Error("Expected Success true")
	}
}

func TestClient_CheckCode(t *testing.T) {
	appID := "app_ayAkIFfeDazdfsDY"
	appSecret := "lClISPPOv0wbiIZEQMlHWDk5wnpdIIES"

	baseUrl := "http://192.168.2.81:8046"

	client := NewClient(appID, appSecret, WithBaseURL(baseUrl))

	req := &CheckCodeRequest{
		Code:   "EUD2X6W6SR",
		UserID: "790",
	}

	resp, err := client.CheckCode(context.Background(), req)
	if resp == nil || err != nil {
		t.Fatalf("CheckCode failed: %v", err)
	}
}

func TestClient_PaymentCallback(t *testing.T) {
	appID := "app_ayAkIFfeDazdfsDY"
	appSecret := "lClISPPOv0wbiIZEQMlHWDk5wnpdIIES"

	baseUrl := "http://192.168.2.81:8046"

	client := NewClient(appID, appSecret, WithBaseURL(baseUrl))

	req := &PaymentCallbackRequest{
		UserID:    "793",
		OrderNo:   "ORD_001",
		Amount:    0.01,
		ProductID: "5664568",
	}

	resp, err := client.PaymentCallback(context.Background(), req)
	if err != nil {
		t.Fatalf("PaymentCallback failed: %v", err)
	}
	if !resp.Success {
		t.Logf("PaymentCallback failed logic: %s", resp.Reason)
	}
}
