// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package personalization

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/algolia/algoliasearch-client-go/v4/algolia/utils"
)

type Option struct {
	optionType string
	name       string
	value      string
}

func QueryParamOption(name string, val any) Option {
	return Option{
		optionType: "query",
		name:       queryParameterToString(name),
		value:      queryParameterToString(val),
	}
}

func HeaderParamOption(name string, val any) Option {
	return Option{
		optionType: "header",
		name:       name,
		value:      parameterToString(val),
	}
}

func (r *ApiCustomDeleteRequest) UnmarshalJSON(b []byte) error {
	req := map[string]json.RawMessage{}
	err := json.Unmarshal(b, &req)
	if err != nil {
		return fmt.Errorf("cannot unmarshal request: %w", err)
	}
	if v, ok := req["path"]; ok {
		err = json.Unmarshal(v, &r.path)
		if err != nil {
			err = json.Unmarshal(b, &r.path)
			if err != nil {
				return fmt.Errorf("cannot unmarshal path: %w", err)
			}
		}
	}
	if v, ok := req["parameters"]; ok {
		err = json.Unmarshal(v, &r.parameters)
		if err != nil {
			err = json.Unmarshal(b, &r.parameters)
			if err != nil {
				return fmt.Errorf("cannot unmarshal parameters: %w", err)
			}
		}
	}

	return nil
}

// ApiCustomDeleteRequest represents the request with all the parameters for the API call.
type ApiCustomDeleteRequest struct {
	path       string
	parameters map[string]interface{}
}

// NewApiCustomDeleteRequest creates an instance of the ApiCustomDeleteRequest to be used for the API call.
func (c *APIClient) NewApiCustomDeleteRequest(path string) ApiCustomDeleteRequest {
	return ApiCustomDeleteRequest{
		path: path,
	}
}

// WithParameters adds the parameters to the ApiCustomDeleteRequest and returns the request for chaining.
func (r ApiCustomDeleteRequest) WithParameters(parameters map[string]interface{}) ApiCustomDeleteRequest {
	r.parameters = parameters
	return r
}

/*
CustomDelete Wraps CustomDeleteWithContext using context.Background.

This method allow you to send requests to the Algolia REST API.

Request can be constructed by NewApiCustomDeleteRequest with parameters below.

	@param path string - Path of the endpoint, anything after \"/1\" must be specified.
	@param parameters map[string]interface{} - Query parameters to apply to the current query.
	@return map[string]interface{}
*/
func (c *APIClient) CustomDelete(r ApiCustomDeleteRequest, opts ...Option) (map[string]interface{}, error) {
	return c.CustomDeleteWithContext(context.Background(), r, opts...)
}

/*
CustomDelete

This method allow you to send requests to the Algolia REST API.

Request can be constructed by NewApiCustomDeleteRequest with parameters below.

	@param path string - Path of the endpoint, anything after \"/1\" must be specified.
	@param parameters map[string]interface{} - Query parameters to apply to the current query.
	@return map[string]interface{}
*/
func (c *APIClient) CustomDeleteWithContext(ctx context.Context, r ApiCustomDeleteRequest, opts ...Option) (map[string]interface{}, error) {
	var (
		postBody    any
		returnValue map[string]interface{}
	)

	requestPath := "/{path}"
	requestPath = strings.ReplaceAll(requestPath, "{path}", parameterToString(r.path))

	headers := make(map[string]string)
	queryParams := url.Values{}
	if r.path == "" {
		return returnValue, reportError("Parameter `path` is required when calling `CustomDelete`.")
	}

	if !utils.IsNilOrEmpty(r.parameters) {
		for k, v := range r.parameters {
			queryParams.Set(k, queryParameterToString(v))
		}
	}

	// optional params if any
	for _, opt := range opts {
		switch opt.optionType {
		case "query":
			queryParams.Set(opt.name, opt.value)
		case "header":
			headers[opt.name] = opt.value
		}
	}

	req, err := c.prepareRequest(ctx, requestPath, http.MethodDelete, postBody, headers, queryParams)
	if err != nil {
		return returnValue, err
	}

	res, resBody, err := c.callAPI(req, false)
	if err != nil {
		return returnValue, err
	}
	if res == nil {
		return returnValue, reportError("res is nil")
	}

	if res.StatusCode >= 300 {
		newErr := &APIError{
			Message: string(resBody),
			Status:  res.StatusCode,
		}

		var v ErrorBase
		err = c.decode(&v, resBody)
		if err != nil {
			newErr.Message = err.Error()
			return returnValue, newErr
		}

		return returnValue, newErr
	}

	err = c.decode(&returnValue, resBody)
	if err != nil {
		return returnValue, reportError("cannot decode result: %w", err)
	}

	return returnValue, nil
}

