# AeAssetCheckedInAmsDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**Id** | **string** | Source&#39;s unique identifier for asset | 
**Vin** | **string** | The Vehicle Identification Number(VIN) of the asset. | 
**Stock** | **string** | The stock number of the asset. | 
**CheckIn** | [**CommonAssetCheckIn**](CommonAssetCheckIn.md) |  | 
**Seller** | [**CommonAmsAccountPointer**](CommonAmsAccountPointer.md) |  | 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeAssetCheckedInAmsDetail

`func NewAeAssetCheckedInAmsDetail(auctionId string, id string, vin string, stock string, checkIn CommonAssetCheckIn, seller CommonAmsAccountPointer, ) *AeAssetCheckedInAmsDetail`

NewAeAssetCheckedInAmsDetail instantiates a new AeAssetCheckedInAmsDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetCheckedInAmsDetailWithDefaults

`func NewAeAssetCheckedInAmsDetailWithDefaults() *AeAssetCheckedInAmsDetail`

NewAeAssetCheckedInAmsDetailWithDefaults instantiates a new AeAssetCheckedInAmsDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeAssetCheckedInAmsDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAssetCheckedInAmsDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAssetCheckedInAmsDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetId

`func (o *AeAssetCheckedInAmsDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AeAssetCheckedInAmsDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AeAssetCheckedInAmsDetail) SetId(v string)`

SetId sets Id field to given value.


### GetVin

`func (o *AeAssetCheckedInAmsDetail) GetVin() string`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *AeAssetCheckedInAmsDetail) GetVinOk() (*string, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *AeAssetCheckedInAmsDetail) SetVin(v string)`

SetVin sets Vin field to given value.


### GetStock

`func (o *AeAssetCheckedInAmsDetail) GetStock() string`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *AeAssetCheckedInAmsDetail) GetStockOk() (*string, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *AeAssetCheckedInAmsDetail) SetStock(v string)`

SetStock sets Stock field to given value.


### GetCheckIn

`func (o *AeAssetCheckedInAmsDetail) GetCheckIn() CommonAssetCheckIn`

GetCheckIn returns the CheckIn field if non-nil, zero value otherwise.

### GetCheckInOk

`func (o *AeAssetCheckedInAmsDetail) GetCheckInOk() (*CommonAssetCheckIn, bool)`

GetCheckInOk returns a tuple with the CheckIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckIn

`func (o *AeAssetCheckedInAmsDetail) SetCheckIn(v CommonAssetCheckIn)`

SetCheckIn sets CheckIn field to given value.


### GetSeller

`func (o *AeAssetCheckedInAmsDetail) GetSeller() CommonAmsAccountPointer`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *AeAssetCheckedInAmsDetail) GetSellerOk() (*CommonAmsAccountPointer, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *AeAssetCheckedInAmsDetail) SetSeller(v CommonAmsAccountPointer)`

SetSeller sets Seller field to given value.


### GetInitiator

`func (o *AeAssetCheckedInAmsDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAssetCheckedInAmsDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAssetCheckedInAmsDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAssetCheckedInAmsDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


