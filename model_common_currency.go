/*
Edge Event Schemas

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package events

import (
	"encoding/json"
	"fmt"
)

// checks if the CommonCurrency type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonCurrency{}

// CommonCurrency struct for CommonCurrency
type CommonCurrency struct {
	// Amount of money in cents so $175.00 would be 17500 cents.
	Cents int32 `json:"cents"`
	CurrencyCode string `json:"currency-code"`
	AdditionalProperties map[string]interface{}
}

type _CommonCurrency CommonCurrency

// NewCommonCurrency instantiates a new CommonCurrency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonCurrency(cents int32, currencyCode string) *CommonCurrency {
	this := CommonCurrency{}
	this.Cents = cents
	this.CurrencyCode = currencyCode
	return &this
}

// NewCommonCurrencyWithDefaults instantiates a new CommonCurrency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonCurrencyWithDefaults() *CommonCurrency {
	this := CommonCurrency{}
	return &this
}

// GetCents returns the Cents field value
func (o *CommonCurrency) GetCents() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Cents
}

// GetCentsOk returns a tuple with the Cents field value
// and a boolean to check if the value has been set.
func (o *CommonCurrency) GetCentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cents, true
}

// SetCents sets field value
func (o *CommonCurrency) SetCents(v int32) {
	o.Cents = v
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *CommonCurrency) GetCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *CommonCurrency) GetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *CommonCurrency) SetCurrencyCode(v string) {
	o.CurrencyCode = v
}

func (o CommonCurrency) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonCurrency) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cents"] = o.Cents
	toSerialize["currency-code"] = o.CurrencyCode

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CommonCurrency) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cents",
		"currency-code",
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

	varCommonCurrency := _CommonCurrency{}

	err = json.Unmarshal(data, &varCommonCurrency)

	if err != nil {
		return err
	}

	*o = CommonCurrency(varCommonCurrency)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cents")
		delete(additionalProperties, "currency-code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCommonCurrency struct {
	value *CommonCurrency
	isSet bool
}

func (v NullableCommonCurrency) Get() *CommonCurrency {
	return v.value
}

func (v *NullableCommonCurrency) Set(val *CommonCurrency) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonCurrency) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonCurrency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonCurrency(val *CommonCurrency) *NullableCommonCurrency {
	return &NullableCommonCurrency{value: val, isSet: true}
}

func (v NullableCommonCurrency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonCurrency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


