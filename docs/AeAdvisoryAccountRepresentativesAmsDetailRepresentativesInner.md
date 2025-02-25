# AeAdvisoryAccountRepresentativesAmsDetailRepresentativesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The representative id an AMS account | 
**AaId** | **NullableString** | The Auction Access ID of the AMS Representative account. | 
**RepNumber** | Pointer to **float32** | Rep number as denoted in ASI, AOS to use same value as representative-id | [optional] 
**PersonInfo** | Pointer to [**CommonAmsPersonInfo**](CommonAmsPersonInfo.md) |  | [optional] 
**InternalNotes** | Pointer to **string** | Any notes on representative | [optional] 
**CanBuy** | Pointer to **bool** | True if representative is allowed to buy a vehicle, false if not | [optional] 

## Methods

### NewAeAdvisoryAccountRepresentativesAmsDetailRepresentativesInner

`func NewAeAdvisoryAccountRepresentativesAmsDetailRepresentativesInner(id string, aaId NullableString, ) *AeAdvisoryAccountRepresentativesAmsDetailRepresentativesInner`

NewAeAdvisoryAccountRepresentativesAmsDetailRepresentativesInner instantiates a new AeAdvisoryAccountRepresentativesAmsDetailRepresentativesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAdvisoryAccountRepresentativesAmsDetailRepresentativesInnerWithDefaults

`func NewAeAdvisoryAccountRepresentativesAmsDetailRepresentativesInnerWithDefaults() *AeAdvisoryAccountRepresentativesAmsDetailRepresentativesInner`

NewAeAdvisoryAccountRepresentativesAmsDetailRepresentativesInnerWithDefaults instantiates a new AeAdvisoryAccountRepresentativesAmsDetailRepresentativesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AeAdvisoryAccountRepresentativesAmsDetailRepresentativesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AeAdvisoryAccountRepresentativesAmsDetailRepresentativesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AeAdvisoryAccountRepresentativesAmsDetailRepresentativesInner) SetId(v string)`

SetId sets Id field to given value.


### GetAaId

`func (o *AeAdvisoryAccountRepresentativesAmsDetailRepresentativesInner) GetAaId() string`

GetAaId returns the AaId field if non-nil, zero value otherwise.

### GetAaIdOk

`func (o *AeAdvisoryAccountRepresentativesAmsDetailRepresentativesInner) GetAaIdOk() (*string, bool)`

GetAaIdOk returns a tuple with the AaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAaId

`func (o *AeAdvisoryAccountRepresentativesAmsDetailRepresentativesInner) SetAaId(v string)`

SetAaId sets AaId field to given value.


### SetAaIdNil

`func (o *AeAdvisoryAccountRepresentativesAmsDetailRepresentativesInner) SetAaIdNil(b bool)`

 SetAaIdNil sets the value for AaId to be an explicit nil

### UnsetAaId
`func (o *AeAdvisoryAccountRepresentativesAmsDetailRepresentativesInner) UnsetAaId()`

UnsetAaId ensures that no value is present for AaId, not even an explicit nil
### GetRepNumber

`func (o *AeAdvisoryAccountRepresentativesAmsDetailRepresentativesInner) GetRepNumber() float32`

GetRepNumber returns the RepNumber field if non-nil, zero value otherwise.

### GetRepNumberOk

`func (o *AeAdvisoryAccountRepresentativesAmsDetailRepresentativesInner) GetRepNumberOk() (*float32, bool)`

GetRepNumberOk returns a tuple with the RepNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepNumber

`func (o *AeAdvisoryAccountRepresentativesAmsDetailRepresentativesInner) SetRepNumber(v float32)`

SetRepNumber sets RepNumber field to given value.

### HasRepNumber

`func (o *AeAdvisoryAccountRepresentativesAmsDetailRepresentativesInner) HasRepNumber() bool`

HasRepNumber returns a boolean if a field has been set.

### GetPersonInfo

`func (o *AeAdvisoryAccountRepresentativesAmsDetailRepresentativesInner) GetPersonInfo() CommonAmsPersonInfo`

GetPersonInfo returns the PersonInfo field if non-nil, zero value otherwise.

### GetPersonInfoOk

`func (o *AeAdvisoryAccountRepresentativesAmsDetailRepresentativesInner) GetPersonInfoOk() (*CommonAmsPersonInfo, bool)`

GetPersonInfoOk returns a tuple with the PersonInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonInfo

`func (o *AeAdvisoryAccountRepresentativesAmsDetailRepresentativesInner) SetPersonInfo(v CommonAmsPersonInfo)`

SetPersonInfo sets PersonInfo field to given value.

### HasPersonInfo

`func (o *AeAdvisoryAccountRepresentativesAmsDetailRepresentativesInner) HasPersonInfo() bool`

HasPersonInfo returns a boolean if a field has been set.

### GetInternalNotes

`func (o *AeAdvisoryAccountRepresentativesAmsDetailRepresentativesInner) GetInternalNotes() string`

GetInternalNotes returns the InternalNotes field if non-nil, zero value otherwise.

### GetInternalNotesOk

`func (o *AeAdvisoryAccountRepresentativesAmsDetailRepresentativesInner) GetInternalNotesOk() (*string, bool)`

GetInternalNotesOk returns a tuple with the InternalNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalNotes

`func (o *AeAdvisoryAccountRepresentativesAmsDetailRepresentativesInner) SetInternalNotes(v string)`

SetInternalNotes sets InternalNotes field to given value.

### HasInternalNotes

`func (o *AeAdvisoryAccountRepresentativesAmsDetailRepresentativesInner) HasInternalNotes() bool`

HasInternalNotes returns a boolean if a field has been set.

### GetCanBuy

`func (o *AeAdvisoryAccountRepresentativesAmsDetailRepresentativesInner) GetCanBuy() bool`

GetCanBuy returns the CanBuy field if non-nil, zero value otherwise.

### GetCanBuyOk

`func (o *AeAdvisoryAccountRepresentativesAmsDetailRepresentativesInner) GetCanBuyOk() (*bool, bool)`

GetCanBuyOk returns a tuple with the CanBuy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanBuy

`func (o *AeAdvisoryAccountRepresentativesAmsDetailRepresentativesInner) SetCanBuy(v bool)`

SetCanBuy sets CanBuy field to given value.

### HasCanBuy

`func (o *AeAdvisoryAccountRepresentativesAmsDetailRepresentativesInner) HasCanBuy() bool`

HasCanBuy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


