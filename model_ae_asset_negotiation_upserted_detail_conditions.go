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

// checks if the AeAssetNegotiationUpsertedDetailConditions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AeAssetNegotiationUpsertedDetailConditions{}

// AeAssetNegotiationUpsertedDetailConditions struct for AeAssetNegotiationUpsertedDetailConditions
type AeAssetNegotiationUpsertedDetailConditions struct {
	Seller CommonAmsAccountPointer `json:"seller"`
	Buyer CommonAmsAccountPointer `json:"buyer"`
	// Light that was on the vehicle when negotiations started
	Lights string `json:"lights"`
	Announcements []string `json:"announcements"`
}

type _AeAssetNegotiationUpsertedDetailConditions AeAssetNegotiationUpsertedDetailConditions

// NewAeAssetNegotiationUpsertedDetailConditions instantiates a new AeAssetNegotiationUpsertedDetailConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAeAssetNegotiationUpsertedDetailConditions(seller CommonAmsAccountPointer, buyer CommonAmsAccountPointer, lights string, announcements []string) *AeAssetNegotiationUpsertedDetailConditions {
	this := AeAssetNegotiationUpsertedDetailConditions{}
	this.Seller = seller
	this.Buyer = buyer
	this.Lights = lights
	this.Announcements = announcements
	return &this
}

// NewAeAssetNegotiationUpsertedDetailConditionsWithDefaults instantiates a new AeAssetNegotiationUpsertedDetailConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAeAssetNegotiationUpsertedDetailConditionsWithDefaults() *AeAssetNegotiationUpsertedDetailConditions {
	this := AeAssetNegotiationUpsertedDetailConditions{}
	return &this
}

// GetSeller returns the Seller field value
func (o *AeAssetNegotiationUpsertedDetailConditions) GetSeller() CommonAmsAccountPointer {
	if o == nil {
		var ret CommonAmsAccountPointer
		return ret
	}

	return o.Seller
}

// GetSellerOk returns a tuple with the Seller field value
// and a boolean to check if the value has been set.
func (o *AeAssetNegotiationUpsertedDetailConditions) GetSellerOk() (*CommonAmsAccountPointer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Seller, true
}

// SetSeller sets field value
func (o *AeAssetNegotiationUpsertedDetailConditions) SetSeller(v CommonAmsAccountPointer) {
	o.Seller = v
}

// GetBuyer returns the Buyer field value
func (o *AeAssetNegotiationUpsertedDetailConditions) GetBuyer() CommonAmsAccountPointer {
	if o == nil {
		var ret CommonAmsAccountPointer
		return ret
	}

	return o.Buyer
}

// GetBuyerOk returns a tuple with the Buyer field value
// and a boolean to check if the value has been set.
func (o *AeAssetNegotiationUpsertedDetailConditions) GetBuyerOk() (*CommonAmsAccountPointer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Buyer, true
}

// SetBuyer sets field value
func (o *AeAssetNegotiationUpsertedDetailConditions) SetBuyer(v CommonAmsAccountPointer) {
	o.Buyer = v
}

// GetLights returns the Lights field value
func (o *AeAssetNegotiationUpsertedDetailConditions) GetLights() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Lights
}

// GetLightsOk returns a tuple with the Lights field value
// and a boolean to check if the value has been set.
func (o *AeAssetNegotiationUpsertedDetailConditions) GetLightsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Lights, true
}

// SetLights sets field value
func (o *AeAssetNegotiationUpsertedDetailConditions) SetLights(v string) {
	o.Lights = v
}

// GetAnnouncements returns the Announcements field value
func (o *AeAssetNegotiationUpsertedDetailConditions) GetAnnouncements() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Announcements
}

// GetAnnouncementsOk returns a tuple with the Announcements field value
// and a boolean to check if the value has been set.
func (o *AeAssetNegotiationUpsertedDetailConditions) GetAnnouncementsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Announcements, true
}

// SetAnnouncements sets field value
func (o *AeAssetNegotiationUpsertedDetailConditions) SetAnnouncements(v []string) {
	o.Announcements = v
}

func (o AeAssetNegotiationUpsertedDetailConditions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AeAssetNegotiationUpsertedDetailConditions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["seller"] = o.Seller
	toSerialize["buyer"] = o.Buyer
	toSerialize["lights"] = o.Lights
	toSerialize["announcements"] = o.Announcements
	return toSerialize, nil
}

func (o *AeAssetNegotiationUpsertedDetailConditions) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"seller",
		"buyer",
		"lights",
		"announcements",
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

	varAeAssetNegotiationUpsertedDetailConditions := _AeAssetNegotiationUpsertedDetailConditions{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAeAssetNegotiationUpsertedDetailConditions)

	if err != nil {
		return err
	}

	*o = AeAssetNegotiationUpsertedDetailConditions(varAeAssetNegotiationUpsertedDetailConditions)

	return err
}

type NullableAeAssetNegotiationUpsertedDetailConditions struct {
	value *AeAssetNegotiationUpsertedDetailConditions
	isSet bool
}

func (v NullableAeAssetNegotiationUpsertedDetailConditions) Get() *AeAssetNegotiationUpsertedDetailConditions {
	return v.value
}

func (v *NullableAeAssetNegotiationUpsertedDetailConditions) Set(val *AeAssetNegotiationUpsertedDetailConditions) {
	v.value = val
	v.isSet = true
}

func (v NullableAeAssetNegotiationUpsertedDetailConditions) IsSet() bool {
	return v.isSet
}

func (v *NullableAeAssetNegotiationUpsertedDetailConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAeAssetNegotiationUpsertedDetailConditions(val *AeAssetNegotiationUpsertedDetailConditions) *NullableAeAssetNegotiationUpsertedDetailConditions {
	return &NullableAeAssetNegotiationUpsertedDetailConditions{value: val, isSet: true}
}

func (v NullableAeAssetNegotiationUpsertedDetailConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAeAssetNegotiationUpsertedDetailConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


