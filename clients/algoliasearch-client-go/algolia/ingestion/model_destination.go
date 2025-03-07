// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// Destination Destinations are Algolia resources like indices or event streams.
type Destination struct {
	// Universally unique identifier (UUID) of a destination resource.
	DestinationID string          `json:"destinationID"`
	Type          DestinationType `json:"type"`
	// Descriptive name for the resource.
	Name  string           `json:"name"`
	Input DestinationInput `json:"input"`
	// Date of creation in RFC3339 format.
	CreatedAt string `json:"createdAt"`
	// Date of last update in RFC3339 format.
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// Universally unique identifier (UUID) of an authentication resource.
	AuthenticationID *string `json:"authenticationID,omitempty"`
}

type DestinationOption func(f *Destination)

func WithDestinationUpdatedAt(val string) DestinationOption {
	return func(f *Destination) {
		f.UpdatedAt = &val
	}
}

func WithDestinationAuthenticationID(val string) DestinationOption {
	return func(f *Destination) {
		f.AuthenticationID = &val
	}
}

// NewDestination instantiates a new Destination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDestination(destinationID string, type_ DestinationType, name string, input DestinationInput, createdAt string, opts ...DestinationOption) *Destination {
	this := &Destination{}
	this.DestinationID = destinationID
	this.Type = type_
	this.Name = name
	this.Input = input
	this.CreatedAt = createdAt
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptyDestination return a pointer to an empty Destination object.
func NewEmptyDestination() *Destination {
	return &Destination{}
}

// GetDestinationID returns the DestinationID field value.
func (o *Destination) GetDestinationID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestinationID
}

// GetDestinationIDOk returns a tuple with the DestinationID field value
// and a boolean to check if the value has been set.
func (o *Destination) GetDestinationIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationID, true
}

// SetDestinationID sets field value.
func (o *Destination) SetDestinationID(v string) *Destination {
	o.DestinationID = v
	return o
}

// GetType returns the Type field value.
func (o *Destination) GetType() DestinationType {
	if o == nil {
		var ret DestinationType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Destination) GetTypeOk() (*DestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *Destination) SetType(v DestinationType) *Destination {
	o.Type = v
	return o
}

// GetName returns the Name field value.
func (o *Destination) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Destination) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *Destination) SetName(v string) *Destination {
	o.Name = v
	return o
}

// GetInput returns the Input field value.
func (o *Destination) GetInput() DestinationInput {
	if o == nil {
		var ret DestinationInput
		return ret
	}

	return o.Input
}

// GetInputOk returns a tuple with the Input field value
// and a boolean to check if the value has been set.
func (o *Destination) GetInputOk() (*DestinationInput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Input, true
}

// SetInput sets field value.
func (o *Destination) SetInput(v *DestinationInput) *Destination {
	o.Input = *v
	return o
}

// GetCreatedAt returns the CreatedAt field value.
func (o *Destination) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Destination) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *Destination) SetCreatedAt(v string) *Destination {
	o.CreatedAt = v
	return o
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Destination) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Destination) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Destination) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *Destination) SetUpdatedAt(v string) *Destination {
	o.UpdatedAt = &v
	return o
}

// GetAuthenticationID returns the AuthenticationID field value if set, zero value otherwise.
func (o *Destination) GetAuthenticationID() string {
	if o == nil || o.AuthenticationID == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationID
}

// GetAuthenticationIDOk returns a tuple with the AuthenticationID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Destination) GetAuthenticationIDOk() (*string, bool) {
	if o == nil || o.AuthenticationID == nil {
		return nil, false
	}
	return o.AuthenticationID, true
}

// HasAuthenticationID returns a boolean if a field has been set.
func (o *Destination) HasAuthenticationID() bool {
	if o != nil && o.AuthenticationID != nil {
		return true
	}

	return false
}

// SetAuthenticationID gets a reference to the given string and assigns it to the AuthenticationID field.
func (o *Destination) SetAuthenticationID(v string) *Destination {
	o.AuthenticationID = &v
	return o
}

func (o Destination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["destinationID"] = o.DestinationID
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["input"] = o.Input
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.AuthenticationID != nil {
		toSerialize["authenticationID"] = o.AuthenticationID
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal Destination: %w", err)
	}

	return serialized, nil
}

func (o Destination) String() string {
	out := ""
	out += fmt.Sprintf("  destinationID=%v\n", o.DestinationID)
	out += fmt.Sprintf("  type=%v\n", o.Type)
	out += fmt.Sprintf("  name=%v\n", o.Name)
	out += fmt.Sprintf("  input=%v\n", o.Input)
	out += fmt.Sprintf("  createdAt=%v\n", o.CreatedAt)
	out += fmt.Sprintf("  updatedAt=%v\n", o.UpdatedAt)
	out += fmt.Sprintf("  authenticationID=%v\n", o.AuthenticationID)
	return fmt.Sprintf("Destination {\n%s}", out)
}

type NullableDestination struct {
	value *Destination
	isSet bool
}

func (v NullableDestination) Get() *Destination {
	return v.value
}

func (v *NullableDestination) Set(val *Destination) {
	v.value = val
	v.isSet = true
}

func (v NullableDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestination(val *Destination) *NullableDestination {
	return &NullableDestination{value: val, isSet: true}
}

func (v NullableDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
