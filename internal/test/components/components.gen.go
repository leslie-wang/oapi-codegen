// Package components provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 DO NOT EDIT.
package components

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/deepmap/oapi-codegen/v2/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

// AdditionalPropertiesObject1 defines model for AdditionalPropertiesObject1.
type AdditionalPropertiesObject1 struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Optional *string `json:"optional,omitempty"`
}

// AdditionalPropertiesObject2 defines model for AdditionalPropertiesObject2.
type AdditionalPropertiesObject2 struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// AdditionalPropertiesObject3 defines model for AdditionalPropertiesObject3.
type AdditionalPropertiesObject3 struct {
	Name                 string                 `json:"name"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// AdditionalPropertiesObject4 defines model for AdditionalPropertiesObject4.
type AdditionalPropertiesObject4 struct {
	Inner                AdditionalPropertiesObject4_Inner `json:"inner"`
	Name                 string                            `json:"name"`
	AdditionalProperties map[string]interface{}            `json:"-"`
}

// AdditionalPropertiesObject4_Inner defines model for AdditionalPropertiesObject4.Inner.
type AdditionalPropertiesObject4_Inner struct {
	Name                 string                 `json:"name"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// AdditionalPropertiesObject5 defines model for AdditionalPropertiesObject5.
type AdditionalPropertiesObject5 map[string]interface{}

// ObjectWithJsonField defines model for ObjectWithJsonField.
type ObjectWithJsonField struct {
	Name   string          `json:"name"`
	Value1 json.RawMessage `json:"value1"`
	Value2 json.RawMessage `json:"value2,omitempty"`
}

// SchemaObject defines model for SchemaObject.
type SchemaObject struct {
	FirstName string `json:"firstName"`
	Role      string `json:"role"`
}

// ResponseObject defines model for ResponseObject.
type ResponseObject struct {
	Field SchemaObject `json:"Field"`
}

// RequestBody defines model for RequestBody.
type RequestBody struct {
	Field SchemaObject `json:"Field"`
}

// ParamsWithAddPropsParams_P1 defines parameters for ParamsWithAddProps.
type ParamsWithAddPropsParams_P1 struct {
	AdditionalProperties map[string]interface{} `json:"-"`
}

// ParamsWithAddPropsParams defines parameters for ParamsWithAddProps.
type ParamsWithAddPropsParams struct {

	// This parameter has additional properties
	P1 ParamsWithAddPropsParams_P1 `json:"p1"`

	// This parameter has an anonymous inner property which needs to be
	// turned into a proper type for additionalProperties to work
	P2 struct {
		Inner map[string]interface{} `json:"inner"`
	} `json:"p2"`
}

// BodyWithAddPropsJSONBody defines parameters for BodyWithAddProps.
type BodyWithAddPropsJSONBody struct {
	Inner                map[string]interface{} `json:"inner"`
	Name                 string                 `json:"name"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// EnsureEverythingIsReferencedRequestBody defines body for EnsureEverythingIsReferenced for application/json ContentType.
type EnsureEverythingIsReferencedJSONRequestBody RequestBody

// BodyWithAddPropsRequestBody defines body for BodyWithAddProps for application/json ContentType.
type BodyWithAddPropsJSONRequestBody BodyWithAddPropsJSONBody

// Getter for additional properties for ParamsWithAddPropsParams_P1. Returns the specified
// element and whether it was found
func (a ParamsWithAddPropsParams_P1) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for ParamsWithAddPropsParams_P1
func (a *ParamsWithAddPropsParams_P1) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for ParamsWithAddPropsParams_P1 to handle AdditionalProperties
func (a *ParamsWithAddPropsParams_P1) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error unmarshaling field %s", fieldName))
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for ParamsWithAddPropsParams_P1 to handle AdditionalProperties
func (a ParamsWithAddPropsParams_P1) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling '%s'", fieldName))
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for BodyWithAddPropsJSONBody. Returns the specified
// element and whether it was found
func (a BodyWithAddPropsJSONBody) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for BodyWithAddPropsJSONBody
func (a *BodyWithAddPropsJSONBody) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for BodyWithAddPropsJSONBody to handle AdditionalProperties
func (a *BodyWithAddPropsJSONBody) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["inner"]; found {
		err = json.Unmarshal(raw, &a.Inner)
		if err != nil {
			return errors.Wrap(err, "error reading 'inner'")
		}
		delete(object, "inner")
	}

	if raw, found := object["name"]; found {
		err = json.Unmarshal(raw, &a.Name)
		if err != nil {
			return errors.Wrap(err, "error reading 'name'")
		}
		delete(object, "name")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error unmarshaling field %s", fieldName))
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for BodyWithAddPropsJSONBody to handle AdditionalProperties
func (a BodyWithAddPropsJSONBody) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["inner"], err = json.Marshal(a.Inner)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error marshaling 'inner'"))
	}

	object["name"], err = json.Marshal(a.Name)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error marshaling 'name'"))
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling '%s'", fieldName))
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for AdditionalPropertiesObject3. Returns the specified
// element and whether it was found
func (a AdditionalPropertiesObject3) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for AdditionalPropertiesObject3
func (a *AdditionalPropertiesObject3) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for AdditionalPropertiesObject3 to handle AdditionalProperties
func (a *AdditionalPropertiesObject3) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["name"]; found {
		err = json.Unmarshal(raw, &a.Name)
		if err != nil {
			return errors.Wrap(err, "error reading 'name'")
		}
		delete(object, "name")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error unmarshaling field %s", fieldName))
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for AdditionalPropertiesObject3 to handle AdditionalProperties
func (a AdditionalPropertiesObject3) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["name"], err = json.Marshal(a.Name)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error marshaling 'name'"))
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling '%s'", fieldName))
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for AdditionalPropertiesObject4. Returns the specified
// element and whether it was found
func (a AdditionalPropertiesObject4) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for AdditionalPropertiesObject4
func (a *AdditionalPropertiesObject4) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for AdditionalPropertiesObject4 to handle AdditionalProperties
func (a *AdditionalPropertiesObject4) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["inner"]; found {
		err = json.Unmarshal(raw, &a.Inner)
		if err != nil {
			return errors.Wrap(err, "error reading 'inner'")
		}
		delete(object, "inner")
	}

	if raw, found := object["name"]; found {
		err = json.Unmarshal(raw, &a.Name)
		if err != nil {
			return errors.Wrap(err, "error reading 'name'")
		}
		delete(object, "name")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error unmarshaling field %s", fieldName))
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for AdditionalPropertiesObject4 to handle AdditionalProperties
func (a AdditionalPropertiesObject4) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["inner"], err = json.Marshal(a.Inner)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error marshaling 'inner'"))
	}

	object["name"], err = json.Marshal(a.Name)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error marshaling 'name'"))
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling '%s'", fieldName))
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for AdditionalPropertiesObject4_Inner. Returns the specified
// element and whether it was found
func (a AdditionalPropertiesObject4_Inner) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for AdditionalPropertiesObject4_Inner
func (a *AdditionalPropertiesObject4_Inner) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for AdditionalPropertiesObject4_Inner to handle AdditionalProperties
func (a *AdditionalPropertiesObject4_Inner) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["name"]; found {
		err = json.Unmarshal(raw, &a.Name)
		if err != nil {
			return errors.Wrap(err, "error reading 'name'")
		}
		delete(object, "name")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error unmarshaling field %s", fieldName))
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for AdditionalPropertiesObject4_Inner to handle AdditionalProperties
func (a AdditionalPropertiesObject4_Inner) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["name"], err = json.Marshal(a.Name)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error marshaling 'name'"))
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling '%s'", fieldName))
		}
	}
	return json.Marshal(object)
}

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A callback for modifying requests which are generated before sending over
	// the network.
	RequestEditor RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = http.DefaultClient
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditor = fn
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// EnsureEverythingIsReferenced request  with any body
	EnsureEverythingIsReferencedWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error)

	EnsureEverythingIsReferenced(ctx context.Context, body EnsureEverythingIsReferencedJSONRequestBody) (*http.Response, error)

	// ParamsWithAddProps request
	ParamsWithAddProps(ctx context.Context, params *ParamsWithAddPropsParams) (*http.Response, error)

	// BodyWithAddProps request  with any body
	BodyWithAddPropsWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error)

	BodyWithAddProps(ctx context.Context, body BodyWithAddPropsJSONRequestBody) (*http.Response, error)
}

func (c *Client) EnsureEverythingIsReferencedWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewEnsureEverythingIsReferencedRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) EnsureEverythingIsReferenced(ctx context.Context, body EnsureEverythingIsReferencedJSONRequestBody) (*http.Response, error) {
	req, err := NewEnsureEverythingIsReferencedRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) ParamsWithAddProps(ctx context.Context, params *ParamsWithAddPropsParams) (*http.Response, error) {
	req, err := NewParamsWithAddPropsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) BodyWithAddPropsWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewBodyWithAddPropsRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) BodyWithAddProps(ctx context.Context, body BodyWithAddPropsJSONRequestBody) (*http.Response, error) {
	req, err := NewBodyWithAddPropsRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

// NewEnsureEverythingIsReferencedRequest calls the generic EnsureEverythingIsReferenced builder with application/json body
func NewEnsureEverythingIsReferencedRequest(server string, body EnsureEverythingIsReferencedJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewEnsureEverythingIsReferencedRequestWithBody(server, "application/json", bodyReader)
}

// NewEnsureEverythingIsReferencedRequestWithBody generates requests for EnsureEverythingIsReferenced with any type of body
func NewEnsureEverythingIsReferencedRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/ensure-everything-is-referenced")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryUrl.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// NewParamsWithAddPropsRequest generates requests for ParamsWithAddProps
func NewParamsWithAddPropsRequest(server string, params *ParamsWithAddPropsParams) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/params_with_add_props")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	queryValues := queryUrl.Query()

	if queryFrag, err := runtime.StyleParam("simple", true, "p1", params.P1); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if queryFrag, err := runtime.StyleParam("form", true, "p2", params.P2); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	queryUrl.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewBodyWithAddPropsRequest calls the generic BodyWithAddProps builder with application/json body
func NewBodyWithAddPropsRequest(server string, body BodyWithAddPropsJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewBodyWithAddPropsRequestWithBody(server, "application/json", bodyReader)
}

// NewBodyWithAddPropsRequestWithBody generates requests for BodyWithAddProps with any type of body
func NewBodyWithAddPropsRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/params_with_add_props")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryUrl.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// EnsureEverythingIsReferenced request  with any body
	EnsureEverythingIsReferencedWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*EnsureEverythingIsReferencedResponse, error)

	EnsureEverythingIsReferencedWithResponse(ctx context.Context, body EnsureEverythingIsReferencedJSONRequestBody) (*EnsureEverythingIsReferencedResponse, error)

	// ParamsWithAddProps request
	ParamsWithAddPropsWithResponse(ctx context.Context, params *ParamsWithAddPropsParams) (*ParamsWithAddPropsResponse, error)

	// BodyWithAddProps request  with any body
	BodyWithAddPropsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*BodyWithAddPropsResponse, error)

	BodyWithAddPropsWithResponse(ctx context.Context, body BodyWithAddPropsJSONRequestBody) (*BodyWithAddPropsResponse, error)
}

type EnsureEverythingIsReferencedResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {

		// Has additional properties with schema for dictionaries
		Five *AdditionalPropertiesObject5 `json:"five,omitempty"`

		// Has anonymous field which has additional properties
		Four      *AdditionalPropertiesObject4 `json:"four,omitempty"`
		JsonField *ObjectWithJsonField         `json:"jsonField,omitempty"`

		// Has additional properties of type int
		One *AdditionalPropertiesObject1 `json:"one,omitempty"`

		// Allows any additional property
		Three *AdditionalPropertiesObject3 `json:"three,omitempty"`

		// Does not allow additional properties
		Two *AdditionalPropertiesObject2 `json:"two,omitempty"`
	}
	JSONDefault *struct {
		Field SchemaObject `json:"Field"`
	}
}

// Status returns HTTPResponse.Status
func (r EnsureEverythingIsReferencedResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r EnsureEverythingIsReferencedResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ParamsWithAddPropsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r ParamsWithAddPropsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ParamsWithAddPropsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type BodyWithAddPropsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r BodyWithAddPropsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r BodyWithAddPropsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// EnsureEverythingIsReferencedWithBodyWithResponse request with arbitrary body returning *EnsureEverythingIsReferencedResponse
func (c *ClientWithResponses) EnsureEverythingIsReferencedWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*EnsureEverythingIsReferencedResponse, error) {
	rsp, err := c.EnsureEverythingIsReferencedWithBody(ctx, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParseEnsureEverythingIsReferencedResponse(rsp)
}

func (c *ClientWithResponses) EnsureEverythingIsReferencedWithResponse(ctx context.Context, body EnsureEverythingIsReferencedJSONRequestBody) (*EnsureEverythingIsReferencedResponse, error) {
	rsp, err := c.EnsureEverythingIsReferenced(ctx, body)
	if err != nil {
		return nil, err
	}
	return ParseEnsureEverythingIsReferencedResponse(rsp)
}

// ParamsWithAddPropsWithResponse request returning *ParamsWithAddPropsResponse
func (c *ClientWithResponses) ParamsWithAddPropsWithResponse(ctx context.Context, params *ParamsWithAddPropsParams) (*ParamsWithAddPropsResponse, error) {
	rsp, err := c.ParamsWithAddProps(ctx, params)
	if err != nil {
		return nil, err
	}
	return ParseParamsWithAddPropsResponse(rsp)
}

// BodyWithAddPropsWithBodyWithResponse request with arbitrary body returning *BodyWithAddPropsResponse
func (c *ClientWithResponses) BodyWithAddPropsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*BodyWithAddPropsResponse, error) {
	rsp, err := c.BodyWithAddPropsWithBody(ctx, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParseBodyWithAddPropsResponse(rsp)
}

func (c *ClientWithResponses) BodyWithAddPropsWithResponse(ctx context.Context, body BodyWithAddPropsJSONRequestBody) (*BodyWithAddPropsResponse, error) {
	rsp, err := c.BodyWithAddProps(ctx, body)
	if err != nil {
		return nil, err
	}
	return ParseBodyWithAddPropsResponse(rsp)
}

// ParseEnsureEverythingIsReferencedResponse parses an HTTP response from a EnsureEverythingIsReferencedWithResponse call
func ParseEnsureEverythingIsReferencedResponse(rsp *http.Response) (*EnsureEverythingIsReferencedResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &EnsureEverythingIsReferencedResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {

			// Has additional properties with schema for dictionaries
			Five *AdditionalPropertiesObject5 `json:"five,omitempty"`

			// Has anonymous field which has additional properties
			Four      *AdditionalPropertiesObject4 `json:"four,omitempty"`
			JsonField *ObjectWithJsonField         `json:"jsonField,omitempty"`

			// Has additional properties of type int
			One *AdditionalPropertiesObject1 `json:"one,omitempty"`

			// Allows any additional property
			Three *AdditionalPropertiesObject3 `json:"three,omitempty"`

			// Does not allow additional properties
			Two *AdditionalPropertiesObject2 `json:"two,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest struct {
			Field SchemaObject `json:"Field"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	case true:
		// Content-type (text/plain) unsupported

	}

	return response, nil
}

// ParseParamsWithAddPropsResponse parses an HTTP response from a ParamsWithAddPropsWithResponse call
func ParseParamsWithAddPropsResponse(rsp *http.Response) (*ParamsWithAddPropsResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &ParamsWithAddPropsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	}

	return response, nil
}

// ParseBodyWithAddPropsResponse parses an HTTP response from a BodyWithAddPropsWithResponse call
func ParseBodyWithAddPropsResponse(rsp *http.Response) (*BodyWithAddPropsResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &BodyWithAddPropsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /ensure-everything-is-referenced)
	EnsureEverythingIsReferenced(ctx echo.Context) error

	// (GET /params_with_add_props)
	ParamsWithAddProps(ctx echo.Context, params ParamsWithAddPropsParams) error

	// (POST /params_with_add_props)
	BodyWithAddProps(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// EnsureEverythingIsReferenced converts echo context to params.
func (w *ServerInterfaceWrapper) EnsureEverythingIsReferenced(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.EnsureEverythingIsReferenced(ctx)
	return err
}

// ParamsWithAddProps converts echo context to params.
func (w *ServerInterfaceWrapper) ParamsWithAddProps(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params ParamsWithAddPropsParams
	// ------------- Required query parameter "p1" -------------

	err = runtime.BindQueryParameter("simple", true, true, "p1", ctx.QueryParams(), &params.P1)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter p1: %s", err))
	}

	// ------------- Required query parameter "p2" -------------

	err = runtime.BindQueryParameter("form", true, true, "p2", ctx.QueryParams(), &params.P2)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter p2: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ParamsWithAddProps(ctx, params)
	return err
}

// BodyWithAddProps converts echo context to params.
func (w *ServerInterfaceWrapper) BodyWithAddProps(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.BodyWithAddProps(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/ensure-everything-is-referenced", wrapper.EnsureEverythingIsReferenced)
	router.GET(baseURL+"/params_with_add_props", wrapper.ParamsWithAddProps)
	router.POST(baseURL+"/params_with_add_props", wrapper.BodyWithAddProps)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9RXTW/jNhD9KwTboxLHzvaim4tu0RRoG+wG6GFjBLQ4ipjKQy05sldY6L8XQ8nflNdx",
	"9rK5RJY5X49v3oy/yswuKouA5GX6VTr4XIOnX602EF582Lxo+GNmkQCJH1VVlSZTZCyOXrxFfuezAhaK",
	"nypnK3DUe/ndQKn54WcHuUzlT6Nt2FFn5Ecfw/9/5i+QkWzbJCRjHGiZfuo9zPg1wRcaVaUyByGpqUCm",
	"0pMz+Cxb/mMfvrLo18V0H/oYP1o9idTgM2cqzlGmciq8WVQliHWRwm6D9Vmwo6nWhk1Ueb+poktrHAqP",
	"fL0T3yDBMzh5FP4P5cXWVmwREjYXbCwMkkwOoDM67hvVAiJVJ9JWXYAYJPuYBhcJR5gl66NrRJITKEyG",
	"UchV6eGw8N8seIGWhCpLu4pj8Na6v1Npt8OlkauPKptyQV4obCJVNUc1vSL316X97nVpByaixWZhay9y",
	"bi2xKkxWiGKIo8f3gwjuW2G/a/nnmXd5JZeg+Mup7j5fuc7v+5WhQnRORG6d0CYLh1wH+FHqXYR/DRV/",
	"eosbUT0L5UQuVVlDULDcuoUimcqg28nA0ckZR+Nt10eKob8H1VHuuXGe/h4qwNnyDAKEU8mOq1kYBQZz",
	"y8alyQA9bJGSf909sHcyxO7lA3gSH8EtA42W4Hx3jePrm+ubTmABVWVkKm+vb67H3BmKipD/CNDXDq5g",
	"Ca6hwuDzlfFXDnJwgBmE23qGUPg+Rx4K4wWgrqxBEvDFePLCW0GFIrFlnMgUijmIzIEi0MKgoML4R/QV",
	"ZEKhDjI7B1G5GkE/8o0xvmFK32mZyvchwfeb/O78h212yc4+0wyRfm/lGe3uO4frw+Tm5g07Q26W8K3G",
	"O9XLbSJzW7vLXbxjFy+7jXbKT6w3mSz4hiLGgZeFgzf4uA0+VvZyD5PQYged3K9XuapLGmZKT4bRwSLZ",
	"WY8q5dTCP7EKPimtn/j+/WCLTAW3WaeZwRIInA+kV2JuddOPsF4L4pIb6Yj7kAVf3FTr+5ACd/Q6gEw/",
	"RZt1c2J4ZoZgvKXKzzW4Zj2UUlmN5a5mdbNy2wenJuqRoHpqgmx1q61sk3OyxZ3xHwbmZmfpQUQA7QVZ",
	"MYdHpNphEBuyQvUnu4WVh1YsW7ZcWfffMAKTkwi8atWITIpDssZWhFnbzvYEC+uybBNZWR9hXxjiote+",
	"XbqxuimD/K02DjKKApIwTx/xJPBM7JhthLMstweMdRf+8Dx/fTv3GnaW9Qt3uPX2vr6n9pAr7fHFtW37",
	"fwAAAP//Xv36ip0PAAA=",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
