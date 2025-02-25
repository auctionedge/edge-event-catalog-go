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

// checks if the AeAssetGatepassVoidedDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AeAssetGatepassVoidedDetail{}

// AeAssetGatepassVoidedDetail Notification that a gatepass document has been voided
type AeAssetGatepassVoidedDetail struct {
	// Auction Edge unique identifier for an auction.
	AuctionId string `json:"auction-id"`
	Asset *CommonAmsAssetPointer `json:"asset,omitempty"`
	// The Vehicle Identification Number(VIN) of the asset.
	Vin string `json:"vin"`
	// The stock number of the asset.
	Stock string `json:"stock"`
	Initiator *CommonInitiator `json:"initiator,omitempty"`
}

type _AeAssetGatepassVoidedDetail AeAssetGatepassVoidedDetail

// NewAeAssetGatepassVoidedDetail instantiates a new AeAssetGatepassVoidedDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAeAssetGatepassVoidedDetail(auctionId string, vin string, stock string) *AeAssetGatepassVoidedDetail {
	this := AeAssetGatepassVoidedDetail{}
	this.AuctionId = auctionId
	this.Vin = vin
	this.Stock = stock
	return &this
}

// NewAeAssetGatepassVoidedDetailWithDefaults instantiates a new AeAssetGatepassVoidedDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAeAssetGatepassVoidedDetailWithDefaults() *AeAssetGatepassVoidedDetail {
	this := AeAssetGatepassVoidedDetail{}
	return &this
}

// GetAuctionId returns the AuctionId field value
func (o *AeAssetGatepassVoidedDetail) GetAuctionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuctionId
}

// GetAuctionIdOk returns a tuple with the AuctionId field value
// and a boolean to check if the value has been set.
func (o *AeAssetGatepassVoidedDetail) GetAuctionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuctionId, true
}

// SetAuctionId sets field value
func (o *AeAssetGatepassVoidedDetail) SetAuctionId(v string) {
	o.AuctionId = v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *AeAssetGatepassVoidedDetail) GetAsset() CommonAmsAssetPointer {
	if o == nil || IsNil(o.Asset) {
		var ret CommonAmsAssetPointer
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAssetGatepassVoidedDetail) GetAssetOk() (*CommonAmsAssetPointer, bool) {
	if o == nil || IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *AeAssetGatepassVoidedDetail) HasAsset() bool {
	if o != nil && !IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given CommonAmsAssetPointer and assigns it to the Asset field.
func (o *AeAssetGatepassVoidedDetail) SetAsset(v CommonAmsAssetPointer) {
	o.Asset = &v
}

// GetVin returns the Vin field value
func (o *AeAssetGatepassVoidedDetail) GetVin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vin
}

// GetVinOk returns a tuple with the Vin field value
// and a boolean to check if the value has been set.
func (o *AeAssetGatepassVoidedDetail) GetVinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vin, true
}

// SetVin sets field value
func (o *AeAssetGatepassVoidedDetail) SetVin(v string) {
	o.Vin = v
}

// GetStock returns the Stock field value
func (o *AeAssetGatepassVoidedDetail) GetStock() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Stock
}

// GetStockOk returns a tuple with the Stock field value
// and a boolean to check if the value has been set.
func (o *AeAssetGatepassVoidedDetail) GetStockOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stock, true
}

// SetStock sets field value
func (o *AeAssetGatepassVoidedDetail) SetStock(v string) {
	o.Stock = v
}

// GetInitiator returns the Initiator field value if set, zero value otherwise.
func (o *AeAssetGatepassVoidedDetail) GetInitiator() CommonInitiator {
	if o == nil || IsNil(o.Initiator) {
		var ret CommonInitiator
		return ret
	}
	return *o.Initiator
}

// GetInitiatorOk returns a tuple with the Initiator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAssetGatepassVoidedDetail) GetInitiatorOk() (*CommonInitiator, bool) {
	if o == nil || IsNil(o.Initiator) {
		return nil, false
	}
	return o.Initiator, true
}

// HasInitiator returns a boolean if a field has been set.
func (o *AeAssetGatepassVoidedDetail) HasInitiator() bool {
	if o != nil && !IsNil(o.Initiator) {
		return true
	}

	return false
}

// SetInitiator gets a reference to the given CommonInitiator and assigns it to the Initiator field.
func (o *AeAssetGatepassVoidedDetail) SetInitiator(v CommonInitiator) {
	o.Initiator = &v
}

func (o AeAssetGatepassVoidedDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AeAssetGatepassVoidedDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["auction-id"] = o.AuctionId
	if !IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	toSerialize["vin"] = o.Vin
	toSerialize["stock"] = o.Stock
	if !IsNil(o.Initiator) {
		toSerialize["initiator"] = o.Initiator
	}
	return toSerialize, nil
}

func (o *AeAssetGatepassVoidedDetail) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"auction-id",
		"vin",
		"stock",
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

	varAeAssetGatepassVoidedDetail := _AeAssetGatepassVoidedDetail{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAeAssetGatepassVoidedDetail)

	if err != nil {
		return err
	}

	*o = AeAssetGatepassVoidedDetail(varAeAssetGatepassVoidedDetail)

	return err
}

type NullableAeAssetGatepassVoidedDetail struct {
	value *AeAssetGatepassVoidedDetail
	isSet bool
}

func (v NullableAeAssetGatepassVoidedDetail) Get() *AeAssetGatepassVoidedDetail {
	return v.value
}

func (v *NullableAeAssetGatepassVoidedDetail) Set(val *AeAssetGatepassVoidedDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableAeAssetGatepassVoidedDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableAeAssetGatepassVoidedDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAeAssetGatepassVoidedDetail(val *AeAssetGatepassVoidedDetail) *NullableAeAssetGatepassVoidedDetail {
	return &NullableAeAssetGatepassVoidedDetail{value: val, isSet: true}
}

func (v NullableAeAssetGatepassVoidedDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAeAssetGatepassVoidedDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


