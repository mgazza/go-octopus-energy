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

// ListElectricityTariffDayUnitRatesReader is a Reader for the ListElectricityTariffDayUnitRates structure.
type ListElectricityTariffDayUnitRatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListElectricityTariffDayUnitRatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListElectricityTariffDayUnitRatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/products/{product_code}/electricity-tariffs/{tariff_code}/day-unit-rates/] List Electricity Tariff Day Unit Rates", response, response.Code())
	}
}

// NewListElectricityTariffDayUnitRatesOK creates a ListElectricityTariffDayUnitRatesOK with default headers values
func NewListElectricityTariffDayUnitRatesOK() *ListElectricityTariffDayUnitRatesOK {
	return &ListElectricityTariffDayUnitRatesOK{}
}

/*
ListElectricityTariffDayUnitRatesOK describes a response with status code 200, with default header values.

ListElectricityTariffDayUnitRatesOK list electricity tariff day unit rates o k
*/
type ListElectricityTariffDayUnitRatesOK struct {
	Payload *models.PaginatedHistoricalChargeList
}

// IsSuccess returns true when this list electricity tariff day unit rates o k response has a 2xx status code
func (o *ListElectricityTariffDayUnitRatesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list electricity tariff day unit rates o k response has a 3xx status code
func (o *ListElectricityTariffDayUnitRatesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list electricity tariff day unit rates o k response has a 4xx status code
func (o *ListElectricityTariffDayUnitRatesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list electricity tariff day unit rates o k response has a 5xx status code
func (o *ListElectricityTariffDayUnitRatesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list electricity tariff day unit rates o k response a status code equal to that given
func (o *ListElectricityTariffDayUnitRatesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list electricity tariff day unit rates o k response
func (o *ListElectricityTariffDayUnitRatesOK) Code() int {
	return 200
}

func (o *ListElectricityTariffDayUnitRatesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/products/{product_code}/electricity-tariffs/{tariff_code}/day-unit-rates/][%d] listElectricityTariffDayUnitRatesOK %s", 200, payload)
}

func (o *ListElectricityTariffDayUnitRatesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/products/{product_code}/electricity-tariffs/{tariff_code}/day-unit-rates/][%d] listElectricityTariffDayUnitRatesOK %s", 200, payload)
}

func (o *ListElectricityTariffDayUnitRatesOK) GetPayload() *models.PaginatedHistoricalChargeList {
	return o.Payload
}

func (o *ListElectricityTariffDayUnitRatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaginatedHistoricalChargeList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
