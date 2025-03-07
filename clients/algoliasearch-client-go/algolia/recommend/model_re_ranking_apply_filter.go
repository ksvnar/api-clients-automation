// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package recommend

import (
	"encoding/json"
	"fmt"
)

// ReRankingApplyFilter - Restrict [Dynamic Re-Ranking](https://www.algolia.com/doc/guides/algolia-ai/re-ranking/) to records that match these filters.
type ReRankingApplyFilter struct {
	ArrayOfReRankingApplyFilter *[]ReRankingApplyFilter
	String                      *string
}

// []ReRankingApplyFilterAsReRankingApplyFilter is a convenience function that returns []ReRankingApplyFilter wrapped in ReRankingApplyFilter.
func ArrayOfReRankingApplyFilterAsReRankingApplyFilter(v []ReRankingApplyFilter) *ReRankingApplyFilter {
	return &ReRankingApplyFilter{
		ArrayOfReRankingApplyFilter: &v,
	}
}

// stringAsReRankingApplyFilter is a convenience function that returns string wrapped in ReRankingApplyFilter.
func StringAsReRankingApplyFilter(v string) *ReRankingApplyFilter {
	return &ReRankingApplyFilter{
		String: &v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct.
func (dst *ReRankingApplyFilter) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal data into ArrayOfReRankingApplyFilter
	err = newStrictDecoder(data).Decode(&dst.ArrayOfReRankingApplyFilter)
	if err == nil && validateStruct(dst.ArrayOfReRankingApplyFilter) == nil {
		jsonArrayOfReRankingApplyFilter, _ := json.Marshal(dst.ArrayOfReRankingApplyFilter)
		if string(jsonArrayOfReRankingApplyFilter) == "{}" { // empty struct
			dst.ArrayOfReRankingApplyFilter = nil
		} else {
			return nil
		}
	} else {
		dst.ArrayOfReRankingApplyFilter = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil && validateStruct(dst.String) == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("Data failed to match schemas in oneOf(ReRankingApplyFilter)")
}

// Marshal data from the first non-nil pointers in the struct to JSON.
func (src ReRankingApplyFilter) MarshalJSON() ([]byte, error) {
	if src.ArrayOfReRankingApplyFilter != nil {
		serialized, err := json.Marshal(&src.ArrayOfReRankingApplyFilter)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal one of ArrayOfReRankingApplyFilter of ReRankingApplyFilter: %w", err)
		}

		return serialized, nil
	}

	if src.String != nil {
		serialized, err := json.Marshal(&src.String)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal one of String of ReRankingApplyFilter: %w", err)
		}

		return serialized, nil
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance.
func (obj ReRankingApplyFilter) GetActualInstance() any {
	if obj.ArrayOfReRankingApplyFilter != nil {
		return *obj.ArrayOfReRankingApplyFilter
	}

	if obj.String != nil {
		return *obj.String
	}

	// all schemas are nil
	return nil
}

type NullableReRankingApplyFilter struct {
	value *ReRankingApplyFilter
	isSet bool
}

func (v NullableReRankingApplyFilter) Get() *ReRankingApplyFilter {
	return v.value
}

func (v *NullableReRankingApplyFilter) Set(val *ReRankingApplyFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableReRankingApplyFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableReRankingApplyFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReRankingApplyFilter(val *ReRankingApplyFilter) *NullableReRankingApplyFilter {
	return &NullableReRankingApplyFilter{value: val, isSet: true}
}

func (v NullableReRankingApplyFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableReRankingApplyFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
