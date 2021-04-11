package handlers

import (
	"context"
	"net/http"

	protos "github.com/izumiya/working/currency/protos/currency"
	"github.com/izumiya/working/product-api/data"
)

// swagger:route GET /products products listProducts
// Returns a list of products from the database
// responses:
//   200: productsResponse

// ListAll handles GET requests and returns all current products
func (p *Products) ListAll(rw http.ResponseWriter, r *http.Request) {
	p.l.Info("[DEBUG] get all records")

	rw.Header().Add("Content-Type", "application/json")

	prods := data.GetProducts()

	err := prods.ToJSON(rw)
	if err != nil {
		// we should never be here but log the error just in case
		p.l.Error("[ERROR] serializing product", "error", err)
	}
}

// swagger:route GET /products/{id} products listSingleProduct
// Return a list of products from the database
// responses:
//   200: productResponse
//   404: errorResponse

// ListSingleProduct handles GET requests
func (p *Products) ListSingleProduct(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	p.l.Info("[DEBUG] get record id", "id", id)

	prod, err := data.GetProductByID(id)

	switch err {
	case nil:
	case data.ErrProductNotFound:
		p.l.Error("fetching product", "error", err)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	default:
		p.l.Error("fetching product", "error", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	// get exchange rate
	rr := &protos.RateRequest{
		Base:        protos.Currencies(protos.Currencies_value["EUR"]),
		Destination: protos.Currencies(protos.Currencies_value["GBP"]),
	}
	resp, err := p.cc.GetRate(context.Background(), rr)
	if err != nil {
		p.l.Error("error getting new rate", "error", err)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	p.l.Info("resp", "response", resp)

	prod.Price = prod.Price * resp.Rate

	err = data.ToJSON(prod, rw)
	if err != nil {
		// we should never be here but log the error just in case
		p.l.Error("serializing product", "error", err)
	}
}
