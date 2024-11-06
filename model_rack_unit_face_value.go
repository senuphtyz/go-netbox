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

// RackUnitFaceValue * `front` - Front * `rear` - Rear
type RackUnitFaceValue string

// List of RackUnit_face_value
const (
	RACKUNITFACEVALUE_FRONT RackUnitFaceValue = "front"
	RACKUNITFACEVALUE_REAR  RackUnitFaceValue = "rear"
)

// All allowed values of RackUnitFaceValue enum
var AllowedRackUnitFaceValueEnumValues = []RackUnitFaceValue{
	"front",
	"rear",
}

func (v *RackUnitFaceValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RackUnitFaceValue(value)
	for _, existing := range AllowedRackUnitFaceValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RackUnitFaceValue", value)
}

// NewRackUnitFaceValueFromValue returns a pointer to a valid RackUnitFaceValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRackUnitFaceValueFromValue(v string) (*RackUnitFaceValue, error) {
	ev := RackUnitFaceValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RackUnitFaceValue: valid values are %v", v, AllowedRackUnitFaceValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RackUnitFaceValue) IsValid() bool {
	for _, existing := range AllowedRackUnitFaceValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RackUnit_face_value value
func (v RackUnitFaceValue) Ptr() *RackUnitFaceValue {
	return &v
}

type NullableRackUnitFaceValue struct {
	value *RackUnitFaceValue
	isSet bool
}

func (v NullableRackUnitFaceValue) Get() *RackUnitFaceValue {
	return v.value
}

func (v *NullableRackUnitFaceValue) Set(val *RackUnitFaceValue) {
	v.value = val
	v.isSet = true
}

func (v NullableRackUnitFaceValue) IsSet() bool {
	return v.isSet
}

func (v *NullableRackUnitFaceValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRackUnitFaceValue(val *RackUnitFaceValue) *NullableRackUnitFaceValue {
	return &NullableRackUnitFaceValue{value: val, isSet: true}
}

func (v NullableRackUnitFaceValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRackUnitFaceValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
