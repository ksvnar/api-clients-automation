// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package suggestions

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/algolia/algoliasearch-client-go/v4/algolia/call"
)

type Option struct {
	optionType string
	name       string
	value      string
}

func QueryParamOption(name string, val any) Option {
	return Option{
		optionType: "query",
		name:       name,
		value:      parameterToString(val),
	}
}

func HeaderParamOption(name string, val any) Option {
	return Option{
		optionType: "header",
		name:       name,
		value:      parameterToString(val),
	}
}

type ApiCreateConfigRequest struct {
	querySuggestionsIndexWithIndexParam *QuerySuggestionsIndexWithIndexParam
}

func (r *ApiCreateConfigRequest) UnmarshalJSON(b []byte) error {
	req := map[string]json.RawMessage{}
	err := json.Unmarshal(b, &req)
	if err != nil {
		return err
	}
	if v, ok := req["querySuggestionsIndexWithIndexParam"]; ok { //querySuggestionsIndexWithIndexParam
		err = json.Unmarshal(v, &r.querySuggestionsIndexWithIndexParam)
		if err != nil {
			err = json.Unmarshal(b, &r.querySuggestionsIndexWithIndexParam)
			if err != nil {
				return err
			}
		}
	} else {
		err = json.Unmarshal(b, &r.querySuggestionsIndexWithIndexParam)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r ApiCreateConfigRequest) WithQuerySuggestionsIndexWithIndexParam(querySuggestionsIndexWithIndexParam *QuerySuggestionsIndexWithIndexParam) ApiCreateConfigRequest {
	r.querySuggestionsIndexWithIndexParam = querySuggestionsIndexWithIndexParam
	return r
}

// @return ApiCreateConfigRequest
func (c *APIClient) NewApiCreateConfigRequest() ApiCreateConfigRequest {
	return ApiCreateConfigRequest{}
}

// CreateConfig wraps CreateConfigWithContext using context.Background.
func (c *APIClient) CreateConfig(r ApiCreateConfigRequest, opts ...Option) (*SuccessResponse, error) {
	return c.CreateConfigWithContext(context.Background(), r, opts...)
}

// @return SuccessResponse
func (c *APIClient) CreateConfigWithContext(ctx context.Context, r ApiCreateConfigRequest, opts ...Option) (*SuccessResponse, error) {
	var (
		postBody    any
		returnValue *SuccessResponse
	)

	requestPath := "/1/configs"

	headers := make(map[string]string)
	queryParams := url.Values{}
	if r.querySuggestionsIndexWithIndexParam == nil {
		return returnValue, reportError("querySuggestionsIndexWithIndexParam is required and must be specified")
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
	postBody = r.querySuggestionsIndexWithIndexParam
	req, err := c.prepareRequest(ctx, requestPath, http.MethodPost, postBody, headers, queryParams)
	if err != nil {
		return returnValue, err
	}

	res, err := c.callAPI(req, call.Write)
	if err != nil {
		return returnValue, err
	}
	if res == nil {
		return returnValue, reportError("res is nil")
	}

	resBody, err := io.ReadAll(res.Body)
	res.Body.Close()
	res.Body = io.NopCloser(bytes.NewBuffer(resBody))
	if err != nil {
		return returnValue, err
	}

	if res.StatusCode >= 300 {
		newErr := &APIError{
			Message: string(resBody),
			Status:  res.StatusCode,
		}
		if res.StatusCode == 400 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 401 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 403 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 422 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 500 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
		}
		return returnValue, newErr
	}

	err = c.decode(&returnValue, resBody, res.Header.Get("Content-Type"))
	if err != nil {
		return returnValue, reportError("cannot decode result: %w", err)
	}

	return returnValue, nil
}

type ApiDelRequest struct {
	path       string
	parameters map[string]interface{}
}

func (r *ApiDelRequest) UnmarshalJSON(b []byte) error {
	req := map[string]json.RawMessage{}
	err := json.Unmarshal(b, &req)
	if err != nil {
		return err
	}
	if v, ok := req["path"]; ok { //path
		err = json.Unmarshal(v, &r.path)
		if err != nil {
			err = json.Unmarshal(b, &r.path)
			if err != nil {
				return err
			}
		}
	}
	if v, ok := req["parameters"]; ok { //parameters
		err = json.Unmarshal(v, &r.parameters)
		if err != nil {
			err = json.Unmarshal(b, &r.parameters)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Query parameters to be applied to the current query.
func (r ApiDelRequest) WithParameters(parameters map[string]interface{}) ApiDelRequest {
	r.parameters = parameters
	return r
}

// @return ApiDelRequest
func (c *APIClient) NewApiDelRequest(path string) ApiDelRequest {
	return ApiDelRequest{
		path: path,
	}
}

// Del wraps DelWithContext using context.Background.
func (c *APIClient) Del(r ApiDelRequest, opts ...Option) (map[string]interface{}, error) {
	return c.DelWithContext(context.Background(), r, opts...)
}

// @return map[string]interface{}
func (c *APIClient) DelWithContext(ctx context.Context, r ApiDelRequest, opts ...Option) (map[string]interface{}, error) {
	var (
		postBody    any
		returnValue map[string]interface{}
	)

	requestPath := "/1{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(parameterToString(r.path)), -1)

	headers := make(map[string]string)
	queryParams := url.Values{}

	if !isNilorEmpty(r.parameters) {
		for k, v := range r.parameters {
			queryParams.Set(k, parameterToString(v))
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

	res, err := c.callAPI(req, call.Write)
	if err != nil {
		return returnValue, err
	}
	if res == nil {
		return returnValue, reportError("res is nil")
	}

	resBody, err := io.ReadAll(res.Body)
	res.Body.Close()
	res.Body = io.NopCloser(bytes.NewBuffer(resBody))
	if err != nil {
		return returnValue, err
	}

	if res.StatusCode >= 300 {
		newErr := &APIError{
			Message: string(resBody),
			Status:  res.StatusCode,
		}
		if res.StatusCode == 400 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 402 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 403 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 404 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
		}
		return returnValue, newErr
	}

	err = c.decode(&returnValue, resBody, res.Header.Get("Content-Type"))
	if err != nil {
		return returnValue, reportError("cannot decode result: %w", err)
	}

	return returnValue, nil
}

type ApiDeleteConfigRequest struct {
	indexName string
}

func (r *ApiDeleteConfigRequest) UnmarshalJSON(b []byte) error {
	req := map[string]json.RawMessage{}
	err := json.Unmarshal(b, &req)
	if err != nil {
		return err
	}
	if v, ok := req["indexName"]; ok { //indexName
		err = json.Unmarshal(v, &r.indexName)
		if err != nil {
			err = json.Unmarshal(b, &r.indexName)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// @return ApiDeleteConfigRequest
func (c *APIClient) NewApiDeleteConfigRequest(indexName string) ApiDeleteConfigRequest {
	return ApiDeleteConfigRequest{
		indexName: indexName,
	}
}

// DeleteConfig wraps DeleteConfigWithContext using context.Background.
func (c *APIClient) DeleteConfig(r ApiDeleteConfigRequest, opts ...Option) (*SuccessResponse, error) {
	return c.DeleteConfigWithContext(context.Background(), r, opts...)
}

// @return SuccessResponse
func (c *APIClient) DeleteConfigWithContext(ctx context.Context, r ApiDeleteConfigRequest, opts ...Option) (*SuccessResponse, error) {
	var (
		postBody    any
		returnValue *SuccessResponse
	)

	requestPath := "/1/configs/{indexName}"
	requestPath = strings.Replace(requestPath, "{"+"indexName"+"}", url.PathEscape(parameterToString(r.indexName)), -1)

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

	req, err := c.prepareRequest(ctx, requestPath, http.MethodDelete, postBody, headers, queryParams)
	if err != nil {
		return returnValue, err
	}

	res, err := c.callAPI(req, call.Write)
	if err != nil {
		return returnValue, err
	}
	if res == nil {
		return returnValue, reportError("res is nil")
	}

	resBody, err := io.ReadAll(res.Body)
	res.Body.Close()
	res.Body = io.NopCloser(bytes.NewBuffer(resBody))
	if err != nil {
		return returnValue, err
	}

	if res.StatusCode >= 300 {
		newErr := &APIError{
			Message: string(resBody),
			Status:  res.StatusCode,
		}
		if res.StatusCode == 401 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 403 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 500 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
		}
		return returnValue, newErr
	}

	err = c.decode(&returnValue, resBody, res.Header.Get("Content-Type"))
	if err != nil {
		return returnValue, reportError("cannot decode result: %w", err)
	}

	return returnValue, nil
}

type ApiGetRequest struct {
	path       string
	parameters map[string]interface{}
}

func (r *ApiGetRequest) UnmarshalJSON(b []byte) error {
	req := map[string]json.RawMessage{}
	err := json.Unmarshal(b, &req)
	if err != nil {
		return err
	}
	if v, ok := req["path"]; ok { //path
		err = json.Unmarshal(v, &r.path)
		if err != nil {
			err = json.Unmarshal(b, &r.path)
			if err != nil {
				return err
			}
		}
	}
	if v, ok := req["parameters"]; ok { //parameters
		err = json.Unmarshal(v, &r.parameters)
		if err != nil {
			err = json.Unmarshal(b, &r.parameters)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Query parameters to be applied to the current query.
func (r ApiGetRequest) WithParameters(parameters map[string]interface{}) ApiGetRequest {
	r.parameters = parameters
	return r
}

// @return ApiGetRequest
func (c *APIClient) NewApiGetRequest(path string) ApiGetRequest {
	return ApiGetRequest{
		path: path,
	}
}

// Get wraps GetWithContext using context.Background.
func (c *APIClient) Get(r ApiGetRequest, opts ...Option) (map[string]interface{}, error) {
	return c.GetWithContext(context.Background(), r, opts...)
}

// @return map[string]interface{}
func (c *APIClient) GetWithContext(ctx context.Context, r ApiGetRequest, opts ...Option) (map[string]interface{}, error) {
	var (
		postBody    any
		returnValue map[string]interface{}
	)

	requestPath := "/1{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(parameterToString(r.path)), -1)

	headers := make(map[string]string)
	queryParams := url.Values{}

	if !isNilorEmpty(r.parameters) {
		for k, v := range r.parameters {
			queryParams.Set(k, parameterToString(v))
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

	res, err := c.callAPI(req, call.Write)
	if err != nil {
		return returnValue, err
	}
	if res == nil {
		return returnValue, reportError("res is nil")
	}

	resBody, err := io.ReadAll(res.Body)
	res.Body.Close()
	res.Body = io.NopCloser(bytes.NewBuffer(resBody))
	if err != nil {
		return returnValue, err
	}

	if res.StatusCode >= 300 {
		newErr := &APIError{
			Message: string(resBody),
			Status:  res.StatusCode,
		}
		if res.StatusCode == 400 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 402 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 403 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 404 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
		}
		return returnValue, newErr
	}

	err = c.decode(&returnValue, resBody, res.Header.Get("Content-Type"))
	if err != nil {
		return returnValue, reportError("cannot decode result: %w", err)
	}

	return returnValue, nil
}

type ApiGetAllConfigsRequest struct {
}

func (r *ApiGetAllConfigsRequest) UnmarshalJSON(b []byte) error {
	req := map[string]json.RawMessage{}
	err := json.Unmarshal(b, &req)
	if err != nil {
		return err
	}

	return nil
}

// @return ApiGetAllConfigsRequest
func (c *APIClient) NewApiGetAllConfigsRequest() ApiGetAllConfigsRequest {
	return ApiGetAllConfigsRequest{}
}

// GetAllConfigs wraps GetAllConfigsWithContext using context.Background.
func (c *APIClient) GetAllConfigs(r ApiGetAllConfigsRequest, opts ...Option) ([]QuerySuggestionsIndex, error) {
	return c.GetAllConfigsWithContext(context.Background(), r, opts...)
}

// @return []QuerySuggestionsIndex
func (c *APIClient) GetAllConfigsWithContext(ctx context.Context, r ApiGetAllConfigsRequest, opts ...Option) ([]QuerySuggestionsIndex, error) {
	var (
		postBody    any
		returnValue []QuerySuggestionsIndex
	)

	requestPath := "/1/configs"

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

	res, err := c.callAPI(req, call.Write)
	if err != nil {
		return returnValue, err
	}
	if res == nil {
		return returnValue, reportError("res is nil")
	}

	resBody, err := io.ReadAll(res.Body)
	res.Body.Close()
	res.Body = io.NopCloser(bytes.NewBuffer(resBody))
	if err != nil {
		return returnValue, err
	}

	if res.StatusCode >= 300 {
		newErr := &APIError{
			Message: string(resBody),
			Status:  res.StatusCode,
		}
		if res.StatusCode == 401 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 403 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 422 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 500 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
		}
		return returnValue, newErr
	}

	err = c.decode(&returnValue, resBody, res.Header.Get("Content-Type"))
	if err != nil {
		return returnValue, reportError("cannot decode result: %w", err)
	}

	return returnValue, nil
}

type ApiGetConfigRequest struct {
	indexName string
}

func (r *ApiGetConfigRequest) UnmarshalJSON(b []byte) error {
	req := map[string]json.RawMessage{}
	err := json.Unmarshal(b, &req)
	if err != nil {
		return err
	}
	if v, ok := req["indexName"]; ok { //indexName
		err = json.Unmarshal(v, &r.indexName)
		if err != nil {
			err = json.Unmarshal(b, &r.indexName)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// @return ApiGetConfigRequest
func (c *APIClient) NewApiGetConfigRequest(indexName string) ApiGetConfigRequest {
	return ApiGetConfigRequest{
		indexName: indexName,
	}
}

// GetConfig wraps GetConfigWithContext using context.Background.
func (c *APIClient) GetConfig(r ApiGetConfigRequest, opts ...Option) (*QuerySuggestionsIndex, error) {
	return c.GetConfigWithContext(context.Background(), r, opts...)
}

// @return QuerySuggestionsIndex
func (c *APIClient) GetConfigWithContext(ctx context.Context, r ApiGetConfigRequest, opts ...Option) (*QuerySuggestionsIndex, error) {
	var (
		postBody    any
		returnValue *QuerySuggestionsIndex
	)

	requestPath := "/1/configs/{indexName}"
	requestPath = strings.Replace(requestPath, "{"+"indexName"+"}", url.PathEscape(parameterToString(r.indexName)), -1)

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

	res, err := c.callAPI(req, call.Write)
	if err != nil {
		return returnValue, err
	}
	if res == nil {
		return returnValue, reportError("res is nil")
	}

	resBody, err := io.ReadAll(res.Body)
	res.Body.Close()
	res.Body = io.NopCloser(bytes.NewBuffer(resBody))
	if err != nil {
		return returnValue, err
	}

	if res.StatusCode >= 300 {
		newErr := &APIError{
			Message: string(resBody),
			Status:  res.StatusCode,
		}
		if res.StatusCode == 400 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 401 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 403 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 404 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 500 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
		}
		return returnValue, newErr
	}

	err = c.decode(&returnValue, resBody, res.Header.Get("Content-Type"))
	if err != nil {
		return returnValue, reportError("cannot decode result: %w", err)
	}

	return returnValue, nil
}

type ApiGetConfigStatusRequest struct {
	indexName string
}

func (r *ApiGetConfigStatusRequest) UnmarshalJSON(b []byte) error {
	req := map[string]json.RawMessage{}
	err := json.Unmarshal(b, &req)
	if err != nil {
		return err
	}
	if v, ok := req["indexName"]; ok { //indexName
		err = json.Unmarshal(v, &r.indexName)
		if err != nil {
			err = json.Unmarshal(b, &r.indexName)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// @return ApiGetConfigStatusRequest
func (c *APIClient) NewApiGetConfigStatusRequest(indexName string) ApiGetConfigStatusRequest {
	return ApiGetConfigStatusRequest{
		indexName: indexName,
	}
}

// GetConfigStatus wraps GetConfigStatusWithContext using context.Background.
func (c *APIClient) GetConfigStatus(r ApiGetConfigStatusRequest, opts ...Option) (*Status, error) {
	return c.GetConfigStatusWithContext(context.Background(), r, opts...)
}

// @return Status
func (c *APIClient) GetConfigStatusWithContext(ctx context.Context, r ApiGetConfigStatusRequest, opts ...Option) (*Status, error) {
	var (
		postBody    any
		returnValue *Status
	)

	requestPath := "/1/configs/{indexName}/status"
	requestPath = strings.Replace(requestPath, "{"+"indexName"+"}", url.PathEscape(parameterToString(r.indexName)), -1)

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

	res, err := c.callAPI(req, call.Write)
	if err != nil {
		return returnValue, err
	}
	if res == nil {
		return returnValue, reportError("res is nil")
	}

	resBody, err := io.ReadAll(res.Body)
	res.Body.Close()
	res.Body = io.NopCloser(bytes.NewBuffer(resBody))
	if err != nil {
		return returnValue, err
	}

	if res.StatusCode >= 300 {
		newErr := &APIError{
			Message: string(resBody),
			Status:  res.StatusCode,
		}
		if res.StatusCode == 401 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 403 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 500 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
		}
		return returnValue, newErr
	}

	err = c.decode(&returnValue, resBody, res.Header.Get("Content-Type"))
	if err != nil {
		return returnValue, reportError("cannot decode result: %w", err)
	}

	return returnValue, nil
}

type ApiGetLogFileRequest struct {
	indexName string
}

func (r *ApiGetLogFileRequest) UnmarshalJSON(b []byte) error {
	req := map[string]json.RawMessage{}
	err := json.Unmarshal(b, &req)
	if err != nil {
		return err
	}
	if v, ok := req["indexName"]; ok { //indexName
		err = json.Unmarshal(v, &r.indexName)
		if err != nil {
			err = json.Unmarshal(b, &r.indexName)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// @return ApiGetLogFileRequest
func (c *APIClient) NewApiGetLogFileRequest(indexName string) ApiGetLogFileRequest {
	return ApiGetLogFileRequest{
		indexName: indexName,
	}
}

// GetLogFile wraps GetLogFileWithContext using context.Background.
func (c *APIClient) GetLogFile(r ApiGetLogFileRequest, opts ...Option) ([]LogFile, error) {
	return c.GetLogFileWithContext(context.Background(), r, opts...)
}

// @return []LogFile
func (c *APIClient) GetLogFileWithContext(ctx context.Context, r ApiGetLogFileRequest, opts ...Option) ([]LogFile, error) {
	var (
		postBody    any
		returnValue []LogFile
	)

	requestPath := "/1/logs/{indexName}"
	requestPath = strings.Replace(requestPath, "{"+"indexName"+"}", url.PathEscape(parameterToString(r.indexName)), -1)

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

	res, err := c.callAPI(req, call.Write)
	if err != nil {
		return returnValue, err
	}
	if res == nil {
		return returnValue, reportError("res is nil")
	}

	resBody, err := io.ReadAll(res.Body)
	res.Body.Close()
	res.Body = io.NopCloser(bytes.NewBuffer(resBody))
	if err != nil {
		return returnValue, err
	}

	if res.StatusCode >= 300 {
		newErr := &APIError{
			Message: string(resBody),
			Status:  res.StatusCode,
		}
		if res.StatusCode == 401 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 403 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 404 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 500 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
		}
		return returnValue, newErr
	}

	err = c.decode(&returnValue, resBody, res.Header.Get("Content-Type"))
	if err != nil {
		return returnValue, reportError("cannot decode result: %w", err)
	}

	return returnValue, nil
}

type ApiPostRequest struct {
	path       string
	parameters map[string]interface{}
	body       map[string]interface{}
}

func (r *ApiPostRequest) UnmarshalJSON(b []byte) error {
	req := map[string]json.RawMessage{}
	err := json.Unmarshal(b, &req)
	if err != nil {
		return err
	}
	if v, ok := req["path"]; ok { //path
		err = json.Unmarshal(v, &r.path)
		if err != nil {
			err = json.Unmarshal(b, &r.path)
			if err != nil {
				return err
			}
		}
	}
	if v, ok := req["parameters"]; ok { //parameters
		err = json.Unmarshal(v, &r.parameters)
		if err != nil {
			err = json.Unmarshal(b, &r.parameters)
			if err != nil {
				return err
			}
		}
	}
	if v, ok := req["body"]; ok { //body
		err = json.Unmarshal(v, &r.body)
		if err != nil {
			err = json.Unmarshal(b, &r.body)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Query parameters to be applied to the current query.
func (r ApiPostRequest) WithParameters(parameters map[string]interface{}) ApiPostRequest {
	r.parameters = parameters
	return r
}

// The parameters to send with the custom request.
func (r ApiPostRequest) WithBody(body map[string]interface{}) ApiPostRequest {
	r.body = body
	return r
}

// @return ApiPostRequest
func (c *APIClient) NewApiPostRequest(path string) ApiPostRequest {
	return ApiPostRequest{
		path: path,
	}
}

// Post wraps PostWithContext using context.Background.
func (c *APIClient) Post(r ApiPostRequest, opts ...Option) (map[string]interface{}, error) {
	return c.PostWithContext(context.Background(), r, opts...)
}

// @return map[string]interface{}
func (c *APIClient) PostWithContext(ctx context.Context, r ApiPostRequest, opts ...Option) (map[string]interface{}, error) {
	var (
		postBody    any
		returnValue map[string]interface{}
	)

	requestPath := "/1{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(parameterToString(r.path)), -1)

	headers := make(map[string]string)
	queryParams := url.Values{}

	if !isNilorEmpty(r.parameters) {
		for k, v := range r.parameters {
			queryParams.Set(k, parameterToString(v))
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
	if isNilorEmpty(r.body) {
		postBody = "{}"
	} else {
		postBody = r.body
	}
	req, err := c.prepareRequest(ctx, requestPath, http.MethodPost, postBody, headers, queryParams)
	if err != nil {
		return returnValue, err
	}

	res, err := c.callAPI(req, call.Write)
	if err != nil {
		return returnValue, err
	}
	if res == nil {
		return returnValue, reportError("res is nil")
	}

	resBody, err := io.ReadAll(res.Body)
	res.Body.Close()
	res.Body = io.NopCloser(bytes.NewBuffer(resBody))
	if err != nil {
		return returnValue, err
	}

	if res.StatusCode >= 300 {
		newErr := &APIError{
			Message: string(resBody),
			Status:  res.StatusCode,
		}
		if res.StatusCode == 400 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 402 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 403 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 404 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
		}
		return returnValue, newErr
	}

	err = c.decode(&returnValue, resBody, res.Header.Get("Content-Type"))
	if err != nil {
		return returnValue, reportError("cannot decode result: %w", err)
	}

	return returnValue, nil
}

type ApiPutRequest struct {
	path       string
	parameters map[string]interface{}
	body       map[string]interface{}
}

func (r *ApiPutRequest) UnmarshalJSON(b []byte) error {
	req := map[string]json.RawMessage{}
	err := json.Unmarshal(b, &req)
	if err != nil {
		return err
	}
	if v, ok := req["path"]; ok { //path
		err = json.Unmarshal(v, &r.path)
		if err != nil {
			err = json.Unmarshal(b, &r.path)
			if err != nil {
				return err
			}
		}
	}
	if v, ok := req["parameters"]; ok { //parameters
		err = json.Unmarshal(v, &r.parameters)
		if err != nil {
			err = json.Unmarshal(b, &r.parameters)
			if err != nil {
				return err
			}
		}
	}
	if v, ok := req["body"]; ok { //body
		err = json.Unmarshal(v, &r.body)
		if err != nil {
			err = json.Unmarshal(b, &r.body)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Query parameters to be applied to the current query.
func (r ApiPutRequest) WithParameters(parameters map[string]interface{}) ApiPutRequest {
	r.parameters = parameters
	return r
}

// The parameters to send with the custom request.
func (r ApiPutRequest) WithBody(body map[string]interface{}) ApiPutRequest {
	r.body = body
	return r
}

// @return ApiPutRequest
func (c *APIClient) NewApiPutRequest(path string) ApiPutRequest {
	return ApiPutRequest{
		path: path,
	}
}

// Put wraps PutWithContext using context.Background.
func (c *APIClient) Put(r ApiPutRequest, opts ...Option) (map[string]interface{}, error) {
	return c.PutWithContext(context.Background(), r, opts...)
}

// @return map[string]interface{}
func (c *APIClient) PutWithContext(ctx context.Context, r ApiPutRequest, opts ...Option) (map[string]interface{}, error) {
	var (
		postBody    any
		returnValue map[string]interface{}
	)

	requestPath := "/1{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(parameterToString(r.path)), -1)

	headers := make(map[string]string)
	queryParams := url.Values{}

	if !isNilorEmpty(r.parameters) {
		for k, v := range r.parameters {
			queryParams.Set(k, parameterToString(v))
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
	if isNilorEmpty(r.body) {
		postBody = "{}"
	} else {
		postBody = r.body
	}
	req, err := c.prepareRequest(ctx, requestPath, http.MethodPut, postBody, headers, queryParams)
	if err != nil {
		return returnValue, err
	}

	res, err := c.callAPI(req, call.Write)
	if err != nil {
		return returnValue, err
	}
	if res == nil {
		return returnValue, reportError("res is nil")
	}

	resBody, err := io.ReadAll(res.Body)
	res.Body.Close()
	res.Body = io.NopCloser(bytes.NewBuffer(resBody))
	if err != nil {
		return returnValue, err
	}

	if res.StatusCode >= 300 {
		newErr := &APIError{
			Message: string(resBody),
			Status:  res.StatusCode,
		}
		if res.StatusCode == 400 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 402 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 403 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 404 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
		}
		return returnValue, newErr
	}

	err = c.decode(&returnValue, resBody, res.Header.Get("Content-Type"))
	if err != nil {
		return returnValue, reportError("cannot decode result: %w", err)
	}

	return returnValue, nil
}

type ApiUpdateConfigRequest struct {
	indexName                  string
	querySuggestionsIndexParam *QuerySuggestionsIndexParam
}

func (r *ApiUpdateConfigRequest) UnmarshalJSON(b []byte) error {
	req := map[string]json.RawMessage{}
	err := json.Unmarshal(b, &req)
	if err != nil {
		return err
	}
	if v, ok := req["indexName"]; ok { //indexName
		err = json.Unmarshal(v, &r.indexName)
		if err != nil {
			err = json.Unmarshal(b, &r.indexName)
			if err != nil {
				return err
			}
		}
	}
	if v, ok := req["querySuggestionsIndexParam"]; ok { //querySuggestionsIndexParam
		err = json.Unmarshal(v, &r.querySuggestionsIndexParam)
		if err != nil {
			err = json.Unmarshal(b, &r.querySuggestionsIndexParam)
			if err != nil {
				return err
			}
		}
	} else {
		err = json.Unmarshal(b, &r.querySuggestionsIndexParam)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r ApiUpdateConfigRequest) WithQuerySuggestionsIndexParam(querySuggestionsIndexParam *QuerySuggestionsIndexParam) ApiUpdateConfigRequest {
	r.querySuggestionsIndexParam = querySuggestionsIndexParam
	return r
}

// @return ApiUpdateConfigRequest
func (c *APIClient) NewApiUpdateConfigRequest(indexName string) ApiUpdateConfigRequest {
	return ApiUpdateConfigRequest{
		indexName: indexName,
	}
}

// UpdateConfig wraps UpdateConfigWithContext using context.Background.
func (c *APIClient) UpdateConfig(r ApiUpdateConfigRequest, opts ...Option) (*SuccessResponse, error) {
	return c.UpdateConfigWithContext(context.Background(), r, opts...)
}

// @return SuccessResponse
func (c *APIClient) UpdateConfigWithContext(ctx context.Context, r ApiUpdateConfigRequest, opts ...Option) (*SuccessResponse, error) {
	var (
		postBody    any
		returnValue *SuccessResponse
	)

	requestPath := "/1/configs/{indexName}"
	requestPath = strings.Replace(requestPath, "{"+"indexName"+"}", url.PathEscape(parameterToString(r.indexName)), -1)

	headers := make(map[string]string)
	queryParams := url.Values{}
	if r.querySuggestionsIndexParam == nil {
		return returnValue, reportError("querySuggestionsIndexParam is required and must be specified")
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
	postBody = r.querySuggestionsIndexParam
	req, err := c.prepareRequest(ctx, requestPath, http.MethodPut, postBody, headers, queryParams)
	if err != nil {
		return returnValue, err
	}

	res, err := c.callAPI(req, call.Write)
	if err != nil {
		return returnValue, err
	}
	if res == nil {
		return returnValue, reportError("res is nil")
	}

	resBody, err := io.ReadAll(res.Body)
	res.Body.Close()
	res.Body = io.NopCloser(bytes.NewBuffer(resBody))
	if err != nil {
		return returnValue, err
	}

	if res.StatusCode >= 300 {
		newErr := &APIError{
			Message: string(resBody),
			Status:  res.StatusCode,
		}
		if res.StatusCode == 401 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 500 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
		}
		return returnValue, newErr
	}

	err = c.decode(&returnValue, resBody, res.Header.Get("Content-Type"))
	if err != nil {
		return returnValue, reportError("cannot decode result: %w", err)
	}

	return returnValue, nil
}
