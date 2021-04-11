package handlers

import (
	"net/http"

	"github.com/izumiya/working/product-api/data"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Returns a list of products
// responses:
//   201: noContent
//   422: errorValidation
//   501: errorResponse

// DeleteProduct deletes a product from  the database
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	p.l.Debug("delete record", "id", id)

	err := data.DeleteProduct(id)

	if err == data.ErrProductNotFound {
		p.l.Error("unable to delete record id does not exist", "id", id, "error", err)
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	if err != nil {
		p.l.Error("unable to delete product", "error", err)
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}
}
