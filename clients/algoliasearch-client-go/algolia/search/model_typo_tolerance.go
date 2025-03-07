// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package search

import (
	"encoding/json"
	"fmt"
)

// TypoTolerance - Whether [typo tolerance](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/typo-tolerance/) is enabled and how it is applied.  If typo tolerance is true, `min`, or `strict`, [word splitting and concetenation](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/in-depth/splitting-and-concatenation/) is also active.
type TypoTolerance struct {
	TypoToleranceEnum *TypoToleranceEnum
	Bool              *bool
}

// boolAsTypoTolerance is a convenience function that returns bool wrapped in TypoTolerance.
func BoolAsTypoTolerance(v bool) *TypoTolerance {
	return &TypoTolerance{
		Bool: &v,
	}
}

// TypoToleranceEnumAsTypoTolerance is a convenience function that returns TypoToleranceEnum wrapped in TypoTolerance.
func TypoToleranceEnumAsTypoTolerance(v TypoToleranceEnum) *TypoTolerance {
	return &TypoTolerance{
		TypoToleranceEnum: &v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct.
func (dst *TypoTolerance) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal data into TypoToleranceEnum
	err = newStrictDecoder(data).Decode(&dst.TypoToleranceEnum)
	if err == nil && validateStruct(dst.TypoToleranceEnum) == nil {
		jsonTypoToleranceEnum, _ := json.Marshal(dst.TypoToleranceEnum)
		if string(jsonTypoToleranceEnum) == "{}" { // empty struct
			dst.TypoToleranceEnum = nil
		} else {
			return nil
		}
	} else {
		dst.TypoToleranceEnum = nil
	}

	// try to unmarshal data into Bool
	err = newStrictDecoder(data).Decode(&dst.Bool)
	if err == nil && validateStruct(dst.Bool) == nil {
		jsonBool, _ := json.Marshal(dst.Bool)
		if string(jsonBool) == "{}" { // empty struct
			dst.Bool = nil
		} else {
			return nil
		}
	} else {
		dst.Bool = nil
	}

	return fmt.Errorf("Data failed to match schemas in oneOf(TypoTolerance)")
}

// Marshal data from the first non-nil pointers in the struct to JSON.
func (src TypoTolerance) MarshalJSON() ([]byte, error) {
	if src.TypoToleranceEnum != nil {
		serialized, err := json.Marshal(&src.TypoToleranceEnum)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal one of TypoToleranceEnum of TypoTolerance: %w", err)
		}

		return serialized, nil
	}

	if src.Bool != nil {
		serialized, err := json.Marshal(&src.Bool)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal one of Bool of TypoTolerance: %w", err)
		}

		return serialized, nil
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance.
func (obj TypoTolerance) GetActualInstance() any {
	if obj.TypoToleranceEnum != nil {
		return *obj.TypoToleranceEnum
	}

	if obj.Bool != nil {
		return *obj.Bool
	}

	// all schemas are nil
	return nil
}

type NullableTypoTolerance struct {
	value *TypoTolerance
	isSet bool
}

func (v NullableTypoTolerance) Get() *TypoTolerance {
	return v.value
}

func (v *NullableTypoTolerance) Set(val *TypoTolerance) {
	v.value = val
	v.isSet = true
}

func (v NullableTypoTolerance) IsSet() bool {
	return v.isSet
}

func (v *NullableTypoTolerance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypoTolerance(val *TypoTolerance) *NullableTypoTolerance {
	return &NullableTypoTolerance{value: val, isSet: true}
}

func (v NullableTypoTolerance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableTypoTolerance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
