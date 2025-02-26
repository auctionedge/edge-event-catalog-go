/*
Edge Event Schemas

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package events

import (
	"encoding/json"
)

// checks if the CommonServiceParameter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonServiceParameter{}

// CommonServiceParameter struct for CommonServiceParameter
type CommonServiceParameter struct {
	// Service parameter name
	Name *string `json:"name,omitempty"`
	// Service parameter value checked against
	Value *string `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CommonServiceParameter CommonServiceParameter

// NewCommonServiceParameter instantiates a new CommonServiceParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonServiceParameter() *CommonServiceParameter {
	this := CommonServiceParameter{}
	return &this
}

// NewCommonServiceParameterWithDefaults instantiates a new CommonServiceParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonServiceParameterWithDefaults() *CommonServiceParameter {
	this := CommonServiceParameter{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CommonServiceParameter) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonServiceParameter) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CommonServiceParameter) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CommonServiceParameter) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CommonServiceParameter) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonServiceParameter) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CommonServiceParameter) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *CommonServiceParameter) SetValue(v string) {
	o.Value = &v
}

func (o CommonServiceParameter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonServiceParameter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CommonServiceParameter) UnmarshalJSON(data []byte) (err error) {
	varCommonServiceParameter := _CommonServiceParameter{}

	err = json.Unmarshal(data, &varCommonServiceParameter)

	if err != nil {
		return err
	}

	*o = CommonServiceParameter(varCommonServiceParameter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCommonServiceParameter struct {
	value *CommonServiceParameter
	isSet bool
}

func (v NullableCommonServiceParameter) Get() *CommonServiceParameter {
	return v.value
}

func (v *NullableCommonServiceParameter) Set(val *CommonServiceParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonServiceParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonServiceParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonServiceParameter(val *CommonServiceParameter) *NullableCommonServiceParameter {
	return &NullableCommonServiceParameter{value: val, isSet: true}
}

func (v NullableCommonServiceParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonServiceParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


