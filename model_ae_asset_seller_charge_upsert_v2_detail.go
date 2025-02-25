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

// checks if the AeAssetSellerChargeUpsertV2Detail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AeAssetSellerChargeUpsertV2Detail{}

// AeAssetSellerChargeUpsertV2Detail An event denoting an insert or update of a charge on a seller's asset.
type AeAssetSellerChargeUpsertV2Detail struct {
	// Auction Edge unique identifier for an auction.
	AuctionId string `json:"auction-id"`
	Asset CommonAmsAssetPointer `json:"asset"`
	Seller CommonAmsAccountPointer `json:"seller"`
	// Unique identifier referencing the add request event that created the charge
	AddRequestReferenceId NullableString `json:"add-request-reference-id,omitempty"`
	Charge CommonCharge `json:"charge"`
	Initiator *CommonInitiator `json:"initiator,omitempty"`
}

type _AeAssetSellerChargeUpsertV2Detail AeAssetSellerChargeUpsertV2Detail

// NewAeAssetSellerChargeUpsertV2Detail instantiates a new AeAssetSellerChargeUpsertV2Detail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAeAssetSellerChargeUpsertV2Detail(auctionId string, asset CommonAmsAssetPointer, seller CommonAmsAccountPointer, charge CommonCharge) *AeAssetSellerChargeUpsertV2Detail {
	this := AeAssetSellerChargeUpsertV2Detail{}
	this.AuctionId = auctionId
	this.Asset = asset
	this.Seller = seller
	this.Charge = charge
	return &this
}

// NewAeAssetSellerChargeUpsertV2DetailWithDefaults instantiates a new AeAssetSellerChargeUpsertV2Detail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAeAssetSellerChargeUpsertV2DetailWithDefaults() *AeAssetSellerChargeUpsertV2Detail {
	this := AeAssetSellerChargeUpsertV2Detail{}
	return &this
}

// GetAuctionId returns the AuctionId field value
func (o *AeAssetSellerChargeUpsertV2Detail) GetAuctionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuctionId
}

// GetAuctionIdOk returns a tuple with the AuctionId field value
// and a boolean to check if the value has been set.
func (o *AeAssetSellerChargeUpsertV2Detail) GetAuctionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuctionId, true
}

// SetAuctionId sets field value
func (o *AeAssetSellerChargeUpsertV2Detail) SetAuctionId(v string) {
	o.AuctionId = v
}

// GetAsset returns the Asset field value
func (o *AeAssetSellerChargeUpsertV2Detail) GetAsset() CommonAmsAssetPointer {
	if o == nil {
		var ret CommonAmsAssetPointer
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *AeAssetSellerChargeUpsertV2Detail) GetAssetOk() (*CommonAmsAssetPointer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *AeAssetSellerChargeUpsertV2Detail) SetAsset(v CommonAmsAssetPointer) {
	o.Asset = v
}

// GetSeller returns the Seller field value
func (o *AeAssetSellerChargeUpsertV2Detail) GetSeller() CommonAmsAccountPointer {
	if o == nil {
		var ret CommonAmsAccountPointer
		return ret
	}

	return o.Seller
}

// GetSellerOk returns a tuple with the Seller field value
// and a boolean to check if the value has been set.
func (o *AeAssetSellerChargeUpsertV2Detail) GetSellerOk() (*CommonAmsAccountPointer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Seller, true
}

// SetSeller sets field value
func (o *AeAssetSellerChargeUpsertV2Detail) SetSeller(v CommonAmsAccountPointer) {
	o.Seller = v
}

// GetAddRequestReferenceId returns the AddRequestReferenceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AeAssetSellerChargeUpsertV2Detail) GetAddRequestReferenceId() string {
	if o == nil || IsNil(o.AddRequestReferenceId.Get()) {
		var ret string
		return ret
	}
	return *o.AddRequestReferenceId.Get()
}

// GetAddRequestReferenceIdOk returns a tuple with the AddRequestReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AeAssetSellerChargeUpsertV2Detail) GetAddRequestReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddRequestReferenceId.Get(), o.AddRequestReferenceId.IsSet()
}

// HasAddRequestReferenceId returns a boolean if a field has been set.
func (o *AeAssetSellerChargeUpsertV2Detail) HasAddRequestReferenceId() bool {
	if o != nil && o.AddRequestReferenceId.IsSet() {
		return true
	}

	return false
}

