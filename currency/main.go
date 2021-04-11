package main

import (
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/izumiya/working/currency/data"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	protos "github.com/izumiya/working/currency/protos/currency"
	"github.com/izumiya/working/currency/server"
)

func main() {
	log := hclog.Default()

	rates, err := data.NewRates(log)
	if err != nil {
		log.Error("unable to generate rates", "error", err)
	}

	gs := grpc.NewServer()
	cs := server.NewCurrency(rates, log)

	protos.RegisterCurrencyServer(gs, cs)

	reflection.Register(gs)

	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("unable to listen", "error", err)
		os.Exit(1)
	}

	gs.Serve(l)
}
