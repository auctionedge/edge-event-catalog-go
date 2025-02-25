/*
Edge Event Schemas

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package events

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CommonAmsPostalAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonAmsPostalAddress{}

// CommonAmsPostalAddress Postal address of entity
type CommonAmsPostalAddress struct {
	// Address line 1 of location
	Address1 string `json:"address1"`
	// Address line 2 of location
	Address2 *string `json:"address2,omitempty"`
	// Address line 2 of location
	Address3 *string `json:"address3,omitempty"`
	// City of location
	City string `json:"city"`
	// State of location
	State string `json:"state"`
	// Postal code of location
	PostalCode string `json:"postal-code"`
}

type _CommonAmsPostalAddress CommonAmsPostalAddress

// NewCommonAmsPostalAddress instantiates a new CommonAmsPostalAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonAmsPostalAddress(address1 string, city string, state string, postalCode string) *CommonAmsPostalAddress {
	this := CommonAmsPostalAddress{}
	this.Address1 = address1
	this.City = city
	this.State = state
	this.PostalCode = postalCode
	return &this
}

// NewCommonAmsPostalAddressWithDefaults instantiates a new CommonAmsPostalAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonAmsPostalAddressWithDefaults() *CommonAmsPostalAddress {
	this := CommonAmsPostalAddress{}
	return &this
}

// GetAddress1 returns the Address1 field value
func (o *CommonAmsPostalAddress) GetAddress1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address1
}

// GetAddress1Ok returns a tuple with the Address1 field value
// and a boolean to check if the value has been set.
func (o *CommonAmsPostalAddress) GetAddress1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address1, true
}

// SetAddress1 sets field value
func (o *CommonAmsPostalAddress) SetAddress1(v string) {
	o.Address1 = v
}

// GetAddress2 returns the Address2 field value if set, zero value otherwise.
func (o *CommonAmsPostalAddress) GetAddress2() string {
	if o == nil || IsNil(o.Address2) {
		var ret string
		return ret
	}
	return *o.Address2
}

// GetAddress2Ok returns a tuple with the Address2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAmsPostalAddress) GetAddress2Ok() (*string, bool) {
	if o == nil || IsNil(o.Address2) {
		return nil, false
	}
	return o.Address2, true
}

// HasAddress2 returns a boolean if a field has been set.
func (o *CommonAmsPostalAddress) HasAddress2() bool {
	if o != nil && !IsNil(o.Address2) {
		return true
	}

	return false
}

// SetAddress2 gets a reference to the given string and assigns it to the Address2 field.
func (o *CommonAmsPostalAddress) SetAddress2(v string) {
	o.Address2 = &v
}

// GetAddress3 returns the Address3 field value if set, zero value otherwise.
func (o *CommonAmsPostalAddress) GetAddress3() string {
	if o == nil || IsNil(o.Address3) {
		var ret string
		return ret
	}
	return *o.Address3
}

// GetAddress3Ok returns a tuple with the Address3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAmsPostalAddress) GetAddress3Ok() (*string, bool) {
	if o == nil || IsNil(o.Address3) {
		return nil, false
	}
	return o.Address3, true
}

// HasAddress3 returns a boolean if a field has been set.
func (o *CommonAmsPostalAddress) HasAddress3() bool {
	if o != nil && !IsNil(o.Address3) {
		return true
	}

	return false
}

// SetAddress3 gets a reference to the given string and assigns it to the Address3 field.
func (o *CommonAmsPostalAddress) SetAddress3(v string) {
	o.Address3 = &v
}

// GetCity returns the City field value
func (o *CommonAmsPostalAddress) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *CommonAmsPostalAddress) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *CommonAmsPostalAddress) SetCity(v string) {
	o.City = v
}

// GetState returns the State field value
func (o *CommonAmsPostalAddress) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *CommonAmsPostalAddress) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *CommonAmsPostalAddress) SetState(v string) {
	o.State = v
}

// GetPostalCode returns the PostalCode field value
func (o *CommonAmsPostalAddress) GetPostalCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value
// and a boolean to check if the value has been set.
func (o *CommonAmsPostalAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostalCode, true
}

// SetPostalCode sets field value
func (o *CommonAmsPostalAddress) SetPostalCode(v string) {
	o.PostalCode = v
}

func (o CommonAmsPostalAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonAmsPostalAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address1"] = o.Address1
	if !IsNil(o.Address2) {
		toSerialize["address2"] = o.Address2
	}
	if !IsNil(o.Address3) {
		toSerialize["address3"] = o.Address3
	}
	toSerialize["city"] = o.City
	toSerialize["state"] = o.State
	toSerialize["postal-code"] = o.PostalCode
	return toSerialize, nil
}

func (o *CommonAmsPostalAddress) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address1",
		"city",
		"state",
		"postal-code",
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

	varCommonAmsPostalAddress := _CommonAmsPostalAddress{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCommonAmsPostalAddress)

	if err != nil {
		return err
	}

	*o = CommonAmsPostalAddress(varCommonAmsPostalAddress)

	return err
}

type NullableCommonAmsPostalAddress struct {
	value *CommonAmsPostalAddress
	isSet bool
}

func (v NullableCommonAmsPostalAddress) Get() *CommonAmsPostalAddress {
	return v.value
}

func (v *NullableCommonAmsPostalAddress) Set(val *CommonAmsPostalAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonAmsPostalAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonAmsPostalAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonAmsPostalAddress(val *CommonAmsPostalAddress) *NullableCommonAmsPostalAddress {
	return &NullableCommonAmsPostalAddress{value: val, isSet: true}
}

func (v NullableCommonAmsPostalAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonAmsPostalAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


