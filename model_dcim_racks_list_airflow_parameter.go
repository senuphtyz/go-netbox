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

// DcimRacksListAirflowParameter the model 'DcimRacksListAirflowParameter'
type DcimRacksListAirflowParameter string

// List of dcim_racks_list_airflow_parameter
const (
	DCIMRACKSLISTAIRFLOWPARAMETER_FRONT_TO_REAR DcimRacksListAirflowParameter = "front-to-rear"
	DCIMRACKSLISTAIRFLOWPARAMETER_REAR_TO_FRONT DcimRacksListAirflowParameter = "rear-to-front"
)

// All allowed values of DcimRacksListAirflowParameter enum
var AllowedDcimRacksListAirflowParameterEnumValues = []DcimRacksListAirflowParameter{
	"front-to-rear",
	"rear-to-front",
}

func (v *DcimRacksListAirflowParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DcimRacksListAirflowParameter(value)
	for _, existing := range AllowedDcimRacksListAirflowParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DcimRacksListAirflowParameter", value)
}

// NewDcimRacksListAirflowParameterFromValue returns a pointer to a valid DcimRacksListAirflowParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDcimRacksListAirflowParameterFromValue(v string) (*DcimRacksListAirflowParameter, error) {
	ev := DcimRacksListAirflowParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DcimRacksListAirflowParameter: valid values are %v", v, AllowedDcimRacksListAirflowParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DcimRacksListAirflowParameter) IsValid() bool {
	for _, existing := range AllowedDcimRacksListAirflowParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dcim_racks_list_airflow_parameter value
func (v DcimRacksListAirflowParameter) Ptr() *DcimRacksListAirflowParameter {
	return &v
}

type NullableDcimRacksListAirflowParameter struct {
	value *DcimRacksListAirflowParameter
	isSet bool
}

func (v NullableDcimRacksListAirflowParameter) Get() *DcimRacksListAirflowParameter {
	return v.value
}

func (v *NullableDcimRacksListAirflowParameter) Set(val *DcimRacksListAirflowParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableDcimRacksListAirflowParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableDcimRacksListAirflowParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDcimRacksListAirflowParameter(val *DcimRacksListAirflowParameter) *NullableDcimRacksListAirflowParameter {
	return &NullableDcimRacksListAirflowParameter{value: val, isSet: true}
}

func (v NullableDcimRacksListAirflowParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDcimRacksListAirflowParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
