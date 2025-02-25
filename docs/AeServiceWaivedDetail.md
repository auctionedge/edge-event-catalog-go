# AeServiceWaivedDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**Vin** | **string** | The Vehicle Identification Number(VIN) of the asset. | 
**Stock** | **string** | The stock number of the asset. | 
**Service** | [**AeServiceWaivedDetailService**](AeServiceWaivedDetailService.md) |  | 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeServiceWaivedDetail

`func NewAeServiceWaivedDetail(auctionId string, vin string, stock string, service AeServiceWaivedDetailService, ) *AeServiceWaivedDetail`

NewAeServiceWaivedDetail instantiates a new AeServiceWaivedDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeServiceWaivedDetailWithDefaults

`func NewAeServiceWaivedDetailWithDefaults() *AeServiceWaivedDetail`

NewAeServiceWaivedDetailWithDefaults instantiates a new AeServiceWaivedDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeServiceWaivedDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeServiceWaivedDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeServiceWaivedDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetVin

`func (o *AeServiceWaivedDetail) GetVin() string`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *AeServiceWaivedDetail) GetVinOk() (*string, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *AeServiceWaivedDetail) SetVin(v string)`

SetVin sets Vin field to given value.


### GetStock

`func (o *AeServiceWaivedDetail) GetStock() string`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *AeServiceWaivedDetail) GetStockOk() (*string, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *AeServiceWaivedDetail) SetStock(v string)`

SetStock sets Stock field to given value.


### GetService

`func (o *AeServiceWaivedDetail) GetService() AeServiceWaivedDetailService`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *AeServiceWaivedDetail) GetServiceOk() (*AeServiceWaivedDetailService, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *AeServiceWaivedDetail) SetService(v AeServiceWaivedDetailService)`

SetService sets Service field to given value.


### GetInitiator

`func (o *AeServiceWaivedDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeServiceWaivedDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeServiceWaivedDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeServiceWaivedDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


