# AeAssetDealUnsoldDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**Asset** | Pointer to [**CommonAmsAssetPointer**](CommonAmsAssetPointer.md) |  | [optional] 
**DealId** | Pointer to **string** | Unique id for a deal | [optional] 
**Vin** | **string** | The Vehicle Identification Number(VIN) of the asset. | 
**Stock** | **string** | The stock number of the asset. | 
**SaleDate** | **string** | The sale date of the deal. | 
**Lane** | **string** | The lane the asset was in when deal made. | 
**Lot** | **string** | The lot the asset was in when deal made. | 
**UpdatedAt** | Pointer to **time.Time** | The updated date time of the gatepass status | [optional] 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeAssetDealUnsoldDetail

`func NewAeAssetDealUnsoldDetail(auctionId string, vin string, stock string, saleDate string, lane string, lot string, ) *AeAssetDealUnsoldDetail`

NewAeAssetDealUnsoldDetail instantiates a new AeAssetDealUnsoldDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetDealUnsoldDetailWithDefaults

`func NewAeAssetDealUnsoldDetailWithDefaults() *AeAssetDealUnsoldDetail`

NewAeAssetDealUnsoldDetailWithDefaults instantiates a new AeAssetDealUnsoldDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeAssetDealUnsoldDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAssetDealUnsoldDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAssetDealUnsoldDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetAsset

`func (o *AeAssetDealUnsoldDetail) GetAsset() CommonAmsAssetPointer`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *AeAssetDealUnsoldDetail) GetAssetOk() (*CommonAmsAssetPointer, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *AeAssetDealUnsoldDetail) SetAsset(v CommonAmsAssetPointer)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *AeAssetDealUnsoldDetail) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetDealId

`func (o *AeAssetDealUnsoldDetail) GetDealId() string`

GetDealId returns the DealId field if non-nil, zero value otherwise.

### GetDealIdOk

`func (o *AeAssetDealUnsoldDetail) GetDealIdOk() (*string, bool)`

GetDealIdOk returns a tuple with the DealId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealId

`func (o *AeAssetDealUnsoldDetail) SetDealId(v string)`

SetDealId sets DealId field to given value.

### HasDealId

`func (o *AeAssetDealUnsoldDetail) HasDealId() bool`

HasDealId returns a boolean if a field has been set.

### GetVin

`func (o *AeAssetDealUnsoldDetail) GetVin() string`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *AeAssetDealUnsoldDetail) GetVinOk() (*string, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *AeAssetDealUnsoldDetail) SetVin(v string)`

SetVin sets Vin field to given value.


### GetStock

`func (o *AeAssetDealUnsoldDetail) GetStock() string`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *AeAssetDealUnsoldDetail) GetStockOk() (*string, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *AeAssetDealUnsoldDetail) SetStock(v string)`

SetStock sets Stock field to given value.


### GetSaleDate

`func (o *AeAssetDealUnsoldDetail) GetSaleDate() string`

GetSaleDate returns the SaleDate field if non-nil, zero value otherwise.

### GetSaleDateOk

`func (o *AeAssetDealUnsoldDetail) GetSaleDateOk() (*string, bool)`

GetSaleDateOk returns a tuple with the SaleDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleDate

`func (o *AeAssetDealUnsoldDetail) SetSaleDate(v string)`

SetSaleDate sets SaleDate field to given value.


### GetLane

`func (o *AeAssetDealUnsoldDetail) GetLane() string`

GetLane returns the Lane field if non-nil, zero value otherwise.

### GetLaneOk

`func (o *AeAssetDealUnsoldDetail) GetLaneOk() (*string, bool)`

GetLaneOk returns a tuple with the Lane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLane

`func (o *AeAssetDealUnsoldDetail) SetLane(v string)`

SetLane sets Lane field to given value.


### GetLot

`func (o *AeAssetDealUnsoldDetail) GetLot() string`

GetLot returns the Lot field if non-nil, zero value otherwise.

### GetLotOk

`func (o *AeAssetDealUnsoldDetail) GetLotOk() (*string, bool)`

GetLotOk returns a tuple with the Lot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLot

`func (o *AeAssetDealUnsoldDetail) SetLot(v string)`

SetLot sets Lot field to given value.


### GetUpdatedAt

`func (o *AeAssetDealUnsoldDetail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AeAssetDealUnsoldDetail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AeAssetDealUnsoldDetail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AeAssetDealUnsoldDetail) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetInitiator

`func (o *AeAssetDealUnsoldDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAssetDealUnsoldDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAssetDealUnsoldDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAssetDealUnsoldDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