func (r *ApiCustomGetRequest) UnmarshalJSON(b []byte) error {
	req := map[string]json.RawMessage{}
	err := json.Unmarshal(b, &req)
	if err != nil {
		return fmt.Errorf("cannot unmarshal request: %w", err)
	}
	if v, ok := req["path"]; ok {
		err = json.Unmarshal(v, &r.path)
		if err != nil {
			err = json.Unmarshal(b, &r.path)
			if err != nil {
				return fmt.Errorf("cannot unmarshal path: %w", err)
			}
		}
	}
	if v, ok := req["parameters"]; ok {
		err = json.Unmarshal(v, &r.parameters)
		if err != nil {
			err = json.Unmarshal(b, &r.parameters)
			if err != nil {
				return fmt.Errorf("cannot unmarshal parameters: %w", err)
			}
		}
	}

	return nil
}

// ApiCustomGetRequest represents the request with all the parameters for the API call.
type ApiCustomGetRequest struct {
	path       string
	parameters map[string]interface{}
}

// NewApiCustomGetRequest creates an instance of the ApiCustomGetRequest to be used for the API call.
func (c *APIClient) NewApiCustomGetRequest(path string) ApiCustomGetRequest {
	return ApiCustomGetRequest{
		path: path,
	}
}

// WithParameters adds the parameters to the ApiCustomGetRequest and returns the request for chaining.
func (r ApiCustomGetRequest) WithParameters(parameters map[string]interface{}) ApiCustomGetRequest {
	r.parameters = parameters
	return r
}

/*
CustomGet Wraps CustomGetWithContext using context.Background.

This method allow you to send requests to the Algolia REST API.

Request can be constructed by NewApiCustomGetRequest with parameters below.

	@param path string - Path of the endpoint, anything after \"/1\" must be specified.
	@param parameters map[string]interface{} - Query parameters to apply to the current query.
	@return map[string]interface{}
*/
func (c *APIClient) CustomGet(r ApiCustomGetRequest, opts ...Option) (map[string]interface{}, error) {
	return c.CustomGetWithContext(context.Background(), r, opts...)
}

/*
CustomGet

This method allow you to send requests to the Algolia REST API.

Request can be constructed by NewApiCustomGetRequest with parameters below.

	@param path string - Path of the endpoint, anything after \"/1\" must be specified.
	@param parameters map[string]interface{} - Query parameters to apply to the current query.
	@return map[string]interface{}
*/
func (c *APIClient) CustomGetWithContext(ctx context.Context, r ApiCustomGetRequest, opts ...Option) (map[string]interface{}, error) {
	var (
		postBody    any
		returnValue map[string]interface{}
	)

	requestPath := "/{path}"
	requestPath = strings.ReplaceAll(requestPath, "{path}", parameterToString(r.path))

	headers := make(map[string]string)
	queryParams := url.Values{}
	if r.path == "" {
		return returnValue, reportError("Parameter `path` is required when calling `CustomGet`.")
	}

	if !utils.IsNilOrEmpty(r.parameters) {
		for k, v := range r.parameters {
			queryParams.Set(k, queryParameterToString(v))
		}
	}

	// optional params if any
	for _, opt := range opts {
		switch opt.optionType {
		case "query":
			queryParams.Set(opt.name, opt.value)
		case "header":
			headers[opt.name] = opt.value
		}
	}

	req, err := c.prepareRequest(ctx, requestPath, http.MethodGet, postBody, headers, queryParams)
	if err != nil {
		return returnValue, err
	}

	res, resBody, err := c.callAPI(req, false)
	if err != nil {
		return returnValue, err
	}
	if res == nil {
		return returnValue, reportError("res is nil")
	}

	if res.StatusCode >= 300 {
		newErr := &APIError{
			Message: string(resBody),
			Status:  res.StatusCode,
		}

		var v ErrorBase
		err = c.decode(&v, resBody)
		if err != nil {
			newErr.Message = err.Error()
			return returnValue, newErr
		}

		return returnValue, newErr
	}

	err = c.decode(&returnValue, resBody)
	if err != nil {
		return returnValue, reportError("cannot decode result: %w", err)
	}

	return returnValue, nil
}

