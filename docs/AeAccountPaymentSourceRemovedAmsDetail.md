# AeAccountPaymentSourceRemovedAmsDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**Account** | [**CommonAmsAccountPointer**](CommonAmsAccountPointer.md) |  | 
**PaymentSourceId** | **string** | auction unique id for payment source in AMS | 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeAccountPaymentSourceRemovedAmsDetail

`func NewAeAccountPaymentSourceRemovedAmsDetail(auctionId string, account CommonAmsAccountPointer, paymentSourceId string, ) *AeAccountPaymentSourceRemovedAmsDetail`

NewAeAccountPaymentSourceRemovedAmsDetail instantiates a new AeAccountPaymentSourceRemovedAmsDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAccountPaymentSourceRemovedAmsDetailWithDefaults

`func NewAeAccountPaymentSourceRemovedAmsDetailWithDefaults() *AeAccountPaymentSourceRemovedAmsDetail`

NewAeAccountPaymentSourceRemovedAmsDetailWithDefaults instantiates a new AeAccountPaymentSourceRemovedAmsDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeAccountPaymentSourceRemovedAmsDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAccountPaymentSourceRemovedAmsDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAccountPaymentSourceRemovedAmsDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetAccount

`func (o *AeAccountPaymentSourceRemovedAmsDetail) GetAccount() CommonAmsAccountPointer`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AeAccountPaymentSourceRemovedAmsDetail) GetAccountOk() (*CommonAmsAccountPointer, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AeAccountPaymentSourceRemovedAmsDetail) SetAccount(v CommonAmsAccountPointer)`

SetAccount sets Account field to given value.


### GetPaymentSourceId

`func (o *AeAccountPaymentSourceRemovedAmsDetail) GetPaymentSourceId() string`

GetPaymentSourceId returns the PaymentSourceId field if non-nil, zero value otherwise.

### GetPaymentSourceIdOk

`func (o *AeAccountPaymentSourceRemovedAmsDetail) GetPaymentSourceIdOk() (*string, bool)`

GetPaymentSourceIdOk returns a tuple with the PaymentSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSourceId

`func (o *AeAccountPaymentSourceRemovedAmsDetail) SetPaymentSourceId(v string)`

SetPaymentSourceId sets PaymentSourceId field to given value.


### GetInitiator

`func (o *AeAccountPaymentSourceRemovedAmsDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAccountPaymentSourceRemovedAmsDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAccountPaymentSourceRemovedAmsDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAccountPaymentSourceRemovedAmsDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


