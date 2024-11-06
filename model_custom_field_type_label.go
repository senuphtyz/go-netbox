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

// CustomFieldTypeLabel the model 'CustomFieldTypeLabel'
type CustomFieldTypeLabel string

// List of CustomField_type_label
const (
	CUSTOMFIELDTYPELABEL_TEXT CustomFieldTypeLabel = "Text"
	CUSTOMFIELDTYPELABEL_TEXT__LONG CustomFieldTypeLabel = "Text (long)"
	CUSTOMFIELDTYPELABEL_INTEGER CustomFieldTypeLabel = "Integer"
	CUSTOMFIELDTYPELABEL_DECIMAL CustomFieldTypeLabel = "Decimal"
	CUSTOMFIELDTYPELABEL_BOOLEAN__TRUE_FALSE CustomFieldTypeLabel = "Boolean (true/false)"
	CUSTOMFIELDTYPELABEL_DATE CustomFieldTypeLabel = "Date"
	CUSTOMFIELDTYPELABEL_DATE__TIME CustomFieldTypeLabel = "Date & time"
	CUSTOMFIELDTYPELABEL_URL CustomFieldTypeLabel = "URL"
	CUSTOMFIELDTYPELABEL_JSON CustomFieldTypeLabel = "JSON"
	CUSTOMFIELDTYPELABEL_SELECTION CustomFieldTypeLabel = "Selection"
	CUSTOMFIELDTYPELABEL_MULTIPLE_SELECTION CustomFieldTypeLabel = "Multiple selection"
	CUSTOMFIELDTYPELABEL_OBJECT CustomFieldTypeLabel = "Object"
	CUSTOMFIELDTYPELABEL_MULTIPLE_OBJECTS CustomFieldTypeLabel = "Multiple objects"
)

// All allowed values of CustomFieldTypeLabel enum
var AllowedCustomFieldTypeLabelEnumValues = []CustomFieldTypeLabel{
	"Text",
	"Text (long)",
	"Integer",
	"Decimal",
	"Boolean (true/false)",
	"Date",
	"Date & time",
	"URL",
	"JSON",
	"Selection",
	"Multiple selection",
	"Object",
	"Multiple objects",
}

func (v *CustomFieldTypeLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CustomFieldTypeLabel(value)
	for _, existing := range AllowedCustomFieldTypeLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CustomFieldTypeLabel", value)
}

// NewCustomFieldTypeLabelFromValue returns a pointer to a valid CustomFieldTypeLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCustomFieldTypeLabelFromValue(v string) (*CustomFieldTypeLabel, error) {
	ev := CustomFieldTypeLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CustomFieldTypeLabel: valid values are %v", v, AllowedCustomFieldTypeLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CustomFieldTypeLabel) IsValid() bool {
	for _, existing := range AllowedCustomFieldTypeLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomField_type_label value
func (v CustomFieldTypeLabel) Ptr() *CustomFieldTypeLabel {
	return &v
}

type NullableCustomFieldTypeLabel struct {
	value *CustomFieldTypeLabel
	isSet bool
}

func (v NullableCustomFieldTypeLabel) Get() *CustomFieldTypeLabel {
	return v.value
}

func (v *NullableCustomFieldTypeLabel) Set(val *CustomFieldTypeLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomFieldTypeLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomFieldTypeLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomFieldTypeLabel(val *CustomFieldTypeLabel) *NullableCustomFieldTypeLabel {
	return &NullableCustomFieldTypeLabel{value: val, isSet: true}
}

func (v NullableCustomFieldTypeLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomFieldTypeLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

