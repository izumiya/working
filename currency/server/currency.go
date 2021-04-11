package server

import (
	"context"
	"io"
	"time"

	"github.com/hashicorp/go-hclog"
	"github.com/izumiya/working/currency/data"

	protos "github.com/izumiya/working/currency/protos/currency"
)

type Currency struct {
	rates         *data.ExchangeRates
	log           hclog.Logger
	subscriptions map[protos.Currency_SubscribeRatesServer][]*protos.RateRequest
	protos.UnimplementedCurrencyServer
}

func NewCurrency(r *data.ExchangeRates, l hclog.Logger) *Currency {
	c := &Currency{rates: r, log: l, subscriptions: make(map[protos.Currency_SubscribeRatesServer][]*protos.RateRequest)}
	go c.handleUpdates()

	return c
}

func (c *Currency) handleUpdates() {
	ru := c.rates.MonitorRates(5 * time.Second)
	for range ru {
		c.log.Info("Got Updated rates")

		for k, v := range c.subscriptions {

			for _, rr := range v {
				r, err := c.rates.GetRate(rr.Base.String(), rr.Destination.String())
				if err != nil {
					c.log.Error("unable to update rate", "base", rr.GetBase().String(), "destination", rr.GetDestination().String())
				}

				err = k.Send(&protos.RateResponse{Base: rr.Base, Destination: rr.Destination, Rate: r})
				if err != nil {
					c.log.Error("unable to send updated rate", "base", rr.GetBase().String(), "destination", rr.GetDestination().String())
				}
			}
		}
	}
}

func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())

	rate, err := c.rates.GetRate(rr.GetBase().String(), rr.GetDestination().String())
	if err != nil {
		return nil, err
	}

	return &protos.RateResponse{Base: rr.GetBase(), Destination: rr.GetDestination(), Rate: rate}, nil
}

func (c *Currency) SubscribeRates(src protos.Currency_SubscribeRatesServer) error {
	for {
		rr, err := src.Recv()
		if err == io.EOF {
			c.log.Info("client has closed connection")
			return err
		}

		if err != nil {
			c.log.Error("unable to read from client", "error", err)
			return err
		}

		c.log.Info("handle cliunet request", "request", rr)

		rrs, ok := c.subscriptions[src]
		if !ok {
			rrs = []*protos.RateRequest{}
		}

		rrs = append(rrs, rr)
		c.subscriptions[src] = rrs
	}
	return nil
}
