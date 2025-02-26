# AeAssetDealOfferMadeDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**Asset** | Pointer to [**CommonAmsAssetPointer**](CommonAmsAssetPointer.md) |  | [optional] 
**DealId** | Pointer to **string** | Unique id for a deal | [optional] 
**Vin** | **string** | The Vehicle Identification Number(VIN) of the asset. | 
**Stock** | **string** | The stock number of the asset. | 
**Seller** | [**CommonAmsAccountPointer**](CommonAmsAccountPointer.md) |  | 
**Buyer** | [**CommonAmsAccountPointer**](CommonAmsAccountPointer.md) |  | 
**SaleDate** | **string** | The sale date of the deal. | 
**Lane** | **string** | The lane the asset was in when deal made. | 
**Lot** | **string** | The lot the asset was in when deal made. | 
**Amount** | **float32** | The amount the asset sold for. | 
**UpdatedAt** | Pointer to **time.Time** | The updated date time of the gatepass status | [optional] 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeAssetDealOfferMadeDetail

`func NewAeAssetDealOfferMadeDetail(auctionId string, vin string, stock string, seller CommonAmsAccountPointer, buyer CommonAmsAccountPointer, saleDate string, lane string, lot string, amount float32, ) *AeAssetDealOfferMadeDetail`

NewAeAssetDealOfferMadeDetail instantiates a new AeAssetDealOfferMadeDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetDealOfferMadeDetailWithDefaults

`func NewAeAssetDealOfferMadeDetailWithDefaults() *AeAssetDealOfferMadeDetail`

NewAeAssetDealOfferMadeDetailWithDefaults instantiates a new AeAssetDealOfferMadeDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeAssetDealOfferMadeDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAssetDealOfferMadeDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAssetDealOfferMadeDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetAsset

`func (o *AeAssetDealOfferMadeDetail) GetAsset() CommonAmsAssetPointer`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *AeAssetDealOfferMadeDetail) GetAssetOk() (*CommonAmsAssetPointer, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *AeAssetDealOfferMadeDetail) SetAsset(v CommonAmsAssetPointer)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *AeAssetDealOfferMadeDetail) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetDealId

`func (o *AeAssetDealOfferMadeDetail) GetDealId() string`

GetDealId returns the DealId field if non-nil, zero value otherwise.

### GetDealIdOk

`func (o *AeAssetDealOfferMadeDetail) GetDealIdOk() (*string, bool)`

GetDealIdOk returns a tuple with the DealId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealId

`func (o *AeAssetDealOfferMadeDetail) SetDealId(v string)`

SetDealId sets DealId field to given value.

### HasDealId

`func (o *AeAssetDealOfferMadeDetail) HasDealId() bool`

HasDealId returns a boolean if a field has been set.

### GetVin

`func (o *AeAssetDealOfferMadeDetail) GetVin() string`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *AeAssetDealOfferMadeDetail) GetVinOk() (*string, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *AeAssetDealOfferMadeDetail) SetVin(v string)`

SetVin sets Vin field to given value.


### GetStock

`func (o *AeAssetDealOfferMadeDetail) GetStock() string`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *AeAssetDealOfferMadeDetail) GetStockOk() (*string, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *AeAssetDealOfferMadeDetail) SetStock(v string)`

SetStock sets Stock field to given value.


### GetSeller

`func (o *AeAssetDealOfferMadeDetail) GetSeller() CommonAmsAccountPointer`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *AeAssetDealOfferMadeDetail) GetSellerOk() (*CommonAmsAccountPointer, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *AeAssetDealOfferMadeDetail) SetSeller(v CommonAmsAccountPointer)`

SetSeller sets Seller field to given value.


### GetBuyer

`func (o *AeAssetDealOfferMadeDetail) GetBuyer() CommonAmsAccountPointer`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *AeAssetDealOfferMadeDetail) GetBuyerOk() (*CommonAmsAccountPointer, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *AeAssetDealOfferMadeDetail) SetBuyer(v CommonAmsAccountPointer)`

SetBuyer sets Buyer field to given value.


### GetSaleDate

`func (o *AeAssetDealOfferMadeDetail) GetSaleDate() string`

GetSaleDate returns the SaleDate field if non-nil, zero value otherwise.

### GetSaleDateOk

`func (o *AeAssetDealOfferMadeDetail) GetSaleDateOk() (*string, bool)`

GetSaleDateOk returns a tuple with the SaleDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleDate

`func (o *AeAssetDealOfferMadeDetail) SetSaleDate(v string)`

SetSaleDate sets SaleDate field to given value.


### GetLane

`func (o *AeAssetDealOfferMadeDetail) GetLane() string`

GetLane returns the Lane field if non-nil, zero value otherwise.

### GetLaneOk

`func (o *AeAssetDealOfferMadeDetail) GetLaneOk() (*string, bool)`

GetLaneOk returns a tuple with the Lane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLane

`func (o *AeAssetDealOfferMadeDetail) SetLane(v string)`

SetLane sets Lane field to given value.


### GetLot

`func (o *AeAssetDealOfferMadeDetail) GetLot() string`

GetLot returns the Lot field if non-nil, zero value otherwise.

### GetLotOk

`func (o *AeAssetDealOfferMadeDetail) GetLotOk() (*string, bool)`

GetLotOk returns a tuple with the Lot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLot

`func (o *AeAssetDealOfferMadeDetail) SetLot(v string)`

SetLot sets Lot field to given value.


### GetAmount

`func (o *AeAssetDealOfferMadeDetail) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AeAssetDealOfferMadeDetail) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AeAssetDealOfferMadeDetail) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetUpdatedAt

`func (o *AeAssetDealOfferMadeDetail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AeAssetDealOfferMadeDetail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AeAssetDealOfferMadeDetail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AeAssetDealOfferMadeDetail) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetInitiator

`func (o *AeAssetDealOfferMadeDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAssetDealOfferMadeDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAssetDealOfferMadeDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAssetDealOfferMadeDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


