# Title

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exempt** | Pointer to **bool** | Defines whether or not the asset can even have a title.  For example a lawn mower would be title exempt. | [optional] [default to false]
**Number** | Pointer to **string** | A unique identifier, typically a string of 7-8 digits or characters, found on your car&#39;s title document. | [optional] 
**State** | Pointer to **string** | The state in which the asset&#39;s title was issued. | [optional] 
**ReceivedAt** | Pointer to **NullableTime** | The date and time at which the title document was received by the auction. | [optional] 

## Methods

### NewTitle

`func NewTitle() *Title`

NewTitle instantiates a new Title object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTitleWithDefaults

`func NewTitleWithDefaults() *Title`

NewTitleWithDefaults instantiates a new Title object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExempt

`func (o *Title) GetExempt() bool`

GetExempt returns the Exempt field if non-nil, zero value otherwise.

### GetExemptOk

`func (o *Title) GetExemptOk() (*bool, bool)`

GetExemptOk returns a tuple with the Exempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExempt

`func (o *Title) SetExempt(v bool)`

SetExempt sets Exempt field to given value.

### HasExempt

`func (o *Title) HasExempt() bool`

HasExempt returns a boolean if a field has been set.

### GetNumber

`func (o *Title) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Title) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Title) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *Title) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetState

`func (o *Title) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Title) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Title) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Title) HasState() bool`

HasState returns a boolean if a field has been set.

### GetReceivedAt

`func (o *Title) GetReceivedAt() time.Time`

GetReceivedAt returns the ReceivedAt field if non-nil, zero value otherwise.

### GetReceivedAtOk

`func (o *Title) GetReceivedAtOk() (*time.Time, bool)`

GetReceivedAtOk returns a tuple with the ReceivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedAt

`func (o *Title) SetReceivedAt(v time.Time)`

SetReceivedAt sets ReceivedAt field to given value.

### HasReceivedAt

`func (o *Title) HasReceivedAt() bool`

HasReceivedAt returns a boolean if a field has been set.

### SetReceivedAtNil

`func (o *Title) SetReceivedAtNil(b bool)`

 SetReceivedAtNil sets the value for ReceivedAt to be an explicit nil

### UnsetReceivedAt
`func (o *Title) UnsetReceivedAt()`

UnsetReceivedAt ensures that no value is present for ReceivedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


