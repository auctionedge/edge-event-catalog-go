/*
Edge Event Schemas

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package events

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ServiceOrder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceOrder{}

// ServiceOrder struct for ServiceOrder
type ServiceOrder struct {
	Class ServiceClassEnum `json:"class"`
	// Service code
	Code string `json:"code"`
	// Description of service chosen
	Description *string `json:"description,omitempty"`
	// Printer name, default configuration used if not provided
	Printer *string `json:"printer,omitempty"`
	// Not currently used
	CanPrint *bool `json:"canPrint,omitempty"`
	// Normal checks to allow service are bypassed.
	IsOverridden *bool `json:"isOverridden,omitempty"`
	// Fee amount
	FeeAmount NullableFloat32 `json:"feeAmount,omitempty"`
	// Note on desired service (not currently used)
	ServiceNote *string `json:"serviceNote,omitempty"`
}

type _ServiceOrder ServiceOrder

// NewServiceOrder instantiates a new ServiceOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceOrder(class ServiceClassEnum, code string) *ServiceOrder {
	this := ServiceOrder{}
	this.Class = class
	this.Code = code
	return &this
}

// NewServiceOrderWithDefaults instantiates a new ServiceOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceOrderWithDefaults() *ServiceOrder {
	this := ServiceOrder{}
	return &this
}

// GetClass returns the Class field value
func (o *ServiceOrder) GetClass() ServiceClassEnum {
	if o == nil {
		var ret ServiceClassEnum
		return ret
	}

	return o.Class
}

// GetClassOk returns a tuple with the Class field value
// and a boolean to check if the value has been set.
func (o *ServiceOrder) GetClassOk() (*ServiceClassEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Class, true
}

// SetClass sets field value
func (o *ServiceOrder) SetClass(v ServiceClassEnum) {
	o.Class = v
}

// GetCode returns the Code field value
func (o *ServiceOrder) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ServiceOrder) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ServiceOrder) SetCode(v string) {
	o.Code = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServiceOrder) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceOrder) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServiceOrder) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServiceOrder) SetDescription(v string) {
	o.Description = &v
}

// GetPrinter returns the Printer field value if set, zero value otherwise.
func (o *ServiceOrder) GetPrinter() string {
	if o == nil || IsNil(o.Printer) {
		var ret string
		return ret
	}
	return *o.Printer
}

// GetPrinterOk returns a tuple with the Printer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceOrder) GetPrinterOk() (*string, bool) {
	if o == nil || IsNil(o.Printer) {
		return nil, false
	}
	return o.Printer, true
}

// HasPrinter returns a boolean if a field has been set.
func (o *ServiceOrder) HasPrinter() bool {
	if o != nil && !IsNil(o.Printer) {
		return true
	}

	return false
}

// SetPrinter gets a reference to the given string and assigns it to the Printer field.
func (o *ServiceOrder) SetPrinter(v string) {
	o.Printer = &v
}

// GetCanPrint returns the CanPrint field value if set, zero value otherwise.
func (o *ServiceOrder) GetCanPrint() bool {
	if o == nil || IsNil(o.CanPrint) {
		var ret bool
		return ret
	}
	return *o.CanPrint
}

// GetCanPrintOk returns a tuple with the CanPrint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceOrder) GetCanPrintOk() (*bool, bool) {
	if o == nil || IsNil(o.CanPrint) {
		return nil, false
	}
	return o.CanPrint, true
}

// HasCanPrint returns a boolean if a field has been set.
func (o *ServiceOrder) HasCanPrint() bool {
	if o != nil && !IsNil(o.CanPrint) {
		return true
	}

	return false
}

// SetCanPrint gets a reference to the given bool and assigns it to the CanPrint field.
func (o *ServiceOrder) SetCanPrint(v bool) {
	o.CanPrint = &v
}

// GetIsOverridden returns the IsOverridden field value if set, zero value otherwise.
func (o *ServiceOrder) GetIsOverridden() bool {
	if o == nil || IsNil(o.IsOverridden) {
		var ret bool
		return ret
	}
	return *o.IsOverridden
}

// GetIsOverriddenOk returns a tuple with the IsOverridden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceOrder) GetIsOverriddenOk() (*bool, bool) {
	if o == nil || IsNil(o.IsOverridden) {
		return nil, false
	}
	return o.IsOverridden, true
}

// HasIsOverridden returns a boolean if a field has been set.
func (o *ServiceOrder) HasIsOverridden() bool {
	if o != nil && !IsNil(o.IsOverridden) {
		return true
	}

	return false
}

// SetIsOverridden gets a reference to the given bool and assigns it to the IsOverridden field.
func (o *ServiceOrder) SetIsOverridden(v bool) {
	o.IsOverridden = &v
}

// GetFeeAmount returns the FeeAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceOrder) GetFeeAmount() float32 {
	if o == nil || IsNil(o.FeeAmount.Get()) {
		var ret float32
		return ret
	}
	return *o.FeeAmount.Get()
}

// GetFeeAmountOk returns a tuple with the FeeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceOrder) GetFeeAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FeeAmount.Get(), o.FeeAmount.IsSet()
}

// HasFeeAmount returns a boolean if a field has been set.
func (o *ServiceOrder) HasFeeAmount() bool {
	if o != nil && o.FeeAmount.IsSet() {
		return true
	}

	return false
}

// SetFeeAmount gets a reference to the given NullableFloat32 and assigns it to the FeeAmount field.
func (o *ServiceOrder) SetFeeAmount(v float32) {
	o.FeeAmount.Set(&v)
}
// SetFeeAmountNil sets the value for FeeAmount to be an explicit nil
func (o *ServiceOrder) SetFeeAmountNil() {
	o.FeeAmount.Set(nil)
}

// UnsetFeeAmount ensures that no value is present for FeeAmount, not even an explicit nil
func (o *ServiceOrder) UnsetFeeAmount() {
	o.FeeAmount.Unset()
}

// GetServiceNote returns the ServiceNote field value if set, zero value otherwise.
func (o *ServiceOrder) GetServiceNote() string {
	if o == nil || IsNil(o.ServiceNote) {
		var ret string
		return ret
	}
	return *o.ServiceNote
}

// GetServiceNoteOk returns a tuple with the ServiceNote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceOrder) GetServiceNoteOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceNote) {
		return nil, false
	}
	return o.ServiceNote, true
}

// HasServiceNote returns a boolean if a field has been set.
func (o *ServiceOrder) HasServiceNote() bool {
	if o != nil && !IsNil(o.ServiceNote) {
		return true
	}

	return false
}

// SetServiceNote gets a reference to the given string and assigns it to the ServiceNote field.
func (o *ServiceOrder) SetServiceNote(v string) {
	o.ServiceNote = &v
}

func (o ServiceOrder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceOrder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["class"] = o.Class
	toSerialize["code"] = o.Code
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Printer) {
		toSerialize["printer"] = o.Printer
	}
	if !IsNil(o.CanPrint) {
		toSerialize["canPrint"] = o.CanPrint
	}
	if !IsNil(o.IsOverridden) {
		toSerialize["isOverridden"] = o.IsOverridden
	}
	if o.FeeAmount.IsSet() {
		toSerialize["feeAmount"] = o.FeeAmount.Get()
	}
	if !IsNil(o.ServiceNote) {
		toSerialize["serviceNote"] = o.ServiceNote
	}
	return toSerialize, nil
}

func (o *ServiceOrder) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"class",
		"code",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varServiceOrder := _ServiceOrder{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServiceOrder)

	if err != nil {
		return err
	}

	*o = ServiceOrder(varServiceOrder)

	return err
}

type NullableServiceOrder struct {
	value *ServiceOrder
	isSet bool
}

func (v NullableServiceOrder) Get() *ServiceOrder {
	return v.value
}

func (v *NullableServiceOrder) Set(val *ServiceOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceOrder(val *ServiceOrder) *NullableServiceOrder {
	return &NullableServiceOrder{value: val, isSet: true}
}

func (v NullableServiceOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


