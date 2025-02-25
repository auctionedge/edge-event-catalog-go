# AeAssetDealSoldDetail

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
**Year** | Pointer to **int32** | The year the vehicle was manufactured. | [optional] 
**Make** | Pointer to **string** | The manufacturer of the vehicle. | [optional] 
**Model** | Pointer to **string** | The model of the vehicle. | [optional] 
**Trimline** | Pointer to **string** | The trim of the vehicle. | [optional] 
**Mileage** | Pointer to **int32** | The mileage of the asset being sold. | [optional] 
**FuelType** | Pointer to **string** | The type of fuel required by the asset being sold. | [optional] 
**SaleType** | Pointer to **string** | The type of sale the deal was made in. | [optional] 
**Lights** | Pointer to **string** | The lights on the asset when the deal was made. | [optional] 
**Announcements** | Pointer to **string** | The announcments on the asset when deal was made. | [optional] 
**SaleDate** | **string** | The sale date of the deal. | 
**Lane** | **string** | The lane the asset was in when deal made. | 
**Lot** | **string** | The lot the asset was in when deal made. | 
**Amount** | **float32** | The amount the asset sold for. | 
**UpdatedAt** | Pointer to **time.Time** | The updated date time of the deal. | [optional] 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeAssetDealSoldDetail

`func NewAeAssetDealSoldDetail(auctionId string, vin string, stock string, seller CommonAmsAccountPointer, buyer CommonAmsAccountPointer, saleDate string, lane string, lot string, amount float32, ) *AeAssetDealSoldDetail`

NewAeAssetDealSoldDetail instantiates a new AeAssetDealSoldDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetDealSoldDetailWithDefaults

`func NewAeAssetDealSoldDetailWithDefaults() *AeAssetDealSoldDetail`

NewAeAssetDealSoldDetailWithDefaults instantiates a new AeAssetDealSoldDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeAssetDealSoldDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAssetDealSoldDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAssetDealSoldDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetAsset

`func (o *AeAssetDealSoldDetail) GetAsset() CommonAmsAssetPointer`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *AeAssetDealSoldDetail) GetAssetOk() (*CommonAmsAssetPointer, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *AeAssetDealSoldDetail) SetAsset(v CommonAmsAssetPointer)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *AeAssetDealSoldDetail) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetDealId

`func (o *AeAssetDealSoldDetail) GetDealId() string`

GetDealId returns the DealId field if non-nil, zero value otherwise.

### GetDealIdOk

`func (o *AeAssetDealSoldDetail) GetDealIdOk() (*string, bool)`

GetDealIdOk returns a tuple with the DealId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealId

`func (o *AeAssetDealSoldDetail) SetDealId(v string)`

SetDealId sets DealId field to given value.

### HasDealId

`func (o *AeAssetDealSoldDetail) HasDealId() bool`

HasDealId returns a boolean if a field has been set.

### GetVin

`func (o *AeAssetDealSoldDetail) GetVin() string`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *AeAssetDealSoldDetail) GetVinOk() (*string, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *AeAssetDealSoldDetail) SetVin(v string)`

SetVin sets Vin field to given value.


### GetStock

`func (o *AeAssetDealSoldDetail) GetStock() string`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *AeAssetDealSoldDetail) GetStockOk() (*string, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *AeAssetDealSoldDetail) SetStock(v string)`

SetStock sets Stock field to given value.


### GetSeller

`func (o *AeAssetDealSoldDetail) GetSeller() CommonAmsAccountPointer`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *AeAssetDealSoldDetail) GetSellerOk() (*CommonAmsAccountPointer, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *AeAssetDealSoldDetail) SetSeller(v CommonAmsAccountPointer)`

SetSeller sets Seller field to given value.


### GetBuyer

`func (o *AeAssetDealSoldDetail) GetBuyer() CommonAmsAccountPointer`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *AeAssetDealSoldDetail) GetBuyerOk() (*CommonAmsAccountPointer, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *AeAssetDealSoldDetail) SetBuyer(v CommonAmsAccountPointer)`

SetBuyer sets Buyer field to given value.


### GetYear

`func (o *AeAssetDealSoldDetail) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *AeAssetDealSoldDetail) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *AeAssetDealSoldDetail) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *AeAssetDealSoldDetail) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetMake

`func (o *AeAssetDealSoldDetail) GetMake() string`

GetMake returns the Make field if non-nil, zero value otherwise.

### GetMakeOk

`func (o *AeAssetDealSoldDetail) GetMakeOk() (*string, bool)`

GetMakeOk returns a tuple with the Make field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMake

`func (o *AeAssetDealSoldDetail) SetMake(v string)`

SetMake sets Make field to given value.

### HasMake

`func (o *AeAssetDealSoldDetail) HasMake() bool`

HasMake returns a boolean if a field has been set.

### GetModel

`func (o *AeAssetDealSoldDetail) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AeAssetDealSoldDetail) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AeAssetDealSoldDetail) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AeAssetDealSoldDetail) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetTrimline

`func (o *AeAssetDealSoldDetail) GetTrimline() string`

GetTrimline returns the Trimline field if non-nil, zero value otherwise.

### GetTrimlineOk

`func (o *AeAssetDealSoldDetail) GetTrimlineOk() (*string, bool)`

