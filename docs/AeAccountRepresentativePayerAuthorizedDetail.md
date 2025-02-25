# AeAccountRepresentativePayerAuthorizedDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**Account** | [**CommonAmsAccountPointer**](CommonAmsAccountPointer.md) |  | 
**RepAccount** | [**CommonAmsRepPointer**](CommonAmsRepPointer.md) |  | 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeAccountRepresentativePayerAuthorizedDetail

`func NewAeAccountRepresentativePayerAuthorizedDetail(auctionId string, account CommonAmsAccountPointer, repAccount CommonAmsRepPointer, ) *AeAccountRepresentativePayerAuthorizedDetail`

NewAeAccountRepresentativePayerAuthorizedDetail instantiates a new AeAccountRepresentativePayerAuthorizedDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAccountRepresentativePayerAuthorizedDetailWithDefaults

`func NewAeAccountRepresentativePayerAuthorizedDetailWithDefaults() *AeAccountRepresentativePayerAuthorizedDetail`

NewAeAccountRepresentativePayerAuthorizedDetailWithDefaults instantiates a new AeAccountRepresentativePayerAuthorizedDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeAccountRepresentativePayerAuthorizedDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAccountRepresentativePayerAuthorizedDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAccountRepresentativePayerAuthorizedDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetAccount

`func (o *AeAccountRepresentativePayerAuthorizedDetail) GetAccount() CommonAmsAccountPointer`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AeAccountRepresentativePayerAuthorizedDetail) GetAccountOk() (*CommonAmsAccountPointer, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AeAccountRepresentativePayerAuthorizedDetail) SetAccount(v CommonAmsAccountPointer)`

SetAccount sets Account field to given value.


### GetRepAccount

`func (o *AeAccountRepresentativePayerAuthorizedDetail) GetRepAccount() CommonAmsRepPointer`

GetRepAccount returns the RepAccount field if non-nil, zero value otherwise.

### GetRepAccountOk

`func (o *AeAccountRepresentativePayerAuthorizedDetail) GetRepAccountOk() (*CommonAmsRepPointer, bool)`

GetRepAccountOk returns a tuple with the RepAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepAccount

`func (o *AeAccountRepresentativePayerAuthorizedDetail) SetRepAccount(v CommonAmsRepPointer)`

SetRepAccount sets RepAccount field to given value.


### GetInitiator

`func (o *AeAccountRepresentativePayerAuthorizedDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAccountRepresentativePayerAuthorizedDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAccountRepresentativePayerAuthorizedDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAccountRepresentativePayerAuthorizedDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


