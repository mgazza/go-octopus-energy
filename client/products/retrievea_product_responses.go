// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mgazza/go-octopus-energy/models"
)

// RetrieveaProductReader is a Reader for the RetrieveaProduct structure.
type RetrieveaProductReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetrieveaProductReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRetrieveAProductOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/products/{product_code}/] Retrieve a Product", response, response.Code())
	}
}

// NewRetrieveAProductOK creates a RetrieveAProductOK with default headers values
func NewRetrieveAProductOK() *RetrieveAProductOK {
	return &RetrieveAProductOK{}
}

/*
RetrieveAProductOK describes a response with status code 200, with default header values.

RetrieveAProductOK retrieve a product o k
*/
type RetrieveAProductOK struct {
	Payload *models.Product
}

// IsSuccess returns true when this retrieve a product o k response has a 2xx status code
func (o *RetrieveAProductOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this retrieve a product o k response has a 3xx status code
func (o *RetrieveAProductOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this retrieve a product o k response has a 4xx status code
func (o *RetrieveAProductOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this retrieve a product o k response has a 5xx status code
func (o *RetrieveAProductOK) IsServerError() bool {
	return false
}

// IsCode returns true when this retrieve a product o k response a status code equal to that given
func (o *RetrieveAProductOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the retrieve a product o k response
func (o *RetrieveAProductOK) Code() int {
	return 200
}

func (o *RetrieveAProductOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/products/{product_code}/][%d] retrieveAProductOK %s", 200, payload)
}

func (o *RetrieveAProductOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/products/{product_code}/][%d] retrieveAProductOK %s", 200, payload)
}

func (o *RetrieveAProductOK) GetPayload() *models.Product {
	return o.Payload
}

func (o *RetrieveAProductOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Product)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
