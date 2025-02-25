# AeAssetDealOfferAcceptedDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**Asset** | [**CommonAmsAssetPointer**](CommonAmsAssetPointer.md) |  | 
**DealId** | Pointer to **string** | Unique id for a deal | [optional] 
**Seller** | [**CommonAmsAccountPointer**](CommonAmsAccountPointer.md) |  | 
**Buyer** | [**CommonAmsAccountPointer**](CommonAmsAccountPointer.md) |  | 
**SaleDate** | **string** | The sale date of the deal. | 
**Lane** | **string** | The lane the asset was in when deal made. | 
**Lot** | **string** | The lot the asset was in when deal made. | 
**Amount** | Pointer to **float32** | The amount the asset sold for. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The updated date time of the deal. | [optional] 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeAssetDealOfferAcceptedDetail

`func NewAeAssetDealOfferAcceptedDetail(auctionId string, asset CommonAmsAssetPointer, seller CommonAmsAccountPointer, buyer CommonAmsAccountPointer, saleDate string, lane string, lot string, ) *AeAssetDealOfferAcceptedDetail`

NewAeAssetDealOfferAcceptedDetail instantiates a new AeAssetDealOfferAcceptedDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetDealOfferAcceptedDetailWithDefaults

`func NewAeAssetDealOfferAcceptedDetailWithDefaults() *AeAssetDealOfferAcceptedDetail`

NewAeAssetDealOfferAcceptedDetailWithDefaults instantiates a new AeAssetDealOfferAcceptedDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeAssetDealOfferAcceptedDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAssetDealOfferAcceptedDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAssetDealOfferAcceptedDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetAsset

`func (o *AeAssetDealOfferAcceptedDetail) GetAsset() CommonAmsAssetPointer`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *AeAssetDealOfferAcceptedDetail) GetAssetOk() (*CommonAmsAssetPointer, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *AeAssetDealOfferAcceptedDetail) SetAsset(v CommonAmsAssetPointer)`

SetAsset sets Asset field to given value.


### GetDealId

`func (o *AeAssetDealOfferAcceptedDetail) GetDealId() string`

GetDealId returns the DealId field if non-nil, zero value otherwise.

### GetDealIdOk

`func (o *AeAssetDealOfferAcceptedDetail) GetDealIdOk() (*string, bool)`

GetDealIdOk returns a tuple with the DealId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealId

`func (o *AeAssetDealOfferAcceptedDetail) SetDealId(v string)`

SetDealId sets DealId field to given value.

### HasDealId

`func (o *AeAssetDealOfferAcceptedDetail) HasDealId() bool`

HasDealId returns a boolean if a field has been set.

### GetSeller

`func (o *AeAssetDealOfferAcceptedDetail) GetSeller() CommonAmsAccountPointer`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *AeAssetDealOfferAcceptedDetail) GetSellerOk() (*CommonAmsAccountPointer, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *AeAssetDealOfferAcceptedDetail) SetSeller(v CommonAmsAccountPointer)`

SetSeller sets Seller field to given value.


### GetBuyer

`func (o *AeAssetDealOfferAcceptedDetail) GetBuyer() CommonAmsAccountPointer`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *AeAssetDealOfferAcceptedDetail) GetBuyerOk() (*CommonAmsAccountPointer, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *AeAssetDealOfferAcceptedDetail) SetBuyer(v CommonAmsAccountPointer)`

SetBuyer sets Buyer field to given value.


### GetSaleDate

`func (o *AeAssetDealOfferAcceptedDetail) GetSaleDate() string`

GetSaleDate returns the SaleDate field if non-nil, zero value otherwise.

### GetSaleDateOk

`func (o *AeAssetDealOfferAcceptedDetail) GetSaleDateOk() (*string, bool)`

GetSaleDateOk returns a tuple with the SaleDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleDate

`func (o *AeAssetDealOfferAcceptedDetail) SetSaleDate(v string)`

SetSaleDate sets SaleDate field to given value.


### GetLane

`func (o *AeAssetDealOfferAcceptedDetail) GetLane() string`

GetLane returns the Lane field if non-nil, zero value otherwise.

### GetLaneOk

`func (o *AeAssetDealOfferAcceptedDetail) GetLaneOk() (*string, bool)`

GetLaneOk returns a tuple with the Lane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLane

`func (o *AeAssetDealOfferAcceptedDetail) SetLane(v string)`

SetLane sets Lane field to given value.


### GetLot

`func (o *AeAssetDealOfferAcceptedDetail) GetLot() string`

GetLot returns the Lot field if non-nil, zero value otherwise.

### GetLotOk

`func (o *AeAssetDealOfferAcceptedDetail) GetLotOk() (*string, bool)`

GetLotOk returns a tuple with the Lot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLot

`func (o *AeAssetDealOfferAcceptedDetail) SetLot(v string)`

SetLot sets Lot field to given value.


### GetAmount

`func (o *AeAssetDealOfferAcceptedDetail) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AeAssetDealOfferAcceptedDetail) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AeAssetDealOfferAcceptedDetail) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *AeAssetDealOfferAcceptedDetail) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AeAssetDealOfferAcceptedDetail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AeAssetDealOfferAcceptedDetail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AeAssetDealOfferAcceptedDetail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AeAssetDealOfferAcceptedDetail) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetInitiator

`func (o *AeAssetDealOfferAcceptedDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAssetDealOfferAcceptedDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAssetDealOfferAcceptedDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAssetDealOfferAcceptedDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


