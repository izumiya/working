package server

import (
	"context"

	"github.com/hashicorp/go-hclog"

	protos "github.com/izumiya/working/currency/protos/currency"
)

type Currency struct {
	protos.UnimplementedCurrencyServer
	log hclog.Logger
}

func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{log: l}
}

func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())

	return &protos.RateResponse{Rate: 0.5}, nil
}
