// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// AuthAlgoliaInsights Credentials for authenticating with the Algolia Insights API.
type AuthAlgoliaInsights struct {
	// Algolia application ID.
	AppID string `json:"appID"`
	// Algolia API key with the ACL: `search`. This field is `null` in the API response.
	ApiKey string `json:"apiKey"`
}

// NewAuthAlgoliaInsights instantiates a new AuthAlgoliaInsights object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAuthAlgoliaInsights(appID string, apiKey string) *AuthAlgoliaInsights {
	this := &AuthAlgoliaInsights{}
	this.AppID = appID
	this.ApiKey = apiKey
	return this
}

// NewEmptyAuthAlgoliaInsights return a pointer to an empty AuthAlgoliaInsights object.
func NewEmptyAuthAlgoliaInsights() *AuthAlgoliaInsights {
	return &AuthAlgoliaInsights{}
}

// GetAppID returns the AppID field value.
func (o *AuthAlgoliaInsights) GetAppID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppID
}

// GetAppIDOk returns a tuple with the AppID field value
// and a boolean to check if the value has been set.
func (o *AuthAlgoliaInsights) GetAppIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppID, true
}

// SetAppID sets field value.
func (o *AuthAlgoliaInsights) SetAppID(v string) *AuthAlgoliaInsights {
	o.AppID = v
	return o
}

// GetApiKey returns the ApiKey field value.
func (o *AuthAlgoliaInsights) GetApiKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *AuthAlgoliaInsights) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value.
func (o *AuthAlgoliaInsights) SetApiKey(v string) *AuthAlgoliaInsights {
	o.ApiKey = v
	return o
}

func (o AuthAlgoliaInsights) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["appID"] = o.AppID
	}
	if true {
		toSerialize["apiKey"] = o.ApiKey
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal AuthAlgoliaInsights: %w", err)
	}

	return serialized, nil
}

func (o AuthAlgoliaInsights) String() string {
	out := ""
	out += fmt.Sprintf("  appID=%v\n", o.AppID)
	out += fmt.Sprintf("  apiKey=%v\n", o.ApiKey)
	return fmt.Sprintf("AuthAlgoliaInsights {\n%s}", out)
}

type NullableAuthAlgoliaInsights struct {
	value *AuthAlgoliaInsights
	isSet bool
}

func (v NullableAuthAlgoliaInsights) Get() *AuthAlgoliaInsights {
	return v.value
}

func (v *NullableAuthAlgoliaInsights) Set(val *AuthAlgoliaInsights) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthAlgoliaInsights) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthAlgoliaInsights) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthAlgoliaInsights(val *AuthAlgoliaInsights) *NullableAuthAlgoliaInsights {
	return &NullableAuthAlgoliaInsights{value: val, isSet: true}
}

func (v NullableAuthAlgoliaInsights) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableAuthAlgoliaInsights) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
