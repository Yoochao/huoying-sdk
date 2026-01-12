package vendor_callback

import (
	"context"
	"testing"
)

func TestClient_ProcessCode(t *testing.T) {
	appID := "app_"
	appSecret := "sec_"

	baseUrl := "http://192.168.2.81:8046"

	client := NewClient(appID, appSecret, WithBaseURL(baseUrl))

	req := &ProcessCodeRequest{
		Code:   "fe54527f49",
		UserID: "789",
	}

	resp, err := client.ProcessCode(context.Background(), req)
	if err != nil {
		t.Fatalf("ProcessCode failed: %v", err)
	}

	if !resp.Success {
		t.Error("Expected Success true")
	}
}

func TestClient_PaymentCallback(t *testing.T) {
	appID := "app_"
	appSecret := "sec_"

	baseUrl := "http://192.168.2.81:8046"

	client := NewClient(appID, appSecret, WithBaseURL(baseUrl))

	req := &PaymentCallbackRequest{
		UserID:    "789",
		OrderNo:   "ORD_001",
		Amount:    100.0,
		ProductID: "5664560",
	}

	err := client.PaymentCallback(context.Background(), req)
	if err != nil {
		t.Fatalf("PaymentCallback failed: %v", err)
	}
}
