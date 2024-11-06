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

// PatchedWritableConsolePortRequestSpeed Port speed in bits per second  * `1200` - 1200 bps * `2400` - 2400 bps * `4800` - 4800 bps * `9600` - 9600 bps * `19200` - 19.2 kbps * `38400` - 38.4 kbps * `57600` - 57.6 kbps * `115200` - 115.2 kbps
type PatchedWritableConsolePortRequestSpeed int32

// List of PatchedWritableConsolePortRequest_speed
const (
	PATCHEDWRITABLECONSOLEPORTREQUESTSPEED__1200   PatchedWritableConsolePortRequestSpeed = 1200
	PATCHEDWRITABLECONSOLEPORTREQUESTSPEED__2400   PatchedWritableConsolePortRequestSpeed = 2400
	PATCHEDWRITABLECONSOLEPORTREQUESTSPEED__4800   PatchedWritableConsolePortRequestSpeed = 4800
	PATCHEDWRITABLECONSOLEPORTREQUESTSPEED__9600   PatchedWritableConsolePortRequestSpeed = 9600
	PATCHEDWRITABLECONSOLEPORTREQUESTSPEED__19200  PatchedWritableConsolePortRequestSpeed = 19200
	PATCHEDWRITABLECONSOLEPORTREQUESTSPEED__38400  PatchedWritableConsolePortRequestSpeed = 38400
	PATCHEDWRITABLECONSOLEPORTREQUESTSPEED__57600  PatchedWritableConsolePortRequestSpeed = 57600
	PATCHEDWRITABLECONSOLEPORTREQUESTSPEED__115200 PatchedWritableConsolePortRequestSpeed = 115200
)

// All allowed values of PatchedWritableConsolePortRequestSpeed enum
var AllowedPatchedWritableConsolePortRequestSpeedEnumValues = []PatchedWritableConsolePortRequestSpeed{
	1200,
	2400,
	4800,
	9600,
	19200,
	38400,
	57600,
	115200,
}

func (v *PatchedWritableConsolePortRequestSpeed) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PatchedWritableConsolePortRequestSpeed(value)
	for _, existing := range AllowedPatchedWritableConsolePortRequestSpeedEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PatchedWritableConsolePortRequestSpeed", value)
}

// NewPatchedWritableConsolePortRequestSpeedFromValue returns a pointer to a valid PatchedWritableConsolePortRequestSpeed
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPatchedWritableConsolePortRequestSpeedFromValue(v int32) (*PatchedWritableConsolePortRequestSpeed, error) {
	ev := PatchedWritableConsolePortRequestSpeed(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PatchedWritableConsolePortRequestSpeed: valid values are %v", v, AllowedPatchedWritableConsolePortRequestSpeedEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PatchedWritableConsolePortRequestSpeed) IsValid() bool {
	for _, existing := range AllowedPatchedWritableConsolePortRequestSpeedEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchedWritableConsolePortRequest_speed value
func (v PatchedWritableConsolePortRequestSpeed) Ptr() *PatchedWritableConsolePortRequestSpeed {
	return &v
}

type NullablePatchedWritableConsolePortRequestSpeed struct {
	value *PatchedWritableConsolePortRequestSpeed
	isSet bool
}

func (v NullablePatchedWritableConsolePortRequestSpeed) Get() *PatchedWritableConsolePortRequestSpeed {
	return v.value
}

func (v *NullablePatchedWritableConsolePortRequestSpeed) Set(val *PatchedWritableConsolePortRequestSpeed) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableConsolePortRequestSpeed) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableConsolePortRequestSpeed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableConsolePortRequestSpeed(val *PatchedWritableConsolePortRequestSpeed) *NullablePatchedWritableConsolePortRequestSpeed {
	return &NullablePatchedWritableConsolePortRequestSpeed{value: val, isSet: true}
}

func (v NullablePatchedWritableConsolePortRequestSpeed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableConsolePortRequestSpeed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
