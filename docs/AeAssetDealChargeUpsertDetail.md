# AeAssetDealChargeUpsertDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**Asset** | Pointer to [**CommonAmsAssetPointer**](CommonAmsAssetPointer.md) |  | [optional] 
**DealId** | Pointer to **string** | Unique id for a deal | [optional] 
**AssetId** | **string** | Source&#39;s unique identifier for asset | 
**Stock** | **string** | The stock number of the asset. | 
**Vin** | **string** | The Vehicle Identification Number(VIN) of the asset. | 
**Account** | [**CommonAmsAccountPointer**](CommonAmsAccountPointer.md) |  | 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 
**Charge** | [**AeAssetDealChargeUpsertDetailCharge**](AeAssetDealChargeUpsertDetailCharge.md) |  | 

## Methods

### NewAeAssetDealChargeUpsertDetail

`func NewAeAssetDealChargeUpsertDetail(auctionId string, assetId string, stock string, vin string, account CommonAmsAccountPointer, charge AeAssetDealChargeUpsertDetailCharge, ) *AeAssetDealChargeUpsertDetail`

NewAeAssetDealChargeUpsertDetail instantiates a new AeAssetDealChargeUpsertDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetDealChargeUpsertDetailWithDefaults

`func NewAeAssetDealChargeUpsertDetailWithDefaults() *AeAssetDealChargeUpsertDetail`

NewAeAssetDealChargeUpsertDetailWithDefaults instantiates a new AeAssetDealChargeUpsertDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeAssetDealChargeUpsertDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAssetDealChargeUpsertDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAssetDealChargeUpsertDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetAsset

`func (o *AeAssetDealChargeUpsertDetail) GetAsset() CommonAmsAssetPointer`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *AeAssetDealChargeUpsertDetail) GetAssetOk() (*CommonAmsAssetPointer, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *AeAssetDealChargeUpsertDetail) SetAsset(v CommonAmsAssetPointer)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *AeAssetDealChargeUpsertDetail) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetDealId

`func (o *AeAssetDealChargeUpsertDetail) GetDealId() string`

GetDealId returns the DealId field if non-nil, zero value otherwise.

### GetDealIdOk

`func (o *AeAssetDealChargeUpsertDetail) GetDealIdOk() (*string, bool)`

GetDealIdOk returns a tuple with the DealId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealId

`func (o *AeAssetDealChargeUpsertDetail) SetDealId(v string)`

SetDealId sets DealId field to given value.

### HasDealId

`func (o *AeAssetDealChargeUpsertDetail) HasDealId() bool`

HasDealId returns a boolean if a field has been set.

### GetAssetId

`func (o *AeAssetDealChargeUpsertDetail) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *AeAssetDealChargeUpsertDetail) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *AeAssetDealChargeUpsertDetail) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetStock

`func (o *AeAssetDealChargeUpsertDetail) GetStock() string`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *AeAssetDealChargeUpsertDetail) GetStockOk() (*string, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *AeAssetDealChargeUpsertDetail) SetStock(v string)`

SetStock sets Stock field to given value.


### GetVin

`func (o *AeAssetDealChargeUpsertDetail) GetVin() string`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *AeAssetDealChargeUpsertDetail) GetVinOk() (*string, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *AeAssetDealChargeUpsertDetail) SetVin(v string)`

SetVin sets Vin field to given value.


### GetAccount

`func (o *AeAssetDealChargeUpsertDetail) GetAccount() CommonAmsAccountPointer`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AeAssetDealChargeUpsertDetail) GetAccountOk() (*CommonAmsAccountPointer, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AeAssetDealChargeUpsertDetail) SetAccount(v CommonAmsAccountPointer)`

SetAccount sets Account field to given value.


### GetInitiator

`func (o *AeAssetDealChargeUpsertDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAssetDealChargeUpsertDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAssetDealChargeUpsertDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAssetDealChargeUpsertDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.

### GetCharge

`func (o *AeAssetDealChargeUpsertDetail) GetCharge() AeAssetDealChargeUpsertDetailCharge`

GetCharge returns the Charge field if non-nil, zero value otherwise.

### GetChargeOk

`func (o *AeAssetDealChargeUpsertDetail) GetChargeOk() (*AeAssetDealChargeUpsertDetailCharge, bool)`

GetChargeOk returns a tuple with the Charge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharge

`func (o *AeAssetDealChargeUpsertDetail) SetCharge(v AeAssetDealChargeUpsertDetailCharge)`

SetCharge sets Charge field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


