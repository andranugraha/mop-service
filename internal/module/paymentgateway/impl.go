package paymentgateway

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/empnefsi/mop-service/internal/common/logger"
	"github.com/empnefsi/mop-service/internal/common/strings"
	"github.com/empnefsi/mop-service/internal/config"
)

func (m *impl) ChargePayment(ctx context.Context, req *PaymentRequest) (*PaymentResponse, error) {
	cmd := config.GetMidtransURL() + PaymentEndpoint

	payload, err := json.Marshal(req)
	if err != nil {
		logger.Error(ctx, "ChargePayment", "failed to marshal payment request: %v", err.Error())
		return nil, fmt.Errorf("failed to marshal payment request: %w", err)
	}
	request, err := http.NewRequest(http.MethodPost, cmd, bytes.NewBuffer(payload))
	if err != nil {
		logger.Error(ctx, "ChargePayment", "failed to create payment request: %v", err.Error())
		return nil, fmt.Errorf("failed to create payment request: %w", err)
	}

	request.Header.Set("accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", getAuthHeader())

	client := &http.Client{Timeout: 5 * time.Second}
	response, err := client.Do(request)
	if err != nil {
		logger.Error(ctx, "ChargePayment", "failed to send payment request: %v", err.Error())
		return nil, fmt.Errorf("failed to send payment request: %w", err)
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(response.Body)
		logger.Error(ctx, "ChargePayment", "unexpected status code: %d, body: %s", response.StatusCode, string(body))
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", response.StatusCode, string(body))
	}

	var paymentResponse PaymentResponse
	err = json.NewDecoder(response.Body).Decode(&paymentResponse)
	if err != nil {
		logger.Error(ctx, "ChargePayment", "failed to decode payment response: %v", err.Error())
		return nil, fmt.Errorf("failed to decode payment response: %w", err)
	}

	if paymentResponse.StatusCode != "201" {
		logger.Error(ctx, "ChargePayment", "unexpected status code: %s", paymentResponse.StatusCode)
		return nil, fmt.Errorf("unexpected status code: %s", paymentResponse.StatusCode)
	}

	return &paymentResponse, nil
}

func getAuthHeader() string {
	authKey := config.GetMidtransServerKey() + ":"
	return "Basic " + strings.Base64Encode([]byte(authKey))
}
