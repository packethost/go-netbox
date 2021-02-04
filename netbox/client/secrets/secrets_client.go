// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package secrets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new secrets API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for secrets API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	SecretsGenerateRsaKeyPairList(params *SecretsGenerateRsaKeyPairListParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsGenerateRsaKeyPairListOK, error)

	SecretsGetSessionKeyCreate(params *SecretsGetSessionKeyCreateParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsGetSessionKeyCreateCreated, error)

	SecretsSecretRolesBulkDelete(params *SecretsSecretRolesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretRolesBulkDeleteNoContent, error)

	SecretsSecretRolesBulkPartialUpdate(params *SecretsSecretRolesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretRolesBulkPartialUpdateOK, error)

	SecretsSecretRolesBulkUpdate(params *SecretsSecretRolesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretRolesBulkUpdateOK, error)

	SecretsSecretRolesCreate(params *SecretsSecretRolesCreateParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretRolesCreateCreated, error)

	SecretsSecretRolesDelete(params *SecretsSecretRolesDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretRolesDeleteNoContent, error)

	SecretsSecretRolesList(params *SecretsSecretRolesListParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretRolesListOK, error)

	SecretsSecretRolesPartialUpdate(params *SecretsSecretRolesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretRolesPartialUpdateOK, error)

	SecretsSecretRolesRead(params *SecretsSecretRolesReadParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretRolesReadOK, error)

	SecretsSecretRolesUpdate(params *SecretsSecretRolesUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretRolesUpdateOK, error)

	SecretsSecretsBulkDelete(params *SecretsSecretsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretsBulkDeleteNoContent, error)

	SecretsSecretsBulkPartialUpdate(params *SecretsSecretsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretsBulkPartialUpdateOK, error)

	SecretsSecretsBulkUpdate(params *SecretsSecretsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretsBulkUpdateOK, error)

	SecretsSecretsCreate(params *SecretsSecretsCreateParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretsCreateCreated, error)

	SecretsSecretsDelete(params *SecretsSecretsDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretsDeleteNoContent, error)

	SecretsSecretsList(params *SecretsSecretsListParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretsListOK, error)

	SecretsSecretsPartialUpdate(params *SecretsSecretsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretsPartialUpdateOK, error)

	SecretsSecretsRead(params *SecretsSecretsReadParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretsReadOK, error)

	SecretsSecretsUpdate(params *SecretsSecretsUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretsUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  SecretsGenerateRsaKeyPairList this endpoint can be used to generate a new r s a key pair the keys are returned in p e m format

  {
        "public_key": "<public key>",
        "private_key": "<private key>"
    }
*/
func (a *Client) SecretsGenerateRsaKeyPairList(params *SecretsGenerateRsaKeyPairListParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsGenerateRsaKeyPairListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecretsGenerateRsaKeyPairListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "secrets_generate-rsa-key-pair_list",
		Method:             "GET",
		PathPattern:        "/secrets/generate-rsa-key-pair/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SecretsGenerateRsaKeyPairListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SecretsGenerateRsaKeyPairListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for secrets_generate-rsa-key-pair_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SecretsGetSessionKeyCreate Retrieve a temporary session key to use for encrypting and decrypting secrets via the API. The user's private RSA
key is POSTed with the name `private_key`. An example:

    curl -v -X POST -H "Authorization: Token <token>" -H "Accept: application/json; indent=4" \
    --data-urlencode "private_key@<filename>" https://netbox/api/secrets/get-session-key/

This request will yield a base64-encoded session key to be included in an `X-Session-Key` header in future requests:

    {
        "session_key": "+8t4SI6XikgVmB5+/urhozx9O5qCQANyOk1MNe6taRf="
    }

This endpoint accepts one optional parameter: `preserve_key`. If True and a session key exists, the existing session
key will be returned instead of a new one.
*/
func (a *Client) SecretsGetSessionKeyCreate(params *SecretsGetSessionKeyCreateParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsGetSessionKeyCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecretsGetSessionKeyCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "secrets_get-session-key_create",
		Method:             "POST",
		PathPattern:        "/secrets/get-session-key/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SecretsGetSessionKeyCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SecretsGetSessionKeyCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for secrets_get-session-key_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SecretsSecretRolesBulkDelete secrets secret roles bulk delete API
*/
func (a *Client) SecretsSecretRolesBulkDelete(params *SecretsSecretRolesBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretRolesBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecretsSecretRolesBulkDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "secrets_secret-roles_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/secrets/secret-roles/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SecretsSecretRolesBulkDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SecretsSecretRolesBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for secrets_secret-roles_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SecretsSecretRolesBulkPartialUpdate secrets secret roles bulk partial update API
*/
func (a *Client) SecretsSecretRolesBulkPartialUpdate(params *SecretsSecretRolesBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretRolesBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecretsSecretRolesBulkPartialUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "secrets_secret-roles_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/secrets/secret-roles/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SecretsSecretRolesBulkPartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SecretsSecretRolesBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for secrets_secret-roles_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SecretsSecretRolesBulkUpdate secrets secret roles bulk update API
*/
func (a *Client) SecretsSecretRolesBulkUpdate(params *SecretsSecretRolesBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretRolesBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecretsSecretRolesBulkUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "secrets_secret-roles_bulk_update",
		Method:             "PUT",
		PathPattern:        "/secrets/secret-roles/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SecretsSecretRolesBulkUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SecretsSecretRolesBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for secrets_secret-roles_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SecretsSecretRolesCreate secrets secret roles create API
*/
func (a *Client) SecretsSecretRolesCreate(params *SecretsSecretRolesCreateParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretRolesCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecretsSecretRolesCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "secrets_secret-roles_create",
		Method:             "POST",
		PathPattern:        "/secrets/secret-roles/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SecretsSecretRolesCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SecretsSecretRolesCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for secrets_secret-roles_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SecretsSecretRolesDelete secrets secret roles delete API
*/
func (a *Client) SecretsSecretRolesDelete(params *SecretsSecretRolesDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretRolesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecretsSecretRolesDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "secrets_secret-roles_delete",
		Method:             "DELETE",
		PathPattern:        "/secrets/secret-roles/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SecretsSecretRolesDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SecretsSecretRolesDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for secrets_secret-roles_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SecretsSecretRolesList secrets secret roles list API
*/
func (a *Client) SecretsSecretRolesList(params *SecretsSecretRolesListParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretRolesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecretsSecretRolesListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "secrets_secret-roles_list",
		Method:             "GET",
		PathPattern:        "/secrets/secret-roles/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SecretsSecretRolesListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SecretsSecretRolesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for secrets_secret-roles_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SecretsSecretRolesPartialUpdate secrets secret roles partial update API
*/
func (a *Client) SecretsSecretRolesPartialUpdate(params *SecretsSecretRolesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretRolesPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecretsSecretRolesPartialUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "secrets_secret-roles_partial_update",
		Method:             "PATCH",
		PathPattern:        "/secrets/secret-roles/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SecretsSecretRolesPartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SecretsSecretRolesPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for secrets_secret-roles_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SecretsSecretRolesRead secrets secret roles read API
*/
func (a *Client) SecretsSecretRolesRead(params *SecretsSecretRolesReadParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretRolesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecretsSecretRolesReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "secrets_secret-roles_read",
		Method:             "GET",
		PathPattern:        "/secrets/secret-roles/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SecretsSecretRolesReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SecretsSecretRolesReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for secrets_secret-roles_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SecretsSecretRolesUpdate secrets secret roles update API
*/
func (a *Client) SecretsSecretRolesUpdate(params *SecretsSecretRolesUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretRolesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecretsSecretRolesUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "secrets_secret-roles_update",
		Method:             "PUT",
		PathPattern:        "/secrets/secret-roles/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SecretsSecretRolesUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SecretsSecretRolesUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for secrets_secret-roles_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SecretsSecretsBulkDelete secrets secrets bulk delete API
*/
func (a *Client) SecretsSecretsBulkDelete(params *SecretsSecretsBulkDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretsBulkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecretsSecretsBulkDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "secrets_secrets_bulk_delete",
		Method:             "DELETE",
		PathPattern:        "/secrets/secrets/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SecretsSecretsBulkDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SecretsSecretsBulkDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for secrets_secrets_bulk_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SecretsSecretsBulkPartialUpdate secrets secrets bulk partial update API
*/
func (a *Client) SecretsSecretsBulkPartialUpdate(params *SecretsSecretsBulkPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretsBulkPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecretsSecretsBulkPartialUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "secrets_secrets_bulk_partial_update",
		Method:             "PATCH",
		PathPattern:        "/secrets/secrets/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SecretsSecretsBulkPartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SecretsSecretsBulkPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for secrets_secrets_bulk_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SecretsSecretsBulkUpdate secrets secrets bulk update API
*/
func (a *Client) SecretsSecretsBulkUpdate(params *SecretsSecretsBulkUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretsBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecretsSecretsBulkUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "secrets_secrets_bulk_update",
		Method:             "PUT",
		PathPattern:        "/secrets/secrets/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SecretsSecretsBulkUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SecretsSecretsBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for secrets_secrets_bulk_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SecretsSecretsCreate secrets secrets create API
*/
func (a *Client) SecretsSecretsCreate(params *SecretsSecretsCreateParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecretsSecretsCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "secrets_secrets_create",
		Method:             "POST",
		PathPattern:        "/secrets/secrets/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SecretsSecretsCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SecretsSecretsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for secrets_secrets_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SecretsSecretsDelete secrets secrets delete API
*/
func (a *Client) SecretsSecretsDelete(params *SecretsSecretsDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecretsSecretsDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "secrets_secrets_delete",
		Method:             "DELETE",
		PathPattern:        "/secrets/secrets/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SecretsSecretsDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SecretsSecretsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for secrets_secrets_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SecretsSecretsList secrets secrets list API
*/
func (a *Client) SecretsSecretsList(params *SecretsSecretsListParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecretsSecretsListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "secrets_secrets_list",
		Method:             "GET",
		PathPattern:        "/secrets/secrets/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SecretsSecretsListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SecretsSecretsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for secrets_secrets_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SecretsSecretsPartialUpdate secrets secrets partial update API
*/
func (a *Client) SecretsSecretsPartialUpdate(params *SecretsSecretsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecretsSecretsPartialUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "secrets_secrets_partial_update",
		Method:             "PATCH",
		PathPattern:        "/secrets/secrets/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SecretsSecretsPartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SecretsSecretsPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for secrets_secrets_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SecretsSecretsRead secrets secrets read API
*/
func (a *Client) SecretsSecretsRead(params *SecretsSecretsReadParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecretsSecretsReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "secrets_secrets_read",
		Method:             "GET",
		PathPattern:        "/secrets/secrets/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SecretsSecretsReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SecretsSecretsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for secrets_secrets_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SecretsSecretsUpdate secrets secrets update API
*/
func (a *Client) SecretsSecretsUpdate(params *SecretsSecretsUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecretsSecretsUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "secrets_secrets_update",
		Method:             "PUT",
		PathPattern:        "/secrets/secrets/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SecretsSecretsUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SecretsSecretsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for secrets_secrets_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