func (r *ApiCustomPostRequest) UnmarshalJSON(b []byte) error {
	req := map[string]json.RawMessage{}
	err := json.Unmarshal(b, &req)
	if err != nil {
		return fmt.Errorf("cannot unmarshal request: %w", err)
	}
	if v, ok := req["path"]; ok {
		err = json.Unmarshal(v, &r.path)
		if err != nil {
			err = json.Unmarshal(b, &r.path)
			if err != nil {
				return fmt.Errorf("cannot unmarshal path: %w", err)
			}
		}
	}
	if v, ok := req["parameters"]; ok {
		err = json.Unmarshal(v, &r.parameters)
		if err != nil {
			err = json.Unmarshal(b, &r.parameters)
			if err != nil {
				return fmt.Errorf("cannot unmarshal parameters: %w", err)
			}
		}
	}
	if v, ok := req["body"]; ok {
		err = json.Unmarshal(v, &r.body)
		if err != nil {
			err = json.Unmarshal(b, &r.body)
			if err != nil {
				return fmt.Errorf("cannot unmarshal body: %w", err)
			}
		}
	}

	return nil
}

// ApiCustomPostRequest represents the request with all the parameters for the API call.
type ApiCustomPostRequest struct {
	path       string
	parameters map[string]interface{}
	body       map[string]interface{}
}

// NewApiCustomPostRequest creates an instance of the ApiCustomPostRequest to be used for the API call.
func (c *APIClient) NewApiCustomPostRequest(path string) ApiCustomPostRequest {
	return ApiCustomPostRequest{
		path: path,
	}
}

// WithParameters adds the parameters to the ApiCustomPostRequest and returns the request for chaining.
func (r ApiCustomPostRequest) WithParameters(parameters map[string]interface{}) ApiCustomPostRequest {
	r.parameters = parameters
	return r
}

// WithBody adds the body to the ApiCustomPostRequest and returns the request for chaining.
func (r ApiCustomPostRequest) WithBody(body map[string]interface{}) ApiCustomPostRequest {
	r.body = body
	return r
}

/*
CustomPost Wraps CustomPostWithContext using context.Background.

This method allow you to send requests to the Algolia REST API.

Request can be constructed by NewApiCustomPostRequest with parameters below.

	@param path string - Path of the endpoint, anything after \"/1\" must be specified.
	@param parameters map[string]interface{} - Query parameters to apply to the current query.
	@param body map[string]interface{} - Parameters to send with the custom request.
	@return map[string]interface{}
*/
func (c *APIClient) CustomPost(r ApiCustomPostRequest, opts ...Option) (map[string]interface{}, error) {
	return c.CustomPostWithContext(context.Background(), r, opts...)
}

