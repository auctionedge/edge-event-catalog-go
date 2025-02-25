# AeAdvisoryAccountRepresentativesAmsDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**Account** | [**CommonAmsAccountPointer**](CommonAmsAccountPointer.md) |  | 
**UpdatedAt** | **time.Time** | The time in ISO-8601 formatted at which the record was generated | 
**Representatives** | [**[]AeAdvisoryAccountRepresentativesAmsDetailRepresentativesInner**](AeAdvisoryAccountRepresentativesAmsDetailRepresentativesInner.md) | The representatives and their permissions for the account | 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeAdvisoryAccountRepresentativesAmsDetail

`func NewAeAdvisoryAccountRepresentativesAmsDetail(auctionId string, account CommonAmsAccountPointer, updatedAt time.Time, representatives []AeAdvisoryAccountRepresentativesAmsDetailRepresentativesInner, ) *AeAdvisoryAccountRepresentativesAmsDetail`

NewAeAdvisoryAccountRepresentativesAmsDetail instantiates a new AeAdvisoryAccountRepresentativesAmsDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAdvisoryAccountRepresentativesAmsDetailWithDefaults

`func NewAeAdvisoryAccountRepresentativesAmsDetailWithDefaults() *AeAdvisoryAccountRepresentativesAmsDetail`

NewAeAdvisoryAccountRepresentativesAmsDetailWithDefaults instantiates a new AeAdvisoryAccountRepresentativesAmsDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeAdvisoryAccountRepresentativesAmsDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAdvisoryAccountRepresentativesAmsDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAdvisoryAccountRepresentativesAmsDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetAccount

`func (o *AeAdvisoryAccountRepresentativesAmsDetail) GetAccount() CommonAmsAccountPointer`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AeAdvisoryAccountRepresentativesAmsDetail) GetAccountOk() (*CommonAmsAccountPointer, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AeAdvisoryAccountRepresentativesAmsDetail) SetAccount(v CommonAmsAccountPointer)`

SetAccount sets Account field to given value.


### GetUpdatedAt

`func (o *AeAdvisoryAccountRepresentativesAmsDetail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AeAdvisoryAccountRepresentativesAmsDetail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AeAdvisoryAccountRepresentativesAmsDetail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRepresentatives

`func (o *AeAdvisoryAccountRepresentativesAmsDetail) GetRepresentatives() []AeAdvisoryAccountRepresentativesAmsDetailRepresentativesInner`

GetRepresentatives returns the Representatives field if non-nil, zero value otherwise.

### GetRepresentativesOk

`func (o *AeAdvisoryAccountRepresentativesAmsDetail) GetRepresentativesOk() (*[]AeAdvisoryAccountRepresentativesAmsDetailRepresentativesInner, bool)`

GetRepresentativesOk returns a tuple with the Representatives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepresentatives

`func (o *AeAdvisoryAccountRepresentativesAmsDetail) SetRepresentatives(v []AeAdvisoryAccountRepresentativesAmsDetailRepresentativesInner)`

SetRepresentatives sets Representatives field to given value.


### GetInitiator

`func (o *AeAdvisoryAccountRepresentativesAmsDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAdvisoryAccountRepresentativesAmsDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAdvisoryAccountRepresentativesAmsDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAdvisoryAccountRepresentativesAmsDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


