# AeDocumentAssetAddedDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentId** | **string** | Unique identifier for Spark Document | 
**DocumentType** | [**DocumentType**](DocumentType.md) |  | 
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**AssetStock** | **string** | The stock number of the asset. | 
**AssetVin** | **string** | The Vehicle Identification Number(VIN) of the asset. | 
**ServiceId** | **string** | Id to access the service(SparkDocs) that will provide the method to retrieve the resource. | 
**RecordedAt** | **time.Time** | Date and time the document was added | 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeDocumentAssetAddedDetail

`func NewAeDocumentAssetAddedDetail(documentId string, documentType DocumentType, auctionId string, assetStock string, assetVin string, serviceId string, recordedAt time.Time, ) *AeDocumentAssetAddedDetail`

NewAeDocumentAssetAddedDetail instantiates a new AeDocumentAssetAddedDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeDocumentAssetAddedDetailWithDefaults

`func NewAeDocumentAssetAddedDetailWithDefaults() *AeDocumentAssetAddedDetail`

NewAeDocumentAssetAddedDetailWithDefaults instantiates a new AeDocumentAssetAddedDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentId

`func (o *AeDocumentAssetAddedDetail) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *AeDocumentAssetAddedDetail) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *AeDocumentAssetAddedDetail) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.


### GetDocumentType

`func (o *AeDocumentAssetAddedDetail) GetDocumentType() DocumentType`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *AeDocumentAssetAddedDetail) GetDocumentTypeOk() (*DocumentType, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *AeDocumentAssetAddedDetail) SetDocumentType(v DocumentType)`

SetDocumentType sets DocumentType field to given value.


### GetAuctionId

`func (o *AeDocumentAssetAddedDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeDocumentAssetAddedDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeDocumentAssetAddedDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetAssetStock

`func (o *AeDocumentAssetAddedDetail) GetAssetStock() string`

GetAssetStock returns the AssetStock field if non-nil, zero value otherwise.

### GetAssetStockOk

`func (o *AeDocumentAssetAddedDetail) GetAssetStockOk() (*string, bool)`

GetAssetStockOk returns a tuple with the AssetStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetStock

`func (o *AeDocumentAssetAddedDetail) SetAssetStock(v string)`

SetAssetStock sets AssetStock field to given value.


### GetAssetVin

`func (o *AeDocumentAssetAddedDetail) GetAssetVin() string`

GetAssetVin returns the AssetVin field if non-nil, zero value otherwise.

### GetAssetVinOk

`func (o *AeDocumentAssetAddedDetail) GetAssetVinOk() (*string, bool)`

GetAssetVinOk returns a tuple with the AssetVin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetVin

`func (o *AeDocumentAssetAddedDetail) SetAssetVin(v string)`

SetAssetVin sets AssetVin field to given value.


### GetServiceId

`func (o *AeDocumentAssetAddedDetail) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *AeDocumentAssetAddedDetail) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *AeDocumentAssetAddedDetail) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetRecordedAt

`func (o *AeDocumentAssetAddedDetail) GetRecordedAt() time.Time`

GetRecordedAt returns the RecordedAt field if non-nil, zero value otherwise.

### GetRecordedAtOk

`func (o *AeDocumentAssetAddedDetail) GetRecordedAtOk() (*time.Time, bool)`

GetRecordedAtOk returns a tuple with the RecordedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordedAt

`func (o *AeDocumentAssetAddedDetail) SetRecordedAt(v time.Time)`

SetRecordedAt sets RecordedAt field to given value.


### GetInitiator

`func (o *AeDocumentAssetAddedDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeDocumentAssetAddedDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeDocumentAssetAddedDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeDocumentAssetAddedDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