/*
CustomPost

This method allow you to send requests to the Algolia REST API.

Request can be constructed by NewApiCustomPostRequest with parameters below.

	@param path string - Path of the endpoint, anything after \"/1\" must be specified.
	@param parameters map[string]interface{} - Query parameters to apply to the current query.
	@param body map[string]interface{} - Parameters to send with the custom request.
	@return map[string]interface{}
*/
func (c *APIClient) CustomPostWithContext(ctx context.Context, r ApiCustomPostRequest, opts ...Option) (map[string]interface{}, error) {
	var (
		postBody    any
		returnValue map[string]interface{}
	)

	requestPath := "/{path}"
	requestPath = strings.ReplaceAll(requestPath, "{path}", parameterToString(r.path))

	headers := make(map[string]string)
	queryParams := url.Values{}
	if r.path == "" {
		return returnValue, reportError("Parameter `path` is required when calling `CustomPost`.")
	}

	if !utils.IsNilOrEmpty(r.parameters) {
		for k, v := range r.parameters {
			queryParams.Set(k, queryParameterToString(v))
		}
	}

	// optional params if any
	for _, opt := range opts {
		switch opt.optionType {
		case "query":
			queryParams.Set(opt.name, opt.value)
		case "header":
			headers[opt.name] = opt.value
		}
	}

	// body params
	if utils.IsNilOrEmpty(r.body) {
		postBody = "{}"
	} else {
		postBody = r.body
	}
	req, err := c.prepareRequest(ctx, requestPath, http.MethodPost, postBody, headers, queryParams)
	if err != nil {
		return returnValue, err
	}

	res, resBody, err := c.callAPI(req, false)
	if err != nil {
		return returnValue, err
	}
	if res == nil {
		return returnValue, reportError("res is nil")
	}

	if res.StatusCode >= 300 {
		newErr := &APIError{
			Message: string(resBody),
			Status:  res.StatusCode,
		}

		var v ErrorBase
		err = c.decode(&v, resBody)
		if err != nil {
			newErr.Message = err.Error()
			return returnValue, newErr
		}

		return returnValue, newErr
	}

	err = c.decode(&returnValue, resBody)
	if err != nil {
		return returnValue, reportError("cannot decode result: %w", err)
	}

	return returnValue, nil
}

func (r *ApiCustomPutRequest) UnmarshalJSON(b []byte) error {
	req := map[string]json.RawMessage{}
	err := json.Unmarshal(b, &req)
	if err != nil {
		return fmt.Errorf("cannot unmarshal request: %w", err)
	}
	if v, ok := req["path"]; ok {
		err = json.Unmarshal(v, &r.path)
		if err != nil {
			err = json.Unmarshal(b, &r.path)
			if err != nil {
				return fmt.Errorf("cannot unmarshal path: %w", err)
			}
		}
	}
	if v, ok := req["parameters"]; ok {
		err = json.Unmarshal(v, &r.parameters)
		if err != nil {
			err = json.Unmarshal(b, &r.parameters)
			if err != nil {
				return fmt.Errorf("cannot unmarshal parameters: %w", err)
			}
		}
	}
	if v, ok := req["body"]; ok {
		err = json.Unmarshal(v, &r.body)
		if err != nil {
			err = json.Unmarshal(b, &r.body)
			if err != nil {
				return fmt.Errorf("cannot unmarshal body: %w", err)
			}
		}
	}

	return nil
}

// ApiCustomPutRequest represents the request with all the parameters for the API call.
type ApiCustomPutRequest struct {
	path       string
	parameters map[string]interface{}
	body       map[string]interface{}
}

// NewApiCustomPutRequest creates an instance of the ApiCustomPutRequest to be used for the API call.
func (c *APIClient) NewApiCustomPutRequest(path string) ApiCustomPutRequest {
	return ApiCustomPutRequest{
		path: path,
	}
}

// WithParameters adds the parameters to the ApiCustomPutRequest and returns the request for chaining.
func (r ApiCustomPutRequest) WithParameters(parameters map[string]interface{}) ApiCustomPutRequest {
	r.parameters = parameters
	return r
}

// WithBody adds the body to the ApiCustomPutRequest and returns the request for chaining.
func (r ApiCustomPutRequest) WithBody(body map[string]interface{}) ApiCustomPutRequest {
	r.body = body
	return r
}

/*
CustomPut Wraps CustomPutWithContext using context.Background.

This method allow you to send requests to the Algolia REST API.

Request can be constructed by NewApiCustomPutRequest with parameters below.

	@param path string - Path of the endpoint, anything after \"/1\" must be specified.
	@param parameters map[string]interface{} - Query parameters to apply to the current query.
	@param body map[string]interface{} - Parameters to send with the custom request.
	@return map[string]interface{}
*/
func (c *APIClient) CustomPut(r ApiCustomPutRequest, opts ...Option) (map[string]interface{}, error) {
	return c.CustomPutWithContext(context.Background(), r, opts...)
}