// SetAddRequestReferenceId gets a reference to the given NullableString and assigns it to the AddRequestReferenceId field.
func (o *AeAssetSellerChargeUpsertV2Detail) SetAddRequestReferenceId(v string) {
	o.AddRequestReferenceId.Set(&v)
}
// SetAddRequestReferenceIdNil sets the value for AddRequestReferenceId to be an explicit nil
func (o *AeAssetSellerChargeUpsertV2Detail) SetAddRequestReferenceIdNil() {
	o.AddRequestReferenceId.Set(nil)
}

// UnsetAddRequestReferenceId ensures that no value is present for AddRequestReferenceId, not even an explicit nil
func (o *AeAssetSellerChargeUpsertV2Detail) UnsetAddRequestReferenceId() {
	o.AddRequestReferenceId.Unset()
}

// GetCharge returns the Charge field value
func (o *AeAssetSellerChargeUpsertV2Detail) GetCharge() CommonCharge {
	if o == nil {
		var ret CommonCharge
		return ret
	}

	return o.Charge
}

// GetChargeOk returns a tuple with the Charge field value
// and a boolean to check if the value has been set.
func (o *AeAssetSellerChargeUpsertV2Detail) GetChargeOk() (*CommonCharge, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Charge, true
}

// SetCharge sets field value
func (o *AeAssetSellerChargeUpsertV2Detail) SetCharge(v CommonCharge) {
	o.Charge = v
}

// GetInitiator returns the Initiator field value if set, zero value otherwise.
func (o *AeAssetSellerChargeUpsertV2Detail) GetInitiator() CommonInitiator {
	if o == nil || IsNil(o.Initiator) {
		var ret CommonInitiator
		return ret
	}
	return *o.Initiator
}

// GetInitiatorOk returns a tuple with the Initiator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAssetSellerChargeUpsertV2Detail) GetInitiatorOk() (*CommonInitiator, bool) {
	if o == nil || IsNil(o.Initiator) {
		return nil, false
	}
	return o.Initiator, true
}

// HasInitiator returns a boolean if a field has been set.
func (o *AeAssetSellerChargeUpsertV2Detail) HasInitiator() bool {
	if o != nil && !IsNil(o.Initiator) {
		return true
	}

	return false
}

// SetInitiator gets a reference to the given CommonInitiator and assigns it to the Initiator field.
func (o *AeAssetSellerChargeUpsertV2Detail) SetInitiator(v CommonInitiator) {
	o.Initiator = &v
}

func (o AeAssetSellerChargeUpsertV2Detail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AeAssetSellerChargeUpsertV2Detail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["auction-id"] = o.AuctionId
	toSerialize["asset"] = o.Asset
	toSerialize["seller"] = o.Seller
	if o.AddRequestReferenceId.IsSet() {
		toSerialize["add-request-reference-id"] = o.AddRequestReferenceId.Get()
	}
	toSerialize["charge"] = o.Charge
	if !IsNil(o.Initiator) {
		toSerialize["initiator"] = o.Initiator
	}
	return toSerialize, nil
}

func (o *AeAssetSellerChargeUpsertV2Detail) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"auction-id",
		"asset",
		"seller",
		"charge",
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

	varAeAssetSellerChargeUpsertV2Detail := _AeAssetSellerChargeUpsertV2Detail{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAeAssetSellerChargeUpsertV2Detail)

	if err != nil {
		return err
	}

	*o = AeAssetSellerChargeUpsertV2Detail(varAeAssetSellerChargeUpsertV2Detail)

	return err
}

type NullableAeAssetSellerChargeUpsertV2Detail struct {
	value *AeAssetSellerChargeUpsertV2Detail
	isSet bool
}

func (v NullableAeAssetSellerChargeUpsertV2Detail) Get() *AeAssetSellerChargeUpsertV2Detail {
	return v.value
}

func (v *NullableAeAssetSellerChargeUpsertV2Detail) Set(val *AeAssetSellerChargeUpsertV2Detail) {
	v.value = val
	v.isSet = true
}

func (v NullableAeAssetSellerChargeUpsertV2Detail) IsSet() bool {
	return v.isSet
}

func (v *NullableAeAssetSellerChargeUpsertV2Detail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAeAssetSellerChargeUpsertV2Detail(val *AeAssetSellerChargeUpsertV2Detail) *NullableAeAssetSellerChargeUpsertV2Detail {
	return &NullableAeAssetSellerChargeUpsertV2Detail{value: val, isSet: true}
}

func (v NullableAeAssetSellerChargeUpsertV2Detail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAeAssetSellerChargeUpsertV2Detail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


