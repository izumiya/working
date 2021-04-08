package handlers

import (
	"net/http"

	"github.com/izumiya/working/product-api/data"
)

// swagger:route PUT /products/{id} products updateProduct
// Update a product
//
// responses:
//   201: noContent
//   422: errorValidation
//   501: errorResponse

func (p *Products) Update(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	p.l.Println("[DEBUG] update record", id)

	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	err := data.UpdateProduct(id, prod)

	if err == data.ErrProductNotFound {
		http.Error(rw, "product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "unable to update product", http.StatusInternalServerError)
		return
	}
}
