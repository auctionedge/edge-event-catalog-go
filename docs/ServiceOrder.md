# ServiceOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Class** | [**ServiceClassEnum**](ServiceClassEnum.md) |  | 
**Code** | **string** | Service code | 
**Description** | Pointer to **string** | Description of service chosen | [optional] 
**Printer** | Pointer to **string** | Printer name, default configuration used if not provided | [optional] 
**CanPrint** | Pointer to **bool** | Not currently used | [optional] 
**IsOverridden** | Pointer to **bool** | Normal checks to allow service are bypassed. | [optional] 
**FeeAmount** | Pointer to **NullableFloat32** | Fee amount | [optional] 
**ServiceNote** | Pointer to **string** | Note on desired service (not currently used) | [optional] 

## Methods

### NewServiceOrder

`func NewServiceOrder(class ServiceClassEnum, code string, ) *ServiceOrder`

NewServiceOrder instantiates a new ServiceOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceOrderWithDefaults

`func NewServiceOrderWithDefaults() *ServiceOrder`

NewServiceOrderWithDefaults instantiates a new ServiceOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *ServiceOrder) GetClass() ServiceClassEnum`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *ServiceOrder) GetClassOk() (*ServiceClassEnum, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *ServiceOrder) SetClass(v ServiceClassEnum)`

SetClass sets Class field to given value.


### GetCode

`func (o *ServiceOrder) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ServiceOrder) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ServiceOrder) SetCode(v string)`

SetCode sets Code field to given value.


### GetDescription

`func (o *ServiceOrder) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceOrder) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceOrder) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceOrder) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPrinter

`func (o *ServiceOrder) GetPrinter() string`

GetPrinter returns the Printer field if non-nil, zero value otherwise.

### GetPrinterOk

`func (o *ServiceOrder) GetPrinterOk() (*string, bool)`

GetPrinterOk returns a tuple with the Printer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrinter

`func (o *ServiceOrder) SetPrinter(v string)`

SetPrinter sets Printer field to given value.

### HasPrinter

`func (o *ServiceOrder) HasPrinter() bool`

HasPrinter returns a boolean if a field has been set.

### GetCanPrint

`func (o *ServiceOrder) GetCanPrint() bool`

GetCanPrint returns the CanPrint field if non-nil, zero value otherwise.

### GetCanPrintOk

`func (o *ServiceOrder) GetCanPrintOk() (*bool, bool)`

GetCanPrintOk returns a tuple with the CanPrint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanPrint

`func (o *ServiceOrder) SetCanPrint(v bool)`

SetCanPrint sets CanPrint field to given value.

### HasCanPrint

`func (o *ServiceOrder) HasCanPrint() bool`

HasCanPrint returns a boolean if a field has been set.

### GetIsOverridden

`func (o *ServiceOrder) GetIsOverridden() bool`

GetIsOverridden returns the IsOverridden field if non-nil, zero value otherwise.

### GetIsOverriddenOk

`func (o *ServiceOrder) GetIsOverriddenOk() (*bool, bool)`

GetIsOverriddenOk returns a tuple with the IsOverridden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOverridden

`func (o *ServiceOrder) SetIsOverridden(v bool)`

SetIsOverridden sets IsOverridden field to given value.

### HasIsOverridden

`func (o *ServiceOrder) HasIsOverridden() bool`

HasIsOverridden returns a boolean if a field has been set.

### GetFeeAmount

`func (o *ServiceOrder) GetFeeAmount() float32`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *ServiceOrder) GetFeeAmountOk() (*float32, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *ServiceOrder) SetFeeAmount(v float32)`

SetFeeAmount sets FeeAmount field to given value.

### HasFeeAmount

`func (o *ServiceOrder) HasFeeAmount() bool`

HasFeeAmount returns a boolean if a field has been set.

### SetFeeAmountNil

`func (o *ServiceOrder) SetFeeAmountNil(b bool)`

 SetFeeAmountNil sets the value for FeeAmount to be an explicit nil

### UnsetFeeAmount
`func (o *ServiceOrder) UnsetFeeAmount()`

UnsetFeeAmount ensures that no value is present for FeeAmount, not even an explicit nil
### GetServiceNote

`func (o *ServiceOrder) GetServiceNote() string`

GetServiceNote returns the ServiceNote field if non-nil, zero value otherwise.

### GetServiceNoteOk

`func (o *ServiceOrder) GetServiceNoteOk() (*string, bool)`

GetServiceNoteOk returns a tuple with the ServiceNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceNote

`func (o *ServiceOrder) SetServiceNote(v string)`

SetServiceNote sets ServiceNote field to given value.

### HasServiceNote

`func (o *ServiceOrder) HasServiceNote() bool`

HasServiceNote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