GetTrimlineOk returns a tuple with the Trimline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrimline

`func (o *AeAssetDealSoldDetail) SetTrimline(v string)`

SetTrimline sets Trimline field to given value.

### HasTrimline

`func (o *AeAssetDealSoldDetail) HasTrimline() bool`

HasTrimline returns a boolean if a field has been set.

### GetMileage

`func (o *AeAssetDealSoldDetail) GetMileage() int32`

GetMileage returns the Mileage field if non-nil, zero value otherwise.

### GetMileageOk

`func (o *AeAssetDealSoldDetail) GetMileageOk() (*int32, bool)`

GetMileageOk returns a tuple with the Mileage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMileage

`func (o *AeAssetDealSoldDetail) SetMileage(v int32)`

SetMileage sets Mileage field to given value.

### HasMileage

`func (o *AeAssetDealSoldDetail) HasMileage() bool`

HasMileage returns a boolean if a field has been set.

### GetFuelType

`func (o *AeAssetDealSoldDetail) GetFuelType() string`

GetFuelType returns the FuelType field if non-nil, zero value otherwise.

### GetFuelTypeOk

`func (o *AeAssetDealSoldDetail) GetFuelTypeOk() (*string, bool)`

GetFuelTypeOk returns a tuple with the FuelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelType

`func (o *AeAssetDealSoldDetail) SetFuelType(v string)`

SetFuelType sets FuelType field to given value.

### HasFuelType

`func (o *AeAssetDealSoldDetail) HasFuelType() bool`

HasFuelType returns a boolean if a field has been set.

### GetSaleType

`func (o *AeAssetDealSoldDetail) GetSaleType() string`

GetSaleType returns the SaleType field if non-nil, zero value otherwise.

### GetSaleTypeOk

`func (o *AeAssetDealSoldDetail) GetSaleTypeOk() (*string, bool)`

GetSaleTypeOk returns a tuple with the SaleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleType

`func (o *AeAssetDealSoldDetail) SetSaleType(v string)`

SetSaleType sets SaleType field to given value.

### HasSaleType

`func (o *AeAssetDealSoldDetail) HasSaleType() bool`

HasSaleType returns a boolean if a field has been set.

### GetLights

`func (o *AeAssetDealSoldDetail) GetLights() string`

GetLights returns the Lights field if non-nil, zero value otherwise.

### GetLightsOk

`func (o *AeAssetDealSoldDetail) GetLightsOk() (*string, bool)`

GetLightsOk returns a tuple with the Lights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLights

`func (o *AeAssetDealSoldDetail) SetLights(v string)`

SetLights sets Lights field to given value.

### HasLights

`func (o *AeAssetDealSoldDetail) HasLights() bool`

HasLights returns a boolean if a field has been set.

### GetAnnouncements

`func (o *AeAssetDealSoldDetail) GetAnnouncements() string`

GetAnnouncements returns the Announcements field if non-nil, zero value otherwise.

### GetAnnouncementsOk

`func (o *AeAssetDealSoldDetail) GetAnnouncementsOk() (*string, bool)`

GetAnnouncementsOk returns a tuple with the Announcements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncements

`func (o *AeAssetDealSoldDetail) SetAnnouncements(v string)`

SetAnnouncements sets Announcements field to given value.

### HasAnnouncements

`func (o *AeAssetDealSoldDetail) HasAnnouncements() bool`

HasAnnouncements returns a boolean if a field has been set.

### GetSaleDate

`func (o *AeAssetDealSoldDetail) GetSaleDate() string`

GetSaleDate returns the SaleDate field if non-nil, zero value otherwise.

### GetSaleDateOk

`func (o *AeAssetDealSoldDetail) GetSaleDateOk() (*string, bool)`

GetSaleDateOk returns a tuple with the SaleDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleDate

`func (o *AeAssetDealSoldDetail) SetSaleDate(v string)`

SetSaleDate sets SaleDate field to given value.


### GetLane

`func (o *AeAssetDealSoldDetail) GetLane() string`

GetLane returns the Lane field if non-nil, zero value otherwise.

### GetLaneOk

`func (o *AeAssetDealSoldDetail) GetLaneOk() (*string, bool)`

GetLaneOk returns a tuple with the Lane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLane

`func (o *AeAssetDealSoldDetail) SetLane(v string)`

SetLane sets Lane field to given value.


### GetLot

`func (o *AeAssetDealSoldDetail) GetLot() string`

GetLot returns the Lot field if non-nil, zero value otherwise.

### GetLotOk

`func (o *AeAssetDealSoldDetail) GetLotOk() (*string, bool)`

GetLotOk returns a tuple with the Lot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLot

`func (o *AeAssetDealSoldDetail) SetLot(v string)`

SetLot sets Lot field to given value.


### GetAmount

`func (o *AeAssetDealSoldDetail) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AeAssetDealSoldDetail) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AeAssetDealSoldDetail) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetUpdatedAt

`func (o *AeAssetDealSoldDetail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AeAssetDealSoldDetail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AeAssetDealSoldDetail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AeAssetDealSoldDetail) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetInitiator

`func (o *AeAssetDealSoldDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAssetDealSoldDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAssetDealSoldDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAssetDealSoldDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


