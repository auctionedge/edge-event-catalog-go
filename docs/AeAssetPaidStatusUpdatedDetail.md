# AeAssetPaidStatusUpdatedDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**Vin** | **string** | The Vehicle Identification Number(VIN) of the asset. | 
**Stock** | **string** | The stock number of the asset. | 
**IsPaid** | **bool** | True if vehicle is paid for, false if not (reversed) | 
**UpdatedAt** | **time.Time** | ISO-8601 formatted date time | 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeAssetPaidStatusUpdatedDetail

`func NewAeAssetPaidStatusUpdatedDetail(auctionId string, vin string, stock string, isPaid bool, updatedAt time.Time, ) *AeAssetPaidStatusUpdatedDetail`

NewAeAssetPaidStatusUpdatedDetail instantiates a new AeAssetPaidStatusUpdatedDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetPaidStatusUpdatedDetailWithDefaults

`func NewAeAssetPaidStatusUpdatedDetailWithDefaults() *AeAssetPaidStatusUpdatedDetail`

NewAeAssetPaidStatusUpdatedDetailWithDefaults instantiates a new AeAssetPaidStatusUpdatedDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeAssetPaidStatusUpdatedDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAssetPaidStatusUpdatedDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAssetPaidStatusUpdatedDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetVin

`func (o *AeAssetPaidStatusUpdatedDetail) GetVin() string`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *AeAssetPaidStatusUpdatedDetail) GetVinOk() (*string, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *AeAssetPaidStatusUpdatedDetail) SetVin(v string)`

SetVin sets Vin field to given value.


### GetStock

`func (o *AeAssetPaidStatusUpdatedDetail) GetStock() string`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *AeAssetPaidStatusUpdatedDetail) GetStockOk() (*string, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *AeAssetPaidStatusUpdatedDetail) SetStock(v string)`

SetStock sets Stock field to given value.


### GetIsPaid

`func (o *AeAssetPaidStatusUpdatedDetail) GetIsPaid() bool`

GetIsPaid returns the IsPaid field if non-nil, zero value otherwise.

### GetIsPaidOk

`func (o *AeAssetPaidStatusUpdatedDetail) GetIsPaidOk() (*bool, bool)`

GetIsPaidOk returns a tuple with the IsPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaid

`func (o *AeAssetPaidStatusUpdatedDetail) SetIsPaid(v bool)`

SetIsPaid sets IsPaid field to given value.


### GetUpdatedAt

`func (o *AeAssetPaidStatusUpdatedDetail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AeAssetPaidStatusUpdatedDetail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AeAssetPaidStatusUpdatedDetail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetInitiator

`func (o *AeAssetPaidStatusUpdatedDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAssetPaidStatusUpdatedDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAssetPaidStatusUpdatedDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAssetPaidStatusUpdatedDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


