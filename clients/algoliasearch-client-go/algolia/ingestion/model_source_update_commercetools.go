// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// SourceUpdateCommercetools struct for SourceUpdateCommercetools.
type SourceUpdateCommercetools struct {
	StoreKeys []string `json:"storeKeys,omitempty"`
	// Locales for your commercetools stores.
	Locales []string `json:"locales,omitempty"`
	Url     *string  `json:"url,omitempty"`
	// Whether a fallback value is stored in the Algolia record if there's no inventory information about the product.
	FallbackIsInStockValue *bool                      `json:"fallbackIsInStockValue,omitempty"`
	CustomFields           *CommercetoolsCustomFields `json:"customFields,omitempty"`
}

type SourceUpdateCommercetoolsOption func(f *SourceUpdateCommercetools)

func WithSourceUpdateCommercetoolsStoreKeys(val []string) SourceUpdateCommercetoolsOption {
	return func(f *SourceUpdateCommercetools) {
		f.StoreKeys = val
	}
}

func WithSourceUpdateCommercetoolsLocales(val []string) SourceUpdateCommercetoolsOption {
	return func(f *SourceUpdateCommercetools) {
		f.Locales = val
	}
}

func WithSourceUpdateCommercetoolsUrl(val string) SourceUpdateCommercetoolsOption {
	return func(f *SourceUpdateCommercetools) {
		f.Url = &val
	}
}

func WithSourceUpdateCommercetoolsFallbackIsInStockValue(val bool) SourceUpdateCommercetoolsOption {
	return func(f *SourceUpdateCommercetools) {
		f.FallbackIsInStockValue = &val
	}
}

func WithSourceUpdateCommercetoolsCustomFields(val CommercetoolsCustomFields) SourceUpdateCommercetoolsOption {
	return func(f *SourceUpdateCommercetools) {
		f.CustomFields = &val
	}
}

// NewSourceUpdateCommercetools instantiates a new SourceUpdateCommercetools object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSourceUpdateCommercetools(opts ...SourceUpdateCommercetoolsOption) *SourceUpdateCommercetools {
	this := &SourceUpdateCommercetools{}
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptySourceUpdateCommercetools return a pointer to an empty SourceUpdateCommercetools object.
func NewEmptySourceUpdateCommercetools() *SourceUpdateCommercetools {
	return &SourceUpdateCommercetools{}
}

// GetStoreKeys returns the StoreKeys field value if set, zero value otherwise.
func (o *SourceUpdateCommercetools) GetStoreKeys() []string {
	if o == nil || o.StoreKeys == nil {
		var ret []string
		return ret
	}
	return o.StoreKeys
}

// GetStoreKeysOk returns a tuple with the StoreKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceUpdateCommercetools) GetStoreKeysOk() ([]string, bool) {
	if o == nil || o.StoreKeys == nil {
		return nil, false
	}
	return o.StoreKeys, true
}

// HasStoreKeys returns a boolean if a field has been set.
func (o *SourceUpdateCommercetools) HasStoreKeys() bool {
	if o != nil && o.StoreKeys != nil {
		return true
	}

	return false
}

// SetStoreKeys gets a reference to the given []string and assigns it to the StoreKeys field.
func (o *SourceUpdateCommercetools) SetStoreKeys(v []string) *SourceUpdateCommercetools {
	o.StoreKeys = v
	return o
}

// GetLocales returns the Locales field value if set, zero value otherwise.
func (o *SourceUpdateCommercetools) GetLocales() []string {
	if o == nil || o.Locales == nil {
		var ret []string
		return ret
	}
	return o.Locales
}

// GetLocalesOk returns a tuple with the Locales field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceUpdateCommercetools) GetLocalesOk() ([]string, bool) {
	if o == nil || o.Locales == nil {
		return nil, false
	}
	return o.Locales, true
}

// HasLocales returns a boolean if a field has been set.
func (o *SourceUpdateCommercetools) HasLocales() bool {
	if o != nil && o.Locales != nil {
		return true
	}

	return false
}