/*
CustomPut

This method allow you to send requests to the Algolia REST API.

Request can be constructed by NewApiCustomPutRequest with parameters below.

	@param path string - Path of the endpoint, anything after \"/1\" must be specified.
	@param parameters map[string]interface{} - Query parameters to apply to the current query.
	@param body map[string]interface{} - Parameters to send with the custom request.
	@return map[string]interface{}
*/
func (c *APIClient) CustomPutWithContext(ctx context.Context, r ApiCustomPutRequest, opts ...Option) (map[string]interface{}, error) {
	var (
		postBody    any
		returnValue map[string]interface{}
	)

	requestPath := "/{path}"
	requestPath = strings.ReplaceAll(requestPath, "{path}", parameterToString(r.path))

	headers := make(map[string]string)
	queryParams := url.Values{}
	if r.path == "" {
		return returnValue, reportError("Parameter `path` is required when calling `CustomPut`.")
	}

	if !utils.IsNilOrEmpty(r.parameters) {
		for k, v := range r.parameters {
			queryParams.Set(k, queryParameterToString(v))
		}
	}

	// optional params if any
	for _, opt := range opts {
		switch opt.optionType {
		case "query":
			queryParams.Set(opt.name, opt.value)
		case "header":
			headers[opt.name] = opt.value
		}
	}

	// body params
	if utils.IsNilOrEmpty(r.body) {
		postBody = "{}"
	} else {
		postBody = r.body
	}
	req, err := c.prepareRequest(ctx, requestPath, http.MethodPut, postBody, headers, queryParams)
	if err != nil {
		return returnValue, err
	}

	res, resBody, err := c.callAPI(req, false)
	if err != nil {
		return returnValue, err
	}
	if res == nil {
		return returnValue, reportError("res is nil")
	}

	if res.StatusCode >= 300 {
		newErr := &APIError{
			Message: string(resBody),
			Status:  res.StatusCode,
		}

		var v ErrorBase
		err = c.decode(&v, resBody)
		if err != nil {
			newErr.Message = err.Error()
			return returnValue, newErr
		}

		return returnValue, newErr
	}

	err = c.decode(&returnValue, resBody)
	if err != nil {
		return returnValue, reportError("cannot decode result: %w", err)
	}

	return returnValue, nil
}

func (r *ApiDeleteUserProfileRequest) UnmarshalJSON(b []byte) error {
	req := map[string]json.RawMessage{}
	err := json.Unmarshal(b, &req)
	if err != nil {
		return fmt.Errorf("cannot unmarshal request: %w", err)
	}
	if v, ok := req["userToken"]; ok {
		err = json.Unmarshal(v, &r.userToken)
		if err != nil {
			err = json.Unmarshal(b, &r.userToken)
			if err != nil {
				return fmt.Errorf("cannot unmarshal userToken: %w", err)
			}
		}
	}

	return nil
}

// ApiDeleteUserProfileRequest represents the request with all the parameters for the API call.
type ApiDeleteUserProfileRequest struct {
	userToken string
}

// NewApiDeleteUserProfileRequest creates an instance of the ApiDeleteUserProfileRequest to be used for the API call.
func (c *APIClient) NewApiDeleteUserProfileRequest(userToken string) ApiDeleteUserProfileRequest {
	return ApiDeleteUserProfileRequest{
		userToken: userToken,
	}
}

/*
DeleteUserProfile Wraps DeleteUserProfileWithContext using context.Background.

Deletes a user profile.

The response includes a date and time when the user profile can safely be considered deleted.

Required API Key ACLs:
  - recommendation

Request can be constructed by NewApiDeleteUserProfileRequest with parameters below.

	@param userToken string - Unique identifier representing a user for which to fetch the personalization profile.
	@return DeleteUserProfileResponse
*/
func (c *APIClient) DeleteUserProfile(r ApiDeleteUserProfileRequest, opts ...Option) (*DeleteUserProfileResponse, error) {
	return c.DeleteUserProfileWithContext(context.Background(), r, opts...)
}

