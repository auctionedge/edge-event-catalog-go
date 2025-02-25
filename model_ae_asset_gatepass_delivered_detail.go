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

// checks if the AeAssetGatepassDeliveredDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AeAssetGatepassDeliveredDetail{}

// AeAssetGatepassDeliveredDetail Notification that a gatepass document has been delivered
type AeAssetGatepassDeliveredDetail struct {
	// Auction Edge unique identifier for an auction.
	AuctionId string `json:"auction-id"`
	Asset *CommonAmsAssetPointer `json:"asset,omitempty"`
	// The Vehicle Identification Number(VIN) of the asset.
	Vin string `json:"vin"`
	// The stock number of the asset.
	Stock string `json:"stock"`
	// Allowed values:     * 'email' - gatepass delivered via email, targets are the email addresses 
	DeliveryMethod string `json:"delivery-method"`
	Targets []string `json:"targets"`
	Initiator *CommonInitiator `json:"initiator,omitempty"`
}

type _AeAssetGatepassDeliveredDetail AeAssetGatepassDeliveredDetail

// NewAeAssetGatepassDeliveredDetail instantiates a new AeAssetGatepassDeliveredDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAeAssetGatepassDeliveredDetail(auctionId string, vin string, stock string, deliveryMethod string, targets []string) *AeAssetGatepassDeliveredDetail {
	this := AeAssetGatepassDeliveredDetail{}
	this.AuctionId = auctionId
	this.Vin = vin
	this.Stock = stock
	this.DeliveryMethod = deliveryMethod
	this.Targets = targets
	return &this
}

// NewAeAssetGatepassDeliveredDetailWithDefaults instantiates a new AeAssetGatepassDeliveredDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAeAssetGatepassDeliveredDetailWithDefaults() *AeAssetGatepassDeliveredDetail {
	this := AeAssetGatepassDeliveredDetail{}
	return &this
}

// GetAuctionId returns the AuctionId field value
func (o *AeAssetGatepassDeliveredDetail) GetAuctionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuctionId
}

// GetAuctionIdOk returns a tuple with the AuctionId field value
// and a boolean to check if the value has been set.
func (o *AeAssetGatepassDeliveredDetail) GetAuctionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuctionId, true
}

// SetAuctionId sets field value
func (o *AeAssetGatepassDeliveredDetail) SetAuctionId(v string) {
	o.AuctionId = v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *AeAssetGatepassDeliveredDetail) GetAsset() CommonAmsAssetPointer {
	if o == nil || IsNil(o.Asset) {
		var ret CommonAmsAssetPointer
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAssetGatepassDeliveredDetail) GetAssetOk() (*CommonAmsAssetPointer, bool) {
	if o == nil || IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *AeAssetGatepassDeliveredDetail) HasAsset() bool {
	if o != nil && !IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given CommonAmsAssetPointer and assigns it to the Asset field.
func (o *AeAssetGatepassDeliveredDetail) SetAsset(v CommonAmsAssetPointer) {
	o.Asset = &v
}

// GetVin returns the Vin field value
func (o *AeAssetGatepassDeliveredDetail) GetVin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vin
}

// GetVinOk returns a tuple with the Vin field value
// and a boolean to check if the value has been set.
func (o *AeAssetGatepassDeliveredDetail) GetVinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vin, true
}

// SetVin sets field value
func (o *AeAssetGatepassDeliveredDetail) SetVin(v string) {
	o.Vin = v
}

// GetStock returns the Stock field value
func (o *AeAssetGatepassDeliveredDetail) GetStock() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Stock
}

// GetStockOk returns a tuple with the Stock field value
// and a boolean to check if the value has been set.
func (o *AeAssetGatepassDeliveredDetail) GetStockOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stock, true
}

// SetStock sets field value
func (o *AeAssetGatepassDeliveredDetail) SetStock(v string) {
	o.Stock = v
}

// GetDeliveryMethod returns the DeliveryMethod field value
func (o *AeAssetGatepassDeliveredDetail) GetDeliveryMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeliveryMethod
}

// GetDeliveryMethodOk returns a tuple with the DeliveryMethod field value
// and a boolean to check if the value has been set.
func (o *AeAssetGatepassDeliveredDetail) GetDeliveryMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeliveryMethod, true
}

// SetDeliveryMethod sets field value
func (o *AeAssetGatepassDeliveredDetail) SetDeliveryMethod(v string) {
	o.DeliveryMethod = v
}

// GetTargets returns the Targets field value
func (o *AeAssetGatepassDeliveredDetail) GetTargets() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value
// and a boolean to check if the value has been set.
func (o *AeAssetGatepassDeliveredDetail) GetTargetsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Targets, true
}

// SetTargets sets field value
func (o *AeAssetGatepassDeliveredDetail) SetTargets(v []string) {
	o.Targets = v
}

// GetInitiator returns the Initiator field value if set, zero value otherwise.
func (o *AeAssetGatepassDeliveredDetail) GetInitiator() CommonInitiator {
	if o == nil || IsNil(o.Initiator) {
		var ret CommonInitiator
		return ret
	}
	return *o.Initiator
}

// GetInitiatorOk returns a tuple with the Initiator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAssetGatepassDeliveredDetail) GetInitiatorOk() (*CommonInitiator, bool) {
	if o == nil || IsNil(o.Initiator) {
		return nil, false
	}
	return o.Initiator, true
}

// HasInitiator returns a boolean if a field has been set.
func (o *AeAssetGatepassDeliveredDetail) HasInitiator() bool {
	if o != nil && !IsNil(o.Initiator) {
		return true
	}

	return false
}

// SetInitiator gets a reference to the given CommonInitiator and assigns it to the Initiator field.
func (o *AeAssetGatepassDeliveredDetail) SetInitiator(v CommonInitiator) {
	o.Initiator = &v
}

func (o AeAssetGatepassDeliveredDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AeAssetGatepassDeliveredDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["auction-id"] = o.AuctionId
	if !IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	toSerialize["vin"] = o.Vin
	toSerialize["stock"] = o.Stock
	toSerialize["delivery-method"] = o.DeliveryMethod
	toSerialize["targets"] = o.Targets
	if !IsNil(o.Initiator) {
		toSerialize["initiator"] = o.Initiator
	}
	return toSerialize, nil
}

func (o *AeAssetGatepassDeliveredDetail) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"auction-id",
		"vin",
		"stock",
		"delivery-method",
		"targets",
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

	varAeAssetGatepassDeliveredDetail := _AeAssetGatepassDeliveredDetail{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAeAssetGatepassDeliveredDetail)

	if err != nil {
		return err
	}

	*o = AeAssetGatepassDeliveredDetail(varAeAssetGatepassDeliveredDetail)

	return err
}

type NullableAeAssetGatepassDeliveredDetail struct {
	value *AeAssetGatepassDeliveredDetail
	isSet bool
}

func (v NullableAeAssetGatepassDeliveredDetail) Get() *AeAssetGatepassDeliveredDetail {
	return v.value
}

func (v *NullableAeAssetGatepassDeliveredDetail) Set(val *AeAssetGatepassDeliveredDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableAeAssetGatepassDeliveredDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableAeAssetGatepassDeliveredDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAeAssetGatepassDeliveredDetail(val *AeAssetGatepassDeliveredDetail) *NullableAeAssetGatepassDeliveredDetail {
	return &NullableAeAssetGatepassDeliveredDetail{value: val, isSet: true}
}

func (v NullableAeAssetGatepassDeliveredDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAeAssetGatepassDeliveredDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


