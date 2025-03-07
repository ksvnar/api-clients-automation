// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// AuthBasic Credentials for authenticating with user name and password.
type AuthBasic struct {
	// Username.
	Username string `json:"username"`
	// Password. This field is `null` in the API response.
	Password string `json:"password"`
}

// NewAuthBasic instantiates a new AuthBasic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAuthBasic(username string, password string) *AuthBasic {
	this := &AuthBasic{}
	this.Username = username
	this.Password = password
	return this
}

// NewEmptyAuthBasic return a pointer to an empty AuthBasic object.
func NewEmptyAuthBasic() *AuthBasic {
	return &AuthBasic{}
}

// GetUsername returns the Username field value.
func (o *AuthBasic) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *AuthBasic) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value.
func (o *AuthBasic) SetUsername(v string) *AuthBasic {
	o.Username = v
	return o
}

// GetPassword returns the Password field value.
func (o *AuthBasic) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *AuthBasic) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value.
func (o *AuthBasic) SetPassword(v string) *AuthBasic {
	o.Password = v
	return o
}

func (o AuthBasic) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["password"] = o.Password
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal AuthBasic: %w", err)
	}

	return serialized, nil
}

func (o AuthBasic) String() string {
	out := ""
	out += fmt.Sprintf("  username=%v\n", o.Username)
	out += fmt.Sprintf("  password=%v\n", o.Password)
	return fmt.Sprintf("AuthBasic {\n%s}", out)
}

type NullableAuthBasic struct {
	value *AuthBasic
	isSet bool
}

func (v NullableAuthBasic) Get() *AuthBasic {
	return v.value
}

func (v *NullableAuthBasic) Set(val *AuthBasic) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthBasic) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthBasic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthBasic(val *AuthBasic) *NullableAuthBasic {
	return &NullableAuthBasic{value: val, isSet: true}
}

func (v NullableAuthBasic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableAuthBasic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