/*
DeleteUserProfile

Deletes a user profile.

The response includes a date and time when the user profile can safely be considered deleted.

Required API Key ACLs:
  - recommendation

Request can be constructed by NewApiDeleteUserProfileRequest with parameters below.

	@param userToken string - Unique identifier representing a user for which to fetch the personalization profile.
	@return DeleteUserProfileResponse
*/
func (c *APIClient) DeleteUserProfileWithContext(ctx context.Context, r ApiDeleteUserProfileRequest, opts ...Option) (*DeleteUserProfileResponse, error) {
	var (
		postBody    any
		returnValue *DeleteUserProfileResponse
	)

	requestPath := "/1/profiles/{userToken}"
	requestPath = strings.ReplaceAll(requestPath, "{userToken}", url.PathEscape(parameterToString(r.userToken)))

	headers := make(map[string]string)
	queryParams := url.Values{}
	if r.userToken == "" {
		return returnValue, reportError("Parameter `userToken` is required when calling `DeleteUserProfile`.")
	}

	// optional params if any
	for _, opt := range opts {
		switch opt.optionType {
		case "query":
			queryParams.Set(opt.name, opt.value)
		case "header":
			headers[opt.name] = opt.value
		}
	}

	req, err := c.prepareRequest(ctx, requestPath, http.MethodDelete, postBody, headers, queryParams)
	if err != nil {
		return returnValue, err
	}

	res, resBody, err := c.callAPI(req, false)
	if err != nil {
		return returnValue, err
	}
	if res == nil {
		return returnValue, reportError("res is nil")
	}

	if res.StatusCode >= 300 {
		newErr := &APIError{
			Message: string(resBody),
			Status:  res.StatusCode,
		}

		var v ErrorBase
		err = c.decode(&v, resBody)
		if err != nil {
			newErr.Message = err.Error()
			return returnValue, newErr
		}

		return returnValue, newErr
	}

	err = c.decode(&returnValue, resBody)
	if err != nil {
		return returnValue, reportError("cannot decode result: %w", err)
	}

	return returnValue, nil
}

/*
GetPersonalizationStrategy Wraps GetPersonalizationStrategyWithContext using context.Background.

Retrieves the current personalization strategy.

Required API Key ACLs:
  - recommendation

Request can be constructed by NewApiGetPersonalizationStrategyRequest with parameters below.

	@return PersonalizationStrategyParams
*/
func (c *APIClient) GetPersonalizationStrategy(opts ...Option) (*PersonalizationStrategyParams, error) {
	return c.GetPersonalizationStrategyWithContext(context.Background(), opts...)
}

/*
GetPersonalizationStrategy

Retrieves the current personalization strategy.

Required API Key ACLs:
  - recommendation

Request can be constructed by NewApiGetPersonalizationStrategyRequest with parameters below.

	@return PersonalizationStrategyParams
*/
func (c *APIClient) GetPersonalizationStrategyWithContext(ctx context.Context, opts ...Option) (*PersonalizationStrategyParams, error) {
	var (
		postBody    any
		returnValue *PersonalizationStrategyParams
	)

	requestPath := "/1/strategies/personalization"

	headers := make(map[string]string)
	queryParams := url.Values{}

	// optional params if any
	for _, opt := range opts {
		switch opt.optionType {
		case "query":
			queryParams.Set(opt.name, opt.value)
		case "header":
			headers[opt.name] = opt.value
		}
	}

	req, err := c.prepareRequest(ctx, requestPath, http.MethodGet, postBody, headers, queryParams)
	if err != nil {
		return returnValue, err
	}

	res, resBody, err := c.callAPI(req, false)
	if err != nil {
		return returnValue, err
	}
	if res == nil {
		return returnValue, reportError("res is nil")
	}

	if res.StatusCode >= 300 {
		newErr := &APIError{
			Message: string(resBody),
			Status:  res.StatusCode,
		}

		var v ErrorBase
		err = c.decode(&v, resBody)
		if err != nil {
			newErr.Message = err.Error()
			return returnValue, newErr
		}

		return returnValue, newErr
	}

	err = c.decode(&returnValue, resBody)
	if err != nil {
		return returnValue, reportError("cannot decode result: %w", err)
	}

	return returnValue, nil
}

