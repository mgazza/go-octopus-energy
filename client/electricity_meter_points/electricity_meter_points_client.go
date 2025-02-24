// Code generated by go-swagger; DO NOT EDIT.

package electricity_meter_points

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new electricity meter points API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new electricity meter points API client with basic auth credentials.
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

// New creates a new electricity meter points API client with a bearer token for authentication.
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
Client for electricity meter points API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetElectricityMeterPoint(params *GetElectricityMeterPointParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetElectricityMeterPointOK, error)

	ListConsumptionForAnElectricityMeter(params *ListConsumptionForAnElectricityMeterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListConsumptionForAnElectricityMeterOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetElectricityMeterPoint Retrieve the details of a meter-point. This endpoint can be used to get the GSP of a given meter-point.
*/
func (a *Client) GetElectricityMeterPoint(params *GetElectricityMeterPointParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetElectricityMeterPointOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetElectricityMeterPointParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Get Electricity Meter Point",
		Method:             "GET",
		PathPattern:        "/v1/electricity-meter-points/{mpan}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetElectricityMeterPointReader{formats: a.formats},
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
	success, ok := result.(*GetElectricityMeterPointOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Get Electricity Meter Point: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListConsumptionForAnElectricityMeter Return a list of consumption values (in kWh) for half-hour periods for a given meter-point and meter. WARNING: Half-hourly consumption data is only available for smart meters. Requests for consumption data for non-smart meters will return an empty response payload.
*/
func (a *Client) ListConsumptionForAnElectricityMeter(params *ListConsumptionForAnElectricityMeterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListConsumptionForAnElectricityMeterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListConsumptionForAnElectricityMeterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "List consumption for an electricity meter",
		Method:             "GET",
		PathPattern:        "/v1/electricity-meter-points/{mpan}/meters/{serial_number}/consumption/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListConsumptionForAnElectricityMeterReader{formats: a.formats},
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
	success, ok := result.(*ListConsumptionForAnElectricityMeterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for List consumption for an electricity meter: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
