package server

import (
	"context"

	"github.com/hashicorp/go-hclog"
	"github.com/izumiya/working/currency/data"

	protos "github.com/izumiya/working/currency/protos/currency"
)

type Currency struct {
	protos.UnimplementedCurrencyServer
	rate *data.ExchangeRates
	log  hclog.Logger
}

func NewCurrency(r *data.ExchangeRates, l hclog.Logger) *Currency {
	return &Currency{rate: r, log: l}
}

func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())

	rate, err := c.rate.GetRate(rr.GetBase().String(), rr.GetDestination().String())
	if err != nil {
		return nil, err
	}

	return &protos.RateResponse{Rate: rate}, nil
}
