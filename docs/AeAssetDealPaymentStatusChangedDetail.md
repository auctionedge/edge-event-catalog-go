# AeAssetDealPaymentStatusChangedDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**AssetId** | **string** | Source&#39;s unique identifier for asset | 
**Stock** | **string** | The stock number of the asset. | 
**Vin** | **string** | The Vehicle Identification Number(VIN) of the asset. | 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 
**Deal** | [**AeAssetDealPaymentStatusChangedDetailDeal**](AeAssetDealPaymentStatusChangedDetailDeal.md) |  | 

## Methods

### NewAeAssetDealPaymentStatusChangedDetail

`func NewAeAssetDealPaymentStatusChangedDetail(auctionId string, assetId string, stock string, vin string, deal AeAssetDealPaymentStatusChangedDetailDeal, ) *AeAssetDealPaymentStatusChangedDetail`

NewAeAssetDealPaymentStatusChangedDetail instantiates a new AeAssetDealPaymentStatusChangedDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetDealPaymentStatusChangedDetailWithDefaults

`func NewAeAssetDealPaymentStatusChangedDetailWithDefaults() *AeAssetDealPaymentStatusChangedDetail`

NewAeAssetDealPaymentStatusChangedDetailWithDefaults instantiates a new AeAssetDealPaymentStatusChangedDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeAssetDealPaymentStatusChangedDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAssetDealPaymentStatusChangedDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAssetDealPaymentStatusChangedDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetAssetId

`func (o *AeAssetDealPaymentStatusChangedDetail) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *AeAssetDealPaymentStatusChangedDetail) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *AeAssetDealPaymentStatusChangedDetail) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetStock

`func (o *AeAssetDealPaymentStatusChangedDetail) GetStock() string`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *AeAssetDealPaymentStatusChangedDetail) GetStockOk() (*string, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *AeAssetDealPaymentStatusChangedDetail) SetStock(v string)`

SetStock sets Stock field to given value.


### GetVin

`func (o *AeAssetDealPaymentStatusChangedDetail) GetVin() string`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *AeAssetDealPaymentStatusChangedDetail) GetVinOk() (*string, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *AeAssetDealPaymentStatusChangedDetail) SetVin(v string)`

SetVin sets Vin field to given value.


### GetInitiator

`func (o *AeAssetDealPaymentStatusChangedDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAssetDealPaymentStatusChangedDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAssetDealPaymentStatusChangedDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAssetDealPaymentStatusChangedDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.

### GetDeal

`func (o *AeAssetDealPaymentStatusChangedDetail) GetDeal() AeAssetDealPaymentStatusChangedDetailDeal`

GetDeal returns the Deal field if non-nil, zero value otherwise.

### GetDealOk

`func (o *AeAssetDealPaymentStatusChangedDetail) GetDealOk() (*AeAssetDealPaymentStatusChangedDetailDeal, bool)`

GetDealOk returns a tuple with the Deal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeal

`func (o *AeAssetDealPaymentStatusChangedDetail) SetDeal(v AeAssetDealPaymentStatusChangedDetailDeal)`

SetDeal sets Deal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


