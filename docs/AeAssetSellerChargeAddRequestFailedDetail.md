# AeAssetSellerChargeAddRequestFailedDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**Asset** | Pointer to [**CommonAmsAssetPointer**](CommonAmsAssetPointer.md) |  | [optional] 
**AssetId** | **string** | Source&#39;s unique identifier for asset | 
**Vin** | **string** | The Vehicle Identification Number(VIN) of the asset. | 
**Stock** | **string** | The stock number of the asset. | 
**Seller** | [**CommonAmsAccountPointer**](CommonAmsAccountPointer.md) |  | 
**ReferenceId** | **string** | Identifier that can be used to identify this event in a response event. | 
**Reason** | **string** | An explaination for why the charge could not be added. | 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeAssetSellerChargeAddRequestFailedDetail

`func NewAeAssetSellerChargeAddRequestFailedDetail(auctionId string, assetId string, vin string, stock string, seller CommonAmsAccountPointer, referenceId string, reason string, ) *AeAssetSellerChargeAddRequestFailedDetail`

NewAeAssetSellerChargeAddRequestFailedDetail instantiates a new AeAssetSellerChargeAddRequestFailedDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetSellerChargeAddRequestFailedDetailWithDefaults

`func NewAeAssetSellerChargeAddRequestFailedDetailWithDefaults() *AeAssetSellerChargeAddRequestFailedDetail`

NewAeAssetSellerChargeAddRequestFailedDetailWithDefaults instantiates a new AeAssetSellerChargeAddRequestFailedDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeAssetSellerChargeAddRequestFailedDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAssetSellerChargeAddRequestFailedDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAssetSellerChargeAddRequestFailedDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetAsset

`func (o *AeAssetSellerChargeAddRequestFailedDetail) GetAsset() CommonAmsAssetPointer`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *AeAssetSellerChargeAddRequestFailedDetail) GetAssetOk() (*CommonAmsAssetPointer, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *AeAssetSellerChargeAddRequestFailedDetail) SetAsset(v CommonAmsAssetPointer)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *AeAssetSellerChargeAddRequestFailedDetail) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetAssetId

`func (o *AeAssetSellerChargeAddRequestFailedDetail) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *AeAssetSellerChargeAddRequestFailedDetail) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *AeAssetSellerChargeAddRequestFailedDetail) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetVin

`func (o *AeAssetSellerChargeAddRequestFailedDetail) GetVin() string`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *AeAssetSellerChargeAddRequestFailedDetail) GetVinOk() (*string, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *AeAssetSellerChargeAddRequestFailedDetail) SetVin(v string)`

SetVin sets Vin field to given value.


### GetStock

`func (o *AeAssetSellerChargeAddRequestFailedDetail) GetStock() string`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *AeAssetSellerChargeAddRequestFailedDetail) GetStockOk() (*string, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *AeAssetSellerChargeAddRequestFailedDetail) SetStock(v string)`

SetStock sets Stock field to given value.


### GetSeller

`func (o *AeAssetSellerChargeAddRequestFailedDetail) GetSeller() CommonAmsAccountPointer`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *AeAssetSellerChargeAddRequestFailedDetail) GetSellerOk() (*CommonAmsAccountPointer, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *AeAssetSellerChargeAddRequestFailedDetail) SetSeller(v CommonAmsAccountPointer)`

SetSeller sets Seller field to given value.


### GetReferenceId

`func (o *AeAssetSellerChargeAddRequestFailedDetail) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *AeAssetSellerChargeAddRequestFailedDetail) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *AeAssetSellerChargeAddRequestFailedDetail) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetReason

`func (o *AeAssetSellerChargeAddRequestFailedDetail) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AeAssetSellerChargeAddRequestFailedDetail) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AeAssetSellerChargeAddRequestFailedDetail) SetReason(v string)`

SetReason sets Reason field to given value.


### GetInitiator

`func (o *AeAssetSellerChargeAddRequestFailedDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAssetSellerChargeAddRequestFailedDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAssetSellerChargeAddRequestFailedDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAssetSellerChargeAddRequestFailedDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


