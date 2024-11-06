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

// PowerOutletFeedLegLabel the model 'PowerOutletFeedLegLabel'
type PowerOutletFeedLegLabel string

// List of PowerOutlet_feed_leg_label
const (
	POWEROUTLETFEEDLEGLABEL_A PowerOutletFeedLegLabel = "A"
	POWEROUTLETFEEDLEGLABEL_B PowerOutletFeedLegLabel = "B"
	POWEROUTLETFEEDLEGLABEL_C PowerOutletFeedLegLabel = "C"
)

// All allowed values of PowerOutletFeedLegLabel enum
var AllowedPowerOutletFeedLegLabelEnumValues = []PowerOutletFeedLegLabel{
	"A",
	"B",
	"C",
}

func (v *PowerOutletFeedLegLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PowerOutletFeedLegLabel(value)
	for _, existing := range AllowedPowerOutletFeedLegLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PowerOutletFeedLegLabel", value)
}

// NewPowerOutletFeedLegLabelFromValue returns a pointer to a valid PowerOutletFeedLegLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPowerOutletFeedLegLabelFromValue(v string) (*PowerOutletFeedLegLabel, error) {
	ev := PowerOutletFeedLegLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PowerOutletFeedLegLabel: valid values are %v", v, AllowedPowerOutletFeedLegLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PowerOutletFeedLegLabel) IsValid() bool {
	for _, existing := range AllowedPowerOutletFeedLegLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PowerOutlet_feed_leg_label value
func (v PowerOutletFeedLegLabel) Ptr() *PowerOutletFeedLegLabel {
	return &v
}

type NullablePowerOutletFeedLegLabel struct {
	value *PowerOutletFeedLegLabel
	isSet bool
}

func (v NullablePowerOutletFeedLegLabel) Get() *PowerOutletFeedLegLabel {
	return v.value
}

func (v *NullablePowerOutletFeedLegLabel) Set(val *PowerOutletFeedLegLabel) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerOutletFeedLegLabel) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerOutletFeedLegLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerOutletFeedLegLabel(val *PowerOutletFeedLegLabel) *NullablePowerOutletFeedLegLabel {
	return &NullablePowerOutletFeedLegLabel{value: val, isSet: true}
}

func (v NullablePowerOutletFeedLegLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerOutletFeedLegLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

