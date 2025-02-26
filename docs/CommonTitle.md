# CommonTitle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exempt** | Pointer to **bool** | Defines whether or not the asset can even have a title.  For example a lawn mower would be title exempt. | [optional] [default to false]
**Number** | Pointer to **string** | A unique identifier, typically a string of 7-8 digits or characters, found on your car&#39;s title document. | [optional] 
**State** | Pointer to **string** | The state in which the asset&#39;s title was issued. | [optional] 
**ReceivedAt** | Pointer to **NullableTime** | The date and time at which the title document was received by the auction. | [optional] 

## Methods

### NewCommonTitle

`func NewCommonTitle() *CommonTitle`

NewCommonTitle instantiates a new CommonTitle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonTitleWithDefaults

`func NewCommonTitleWithDefaults() *CommonTitle`

NewCommonTitleWithDefaults instantiates a new CommonTitle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExempt

`func (o *CommonTitle) GetExempt() bool`

GetExempt returns the Exempt field if non-nil, zero value otherwise.

### GetExemptOk

`func (o *CommonTitle) GetExemptOk() (*bool, bool)`

GetExemptOk returns a tuple with the Exempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExempt

`func (o *CommonTitle) SetExempt(v bool)`

SetExempt sets Exempt field to given value.

### HasExempt

`func (o *CommonTitle) HasExempt() bool`

HasExempt returns a boolean if a field has been set.

### GetNumber

`func (o *CommonTitle) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *CommonTitle) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *CommonTitle) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *CommonTitle) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetState

`func (o *CommonTitle) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CommonTitle) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CommonTitle) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CommonTitle) HasState() bool`

HasState returns a boolean if a field has been set.

### GetReceivedAt

`func (o *CommonTitle) GetReceivedAt() time.Time`

GetReceivedAt returns the ReceivedAt field if non-nil, zero value otherwise.

### GetReceivedAtOk

`func (o *CommonTitle) GetReceivedAtOk() (*time.Time, bool)`

GetReceivedAtOk returns a tuple with the ReceivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedAt

`func (o *CommonTitle) SetReceivedAt(v time.Time)`

SetReceivedAt sets ReceivedAt field to given value.

### HasReceivedAt

`func (o *CommonTitle) HasReceivedAt() bool`

HasReceivedAt returns a boolean if a field has been set.

### SetReceivedAtNil

`func (o *CommonTitle) SetReceivedAtNil(b bool)`

 SetReceivedAtNil sets the value for ReceivedAt to be an explicit nil

### UnsetReceivedAt
`func (o *CommonTitle) UnsetReceivedAt()`

UnsetReceivedAt ensures that no value is present for ReceivedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


