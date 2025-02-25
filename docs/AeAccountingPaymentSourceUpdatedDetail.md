# AeAccountingPaymentSourceUpdatedDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**Vin** | **string** | The Vehicle Identification Number(VIN) of the asset. | 
**Stock** | **string** | The stock number of the asset. | 
**PriorSource** | **NullableString** | Simple text describing the prior funding source if any. | 
**CurrentSource** | **NullableString** | Simple text describing the current funding source if any. | 
**PriorStatus** | Pointer to [**PaymentStatus**](PaymentStatus.md) |  | [optional] 
**CurrentStatus** | Pointer to [**PaymentStatus**](PaymentStatus.md) |  | [optional] 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeAccountingPaymentSourceUpdatedDetail

`func NewAeAccountingPaymentSourceUpdatedDetail(auctionId string, vin string, stock string, priorSource NullableString, currentSource NullableString, ) *AeAccountingPaymentSourceUpdatedDetail`

NewAeAccountingPaymentSourceUpdatedDetail instantiates a new AeAccountingPaymentSourceUpdatedDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAccountingPaymentSourceUpdatedDetailWithDefaults

`func NewAeAccountingPaymentSourceUpdatedDetailWithDefaults() *AeAccountingPaymentSourceUpdatedDetail`

NewAeAccountingPaymentSourceUpdatedDetailWithDefaults instantiates a new AeAccountingPaymentSourceUpdatedDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeAccountingPaymentSourceUpdatedDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAccountingPaymentSourceUpdatedDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAccountingPaymentSourceUpdatedDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetVin

`func (o *AeAccountingPaymentSourceUpdatedDetail) GetVin() string`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *AeAccountingPaymentSourceUpdatedDetail) GetVinOk() (*string, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *AeAccountingPaymentSourceUpdatedDetail) SetVin(v string)`

SetVin sets Vin field to given value.


### GetStock

`func (o *AeAccountingPaymentSourceUpdatedDetail) GetStock() string`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *AeAccountingPaymentSourceUpdatedDetail) GetStockOk() (*string, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *AeAccountingPaymentSourceUpdatedDetail) SetStock(v string)`

SetStock sets Stock field to given value.


### GetPriorSource

`func (o *AeAccountingPaymentSourceUpdatedDetail) GetPriorSource() string`

GetPriorSource returns the PriorSource field if non-nil, zero value otherwise.

### GetPriorSourceOk

`func (o *AeAccountingPaymentSourceUpdatedDetail) GetPriorSourceOk() (*string, bool)`

GetPriorSourceOk returns a tuple with the PriorSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorSource

`func (o *AeAccountingPaymentSourceUpdatedDetail) SetPriorSource(v string)`

SetPriorSource sets PriorSource field to given value.


### SetPriorSourceNil

`func (o *AeAccountingPaymentSourceUpdatedDetail) SetPriorSourceNil(b bool)`

 SetPriorSourceNil sets the value for PriorSource to be an explicit nil

### UnsetPriorSource
`func (o *AeAccountingPaymentSourceUpdatedDetail) UnsetPriorSource()`

UnsetPriorSource ensures that no value is present for PriorSource, not even an explicit nil
### GetCurrentSource

`func (o *AeAccountingPaymentSourceUpdatedDetail) GetCurrentSource() string`

GetCurrentSource returns the CurrentSource field if non-nil, zero value otherwise.

### GetCurrentSourceOk

`func (o *AeAccountingPaymentSourceUpdatedDetail) GetCurrentSourceOk() (*string, bool)`

GetCurrentSourceOk returns a tuple with the CurrentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSource

`func (o *AeAccountingPaymentSourceUpdatedDetail) SetCurrentSource(v string)`

SetCurrentSource sets CurrentSource field to given value.


### SetCurrentSourceNil

`func (o *AeAccountingPaymentSourceUpdatedDetail) SetCurrentSourceNil(b bool)`

 SetCurrentSourceNil sets the value for CurrentSource to be an explicit nil

### UnsetCurrentSource
`func (o *AeAccountingPaymentSourceUpdatedDetail) UnsetCurrentSource()`

UnsetCurrentSource ensures that no value is present for CurrentSource, not even an explicit nil
### GetPriorStatus

`func (o *AeAccountingPaymentSourceUpdatedDetail) GetPriorStatus() PaymentStatus`

GetPriorStatus returns the PriorStatus field if non-nil, zero value otherwise.

### GetPriorStatusOk

`func (o *AeAccountingPaymentSourceUpdatedDetail) GetPriorStatusOk() (*PaymentStatus, bool)`

GetPriorStatusOk returns a tuple with the PriorStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorStatus

`func (o *AeAccountingPaymentSourceUpdatedDetail) SetPriorStatus(v PaymentStatus)`

SetPriorStatus sets PriorStatus field to given value.

### HasPriorStatus

`func (o *AeAccountingPaymentSourceUpdatedDetail) HasPriorStatus() bool`

HasPriorStatus returns a boolean if a field has been set.

### GetCurrentStatus

`func (o *AeAccountingPaymentSourceUpdatedDetail) GetCurrentStatus() PaymentStatus`

GetCurrentStatus returns the CurrentStatus field if non-nil, zero value otherwise.

### GetCurrentStatusOk

`func (o *AeAccountingPaymentSourceUpdatedDetail) GetCurrentStatusOk() (*PaymentStatus, bool)`

GetCurrentStatusOk returns a tuple with the CurrentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStatus

`func (o *AeAccountingPaymentSourceUpdatedDetail) SetCurrentStatus(v PaymentStatus)`

SetCurrentStatus sets CurrentStatus field to given value.

### HasCurrentStatus

`func (o *AeAccountingPaymentSourceUpdatedDetail) HasCurrentStatus() bool`

HasCurrentStatus returns a boolean if a field has been set.

### GetInitiator

`func (o *AeAccountingPaymentSourceUpdatedDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAccountingPaymentSourceUpdatedDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAccountingPaymentSourceUpdatedDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAccountingPaymentSourceUpdatedDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