func (r *ApiGetUserTokenProfileRequest) UnmarshalJSON(b []byte) error {
	req := map[string]json.RawMessage{}
	err := json.Unmarshal(b, &req)
	if err != nil {
		return fmt.Errorf("cannot unmarshal request: %w", err)
	}
	if v, ok := req["userToken"]; ok {
		err = json.Unmarshal(v, &r.userToken)
		if err != nil {
			err = json.Unmarshal(b, &r.userToken)
			if err != nil {
				return fmt.Errorf("cannot unmarshal userToken: %w", err)
			}
		}
	}

	return nil
}

// ApiGetUserTokenProfileRequest represents the request with all the parameters for the API call.
type ApiGetUserTokenProfileRequest struct {
	userToken string
}

// NewApiGetUserTokenProfileRequest creates an instance of the ApiGetUserTokenProfileRequest to be used for the API call.
func (c *APIClient) NewApiGetUserTokenProfileRequest(userToken string) ApiGetUserTokenProfileRequest {
	return ApiGetUserTokenProfileRequest{
		userToken: userToken,
	}
}

/*
GetUserTokenProfile Wraps GetUserTokenProfileWithContext using context.Background.

Retrieves a user profile and their affinities for different facets.

Required API Key ACLs:
  - recommendation

Request can be constructed by NewApiGetUserTokenProfileRequest with parameters below.

	@param userToken string - Unique identifier representing a user for which to fetch the personalization profile.
	@return GetUserTokenResponse
*/
func (c *APIClient) GetUserTokenProfile(r ApiGetUserTokenProfileRequest, opts ...Option) (*GetUserTokenResponse, error) {
	return c.GetUserTokenProfileWithContext(context.Background(), r, opts...)
}

/*
GetUserTokenProfile

Retrieves a user profile and their affinities for different facets.

Required API Key ACLs:
  - recommendation

Request can be constructed by NewApiGetUserTokenProfileRequest with parameters below.

	@param userToken string - Unique identifier representing a user for which to fetch the personalization profile.
	@return GetUserTokenResponse
*/
func (c *APIClient) GetUserTokenProfileWithContext(ctx context.Context, r ApiGetUserTokenProfileRequest, opts ...Option) (*GetUserTokenResponse, error) {
	var (
		postBody    any
		returnValue *GetUserTokenResponse
	)

	requestPath := "/1/profiles/personalization/{userToken}"
	requestPath = strings.ReplaceAll(requestPath, "{userToken}", url.PathEscape(parameterToString(r.userToken)))

	headers := make(map[string]string)
	queryParams := url.Values{}
	if r.userToken == "" {
		return returnValue, reportError("Parameter `userToken` is required when calling `GetUserTokenProfile`.")
	}

	// optional params if any
	for _, opt := range opts {
		switch opt.optionType {
		case "query":
			queryParams.Set(opt.name, opt.value)
		case "header":
			headers[opt.name] = opt.value
		}
	}

	req, err := c.prepareRequest(ctx, requestPath, http.MethodGet, postBody, headers, queryParams)
	if err != nil {
		return returnValue, err
	}

	res, resBody, err := c.callAPI(req, false)
	if err != nil {
		return returnValue, err
	}
	if res == nil {
		return returnValue, reportError("res is nil")
	}

	if res.StatusCode >= 300 {
		newErr := &APIError{
			Message: string(resBody),
			Status:  res.StatusCode,
		}

		var v ErrorBase
		err = c.decode(&v, resBody)
		if err != nil {
			newErr.Message = err.Error()
			return returnValue, newErr
		}

		return returnValue, newErr
	}

	err = c.decode(&returnValue, resBody)
	if err != nil {
		return returnValue, reportError("cannot decode result: %w", err)
	}

	return returnValue, nil
}

