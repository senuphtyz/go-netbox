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

// LocationStatusValue * `planned` - Planned * `staging` - Staging * `active` - Active * `decommissioning` - Decommissioning * `retired` - Retired
type LocationStatusValue string

// List of Location_status_value
const (
	LOCATIONSTATUSVALUE_PLANNED LocationStatusValue = "planned"
	LOCATIONSTATUSVALUE_STAGING LocationStatusValue = "staging"
	LOCATIONSTATUSVALUE_ACTIVE LocationStatusValue = "active"
	LOCATIONSTATUSVALUE_DECOMMISSIONING LocationStatusValue = "decommissioning"
	LOCATIONSTATUSVALUE_RETIRED LocationStatusValue = "retired"
)

// All allowed values of LocationStatusValue enum
var AllowedLocationStatusValueEnumValues = []LocationStatusValue{
	"planned",
	"staging",
	"active",
	"decommissioning",
	"retired",
}

func (v *LocationStatusValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LocationStatusValue(value)
	for _, existing := range AllowedLocationStatusValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LocationStatusValue", value)
}

// NewLocationStatusValueFromValue returns a pointer to a valid LocationStatusValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLocationStatusValueFromValue(v string) (*LocationStatusValue, error) {
	ev := LocationStatusValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LocationStatusValue: valid values are %v", v, AllowedLocationStatusValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LocationStatusValue) IsValid() bool {
	for _, existing := range AllowedLocationStatusValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Location_status_value value
func (v LocationStatusValue) Ptr() *LocationStatusValue {
	return &v
}

type NullableLocationStatusValue struct {
	value *LocationStatusValue
	isSet bool
}

func (v NullableLocationStatusValue) Get() *LocationStatusValue {
	return v.value
}

func (v *NullableLocationStatusValue) Set(val *LocationStatusValue) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationStatusValue) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationStatusValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationStatusValue(val *LocationStatusValue) *NullableLocationStatusValue {
	return &NullableLocationStatusValue{value: val, isSet: true}
}

func (v NullableLocationStatusValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationStatusValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

