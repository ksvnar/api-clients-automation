// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package insights

import (
	"encoding/json"
	"fmt"
)

// ObjectDataAfterSearch struct for ObjectDataAfterSearch.
type ObjectDataAfterSearch struct {
	// Unique identifier for a search query, used to track purchase events with multiple records that originate from different searches.
	QueryID *string `json:"queryID,omitempty"`
	Price   *Price  `json:"price,omitempty"`
	// Quantity of a product that has been purchased or added to the cart. The total purchase value is the sum of `quantity` multiplied with the `price` for each purchased item.
	Quantity *int32    `json:"quantity,omitempty"`
	Discount *Discount `json:"discount,omitempty"`
}

type ObjectDataAfterSearchOption func(f *ObjectDataAfterSearch)

func WithObjectDataAfterSearchQueryID(val string) ObjectDataAfterSearchOption {
	return func(f *ObjectDataAfterSearch) {
		f.QueryID = &val
	}
}

func WithObjectDataAfterSearchPrice(val Price) ObjectDataAfterSearchOption {
	return func(f *ObjectDataAfterSearch) {
		f.Price = &val
	}
}

func WithObjectDataAfterSearchQuantity(val int32) ObjectDataAfterSearchOption {
	return func(f *ObjectDataAfterSearch) {
		f.Quantity = &val
	}
}

func WithObjectDataAfterSearchDiscount(val Discount) ObjectDataAfterSearchOption {
	return func(f *ObjectDataAfterSearch) {
		f.Discount = &val
	}
}

// NewObjectDataAfterSearch instantiates a new ObjectDataAfterSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObjectDataAfterSearch(opts ...ObjectDataAfterSearchOption) *ObjectDataAfterSearch {
	this := &ObjectDataAfterSearch{}
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptyObjectDataAfterSearch return a pointer to an empty ObjectDataAfterSearch object.
func NewEmptyObjectDataAfterSearch() *ObjectDataAfterSearch {
	return &ObjectDataAfterSearch{}
}

// GetQueryID returns the QueryID field value if set, zero value otherwise.
func (o *ObjectDataAfterSearch) GetQueryID() string {
	if o == nil || o.QueryID == nil {
		var ret string
		return ret
	}
	return *o.QueryID
}

// GetQueryIDOk returns a tuple with the QueryID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectDataAfterSearch) GetQueryIDOk() (*string, bool) {
	if o == nil || o.QueryID == nil {
		return nil, false
	}
	return o.QueryID, true
}

// HasQueryID returns a boolean if a field has been set.
func (o *ObjectDataAfterSearch) HasQueryID() bool {
	if o != nil && o.QueryID != nil {
		return true
	}

	return false
}

// SetQueryID gets a reference to the given string and assigns it to the QueryID field.
func (o *ObjectDataAfterSearch) SetQueryID(v string) *ObjectDataAfterSearch {
	o.QueryID = &v
	return o
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *ObjectDataAfterSearch) GetPrice() Price {
	if o == nil || o.Price == nil {
		var ret Price
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectDataAfterSearch) GetPriceOk() (*Price, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *ObjectDataAfterSearch) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given Price and assigns it to the Price field.
func (o *ObjectDataAfterSearch) SetPrice(v *Price) *ObjectDataAfterSearch {
	o.Price = v
	return o
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *ObjectDataAfterSearch) GetQuantity() int32 {
	if o == nil || o.Quantity == nil {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectDataAfterSearch) GetQuantityOk() (*int32, bool) {
	if o == nil || o.Quantity == nil {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *ObjectDataAfterSearch) HasQuantity() bool {
	if o != nil && o.Quantity != nil {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *ObjectDataAfterSearch) SetQuantity(v int32) *ObjectDataAfterSearch {
	o.Quantity = &v
	return o
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *ObjectDataAfterSearch) GetDiscount() Discount {
	if o == nil || o.Discount == nil {
		var ret Discount
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectDataAfterSearch) GetDiscountOk() (*Discount, bool) {
	if o == nil || o.Discount == nil {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *ObjectDataAfterSearch) HasDiscount() bool {
	if o != nil && o.Discount != nil {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given Discount and assigns it to the Discount field.
func (o *ObjectDataAfterSearch) SetDiscount(v *Discount) *ObjectDataAfterSearch {
	o.Discount = v
	return o
}

func (o ObjectDataAfterSearch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.QueryID != nil {
		toSerialize["queryID"] = o.QueryID
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	if o.Discount != nil {
		toSerialize["discount"] = o.Discount
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal ObjectDataAfterSearch: %w", err)
	}

	return serialized, nil
}

func (o ObjectDataAfterSearch) String() string {
	out := ""
	out += fmt.Sprintf("  queryID=%v\n", o.QueryID)
	out += fmt.Sprintf("  price=%v\n", o.Price)
	out += fmt.Sprintf("  quantity=%v\n", o.Quantity)
	out += fmt.Sprintf("  discount=%v\n", o.Discount)
	return fmt.Sprintf("ObjectDataAfterSearch {\n%s}", out)
}

type NullableObjectDataAfterSearch struct {
	value *ObjectDataAfterSearch
	isSet bool
}

func (v NullableObjectDataAfterSearch) Get() *ObjectDataAfterSearch {
	return v.value
}

func (v *NullableObjectDataAfterSearch) Set(val *ObjectDataAfterSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectDataAfterSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectDataAfterSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectDataAfterSearch(val *ObjectDataAfterSearch) *NullableObjectDataAfterSearch {
	return &NullableObjectDataAfterSearch{value: val, isSet: true}
}

func (v NullableObjectDataAfterSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableObjectDataAfterSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
