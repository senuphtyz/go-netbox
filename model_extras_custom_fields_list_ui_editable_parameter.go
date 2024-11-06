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

// ExtrasCustomFieldsListUiEditableParameter the model 'ExtrasCustomFieldsListUiEditableParameter'
type ExtrasCustomFieldsListUiEditableParameter string

// List of extras_custom_fields_list_ui_editable_parameter
const (
	EXTRASCUSTOMFIELDSLISTUIEDITABLEPARAMETER_HIDDEN ExtrasCustomFieldsListUiEditableParameter = "hidden"
	EXTRASCUSTOMFIELDSLISTUIEDITABLEPARAMETER_NO ExtrasCustomFieldsListUiEditableParameter = "no"
	EXTRASCUSTOMFIELDSLISTUIEDITABLEPARAMETER_YES ExtrasCustomFieldsListUiEditableParameter = "yes"
)

// All allowed values of ExtrasCustomFieldsListUiEditableParameter enum
var AllowedExtrasCustomFieldsListUiEditableParameterEnumValues = []ExtrasCustomFieldsListUiEditableParameter{
	"hidden",
	"no",
	"yes",
}

func (v *ExtrasCustomFieldsListUiEditableParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ExtrasCustomFieldsListUiEditableParameter(value)
	for _, existing := range AllowedExtrasCustomFieldsListUiEditableParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ExtrasCustomFieldsListUiEditableParameter", value)
}

// NewExtrasCustomFieldsListUiEditableParameterFromValue returns a pointer to a valid ExtrasCustomFieldsListUiEditableParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewExtrasCustomFieldsListUiEditableParameterFromValue(v string) (*ExtrasCustomFieldsListUiEditableParameter, error) {
	ev := ExtrasCustomFieldsListUiEditableParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ExtrasCustomFieldsListUiEditableParameter: valid values are %v", v, AllowedExtrasCustomFieldsListUiEditableParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExtrasCustomFieldsListUiEditableParameter) IsValid() bool {
	for _, existing := range AllowedExtrasCustomFieldsListUiEditableParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to extras_custom_fields_list_ui_editable_parameter value
func (v ExtrasCustomFieldsListUiEditableParameter) Ptr() *ExtrasCustomFieldsListUiEditableParameter {
	return &v
}

type NullableExtrasCustomFieldsListUiEditableParameter struct {
	value *ExtrasCustomFieldsListUiEditableParameter
	isSet bool
}

func (v NullableExtrasCustomFieldsListUiEditableParameter) Get() *ExtrasCustomFieldsListUiEditableParameter {
	return v.value
}

func (v *NullableExtrasCustomFieldsListUiEditableParameter) Set(val *ExtrasCustomFieldsListUiEditableParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableExtrasCustomFieldsListUiEditableParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableExtrasCustomFieldsListUiEditableParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtrasCustomFieldsListUiEditableParameter(val *ExtrasCustomFieldsListUiEditableParameter) *NullableExtrasCustomFieldsListUiEditableParameter {
	return &NullableExtrasCustomFieldsListUiEditableParameter{value: val, isSet: true}
}

func (v NullableExtrasCustomFieldsListUiEditableParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtrasCustomFieldsListUiEditableParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

