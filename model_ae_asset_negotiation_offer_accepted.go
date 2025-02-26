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

// checks if the AeAssetNegotiationOfferAccepted type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AeAssetNegotiationOfferAccepted{}

// AeAssetNegotiationOfferAccepted Notification that a service order has been placed
type AeAssetNegotiationOfferAccepted struct {
	Detail AeAssetNegotiationOfferAcceptedDetail `json:"detail"`
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

type _AeAssetNegotiationOfferAccepted AeAssetNegotiationOfferAccepted

// NewAeAssetNegotiationOfferAccepted instantiates a new AeAssetNegotiationOfferAccepted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAeAssetNegotiationOfferAccepted(detail AeAssetNegotiationOfferAcceptedDetail, detailType string, id string, source string) *AeAssetNegotiationOfferAccepted {
	this := AeAssetNegotiationOfferAccepted{}
	this.DetailType = detailType
	this.Id = id
	this.Source = source
	var deprecated bool = false
	this.Deprecated = &deprecated
	return &this
}

// NewAeAssetNegotiationOfferAcceptedWithDefaults instantiates a new AeAssetNegotiationOfferAccepted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAeAssetNegotiationOfferAcceptedWithDefaults() *AeAssetNegotiationOfferAccepted {
	this := AeAssetNegotiationOfferAccepted{}
	var deprecated bool = false
	this.Deprecated = &deprecated
	return &this
}

// GetDetail returns the Detail field value
func (o *AeAssetNegotiationOfferAccepted) GetDetail() AeAssetNegotiationOfferAcceptedDetail {
	if o == nil {
		var ret AeAssetNegotiationOfferAcceptedDetail
		return ret
	}

	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *AeAssetNegotiationOfferAccepted) GetDetailOk() (*AeAssetNegotiationOfferAcceptedDetail, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Detail, true
}

// SetDetail sets field value
func (o *AeAssetNegotiationOfferAccepted) SetDetail(v AeAssetNegotiationOfferAcceptedDetail) {
	o.Detail = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *AeAssetNegotiationOfferAccepted) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAssetNegotiationOfferAccepted) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *AeAssetNegotiationOfferAccepted) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *AeAssetNegotiationOfferAccepted) SetVersion(v string) {
	o.Version = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *AeAssetNegotiationOfferAccepted) GetAccount() string {
	if o == nil || IsNil(o.Account) {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAssetNegotiationOfferAccepted) GetAccountOk() (*string, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *AeAssetNegotiationOfferAccepted) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *AeAssetNegotiationOfferAccepted) SetAccount(v string) {
	o.Account = &v
}

// GetDetailType returns the DetailType field value
func (o *AeAssetNegotiationOfferAccepted) GetDetailType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DetailType
}

// GetDetailTypeOk returns a tuple with the DetailType field value
// and a boolean to check if the value has been set.
func (o *AeAssetNegotiationOfferAccepted) GetDetailTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DetailType, true
}

// SetDetailType sets field value
func (o *AeAssetNegotiationOfferAccepted) SetDetailType(v string) {
	o.DetailType = v
}

// GetId returns the Id field value
func (o *AeAssetNegotiationOfferAccepted) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AeAssetNegotiationOfferAccepted) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AeAssetNegotiationOfferAccepted) SetId(v string) {
	o.Id = v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *AeAssetNegotiationOfferAccepted) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAssetNegotiationOfferAccepted) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *AeAssetNegotiationOfferAccepted) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *AeAssetNegotiationOfferAccepted) SetRegion(v string) {
	o.Region = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *AeAssetNegotiationOfferAccepted) GetResources() []string {
	if o == nil || IsNil(o.Resources) {
		var ret []string
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAssetNegotiationOfferAccepted) GetResourcesOk() ([]string, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *AeAssetNegotiationOfferAccepted) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []string and assigns it to the Resources field.
func (o *AeAssetNegotiationOfferAccepted) SetResources(v []string) {
	o.Resources = v
}

// GetSource returns the Source field value
func (o *AeAssetNegotiationOfferAccepted) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *AeAssetNegotiationOfferAccepted) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *AeAssetNegotiationOfferAccepted) SetSource(v string) {
	o.Source = v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *AeAssetNegotiationOfferAccepted) GetTime() time.Time {
	if o == nil || IsNil(o.Time) {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAssetNegotiationOfferAccepted) GetTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *AeAssetNegotiationOfferAccepted) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *AeAssetNegotiationOfferAccepted) SetTime(v time.Time) {
	o.Time = &v
}

