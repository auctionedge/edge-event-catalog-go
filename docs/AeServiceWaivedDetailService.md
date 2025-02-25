# AeServiceWaivedDetailService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Class** | [**ServiceClassEnum**](ServiceClassEnum.md) |  | 
**Code** | Pointer to **string** | Service code | [optional] 

## Methods

### NewAeServiceWaivedDetailService

`func NewAeServiceWaivedDetailService(class ServiceClassEnum, ) *AeServiceWaivedDetailService`

NewAeServiceWaivedDetailService instantiates a new AeServiceWaivedDetailService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeServiceWaivedDetailServiceWithDefaults

`func NewAeServiceWaivedDetailServiceWithDefaults() *AeServiceWaivedDetailService`

NewAeServiceWaivedDetailServiceWithDefaults instantiates a new AeServiceWaivedDetailService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *AeServiceWaivedDetailService) GetClass() ServiceClassEnum`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *AeServiceWaivedDetailService) GetClassOk() (*ServiceClassEnum, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *AeServiceWaivedDetailService) SetClass(v ServiceClassEnum)`

SetClass sets Class field to given value.


### GetCode

`func (o *AeServiceWaivedDetailService) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AeServiceWaivedDetailService) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AeServiceWaivedDetailService) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AeServiceWaivedDetailService) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


