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

// checks if the AeAssetGatepassSendEmailDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AeAssetGatepassSendEmailDetail{}

// AeAssetGatepassSendEmailDetail Email asset gate pass to the provided email address(es)
type AeAssetGatepassSendEmailDetail struct {
	// Auction Edge unique identifier for an auction.
	AuctionId string `json:"auction-id"`
	Asset *CommonAmsAssetPointer `json:"asset,omitempty"`
	// The Vehicle Identification Number(VIN) of the asset.
	Vin string `json:"vin"`
	// The stock number of the asset.
	Stock string `json:"stock"`
	// The email address of the initiator of the gatepass email request.  Errors will be sent here.
	InitiatorEmail *string `json:"initiator-email,omitempty"`
	SendToEmail []string `json:"send-to-email"`
	GatepassType *GatepassType `json:"gatepass-type,omitempty"`
	// Optional name of the sender.
	SenderName *string `json:"senderName,omitempty"`
	Initiator *CommonInitiator `json:"initiator,omitempty"`
}

type _AeAssetGatepassSendEmailDetail AeAssetGatepassSendEmailDetail

// NewAeAssetGatepassSendEmailDetail instantiates a new AeAssetGatepassSendEmailDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAeAssetGatepassSendEmailDetail(auctionId string, vin string, stock string, sendToEmail []string) *AeAssetGatepassSendEmailDetail {
	this := AeAssetGatepassSendEmailDetail{}
	this.AuctionId = auctionId
	this.Vin = vin
	this.Stock = stock
	this.SendToEmail = sendToEmail
	var gatepassType GatepassType = G
	this.GatepassType = &gatepassType
	return &this
}

// NewAeAssetGatepassSendEmailDetailWithDefaults instantiates a new AeAssetGatepassSendEmailDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAeAssetGatepassSendEmailDetailWithDefaults() *AeAssetGatepassSendEmailDetail {
	this := AeAssetGatepassSendEmailDetail{}
	var gatepassType GatepassType = G
	this.GatepassType = &gatepassType
	return &this
}

// GetAuctionId returns the AuctionId field value
func (o *AeAssetGatepassSendEmailDetail) GetAuctionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuctionId
}

// GetAuctionIdOk returns a tuple with the AuctionId field value
// and a boolean to check if the value has been set.
func (o *AeAssetGatepassSendEmailDetail) GetAuctionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuctionId, true
}

// SetAuctionId sets field value
func (o *AeAssetGatepassSendEmailDetail) SetAuctionId(v string) {
	o.AuctionId = v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *AeAssetGatepassSendEmailDetail) GetAsset() CommonAmsAssetPointer {
	if o == nil || IsNil(o.Asset) {
		var ret CommonAmsAssetPointer
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAssetGatepassSendEmailDetail) GetAssetOk() (*CommonAmsAssetPointer, bool) {
	if o == nil || IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *AeAssetGatepassSendEmailDetail) HasAsset() bool {
	if o != nil && !IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given CommonAmsAssetPointer and assigns it to the Asset field.
func (o *AeAssetGatepassSendEmailDetail) SetAsset(v CommonAmsAssetPointer) {
	o.Asset = &v
}

// GetVin returns the Vin field value
func (o *AeAssetGatepassSendEmailDetail) GetVin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vin
}

// GetVinOk returns a tuple with the Vin field value
// and a boolean to check if the value has been set.
func (o *AeAssetGatepassSendEmailDetail) GetVinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vin, true
}

// SetVin sets field value
func (o *AeAssetGatepassSendEmailDetail) SetVin(v string) {
	o.Vin = v
}

// GetStock returns the Stock field value
func (o *AeAssetGatepassSendEmailDetail) GetStock() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Stock
}

// GetStockOk returns a tuple with the Stock field value
// and a boolean to check if the value has been set.
func (o *AeAssetGatepassSendEmailDetail) GetStockOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stock, true
}

// SetStock sets field value
func (o *AeAssetGatepassSendEmailDetail) SetStock(v string) {
	o.Stock = v
}

// GetInitiatorEmail returns the InitiatorEmail field value if set, zero value otherwise.
func (o *AeAssetGatepassSendEmailDetail) GetInitiatorEmail() string {
	if o == nil || IsNil(o.InitiatorEmail) {
		var ret string
		return ret
	}
	return *o.InitiatorEmail
}

// GetInitiatorEmailOk returns a tuple with the InitiatorEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAssetGatepassSendEmailDetail) GetInitiatorEmailOk() (*string, bool) {
	if o == nil || IsNil(o.InitiatorEmail) {
		return nil, false
	}
	return o.InitiatorEmail, true
}

// HasInitiatorEmail returns a boolean if a field has been set.
func (o *AeAssetGatepassSendEmailDetail) HasInitiatorEmail() bool {
	if o != nil && !IsNil(o.InitiatorEmail) {
		return true
	}

	return false
}

// SetInitiatorEmail gets a reference to the given string and assigns it to the InitiatorEmail field.
func (o *AeAssetGatepassSendEmailDetail) SetInitiatorEmail(v string) {
	o.InitiatorEmail = &v
}

// GetSendToEmail returns the SendToEmail field value
func (o *AeAssetGatepassSendEmailDetail) GetSendToEmail() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SendToEmail
}

