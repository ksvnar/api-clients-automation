// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package recommend

import (
	"encoding/json"
	"fmt"
)

// TrendingItemsModel Trending items model.  Trending items are determined from the number of conversion events collected on them.
type TrendingItemsModel string

// List of trendingItemsModel.
const (
	TRENDINGITEMSMODEL_TRENDING_ITEMS TrendingItemsModel = "trending-items"
)

// All allowed values of TrendingItemsModel enum.
var AllowedTrendingItemsModelEnumValues = []TrendingItemsModel{
	"trending-items",
}

func (v *TrendingItemsModel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return fmt.Errorf("failed to unmarshal value '%s' for enum 'TrendingItemsModel': %w", string(src), err)
	}
	enumTypeValue := TrendingItemsModel(value)
	for _, existing := range AllowedTrendingItemsModelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TrendingItemsModel", value)
}

// NewTrendingItemsModelFromValue returns a pointer to a valid TrendingItemsModel
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTrendingItemsModelFromValue(v string) (*TrendingItemsModel, error) {
	ev := TrendingItemsModel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TrendingItemsModel: valid values are %v", v, AllowedTrendingItemsModelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TrendingItemsModel) IsValid() bool {
	for _, existing := range AllowedTrendingItemsModelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to trendingItemsModel value.
func (v TrendingItemsModel) Ptr() *TrendingItemsModel {
	return &v
}

type NullableTrendingItemsModel struct {
	value *TrendingItemsModel
	isSet bool
}

func (v NullableTrendingItemsModel) Get() *TrendingItemsModel {
	return v.value
}

func (v *NullableTrendingItemsModel) Set(val *TrendingItemsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTrendingItemsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTrendingItemsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrendingItemsModel(val *TrendingItemsModel) *NullableTrendingItemsModel {
	return &NullableTrendingItemsModel{value: val, isSet: true}
}

func (v NullableTrendingItemsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableTrendingItemsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
