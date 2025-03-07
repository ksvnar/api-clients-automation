// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package search

import (
	"encoding/json"
	"fmt"
)

// Acl Access control list permissions.
type Acl string

// List of acl.
const (
	ACL_ADD_OBJECT                   Acl = "addObject"
	ACL_ANALYTICS                    Acl = "analytics"
	ACL_BROWSE                       Acl = "browse"
	ACL_DELETE_OBJECT                Acl = "deleteObject"
	ACL_DELETE_INDEX                 Acl = "deleteIndex"
	ACL_EDIT_SETTINGS                Acl = "editSettings"
	ACL_INFERENCE                    Acl = "inference"
	ACL_LIST_INDEXES                 Acl = "listIndexes"
	ACL_LOGS                         Acl = "logs"
	ACL_PERSONALIZATION              Acl = "personalization"
	ACL_RECOMMENDATION               Acl = "recommendation"
	ACL_SEARCH                       Acl = "search"
	ACL_SEE_UNRETRIEVABLE_ATTRIBUTES Acl = "seeUnretrievableAttributes"
	ACL_SETTINGS                     Acl = "settings"
	ACL_USAGE                        Acl = "usage"
)

// All allowed values of Acl enum.
var AllowedAclEnumValues = []Acl{
	"addObject",
	"analytics",
	"browse",
	"deleteObject",
	"deleteIndex",
	"editSettings",
	"inference",
	"listIndexes",
	"logs",
	"personalization",
	"recommendation",
	"search",
	"seeUnretrievableAttributes",
	"settings",
	"usage",
}

func (v *Acl) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return fmt.Errorf("failed to unmarshal value '%s' for enum 'Acl': %w", string(src), err)
	}
	enumTypeValue := Acl(value)
	for _, existing := range AllowedAclEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Acl", value)
}

// NewAclFromValue returns a pointer to a valid Acl
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAclFromValue(v string) (*Acl, error) {
	ev := Acl(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Acl: valid values are %v", v, AllowedAclEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v Acl) IsValid() bool {
	for _, existing := range AllowedAclEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to acl value.
func (v Acl) Ptr() *Acl {
	return &v
}

type NullableAcl struct {
	value *Acl
	isSet bool
}

func (v NullableAcl) Get() *Acl {
	return v.value
}

func (v *NullableAcl) Set(val *Acl) {
	v.value = val
	v.isSet = true
}

func (v NullableAcl) IsSet() bool {
	return v.isSet
}

func (v *NullableAcl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcl(val *Acl) *NullableAcl {
	return &NullableAcl{value: val, isSet: true}
}

func (v NullableAcl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableAcl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
