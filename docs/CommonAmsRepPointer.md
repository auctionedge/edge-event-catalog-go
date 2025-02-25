# CommonAmsRepPointer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RepId** | **NullableString** | The rep id an AMS account.  This common object is deprecated, please use common.ams.representative-id instead. | 
**AaId** | **NullableString** | The Auction Access ID of the AMS Representative account. | 

## Methods

### NewCommonAmsRepPointer

`func NewCommonAmsRepPointer(repId NullableString, aaId NullableString, ) *CommonAmsRepPointer`

NewCommonAmsRepPointer instantiates a new CommonAmsRepPointer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonAmsRepPointerWithDefaults

`func NewCommonAmsRepPointerWithDefaults() *CommonAmsRepPointer`

NewCommonAmsRepPointerWithDefaults instantiates a new CommonAmsRepPointer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepId

`func (o *CommonAmsRepPointer) GetRepId() string`

GetRepId returns the RepId field if non-nil, zero value otherwise.

### GetRepIdOk

`func (o *CommonAmsRepPointer) GetRepIdOk() (*string, bool)`

GetRepIdOk returns a tuple with the RepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepId

`func (o *CommonAmsRepPointer) SetRepId(v string)`

SetRepId sets RepId field to given value.


### SetRepIdNil

`func (o *CommonAmsRepPointer) SetRepIdNil(b bool)`

 SetRepIdNil sets the value for RepId to be an explicit nil

### UnsetRepId
`func (o *CommonAmsRepPointer) UnsetRepId()`

UnsetRepId ensures that no value is present for RepId, not even an explicit nil
### GetAaId

`func (o *CommonAmsRepPointer) GetAaId() string`

GetAaId returns the AaId field if non-nil, zero value otherwise.

### GetAaIdOk

`func (o *CommonAmsRepPointer) GetAaIdOk() (*string, bool)`

GetAaIdOk returns a tuple with the AaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAaId

`func (o *CommonAmsRepPointer) SetAaId(v string)`

SetAaId sets AaId field to given value.


### SetAaIdNil

`func (o *CommonAmsRepPointer) SetAaIdNil(b bool)`

 SetAaIdNil sets the value for AaId to be an explicit nil

### UnsetAaId
`func (o *CommonAmsRepPointer) UnsetAaId()`

UnsetAaId ensures that no value is present for AaId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


