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

// PowerFeedTypeLabel the model 'PowerFeedTypeLabel'
type PowerFeedTypeLabel string

// List of PowerFeed_type_label
const (
	POWERFEEDTYPELABEL_PRIMARY PowerFeedTypeLabel = "Primary"
	POWERFEEDTYPELABEL_REDUNDANT PowerFeedTypeLabel = "Redundant"
)

// All allowed values of PowerFeedTypeLabel enum
var AllowedPowerFeedTypeLabelEnumValues = []PowerFeedTypeLabel{
	"Primary",
	"Redundant",
}

func (v *PowerFeedTypeLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PowerFeedTypeLabel(value)
	for _, existing := range AllowedPowerFeedTypeLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PowerFeedTypeLabel", value)
}

// NewPowerFeedTypeLabelFromValue returns a pointer to a valid PowerFeedTypeLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPowerFeedTypeLabelFromValue(v string) (*PowerFeedTypeLabel, error) {
	ev := PowerFeedTypeLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PowerFeedTypeLabel: valid values are %v", v, AllowedPowerFeedTypeLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PowerFeedTypeLabel) IsValid() bool {
	for _, existing := range AllowedPowerFeedTypeLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PowerFeed_type_label value
func (v PowerFeedTypeLabel) Ptr() *PowerFeedTypeLabel {
	return &v
}

type NullablePowerFeedTypeLabel struct {
	value *PowerFeedTypeLabel
	isSet bool
}

func (v NullablePowerFeedTypeLabel) Get() *PowerFeedTypeLabel {
	return v.value
}

func (v *NullablePowerFeedTypeLabel) Set(val *PowerFeedTypeLabel) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerFeedTypeLabel) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerFeedTypeLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerFeedTypeLabel(val *PowerFeedTypeLabel) *NullablePowerFeedTypeLabel {
	return &NullablePowerFeedTypeLabel{value: val, isSet: true}
}

func (v NullablePowerFeedTypeLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerFeedTypeLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

