/*
Edge Event Schemas

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package events

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the CommonAssetBuyNow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonAssetBuyNow{}

// CommonAssetBuyNow An object representation for an asset's buy now properties
type CommonAssetBuyNow struct {
	BuyNowPrice CommonCurrency `json:"buy-now-price"`
	// The time that this buy-now price is available for the asset
	BuyNowValidAt time.Time `json:"buy-now-valid-at"`
	// The time that this buy-now price for the asset expires.
	BuyNowValidUntil time.Time `json:"buy-now-valid-until"`
	AdditionalProperties map[string]interface{}
}

type _CommonAssetBuyNow CommonAssetBuyNow

// NewCommonAssetBuyNow instantiates a new CommonAssetBuyNow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonAssetBuyNow(buyNowPrice CommonCurrency, buyNowValidAt time.Time, buyNowValidUntil time.Time) *CommonAssetBuyNow {
	this := CommonAssetBuyNow{}
	this.BuyNowPrice = buyNowPrice
	this.BuyNowValidAt = buyNowValidAt
	this.BuyNowValidUntil = buyNowValidUntil
	return &this
}

// NewCommonAssetBuyNowWithDefaults instantiates a new CommonAssetBuyNow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonAssetBuyNowWithDefaults() *CommonAssetBuyNow {
	this := CommonAssetBuyNow{}
	return &this
}

// GetBuyNowPrice returns the BuyNowPrice field value
func (o *CommonAssetBuyNow) GetBuyNowPrice() CommonCurrency {
	if o == nil {
		var ret CommonCurrency
		return ret
	}

	return o.BuyNowPrice
}

// GetBuyNowPriceOk returns a tuple with the BuyNowPrice field value
// and a boolean to check if the value has been set.
func (o *CommonAssetBuyNow) GetBuyNowPriceOk() (*CommonCurrency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BuyNowPrice, true
}

// SetBuyNowPrice sets field value
func (o *CommonAssetBuyNow) SetBuyNowPrice(v CommonCurrency) {
	o.BuyNowPrice = v
}

// GetBuyNowValidAt returns the BuyNowValidAt field value
func (o *CommonAssetBuyNow) GetBuyNowValidAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.BuyNowValidAt
}

// GetBuyNowValidAtOk returns a tuple with the BuyNowValidAt field value
// and a boolean to check if the value has been set.
func (o *CommonAssetBuyNow) GetBuyNowValidAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BuyNowValidAt, true
}

// SetBuyNowValidAt sets field value
func (o *CommonAssetBuyNow) SetBuyNowValidAt(v time.Time) {
	o.BuyNowValidAt = v
}

// GetBuyNowValidUntil returns the BuyNowValidUntil field value
func (o *CommonAssetBuyNow) GetBuyNowValidUntil() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.BuyNowValidUntil
}

// GetBuyNowValidUntilOk returns a tuple with the BuyNowValidUntil field value
// and a boolean to check if the value has been set.
func (o *CommonAssetBuyNow) GetBuyNowValidUntilOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BuyNowValidUntil, true
}

// SetBuyNowValidUntil sets field value
func (o *CommonAssetBuyNow) SetBuyNowValidUntil(v time.Time) {
	o.BuyNowValidUntil = v
}

func (o CommonAssetBuyNow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonAssetBuyNow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["buy-now-price"] = o.BuyNowPrice
	toSerialize["buy-now-valid-at"] = o.BuyNowValidAt
	toSerialize["buy-now-valid-until"] = o.BuyNowValidUntil

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CommonAssetBuyNow) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"buy-now-price",
		"buy-now-valid-at",
		"buy-now-valid-until",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCommonAssetBuyNow := _CommonAssetBuyNow{}

	err = json.Unmarshal(data, &varCommonAssetBuyNow)

	if err != nil {
		return err
	}

	*o = CommonAssetBuyNow(varCommonAssetBuyNow)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "buy-now-price")
		delete(additionalProperties, "buy-now-valid-at")
		delete(additionalProperties, "buy-now-valid-until")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCommonAssetBuyNow struct {
	value *CommonAssetBuyNow
	isSet bool
}

func (v NullableCommonAssetBuyNow) Get() *CommonAssetBuyNow {
	return v.value
}

func (v *NullableCommonAssetBuyNow) Set(val *CommonAssetBuyNow) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonAssetBuyNow) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonAssetBuyNow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonAssetBuyNow(val *CommonAssetBuyNow) *NullableCommonAssetBuyNow {
	return &NullableCommonAssetBuyNow{value: val, isSet: true}
}

func (v NullableCommonAssetBuyNow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonAssetBuyNow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


