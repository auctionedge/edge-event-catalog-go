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

// checks if the AeAssetDealPaymentStatusChanged type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AeAssetDealPaymentStatusChanged{}

// AeAssetDealPaymentStatusChanged Notification that a service order has been placed
type AeAssetDealPaymentStatusChanged struct {
	Detail AeAssetDealPaymentStatusChangedDetail `json:"detail"`
	// The publisher AWS account number
	Account *string `json:"account,omitempty"`
	// Identifies, in combination with the source field, the fields and values that appear in the detail field.
	DetailType string `json:"detail-type"`
	// A Version 4 UUID that's generated for every event. You can use id to trace events as they move through rules to targets.
	Id string `json:"id"`
	// Identifies the AWS Region where the event originated
	Region *string `json:"region,omitempty"`
	// Identifiers for any resources involved in the event
	Resources []string `json:"resources,omitempty"`
	// Identifies the service that generated the event. May not begin with 'aws.'
	Source string `json:"source"`
	// The event timestamp, which can be specified by the service originating the event. May be before the actual publication of the event. ISO-8601 format
	Time *time.Time `json:"time,omitempty"`
	// UA id of the auction. Not all events will need the identification of an auction.
	AuctionId *string `json:"auction-id,omitempty"`
	// Timestamp after which the consumer should no longer process the message. ISO-8601 format
	ExpiresAt *time.Time `json:"expires-at,omitempty"`
	// True if this version has been deprecated.
	Deprecated *bool `json:"deprecated,omitempty"`
	// MD5 content hash on the entire event content (Library generated)
	IdempotencyKey *string `json:"idempotency-key,omitempty"`
}

type _AeAssetDealPaymentStatusChanged AeAssetDealPaymentStatusChanged

// NewAeAssetDealPaymentStatusChanged instantiates a new AeAssetDealPaymentStatusChanged object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAeAssetDealPaymentStatusChanged(detail AeAssetDealPaymentStatusChangedDetail, detailType string, id string, source string) *AeAssetDealPaymentStatusChanged {
	this := AeAssetDealPaymentStatusChanged{}
	this.DetailType = detailType
	this.Id = id
	this.Source = source
	var deprecated bool = false
	this.Deprecated = &deprecated
	return &this
}

// NewAeAssetDealPaymentStatusChangedWithDefaults instantiates a new AeAssetDealPaymentStatusChanged object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAeAssetDealPaymentStatusChangedWithDefaults() *AeAssetDealPaymentStatusChanged {
	this := AeAssetDealPaymentStatusChanged{}
	var deprecated bool = false
	this.Deprecated = &deprecated
	return &this
}

// GetDetail returns the Detail field value
func (o *AeAssetDealPaymentStatusChanged) GetDetail() AeAssetDealPaymentStatusChangedDetail {
	if o == nil {
		var ret AeAssetDealPaymentStatusChangedDetail
		return ret
	}

	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *AeAssetDealPaymentStatusChanged) GetDetailOk() (*AeAssetDealPaymentStatusChangedDetail, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Detail, true
}

// SetDetail sets field value
func (o *AeAssetDealPaymentStatusChanged) SetDetail(v AeAssetDealPaymentStatusChangedDetail) {
	o.Detail = v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *AeAssetDealPaymentStatusChanged) GetAccount() string {
	if o == nil || IsNil(o.Account) {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAssetDealPaymentStatusChanged) GetAccountOk() (*string, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *AeAssetDealPaymentStatusChanged) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *AeAssetDealPaymentStatusChanged) SetAccount(v string) {
	o.Account = &v
}

// GetDetailType returns the DetailType field value
func (o *AeAssetDealPaymentStatusChanged) GetDetailType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DetailType
}

// GetDetailTypeOk returns a tuple with the DetailType field value
// and a boolean to check if the value has been set.
func (o *AeAssetDealPaymentStatusChanged) GetDetailTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DetailType, true
}

// SetDetailType sets field value
func (o *AeAssetDealPaymentStatusChanged) SetDetailType(v string) {
	o.DetailType = v
}

// GetId returns the Id field value
func (o *AeAssetDealPaymentStatusChanged) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AeAssetDealPaymentStatusChanged) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AeAssetDealPaymentStatusChanged) SetId(v string) {
	o.Id = v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *AeAssetDealPaymentStatusChanged) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAssetDealPaymentStatusChanged) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *AeAssetDealPaymentStatusChanged) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *AeAssetDealPaymentStatusChanged) SetRegion(v string) {
	o.Region = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *AeAssetDealPaymentStatusChanged) GetResources() []string {
	if o == nil || IsNil(o.Resources) {
		var ret []string
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAssetDealPaymentStatusChanged) GetResourcesOk() ([]string, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *AeAssetDealPaymentStatusChanged) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []string and assigns it to the Resources field.
func (o *AeAssetDealPaymentStatusChanged) SetResources(v []string) {
	o.Resources = v
}

// GetSource returns the Source field value
func (o *AeAssetDealPaymentStatusChanged) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *AeAssetDealPaymentStatusChanged) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *AeAssetDealPaymentStatusChanged) SetSource(v string) {
	o.Source = v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *AeAssetDealPaymentStatusChanged) GetTime() time.Time {
	if o == nil || IsNil(o.Time) {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAssetDealPaymentStatusChanged) GetTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *AeAssetDealPaymentStatusChanged) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *AeAssetDealPaymentStatusChanged) SetTime(v time.Time) {
	o.Time = &v
}

