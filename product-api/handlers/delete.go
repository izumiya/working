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

	p.l.Info("[DEBUG] delete record", "id", id)

	err := data.DeleteProduct(id)

	if err == data.ErrProductNotFound {
		http.Error(rw, "product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "unable to delete product", http.StatusInternalServerError)
		return
	}
}
