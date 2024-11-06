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

// ConsolePortTypeValue * `de-9` - DE-9 * `db-25` - DB-25 * `rj-11` - RJ-11 * `rj-12` - RJ-12 * `rj-45` - RJ-45 * `mini-din-8` - Mini-DIN 8 * `usb-a` - USB Type A * `usb-b` - USB Type B * `usb-c` - USB Type C * `usb-mini-a` - USB Mini A * `usb-mini-b` - USB Mini B * `usb-micro-a` - USB Micro A * `usb-micro-b` - USB Micro B * `usb-micro-ab` - USB Micro AB * `other` - Other
type ConsolePortTypeValue string

// List of ConsolePort_type_value
const (
	CONSOLEPORTTYPEVALUE_DE_9         ConsolePortTypeValue = "de-9"
	CONSOLEPORTTYPEVALUE_DB_25        ConsolePortTypeValue = "db-25"
	CONSOLEPORTTYPEVALUE_RJ_11        ConsolePortTypeValue = "rj-11"
	CONSOLEPORTTYPEVALUE_RJ_12        ConsolePortTypeValue = "rj-12"
	CONSOLEPORTTYPEVALUE_RJ_45        ConsolePortTypeValue = "rj-45"
	CONSOLEPORTTYPEVALUE_MINI_DIN_8   ConsolePortTypeValue = "mini-din-8"
	CONSOLEPORTTYPEVALUE_USB_A        ConsolePortTypeValue = "usb-a"
	CONSOLEPORTTYPEVALUE_USB_B        ConsolePortTypeValue = "usb-b"
	CONSOLEPORTTYPEVALUE_USB_C        ConsolePortTypeValue = "usb-c"
	CONSOLEPORTTYPEVALUE_USB_MINI_A   ConsolePortTypeValue = "usb-mini-a"
	CONSOLEPORTTYPEVALUE_USB_MINI_B   ConsolePortTypeValue = "usb-mini-b"
	CONSOLEPORTTYPEVALUE_USB_MICRO_A  ConsolePortTypeValue = "usb-micro-a"
	CONSOLEPORTTYPEVALUE_USB_MICRO_B  ConsolePortTypeValue = "usb-micro-b"
	CONSOLEPORTTYPEVALUE_USB_MICRO_AB ConsolePortTypeValue = "usb-micro-ab"
	CONSOLEPORTTYPEVALUE_OTHER        ConsolePortTypeValue = "other"
	CONSOLEPORTTYPEVALUE_EMPTY        ConsolePortTypeValue = ""
)

// All allowed values of ConsolePortTypeValue enum
var AllowedConsolePortTypeValueEnumValues = []ConsolePortTypeValue{
	"de-9",
	"db-25",
	"rj-11",
	"rj-12",
	"rj-45",
	"mini-din-8",
	"usb-a",
	"usb-b",
	"usb-c",
	"usb-mini-a",
	"usb-mini-b",
	"usb-micro-a",
	"usb-micro-b",
	"usb-micro-ab",
	"other",
	"",
}

func (v *ConsolePortTypeValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConsolePortTypeValue(value)
	for _, existing := range AllowedConsolePortTypeValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConsolePortTypeValue", value)
}

// NewConsolePortTypeValueFromValue returns a pointer to a valid ConsolePortTypeValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConsolePortTypeValueFromValue(v string) (*ConsolePortTypeValue, error) {
	ev := ConsolePortTypeValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConsolePortTypeValue: valid values are %v", v, AllowedConsolePortTypeValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConsolePortTypeValue) IsValid() bool {
	for _, existing := range AllowedConsolePortTypeValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConsolePort_type_value value
func (v ConsolePortTypeValue) Ptr() *ConsolePortTypeValue {
	return &v
}

type NullableConsolePortTypeValue struct {
	value *ConsolePortTypeValue
	isSet bool
}

func (v NullableConsolePortTypeValue) Get() *ConsolePortTypeValue {
	return v.value
}

func (v *NullableConsolePortTypeValue) Set(val *ConsolePortTypeValue) {
	v.value = val
	v.isSet = true
}

func (v NullableConsolePortTypeValue) IsSet() bool {
	return v.isSet
}

func (v *NullableConsolePortTypeValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsolePortTypeValue(val *ConsolePortTypeValue) *NullableConsolePortTypeValue {
	return &NullableConsolePortTypeValue{value: val, isSet: true}
}

func (v NullableConsolePortTypeValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsolePortTypeValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
