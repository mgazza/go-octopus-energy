// Code generated by go-swagger; DO NOT EDIT.

package accounts

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

	"github.com/mgazza/go-octopus-energy/models"
)

// NewRenewaBusinessTariffParams creates a new RenewaBusinessTariffParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRenewaBusinessTariffParams() *RenewaBusinessTariffParams {
	return &RenewaBusinessTariffParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRenewaBusinessTariffParamsWithTimeout creates a new RenewaBusinessTariffParams object
// with the ability to set a timeout on a request.
func NewRenewaBusinessTariffParamsWithTimeout(timeout time.Duration) *RenewaBusinessTariffParams {
	return &RenewaBusinessTariffParams{
		timeout: timeout,
	}
}

// NewRenewaBusinessTariffParamsWithContext creates a new RenewaBusinessTariffParams object
// with the ability to set a context for a request.
func NewRenewaBusinessTariffParamsWithContext(ctx context.Context) *RenewaBusinessTariffParams {
	return &RenewaBusinessTariffParams{
		Context: ctx,
	}
}

// NewRenewaBusinessTariffParamsWithHTTPClient creates a new RenewaBusinessTariffParams object
// with the ability to set a custom HTTPClient for a request.
func NewRenewaBusinessTariffParamsWithHTTPClient(client *http.Client) *RenewaBusinessTariffParams {
	return &RenewaBusinessTariffParams{
		HTTPClient: client,
	}
}

/*
RenewaBusinessTariffParams contains all the parameters to send to the API endpoint

	for the renew a business tariff operation.

	Typically these are written to a http.Request.
*/
type RenewaBusinessTariffParams struct {

	/* AccountNumber.

	   Number of the account
	*/
	AccountNumber string

	// Body.
	Body *models.BusinessTariffRenewal

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the renew a business tariff params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RenewaBusinessTariffParams) WithDefaults() *RenewaBusinessTariffParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the renew a business tariff params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RenewaBusinessTariffParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the renew a business tariff params
func (o *RenewaBusinessTariffParams) WithTimeout(timeout time.Duration) *RenewaBusinessTariffParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the renew a business tariff params
func (o *RenewaBusinessTariffParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the renew a business tariff params
func (o *RenewaBusinessTariffParams) WithContext(ctx context.Context) *RenewaBusinessTariffParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the renew a business tariff params
func (o *RenewaBusinessTariffParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the renew a business tariff params
func (o *RenewaBusinessTariffParams) WithHTTPClient(client *http.Client) *RenewaBusinessTariffParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the renew a business tariff params
func (o *RenewaBusinessTariffParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountNumber adds the accountNumber to the renew a business tariff params
func (o *RenewaBusinessTariffParams) WithAccountNumber(accountNumber string) *RenewaBusinessTariffParams {
	o.SetAccountNumber(accountNumber)
	return o
}

// SetAccountNumber adds the accountNumber to the renew a business tariff params
func (o *RenewaBusinessTariffParams) SetAccountNumber(accountNumber string) {
	o.AccountNumber = accountNumber
}

// WithBody adds the body to the renew a business tariff params
func (o *RenewaBusinessTariffParams) WithBody(body *models.BusinessTariffRenewal) *RenewaBusinessTariffParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the renew a business tariff params
func (o *RenewaBusinessTariffParams) SetBody(body *models.BusinessTariffRenewal) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RenewaBusinessTariffParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param account_number
	if err := r.SetPathParam("account_number", o.AccountNumber); err != nil {
		return err
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
