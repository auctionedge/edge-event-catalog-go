# AeServiceRequestDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**Vin** | **string** | The Vehicle Identification Number(VIN) of the asset. | 
**Stock** | **string** | The stock number of the asset. | 
**Service** | [**ServiceOrder**](ServiceOrder.md) |  | 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeServiceRequestDetail

`func NewAeServiceRequestDetail(auctionId string, vin string, stock string, service ServiceOrder, ) *AeServiceRequestDetail`

NewAeServiceRequestDetail instantiates a new AeServiceRequestDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeServiceRequestDetailWithDefaults

`func NewAeServiceRequestDetailWithDefaults() *AeServiceRequestDetail`

NewAeServiceRequestDetailWithDefaults instantiates a new AeServiceRequestDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeServiceRequestDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeServiceRequestDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeServiceRequestDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetVin

`func (o *AeServiceRequestDetail) GetVin() string`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *AeServiceRequestDetail) GetVinOk() (*string, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *AeServiceRequestDetail) SetVin(v string)`

SetVin sets Vin field to given value.


### GetStock

`func (o *AeServiceRequestDetail) GetStock() string`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *AeServiceRequestDetail) GetStockOk() (*string, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *AeServiceRequestDetail) SetStock(v string)`

SetStock sets Stock field to given value.


### GetService

`func (o *AeServiceRequestDetail) GetService() ServiceOrder`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *AeServiceRequestDetail) GetServiceOk() (*ServiceOrder, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *AeServiceRequestDetail) SetService(v ServiceOrder)`

SetService sets Service field to given value.


### GetInitiator

`func (o *AeServiceRequestDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeServiceRequestDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeServiceRequestDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeServiceRequestDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


