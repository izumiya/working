package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/izumiya/working/product-api/data"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Returns a list of products
// responses:
//   201: noContent

// DeleteProduct deletes a product from  the database
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
	}

	p.l.Println("Handle DELETE Product", id)

	err = data.DeleteProduct(id)

	if err == data.ErrProductNotFound {
		http.Error(rw, "product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "unable to delete product", http.StatusInternalServerError)
		return
	}
}
