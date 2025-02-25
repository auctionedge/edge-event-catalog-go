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

// checks if the AeAssetDealPaymentStatusChangedDetailDealPayment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AeAssetDealPaymentStatusChangedDetailDealPayment{}

// AeAssetDealPaymentStatusChangedDetailDealPayment struct for AeAssetDealPaymentStatusChangedDetailDealPayment
type AeAssetDealPaymentStatusChangedDetailDealPayment struct {
	Status string `json:"status"`
	// The time at which the status changed
	UpdatedAt time.Time `json:"updated-at"`
	AdditionalProperties map[string]interface{}
}

type _AeAssetDealPaymentStatusChangedDetailDealPayment AeAssetDealPaymentStatusChangedDetailDealPayment

// NewAeAssetDealPaymentStatusChangedDetailDealPayment instantiates a new AeAssetDealPaymentStatusChangedDetailDealPayment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAeAssetDealPaymentStatusChangedDetailDealPayment(status string, updatedAt time.Time) *AeAssetDealPaymentStatusChangedDetailDealPayment {
	this := AeAssetDealPaymentStatusChangedDetailDealPayment{}
	this.Status = status
	this.UpdatedAt = updatedAt
	return &this
}

// NewAeAssetDealPaymentStatusChangedDetailDealPaymentWithDefaults instantiates a new AeAssetDealPaymentStatusChangedDetailDealPayment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAeAssetDealPaymentStatusChangedDetailDealPaymentWithDefaults() *AeAssetDealPaymentStatusChangedDetailDealPayment {
	this := AeAssetDealPaymentStatusChangedDetailDealPayment{}
	return &this
}

// GetStatus returns the Status field value
func (o *AeAssetDealPaymentStatusChangedDetailDealPayment) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AeAssetDealPaymentStatusChangedDetailDealPayment) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AeAssetDealPaymentStatusChangedDetailDealPayment) SetStatus(v string) {
	o.Status = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *AeAssetDealPaymentStatusChangedDetailDealPayment) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *AeAssetDealPaymentStatusChangedDetailDealPayment) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *AeAssetDealPaymentStatusChangedDetailDealPayment) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o AeAssetDealPaymentStatusChangedDetailDealPayment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AeAssetDealPaymentStatusChangedDetailDealPayment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["updated-at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AeAssetDealPaymentStatusChangedDetailDealPayment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"updated-at",
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

	varAeAssetDealPaymentStatusChangedDetailDealPayment := _AeAssetDealPaymentStatusChangedDetailDealPayment{}

	err = json.Unmarshal(data, &varAeAssetDealPaymentStatusChangedDetailDealPayment)

	if err != nil {
		return err
	}

	*o = AeAssetDealPaymentStatusChangedDetailDealPayment(varAeAssetDealPaymentStatusChangedDetailDealPayment)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "updated-at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAeAssetDealPaymentStatusChangedDetailDealPayment struct {
	value *AeAssetDealPaymentStatusChangedDetailDealPayment
	isSet bool
}

func (v NullableAeAssetDealPaymentStatusChangedDetailDealPayment) Get() *AeAssetDealPaymentStatusChangedDetailDealPayment {
	return v.value
}

func (v *NullableAeAssetDealPaymentStatusChangedDetailDealPayment) Set(val *AeAssetDealPaymentStatusChangedDetailDealPayment) {
	v.value = val
	v.isSet = true
}

func (v NullableAeAssetDealPaymentStatusChangedDetailDealPayment) IsSet() bool {
	return v.isSet
}

func (v *NullableAeAssetDealPaymentStatusChangedDetailDealPayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAeAssetDealPaymentStatusChangedDetailDealPayment(val *AeAssetDealPaymentStatusChangedDetailDealPayment) *NullableAeAssetDealPaymentStatusChangedDetailDealPayment {
	return &NullableAeAssetDealPaymentStatusChangedDetailDealPayment{value: val, isSet: true}
}

func (v NullableAeAssetDealPaymentStatusChangedDetailDealPayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAeAssetDealPaymentStatusChangedDetailDealPayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


