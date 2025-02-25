# AeDocumentAssetRemovedDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentId** | **string** | Unique identifier for Spark Document | 
**DocumentType** | [**DocumentType**](DocumentType.md) |  | 
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**AssetStock** | **string** | The stock number of the asset. | 
**AssetVin** | **string** | The Vehicle Identification Number(VIN) of the asset. | 
**RecordedAt** | **time.Time** | Date and time the document was added | 

## Methods

### NewAeDocumentAssetRemovedDetail

`func NewAeDocumentAssetRemovedDetail(documentId string, documentType DocumentType, auctionId string, assetStock string, assetVin string, recordedAt time.Time, ) *AeDocumentAssetRemovedDetail`

NewAeDocumentAssetRemovedDetail instantiates a new AeDocumentAssetRemovedDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeDocumentAssetRemovedDetailWithDefaults

`func NewAeDocumentAssetRemovedDetailWithDefaults() *AeDocumentAssetRemovedDetail`

NewAeDocumentAssetRemovedDetailWithDefaults instantiates a new AeDocumentAssetRemovedDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentId

`func (o *AeDocumentAssetRemovedDetail) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *AeDocumentAssetRemovedDetail) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *AeDocumentAssetRemovedDetail) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.


### GetDocumentType

`func (o *AeDocumentAssetRemovedDetail) GetDocumentType() DocumentType`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *AeDocumentAssetRemovedDetail) GetDocumentTypeOk() (*DocumentType, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *AeDocumentAssetRemovedDetail) SetDocumentType(v DocumentType)`

SetDocumentType sets DocumentType field to given value.


### GetAuctionId

`func (o *AeDocumentAssetRemovedDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeDocumentAssetRemovedDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeDocumentAssetRemovedDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetAssetStock

`func (o *AeDocumentAssetRemovedDetail) GetAssetStock() string`

GetAssetStock returns the AssetStock field if non-nil, zero value otherwise.

### GetAssetStockOk

`func (o *AeDocumentAssetRemovedDetail) GetAssetStockOk() (*string, bool)`

GetAssetStockOk returns a tuple with the AssetStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetStock

`func (o *AeDocumentAssetRemovedDetail) SetAssetStock(v string)`

SetAssetStock sets AssetStock field to given value.


### GetAssetVin

`func (o *AeDocumentAssetRemovedDetail) GetAssetVin() string`

GetAssetVin returns the AssetVin field if non-nil, zero value otherwise.

### GetAssetVinOk

`func (o *AeDocumentAssetRemovedDetail) GetAssetVinOk() (*string, bool)`

GetAssetVinOk returns a tuple with the AssetVin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetVin

`func (o *AeDocumentAssetRemovedDetail) SetAssetVin(v string)`

SetAssetVin sets AssetVin field to given value.


### GetRecordedAt

`func (o *AeDocumentAssetRemovedDetail) GetRecordedAt() time.Time`

GetRecordedAt returns the RecordedAt field if non-nil, zero value otherwise.

### GetRecordedAtOk

`func (o *AeDocumentAssetRemovedDetail) GetRecordedAtOk() (*time.Time, bool)`

GetRecordedAtOk returns a tuple with the RecordedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordedAt

`func (o *AeDocumentAssetRemovedDetail) SetRecordedAt(v time.Time)`

SetRecordedAt sets RecordedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


