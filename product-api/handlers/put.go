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

	p.l.Debug("update record", "id", id)

	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	err := p.productsDB.UpdateProduct(id, prod)

	if err == data.ErrProductNotFound {
		p.l.Error("unable to update product id not exists", "id", id, "error", err)
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	if err != nil {
		p.l.Error("unable to update product", "error", err)
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}
}
