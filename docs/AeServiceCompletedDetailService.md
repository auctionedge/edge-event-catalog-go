# AeServiceCompletedDetailService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Class** | [**ServiceClassEnum**](ServiceClassEnum.md) |  | 
**Code** | **string** | Service code | 
**FeeAmount** | Pointer to **NullableFloat32** | Fee amount | [optional] 

## Methods

### NewAeServiceCompletedDetailService

`func NewAeServiceCompletedDetailService(class ServiceClassEnum, code string, ) *AeServiceCompletedDetailService`

NewAeServiceCompletedDetailService instantiates a new AeServiceCompletedDetailService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeServiceCompletedDetailServiceWithDefaults

`func NewAeServiceCompletedDetailServiceWithDefaults() *AeServiceCompletedDetailService`

NewAeServiceCompletedDetailServiceWithDefaults instantiates a new AeServiceCompletedDetailService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *AeServiceCompletedDetailService) GetClass() ServiceClassEnum`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *AeServiceCompletedDetailService) GetClassOk() (*ServiceClassEnum, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *AeServiceCompletedDetailService) SetClass(v ServiceClassEnum)`

SetClass sets Class field to given value.


### GetCode

`func (o *AeServiceCompletedDetailService) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AeServiceCompletedDetailService) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AeServiceCompletedDetailService) SetCode(v string)`

SetCode sets Code field to given value.


### GetFeeAmount

`func (o *AeServiceCompletedDetailService) GetFeeAmount() float32`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *AeServiceCompletedDetailService) GetFeeAmountOk() (*float32, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *AeServiceCompletedDetailService) SetFeeAmount(v float32)`

SetFeeAmount sets FeeAmount field to given value.

### HasFeeAmount

`func (o *AeServiceCompletedDetailService) HasFeeAmount() bool`

HasFeeAmount returns a boolean if a field has been set.

### SetFeeAmountNil

`func (o *AeServiceCompletedDetailService) SetFeeAmountNil(b bool)`

 SetFeeAmountNil sets the value for FeeAmount to be an explicit nil

### UnsetFeeAmount
`func (o *AeServiceCompletedDetailService) UnsetFeeAmount()`

UnsetFeeAmount ensures that no value is present for FeeAmount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


