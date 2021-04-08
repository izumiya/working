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
	p.l.Println("[DEBUG] get all records")

	prods := data.GetProducts()

	err := prods.ToJSON(rw)
	if err != nil {
		// we should never be here but log the error just in case
		p.l.Println("[ERROR] serializing product", err)
	}
}

// swagger:route GET /products/{id} products listSingle
// Return a list of products from the database
// responses:
//   200: productResponse
//   404: errorResponse

// ListSingle handles GET requests
func (p *Products) ListSingle(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	p.l.Println("[DEBUG] get record id", id)

	prod, err := data.GetProductByID(id)

	err = prod.ToJSON(rw)
	if err != nil {
		// we should never be here but log the error just in case
		p.l.Println("[ERROR] serializing product", err)
	}
}
