package vendor_callback

import (
	"context"
	"testing"
)

func TestClient_ProcessCode(t *testing.T) {
	appID := "test_app_id"
	appSecret := "test_app_secret"

	baseUrl := ""

	client := NewClient(appID, appSecret, WithBaseURL(baseUrl))

	req := &ProcessCodeRequest{
		Code:       "ABC-123",
		UserID:     "u_888",
		ActionTime: 1700000000,
	}

	resp, err := client.ProcessCode(context.Background(), req)
	if err != nil {
		t.Fatalf("ProcessCode failed: %v", err)
	}

	if !resp.Success {
		t.Error("Expected Success true")
	}
	if resp.Type != "invitation" {
		t.Errorf("Expected type invitation, got %s", resp.Type)
	}
	if resp.Reward.Value != 7 {
		t.Errorf("Expected reward value 7, got %d", resp.Reward.Value)
	}
}

func TestClient_PaymentCallback(t *testing.T) {
	appID := "test_app_id"
	appSecret := "test_app_secret"

	baseUrl := ""

	client := NewClient(appID, appSecret, WithBaseURL(baseUrl))

	req := &PaymentCallbackRequest{
		UserID:    "u_888",
		OrderNo:   "ORD_001",
		Amount:    100.0,
		ProductID: 1001,
	}

	err := client.PaymentCallback(context.Background(), req)
	if err != nil {
		t.Fatalf("PaymentCallback failed: %v", err)
	}
}
