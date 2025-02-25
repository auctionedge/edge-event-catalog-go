# AeAccountRemovedAmsDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Source&#39;s unique identifier for account | 
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**Account** | Pointer to [**CommonAmsAccountPointer**](CommonAmsAccountPointer.md) |  | [optional] 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeAccountRemovedAmsDetail

`func NewAeAccountRemovedAmsDetail(id string, auctionId string, ) *AeAccountRemovedAmsDetail`

NewAeAccountRemovedAmsDetail instantiates a new AeAccountRemovedAmsDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAccountRemovedAmsDetailWithDefaults

`func NewAeAccountRemovedAmsDetailWithDefaults() *AeAccountRemovedAmsDetail`

NewAeAccountRemovedAmsDetailWithDefaults instantiates a new AeAccountRemovedAmsDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AeAccountRemovedAmsDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AeAccountRemovedAmsDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AeAccountRemovedAmsDetail) SetId(v string)`

SetId sets Id field to given value.


### GetAuctionId

`func (o *AeAccountRemovedAmsDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAccountRemovedAmsDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAccountRemovedAmsDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetAccount

`func (o *AeAccountRemovedAmsDetail) GetAccount() CommonAmsAccountPointer`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AeAccountRemovedAmsDetail) GetAccountOk() (*CommonAmsAccountPointer, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AeAccountRemovedAmsDetail) SetAccount(v CommonAmsAccountPointer)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AeAccountRemovedAmsDetail) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetInitiator

`func (o *AeAccountRemovedAmsDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAccountRemovedAmsDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAccountRemovedAmsDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAccountRemovedAmsDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


