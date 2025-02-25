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

// checks if the AeAssetSellerChargeRemoveFailedDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AeAssetSellerChargeRemoveFailedDetail{}

// AeAssetSellerChargeRemoveFailedDetail triggered when a seller charge fails to be removed for any reason.
type AeAssetSellerChargeRemoveFailedDetail struct {
	// Auction Edge unique identifier for an auction.
	AuctionId string `json:"auction-id"`
	Asset CommonAmsAssetPointer `json:"asset"`
	Seller CommonAmsAccountPointer `json:"seller"`
	// Unique identifier of the charge to be removed.
	ChargeId string `json:"charge-id"`
	// Unique identifier referencing a remove request event that requested the charge be removed
	RemoveRequestReferenceId NullableString `json:"remove-request-reference-id,omitempty"`
	// An explaination for why the charge could not be added.
	Reason string `json:"reason"`
	Initiator *CommonInitiator `json:"initiator,omitempty"`
}

type _AeAssetSellerChargeRemoveFailedDetail AeAssetSellerChargeRemoveFailedDetail

// NewAeAssetSellerChargeRemoveFailedDetail instantiates a new AeAssetSellerChargeRemoveFailedDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAeAssetSellerChargeRemoveFailedDetail(auctionId string, asset CommonAmsAssetPointer, seller CommonAmsAccountPointer, chargeId string, reason string) *AeAssetSellerChargeRemoveFailedDetail {
	this := AeAssetSellerChargeRemoveFailedDetail{}
	this.AuctionId = auctionId
	this.Asset = asset
	this.Seller = seller
	this.ChargeId = chargeId
	this.Reason = reason
	return &this
}

// NewAeAssetSellerChargeRemoveFailedDetailWithDefaults instantiates a new AeAssetSellerChargeRemoveFailedDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAeAssetSellerChargeRemoveFailedDetailWithDefaults() *AeAssetSellerChargeRemoveFailedDetail {
	this := AeAssetSellerChargeRemoveFailedDetail{}
	return &this
}

// GetAuctionId returns the AuctionId field value
func (o *AeAssetSellerChargeRemoveFailedDetail) GetAuctionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuctionId
}

// GetAuctionIdOk returns a tuple with the AuctionId field value
// and a boolean to check if the value has been set.
func (o *AeAssetSellerChargeRemoveFailedDetail) GetAuctionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuctionId, true
}

// SetAuctionId sets field value
func (o *AeAssetSellerChargeRemoveFailedDetail) SetAuctionId(v string) {
	o.AuctionId = v
}

// GetAsset returns the Asset field value
func (o *AeAssetSellerChargeRemoveFailedDetail) GetAsset() CommonAmsAssetPointer {
	if o == nil {
		var ret CommonAmsAssetPointer
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *AeAssetSellerChargeRemoveFailedDetail) GetAssetOk() (*CommonAmsAssetPointer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *AeAssetSellerChargeRemoveFailedDetail) SetAsset(v CommonAmsAssetPointer) {
	o.Asset = v
}

// GetSeller returns the Seller field value
func (o *AeAssetSellerChargeRemoveFailedDetail) GetSeller() CommonAmsAccountPointer {
	if o == nil {
		var ret CommonAmsAccountPointer
		return ret
	}

	return o.Seller
}

// GetSellerOk returns a tuple with the Seller field value
// and a boolean to check if the value has been set.
func (o *AeAssetSellerChargeRemoveFailedDetail) GetSellerOk() (*CommonAmsAccountPointer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Seller, true
}

// SetSeller sets field value
func (o *AeAssetSellerChargeRemoveFailedDetail) SetSeller(v CommonAmsAccountPointer) {
	o.Seller = v
}

// GetChargeId returns the ChargeId field value
func (o *AeAssetSellerChargeRemoveFailedDetail) GetChargeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChargeId
}

// GetChargeIdOk returns a tuple with the ChargeId field value
// and a boolean to check if the value has been set.
func (o *AeAssetSellerChargeRemoveFailedDetail) GetChargeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChargeId, true
}

// SetChargeId sets field value
func (o *AeAssetSellerChargeRemoveFailedDetail) SetChargeId(v string) {
	o.ChargeId = v
}

// GetRemoveRequestReferenceId returns the RemoveRequestReferenceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AeAssetSellerChargeRemoveFailedDetail) GetRemoveRequestReferenceId() string {
	if o == nil || IsNil(o.RemoveRequestReferenceId.Get()) {
		var ret string
		return ret
	}
	return *o.RemoveRequestReferenceId.Get()
}

// GetRemoveRequestReferenceIdOk returns a tuple with the RemoveRequestReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AeAssetSellerChargeRemoveFailedDetail) GetRemoveRequestReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemoveRequestReferenceId.Get(), o.RemoveRequestReferenceId.IsSet()
}

