// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteProductReader is a Reader for the DeleteProduct structure.
type DeleteProductReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProductReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDeleteProductCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewDeleteProductUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewDeleteProductNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteProductCreated creates a DeleteProductCreated with default headers values
func NewDeleteProductCreated() *DeleteProductCreated {
	return &DeleteProductCreated{}
}

/* DeleteProductCreated describes a response with status code 201, with default header values.

No content is returned by this API endpoint
*/
type DeleteProductCreated struct {
}

func (o *DeleteProductCreated) Error() string {
	return fmt.Sprintf("[DELETE /products/{id}][%d] deleteProductCreated ", 201)
}

func (o *DeleteProductCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProductUnprocessableEntity creates a DeleteProductUnprocessableEntity with default headers values
func NewDeleteProductUnprocessableEntity() *DeleteProductUnprocessableEntity {
	return &DeleteProductUnprocessableEntity{}
}

/* DeleteProductUnprocessableEntity describes a response with status code 422, with default header values.

DeleteProductUnprocessableEntity delete product unprocessable entity
*/
type DeleteProductUnprocessableEntity struct {
}

func (o *DeleteProductUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /products/{id}][%d] deleteProductUnprocessableEntity ", 422)
}

func (o *DeleteProductUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProductNotImplemented creates a DeleteProductNotImplemented with default headers values
func NewDeleteProductNotImplemented() *DeleteProductNotImplemented {
	return &DeleteProductNotImplemented{}
}

/* DeleteProductNotImplemented describes a response with status code 501, with default header values.

DeleteProductNotImplemented delete product not implemented
*/
type DeleteProductNotImplemented struct {
}

func (o *DeleteProductNotImplemented) Error() string {
	return fmt.Sprintf("[DELETE /products/{id}][%d] deleteProductNotImplemented ", 501)
}

func (o *DeleteProductNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}