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

// DeviceTypeSubdeviceRoleLabel the model 'DeviceTypeSubdeviceRoleLabel'
type DeviceTypeSubdeviceRoleLabel string

// List of DeviceType_subdevice_role_label
const (
	DEVICETYPESUBDEVICEROLELABEL_PARENT DeviceTypeSubdeviceRoleLabel = "Parent"
	DEVICETYPESUBDEVICEROLELABEL_CHILD  DeviceTypeSubdeviceRoleLabel = "Child"
)

// All allowed values of DeviceTypeSubdeviceRoleLabel enum
var AllowedDeviceTypeSubdeviceRoleLabelEnumValues = []DeviceTypeSubdeviceRoleLabel{
	"Parent",
	"Child",
}

func (v *DeviceTypeSubdeviceRoleLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeviceTypeSubdeviceRoleLabel(value)
	for _, existing := range AllowedDeviceTypeSubdeviceRoleLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeviceTypeSubdeviceRoleLabel", value)
}

// NewDeviceTypeSubdeviceRoleLabelFromValue returns a pointer to a valid DeviceTypeSubdeviceRoleLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeviceTypeSubdeviceRoleLabelFromValue(v string) (*DeviceTypeSubdeviceRoleLabel, error) {
	ev := DeviceTypeSubdeviceRoleLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeviceTypeSubdeviceRoleLabel: valid values are %v", v, AllowedDeviceTypeSubdeviceRoleLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeviceTypeSubdeviceRoleLabel) IsValid() bool {
	for _, existing := range AllowedDeviceTypeSubdeviceRoleLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeviceType_subdevice_role_label value
func (v DeviceTypeSubdeviceRoleLabel) Ptr() *DeviceTypeSubdeviceRoleLabel {
	return &v
}

type NullableDeviceTypeSubdeviceRoleLabel struct {
	value *DeviceTypeSubdeviceRoleLabel
	isSet bool
}

func (v NullableDeviceTypeSubdeviceRoleLabel) Get() *DeviceTypeSubdeviceRoleLabel {
	return v.value
}

func (v *NullableDeviceTypeSubdeviceRoleLabel) Set(val *DeviceTypeSubdeviceRoleLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceTypeSubdeviceRoleLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceTypeSubdeviceRoleLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceTypeSubdeviceRoleLabel(val *DeviceTypeSubdeviceRoleLabel) *NullableDeviceTypeSubdeviceRoleLabel {
	return &NullableDeviceTypeSubdeviceRoleLabel{value: val, isSet: true}
}

func (v NullableDeviceTypeSubdeviceRoleLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceTypeSubdeviceRoleLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
