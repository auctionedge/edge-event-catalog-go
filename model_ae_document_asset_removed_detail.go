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

// checks if the AeDocumentAssetRemovedDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AeDocumentAssetRemovedDetail{}

// AeDocumentAssetRemovedDetail Document removed and no long available
type AeDocumentAssetRemovedDetail struct {
	// Unique identifier for Spark Document
	DocumentId string `json:"document-id"`
	DocumentType DocumentType `json:"document-type"`
	// Auction Edge unique identifier for an auction.
	AuctionId string `json:"auction-id"`
	// The stock number of the asset.
	AssetStock string `json:"asset-stock"`
	// The Vehicle Identification Number(VIN) of the asset.
	AssetVin string `json:"asset-vin"`
	// Date and time the document was added
	RecordedAt time.Time `json:"recorded-at"`
}

type _AeDocumentAssetRemovedDetail AeDocumentAssetRemovedDetail

// NewAeDocumentAssetRemovedDetail instantiates a new AeDocumentAssetRemovedDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAeDocumentAssetRemovedDetail(documentId string, documentType DocumentType, auctionId string, assetStock string, assetVin string, recordedAt time.Time) *AeDocumentAssetRemovedDetail {
	this := AeDocumentAssetRemovedDetail{}
	this.DocumentId = documentId
	this.DocumentType = documentType
	this.AuctionId = auctionId
	this.AssetStock = assetStock
	this.AssetVin = assetVin
	this.RecordedAt = recordedAt
	return &this
}

// NewAeDocumentAssetRemovedDetailWithDefaults instantiates a new AeDocumentAssetRemovedDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAeDocumentAssetRemovedDetailWithDefaults() *AeDocumentAssetRemovedDetail {
	this := AeDocumentAssetRemovedDetail{}
	return &this
}

// GetDocumentId returns the DocumentId field value
func (o *AeDocumentAssetRemovedDetail) GetDocumentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value
// and a boolean to check if the value has been set.
func (o *AeDocumentAssetRemovedDetail) GetDocumentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocumentId, true
}

// SetDocumentId sets field value
func (o *AeDocumentAssetRemovedDetail) SetDocumentId(v string) {
	o.DocumentId = v
}

// GetDocumentType returns the DocumentType field value
func (o *AeDocumentAssetRemovedDetail) GetDocumentType() DocumentType {
	if o == nil {
		var ret DocumentType
		return ret
	}

	return o.DocumentType
}

// GetDocumentTypeOk returns a tuple with the DocumentType field value
// and a boolean to check if the value has been set.
func (o *AeDocumentAssetRemovedDetail) GetDocumentTypeOk() (*DocumentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocumentType, true
}

// SetDocumentType sets field value
func (o *AeDocumentAssetRemovedDetail) SetDocumentType(v DocumentType) {
	o.DocumentType = v
}

// GetAuctionId returns the AuctionId field value
func (o *AeDocumentAssetRemovedDetail) GetAuctionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuctionId
}

// GetAuctionIdOk returns a tuple with the AuctionId field value
// and a boolean to check if the value has been set.
func (o *AeDocumentAssetRemovedDetail) GetAuctionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuctionId, true
}

// SetAuctionId sets field value
func (o *AeDocumentAssetRemovedDetail) SetAuctionId(v string) {
	o.AuctionId = v
}

// GetAssetStock returns the AssetStock field value
func (o *AeDocumentAssetRemovedDetail) GetAssetStock() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetStock
}

// GetAssetStockOk returns a tuple with the AssetStock field value
// and a boolean to check if the value has been set.
func (o *AeDocumentAssetRemovedDetail) GetAssetStockOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetStock, true
}

// SetAssetStock sets field value
func (o *AeDocumentAssetRemovedDetail) SetAssetStock(v string) {
	o.AssetStock = v
}

// GetAssetVin returns the AssetVin field value
func (o *AeDocumentAssetRemovedDetail) GetAssetVin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetVin
}

// GetAssetVinOk returns a tuple with the AssetVin field value
// and a boolean to check if the value has been set.
func (o *AeDocumentAssetRemovedDetail) GetAssetVinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetVin, true
}

// SetAssetVin sets field value
func (o *AeDocumentAssetRemovedDetail) SetAssetVin(v string) {
	o.AssetVin = v
}

// GetRecordedAt returns the RecordedAt field value
func (o *AeDocumentAssetRemovedDetail) GetRecordedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.RecordedAt
}

// GetRecordedAtOk returns a tuple with the RecordedAt field value
// and a boolean to check if the value has been set.
func (o *AeDocumentAssetRemovedDetail) GetRecordedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecordedAt, true
}

// SetRecordedAt sets field value
func (o *AeDocumentAssetRemovedDetail) SetRecordedAt(v time.Time) {
	o.RecordedAt = v
}

func (o AeDocumentAssetRemovedDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AeDocumentAssetRemovedDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["document-id"] = o.DocumentId
	toSerialize["document-type"] = o.DocumentType
	toSerialize["auction-id"] = o.AuctionId
	toSerialize["asset-stock"] = o.AssetStock
	toSerialize["asset-vin"] = o.AssetVin
	toSerialize["recorded-at"] = o.RecordedAt
	return toSerialize, nil
}

func (o *AeDocumentAssetRemovedDetail) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"document-id",
		"document-type",
		"auction-id",
		"asset-stock",
		"asset-vin",
		"recorded-at",
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

	varAeDocumentAssetRemovedDetail := _AeDocumentAssetRemovedDetail{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAeDocumentAssetRemovedDetail)

	if err != nil {
		return err
	}

	*o = AeDocumentAssetRemovedDetail(varAeDocumentAssetRemovedDetail)

	return err
}

type NullableAeDocumentAssetRemovedDetail struct {
	value *AeDocumentAssetRemovedDetail
	isSet bool
}

func (v NullableAeDocumentAssetRemovedDetail) Get() *AeDocumentAssetRemovedDetail {
	return v.value
}

func (v *NullableAeDocumentAssetRemovedDetail) Set(val *AeDocumentAssetRemovedDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableAeDocumentAssetRemovedDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableAeDocumentAssetRemovedDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAeDocumentAssetRemovedDetail(val *AeDocumentAssetRemovedDetail) *NullableAeDocumentAssetRemovedDetail {
	return &NullableAeDocumentAssetRemovedDetail{value: val, isSet: true}
}

func (v NullableAeDocumentAssetRemovedDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAeDocumentAssetRemovedDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


