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

// checks if the AeAssetSellerChargeUpsert type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AeAssetSellerChargeUpsert{}

// AeAssetSellerChargeUpsert Notification that a service order has been placed
type AeAssetSellerChargeUpsert struct {
	Detail AeAssetSellerChargeUpsertDetail `json:"detail"`
	// Optional version
	Version *string `json:"version,omitempty"`
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

type _AeAssetSellerChargeUpsert AeAssetSellerChargeUpsert

// NewAeAssetSellerChargeUpsert instantiates a new AeAssetSellerChargeUpsert object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAeAssetSellerChargeUpsert(detail AeAssetSellerChargeUpsertDetail, detailType string, id string, source string) *AeAssetSellerChargeUpsert {
	this := AeAssetSellerChargeUpsert{}
	this.DetailType = detailType
	this.Id = id
	this.Source = source
	var deprecated bool = false
	this.Deprecated = &deprecated
	return &this
}

// NewAeAssetSellerChargeUpsertWithDefaults instantiates a new AeAssetSellerChargeUpsert object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAeAssetSellerChargeUpsertWithDefaults() *AeAssetSellerChargeUpsert {
	this := AeAssetSellerChargeUpsert{}
	var deprecated bool = false
	this.Deprecated = &deprecated
	return &this
}

// GetDetail returns the Detail field value
func (o *AeAssetSellerChargeUpsert) GetDetail() AeAssetSellerChargeUpsertDetail {
	if o == nil {
		var ret AeAssetSellerChargeUpsertDetail
		return ret
	}

	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *AeAssetSellerChargeUpsert) GetDetailOk() (*AeAssetSellerChargeUpsertDetail, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Detail, true
}

// SetDetail sets field value
func (o *AeAssetSellerChargeUpsert) SetDetail(v AeAssetSellerChargeUpsertDetail) {
	o.Detail = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *AeAssetSellerChargeUpsert) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAssetSellerChargeUpsert) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *AeAssetSellerChargeUpsert) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *AeAssetSellerChargeUpsert) SetVersion(v string) {
	o.Version = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *AeAssetSellerChargeUpsert) GetAccount() string {
	if o == nil || IsNil(o.Account) {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAssetSellerChargeUpsert) GetAccountOk() (*string, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *AeAssetSellerChargeUpsert) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *AeAssetSellerChargeUpsert) SetAccount(v string) {
	o.Account = &v
}

// GetDetailType returns the DetailType field value
func (o *AeAssetSellerChargeUpsert) GetDetailType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DetailType
}

// GetDetailTypeOk returns a tuple with the DetailType field value
// and a boolean to check if the value has been set.
func (o *AeAssetSellerChargeUpsert) GetDetailTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DetailType, true
}

// SetDetailType sets field value
func (o *AeAssetSellerChargeUpsert) SetDetailType(v string) {
	o.DetailType = v
}

// GetId returns the Id field value
func (o *AeAssetSellerChargeUpsert) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AeAssetSellerChargeUpsert) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AeAssetSellerChargeUpsert) SetId(v string) {
	o.Id = v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *AeAssetSellerChargeUpsert) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAssetSellerChargeUpsert) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *AeAssetSellerChargeUpsert) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *AeAssetSellerChargeUpsert) SetRegion(v string) {
	o.Region = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *AeAssetSellerChargeUpsert) GetResources() []string {
	if o == nil || IsNil(o.Resources) {
		var ret []string
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAssetSellerChargeUpsert) GetResourcesOk() ([]string, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *AeAssetSellerChargeUpsert) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []string and assigns it to the Resources field.
func (o *AeAssetSellerChargeUpsert) SetResources(v []string) {
	o.Resources = v
}

// GetSource returns the Source field value
func (o *AeAssetSellerChargeUpsert) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *AeAssetSellerChargeUpsert) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *AeAssetSellerChargeUpsert) SetSource(v string) {
	o.Source = v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *AeAssetSellerChargeUpsert) GetTime() time.Time {
	if o == nil || IsNil(o.Time) {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAssetSellerChargeUpsert) GetTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *AeAssetSellerChargeUpsert) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *AeAssetSellerChargeUpsert) SetTime(v time.Time) {
	o.Time = &v
}