// GetAuctionId returns the AuctionId field value if set, zero value otherwise.
func (o *AeAssetNegotiationOfferAccepted) GetAuctionId() string {
	if o == nil || IsNil(o.AuctionId) {
		var ret string
		return ret
	}
	return *o.AuctionId
}

// GetAuctionIdOk returns a tuple with the AuctionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAssetNegotiationOfferAccepted) GetAuctionIdOk() (*string, bool) {
	if o == nil || IsNil(o.AuctionId) {
		return nil, false
	}
	return o.AuctionId, true
}

// HasAuctionId returns a boolean if a field has been set.
func (o *AeAssetNegotiationOfferAccepted) HasAuctionId() bool {
	if o != nil && !IsNil(o.AuctionId) {
		return true
	}

	return false
}

// SetAuctionId gets a reference to the given string and assigns it to the AuctionId field.
func (o *AeAssetNegotiationOfferAccepted) SetAuctionId(v string) {
	o.AuctionId = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *AeAssetNegotiationOfferAccepted) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAssetNegotiationOfferAccepted) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *AeAssetNegotiationOfferAccepted) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *AeAssetNegotiationOfferAccepted) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetDeprecated returns the Deprecated field value if set, zero value otherwise.
func (o *AeAssetNegotiationOfferAccepted) GetDeprecated() bool {
	if o == nil || IsNil(o.Deprecated) {
		var ret bool
		return ret
	}
	return *o.Deprecated
}

// GetDeprecatedOk returns a tuple with the Deprecated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAssetNegotiationOfferAccepted) GetDeprecatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Deprecated) {
		return nil, false
	}
	return o.Deprecated, true
}

// HasDeprecated returns a boolean if a field has been set.
func (o *AeAssetNegotiationOfferAccepted) HasDeprecated() bool {
	if o != nil && !IsNil(o.Deprecated) {
		return true
	}

	return false
}

// SetDeprecated gets a reference to the given bool and assigns it to the Deprecated field.
func (o *AeAssetNegotiationOfferAccepted) SetDeprecated(v bool) {
	o.Deprecated = &v
}

// GetIdempotencyKey returns the IdempotencyKey field value if set, zero value otherwise.
func (o *AeAssetNegotiationOfferAccepted) GetIdempotencyKey() string {
	if o == nil || IsNil(o.IdempotencyKey) {
		var ret string
		return ret
	}
	return *o.IdempotencyKey
}

// GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAssetNegotiationOfferAccepted) GetIdempotencyKeyOk() (*string, bool) {
	if o == nil || IsNil(o.IdempotencyKey) {
		return nil, false
	}
	return o.IdempotencyKey, true
}

// HasIdempotencyKey returns a boolean if a field has been set.
func (o *AeAssetNegotiationOfferAccepted) HasIdempotencyKey() bool {
	if o != nil && !IsNil(o.IdempotencyKey) {
		return true
	}

	return false
}

// SetIdempotencyKey gets a reference to the given string and assigns it to the IdempotencyKey field.
func (o *AeAssetNegotiationOfferAccepted) SetIdempotencyKey(v string) {
	o.IdempotencyKey = &v
}

func (o AeAssetNegotiationOfferAccepted) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AeAssetNegotiationOfferAccepted) ToMap() (map[string]interface{}, error) {
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

func (o *AeAssetNegotiationOfferAccepted) UnmarshalJSON(data []byte) (err error) {
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

	varAeAssetNegotiationOfferAccepted := _AeAssetNegotiationOfferAccepted{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAeAssetNegotiationOfferAccepted)

	if err != nil {
		return err
	}

	*o = AeAssetNegotiationOfferAccepted(varAeAssetNegotiationOfferAccepted)

	return err
}

type NullableAeAssetNegotiationOfferAccepted struct {
	value *AeAssetNegotiationOfferAccepted
	isSet bool
}

func (v NullableAeAssetNegotiationOfferAccepted) Get() *AeAssetNegotiationOfferAccepted {
	return v.value
}

func (v *NullableAeAssetNegotiationOfferAccepted) Set(val *AeAssetNegotiationOfferAccepted) {
	v.value = val
	v.isSet = true
}

func (v NullableAeAssetNegotiationOfferAccepted) IsSet() bool {
	return v.isSet
}

func (v *NullableAeAssetNegotiationOfferAccepted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAeAssetNegotiationOfferAccepted(val *AeAssetNegotiationOfferAccepted) *NullableAeAssetNegotiationOfferAccepted {
	return &NullableAeAssetNegotiationOfferAccepted{value: val, isSet: true}
}

func (v NullableAeAssetNegotiationOfferAccepted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAeAssetNegotiationOfferAccepted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


