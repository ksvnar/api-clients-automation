// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package search

import (
	"encoding/json"
	"fmt"
)

// SynonymType Synonym type.
type SynonymType string

// List of SynonymType.
const (
	SYNONYMTYPE_SYNONYM        SynonymType = "synonym"
	SYNONYMTYPE_ONEWAYSYNONYM  SynonymType = "onewaysynonym"
	SYNONYMTYPE_ALTCORRECTION1 SynonymType = "altcorrection1"
	SYNONYMTYPE_ALTCORRECTION2 SynonymType = "altcorrection2"
	SYNONYMTYPE_PLACEHOLDER    SynonymType = "placeholder"
)

// All allowed values of SynonymType enum.
var AllowedSynonymTypeEnumValues = []SynonymType{
	"synonym",
	"onewaysynonym",
	"altcorrection1",
	"altcorrection2",
	"placeholder",
}

func (v *SynonymType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return fmt.Errorf("failed to unmarshal value '%s' for enum 'SynonymType': %w", string(src), err)
	}
	enumTypeValue := SynonymType(value)
	for _, existing := range AllowedSynonymTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SynonymType", value)
}

// NewSynonymTypeFromValue returns a pointer to a valid SynonymType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSynonymTypeFromValue(v string) (*SynonymType, error) {
	ev := SynonymType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SynonymType: valid values are %v", v, AllowedSynonymTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SynonymType) IsValid() bool {
	for _, existing := range AllowedSynonymTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SynonymType value.
func (v SynonymType) Ptr() *SynonymType {
	return &v
}

type NullableSynonymType struct {
	value *SynonymType
	isSet bool
}

func (v NullableSynonymType) Get() *SynonymType {
	return v.value
}

func (v *NullableSynonymType) Set(val *SynonymType) {
	v.value = val
	v.isSet = true
}

func (v NullableSynonymType) IsSet() bool {
	return v.isSet
}

func (v *NullableSynonymType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSynonymType(val *SynonymType) *NullableSynonymType {
	return &NullableSynonymType{value: val, isSet: true}
}

func (v NullableSynonymType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableSynonymType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
