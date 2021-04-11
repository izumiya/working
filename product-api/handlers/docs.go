package handlers

import "github.com/izumiya/working/product-api/data"

// A list of products returns in the response
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in the system
	// in: body
	Body []data.Product
}

// swagger:response productResponse
type productResponseWrapper struct {
	// in: body
	Body data.Product
}

// No content is returned by this API endpoint
// swagger:response noContent
type noContentResponseWrapper struct {
}

// swagger:response errorValidation
type errorValidationWrapper struct {
}

// swagger:response errorResponse
type errorResponseWrapper struct {
}

// swagger:parameters listSingleProduct deleteProduct updateProduct
type productIDParameterWrapper struct {
	// The id of the product for which the operation relates
	// in: path
	// required: true
	ID int `json:"id"`
}

