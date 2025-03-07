// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package analytics

import (
	"encoding/json"
	"fmt"
)

// GetTopSearchesResponse - struct for GetTopSearchesResponse.
type GetTopSearchesResponse struct {
	TopSearchesResponse                     *TopSearchesResponse
	TopSearchesResponseWithAnalytics        *TopSearchesResponseWithAnalytics
	TopSearchesResponseWithRevenueAnalytics *TopSearchesResponseWithRevenueAnalytics
}

// TopSearchesResponseAsGetTopSearchesResponse is a convenience function that returns TopSearchesResponse wrapped in GetTopSearchesResponse.
func TopSearchesResponseAsGetTopSearchesResponse(v *TopSearchesResponse) *GetTopSearchesResponse {
	return &GetTopSearchesResponse{
		TopSearchesResponse: v,
	}
}

// TopSearchesResponseWithAnalyticsAsGetTopSearchesResponse is a convenience function that returns TopSearchesResponseWithAnalytics wrapped in GetTopSearchesResponse.
func TopSearchesResponseWithAnalyticsAsGetTopSearchesResponse(v *TopSearchesResponseWithAnalytics) *GetTopSearchesResponse {
	return &GetTopSearchesResponse{
		TopSearchesResponseWithAnalytics: v,
	}
}

// TopSearchesResponseWithRevenueAnalyticsAsGetTopSearchesResponse is a convenience function that returns TopSearchesResponseWithRevenueAnalytics wrapped in GetTopSearchesResponse.
func TopSearchesResponseWithRevenueAnalyticsAsGetTopSearchesResponse(v *TopSearchesResponseWithRevenueAnalytics) *GetTopSearchesResponse {
	return &GetTopSearchesResponse{
		TopSearchesResponseWithRevenueAnalytics: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct.
func (dst *GetTopSearchesResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal data into TopSearchesResponse
	err = newStrictDecoder(data).Decode(&dst.TopSearchesResponse)
	if err == nil && validateStruct(dst.TopSearchesResponse) == nil {
		jsonTopSearchesResponse, _ := json.Marshal(dst.TopSearchesResponse)
		if string(jsonTopSearchesResponse) == "{}" { // empty struct
			dst.TopSearchesResponse = nil
		} else {
			return nil
		}
	} else {
		dst.TopSearchesResponse = nil
	}

	// try to unmarshal data into TopSearchesResponseWithAnalytics
	err = newStrictDecoder(data).Decode(&dst.TopSearchesResponseWithAnalytics)
	if err == nil && validateStruct(dst.TopSearchesResponseWithAnalytics) == nil {
		jsonTopSearchesResponseWithAnalytics, _ := json.Marshal(dst.TopSearchesResponseWithAnalytics)
		if string(jsonTopSearchesResponseWithAnalytics) == "{}" { // empty struct
			dst.TopSearchesResponseWithAnalytics = nil
		} else {
			return nil
		}
	} else {
		dst.TopSearchesResponseWithAnalytics = nil
	}

	// try to unmarshal data into TopSearchesResponseWithRevenueAnalytics
	err = newStrictDecoder(data).Decode(&dst.TopSearchesResponseWithRevenueAnalytics)
	if err == nil && validateStruct(dst.TopSearchesResponseWithRevenueAnalytics) == nil {
		jsonTopSearchesResponseWithRevenueAnalytics, _ := json.Marshal(dst.TopSearchesResponseWithRevenueAnalytics)
		if string(jsonTopSearchesResponseWithRevenueAnalytics) == "{}" { // empty struct
			dst.TopSearchesResponseWithRevenueAnalytics = nil
		} else {
			return nil
		}
	} else {
		dst.TopSearchesResponseWithRevenueAnalytics = nil
	}

	return fmt.Errorf("Data failed to match schemas in oneOf(GetTopSearchesResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON.
func (src GetTopSearchesResponse) MarshalJSON() ([]byte, error) {
	if src.TopSearchesResponse != nil {
		serialized, err := json.Marshal(&src.TopSearchesResponse)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal one of TopSearchesResponse of GetTopSearchesResponse: %w", err)
		}

		return serialized, nil
	}

	if src.TopSearchesResponseWithAnalytics != nil {
		serialized, err := json.Marshal(&src.TopSearchesResponseWithAnalytics)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal one of TopSearchesResponseWithAnalytics of GetTopSearchesResponse: %w", err)
		}

		return serialized, nil
	}

	if src.TopSearchesResponseWithRevenueAnalytics != nil {
		serialized, err := json.Marshal(&src.TopSearchesResponseWithRevenueAnalytics)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal one of TopSearchesResponseWithRevenueAnalytics of GetTopSearchesResponse: %w", err)
		}

		return serialized, nil
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance.
func (obj GetTopSearchesResponse) GetActualInstance() any {
	if obj.TopSearchesResponse != nil {
		return *obj.TopSearchesResponse
	}

	if obj.TopSearchesResponseWithAnalytics != nil {
		return *obj.TopSearchesResponseWithAnalytics
	}

	if obj.TopSearchesResponseWithRevenueAnalytics != nil {
		return *obj.TopSearchesResponseWithRevenueAnalytics
	}

	// all schemas are nil
	return nil
}

type NullableGetTopSearchesResponse struct {
	value *GetTopSearchesResponse
	isSet bool
}

func (v NullableGetTopSearchesResponse) Get() *GetTopSearchesResponse {
	return v.value
}

func (v *NullableGetTopSearchesResponse) Set(val *GetTopSearchesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTopSearchesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTopSearchesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTopSearchesResponse(val *GetTopSearchesResponse) *NullableGetTopSearchesResponse {
	return &NullableGetTopSearchesResponse{value: val, isSet: true}
}

func (v NullableGetTopSearchesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableGetTopSearchesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
