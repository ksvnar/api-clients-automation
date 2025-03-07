// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// ListDestinationsResponse struct for ListDestinationsResponse.
type ListDestinationsResponse struct {
	Destinations []Destination `json:"destinations"`
	Pagination   Pagination    `json:"pagination"`
}

// NewListDestinationsResponse instantiates a new ListDestinationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListDestinationsResponse(destinations []Destination, pagination Pagination) *ListDestinationsResponse {
	this := &ListDestinationsResponse{}
	this.Destinations = destinations
	this.Pagination = pagination
	return this
}

// NewEmptyListDestinationsResponse return a pointer to an empty ListDestinationsResponse object.
func NewEmptyListDestinationsResponse() *ListDestinationsResponse {
	return &ListDestinationsResponse{}
}

// GetDestinations returns the Destinations field value.
func (o *ListDestinationsResponse) GetDestinations() []Destination {
	if o == nil {
		var ret []Destination
		return ret
	}

	return o.Destinations
}

// GetDestinationsOk returns a tuple with the Destinations field value
// and a boolean to check if the value has been set.
func (o *ListDestinationsResponse) GetDestinationsOk() ([]Destination, bool) {
	if o == nil {
		return nil, false
	}
	return o.Destinations, true
}

// SetDestinations sets field value.
func (o *ListDestinationsResponse) SetDestinations(v []Destination) *ListDestinationsResponse {
	o.Destinations = v
	return o
}

// GetPagination returns the Pagination field value.
func (o *ListDestinationsResponse) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ListDestinationsResponse) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value.
func (o *ListDestinationsResponse) SetPagination(v *Pagination) *ListDestinationsResponse {
	o.Pagination = *v
	return o
}

func (o ListDestinationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["destinations"] = o.Destinations
	}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal ListDestinationsResponse: %w", err)
	}

	return serialized, nil
}

func (o ListDestinationsResponse) String() string {
	out := ""
	out += fmt.Sprintf("  destinations=%v\n", o.Destinations)
	out += fmt.Sprintf("  pagination=%v\n", o.Pagination)
	return fmt.Sprintf("ListDestinationsResponse {\n%s}", out)
}

type NullableListDestinationsResponse struct {
	value *ListDestinationsResponse
	isSet bool
}

func (v NullableListDestinationsResponse) Get() *ListDestinationsResponse {
	return v.value
}

func (v *NullableListDestinationsResponse) Set(val *ListDestinationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListDestinationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListDestinationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDestinationsResponse(val *ListDestinationsResponse) *NullableListDestinationsResponse {
	return &NullableListDestinationsResponse{value: val, isSet: true}
}

func (v NullableListDestinationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableListDestinationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
