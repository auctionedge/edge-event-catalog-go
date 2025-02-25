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

// checks if the CommonAmsRepresentativePointer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonAmsRepresentativePointer{}

// CommonAmsRepresentativePointer Identifiers that reference an AMS representative.
type CommonAmsRepresentativePointer struct {
	// The representative id an AMS account
	Id string `json:"id"`
	// The Auction Access ID of the AMS Representative account.
	AaId NullableString `json:"aa-id"`
	// Rep number as denoted in ASI, AOS to use same value as representative-id
	RepNumber *float32 `json:"rep-number,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CommonAmsRepresentativePointer CommonAmsRepresentativePointer

// NewCommonAmsRepresentativePointer instantiates a new CommonAmsRepresentativePointer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonAmsRepresentativePointer(id string, aaId NullableString) *CommonAmsRepresentativePointer {
	this := CommonAmsRepresentativePointer{}
	this.Id = id
	this.AaId = aaId
	return &this
}

// NewCommonAmsRepresentativePointerWithDefaults instantiates a new CommonAmsRepresentativePointer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonAmsRepresentativePointerWithDefaults() *CommonAmsRepresentativePointer {
	this := CommonAmsRepresentativePointer{}
	return &this
}

// GetId returns the Id field value
func (o *CommonAmsRepresentativePointer) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CommonAmsRepresentativePointer) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CommonAmsRepresentativePointer) SetId(v string) {
	o.Id = v
}

// GetAaId returns the AaId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CommonAmsRepresentativePointer) GetAaId() string {
	if o == nil || o.AaId.Get() == nil {
		var ret string
		return ret
	}

	return *o.AaId.Get()
}

// GetAaIdOk returns a tuple with the AaId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonAmsRepresentativePointer) GetAaIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AaId.Get(), o.AaId.IsSet()
}

// SetAaId sets field value
func (o *CommonAmsRepresentativePointer) SetAaId(v string) {
	o.AaId.Set(&v)
}

// GetRepNumber returns the RepNumber field value if set, zero value otherwise.
func (o *CommonAmsRepresentativePointer) GetRepNumber() float32 {
	if o == nil || IsNil(o.RepNumber) {
		var ret float32
		return ret
	}
	return *o.RepNumber
}

// GetRepNumberOk returns a tuple with the RepNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAmsRepresentativePointer) GetRepNumberOk() (*float32, bool) {
	if o == nil || IsNil(o.RepNumber) {
		return nil, false
	}
	return o.RepNumber, true
}

// HasRepNumber returns a boolean if a field has been set.
func (o *CommonAmsRepresentativePointer) HasRepNumber() bool {
	if o != nil && !IsNil(o.RepNumber) {
		return true
	}

	return false
}

// SetRepNumber gets a reference to the given float32 and assigns it to the RepNumber field.
func (o *CommonAmsRepresentativePointer) SetRepNumber(v float32) {
	o.RepNumber = &v
}

func (o CommonAmsRepresentativePointer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonAmsRepresentativePointer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["aa-id"] = o.AaId.Get()
	if !IsNil(o.RepNumber) {
		toSerialize["rep-number"] = o.RepNumber
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CommonAmsRepresentativePointer) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"aa-id",
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

	varCommonAmsRepresentativePointer := _CommonAmsRepresentativePointer{}

	err = json.Unmarshal(data, &varCommonAmsRepresentativePointer)

	if err != nil {
		return err
	}

	*o = CommonAmsRepresentativePointer(varCommonAmsRepresentativePointer)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "aa-id")
		delete(additionalProperties, "rep-number")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCommonAmsRepresentativePointer struct {
	value *CommonAmsRepresentativePointer
	isSet bool
}

func (v NullableCommonAmsRepresentativePointer) Get() *CommonAmsRepresentativePointer {
	return v.value
}

func (v *NullableCommonAmsRepresentativePointer) Set(val *CommonAmsRepresentativePointer) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonAmsRepresentativePointer) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonAmsRepresentativePointer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonAmsRepresentativePointer(val *CommonAmsRepresentativePointer) *NullableCommonAmsRepresentativePointer {
	return &NullableCommonAmsRepresentativePointer{value: val, isSet: true}
}

func (v NullableCommonAmsRepresentativePointer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonAmsRepresentativePointer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


