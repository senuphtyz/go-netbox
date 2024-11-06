/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.1.4 (4.1)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// InterfaceModeValue * `access` - Access * `tagged` - Tagged * `tagged-all` - Tagged (All)
type InterfaceModeValue string

// List of Interface_mode_value
const (
	INTERFACEMODEVALUE_ACCESS InterfaceModeValue = "access"
	INTERFACEMODEVALUE_TAGGED InterfaceModeValue = "tagged"
	INTERFACEMODEVALUE_TAGGED_ALL InterfaceModeValue = "tagged-all"
	INTERFACEMODEVALUE_EMPTY InterfaceModeValue = ""
)

// All allowed values of InterfaceModeValue enum
var AllowedInterfaceModeValueEnumValues = []InterfaceModeValue{
	"access",
	"tagged",
	"tagged-all",
	"",
}

func (v *InterfaceModeValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InterfaceModeValue(value)
	for _, existing := range AllowedInterfaceModeValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InterfaceModeValue", value)
}

// NewInterfaceModeValueFromValue returns a pointer to a valid InterfaceModeValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInterfaceModeValueFromValue(v string) (*InterfaceModeValue, error) {
	ev := InterfaceModeValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InterfaceModeValue: valid values are %v", v, AllowedInterfaceModeValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InterfaceModeValue) IsValid() bool {
	for _, existing := range AllowedInterfaceModeValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Interface_mode_value value
func (v InterfaceModeValue) Ptr() *InterfaceModeValue {
	return &v
}

type NullableInterfaceModeValue struct {
	value *InterfaceModeValue
	isSet bool
}

func (v NullableInterfaceModeValue) Get() *InterfaceModeValue {
	return v.value
}

func (v *NullableInterfaceModeValue) Set(val *InterfaceModeValue) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceModeValue) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceModeValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceModeValue(val *InterfaceModeValue) *NullableInterfaceModeValue {
	return &NullableInterfaceModeValue{value: val, isSet: true}
}

func (v NullableInterfaceModeValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceModeValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

