# AeAccountPaymentSourceUpsertAmsDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**Account** | [**CommonAmsAccountPointer**](CommonAmsAccountPointer.md) |  | 
**DisplayName** | **string** | display name of the account&#39;s payment source | 
**PaymentSourceId** | **string** | auction unique id for payment source in AMS | 
**PaymentVendorId** | **string** | auction unique id for payment vendor in AMS | 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeAccountPaymentSourceUpsertAmsDetail

`func NewAeAccountPaymentSourceUpsertAmsDetail(auctionId string, account CommonAmsAccountPointer, displayName string, paymentSourceId string, paymentVendorId string, ) *AeAccountPaymentSourceUpsertAmsDetail`

NewAeAccountPaymentSourceUpsertAmsDetail instantiates a new AeAccountPaymentSourceUpsertAmsDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAccountPaymentSourceUpsertAmsDetailWithDefaults

`func NewAeAccountPaymentSourceUpsertAmsDetailWithDefaults() *AeAccountPaymentSourceUpsertAmsDetail`

NewAeAccountPaymentSourceUpsertAmsDetailWithDefaults instantiates a new AeAccountPaymentSourceUpsertAmsDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeAccountPaymentSourceUpsertAmsDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAccountPaymentSourceUpsertAmsDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAccountPaymentSourceUpsertAmsDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetAccount

`func (o *AeAccountPaymentSourceUpsertAmsDetail) GetAccount() CommonAmsAccountPointer`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AeAccountPaymentSourceUpsertAmsDetail) GetAccountOk() (*CommonAmsAccountPointer, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AeAccountPaymentSourceUpsertAmsDetail) SetAccount(v CommonAmsAccountPointer)`

SetAccount sets Account field to given value.


### GetDisplayName

`func (o *AeAccountPaymentSourceUpsertAmsDetail) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AeAccountPaymentSourceUpsertAmsDetail) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AeAccountPaymentSourceUpsertAmsDetail) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetPaymentSourceId

`func (o *AeAccountPaymentSourceUpsertAmsDetail) GetPaymentSourceId() string`

GetPaymentSourceId returns the PaymentSourceId field if non-nil, zero value otherwise.

### GetPaymentSourceIdOk

`func (o *AeAccountPaymentSourceUpsertAmsDetail) GetPaymentSourceIdOk() (*string, bool)`

GetPaymentSourceIdOk returns a tuple with the PaymentSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSourceId

`func (o *AeAccountPaymentSourceUpsertAmsDetail) SetPaymentSourceId(v string)`

SetPaymentSourceId sets PaymentSourceId field to given value.


### GetPaymentVendorId

`func (o *AeAccountPaymentSourceUpsertAmsDetail) GetPaymentVendorId() string`

GetPaymentVendorId returns the PaymentVendorId field if non-nil, zero value otherwise.

### GetPaymentVendorIdOk

`func (o *AeAccountPaymentSourceUpsertAmsDetail) GetPaymentVendorIdOk() (*string, bool)`

GetPaymentVendorIdOk returns a tuple with the PaymentVendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentVendorId

`func (o *AeAccountPaymentSourceUpsertAmsDetail) SetPaymentVendorId(v string)`

SetPaymentVendorId sets PaymentVendorId field to given value.


### GetInitiator

`func (o *AeAccountPaymentSourceUpsertAmsDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAccountPaymentSourceUpsertAmsDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAccountPaymentSourceUpsertAmsDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAccountPaymentSourceUpsertAmsDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