func (r *ApiSetPersonalizationStrategyRequest) UnmarshalJSON(b []byte) error {
	req := map[string]json.RawMessage{}
	err := json.Unmarshal(b, &req)
	if err != nil {
		return fmt.Errorf("cannot unmarshal request: %w", err)
	}
	if v, ok := req["personalizationStrategyParams"]; ok {
		err = json.Unmarshal(v, &r.personalizationStrategyParams)
		if err != nil {
			err = json.Unmarshal(b, &r.personalizationStrategyParams)
			if err != nil {
				return fmt.Errorf("cannot unmarshal personalizationStrategyParams: %w", err)
			}
		}
	} else {
		err = json.Unmarshal(b, &r.personalizationStrategyParams)
		if err != nil {
			return fmt.Errorf("cannot unmarshal body parameter personalizationStrategyParams: %w", err)
		}
	}

	return nil
}

// ApiSetPersonalizationStrategyRequest represents the request with all the parameters for the API call.
type ApiSetPersonalizationStrategyRequest struct {
	personalizationStrategyParams *PersonalizationStrategyParams
}

// NewApiSetPersonalizationStrategyRequest creates an instance of the ApiSetPersonalizationStrategyRequest to be used for the API call.
func (c *APIClient) NewApiSetPersonalizationStrategyRequest(personalizationStrategyParams *PersonalizationStrategyParams) ApiSetPersonalizationStrategyRequest {
	return ApiSetPersonalizationStrategyRequest{
		personalizationStrategyParams: personalizationStrategyParams,
	}
}

/*
SetPersonalizationStrategy Wraps SetPersonalizationStrategyWithContext using context.Background.

Creates a new personalization strategy.

Required API Key ACLs:
  - recommendation

Request can be constructed by NewApiSetPersonalizationStrategyRequest with parameters below.

	@param personalizationStrategyParams PersonalizationStrategyParams
	@return SetPersonalizationStrategyResponse
*/
func (c *APIClient) SetPersonalizationStrategy(r ApiSetPersonalizationStrategyRequest, opts ...Option) (*SetPersonalizationStrategyResponse, error) {
	return c.SetPersonalizationStrategyWithContext(context.Background(), r, opts...)
}

/*
SetPersonalizationStrategy

Creates a new personalization strategy.

Required API Key ACLs:
  - recommendation

Request can be constructed by NewApiSetPersonalizationStrategyRequest with parameters below.

	@param personalizationStrategyParams PersonalizationStrategyParams
	@return SetPersonalizationStrategyResponse
*/
func (c *APIClient) SetPersonalizationStrategyWithContext(ctx context.Context, r ApiSetPersonalizationStrategyRequest, opts ...Option) (*SetPersonalizationStrategyResponse, error) {
	var (
		postBody    any
		returnValue *SetPersonalizationStrategyResponse
	)

	requestPath := "/1/strategies/personalization"

	headers := make(map[string]string)
	queryParams := url.Values{}

	if r.personalizationStrategyParams == nil {
		return returnValue, reportError("Parameter `personalizationStrategyParams` is required when calling `SetPersonalizationStrategy`.")
	}

	// optional params if any
	for _, opt := range opts {
		switch opt.optionType {
		case "query":
			queryParams.Set(opt.name, opt.value)
		case "header":
			headers[opt.name] = opt.value
		}
	}

	// body params
	postBody = r.personalizationStrategyParams
	req, err := c.prepareRequest(ctx, requestPath, http.MethodPost, postBody, headers, queryParams)
	if err != nil {
		return returnValue, err
	}

	res, resBody, err := c.callAPI(req, false)
	if err != nil {
		return returnValue, err
	}
	if res == nil {
		return returnValue, reportError("res is nil")
	}

	if res.StatusCode >= 300 {
		newErr := &APIError{
			Message: string(resBody),
			Status:  res.StatusCode,
		}

		var v ErrorBase
		err = c.decode(&v, resBody)
		if err != nil {
			newErr.Message = err.Error()
			return returnValue, newErr
		}

		return returnValue, newErr
	}

	err = c.decode(&returnValue, resBody)
	if err != nil {
		return returnValue, reportError("cannot decode result: %w", err)
	}

	return returnValue, nil
}
