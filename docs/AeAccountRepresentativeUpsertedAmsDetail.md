# AeAccountRepresentativeUpsertedAmsDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The representative id an AMS account | 
**AaId** | **NullableString** | The Auction Access ID of the AMS Representative account. | 
**RepNumber** | Pointer to **float32** | Rep number as denoted in ASI, AOS to use same value as representative-id | [optional] 
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**AccountId** | Pointer to **string** | Source&#39;s unique identifier for account | [optional] 
**PersonInfo** | Pointer to [**CommonAmsPersonInfo**](CommonAmsPersonInfo.md) |  | [optional] 
**InternalNotes** | Pointer to **string** | Any notes on representative | [optional] 
**CanBuy** | Pointer to **bool** | True if representative is allowed to buy a vehicle, false if not | [optional] 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeAccountRepresentativeUpsertedAmsDetail

`func NewAeAccountRepresentativeUpsertedAmsDetail(id string, aaId NullableString, auctionId string, ) *AeAccountRepresentativeUpsertedAmsDetail`

NewAeAccountRepresentativeUpsertedAmsDetail instantiates a new AeAccountRepresentativeUpsertedAmsDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAccountRepresentativeUpsertedAmsDetailWithDefaults

`func NewAeAccountRepresentativeUpsertedAmsDetailWithDefaults() *AeAccountRepresentativeUpsertedAmsDetail`

NewAeAccountRepresentativeUpsertedAmsDetailWithDefaults instantiates a new AeAccountRepresentativeUpsertedAmsDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AeAccountRepresentativeUpsertedAmsDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AeAccountRepresentativeUpsertedAmsDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AeAccountRepresentativeUpsertedAmsDetail) SetId(v string)`

SetId sets Id field to given value.


### GetAaId

`func (o *AeAccountRepresentativeUpsertedAmsDetail) GetAaId() string`

GetAaId returns the AaId field if non-nil, zero value otherwise.

### GetAaIdOk

`func (o *AeAccountRepresentativeUpsertedAmsDetail) GetAaIdOk() (*string, bool)`

GetAaIdOk returns a tuple with the AaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAaId

`func (o *AeAccountRepresentativeUpsertedAmsDetail) SetAaId(v string)`

SetAaId sets AaId field to given value.


### SetAaIdNil

`func (o *AeAccountRepresentativeUpsertedAmsDetail) SetAaIdNil(b bool)`

 SetAaIdNil sets the value for AaId to be an explicit nil

### UnsetAaId
`func (o *AeAccountRepresentativeUpsertedAmsDetail) UnsetAaId()`

UnsetAaId ensures that no value is present for AaId, not even an explicit nil
### GetRepNumber

`func (o *AeAccountRepresentativeUpsertedAmsDetail) GetRepNumber() float32`

GetRepNumber returns the RepNumber field if non-nil, zero value otherwise.

### GetRepNumberOk

`func (o *AeAccountRepresentativeUpsertedAmsDetail) GetRepNumberOk() (*float32, bool)`

GetRepNumberOk returns a tuple with the RepNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepNumber

`func (o *AeAccountRepresentativeUpsertedAmsDetail) SetRepNumber(v float32)`

SetRepNumber sets RepNumber field to given value.

### HasRepNumber

`func (o *AeAccountRepresentativeUpsertedAmsDetail) HasRepNumber() bool`

HasRepNumber returns a boolean if a field has been set.

### GetAuctionId

`func (o *AeAccountRepresentativeUpsertedAmsDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAccountRepresentativeUpsertedAmsDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAccountRepresentativeUpsertedAmsDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetAccountId

`func (o *AeAccountRepresentativeUpsertedAmsDetail) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AeAccountRepresentativeUpsertedAmsDetail) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AeAccountRepresentativeUpsertedAmsDetail) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AeAccountRepresentativeUpsertedAmsDetail) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetPersonInfo

`func (o *AeAccountRepresentativeUpsertedAmsDetail) GetPersonInfo() CommonAmsPersonInfo`

GetPersonInfo returns the PersonInfo field if non-nil, zero value otherwise.

### GetPersonInfoOk

`func (o *AeAccountRepresentativeUpsertedAmsDetail) GetPersonInfoOk() (*CommonAmsPersonInfo, bool)`

GetPersonInfoOk returns a tuple with the PersonInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonInfo

`func (o *AeAccountRepresentativeUpsertedAmsDetail) SetPersonInfo(v CommonAmsPersonInfo)`

SetPersonInfo sets PersonInfo field to given value.

### HasPersonInfo

`func (o *AeAccountRepresentativeUpsertedAmsDetail) HasPersonInfo() bool`

HasPersonInfo returns a boolean if a field has been set.

### GetInternalNotes

`func (o *AeAccountRepresentativeUpsertedAmsDetail) GetInternalNotes() string`

GetInternalNotes returns the InternalNotes field if non-nil, zero value otherwise.

### GetInternalNotesOk

`func (o *AeAccountRepresentativeUpsertedAmsDetail) GetInternalNotesOk() (*string, bool)`

GetInternalNotesOk returns a tuple with the InternalNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalNotes

`func (o *AeAccountRepresentativeUpsertedAmsDetail) SetInternalNotes(v string)`

SetInternalNotes sets InternalNotes field to given value.

### HasInternalNotes

`func (o *AeAccountRepresentativeUpsertedAmsDetail) HasInternalNotes() bool`

HasInternalNotes returns a boolean if a field has been set.

### GetCanBuy

`func (o *AeAccountRepresentativeUpsertedAmsDetail) GetCanBuy() bool`

GetCanBuy returns the CanBuy field if non-nil, zero value otherwise.

### GetCanBuyOk

`func (o *AeAccountRepresentativeUpsertedAmsDetail) GetCanBuyOk() (*bool, bool)`

GetCanBuyOk returns a tuple with the CanBuy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanBuy

`func (o *AeAccountRepresentativeUpsertedAmsDetail) SetCanBuy(v bool)`

SetCanBuy sets CanBuy field to given value.

### HasCanBuy

`func (o *AeAccountRepresentativeUpsertedAmsDetail) HasCanBuy() bool`

HasCanBuy returns a boolean if a field has been set.

### GetInitiator

`func (o *AeAccountRepresentativeUpsertedAmsDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAccountRepresentativeUpsertedAmsDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAccountRepresentativeUpsertedAmsDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAccountRepresentativeUpsertedAmsDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


