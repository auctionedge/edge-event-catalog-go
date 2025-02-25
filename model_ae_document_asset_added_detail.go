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

// checks if the AeDocumentAssetAddedDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AeDocumentAssetAddedDetail{}

// AeDocumentAssetAddedDetail Document uploaded and available
type AeDocumentAssetAddedDetail struct {
	// Unique identifier for Spark Document
	DocumentId string `json:"document-id"`
	DocumentType DocumentType `json:"document-type"`
	// Auction Edge unique identifier for an auction.
	AuctionId string `json:"auction-id"`
	// The stock number of the asset.
	AssetStock string `json:"asset-stock"`
	// The Vehicle Identification Number(VIN) of the asset.
	AssetVin string `json:"asset-vin"`
	// Id to access the service(SparkDocs) that will provide the method to retrieve the resource.
	ServiceId string `json:"service-id"`
	// Date and time the document was added
	RecordedAt time.Time `json:"recorded-at"`
	Initiator *CommonInitiator `json:"initiator,omitempty"`
}

type _AeDocumentAssetAddedDetail AeDocumentAssetAddedDetail

// NewAeDocumentAssetAddedDetail instantiates a new AeDocumentAssetAddedDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAeDocumentAssetAddedDetail(documentId string, documentType DocumentType, auctionId string, assetStock string, assetVin string, serviceId string, recordedAt time.Time) *AeDocumentAssetAddedDetail {
	this := AeDocumentAssetAddedDetail{}
	this.DocumentId = documentId
	this.DocumentType = documentType
	this.AuctionId = auctionId
	this.AssetStock = assetStock
	this.AssetVin = assetVin
	this.ServiceId = serviceId
	this.RecordedAt = recordedAt
	return &this
}

// NewAeDocumentAssetAddedDetailWithDefaults instantiates a new AeDocumentAssetAddedDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAeDocumentAssetAddedDetailWithDefaults() *AeDocumentAssetAddedDetail {
	this := AeDocumentAssetAddedDetail{}
	return &this
}

// GetDocumentId returns the DocumentId field value
func (o *AeDocumentAssetAddedDetail) GetDocumentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value
// and a boolean to check if the value has been set.
func (o *AeDocumentAssetAddedDetail) GetDocumentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocumentId, true
}

// SetDocumentId sets field value
func (o *AeDocumentAssetAddedDetail) SetDocumentId(v string) {
	o.DocumentId = v
}

// GetDocumentType returns the DocumentType field value
func (o *AeDocumentAssetAddedDetail) GetDocumentType() DocumentType {
	if o == nil {
		var ret DocumentType
		return ret
	}

	return o.DocumentType
}

// GetDocumentTypeOk returns a tuple with the DocumentType field value
// and a boolean to check if the value has been set.
func (o *AeDocumentAssetAddedDetail) GetDocumentTypeOk() (*DocumentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocumentType, true
}

// SetDocumentType sets field value
func (o *AeDocumentAssetAddedDetail) SetDocumentType(v DocumentType) {
	o.DocumentType = v
}

// GetAuctionId returns the AuctionId field value
func (o *AeDocumentAssetAddedDetail) GetAuctionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuctionId
}

// GetAuctionIdOk returns a tuple with the AuctionId field value
// and a boolean to check if the value has been set.
func (o *AeDocumentAssetAddedDetail) GetAuctionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuctionId, true
}

// SetAuctionId sets field value
func (o *AeDocumentAssetAddedDetail) SetAuctionId(v string) {
	o.AuctionId = v
}

// GetAssetStock returns the AssetStock field value
func (o *AeDocumentAssetAddedDetail) GetAssetStock() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetStock
}

// GetAssetStockOk returns a tuple with the AssetStock field value
// and a boolean to check if the value has been set.
func (o *AeDocumentAssetAddedDetail) GetAssetStockOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetStock, true
}

// SetAssetStock sets field value
func (o *AeDocumentAssetAddedDetail) SetAssetStock(v string) {
	o.AssetStock = v
}

