# AeServiceCompletedDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**Vin** | **string** | The Vehicle Identification Number(VIN) of the asset. | 
**Stock** | **string** | The stock number of the asset. | 
**Service** | [**AeServiceCompletedDetailService**](AeServiceCompletedDetailService.md) |  | 
**Passed** | Pointer to **bool** | Result of service (pass/fail) | [optional] 
**Messages** | Pointer to **[]string** |  | [optional] 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeServiceCompletedDetail

`func NewAeServiceCompletedDetail(auctionId string, vin string, stock string, service AeServiceCompletedDetailService, ) *AeServiceCompletedDetail`

NewAeServiceCompletedDetail instantiates a new AeServiceCompletedDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeServiceCompletedDetailWithDefaults

`func NewAeServiceCompletedDetailWithDefaults() *AeServiceCompletedDetail`

NewAeServiceCompletedDetailWithDefaults instantiates a new AeServiceCompletedDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeServiceCompletedDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeServiceCompletedDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeServiceCompletedDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetVin

`func (o *AeServiceCompletedDetail) GetVin() string`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *AeServiceCompletedDetail) GetVinOk() (*string, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *AeServiceCompletedDetail) SetVin(v string)`

SetVin sets Vin field to given value.


### GetStock

`func (o *AeServiceCompletedDetail) GetStock() string`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *AeServiceCompletedDetail) GetStockOk() (*string, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *AeServiceCompletedDetail) SetStock(v string)`

SetStock sets Stock field to given value.


### GetService

`func (o *AeServiceCompletedDetail) GetService() AeServiceCompletedDetailService`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *AeServiceCompletedDetail) GetServiceOk() (*AeServiceCompletedDetailService, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *AeServiceCompletedDetail) SetService(v AeServiceCompletedDetailService)`

SetService sets Service field to given value.


### GetPassed

`func (o *AeServiceCompletedDetail) GetPassed() bool`

GetPassed returns the Passed field if non-nil, zero value otherwise.

### GetPassedOk

`func (o *AeServiceCompletedDetail) GetPassedOk() (*bool, bool)`

GetPassedOk returns a tuple with the Passed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassed

`func (o *AeServiceCompletedDetail) SetPassed(v bool)`

SetPassed sets Passed field to given value.

### HasPassed

`func (o *AeServiceCompletedDetail) HasPassed() bool`

HasPassed returns a boolean if a field has been set.

### GetMessages

`func (o *AeServiceCompletedDetail) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *AeServiceCompletedDetail) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *AeServiceCompletedDetail) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *AeServiceCompletedDetail) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetInitiator

`func (o *AeServiceCompletedDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeServiceCompletedDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeServiceCompletedDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeServiceCompletedDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


