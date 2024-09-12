package paymentgateway

import (
	"context"
	"sync"
)

var (
	module Module
	once   sync.Once
)

type Module interface {
	ChargePayment(ctx context.Context, req *PaymentRequest) (*PaymentResponse, error)
}

type impl struct {
}

func GetModule() Module {
	once.Do(func() {
		module = &impl{}
	})
	return module
}
