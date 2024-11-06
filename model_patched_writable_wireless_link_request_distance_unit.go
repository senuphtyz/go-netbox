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

// PatchedWritableWirelessLinkRequestDistanceUnit * `km` - Kilometers * `m` - Meters * `mi` - Miles * `ft` - Feet
type PatchedWritableWirelessLinkRequestDistanceUnit string

// List of PatchedWritableWirelessLinkRequest_distance_unit
const (
	PATCHEDWRITABLEWIRELESSLINKREQUESTDISTANCEUNIT_KM    PatchedWritableWirelessLinkRequestDistanceUnit = "km"
	PATCHEDWRITABLEWIRELESSLINKREQUESTDISTANCEUNIT_M     PatchedWritableWirelessLinkRequestDistanceUnit = "m"
	PATCHEDWRITABLEWIRELESSLINKREQUESTDISTANCEUNIT_MI    PatchedWritableWirelessLinkRequestDistanceUnit = "mi"
	PATCHEDWRITABLEWIRELESSLINKREQUESTDISTANCEUNIT_FT    PatchedWritableWirelessLinkRequestDistanceUnit = "ft"
	PATCHEDWRITABLEWIRELESSLINKREQUESTDISTANCEUNIT_EMPTY PatchedWritableWirelessLinkRequestDistanceUnit = ""
)

// All allowed values of PatchedWritableWirelessLinkRequestDistanceUnit enum
var AllowedPatchedWritableWirelessLinkRequestDistanceUnitEnumValues = []PatchedWritableWirelessLinkRequestDistanceUnit{
	"km",
	"m",
	"mi",
	"ft",
	"",
}

func (v *PatchedWritableWirelessLinkRequestDistanceUnit) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PatchedWritableWirelessLinkRequestDistanceUnit(value)
	for _, existing := range AllowedPatchedWritableWirelessLinkRequestDistanceUnitEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PatchedWritableWirelessLinkRequestDistanceUnit", value)
}

// NewPatchedWritableWirelessLinkRequestDistanceUnitFromValue returns a pointer to a valid PatchedWritableWirelessLinkRequestDistanceUnit
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPatchedWritableWirelessLinkRequestDistanceUnitFromValue(v string) (*PatchedWritableWirelessLinkRequestDistanceUnit, error) {
	ev := PatchedWritableWirelessLinkRequestDistanceUnit(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PatchedWritableWirelessLinkRequestDistanceUnit: valid values are %v", v, AllowedPatchedWritableWirelessLinkRequestDistanceUnitEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PatchedWritableWirelessLinkRequestDistanceUnit) IsValid() bool {
	for _, existing := range AllowedPatchedWritableWirelessLinkRequestDistanceUnitEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchedWritableWirelessLinkRequest_distance_unit value
func (v PatchedWritableWirelessLinkRequestDistanceUnit) Ptr() *PatchedWritableWirelessLinkRequestDistanceUnit {
	return &v
}

type NullablePatchedWritableWirelessLinkRequestDistanceUnit struct {
	value *PatchedWritableWirelessLinkRequestDistanceUnit
	isSet bool
}

func (v NullablePatchedWritableWirelessLinkRequestDistanceUnit) Get() *PatchedWritableWirelessLinkRequestDistanceUnit {
	return v.value
}

func (v *NullablePatchedWritableWirelessLinkRequestDistanceUnit) Set(val *PatchedWritableWirelessLinkRequestDistanceUnit) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableWirelessLinkRequestDistanceUnit) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableWirelessLinkRequestDistanceUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableWirelessLinkRequestDistanceUnit(val *PatchedWritableWirelessLinkRequestDistanceUnit) *NullablePatchedWritableWirelessLinkRequestDistanceUnit {
	return &NullablePatchedWritableWirelessLinkRequestDistanceUnit{value: val, isSet: true}
}

func (v NullablePatchedWritableWirelessLinkRequestDistanceUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableWirelessLinkRequestDistanceUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
