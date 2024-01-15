// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package insights

import (
	"encoding/json"
	"fmt"
)

// ConvertedObjectIDs Use this event to track when users convert on items unrelated to a previous Algolia request. For example, if you don't use Algolia to build your category pages, use this event.  To track conversion events related to Algolia requests, use the \"Converted object IDs after search\" event.
type ConvertedObjectIDs struct {
	// The name of the event, up to 64 ASCII characters.  Consider naming events consistently—for example, by adopting Segment's [object-action](https://segment.com/academy/collecting-data/naming-conventions-for-clean-data/#the-object-action-framework) framework.
	EventName string          `json:"eventName"`
	EventType ConversionEvent `json:"eventType"`
	// The name of an Algolia index.
	Index string `json:"index"`
	// The object IDs of the records that are part of the event.
	ObjectIDs []string `json:"objectIDs"`
	// An anonymous or pseudonymous user identifier.  > **Note**: Never include personally identifiable information in user tokens.
	UserToken string `json:"userToken"`
	// An identifier for authenticated users.  > **Note**: Never include personally identifiable information in user tokens.
	AuthenticatedUserToken *string `json:"authenticatedUserToken,omitempty"`
	// The timestamp of the event in milliseconds in [Unix epoch time](https://wikipedia.org/wiki/Unix_time). By default, the Insights API uses the time it receives an event as its timestamp.
	Timestamp *int64 `json:"timestamp,omitempty"`
}

type ConvertedObjectIDsOption func(f *ConvertedObjectIDs)

func WithConvertedObjectIDsAuthenticatedUserToken(val string) ConvertedObjectIDsOption {
	return func(f *ConvertedObjectIDs) {
		f.AuthenticatedUserToken = &val
	}
}

func WithConvertedObjectIDsTimestamp(val int64) ConvertedObjectIDsOption {
	return func(f *ConvertedObjectIDs) {
		f.Timestamp = &val
	}
}

// NewConvertedObjectIDs instantiates a new ConvertedObjectIDs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewConvertedObjectIDs(eventName string, eventType ConversionEvent, index string, objectIDs []string, userToken string, opts ...ConvertedObjectIDsOption) *ConvertedObjectIDs {
	this := &ConvertedObjectIDs{}
	this.EventName = eventName
	this.EventType = eventType
	this.Index = index
	this.ObjectIDs = objectIDs
	this.UserToken = userToken
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptyConvertedObjectIDs return a pointer to an empty ConvertedObjectIDs object.
func NewEmptyConvertedObjectIDs() *ConvertedObjectIDs {
	return &ConvertedObjectIDs{}
}

// GetEventName returns the EventName field value.
func (o *ConvertedObjectIDs) GetEventName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value
// and a boolean to check if the value has been set.
func (o *ConvertedObjectIDs) GetEventNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventName, true
}

// SetEventName sets field value.
func (o *ConvertedObjectIDs) SetEventName(v string) *ConvertedObjectIDs {
	o.EventName = v
	return o
}

// GetEventType returns the EventType field value.
func (o *ConvertedObjectIDs) GetEventType() ConversionEvent {
	if o == nil {
		var ret ConversionEvent
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *ConvertedObjectIDs) GetEventTypeOk() (*ConversionEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value.
func (o *ConvertedObjectIDs) SetEventType(v ConversionEvent) *ConvertedObjectIDs {
	o.EventType = v
	return o
}

// GetIndex returns the Index field value.
func (o *ConvertedObjectIDs) GetIndex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *ConvertedObjectIDs) GetIndexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value.
func (o *ConvertedObjectIDs) SetIndex(v string) *ConvertedObjectIDs {
	o.Index = v
	return o
}

// GetObjectIDs returns the ObjectIDs field value.
func (o *ConvertedObjectIDs) GetObjectIDs() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ObjectIDs
}

// GetObjectIDsOk returns a tuple with the ObjectIDs field value
// and a boolean to check if the value has been set.
func (o *ConvertedObjectIDs) GetObjectIDsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObjectIDs, true
}

