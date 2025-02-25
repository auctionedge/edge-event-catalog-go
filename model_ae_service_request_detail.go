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

// checks if the AeServiceRequestDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AeServiceRequestDetail{}

// AeServiceRequestDetail Request that a service be performed
type AeServiceRequestDetail struct {
	// Auction Edge unique identifier for an auction.
	AuctionId string `json:"auction-id"`
	// The Vehicle Identification Number(VIN) of the asset.
	Vin string `json:"vin"`
	// The stock number of the asset.
	Stock string `json:"stock"`
	Service ServiceOrder `json:"service"`
	Initiator *CommonInitiator `json:"initiator,omitempty"`
}

type _AeServiceRequestDetail AeServiceRequestDetail

// NewAeServiceRequestDetail instantiates a new AeServiceRequestDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAeServiceRequestDetail(auctionId string, vin string, stock string, service ServiceOrder) *AeServiceRequestDetail {
	this := AeServiceRequestDetail{}
	this.AuctionId = auctionId
	this.Vin = vin
	this.Stock = stock
	this.Service = service
	return &this
}

// NewAeServiceRequestDetailWithDefaults instantiates a new AeServiceRequestDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAeServiceRequestDetailWithDefaults() *AeServiceRequestDetail {
	this := AeServiceRequestDetail{}
	return &this
}

// GetAuctionId returns the AuctionId field value
func (o *AeServiceRequestDetail) GetAuctionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuctionId
}

// GetAuctionIdOk returns a tuple with the AuctionId field value
// and a boolean to check if the value has been set.
func (o *AeServiceRequestDetail) GetAuctionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuctionId, true
}

// SetAuctionId sets field value
func (o *AeServiceRequestDetail) SetAuctionId(v string) {
	o.AuctionId = v
}

// GetVin returns the Vin field value
func (o *AeServiceRequestDetail) GetVin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vin
}

// GetVinOk returns a tuple with the Vin field value
// and a boolean to check if the value has been set.
func (o *AeServiceRequestDetail) GetVinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vin, true
}

// SetVin sets field value
func (o *AeServiceRequestDetail) SetVin(v string) {
	o.Vin = v
}

// GetStock returns the Stock field value
func (o *AeServiceRequestDetail) GetStock() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Stock
}

// GetStockOk returns a tuple with the Stock field value
// and a boolean to check if the value has been set.
func (o *AeServiceRequestDetail) GetStockOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stock, true
}

// SetStock sets field value
func (o *AeServiceRequestDetail) SetStock(v string) {
	o.Stock = v
}

// GetService returns the Service field value
func (o *AeServiceRequestDetail) GetService() ServiceOrder {
	if o == nil {
		var ret ServiceOrder
		return ret
	}

	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *AeServiceRequestDetail) GetServiceOk() (*ServiceOrder, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value
func (o *AeServiceRequestDetail) SetService(v ServiceOrder) {
	o.Service = v
}

// GetInitiator returns the Initiator field value if set, zero value otherwise.
func (o *AeServiceRequestDetail) GetInitiator() CommonInitiator {
	if o == nil || IsNil(o.Initiator) {
		var ret CommonInitiator
		return ret
	}
	return *o.Initiator
}

// GetInitiatorOk returns a tuple with the Initiator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeServiceRequestDetail) GetInitiatorOk() (*CommonInitiator, bool) {
	if o == nil || IsNil(o.Initiator) {
		return nil, false
	}
	return o.Initiator, true
}

// HasInitiator returns a boolean if a field has been set.
func (o *AeServiceRequestDetail) HasInitiator() bool {
	if o != nil && !IsNil(o.Initiator) {
		return true
	}

	return false
}

// SetInitiator gets a reference to the given CommonInitiator and assigns it to the Initiator field.
func (o *AeServiceRequestDetail) SetInitiator(v CommonInitiator) {
	o.Initiator = &v
}

func (o AeServiceRequestDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AeServiceRequestDetail) ToMap() (map[string]interface{}, error) {
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

func (o *AeServiceRequestDetail) UnmarshalJSON(data []byte) (err error) {
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

	varAeServiceRequestDetail := _AeServiceRequestDetail{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAeServiceRequestDetail)

	if err != nil {
		return err
	}

	*o = AeServiceRequestDetail(varAeServiceRequestDetail)

	return err
}

type NullableAeServiceRequestDetail struct {
	value *AeServiceRequestDetail
	isSet bool
}

func (v NullableAeServiceRequestDetail) Get() *AeServiceRequestDetail {
	return v.value
}

func (v *NullableAeServiceRequestDetail) Set(val *AeServiceRequestDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableAeServiceRequestDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableAeServiceRequestDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAeServiceRequestDetail(val *AeServiceRequestDetail) *NullableAeServiceRequestDetail {
	return &NullableAeServiceRequestDetail{value: val, isSet: true}
}

func (v NullableAeServiceRequestDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAeServiceRequestDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


