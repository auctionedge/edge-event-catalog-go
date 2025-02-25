/*
Edge Event Schemas

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package events

import (
	"encoding/json"
	"fmt"
)

// GatepassType Allowed values:     * 'G' - Buyer or Seller gatepass (depending on asset status) - Default     * 'T' - Test drive gatepass     * 'R' - Repair gatepass 
type GatepassType string

// List of gatepass-type
const (
	G GatepassType = "G"
	T GatepassType = "T"
	R GatepassType = "R"
)

// All allowed values of GatepassType enum
var AllowedGatepassTypeEnumValues = []GatepassType{
	"G",
	"T",
	"R",
}

func (v *GatepassType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GatepassType(value)
	for _, existing := range AllowedGatepassTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GatepassType", value)
}

// NewGatepassTypeFromValue returns a pointer to a valid GatepassType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGatepassTypeFromValue(v string) (*GatepassType, error) {
	ev := GatepassType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GatepassType: valid values are %v", v, AllowedGatepassTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GatepassType) IsValid() bool {
	for _, existing := range AllowedGatepassTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to gatepass-type value
func (v GatepassType) Ptr() *GatepassType {
	return &v
}

type NullableGatepassType struct {
	value *GatepassType
	isSet bool
}

func (v NullableGatepassType) Get() *GatepassType {
	return v.value
}

func (v *NullableGatepassType) Set(val *GatepassType) {
	v.value = val
	v.isSet = true
}

func (v NullableGatepassType) IsSet() bool {
	return v.isSet
}

func (v *NullableGatepassType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatepassType(val *GatepassType) *NullableGatepassType {
	return &NullableGatepassType{value: val, isSet: true}
}

func (v NullableGatepassType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatepassType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

