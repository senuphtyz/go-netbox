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

// End1 * `A` - A * `B` - B
type End1 string

// List of End_1
const (
	END1_A End1 = "A"
	END1_B End1 = "B"
)

// All allowed values of End1 enum
var AllowedEnd1EnumValues = []End1{
	"A",
	"B",
}

func (v *End1) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := End1(value)
	for _, existing := range AllowedEnd1EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid End1", value)
}

// NewEnd1FromValue returns a pointer to a valid End1
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnd1FromValue(v string) (*End1, error) {
	ev := End1(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for End1: valid values are %v", v, AllowedEnd1EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v End1) IsValid() bool {
	for _, existing := range AllowedEnd1EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to End_1 value
func (v End1) Ptr() *End1 {
	return &v
}

type NullableEnd1 struct {
	value *End1
	isSet bool
}

func (v NullableEnd1) Get() *End1 {
	return v.value
}

func (v *NullableEnd1) Set(val *End1) {
	v.value = val
	v.isSet = true
}

func (v NullableEnd1) IsSet() bool {
	return v.isSet
}

func (v *NullableEnd1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnd1(val *End1) *NullableEnd1 {
	return &NullableEnd1{value: val, isSet: true}
}

func (v NullableEnd1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnd1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

