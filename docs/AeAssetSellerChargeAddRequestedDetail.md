# AeAssetSellerChargeAddRequestedDetail

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
**Charge** | [**AeAssetSellerChargeAddRequestedDetailCharge**](AeAssetSellerChargeAddRequestedDetailCharge.md) |  | 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeAssetSellerChargeAddRequestedDetail

`func NewAeAssetSellerChargeAddRequestedDetail(auctionId string, assetId string, vin string, stock string, seller CommonAmsAccountPointer, referenceId string, charge AeAssetSellerChargeAddRequestedDetailCharge, ) *AeAssetSellerChargeAddRequestedDetail`

NewAeAssetSellerChargeAddRequestedDetail instantiates a new AeAssetSellerChargeAddRequestedDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetSellerChargeAddRequestedDetailWithDefaults

`func NewAeAssetSellerChargeAddRequestedDetailWithDefaults() *AeAssetSellerChargeAddRequestedDetail`

NewAeAssetSellerChargeAddRequestedDetailWithDefaults instantiates a new AeAssetSellerChargeAddRequestedDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeAssetSellerChargeAddRequestedDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAssetSellerChargeAddRequestedDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAssetSellerChargeAddRequestedDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetAsset

`func (o *AeAssetSellerChargeAddRequestedDetail) GetAsset() CommonAmsAssetPointer`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *AeAssetSellerChargeAddRequestedDetail) GetAssetOk() (*CommonAmsAssetPointer, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *AeAssetSellerChargeAddRequestedDetail) SetAsset(v CommonAmsAssetPointer)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *AeAssetSellerChargeAddRequestedDetail) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetAssetId

`func (o *AeAssetSellerChargeAddRequestedDetail) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *AeAssetSellerChargeAddRequestedDetail) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *AeAssetSellerChargeAddRequestedDetail) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetVin

`func (o *AeAssetSellerChargeAddRequestedDetail) GetVin() string`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *AeAssetSellerChargeAddRequestedDetail) GetVinOk() (*string, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *AeAssetSellerChargeAddRequestedDetail) SetVin(v string)`

SetVin sets Vin field to given value.


### GetStock

`func (o *AeAssetSellerChargeAddRequestedDetail) GetStock() string`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *AeAssetSellerChargeAddRequestedDetail) GetStockOk() (*string, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *AeAssetSellerChargeAddRequestedDetail) SetStock(v string)`

SetStock sets Stock field to given value.


### GetSeller

`func (o *AeAssetSellerChargeAddRequestedDetail) GetSeller() CommonAmsAccountPointer`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *AeAssetSellerChargeAddRequestedDetail) GetSellerOk() (*CommonAmsAccountPointer, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *AeAssetSellerChargeAddRequestedDetail) SetSeller(v CommonAmsAccountPointer)`

SetSeller sets Seller field to given value.


### GetReferenceId

`func (o *AeAssetSellerChargeAddRequestedDetail) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *AeAssetSellerChargeAddRequestedDetail) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *AeAssetSellerChargeAddRequestedDetail) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetCharge

`func (o *AeAssetSellerChargeAddRequestedDetail) GetCharge() AeAssetSellerChargeAddRequestedDetailCharge`

GetCharge returns the Charge field if non-nil, zero value otherwise.

### GetChargeOk

`func (o *AeAssetSellerChargeAddRequestedDetail) GetChargeOk() (*AeAssetSellerChargeAddRequestedDetailCharge, bool)`

GetChargeOk returns a tuple with the Charge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharge

`func (o *AeAssetSellerChargeAddRequestedDetail) SetCharge(v AeAssetSellerChargeAddRequestedDetailCharge)`

SetCharge sets Charge field to given value.


### GetInitiator

`func (o *AeAssetSellerChargeAddRequestedDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAssetSellerChargeAddRequestedDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAssetSellerChargeAddRequestedDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAssetSellerChargeAddRequestedDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


