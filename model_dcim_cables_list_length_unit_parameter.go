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

// DcimCablesListLengthUnitParameter the model 'DcimCablesListLengthUnitParameter'
type DcimCablesListLengthUnitParameter string

// List of dcim_cables_list_length_unit_parameter
const (
	DCIMCABLESLISTLENGTHUNITPARAMETER_CM DcimCablesListLengthUnitParameter = "cm"
	DCIMCABLESLISTLENGTHUNITPARAMETER_FT DcimCablesListLengthUnitParameter = "ft"
	DCIMCABLESLISTLENGTHUNITPARAMETER_IN DcimCablesListLengthUnitParameter = "in"
	DCIMCABLESLISTLENGTHUNITPARAMETER_KM DcimCablesListLengthUnitParameter = "km"
	DCIMCABLESLISTLENGTHUNITPARAMETER_M DcimCablesListLengthUnitParameter = "m"
	DCIMCABLESLISTLENGTHUNITPARAMETER_MI DcimCablesListLengthUnitParameter = "mi"
)

// All allowed values of DcimCablesListLengthUnitParameter enum
var AllowedDcimCablesListLengthUnitParameterEnumValues = []DcimCablesListLengthUnitParameter{
	"cm",
	"ft",
	"in",
	"km",
	"m",
	"mi",
}

func (v *DcimCablesListLengthUnitParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DcimCablesListLengthUnitParameter(value)
	for _, existing := range AllowedDcimCablesListLengthUnitParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DcimCablesListLengthUnitParameter", value)
}

// NewDcimCablesListLengthUnitParameterFromValue returns a pointer to a valid DcimCablesListLengthUnitParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDcimCablesListLengthUnitParameterFromValue(v string) (*DcimCablesListLengthUnitParameter, error) {
	ev := DcimCablesListLengthUnitParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DcimCablesListLengthUnitParameter: valid values are %v", v, AllowedDcimCablesListLengthUnitParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DcimCablesListLengthUnitParameter) IsValid() bool {
	for _, existing := range AllowedDcimCablesListLengthUnitParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dcim_cables_list_length_unit_parameter value
func (v DcimCablesListLengthUnitParameter) Ptr() *DcimCablesListLengthUnitParameter {
	return &v
}

type NullableDcimCablesListLengthUnitParameter struct {
	value *DcimCablesListLengthUnitParameter
	isSet bool
}

func (v NullableDcimCablesListLengthUnitParameter) Get() *DcimCablesListLengthUnitParameter {
	return v.value
}

func (v *NullableDcimCablesListLengthUnitParameter) Set(val *DcimCablesListLengthUnitParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableDcimCablesListLengthUnitParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableDcimCablesListLengthUnitParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDcimCablesListLengthUnitParameter(val *DcimCablesListLengthUnitParameter) *NullableDcimCablesListLengthUnitParameter {
	return &NullableDcimCablesListLengthUnitParameter{value: val, isSet: true}
}

func (v NullableDcimCablesListLengthUnitParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDcimCablesListLengthUnitParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

