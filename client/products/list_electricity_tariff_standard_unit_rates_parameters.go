// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewListElectricityTariffStandardUnitRatesParams creates a new ListElectricityTariffStandardUnitRatesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListElectricityTariffStandardUnitRatesParams() *ListElectricityTariffStandardUnitRatesParams {
	return &ListElectricityTariffStandardUnitRatesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListElectricityTariffStandardUnitRatesParamsWithTimeout creates a new ListElectricityTariffStandardUnitRatesParams object
// with the ability to set a timeout on a request.
func NewListElectricityTariffStandardUnitRatesParamsWithTimeout(timeout time.Duration) *ListElectricityTariffStandardUnitRatesParams {
	return &ListElectricityTariffStandardUnitRatesParams{
		timeout: timeout,
	}
}

// NewListElectricityTariffStandardUnitRatesParamsWithContext creates a new ListElectricityTariffStandardUnitRatesParams object
// with the ability to set a context for a request.
func NewListElectricityTariffStandardUnitRatesParamsWithContext(ctx context.Context) *ListElectricityTariffStandardUnitRatesParams {
	return &ListElectricityTariffStandardUnitRatesParams{
		Context: ctx,
	}
}

// NewListElectricityTariffStandardUnitRatesParamsWithHTTPClient creates a new ListElectricityTariffStandardUnitRatesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListElectricityTariffStandardUnitRatesParamsWithHTTPClient(client *http.Client) *ListElectricityTariffStandardUnitRatesParams {
	return &ListElectricityTariffStandardUnitRatesParams{
		HTTPClient: client,
	}
}

/*
ListElectricityTariffStandardUnitRatesParams contains all the parameters to send to the API endpoint

	for the list electricity tariff standard unit rates operation.

	Typically these are written to a http.Request.
*/
type ListElectricityTariffStandardUnitRatesParams struct {

	/* Page.

	   A page number within the paginated result set.
	*/
	Page *int64

	/* PageSize.

	   Page size of returned results. Default is 100, maximum is 1500 to give up to a month of half-hourly prices.
	*/
	PageSize *int64

	/* PeriodFrom.

	   Show results from the given datetime (inclusive). This parameter can be provided on its own.

	   Format: datetime
	*/
	PeriodFrom *strfmt.DateTime

	/* PeriodTo.

	   Show results up to the given datetime (exclusive). You must also provide the period_from parameter in order to create a range.

	   Format: datetime
	*/
	PeriodTo *strfmt.DateTime

	/* ProductCode.

	   The code of the product to be retrieved, for example VAR-17-01-11.
	*/
	ProductCode string

	/* TariffCode.

	   The code of the tariff to be retrieved, for example E-1R-VAR-17-01-11-A.
	*/
	TariffCode string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list electricity tariff standard unit rates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListElectricityTariffStandardUnitRatesParams) WithDefaults() *ListElectricityTariffStandardUnitRatesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list electricity tariff standard unit rates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListElectricityTariffStandardUnitRatesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list electricity tariff standard unit rates params
