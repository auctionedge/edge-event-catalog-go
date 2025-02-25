# AeAccountRepresentativeRemovedAmsDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The representative id an AMS account | 
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**AccountId** | **string** | Source&#39;s unique identifier for account | 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeAccountRepresentativeRemovedAmsDetail

`func NewAeAccountRepresentativeRemovedAmsDetail(id string, auctionId string, accountId string, ) *AeAccountRepresentativeRemovedAmsDetail`

NewAeAccountRepresentativeRemovedAmsDetail instantiates a new AeAccountRepresentativeRemovedAmsDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAccountRepresentativeRemovedAmsDetailWithDefaults

`func NewAeAccountRepresentativeRemovedAmsDetailWithDefaults() *AeAccountRepresentativeRemovedAmsDetail`

NewAeAccountRepresentativeRemovedAmsDetailWithDefaults instantiates a new AeAccountRepresentativeRemovedAmsDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AeAccountRepresentativeRemovedAmsDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AeAccountRepresentativeRemovedAmsDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AeAccountRepresentativeRemovedAmsDetail) SetId(v string)`

SetId sets Id field to given value.


### GetAuctionId

`func (o *AeAccountRepresentativeRemovedAmsDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAccountRepresentativeRemovedAmsDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAccountRepresentativeRemovedAmsDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetAccountId

`func (o *AeAccountRepresentativeRemovedAmsDetail) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AeAccountRepresentativeRemovedAmsDetail) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AeAccountRepresentativeRemovedAmsDetail) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetInitiator

`func (o *AeAccountRepresentativeRemovedAmsDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAccountRepresentativeRemovedAmsDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAccountRepresentativeRemovedAmsDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAccountRepresentativeRemovedAmsDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


