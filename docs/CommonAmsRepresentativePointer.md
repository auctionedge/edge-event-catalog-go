# CommonAmsRepresentativePointer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The representative id an AMS account | 
**AaId** | **NullableString** | The Auction Access ID of the AMS Representative account. | 
**RepNumber** | Pointer to **float32** | Rep number as denoted in ASI, AOS to use same value as representative-id | [optional] 

## Methods

### NewCommonAmsRepresentativePointer

`func NewCommonAmsRepresentativePointer(id string, aaId NullableString, ) *CommonAmsRepresentativePointer`

NewCommonAmsRepresentativePointer instantiates a new CommonAmsRepresentativePointer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonAmsRepresentativePointerWithDefaults

`func NewCommonAmsRepresentativePointerWithDefaults() *CommonAmsRepresentativePointer`

NewCommonAmsRepresentativePointerWithDefaults instantiates a new CommonAmsRepresentativePointer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommonAmsRepresentativePointer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommonAmsRepresentativePointer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommonAmsRepresentativePointer) SetId(v string)`

SetId sets Id field to given value.


### GetAaId

`func (o *CommonAmsRepresentativePointer) GetAaId() string`

GetAaId returns the AaId field if non-nil, zero value otherwise.

### GetAaIdOk

`func (o *CommonAmsRepresentativePointer) GetAaIdOk() (*string, bool)`

GetAaIdOk returns a tuple with the AaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAaId

`func (o *CommonAmsRepresentativePointer) SetAaId(v string)`

SetAaId sets AaId field to given value.


### SetAaIdNil

`func (o *CommonAmsRepresentativePointer) SetAaIdNil(b bool)`

 SetAaIdNil sets the value for AaId to be an explicit nil

### UnsetAaId
`func (o *CommonAmsRepresentativePointer) UnsetAaId()`

UnsetAaId ensures that no value is present for AaId, not even an explicit nil
### GetRepNumber

`func (o *CommonAmsRepresentativePointer) GetRepNumber() float32`

GetRepNumber returns the RepNumber field if non-nil, zero value otherwise.

### GetRepNumberOk

`func (o *CommonAmsRepresentativePointer) GetRepNumberOk() (*float32, bool)`

GetRepNumberOk returns a tuple with the RepNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepNumber

`func (o *CommonAmsRepresentativePointer) SetRepNumber(v float32)`

SetRepNumber sets RepNumber field to given value.

### HasRepNumber

`func (o *CommonAmsRepresentativePointer) HasRepNumber() bool`

HasRepNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