// HasRemoveRequestReferenceId returns a boolean if a field has been set.
func (o *AeAssetSellerChargeRemoveFailedDetail) HasRemoveRequestReferenceId() bool {
	if o != nil && o.RemoveRequestReferenceId.IsSet() {
		return true
	}

	return false
}

// SetRemoveRequestReferenceId gets a reference to the given NullableString and assigns it to the RemoveRequestReferenceId field.
func (o *AeAssetSellerChargeRemoveFailedDetail) SetRemoveRequestReferenceId(v string) {
	o.RemoveRequestReferenceId.Set(&v)
}
// SetRemoveRequestReferenceIdNil sets the value for RemoveRequestReferenceId to be an explicit nil
func (o *AeAssetSellerChargeRemoveFailedDetail) SetRemoveRequestReferenceIdNil() {
	o.RemoveRequestReferenceId.Set(nil)
}

// UnsetRemoveRequestReferenceId ensures that no value is present for RemoveRequestReferenceId, not even an explicit nil
func (o *AeAssetSellerChargeRemoveFailedDetail) UnsetRemoveRequestReferenceId() {
	o.RemoveRequestReferenceId.Unset()
}

// GetReason returns the Reason field value
func (o *AeAssetSellerChargeRemoveFailedDetail) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *AeAssetSellerChargeRemoveFailedDetail) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *AeAssetSellerChargeRemoveFailedDetail) SetReason(v string) {
	o.Reason = v
}

// GetInitiator returns the Initiator field value if set, zero value otherwise.
func (o *AeAssetSellerChargeRemoveFailedDetail) GetInitiator() CommonInitiator {
	if o == nil || IsNil(o.Initiator) {
		var ret CommonInitiator
		return ret
	}
	return *o.Initiator
}

// GetInitiatorOk returns a tuple with the Initiator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAssetSellerChargeRemoveFailedDetail) GetInitiatorOk() (*CommonInitiator, bool) {
	if o == nil || IsNil(o.Initiator) {
		return nil, false
	}
	return o.Initiator, true
}

// HasInitiator returns a boolean if a field has been set.
func (o *AeAssetSellerChargeRemoveFailedDetail) HasInitiator() bool {
	if o != nil && !IsNil(o.Initiator) {
		return true
	}

	return false
}

// SetInitiator gets a reference to the given CommonInitiator and assigns it to the Initiator field.
func (o *AeAssetSellerChargeRemoveFailedDetail) SetInitiator(v CommonInitiator) {
	o.Initiator = &v
}

func (o AeAssetSellerChargeRemoveFailedDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AeAssetSellerChargeRemoveFailedDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["auction-id"] = o.AuctionId
	toSerialize["asset"] = o.Asset
	toSerialize["seller"] = o.Seller
	toSerialize["charge-id"] = o.ChargeId
	if o.RemoveRequestReferenceId.IsSet() {
		toSerialize["remove-request-reference-id"] = o.RemoveRequestReferenceId.Get()
	}
	toSerialize["reason"] = o.Reason
	if !IsNil(o.Initiator) {
		toSerialize["initiator"] = o.Initiator
	}
	return toSerialize, nil
}

func (o *AeAssetSellerChargeRemoveFailedDetail) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"auction-id",
		"asset",
		"seller",
		"charge-id",
		"reason",
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

	varAeAssetSellerChargeRemoveFailedDetail := _AeAssetSellerChargeRemoveFailedDetail{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAeAssetSellerChargeRemoveFailedDetail)

	if err != nil {
		return err
	}

	*o = AeAssetSellerChargeRemoveFailedDetail(varAeAssetSellerChargeRemoveFailedDetail)

	return err
}

type NullableAeAssetSellerChargeRemoveFailedDetail struct {
	value *AeAssetSellerChargeRemoveFailedDetail
	isSet bool
}

func (v NullableAeAssetSellerChargeRemoveFailedDetail) Get() *AeAssetSellerChargeRemoveFailedDetail {
	return v.value
}

func (v *NullableAeAssetSellerChargeRemoveFailedDetail) Set(val *AeAssetSellerChargeRemoveFailedDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableAeAssetSellerChargeRemoveFailedDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableAeAssetSellerChargeRemoveFailedDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAeAssetSellerChargeRemoveFailedDetail(val *AeAssetSellerChargeRemoveFailedDetail) *NullableAeAssetSellerChargeRemoveFailedDetail {
	return &NullableAeAssetSellerChargeRemoveFailedDetail{value: val, isSet: true}
}

func (v NullableAeAssetSellerChargeRemoveFailedDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAeAssetSellerChargeRemoveFailedDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


