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
	"bytes"
	"fmt"
)

// checks if the AeAssetSaleListingUpsertedAmsDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AeAssetSaleListingUpsertedAmsDetail{}

// AeAssetSaleListingUpsertedAmsDetail A sale listing for an asset has been created or modified by the AMS
type AeAssetSaleListingUpsertedAmsDetail struct {
	// Auction Edge unique identifier for an auction.
	AuctionId string `json:"auction-id"`
	Asset CommonAmsAssetPointer `json:"asset"`
	SaleListing CommonAssetSaleListing `json:"sale-listing"`
	// The updated date time of the deal.
	UpdatedAt time.Time `json:"updated-at"`
	Initiator CommonInitiator `json:"initiator"`
}

type _AeAssetSaleListingUpsertedAmsDetail AeAssetSaleListingUpsertedAmsDetail

// NewAeAssetSaleListingUpsertedAmsDetail instantiates a new AeAssetSaleListingUpsertedAmsDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAeAssetSaleListingUpsertedAmsDetail(auctionId string, asset CommonAmsAssetPointer, saleListing CommonAssetSaleListing, updatedAt time.Time, initiator CommonInitiator) *AeAssetSaleListingUpsertedAmsDetail {
	this := AeAssetSaleListingUpsertedAmsDetail{}
	this.AuctionId = auctionId
	this.Asset = asset
	this.SaleListing = saleListing
	this.UpdatedAt = updatedAt
	this.Initiator = initiator
	return &this
}

// NewAeAssetSaleListingUpsertedAmsDetailWithDefaults instantiates a new AeAssetSaleListingUpsertedAmsDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAeAssetSaleListingUpsertedAmsDetailWithDefaults() *AeAssetSaleListingUpsertedAmsDetail {
	this := AeAssetSaleListingUpsertedAmsDetail{}
	return &this
}

// GetAuctionId returns the AuctionId field value
func (o *AeAssetSaleListingUpsertedAmsDetail) GetAuctionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuctionId
}

// GetAuctionIdOk returns a tuple with the AuctionId field value
// and a boolean to check if the value has been set.
func (o *AeAssetSaleListingUpsertedAmsDetail) GetAuctionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuctionId, true
}

// SetAuctionId sets field value
func (o *AeAssetSaleListingUpsertedAmsDetail) SetAuctionId(v string) {
	o.AuctionId = v
}

// GetAsset returns the Asset field value
func (o *AeAssetSaleListingUpsertedAmsDetail) GetAsset() CommonAmsAssetPointer {
	if o == nil {
		var ret CommonAmsAssetPointer
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *AeAssetSaleListingUpsertedAmsDetail) GetAssetOk() (*CommonAmsAssetPointer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *AeAssetSaleListingUpsertedAmsDetail) SetAsset(v CommonAmsAssetPointer) {
	o.Asset = v
}

// GetSaleListing returns the SaleListing field value
func (o *AeAssetSaleListingUpsertedAmsDetail) GetSaleListing() CommonAssetSaleListing {
	if o == nil {
		var ret CommonAssetSaleListing
		return ret
	}

	return o.SaleListing
}

// GetSaleListingOk returns a tuple with the SaleListing field value
// and a boolean to check if the value has been set.
func (o *AeAssetSaleListingUpsertedAmsDetail) GetSaleListingOk() (*CommonAssetSaleListing, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SaleListing, true
}

// SetSaleListing sets field value
func (o *AeAssetSaleListingUpsertedAmsDetail) SetSaleListing(v CommonAssetSaleListing) {
	o.SaleListing = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *AeAssetSaleListingUpsertedAmsDetail) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *AeAssetSaleListingUpsertedAmsDetail) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *AeAssetSaleListingUpsertedAmsDetail) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetInitiator returns the Initiator field value
func (o *AeAssetSaleListingUpsertedAmsDetail) GetInitiator() CommonInitiator {
	if o == nil {
		var ret CommonInitiator
		return ret
	}

	return o.Initiator
}

// GetInitiatorOk returns a tuple with the Initiator field value
// and a boolean to check if the value has been set.
func (o *AeAssetSaleListingUpsertedAmsDetail) GetInitiatorOk() (*CommonInitiator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Initiator, true
}

// SetInitiator sets field value
func (o *AeAssetSaleListingUpsertedAmsDetail) SetInitiator(v CommonInitiator) {
	o.Initiator = v
}

func (o AeAssetSaleListingUpsertedAmsDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AeAssetSaleListingUpsertedAmsDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["auction-id"] = o.AuctionId
	toSerialize["asset"] = o.Asset
	toSerialize["sale-listing"] = o.SaleListing
	toSerialize["updated-at"] = o.UpdatedAt
	toSerialize["initiator"] = o.Initiator
	return toSerialize, nil
}

func (o *AeAssetSaleListingUpsertedAmsDetail) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"auction-id",
		"asset",
		"sale-listing",
		"updated-at",
		"initiator",
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

	varAeAssetSaleListingUpsertedAmsDetail := _AeAssetSaleListingUpsertedAmsDetail{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAeAssetSaleListingUpsertedAmsDetail)

	if err != nil {
		return err
	}

	*o = AeAssetSaleListingUpsertedAmsDetail(varAeAssetSaleListingUpsertedAmsDetail)

	return err
}

type NullableAeAssetSaleListingUpsertedAmsDetail struct {
	value *AeAssetSaleListingUpsertedAmsDetail
	isSet bool
}

func (v NullableAeAssetSaleListingUpsertedAmsDetail) Get() *AeAssetSaleListingUpsertedAmsDetail {
	return v.value
}

func (v *NullableAeAssetSaleListingUpsertedAmsDetail) Set(val *AeAssetSaleListingUpsertedAmsDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableAeAssetSaleListingUpsertedAmsDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableAeAssetSaleListingUpsertedAmsDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAeAssetSaleListingUpsertedAmsDetail(val *AeAssetSaleListingUpsertedAmsDetail) *NullableAeAssetSaleListingUpsertedAmsDetail {
	return &NullableAeAssetSaleListingUpsertedAmsDetail{value: val, isSet: true}
}

func (v NullableAeAssetSaleListingUpsertedAmsDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAeAssetSaleListingUpsertedAmsDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


