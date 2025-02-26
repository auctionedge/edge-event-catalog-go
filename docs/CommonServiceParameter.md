# CommonServiceParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Service parameter name | [optional] 
**Value** | Pointer to **string** | Service parameter value checked against | [optional] 

## Methods

### NewCommonServiceParameter

`func NewCommonServiceParameter() *CommonServiceParameter`

NewCommonServiceParameter instantiates a new CommonServiceParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonServiceParameterWithDefaults

`func NewCommonServiceParameterWithDefaults() *CommonServiceParameter`

NewCommonServiceParameterWithDefaults instantiates a new CommonServiceParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CommonServiceParameter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommonServiceParameter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommonServiceParameter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CommonServiceParameter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *CommonServiceParameter) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CommonServiceParameter) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CommonServiceParameter) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CommonServiceParameter) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


