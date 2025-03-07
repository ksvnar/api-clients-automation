// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package search

import (
	"encoding/json"
	"fmt"
)

// ExactOnSingleWordQuery Determines how the [Exact ranking criterion](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/override-search-engine-defaults/in-depth/adjust-exact-settings/#turn-off-exact-for-some-attributes) is computed when the search query has only one word.  - `attribute`.   The Exact ranking criterion is 1 if the query word and attribute value are the same.   For example, a search for \"road\" will match the value \"road\", but not \"road trip\".  - `none`.   The Exact ranking criterion is ignored on single-word searches.  - `word`.   The Exact ranking criterion is 1 if the query word is found in the attribute value.   The query word must have at least 3 characters and must not be a stop word.   Only exact matches will be highlighted,   partial and prefix matches won't.
type ExactOnSingleWordQuery string

// List of exactOnSingleWordQuery.
const (
	EXACTONSINGLEWORDQUERY_ATTRIBUTE ExactOnSingleWordQuery = "attribute"
	EXACTONSINGLEWORDQUERY_NONE      ExactOnSingleWordQuery = "none"
	EXACTONSINGLEWORDQUERY_WORD      ExactOnSingleWordQuery = "word"
)

// All allowed values of ExactOnSingleWordQuery enum.
var AllowedExactOnSingleWordQueryEnumValues = []ExactOnSingleWordQuery{
	"attribute",
	"none",
	"word",
}

func (v *ExactOnSingleWordQuery) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return fmt.Errorf("failed to unmarshal value '%s' for enum 'ExactOnSingleWordQuery': %w", string(src), err)
	}
	enumTypeValue := ExactOnSingleWordQuery(value)
	for _, existing := range AllowedExactOnSingleWordQueryEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ExactOnSingleWordQuery", value)
}

// NewExactOnSingleWordQueryFromValue returns a pointer to a valid ExactOnSingleWordQuery
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExactOnSingleWordQueryFromValue(v string) (*ExactOnSingleWordQuery, error) {
	ev := ExactOnSingleWordQuery(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ExactOnSingleWordQuery: valid values are %v", v, AllowedExactOnSingleWordQueryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExactOnSingleWordQuery) IsValid() bool {
	for _, existing := range AllowedExactOnSingleWordQueryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to exactOnSingleWordQuery value.
func (v ExactOnSingleWordQuery) Ptr() *ExactOnSingleWordQuery {
	return &v
}

type NullableExactOnSingleWordQuery struct {
	value *ExactOnSingleWordQuery
	isSet bool
}

func (v NullableExactOnSingleWordQuery) Get() *ExactOnSingleWordQuery {
	return v.value
}

func (v *NullableExactOnSingleWordQuery) Set(val *ExactOnSingleWordQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableExactOnSingleWordQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableExactOnSingleWordQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExactOnSingleWordQuery(val *ExactOnSingleWordQuery) *NullableExactOnSingleWordQuery {
	return &NullableExactOnSingleWordQuery{value: val, isSet: true}
}

func (v NullableExactOnSingleWordQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableExactOnSingleWordQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