// GetAuctionId returns the AuctionId field value if set, zero value otherwise.
func (o *AeAssetSellerChargeUpsert) GetAuctionId() string {
	if o == nil || IsNil(o.AuctionId) {
		var ret string
		return ret
	}
	return *o.AuctionId
}

// GetAuctionIdOk returns a tuple with the AuctionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAssetSellerChargeUpsert) GetAuctionIdOk() (*string, bool) {
	if o == nil || IsNil(o.AuctionId) {
		return nil, false
	}
	return o.AuctionId, true
}

// HasAuctionId returns a boolean if a field has been set.
func (o *AeAssetSellerChargeUpsert) HasAuctionId() bool {
	if o != nil && !IsNil(o.AuctionId) {
		return true
	}

	return false
}

// SetAuctionId gets a reference to the given string and assigns it to the AuctionId field.
func (o *AeAssetSellerChargeUpsert) SetAuctionId(v string) {
	o.AuctionId = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *AeAssetSellerChargeUpsert) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAssetSellerChargeUpsert) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *AeAssetSellerChargeUpsert) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *AeAssetSellerChargeUpsert) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetDeprecated returns the Deprecated field value if set, zero value otherwise.
func (o *AeAssetSellerChargeUpsert) GetDeprecated() bool {
	if o == nil || IsNil(o.Deprecated) {
		var ret bool
		return ret
	}
	return *o.Deprecated
}

// GetDeprecatedOk returns a tuple with the Deprecated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAssetSellerChargeUpsert) GetDeprecatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Deprecated) {
		return nil, false
	}
	return o.Deprecated, true
}

// HasDeprecated returns a boolean if a field has been set.
func (o *AeAssetSellerChargeUpsert) HasDeprecated() bool {
	if o != nil && !IsNil(o.Deprecated) {
		return true
	}

	return false
}

// SetDeprecated gets a reference to the given bool and assigns it to the Deprecated field.
func (o *AeAssetSellerChargeUpsert) SetDeprecated(v bool) {
	o.Deprecated = &v
}

// GetIdempotencyKey returns the IdempotencyKey field value if set, zero value otherwise.
func (o *AeAssetSellerChargeUpsert) GetIdempotencyKey() string {
	if o == nil || IsNil(o.IdempotencyKey) {
		var ret string
		return ret
	}
	return *o.IdempotencyKey
}

// GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAssetSellerChargeUpsert) GetIdempotencyKeyOk() (*string, bool) {
	if o == nil || IsNil(o.IdempotencyKey) {
		return nil, false
	}
	return o.IdempotencyKey, true
}

// HasIdempotencyKey returns a boolean if a field has been set.
func (o *AeAssetSellerChargeUpsert) HasIdempotencyKey() bool {
	if o != nil && !IsNil(o.IdempotencyKey) {
		return true
	}

	return false
}

// SetIdempotencyKey gets a reference to the given string and assigns it to the IdempotencyKey field.
func (o *AeAssetSellerChargeUpsert) SetIdempotencyKey(v string) {
	o.IdempotencyKey = &v
}

func (o AeAssetSellerChargeUpsert) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AeAssetSellerChargeUpsert) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["detail"] = o.Detail
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
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

func (o *AeAssetSellerChargeUpsert) UnmarshalJSON(data []byte) (err error) {
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

	varAeAssetSellerChargeUpsert := _AeAssetSellerChargeUpsert{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAeAssetSellerChargeUpsert)

	if err != nil {
		return err
	}

	*o = AeAssetSellerChargeUpsert(varAeAssetSellerChargeUpsert)

	return err
}

type NullableAeAssetSellerChargeUpsert struct {
	value *AeAssetSellerChargeUpsert
	isSet bool
}

func (v NullableAeAssetSellerChargeUpsert) Get() *AeAssetSellerChargeUpsert {
	return v.value
}

func (v *NullableAeAssetSellerChargeUpsert) Set(val *AeAssetSellerChargeUpsert) {
	v.value = val
	v.isSet = true
}

func (v NullableAeAssetSellerChargeUpsert) IsSet() bool {
	return v.isSet
}

func (v *NullableAeAssetSellerChargeUpsert) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAeAssetSellerChargeUpsert(val *AeAssetSellerChargeUpsert) *NullableAeAssetSellerChargeUpsert {
	return &NullableAeAssetSellerChargeUpsert{value: val, isSet: true}
}

func (v NullableAeAssetSellerChargeUpsert) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAeAssetSellerChargeUpsert) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