// SetLocales gets a reference to the given []string and assigns it to the Locales field.
func (o *SourceUpdateCommercetools) SetLocales(v []string) *SourceUpdateCommercetools {
	o.Locales = v
	return o
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SourceUpdateCommercetools) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceUpdateCommercetools) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SourceUpdateCommercetools) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SourceUpdateCommercetools) SetUrl(v string) *SourceUpdateCommercetools {
	o.Url = &v
	return o
}

// GetFallbackIsInStockValue returns the FallbackIsInStockValue field value if set, zero value otherwise.
func (o *SourceUpdateCommercetools) GetFallbackIsInStockValue() bool {
	if o == nil || o.FallbackIsInStockValue == nil {
		var ret bool
		return ret
	}
	return *o.FallbackIsInStockValue
}

// GetFallbackIsInStockValueOk returns a tuple with the FallbackIsInStockValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceUpdateCommercetools) GetFallbackIsInStockValueOk() (*bool, bool) {
	if o == nil || o.FallbackIsInStockValue == nil {
		return nil, false
	}
	return o.FallbackIsInStockValue, true
}

// HasFallbackIsInStockValue returns a boolean if a field has been set.
func (o *SourceUpdateCommercetools) HasFallbackIsInStockValue() bool {
	if o != nil && o.FallbackIsInStockValue != nil {
		return true
	}

	return false
}

// SetFallbackIsInStockValue gets a reference to the given bool and assigns it to the FallbackIsInStockValue field.
func (o *SourceUpdateCommercetools) SetFallbackIsInStockValue(v bool) *SourceUpdateCommercetools {
	o.FallbackIsInStockValue = &v
	return o
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *SourceUpdateCommercetools) GetCustomFields() CommercetoolsCustomFields {
	if o == nil || o.CustomFields == nil {
		var ret CommercetoolsCustomFields
		return ret
	}
	return *o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceUpdateCommercetools) GetCustomFieldsOk() (*CommercetoolsCustomFields, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *SourceUpdateCommercetools) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given CommercetoolsCustomFields and assigns it to the CustomFields field.
func (o *SourceUpdateCommercetools) SetCustomFields(v *CommercetoolsCustomFields) *SourceUpdateCommercetools {
	o.CustomFields = v
	return o
}

func (o SourceUpdateCommercetools) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.StoreKeys != nil {
		toSerialize["storeKeys"] = o.StoreKeys
	}
	if o.Locales != nil {
		toSerialize["locales"] = o.Locales
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.FallbackIsInStockValue != nil {
		toSerialize["fallbackIsInStockValue"] = o.FallbackIsInStockValue
	}
	if o.CustomFields != nil {
		toSerialize["customFields"] = o.CustomFields
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal SourceUpdateCommercetools: %w", err)
	}

	return serialized, nil
}

func (o SourceUpdateCommercetools) String() string {
	out := ""
	out += fmt.Sprintf("  storeKeys=%v\n", o.StoreKeys)
	out += fmt.Sprintf("  locales=%v\n", o.Locales)
	out += fmt.Sprintf("  url=%v\n", o.Url)
	out += fmt.Sprintf("  fallbackIsInStockValue=%v\n", o.FallbackIsInStockValue)
	out += fmt.Sprintf("  customFields=%v\n", o.CustomFields)
	return fmt.Sprintf("SourceUpdateCommercetools {\n%s}", out)
}

type NullableSourceUpdateCommercetools struct {
	value *SourceUpdateCommercetools
	isSet bool
}

func (v NullableSourceUpdateCommercetools) Get() *SourceUpdateCommercetools {
	return v.value
}

func (v *NullableSourceUpdateCommercetools) Set(val *SourceUpdateCommercetools) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceUpdateCommercetools) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceUpdateCommercetools) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceUpdateCommercetools(val *SourceUpdateCommercetools) *NullableSourceUpdateCommercetools {
	return &NullableSourceUpdateCommercetools{value: val, isSet: true}
}

func (v NullableSourceUpdateCommercetools) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableSourceUpdateCommercetools) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
