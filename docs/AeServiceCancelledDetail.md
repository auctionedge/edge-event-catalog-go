# AeServiceCancelledDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**Vin** | **string** | The Vehicle Identification Number(VIN) of the asset. | 
**Stock** | **string** | The stock number of the asset. | 
**Service** | [**AeServiceOrderPlacedDetailService**](AeServiceOrderPlacedDetailService.md) |  | 
**Reason** | Pointer to **string** | Reason that the service was cancelled | [optional] 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeServiceCancelledDetail

`func NewAeServiceCancelledDetail(auctionId string, vin string, stock string, service AeServiceOrderPlacedDetailService, ) *AeServiceCancelledDetail`

NewAeServiceCancelledDetail instantiates a new AeServiceCancelledDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeServiceCancelledDetailWithDefaults

`func NewAeServiceCancelledDetailWithDefaults() *AeServiceCancelledDetail`

NewAeServiceCancelledDetailWithDefaults instantiates a new AeServiceCancelledDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeServiceCancelledDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeServiceCancelledDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeServiceCancelledDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetVin

`func (o *AeServiceCancelledDetail) GetVin() string`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *AeServiceCancelledDetail) GetVinOk() (*string, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *AeServiceCancelledDetail) SetVin(v string)`

SetVin sets Vin field to given value.


### GetStock

`func (o *AeServiceCancelledDetail) GetStock() string`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *AeServiceCancelledDetail) GetStockOk() (*string, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *AeServiceCancelledDetail) SetStock(v string)`

SetStock sets Stock field to given value.


### GetService

`func (o *AeServiceCancelledDetail) GetService() AeServiceOrderPlacedDetailService`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *AeServiceCancelledDetail) GetServiceOk() (*AeServiceOrderPlacedDetailService, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *AeServiceCancelledDetail) SetService(v AeServiceOrderPlacedDetailService)`

SetService sets Service field to given value.


### GetReason

`func (o *AeServiceCancelledDetail) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AeServiceCancelledDetail) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AeServiceCancelledDetail) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AeServiceCancelledDetail) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetInitiator

`func (o *AeServiceCancelledDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeServiceCancelledDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeServiceCancelledDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeServiceCancelledDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


