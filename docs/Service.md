# Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Source&#39;s unique identifier for service | 
**ServiceClass** | [**ServiceClassEnum**](ServiceClassEnum.md) |  | 
**Name** | **string** | The name of the inspection option. | 
**Description** | **string** | Description of service being performed for information to be presented | 
**SpecificDetails** | Pointer to [**ServiceSpecificDetails**](ServiceSpecificDetails.md) |  | [optional] 
**RejectReason** | Pointer to **string** | Reason why this service option is not eligible | [optional] 
**Parameters** | Pointer to [**[]CommonServiceParameter**](CommonServiceParameter.md) |  | [optional] 

## Methods

### NewService

`func NewService(id string, serviceClass ServiceClassEnum, name string, description string, ) *Service`

NewService instantiates a new Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWithDefaults

`func NewServiceWithDefaults() *Service`

NewServiceWithDefaults instantiates a new Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Service) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Service) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Service) SetId(v string)`

SetId sets Id field to given value.


### GetServiceClass

`func (o *Service) GetServiceClass() ServiceClassEnum`

GetServiceClass returns the ServiceClass field if non-nil, zero value otherwise.

### GetServiceClassOk

`func (o *Service) GetServiceClassOk() (*ServiceClassEnum, bool)`

GetServiceClassOk returns a tuple with the ServiceClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceClass

`func (o *Service) SetServiceClass(v ServiceClassEnum)`

SetServiceClass sets ServiceClass field to given value.


### GetName

`func (o *Service) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Service) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Service) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Service) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Service) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Service) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSpecificDetails

`func (o *Service) GetSpecificDetails() ServiceSpecificDetails`

GetSpecificDetails returns the SpecificDetails field if non-nil, zero value otherwise.

### GetSpecificDetailsOk

`func (o *Service) GetSpecificDetailsOk() (*ServiceSpecificDetails, bool)`

GetSpecificDetailsOk returns a tuple with the SpecificDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificDetails

`func (o *Service) SetSpecificDetails(v ServiceSpecificDetails)`

SetSpecificDetails sets SpecificDetails field to given value.

### HasSpecificDetails

`func (o *Service) HasSpecificDetails() bool`

HasSpecificDetails returns a boolean if a field has been set.

### GetRejectReason

`func (o *Service) GetRejectReason() string`

GetRejectReason returns the RejectReason field if non-nil, zero value otherwise.

### GetRejectReasonOk

`func (o *Service) GetRejectReasonOk() (*string, bool)`

GetRejectReasonOk returns a tuple with the RejectReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectReason

`func (o *Service) SetRejectReason(v string)`

SetRejectReason sets RejectReason field to given value.

### HasRejectReason

`func (o *Service) HasRejectReason() bool`

HasRejectReason returns a boolean if a field has been set.

### GetParameters

`func (o *Service) GetParameters() []CommonServiceParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *Service) GetParametersOk() (*[]CommonServiceParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *Service) SetParameters(v []CommonServiceParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *Service) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


