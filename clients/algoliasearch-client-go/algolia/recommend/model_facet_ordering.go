// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package recommend

import (
	"encoding/json"
	"fmt"
)

// FacetOrdering Order of facet names and facet values in your UI.
type FacetOrdering struct {
	Facets *Facets `json:"facets,omitempty"`
	// Order of facet values. One object for each facet.
	Values *map[string]Value `json:"values,omitempty"`
}

type FacetOrderingOption func(f *FacetOrdering)

func WithFacetOrderingFacets(val Facets) FacetOrderingOption {
	return func(f *FacetOrdering) {
		f.Facets = &val
	}
}

func WithFacetOrderingValues(val map[string]Value) FacetOrderingOption {
	return func(f *FacetOrdering) {
		f.Values = &val
	}
}

// NewFacetOrdering instantiates a new FacetOrdering object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFacetOrdering(opts ...FacetOrderingOption) *FacetOrdering {
	this := &FacetOrdering{}
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptyFacetOrdering return a pointer to an empty FacetOrdering object.
func NewEmptyFacetOrdering() *FacetOrdering {
	return &FacetOrdering{}
}

// GetFacets returns the Facets field value if set, zero value otherwise.
func (o *FacetOrdering) GetFacets() Facets {
	if o == nil || o.Facets == nil {
		var ret Facets
		return ret
	}
	return *o.Facets
}

// GetFacetsOk returns a tuple with the Facets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacetOrdering) GetFacetsOk() (*Facets, bool) {
	if o == nil || o.Facets == nil {
		return nil, false
	}
	return o.Facets, true
}

// HasFacets returns a boolean if a field has been set.
func (o *FacetOrdering) HasFacets() bool {
	if o != nil && o.Facets != nil {
		return true
	}

	return false
}

// SetFacets gets a reference to the given Facets and assigns it to the Facets field.
func (o *FacetOrdering) SetFacets(v *Facets) *FacetOrdering {
	o.Facets = v
	return o
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *FacetOrdering) GetValues() map[string]Value {
	if o == nil || o.Values == nil {
		var ret map[string]Value
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacetOrdering) GetValuesOk() (*map[string]Value, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *FacetOrdering) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given map[string]Value and assigns it to the Values field.
func (o *FacetOrdering) SetValues(v map[string]Value) *FacetOrdering {
	o.Values = &v
	return o
}

func (o FacetOrdering) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Facets != nil {
		toSerialize["facets"] = o.Facets
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal FacetOrdering: %w", err)
	}

	return serialized, nil
}

func (o FacetOrdering) String() string {
	out := ""
	out += fmt.Sprintf("  facets=%v\n", o.Facets)
	out += fmt.Sprintf("  values=%v\n", o.Values)
	return fmt.Sprintf("FacetOrdering {\n%s}", out)
}

type NullableFacetOrdering struct {
	value *FacetOrdering
	isSet bool
}

func (v NullableFacetOrdering) Get() *FacetOrdering {
	return v.value
}

func (v *NullableFacetOrdering) Set(val *FacetOrdering) {
	v.value = val
	v.isSet = true
}

func (v NullableFacetOrdering) IsSet() bool {
	return v.isSet
}

func (v *NullableFacetOrdering) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFacetOrdering(val *FacetOrdering) *NullableFacetOrdering {
	return &NullableFacetOrdering{value: val, isSet: true}
}

func (v NullableFacetOrdering) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableFacetOrdering) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
