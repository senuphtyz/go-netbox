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

// DeviceTypeWeightUnitLabel the model 'DeviceTypeWeightUnitLabel'
type DeviceTypeWeightUnitLabel string

// List of DeviceType_weight_unit_label
const (
	DEVICETYPEWEIGHTUNITLABEL_KILOGRAMS DeviceTypeWeightUnitLabel = "Kilograms"
	DEVICETYPEWEIGHTUNITLABEL_GRAMS DeviceTypeWeightUnitLabel = "Grams"
	DEVICETYPEWEIGHTUNITLABEL_POUNDS DeviceTypeWeightUnitLabel = "Pounds"
	DEVICETYPEWEIGHTUNITLABEL_OUNCES DeviceTypeWeightUnitLabel = "Ounces"
)

// All allowed values of DeviceTypeWeightUnitLabel enum
var AllowedDeviceTypeWeightUnitLabelEnumValues = []DeviceTypeWeightUnitLabel{
	"Kilograms",
	"Grams",
	"Pounds",
	"Ounces",
}

func (v *DeviceTypeWeightUnitLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeviceTypeWeightUnitLabel(value)
	for _, existing := range AllowedDeviceTypeWeightUnitLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeviceTypeWeightUnitLabel", value)
}

// NewDeviceTypeWeightUnitLabelFromValue returns a pointer to a valid DeviceTypeWeightUnitLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeviceTypeWeightUnitLabelFromValue(v string) (*DeviceTypeWeightUnitLabel, error) {
	ev := DeviceTypeWeightUnitLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeviceTypeWeightUnitLabel: valid values are %v", v, AllowedDeviceTypeWeightUnitLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeviceTypeWeightUnitLabel) IsValid() bool {
	for _, existing := range AllowedDeviceTypeWeightUnitLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeviceType_weight_unit_label value
func (v DeviceTypeWeightUnitLabel) Ptr() *DeviceTypeWeightUnitLabel {
	return &v
}

type NullableDeviceTypeWeightUnitLabel struct {
	value *DeviceTypeWeightUnitLabel
	isSet bool
}

func (v NullableDeviceTypeWeightUnitLabel) Get() *DeviceTypeWeightUnitLabel {
	return v.value
}

func (v *NullableDeviceTypeWeightUnitLabel) Set(val *DeviceTypeWeightUnitLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceTypeWeightUnitLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceTypeWeightUnitLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceTypeWeightUnitLabel(val *DeviceTypeWeightUnitLabel) *NullableDeviceTypeWeightUnitLabel {
	return &NullableDeviceTypeWeightUnitLabel{value: val, isSet: true}
}

func (v NullableDeviceTypeWeightUnitLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceTypeWeightUnitLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

