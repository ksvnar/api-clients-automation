// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// Window The window parameter represents the holds the dates used to query the Observability data from the database in a given window.
type Window struct {
	// A date in format RFC3339 representing the oldest possible data in query window.
	StartDate string `json:"startDate"`
	// A date in format RFC3339 representing the newest possible data in query window.
	EndDate string `json:"endDate"`
}

// NewWindow instantiates a new Window object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWindow(startDate string, endDate string) *Window {
	this := &Window{}
	this.StartDate = startDate
	this.EndDate = endDate
	return this
}

// NewEmptyWindow return a pointer to an empty Window object.
func NewEmptyWindow() *Window {
	return &Window{}
}

// GetStartDate returns the StartDate field value.
func (o *Window) GetStartDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *Window) GetStartDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value.
func (o *Window) SetStartDate(v string) *Window {
	o.StartDate = v
	return o
}

// GetEndDate returns the EndDate field value.
func (o *Window) GetEndDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value
// and a boolean to check if the value has been set.
func (o *Window) GetEndDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndDate, true
}

// SetEndDate sets field value.
func (o *Window) SetEndDate(v string) *Window {
	o.EndDate = v
	return o
}

func (o Window) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["startDate"] = o.StartDate
	}
	if true {
		toSerialize["endDate"] = o.EndDate
	}
	return json.Marshal(toSerialize)
}

func (o Window) String() string {
	out := ""
	out += fmt.Sprintf("  startDate=%v\n", o.StartDate)
	out += fmt.Sprintf("  endDate=%v\n", o.EndDate)
	return fmt.Sprintf("Window {\n%s}", out)
}

type NullableWindow struct {
	value *Window
	isSet bool
}

func (v NullableWindow) Get() *Window {
	return v.value
}

func (v *NullableWindow) Set(val *Window) {
	v.value = val
	v.isSet = true
}

func (v NullableWindow) IsSet() bool {
	return v.isSet
}

func (v *NullableWindow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWindow(val *Window) *NullableWindow {
	return &NullableWindow{value: val, isSet: true}
}

func (v NullableWindow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWindow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
