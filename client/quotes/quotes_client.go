// Code generated by go-swagger; DO NOT EDIT.

package quotes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new quotes API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new quotes API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new quotes API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for quotes API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateaQuote(params *CreateaQuoteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateAQuoteCreated, error)

	ShareaQuoteViaEmail(params *ShareaQuoteViaEmailParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ShareAQuoteViaEmailNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateaQuote At least one electricity or gas meter-point must be included. A maximum of one electricity and one gas meter-point can be included at this time. For all meter-points, at least one of gsp or postcode must be included. All meter-points must be based in the UK, excluding Northern Ireland. All meter-points must belong to the same gsp or postcode. WARNING: This endpoint is only available to partner organisations.
*/
func (a *Client) CreateaQuote(params *CreateaQuoteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateAQuoteCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateaQuoteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Create a quote",
		Method:             "POST",
		PathPattern:        "/v1/quotes/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateaQuoteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateAQuoteCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Create a quote: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ShareaQuoteViaEmail Use this endpoint after quote creation to send a quote summary email to the specified recipients if they wish to enact the quote at a later time. A quote share record is saved to the database for each recipient’s email address. WARNING: This endpoint is only available to partner organisations.
*/
func (a *Client) ShareaQuoteViaEmail(params *ShareaQuoteViaEmailParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ShareAQuoteViaEmailNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewShareaQuoteViaEmailParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Share a quote via email",
		Method:             "POST",
		PathPattern:        "/v1/quotes/{code}/products/{product_id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ShareaQuoteViaEmailReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ShareAQuoteViaEmailNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Share a quote via email: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
