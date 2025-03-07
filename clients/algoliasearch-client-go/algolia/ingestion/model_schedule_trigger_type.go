// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// ScheduleTriggerType Task runs on a schedule.
type ScheduleTriggerType string

// List of ScheduleTriggerType.
const (
	SCHEDULETRIGGERTYPE_SCHEDULE ScheduleTriggerType = "schedule"
)

// All allowed values of ScheduleTriggerType enum.
var AllowedScheduleTriggerTypeEnumValues = []ScheduleTriggerType{
	"schedule",
}

func (v *ScheduleTriggerType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return fmt.Errorf("failed to unmarshal value '%s' for enum 'ScheduleTriggerType': %w", string(src), err)
	}
	enumTypeValue := ScheduleTriggerType(value)
	for _, existing := range AllowedScheduleTriggerTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ScheduleTriggerType", value)
}

// NewScheduleTriggerTypeFromValue returns a pointer to a valid ScheduleTriggerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScheduleTriggerTypeFromValue(v string) (*ScheduleTriggerType, error) {
	ev := ScheduleTriggerType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ScheduleTriggerType: valid values are %v", v, AllowedScheduleTriggerTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScheduleTriggerType) IsValid() bool {
	for _, existing := range AllowedScheduleTriggerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScheduleTriggerType value.
func (v ScheduleTriggerType) Ptr() *ScheduleTriggerType {
	return &v
}

type NullableScheduleTriggerType struct {
	value *ScheduleTriggerType
	isSet bool
}

func (v NullableScheduleTriggerType) Get() *ScheduleTriggerType {
	return v.value
}

func (v *NullableScheduleTriggerType) Set(val *ScheduleTriggerType) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleTriggerType) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleTriggerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleTriggerType(val *ScheduleTriggerType) *NullableScheduleTriggerType {
	return &NullableScheduleTriggerType{value: val, isSet: true}
}

func (v NullableScheduleTriggerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableScheduleTriggerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
