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

// checks if the AeAssetDealPaymentApplySourceDetailDealPayment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AeAssetDealPaymentApplySourceDetailDealPayment{}

// AeAssetDealPaymentApplySourceDetailDealPayment struct for AeAssetDealPaymentApplySourceDetailDealPayment
type AeAssetDealPaymentApplySourceDetailDealPayment struct {
	// Payment Source ID.
	SourceId string `json:"source-id"`
}

type _AeAssetDealPaymentApplySourceDetailDealPayment AeAssetDealPaymentApplySourceDetailDealPayment

// NewAeAssetDealPaymentApplySourceDetailDealPayment instantiates a new AeAssetDealPaymentApplySourceDetailDealPayment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAeAssetDealPaymentApplySourceDetailDealPayment(sourceId string) *AeAssetDealPaymentApplySourceDetailDealPayment {
	this := AeAssetDealPaymentApplySourceDetailDealPayment{}
	this.SourceId = sourceId
	return &this
}

// NewAeAssetDealPaymentApplySourceDetailDealPaymentWithDefaults instantiates a new AeAssetDealPaymentApplySourceDetailDealPayment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAeAssetDealPaymentApplySourceDetailDealPaymentWithDefaults() *AeAssetDealPaymentApplySourceDetailDealPayment {
	this := AeAssetDealPaymentApplySourceDetailDealPayment{}
	return &this
}

// GetSourceId returns the SourceId field value
func (o *AeAssetDealPaymentApplySourceDetailDealPayment) GetSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value
// and a boolean to check if the value has been set.
func (o *AeAssetDealPaymentApplySourceDetailDealPayment) GetSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceId, true
}

// SetSourceId sets field value
func (o *AeAssetDealPaymentApplySourceDetailDealPayment) SetSourceId(v string) {
	o.SourceId = v
}

func (o AeAssetDealPaymentApplySourceDetailDealPayment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AeAssetDealPaymentApplySourceDetailDealPayment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source-id"] = o.SourceId
	return toSerialize, nil
}

func (o *AeAssetDealPaymentApplySourceDetailDealPayment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"source-id",
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

	varAeAssetDealPaymentApplySourceDetailDealPayment := _AeAssetDealPaymentApplySourceDetailDealPayment{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAeAssetDealPaymentApplySourceDetailDealPayment)

	if err != nil {
		return err
	}

	*o = AeAssetDealPaymentApplySourceDetailDealPayment(varAeAssetDealPaymentApplySourceDetailDealPayment)

	return err
}

type NullableAeAssetDealPaymentApplySourceDetailDealPayment struct {
	value *AeAssetDealPaymentApplySourceDetailDealPayment
	isSet bool
}

func (v NullableAeAssetDealPaymentApplySourceDetailDealPayment) Get() *AeAssetDealPaymentApplySourceDetailDealPayment {
	return v.value
}

func (v *NullableAeAssetDealPaymentApplySourceDetailDealPayment) Set(val *AeAssetDealPaymentApplySourceDetailDealPayment) {
	v.value = val
	v.isSet = true
}

func (v NullableAeAssetDealPaymentApplySourceDetailDealPayment) IsSet() bool {
	return v.isSet
}

func (v *NullableAeAssetDealPaymentApplySourceDetailDealPayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAeAssetDealPaymentApplySourceDetailDealPayment(val *AeAssetDealPaymentApplySourceDetailDealPayment) *NullableAeAssetDealPaymentApplySourceDetailDealPayment {
	return &NullableAeAssetDealPaymentApplySourceDetailDealPayment{value: val, isSet: true}
}

func (v NullableAeAssetDealPaymentApplySourceDetailDealPayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAeAssetDealPaymentApplySourceDetailDealPayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


