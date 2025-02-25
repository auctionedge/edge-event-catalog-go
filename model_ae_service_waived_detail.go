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

// checks if the AeServiceWaivedDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AeServiceWaivedDetail{}

// AeServiceWaivedDetail Notification that a service has been waived
type AeServiceWaivedDetail struct {
	// Auction Edge unique identifier for an auction.
	AuctionId string `json:"auction-id"`
	// The Vehicle Identification Number(VIN) of the asset.
	Vin string `json:"vin"`
	// The stock number of the asset.
	Stock string `json:"stock"`
	Service AeServiceWaivedDetailService `json:"service"`
	Initiator *CommonInitiator `json:"initiator,omitempty"`
}

type _AeServiceWaivedDetail AeServiceWaivedDetail

// NewAeServiceWaivedDetail instantiates a new AeServiceWaivedDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAeServiceWaivedDetail(auctionId string, vin string, stock string, service AeServiceWaivedDetailService) *AeServiceWaivedDetail {
	this := AeServiceWaivedDetail{}
	this.AuctionId = auctionId
	this.Vin = vin
	this.Stock = stock
	this.Service = service
	return &this
}

// NewAeServiceWaivedDetailWithDefaults instantiates a new AeServiceWaivedDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAeServiceWaivedDetailWithDefaults() *AeServiceWaivedDetail {
	this := AeServiceWaivedDetail{}
	return &this
}

// GetAuctionId returns the AuctionId field value
func (o *AeServiceWaivedDetail) GetAuctionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuctionId
}

// GetAuctionIdOk returns a tuple with the AuctionId field value
// and a boolean to check if the value has been set.
func (o *AeServiceWaivedDetail) GetAuctionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuctionId, true
}

// SetAuctionId sets field value
func (o *AeServiceWaivedDetail) SetAuctionId(v string) {
	o.AuctionId = v
}

// GetVin returns the Vin field value
func (o *AeServiceWaivedDetail) GetVin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vin
}

// GetVinOk returns a tuple with the Vin field value
// and a boolean to check if the value has been set.
func (o *AeServiceWaivedDetail) GetVinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vin, true
}

// SetVin sets field value
func (o *AeServiceWaivedDetail) SetVin(v string) {
	o.Vin = v
}

// GetStock returns the Stock field value
func (o *AeServiceWaivedDetail) GetStock() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Stock
}

// GetStockOk returns a tuple with the Stock field value
// and a boolean to check if the value has been set.
func (o *AeServiceWaivedDetail) GetStockOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stock, true
}

// SetStock sets field value
func (o *AeServiceWaivedDetail) SetStock(v string) {
	o.Stock = v
}

// GetService returns the Service field value
func (o *AeServiceWaivedDetail) GetService() AeServiceWaivedDetailService {
	if o == nil {
		var ret AeServiceWaivedDetailService
		return ret
	}

	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *AeServiceWaivedDetail) GetServiceOk() (*AeServiceWaivedDetailService, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value
func (o *AeServiceWaivedDetail) SetService(v AeServiceWaivedDetailService) {
	o.Service = v
}

// GetInitiator returns the Initiator field value if set, zero value otherwise.
func (o *AeServiceWaivedDetail) GetInitiator() CommonInitiator {
	if o == nil || IsNil(o.Initiator) {
		var ret CommonInitiator
		return ret
	}
	return *o.Initiator
}

// GetInitiatorOk returns a tuple with the Initiator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeServiceWaivedDetail) GetInitiatorOk() (*CommonInitiator, bool) {
	if o == nil || IsNil(o.Initiator) {
		return nil, false
	}
	return o.Initiator, true
}

// HasInitiator returns a boolean if a field has been set.
func (o *AeServiceWaivedDetail) HasInitiator() bool {
	if o != nil && !IsNil(o.Initiator) {
		return true
	}

	return false
}

// SetInitiator gets a reference to the given CommonInitiator and assigns it to the Initiator field.
func (o *AeServiceWaivedDetail) SetInitiator(v CommonInitiator) {
	o.Initiator = &v
}

func (o AeServiceWaivedDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AeServiceWaivedDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["auction-id"] = o.AuctionId
	toSerialize["vin"] = o.Vin
	toSerialize["stock"] = o.Stock
	toSerialize["service"] = o.Service
	if !IsNil(o.Initiator) {
		toSerialize["initiator"] = o.Initiator
	}
	return toSerialize, nil
}

func (o *AeServiceWaivedDetail) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"auction-id",
		"vin",
		"stock",
		"service",
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

	varAeServiceWaivedDetail := _AeServiceWaivedDetail{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAeServiceWaivedDetail)

	if err != nil {
		return err
	}

	*o = AeServiceWaivedDetail(varAeServiceWaivedDetail)

	return err
}

type NullableAeServiceWaivedDetail struct {
	value *AeServiceWaivedDetail
	isSet bool
}

func (v NullableAeServiceWaivedDetail) Get() *AeServiceWaivedDetail {
	return v.value
}

func (v *NullableAeServiceWaivedDetail) Set(val *AeServiceWaivedDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableAeServiceWaivedDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableAeServiceWaivedDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAeServiceWaivedDetail(val *AeServiceWaivedDetail) *NullableAeServiceWaivedDetail {
	return &NullableAeServiceWaivedDetail{value: val, isSet: true}
}

func (v NullableAeServiceWaivedDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAeServiceWaivedDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


