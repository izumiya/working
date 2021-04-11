package handlers

import (
	"net/http"

	"github.com/izumiya/working/product-api/data"
)

// swagger:route GET /products products listProducts
// Returns a list of products from the database
// responses:
//   200: productsResponse

// ListAll handles GET requests and returns all current products
func (p *Products) ListAll(rw http.ResponseWriter, r *http.Request) {
	p.l.Debug("get all records")

	rw.Header().Add("Content-Type", "application/json")

	cur := r.URL.Query().Get("currency")

	prods, err := p.productsDB.GetProducts(cur)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	err = data.ToJSON(prods, rw)
	if err != nil {
		// we should never be here but log the error just in case
		p.l.Error("unable to serializing product", "error", err)
	}
}

// swagger:route GET /products/{id} products listSingleProduct
// Return a list of products from the database
// responses:
//   200: productResponse
//   404: errorResponse

// listSingleProduct handles GET requests
func (p *Products) ListSingle(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	id := getProductID(r)
	cur := r.URL.Query().Get("currency")

	p.l.Debug("get record", "id", id)

	prod, err := p.productsDB.GetProductByID(id, cur)

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

	err = data.ToJSON(prod, rw)
	if err != nil {
		// we should never be here but log the error just in case
		p.l.Error("serializing product", "error", err)
	}
}
