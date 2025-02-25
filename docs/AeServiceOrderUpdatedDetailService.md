# AeServiceOrderUpdatedDetailService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Class** | [**ServiceClassEnum**](ServiceClassEnum.md) |  | 
**Code** | **string** | Service code | 
**Status** | **string** | Status of current state of service | 
**Note** | Pointer to **string** | Notes, summaries, or details about service that may be shared with end user | [optional] 

## Methods

### NewAeServiceOrderUpdatedDetailService

`func NewAeServiceOrderUpdatedDetailService(class ServiceClassEnum, code string, status string, ) *AeServiceOrderUpdatedDetailService`

NewAeServiceOrderUpdatedDetailService instantiates a new AeServiceOrderUpdatedDetailService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeServiceOrderUpdatedDetailServiceWithDefaults

`func NewAeServiceOrderUpdatedDetailServiceWithDefaults() *AeServiceOrderUpdatedDetailService`

NewAeServiceOrderUpdatedDetailServiceWithDefaults instantiates a new AeServiceOrderUpdatedDetailService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *AeServiceOrderUpdatedDetailService) GetClass() ServiceClassEnum`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *AeServiceOrderUpdatedDetailService) GetClassOk() (*ServiceClassEnum, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *AeServiceOrderUpdatedDetailService) SetClass(v ServiceClassEnum)`

SetClass sets Class field to given value.


### GetCode

`func (o *AeServiceOrderUpdatedDetailService) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AeServiceOrderUpdatedDetailService) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AeServiceOrderUpdatedDetailService) SetCode(v string)`

SetCode sets Code field to given value.


### GetStatus

`func (o *AeServiceOrderUpdatedDetailService) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AeServiceOrderUpdatedDetailService) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AeServiceOrderUpdatedDetailService) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetNote

`func (o *AeServiceOrderUpdatedDetailService) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *AeServiceOrderUpdatedDetailService) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *AeServiceOrderUpdatedDetailService) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *AeServiceOrderUpdatedDetailService) HasNote() bool`

HasNote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