// GetSendToEmailOk returns a tuple with the SendToEmail field value
// and a boolean to check if the value has been set.
func (o *AeAssetGatepassSendEmailDetail) GetSendToEmailOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SendToEmail, true
}

// SetSendToEmail sets field value
func (o *AeAssetGatepassSendEmailDetail) SetSendToEmail(v []string) {
	o.SendToEmail = v
}

// GetGatepassType returns the GatepassType field value if set, zero value otherwise.
func (o *AeAssetGatepassSendEmailDetail) GetGatepassType() GatepassType {
	if o == nil || IsNil(o.GatepassType) {
		var ret GatepassType
		return ret
	}
	return *o.GatepassType
}

// GetGatepassTypeOk returns a tuple with the GatepassType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAssetGatepassSendEmailDetail) GetGatepassTypeOk() (*GatepassType, bool) {
	if o == nil || IsNil(o.GatepassType) {
		return nil, false
	}
	return o.GatepassType, true
}

// HasGatepassType returns a boolean if a field has been set.
func (o *AeAssetGatepassSendEmailDetail) HasGatepassType() bool {
	if o != nil && !IsNil(o.GatepassType) {
		return true
	}

	return false
}

// SetGatepassType gets a reference to the given GatepassType and assigns it to the GatepassType field.
func (o *AeAssetGatepassSendEmailDetail) SetGatepassType(v GatepassType) {
	o.GatepassType = &v
}

// GetSenderName returns the SenderName field value if set, zero value otherwise.
func (o *AeAssetGatepassSendEmailDetail) GetSenderName() string {
	if o == nil || IsNil(o.SenderName) {
		var ret string
		return ret
	}
	return *o.SenderName
}

// GetSenderNameOk returns a tuple with the SenderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAssetGatepassSendEmailDetail) GetSenderNameOk() (*string, bool) {
	if o == nil || IsNil(o.SenderName) {
		return nil, false
	}
	return o.SenderName, true
}

// HasSenderName returns a boolean if a field has been set.
func (o *AeAssetGatepassSendEmailDetail) HasSenderName() bool {
	if o != nil && !IsNil(o.SenderName) {
		return true
	}

	return false
}

// SetSenderName gets a reference to the given string and assigns it to the SenderName field.
func (o *AeAssetGatepassSendEmailDetail) SetSenderName(v string) {
	o.SenderName = &v
}

// GetInitiator returns the Initiator field value if set, zero value otherwise.
func (o *AeAssetGatepassSendEmailDetail) GetInitiator() CommonInitiator {
	if o == nil || IsNil(o.Initiator) {
		var ret CommonInitiator
		return ret
	}
	return *o.Initiator
}

// GetInitiatorOk returns a tuple with the Initiator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAssetGatepassSendEmailDetail) GetInitiatorOk() (*CommonInitiator, bool) {
	if o == nil || IsNil(o.Initiator) {
		return nil, false
	}
	return o.Initiator, true
}

// HasInitiator returns a boolean if a field has been set.
func (o *AeAssetGatepassSendEmailDetail) HasInitiator() bool {
	if o != nil && !IsNil(o.Initiator) {
		return true
	}

	return false
}

// SetInitiator gets a reference to the given CommonInitiator and assigns it to the Initiator field.
func (o *AeAssetGatepassSendEmailDetail) SetInitiator(v CommonInitiator) {
	o.Initiator = &v
}

func (o AeAssetGatepassSendEmailDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AeAssetGatepassSendEmailDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["auction-id"] = o.AuctionId
	if !IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	toSerialize["vin"] = o.Vin
	toSerialize["stock"] = o.Stock
	if !IsNil(o.InitiatorEmail) {
		toSerialize["initiator-email"] = o.InitiatorEmail
	}
	toSerialize["send-to-email"] = o.SendToEmail
	if !IsNil(o.GatepassType) {
		toSerialize["gatepass-type"] = o.GatepassType
	}
	if !IsNil(o.SenderName) {
		toSerialize["senderName"] = o.SenderName
	}
	if !IsNil(o.Initiator) {
		toSerialize["initiator"] = o.Initiator
	}
	return toSerialize, nil
}

func (o *AeAssetGatepassSendEmailDetail) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"auction-id",
		"vin",
		"stock",
		"send-to-email",
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

	varAeAssetGatepassSendEmailDetail := _AeAssetGatepassSendEmailDetail{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAeAssetGatepassSendEmailDetail)

	if err != nil {
		return err
	}

	*o = AeAssetGatepassSendEmailDetail(varAeAssetGatepassSendEmailDetail)

	return err
}

type NullableAeAssetGatepassSendEmailDetail struct {
	value *AeAssetGatepassSendEmailDetail
	isSet bool
}

func (v NullableAeAssetGatepassSendEmailDetail) Get() *AeAssetGatepassSendEmailDetail {
	return v.value
}

func (v *NullableAeAssetGatepassSendEmailDetail) Set(val *AeAssetGatepassSendEmailDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableAeAssetGatepassSendEmailDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableAeAssetGatepassSendEmailDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAeAssetGatepassSendEmailDetail(val *AeAssetGatepassSendEmailDetail) *NullableAeAssetGatepassSendEmailDetail {
	return &NullableAeAssetGatepassSendEmailDetail{value: val, isSet: true}
}

func (v NullableAeAssetGatepassSendEmailDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAeAssetGatepassSendEmailDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