func (o *ListElectricityTariffStandardUnitRatesParams) WithTimeout(timeout time.Duration) *ListElectricityTariffStandardUnitRatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list electricity tariff standard unit rates params
func (o *ListElectricityTariffStandardUnitRatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list electricity tariff standard unit rates params
func (o *ListElectricityTariffStandardUnitRatesParams) WithContext(ctx context.Context) *ListElectricityTariffStandardUnitRatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list electricity tariff standard unit rates params
func (o *ListElectricityTariffStandardUnitRatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list electricity tariff standard unit rates params
func (o *ListElectricityTariffStandardUnitRatesParams) WithHTTPClient(client *http.Client) *ListElectricityTariffStandardUnitRatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list electricity tariff standard unit rates params
func (o *ListElectricityTariffStandardUnitRatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the list electricity tariff standard unit rates params
func (o *ListElectricityTariffStandardUnitRatesParams) WithPage(page *int64) *ListElectricityTariffStandardUnitRatesParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the list electricity tariff standard unit rates params
func (o *ListElectricityTariffStandardUnitRatesParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the list electricity tariff standard unit rates params
func (o *ListElectricityTariffStandardUnitRatesParams) WithPageSize(pageSize *int64) *ListElectricityTariffStandardUnitRatesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the list electricity tariff standard unit rates params
func (o *ListElectricityTariffStandardUnitRatesParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithPeriodFrom adds the periodFrom to the list electricity tariff standard unit rates params
func (o *ListElectricityTariffStandardUnitRatesParams) WithPeriodFrom(periodFrom *strfmt.DateTime) *ListElectricityTariffStandardUnitRatesParams {
	o.SetPeriodFrom(periodFrom)
	return o
}

// SetPeriodFrom adds the periodFrom to the list electricity tariff standard unit rates params
func (o *ListElectricityTariffStandardUnitRatesParams) SetPeriodFrom(periodFrom *strfmt.DateTime) {
	o.PeriodFrom = periodFrom
}

// WithPeriodTo adds the periodTo to the list electricity tariff standard unit rates params
func (o *ListElectricityTariffStandardUnitRatesParams) WithPeriodTo(periodTo *strfmt.DateTime) *ListElectricityTariffStandardUnitRatesParams {
	o.SetPeriodTo(periodTo)
	return o
}

// SetPeriodTo adds the periodTo to the list electricity tariff standard unit rates params
func (o *ListElectricityTariffStandardUnitRatesParams) SetPeriodTo(periodTo *strfmt.DateTime) {
	o.PeriodTo = periodTo
}

// WithProductCode adds the productCode to the list electricity tariff standard unit rates params
func (o *ListElectricityTariffStandardUnitRatesParams) WithProductCode(productCode string) *ListElectricityTariffStandardUnitRatesParams {
	o.SetProductCode(productCode)
	return o
}

// SetProductCode adds the productCode to the list electricity tariff standard unit rates params
func (o *ListElectricityTariffStandardUnitRatesParams) SetProductCode(productCode string) {
	o.ProductCode = productCode
}

// WithTariffCode adds the tariffCode to the list electricity tariff standard unit rates params
func (o *ListElectricityTariffStandardUnitRatesParams) WithTariffCode(tariffCode string) *ListElectricityTariffStandardUnitRatesParams {
	o.SetTariffCode(tariffCode)
	return o
}

// SetTariffCode adds the tariffCode to the list electricity tariff standard unit rates params
func (o *ListElectricityTariffStandardUnitRatesParams) SetTariffCode(tariffCode string) {
	o.TariffCode = tariffCode
}

// WriteToRequest writes these params to a swagger request
func (o *ListElectricityTariffStandardUnitRatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Page != nil {

		// query param page
		var qrPage int64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param page_size
		var qrPageSize int64

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("page_size", qPageSize); err != nil {
				return err
			}
		}
	}

	if o.PeriodFrom != nil {

		// query param period_from
		var qrPeriodFrom strfmt.DateTime

		if o.PeriodFrom != nil {
			qrPeriodFrom = *o.PeriodFrom
		}
		qPeriodFrom := qrPeriodFrom.String()
		if qPeriodFrom != "" {

			if err := r.SetQueryParam("period_from", qPeriodFrom); err != nil {
				return err
			}
		}
	}

	if o.PeriodTo != nil {

		// query param period_to
		var qrPeriodTo strfmt.DateTime

		if o.PeriodTo != nil {
			qrPeriodTo = *o.PeriodTo
		}
		qPeriodTo := qrPeriodTo.String()
		if qPeriodTo != "" {

			if err := r.SetQueryParam("period_to", qPeriodTo); err != nil {
				return err
			}
		}
	}

	// path param product_code
	if err := r.SetPathParam("product_code", o.ProductCode); err != nil {
		return err
	}

	// path param tariff_code
	if err := r.SetPathParam("tariff_code", o.TariffCode); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
