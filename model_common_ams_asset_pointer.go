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

// checks if the CommonAmsAssetPointer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonAmsAssetPointer{}

// CommonAmsAssetPointer Identifiers that reference an asset.
type CommonAmsAssetPointer struct {
	// Source's unique identifier for asset
	Id string `json:"id"`
	// The stock number of the asset.
	Stock string `json:"stock"`
	// The Vehicle Identification Number(VIN) of the asset.
	Vin string `json:"vin"`
	AdditionalProperties map[string]interface{}
}

type _CommonAmsAssetPointer CommonAmsAssetPointer

// NewCommonAmsAssetPointer instantiates a new CommonAmsAssetPointer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonAmsAssetPointer(id string, stock string, vin string) *CommonAmsAssetPointer {
	this := CommonAmsAssetPointer{}
	this.Id = id
	this.Stock = stock
	this.Vin = vin
	return &this
}

// NewCommonAmsAssetPointerWithDefaults instantiates a new CommonAmsAssetPointer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonAmsAssetPointerWithDefaults() *CommonAmsAssetPointer {
	this := CommonAmsAssetPointer{}
	return &this
}

// GetId returns the Id field value
func (o *CommonAmsAssetPointer) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CommonAmsAssetPointer) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CommonAmsAssetPointer) SetId(v string) {
	o.Id = v
}

// GetStock returns the Stock field value
func (o *CommonAmsAssetPointer) GetStock() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Stock
}

// GetStockOk returns a tuple with the Stock field value
// and a boolean to check if the value has been set.
func (o *CommonAmsAssetPointer) GetStockOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stock, true
}

// SetStock sets field value
func (o *CommonAmsAssetPointer) SetStock(v string) {
	o.Stock = v
}

// GetVin returns the Vin field value
func (o *CommonAmsAssetPointer) GetVin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vin
}

// GetVinOk returns a tuple with the Vin field value
// and a boolean to check if the value has been set.
func (o *CommonAmsAssetPointer) GetVinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vin, true
}

// SetVin sets field value
func (o *CommonAmsAssetPointer) SetVin(v string) {
	o.Vin = v
}

func (o CommonAmsAssetPointer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonAmsAssetPointer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["stock"] = o.Stock
	toSerialize["vin"] = o.Vin

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CommonAmsAssetPointer) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"stock",
		"vin",
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

	varCommonAmsAssetPointer := _CommonAmsAssetPointer{}

	err = json.Unmarshal(data, &varCommonAmsAssetPointer)

	if err != nil {
		return err
	}

	*o = CommonAmsAssetPointer(varCommonAmsAssetPointer)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "stock")
		delete(additionalProperties, "vin")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCommonAmsAssetPointer struct {
	value *CommonAmsAssetPointer
	isSet bool
}

func (v NullableCommonAmsAssetPointer) Get() *CommonAmsAssetPointer {
	return v.value
}

func (v *NullableCommonAmsAssetPointer) Set(val *CommonAmsAssetPointer) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonAmsAssetPointer) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonAmsAssetPointer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonAmsAssetPointer(val *CommonAmsAssetPointer) *NullableCommonAmsAssetPointer {
	return &NullableCommonAmsAssetPointer{value: val, isSet: true}
}

func (v NullableCommonAmsAssetPointer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonAmsAssetPointer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