// SetObjectIDs sets field value.
func (o *ConvertedObjectIDs) SetObjectIDs(v []string) *ConvertedObjectIDs {
	o.ObjectIDs = v
	return o
}

// GetUserToken returns the UserToken field value.
func (o *ConvertedObjectIDs) GetUserToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value
// and a boolean to check if the value has been set.
func (o *ConvertedObjectIDs) GetUserTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserToken, true
}

// SetUserToken sets field value.
func (o *ConvertedObjectIDs) SetUserToken(v string) *ConvertedObjectIDs {
	o.UserToken = v
	return o
}

// GetAuthenticatedUserToken returns the AuthenticatedUserToken field value if set, zero value otherwise.
func (o *ConvertedObjectIDs) GetAuthenticatedUserToken() string {
	if o == nil || o.AuthenticatedUserToken == nil {
		var ret string
		return ret
	}
	return *o.AuthenticatedUserToken
}

// GetAuthenticatedUserTokenOk returns a tuple with the AuthenticatedUserToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertedObjectIDs) GetAuthenticatedUserTokenOk() (*string, bool) {
	if o == nil || o.AuthenticatedUserToken == nil {
		return nil, false
	}
	return o.AuthenticatedUserToken, true
}

// HasAuthenticatedUserToken returns a boolean if a field has been set.
func (o *ConvertedObjectIDs) HasAuthenticatedUserToken() bool {
	if o != nil && o.AuthenticatedUserToken != nil {
		return true
	}

	return false
}

// SetAuthenticatedUserToken gets a reference to the given string and assigns it to the AuthenticatedUserToken field.
func (o *ConvertedObjectIDs) SetAuthenticatedUserToken(v string) *ConvertedObjectIDs {
	o.AuthenticatedUserToken = &v
	return o
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ConvertedObjectIDs) GetTimestamp() int64 {
	if o == nil || o.Timestamp == nil {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertedObjectIDs) GetTimestampOk() (*int64, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ConvertedObjectIDs) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *ConvertedObjectIDs) SetTimestamp(v int64) *ConvertedObjectIDs {
	o.Timestamp = &v
	return o
}

func (o ConvertedObjectIDs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["eventName"] = o.EventName
	}
	if true {
		toSerialize["eventType"] = o.EventType
	}
	if true {
		toSerialize["index"] = o.Index
	}
	if true {
		toSerialize["objectIDs"] = o.ObjectIDs
	}
	if true {
		toSerialize["userToken"] = o.UserToken
	}
	if o.AuthenticatedUserToken != nil {
		toSerialize["authenticatedUserToken"] = o.AuthenticatedUserToken
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	return json.Marshal(toSerialize)
}

func (o ConvertedObjectIDs) String() string {
	out := ""
	out += fmt.Sprintf("  eventName=%v\n", o.EventName)
	out += fmt.Sprintf("  eventType=%v\n", o.EventType)
	out += fmt.Sprintf("  index=%v\n", o.Index)
	out += fmt.Sprintf("  objectIDs=%v\n", o.ObjectIDs)
	out += fmt.Sprintf("  userToken=%v\n", o.UserToken)
	out += fmt.Sprintf("  authenticatedUserToken=%v\n", o.AuthenticatedUserToken)
	out += fmt.Sprintf("  timestamp=%v\n", o.Timestamp)
	return fmt.Sprintf("ConvertedObjectIDs {\n%s}", out)
}

type NullableConvertedObjectIDs struct {
	value *ConvertedObjectIDs
	isSet bool
}

func (v NullableConvertedObjectIDs) Get() *ConvertedObjectIDs {
	return v.value
}

func (v *NullableConvertedObjectIDs) Set(val *ConvertedObjectIDs) {
	v.value = val
	v.isSet = true
}

func (v NullableConvertedObjectIDs) IsSet() bool {
	return v.isSet
}

func (v *NullableConvertedObjectIDs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvertedObjectIDs(val *ConvertedObjectIDs) *NullableConvertedObjectIDs {
	return &NullableConvertedObjectIDs{value: val, isSet: true}
}

func (v NullableConvertedObjectIDs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvertedObjectIDs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
