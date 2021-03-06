module github.com/izumiya/working/product-api

go 1.16

require (
	github.com/go-openapi/runtime v0.19.27
	github.com/go-playground/validator/v10 v10.5.0
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/hashicorp/go-hclog v0.16.0
	github.com/izumiya/working/currency v0.0.0
	github.com/nicholasjackson/env v0.6.0
	golang.org/x/tools v0.0.0-20190617190820-da514acc4774
	google.golang.org/grpc v1.37.0
)

replace github.com/izumiya/working/currency => ../currency
