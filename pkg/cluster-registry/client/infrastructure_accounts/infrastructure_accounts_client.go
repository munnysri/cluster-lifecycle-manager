// Code generated by go-swagger; DO NOT EDIT.

package infrastructure_accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new infrastructure accounts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for infrastructure accounts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateInfrastructureAccount creates infrastructure account

Creates a new infrastructure account

*/
func (a *Client) CreateInfrastructureAccount(params *CreateInfrastructureAccountParams, authInfo runtime.ClientAuthInfoWriter) (*CreateInfrastructureAccountCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateInfrastructureAccountParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createInfrastructureAccount",
		Method:             "POST",
		PathPattern:        "/infrastructure-accounts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateInfrastructureAccountReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateInfrastructureAccountCreated), nil

}

/*
GetInfrastructureAccount gets single infrastructure account

Read information regarding the infrastructure account.

*/
func (a *Client) GetInfrastructureAccount(params *GetInfrastructureAccountParams, authInfo runtime.ClientAuthInfoWriter) (*GetInfrastructureAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInfrastructureAccountParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getInfrastructureAccount",
		Method:             "GET",
		PathPattern:        "/infrastructure-accounts/{account_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInfrastructureAccountReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetInfrastructureAccountOK), nil

}

/*
ListInfrastructureAccounts lists all registered infrastructure accounts
*/
func (a *Client) ListInfrastructureAccounts(params *ListInfrastructureAccountsParams, authInfo runtime.ClientAuthInfoWriter) (*ListInfrastructureAccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListInfrastructureAccountsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listInfrastructureAccounts",
		Method:             "GET",
		PathPattern:        "/infrastructure-accounts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListInfrastructureAccountsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListInfrastructureAccountsOK), nil

}

/*
UpdateInfrastructureAccount updates infrastructure account

update an infrastructure account.
*/
func (a *Client) UpdateInfrastructureAccount(params *UpdateInfrastructureAccountParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateInfrastructureAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateInfrastructureAccountParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateInfrastructureAccount",
		Method:             "PATCH",
		PathPattern:        "/infrastructure-accounts/{account_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateInfrastructureAccountReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateInfrastructureAccountOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
