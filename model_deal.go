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

// checks if the Deal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Deal{}

// Deal Details on the current deal for the asset.
type Deal struct {
	AmsSaleListing DealAmsSaleListing `json:"ams-sale-listing"`
	Buyer CommonAmsAccountPointer `json:"buyer"`
	// The date time that asset was sold at.
	SoldAt NullableTime `json:"sold-at,omitempty"`
	// The amount that the asset was sold for.
	SoldAmountUsd float32 `json:"sold-amount-usd"`
	// Whether or not the current deal is \"on offer\" / \"on IF\"
	OnOffer *bool `json:"on-offer,omitempty"`
	// Whether or not the current deal is being arbitrated.
	InArbitration *bool `json:"in-arbitration,omitempty"`
}

type _Deal Deal

// NewDeal instantiates a new Deal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeal(amsSaleListing DealAmsSaleListing, buyer CommonAmsAccountPointer, soldAmountUsd float32) *Deal {
	this := Deal{}
	this.AmsSaleListing = amsSaleListing
	this.Buyer = buyer
	this.SoldAmountUsd = soldAmountUsd
	var onOffer bool = false
	this.OnOffer = &onOffer
	var inArbitration bool = false
	this.InArbitration = &inArbitration
	return &this
}

// NewDealWithDefaults instantiates a new Deal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDealWithDefaults() *Deal {
	this := Deal{}
	var onOffer bool = false
	this.OnOffer = &onOffer
	var inArbitration bool = false
	this.InArbitration = &inArbitration
	return &this
}

// GetAmsSaleListing returns the AmsSaleListing field value
func (o *Deal) GetAmsSaleListing() DealAmsSaleListing {
	if o == nil {
		var ret DealAmsSaleListing
		return ret
	}

	return o.AmsSaleListing
}

// GetAmsSaleListingOk returns a tuple with the AmsSaleListing field value
// and a boolean to check if the value has been set.
func (o *Deal) GetAmsSaleListingOk() (*DealAmsSaleListing, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmsSaleListing, true
}

// SetAmsSaleListing sets field value
func (o *Deal) SetAmsSaleListing(v DealAmsSaleListing) {
	o.AmsSaleListing = v
}

// GetBuyer returns the Buyer field value
func (o *Deal) GetBuyer() CommonAmsAccountPointer {
	if o == nil {
		var ret CommonAmsAccountPointer
		return ret
	}

	return o.Buyer
}

// GetBuyerOk returns a tuple with the Buyer field value
// and a boolean to check if the value has been set.
func (o *Deal) GetBuyerOk() (*CommonAmsAccountPointer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Buyer, true
}

// SetBuyer sets field value
func (o *Deal) SetBuyer(v CommonAmsAccountPointer) {
	o.Buyer = v
}

// GetSoldAt returns the SoldAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Deal) GetSoldAt() time.Time {
	if o == nil || IsNil(o.SoldAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.SoldAt.Get()
}

// GetSoldAtOk returns a tuple with the SoldAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Deal) GetSoldAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.SoldAt.Get(), o.SoldAt.IsSet()
}

// HasSoldAt returns a boolean if a field has been set.
func (o *Deal) HasSoldAt() bool {
	if o != nil && o.SoldAt.IsSet() {
		return true
	}

	return false
}

// SetSoldAt gets a reference to the given NullableTime and assigns it to the SoldAt field.
func (o *Deal) SetSoldAt(v time.Time) {
	o.SoldAt.Set(&v)
}
// SetSoldAtNil sets the value for SoldAt to be an explicit nil
func (o *Deal) SetSoldAtNil() {
	o.SoldAt.Set(nil)
}

// UnsetSoldAt ensures that no value is present for SoldAt, not even an explicit nil
func (o *Deal) UnsetSoldAt() {
	o.SoldAt.Unset()
}

// GetSoldAmountUsd returns the SoldAmountUsd field value
func (o *Deal) GetSoldAmountUsd() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SoldAmountUsd
}

// GetSoldAmountUsdOk returns a tuple with the SoldAmountUsd field value
// and a boolean to check if the value has been set.
func (o *Deal) GetSoldAmountUsdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SoldAmountUsd, true
}

// SetSoldAmountUsd sets field value
func (o *Deal) SetSoldAmountUsd(v float32) {
	o.SoldAmountUsd = v
}

// GetOnOffer returns the OnOffer field value if set, zero value otherwise.
func (o *Deal) GetOnOffer() bool {
	if o == nil || IsNil(o.OnOffer) {
		var ret bool
		return ret
	}
	return *o.OnOffer
}

// GetOnOfferOk returns a tuple with the OnOffer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deal) GetOnOfferOk() (*bool, bool) {
	if o == nil || IsNil(o.OnOffer) {
		return nil, false
	}
	return o.OnOffer, true
}

// HasOnOffer returns a boolean if a field has been set.
func (o *Deal) HasOnOffer() bool {
	if o != nil && !IsNil(o.OnOffer) {
		return true
	}

	return false
}

// SetOnOffer gets a reference to the given bool and assigns it to the OnOffer field.
func (o *Deal) SetOnOffer(v bool) {
	o.OnOffer = &v
}

// GetInArbitration returns the InArbitration field value if set, zero value otherwise.
func (o *Deal) GetInArbitration() bool {
	if o == nil || IsNil(o.InArbitration) {
		var ret bool
		return ret
	}
	return *o.InArbitration
}

// GetInArbitrationOk returns a tuple with the InArbitration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deal) GetInArbitrationOk() (*bool, bool) {
	if o == nil || IsNil(o.InArbitration) {
		return nil, false
	}
	return o.InArbitration, true
}

// HasInArbitration returns a boolean if a field has been set.
func (o *Deal) HasInArbitration() bool {
	if o != nil && !IsNil(o.InArbitration) {
		return true
	}

	return false
}

// SetInArbitration gets a reference to the given bool and assigns it to the InArbitration field.
func (o *Deal) SetInArbitration(v bool) {
	o.InArbitration = &v
}

func (o Deal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Deal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ams-sale-listing"] = o.AmsSaleListing
	toSerialize["buyer"] = o.Buyer
	if o.SoldAt.IsSet() {
		toSerialize["sold-at"] = o.SoldAt.Get()
	}
	toSerialize["sold-amount-usd"] = o.SoldAmountUsd
	if !IsNil(o.OnOffer) {
		toSerialize["on-offer"] = o.OnOffer
	}
	if !IsNil(o.InArbitration) {
		toSerialize["in-arbitration"] = o.InArbitration
	}
	return toSerialize, nil
}

func (o *Deal) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ams-sale-listing",
		"buyer",
		"sold-amount-usd",
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

	varDeal := _Deal{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeal)

	if err != nil {
		return err
	}

	*o = Deal(varDeal)

	return err
}

type NullableDeal struct {
	value *Deal
	isSet bool
}

func (v NullableDeal) Get() *Deal {
	return v.value
}

func (v *NullableDeal) Set(val *Deal) {
	v.value = val
	v.isSet = true
}

func (v NullableDeal) IsSet() bool {
	return v.isSet
}

func (v *NullableDeal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeal(val *Deal) *NullableDeal {
	return &NullableDeal{value: val, isSet: true}
}

func (v NullableDeal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


