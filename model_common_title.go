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
)

// checks if the CommonTitle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonTitle{}

// CommonTitle Attributes defining the title document for the asset.
type CommonTitle struct {
	// Defines whether or not the asset can even have a title.  For example a lawn mower would be title exempt.
	Exempt *bool `json:"exempt,omitempty"`
	// A unique identifier, typically a string of 7-8 digits or characters, found on your car's title document.
	Number *string `json:"number,omitempty"`
	// The state in which the asset's title was issued.
	State *string `json:"state,omitempty"`
	// The date and time at which the title document was received by the auction.
	ReceivedAt NullableTime `json:"received-at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CommonTitle CommonTitle

// NewCommonTitle instantiates a new CommonTitle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonTitle() *CommonTitle {
	this := CommonTitle{}
	var exempt bool = false
	this.Exempt = &exempt
	return &this
}

// NewCommonTitleWithDefaults instantiates a new CommonTitle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonTitleWithDefaults() *CommonTitle {
	this := CommonTitle{}
	var exempt bool = false
	this.Exempt = &exempt
	return &this
}

// GetExempt returns the Exempt field value if set, zero value otherwise.
func (o *CommonTitle) GetExempt() bool {
	if o == nil || IsNil(o.Exempt) {
		var ret bool
		return ret
	}
	return *o.Exempt
}

// GetExemptOk returns a tuple with the Exempt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonTitle) GetExemptOk() (*bool, bool) {
	if o == nil || IsNil(o.Exempt) {
		return nil, false
	}
	return o.Exempt, true
}

// HasExempt returns a boolean if a field has been set.
func (o *CommonTitle) HasExempt() bool {
	if o != nil && !IsNil(o.Exempt) {
		return true
	}

	return false
}

// SetExempt gets a reference to the given bool and assigns it to the Exempt field.
func (o *CommonTitle) SetExempt(v bool) {
	o.Exempt = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *CommonTitle) GetNumber() string {
	if o == nil || IsNil(o.Number) {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonTitle) GetNumberOk() (*string, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *CommonTitle) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *CommonTitle) SetNumber(v string) {
	o.Number = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CommonTitle) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonTitle) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CommonTitle) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *CommonTitle) SetState(v string) {
	o.State = &v
}

// GetReceivedAt returns the ReceivedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonTitle) GetReceivedAt() time.Time {
	if o == nil || IsNil(o.ReceivedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ReceivedAt.Get()
}

// GetReceivedAtOk returns a tuple with the ReceivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonTitle) GetReceivedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReceivedAt.Get(), o.ReceivedAt.IsSet()
}

// HasReceivedAt returns a boolean if a field has been set.
func (o *CommonTitle) HasReceivedAt() bool {
	if o != nil && o.ReceivedAt.IsSet() {
		return true
	}

	return false
}

// SetReceivedAt gets a reference to the given NullableTime and assigns it to the ReceivedAt field.
func (o *CommonTitle) SetReceivedAt(v time.Time) {
	o.ReceivedAt.Set(&v)
}
// SetReceivedAtNil sets the value for ReceivedAt to be an explicit nil
func (o *CommonTitle) SetReceivedAtNil() {
	o.ReceivedAt.Set(nil)
}

// UnsetReceivedAt ensures that no value is present for ReceivedAt, not even an explicit nil
func (o *CommonTitle) UnsetReceivedAt() {
	o.ReceivedAt.Unset()
}

func (o CommonTitle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonTitle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Exempt) {
		toSerialize["exempt"] = o.Exempt
	}
	if !IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if o.ReceivedAt.IsSet() {
		toSerialize["received-at"] = o.ReceivedAt.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CommonTitle) UnmarshalJSON(data []byte) (err error) {
	varCommonTitle := _CommonTitle{}

	err = json.Unmarshal(data, &varCommonTitle)

	if err != nil {
		return err
	}

	*o = CommonTitle(varCommonTitle)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "exempt")
		delete(additionalProperties, "number")
		delete(additionalProperties, "state")
		delete(additionalProperties, "received-at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCommonTitle struct {
	value *CommonTitle
	isSet bool
}

func (v NullableCommonTitle) Get() *CommonTitle {
	return v.value
}

func (v *NullableCommonTitle) Set(val *CommonTitle) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonTitle) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonTitle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonTitle(val *CommonTitle) *NullableCommonTitle {
	return &NullableCommonTitle{value: val, isSet: true}
}

func (v NullableCommonTitle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonTitle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


