# AeAssetDealPaymentApplySourceDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**AssetId** | Pointer to **string** | Source&#39;s unique identifier for asset | [optional] 
**Stock** | **string** | The stock number of the asset. | 
**Vin** | **string** | The Vehicle Identification Number(VIN) of the asset. | 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 
**Deal** | [**AeAssetDealPaymentApplySourceDetailDeal**](AeAssetDealPaymentApplySourceDetailDeal.md) |  | 

## Methods

### NewAeAssetDealPaymentApplySourceDetail

`func NewAeAssetDealPaymentApplySourceDetail(auctionId string, stock string, vin string, deal AeAssetDealPaymentApplySourceDetailDeal, ) *AeAssetDealPaymentApplySourceDetail`

NewAeAssetDealPaymentApplySourceDetail instantiates a new AeAssetDealPaymentApplySourceDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetDealPaymentApplySourceDetailWithDefaults

`func NewAeAssetDealPaymentApplySourceDetailWithDefaults() *AeAssetDealPaymentApplySourceDetail`

NewAeAssetDealPaymentApplySourceDetailWithDefaults instantiates a new AeAssetDealPaymentApplySourceDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeAssetDealPaymentApplySourceDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAssetDealPaymentApplySourceDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAssetDealPaymentApplySourceDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetAssetId

`func (o *AeAssetDealPaymentApplySourceDetail) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *AeAssetDealPaymentApplySourceDetail) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *AeAssetDealPaymentApplySourceDetail) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *AeAssetDealPaymentApplySourceDetail) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetStock

`func (o *AeAssetDealPaymentApplySourceDetail) GetStock() string`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *AeAssetDealPaymentApplySourceDetail) GetStockOk() (*string, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *AeAssetDealPaymentApplySourceDetail) SetStock(v string)`

SetStock sets Stock field to given value.


### GetVin

`func (o *AeAssetDealPaymentApplySourceDetail) GetVin() string`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *AeAssetDealPaymentApplySourceDetail) GetVinOk() (*string, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *AeAssetDealPaymentApplySourceDetail) SetVin(v string)`

SetVin sets Vin field to given value.


### GetInitiator

`func (o *AeAssetDealPaymentApplySourceDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAssetDealPaymentApplySourceDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAssetDealPaymentApplySourceDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAssetDealPaymentApplySourceDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.

### GetDeal

`func (o *AeAssetDealPaymentApplySourceDetail) GetDeal() AeAssetDealPaymentApplySourceDetailDeal`

GetDeal returns the Deal field if non-nil, zero value otherwise.

### GetDealOk

`func (o *AeAssetDealPaymentApplySourceDetail) GetDealOk() (*AeAssetDealPaymentApplySourceDetailDeal, bool)`

GetDealOk returns a tuple with the Deal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeal

`func (o *AeAssetDealPaymentApplySourceDetail) SetDeal(v AeAssetDealPaymentApplySourceDetailDeal)`

SetDeal sets Deal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


