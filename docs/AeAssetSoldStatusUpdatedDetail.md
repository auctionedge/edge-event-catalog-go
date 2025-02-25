# AeAssetSoldStatusUpdatedDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**Vin** | **string** | The Vehicle Identification Number(VIN) of the asset. | 
**Stock** | **string** | The stock number of the asset. | 
**Amount** | **NullableFloat32** | Amount asset was sold for, will be null if reversed | 
**UpdatedAt** | **time.Time** | ISO-8601 formatted date time | 
**IsSold** | **bool** | True if vehicle is sold, false if not (reversed) | 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeAssetSoldStatusUpdatedDetail

`func NewAeAssetSoldStatusUpdatedDetail(auctionId string, vin string, stock string, amount NullableFloat32, updatedAt time.Time, isSold bool, ) *AeAssetSoldStatusUpdatedDetail`

NewAeAssetSoldStatusUpdatedDetail instantiates a new AeAssetSoldStatusUpdatedDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetSoldStatusUpdatedDetailWithDefaults

`func NewAeAssetSoldStatusUpdatedDetailWithDefaults() *AeAssetSoldStatusUpdatedDetail`

NewAeAssetSoldStatusUpdatedDetailWithDefaults instantiates a new AeAssetSoldStatusUpdatedDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeAssetSoldStatusUpdatedDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAssetSoldStatusUpdatedDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAssetSoldStatusUpdatedDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetVin

`func (o *AeAssetSoldStatusUpdatedDetail) GetVin() string`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *AeAssetSoldStatusUpdatedDetail) GetVinOk() (*string, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *AeAssetSoldStatusUpdatedDetail) SetVin(v string)`

SetVin sets Vin field to given value.


### GetStock

`func (o *AeAssetSoldStatusUpdatedDetail) GetStock() string`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *AeAssetSoldStatusUpdatedDetail) GetStockOk() (*string, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *AeAssetSoldStatusUpdatedDetail) SetStock(v string)`

SetStock sets Stock field to given value.


### GetAmount

`func (o *AeAssetSoldStatusUpdatedDetail) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AeAssetSoldStatusUpdatedDetail) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AeAssetSoldStatusUpdatedDetail) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### SetAmountNil

`func (o *AeAssetSoldStatusUpdatedDetail) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *AeAssetSoldStatusUpdatedDetail) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetUpdatedAt

`func (o *AeAssetSoldStatusUpdatedDetail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AeAssetSoldStatusUpdatedDetail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AeAssetSoldStatusUpdatedDetail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetIsSold

`func (o *AeAssetSoldStatusUpdatedDetail) GetIsSold() bool`

GetIsSold returns the IsSold field if non-nil, zero value otherwise.

### GetIsSoldOk

`func (o *AeAssetSoldStatusUpdatedDetail) GetIsSoldOk() (*bool, bool)`

GetIsSoldOk returns a tuple with the IsSold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSold

`func (o *AeAssetSoldStatusUpdatedDetail) SetIsSold(v bool)`

SetIsSold sets IsSold field to given value.


### GetInitiator

`func (o *AeAssetSoldStatusUpdatedDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAssetSoldStatusUpdatedDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAssetSoldStatusUpdatedDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAssetSoldStatusUpdatedDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


