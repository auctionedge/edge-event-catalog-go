# AeAssetDealOfferRejectedDetail

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
**UpdatedAt** | Pointer to **time.Time** | The updated date time of the gatepass status | [optional] 
**RejectedReason** | **string** | The reason the offer was rejected | 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeAssetDealOfferRejectedDetail

`func NewAeAssetDealOfferRejectedDetail(auctionId string, asset CommonAmsAssetPointer, seller CommonAmsAccountPointer, buyer CommonAmsAccountPointer, saleDate string, lane string, lot string, rejectedReason string, ) *AeAssetDealOfferRejectedDetail`

NewAeAssetDealOfferRejectedDetail instantiates a new AeAssetDealOfferRejectedDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetDealOfferRejectedDetailWithDefaults

`func NewAeAssetDealOfferRejectedDetailWithDefaults() *AeAssetDealOfferRejectedDetail`

NewAeAssetDealOfferRejectedDetailWithDefaults instantiates a new AeAssetDealOfferRejectedDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeAssetDealOfferRejectedDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAssetDealOfferRejectedDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAssetDealOfferRejectedDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetAsset

`func (o *AeAssetDealOfferRejectedDetail) GetAsset() CommonAmsAssetPointer`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *AeAssetDealOfferRejectedDetail) GetAssetOk() (*CommonAmsAssetPointer, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *AeAssetDealOfferRejectedDetail) SetAsset(v CommonAmsAssetPointer)`

SetAsset sets Asset field to given value.


### GetDealId

`func (o *AeAssetDealOfferRejectedDetail) GetDealId() string`

GetDealId returns the DealId field if non-nil, zero value otherwise.

### GetDealIdOk

`func (o *AeAssetDealOfferRejectedDetail) GetDealIdOk() (*string, bool)`

GetDealIdOk returns a tuple with the DealId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealId

`func (o *AeAssetDealOfferRejectedDetail) SetDealId(v string)`

SetDealId sets DealId field to given value.

### HasDealId

`func (o *AeAssetDealOfferRejectedDetail) HasDealId() bool`

HasDealId returns a boolean if a field has been set.

### GetSeller

`func (o *AeAssetDealOfferRejectedDetail) GetSeller() CommonAmsAccountPointer`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *AeAssetDealOfferRejectedDetail) GetSellerOk() (*CommonAmsAccountPointer, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *AeAssetDealOfferRejectedDetail) SetSeller(v CommonAmsAccountPointer)`

SetSeller sets Seller field to given value.


### GetBuyer

`func (o *AeAssetDealOfferRejectedDetail) GetBuyer() CommonAmsAccountPointer`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *AeAssetDealOfferRejectedDetail) GetBuyerOk() (*CommonAmsAccountPointer, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *AeAssetDealOfferRejectedDetail) SetBuyer(v CommonAmsAccountPointer)`

SetBuyer sets Buyer field to given value.


### GetSaleDate

`func (o *AeAssetDealOfferRejectedDetail) GetSaleDate() string`

GetSaleDate returns the SaleDate field if non-nil, zero value otherwise.

### GetSaleDateOk

`func (o *AeAssetDealOfferRejectedDetail) GetSaleDateOk() (*string, bool)`

GetSaleDateOk returns a tuple with the SaleDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleDate

`func (o *AeAssetDealOfferRejectedDetail) SetSaleDate(v string)`

SetSaleDate sets SaleDate field to given value.


### GetLane

`func (o *AeAssetDealOfferRejectedDetail) GetLane() string`

GetLane returns the Lane field if non-nil, zero value otherwise.

### GetLaneOk

`func (o *AeAssetDealOfferRejectedDetail) GetLaneOk() (*string, bool)`

GetLaneOk returns a tuple with the Lane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLane

`func (o *AeAssetDealOfferRejectedDetail) SetLane(v string)`

SetLane sets Lane field to given value.


### GetLot

`func (o *AeAssetDealOfferRejectedDetail) GetLot() string`

GetLot returns the Lot field if non-nil, zero value otherwise.

### GetLotOk

`func (o *AeAssetDealOfferRejectedDetail) GetLotOk() (*string, bool)`

GetLotOk returns a tuple with the Lot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLot

`func (o *AeAssetDealOfferRejectedDetail) SetLot(v string)`

SetLot sets Lot field to given value.


### GetAmount

`func (o *AeAssetDealOfferRejectedDetail) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AeAssetDealOfferRejectedDetail) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AeAssetDealOfferRejectedDetail) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *AeAssetDealOfferRejectedDetail) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AeAssetDealOfferRejectedDetail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AeAssetDealOfferRejectedDetail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AeAssetDealOfferRejectedDetail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AeAssetDealOfferRejectedDetail) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetRejectedReason

`func (o *AeAssetDealOfferRejectedDetail) GetRejectedReason() string`

GetRejectedReason returns the RejectedReason field if non-nil, zero value otherwise.

### GetRejectedReasonOk

`func (o *AeAssetDealOfferRejectedDetail) GetRejectedReasonOk() (*string, bool)`

GetRejectedReasonOk returns a tuple with the RejectedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedReason

`func (o *AeAssetDealOfferRejectedDetail) SetRejectedReason(v string)`

SetRejectedReason sets RejectedReason field to given value.


### GetInitiator

`func (o *AeAssetDealOfferRejectedDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAssetDealOfferRejectedDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAssetDealOfferRejectedDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAssetDealOfferRejectedDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


