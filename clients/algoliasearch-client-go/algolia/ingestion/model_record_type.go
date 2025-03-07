// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// RecordType Record type for ecommerce sources.
type RecordType string

// List of RecordType.
const (
	RECORDTYPE_PRODUCT RecordType = "product"
	RECORDTYPE_VARIANT RecordType = "variant"
)

// All allowed values of RecordType enum.
var AllowedRecordTypeEnumValues = []RecordType{
	"product",
	"variant",
}

func (v *RecordType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return fmt.Errorf("failed to unmarshal value '%s' for enum 'RecordType': %w", string(src), err)
	}
	enumTypeValue := RecordType(value)
	for _, existing := range AllowedRecordTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RecordType", value)
}

// NewRecordTypeFromValue returns a pointer to a valid RecordType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRecordTypeFromValue(v string) (*RecordType, error) {
	ev := RecordType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RecordType: valid values are %v", v, AllowedRecordTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RecordType) IsValid() bool {
	for _, existing := range AllowedRecordTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RecordType value.
func (v RecordType) Ptr() *RecordType {
	return &v
}

type NullableRecordType struct {
	value *RecordType
	isSet bool
}

func (v NullableRecordType) Get() *RecordType {
	return v.value
}

func (v *NullableRecordType) Set(val *RecordType) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordType) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordType(val *RecordType) *NullableRecordType {
	return &NullableRecordType{value: val, isSet: true}
}

func (v NullableRecordType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableRecordType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