// GetAuctionId returns the AuctionId field value if set, zero value otherwise.
func (o *AeAssetDealPaymentStatusChanged) GetAuctionId() string {
	if o == nil || IsNil(o.AuctionId) {
		var ret string
		return ret
	}
	return *o.AuctionId
}

// GetAuctionIdOk returns a tuple with the AuctionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAssetDealPaymentStatusChanged) GetAuctionIdOk() (*string, bool) {
	if o == nil || IsNil(o.AuctionId) {
		return nil, false
	}
	return o.AuctionId, true
}

// HasAuctionId returns a boolean if a field has been set.
func (o *AeAssetDealPaymentStatusChanged) HasAuctionId() bool {
	if o != nil && !IsNil(o.AuctionId) {
		return true
	}

	return false
}

// SetAuctionId gets a reference to the given string and assigns it to the AuctionId field.
func (o *AeAssetDealPaymentStatusChanged) SetAuctionId(v string) {
	o.AuctionId = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *AeAssetDealPaymentStatusChanged) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAssetDealPaymentStatusChanged) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *AeAssetDealPaymentStatusChanged) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *AeAssetDealPaymentStatusChanged) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetDeprecated returns the Deprecated field value if set, zero value otherwise.
func (o *AeAssetDealPaymentStatusChanged) GetDeprecated() bool {
	if o == nil || IsNil(o.Deprecated) {
		var ret bool
		return ret
	}
	return *o.Deprecated
}

// GetDeprecatedOk returns a tuple with the Deprecated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAssetDealPaymentStatusChanged) GetDeprecatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Deprecated) {
		return nil, false
	}
	return o.Deprecated, true
}

// HasDeprecated returns a boolean if a field has been set.
func (o *AeAssetDealPaymentStatusChanged) HasDeprecated() bool {
	if o != nil && !IsNil(o.Deprecated) {
		return true
	}

	return false
}

// SetDeprecated gets a reference to the given bool and assigns it to the Deprecated field.
func (o *AeAssetDealPaymentStatusChanged) SetDeprecated(v bool) {
	o.Deprecated = &v
}

// GetIdempotencyKey returns the IdempotencyKey field value if set, zero value otherwise.
func (o *AeAssetDealPaymentStatusChanged) GetIdempotencyKey() string {
	if o == nil || IsNil(o.IdempotencyKey) {
		var ret string
		return ret
	}
	return *o.IdempotencyKey
}

// GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAssetDealPaymentStatusChanged) GetIdempotencyKeyOk() (*string, bool) {
	if o == nil || IsNil(o.IdempotencyKey) {
		return nil, false
	}
	return o.IdempotencyKey, true
}

// HasIdempotencyKey returns a boolean if a field has been set.
func (o *AeAssetDealPaymentStatusChanged) HasIdempotencyKey() bool {
	if o != nil && !IsNil(o.IdempotencyKey) {
		return true
	}

	return false
}

// SetIdempotencyKey gets a reference to the given string and assigns it to the IdempotencyKey field.
func (o *AeAssetDealPaymentStatusChanged) SetIdempotencyKey(v string) {
	o.IdempotencyKey = &v
}

func (o AeAssetDealPaymentStatusChanged) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AeAssetDealPaymentStatusChanged) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["detail"] = o.Detail
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	toSerialize["detail-type"] = o.DetailType
	toSerialize["id"] = o.Id
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	toSerialize["source"] = o.Source
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.AuctionId) {
		toSerialize["auction-id"] = o.AuctionId
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expires-at"] = o.ExpiresAt
	}
	if !IsNil(o.Deprecated) {
		toSerialize["deprecated"] = o.Deprecated
	}
	if !IsNil(o.IdempotencyKey) {
		toSerialize["idempotency-key"] = o.IdempotencyKey
	}
	return toSerialize, nil
}

func (o *AeAssetDealPaymentStatusChanged) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"detail",
		"detail-type",
		"id",
		"source",
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

	varAeAssetDealPaymentStatusChanged := _AeAssetDealPaymentStatusChanged{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAeAssetDealPaymentStatusChanged)

	if err != nil {
		return err
	}

	*o = AeAssetDealPaymentStatusChanged(varAeAssetDealPaymentStatusChanged)

	return err
}

type NullableAeAssetDealPaymentStatusChanged struct {
	value *AeAssetDealPaymentStatusChanged
	isSet bool
}

func (v NullableAeAssetDealPaymentStatusChanged) Get() *AeAssetDealPaymentStatusChanged {
	return v.value
}

func (v *NullableAeAssetDealPaymentStatusChanged) Set(val *AeAssetDealPaymentStatusChanged) {
	v.value = val
	v.isSet = true
}

func (v NullableAeAssetDealPaymentStatusChanged) IsSet() bool {
	return v.isSet
}

func (v *NullableAeAssetDealPaymentStatusChanged) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAeAssetDealPaymentStatusChanged(val *AeAssetDealPaymentStatusChanged) *NullableAeAssetDealPaymentStatusChanged {
	return &NullableAeAssetDealPaymentStatusChanged{value: val, isSet: true}
}

func (v NullableAeAssetDealPaymentStatusChanged) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAeAssetDealPaymentStatusChanged) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


