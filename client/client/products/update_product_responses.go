// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateProductReader is a Reader for the UpdateProduct structure.
type UpdateProductReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProductReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateProductCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewUpdateProductUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewUpdateProductNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateProductCreated creates a UpdateProductCreated with default headers values
func NewUpdateProductCreated() *UpdateProductCreated {
	return &UpdateProductCreated{}
}

/* UpdateProductCreated describes a response with status code 201, with default header values.

No content is returned by this API endpoint
*/
type UpdateProductCreated struct {
}

func (o *UpdateProductCreated) Error() string {
	return fmt.Sprintf("[PUT /products/{id}][%d] updateProductCreated ", 201)
}

func (o *UpdateProductCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateProductUnprocessableEntity creates a UpdateProductUnprocessableEntity with default headers values
func NewUpdateProductUnprocessableEntity() *UpdateProductUnprocessableEntity {
	return &UpdateProductUnprocessableEntity{}
}

/* UpdateProductUnprocessableEntity describes a response with status code 422, with default header values.

UpdateProductUnprocessableEntity update product unprocessable entity
*/
type UpdateProductUnprocessableEntity struct {
}

func (o *UpdateProductUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /products/{id}][%d] updateProductUnprocessableEntity ", 422)
}

func (o *UpdateProductUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateProductNotImplemented creates a UpdateProductNotImplemented with default headers values
func NewUpdateProductNotImplemented() *UpdateProductNotImplemented {
	return &UpdateProductNotImplemented{}
}

/* UpdateProductNotImplemented describes a response with status code 501, with default header values.

UpdateProductNotImplemented update product not implemented
*/
type UpdateProductNotImplemented struct {
}

func (o *UpdateProductNotImplemented) Error() string {
	return fmt.Sprintf("[PUT /products/{id}][%d] updateProductNotImplemented ", 501)
}

func (o *UpdateProductNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
