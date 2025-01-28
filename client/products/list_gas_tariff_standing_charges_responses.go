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

// ListGasTariffStandingChargesReader is a Reader for the ListGasTariffStandingCharges structure.
type ListGasTariffStandingChargesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListGasTariffStandingChargesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListGasTariffStandingChargesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/products/{product_code}/gas-tariffs/{tariff_code}/standing-charges/] List Gas Tariff Standing Charges", response, response.Code())
	}
}

// NewListGasTariffStandingChargesOK creates a ListGasTariffStandingChargesOK with default headers values
func NewListGasTariffStandingChargesOK() *ListGasTariffStandingChargesOK {
	return &ListGasTariffStandingChargesOK{}
}

/*
ListGasTariffStandingChargesOK describes a response with status code 200, with default header values.

ListGasTariffStandingChargesOK list gas tariff standing charges o k
*/
type ListGasTariffStandingChargesOK struct {
	Payload *models.PaginatedHistoricalChargeList
}

// IsSuccess returns true when this list gas tariff standing charges o k response has a 2xx status code
func (o *ListGasTariffStandingChargesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list gas tariff standing charges o k response has a 3xx status code
func (o *ListGasTariffStandingChargesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list gas tariff standing charges o k response has a 4xx status code
func (o *ListGasTariffStandingChargesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list gas tariff standing charges o k response has a 5xx status code
func (o *ListGasTariffStandingChargesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list gas tariff standing charges o k response a status code equal to that given
func (o *ListGasTariffStandingChargesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list gas tariff standing charges o k response
func (o *ListGasTariffStandingChargesOK) Code() int {
	return 200
}

func (o *ListGasTariffStandingChargesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/products/{product_code}/gas-tariffs/{tariff_code}/standing-charges/][%d] listGasTariffStandingChargesOK %s", 200, payload)
}

func (o *ListGasTariffStandingChargesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/products/{product_code}/gas-tariffs/{tariff_code}/standing-charges/][%d] listGasTariffStandingChargesOK %s", 200, payload)
}

func (o *ListGasTariffStandingChargesOK) GetPayload() *models.PaginatedHistoricalChargeList {
	return o.Payload
}

func (o *ListGasTariffStandingChargesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaginatedHistoricalChargeList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