// GetAssetVin returns the AssetVin field value
func (o *AeDocumentAssetAddedDetail) GetAssetVin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetVin
}

// GetAssetVinOk returns a tuple with the AssetVin field value
// and a boolean to check if the value has been set.
func (o *AeDocumentAssetAddedDetail) GetAssetVinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetVin, true
}

// SetAssetVin sets field value
func (o *AeDocumentAssetAddedDetail) SetAssetVin(v string) {
	o.AssetVin = v
}

// GetServiceId returns the ServiceId field value
func (o *AeDocumentAssetAddedDetail) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *AeDocumentAssetAddedDetail) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *AeDocumentAssetAddedDetail) SetServiceId(v string) {
	o.ServiceId = v
}

// GetRecordedAt returns the RecordedAt field value
func (o *AeDocumentAssetAddedDetail) GetRecordedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.RecordedAt
}

// GetRecordedAtOk returns a tuple with the RecordedAt field value
// and a boolean to check if the value has been set.
func (o *AeDocumentAssetAddedDetail) GetRecordedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecordedAt, true
}

// SetRecordedAt sets field value
func (o *AeDocumentAssetAddedDetail) SetRecordedAt(v time.Time) {
	o.RecordedAt = v
}

// GetInitiator returns the Initiator field value if set, zero value otherwise.
func (o *AeDocumentAssetAddedDetail) GetInitiator() CommonInitiator {
	if o == nil || IsNil(o.Initiator) {
		var ret CommonInitiator
		return ret
	}
	return *o.Initiator
}

// GetInitiatorOk returns a tuple with the Initiator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeDocumentAssetAddedDetail) GetInitiatorOk() (*CommonInitiator, bool) {
	if o == nil || IsNil(o.Initiator) {
		return nil, false
	}
	return o.Initiator, true
}

// HasInitiator returns a boolean if a field has been set.
func (o *AeDocumentAssetAddedDetail) HasInitiator() bool {
	if o != nil && !IsNil(o.Initiator) {
		return true
	}

	return false
}

// SetInitiator gets a reference to the given CommonInitiator and assigns it to the Initiator field.
func (o *AeDocumentAssetAddedDetail) SetInitiator(v CommonInitiator) {
	o.Initiator = &v
}

func (o AeDocumentAssetAddedDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AeDocumentAssetAddedDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["document-id"] = o.DocumentId
	toSerialize["document-type"] = o.DocumentType
	toSerialize["auction-id"] = o.AuctionId
	toSerialize["asset-stock"] = o.AssetStock
	toSerialize["asset-vin"] = o.AssetVin
	toSerialize["service-id"] = o.ServiceId
	toSerialize["recorded-at"] = o.RecordedAt
	if !IsNil(o.Initiator) {
		toSerialize["initiator"] = o.Initiator
	}
	return toSerialize, nil
}

func (o *AeDocumentAssetAddedDetail) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"document-id",
		"document-type",
		"auction-id",
		"asset-stock",
		"asset-vin",
		"service-id",
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

	varAeDocumentAssetAddedDetail := _AeDocumentAssetAddedDetail{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAeDocumentAssetAddedDetail)

	if err != nil {
		return err
	}

	*o = AeDocumentAssetAddedDetail(varAeDocumentAssetAddedDetail)

	return err
}

type NullableAeDocumentAssetAddedDetail struct {
	value *AeDocumentAssetAddedDetail
	isSet bool
}

func (v NullableAeDocumentAssetAddedDetail) Get() *AeDocumentAssetAddedDetail {
	return v.value
}

func (v *NullableAeDocumentAssetAddedDetail) Set(val *AeDocumentAssetAddedDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableAeDocumentAssetAddedDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableAeDocumentAssetAddedDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAeDocumentAssetAddedDetail(val *AeDocumentAssetAddedDetail) *NullableAeDocumentAssetAddedDetail {
	return &NullableAeDocumentAssetAddedDetail{value: val, isSet: true}
}

func (v NullableAeDocumentAssetAddedDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAeDocumentAssetAddedDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


