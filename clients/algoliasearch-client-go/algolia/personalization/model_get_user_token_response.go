// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package personalization

import (
	"encoding/json"
	"fmt"
)

// GetUserTokenResponse struct for GetUserTokenResponse.
type GetUserTokenResponse struct {
	// Unique pseudonymous or anonymous user identifier.  This helps with analytics and click and conversion events. For more information, see [user token](https://www.algolia.com/doc/guides/sending-events/concepts/usertoken/).
	UserToken string `json:"userToken"`
	// Date and time of the last event from this user, in RFC 3339 format.
	LastEventAt string `json:"lastEventAt"`
	// Scores for different facet values.  Scores represent the user affinity for a user profile towards specific facet values, given the personalization strategy and past events.
	Scores map[string]interface{} `json:"scores"`
}

// NewGetUserTokenResponse instantiates a new GetUserTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGetUserTokenResponse(userToken string, lastEventAt string, scores map[string]interface{}) *GetUserTokenResponse {
	this := &GetUserTokenResponse{}
	this.UserToken = userToken
	this.LastEventAt = lastEventAt
	this.Scores = scores
	return this
}

// NewEmptyGetUserTokenResponse return a pointer to an empty GetUserTokenResponse object.
func NewEmptyGetUserTokenResponse() *GetUserTokenResponse {
	return &GetUserTokenResponse{}
}

// GetUserToken returns the UserToken field value.
func (o *GetUserTokenResponse) GetUserToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value
// and a boolean to check if the value has been set.
func (o *GetUserTokenResponse) GetUserTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserToken, true
}

// SetUserToken sets field value.
func (o *GetUserTokenResponse) SetUserToken(v string) *GetUserTokenResponse {
	o.UserToken = v
	return o
}

// GetLastEventAt returns the LastEventAt field value.
func (o *GetUserTokenResponse) GetLastEventAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastEventAt
}

// GetLastEventAtOk returns a tuple with the LastEventAt field value
// and a boolean to check if the value has been set.
func (o *GetUserTokenResponse) GetLastEventAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastEventAt, true
}

// SetLastEventAt sets field value.
func (o *GetUserTokenResponse) SetLastEventAt(v string) *GetUserTokenResponse {
	o.LastEventAt = v
	return o
}

// GetScores returns the Scores field value.
func (o *GetUserTokenResponse) GetScores() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Scores
}

// GetScoresOk returns a tuple with the Scores field value
// and a boolean to check if the value has been set.
func (o *GetUserTokenResponse) GetScoresOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scores, true
}

// SetScores sets field value.
func (o *GetUserTokenResponse) SetScores(v map[string]interface{}) *GetUserTokenResponse {
	o.Scores = v
	return o
}

func (o GetUserTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["userToken"] = o.UserToken
	}
	if true {
		toSerialize["lastEventAt"] = o.LastEventAt
	}
	if true {
		toSerialize["scores"] = o.Scores
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal GetUserTokenResponse: %w", err)
	}

	return serialized, nil
}

func (o GetUserTokenResponse) String() string {
	out := ""
	out += fmt.Sprintf("  userToken=%v\n", o.UserToken)
	out += fmt.Sprintf("  lastEventAt=%v\n", o.LastEventAt)
	out += fmt.Sprintf("  scores=%v\n", o.Scores)
	return fmt.Sprintf("GetUserTokenResponse {\n%s}", out)
}

type NullableGetUserTokenResponse struct {
	value *GetUserTokenResponse
	isSet bool
}

func (v NullableGetUserTokenResponse) Get() *GetUserTokenResponse {
	return v.value
}

func (v *NullableGetUserTokenResponse) Set(val *GetUserTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserTokenResponse(val *GetUserTokenResponse) *NullableGetUserTokenResponse {
	return &NullableGetUserTokenResponse{value: val, isSet: true}
}

func (v NullableGetUserTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableGetUserTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
