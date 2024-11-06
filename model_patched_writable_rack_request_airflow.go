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

// PatchedWritableRackRequestAirflow * `front-to-rear` - Front to rear * `rear-to-front` - Rear to front
type PatchedWritableRackRequestAirflow string

// List of PatchedWritableRackRequest_airflow
const (
	PATCHEDWRITABLERACKREQUESTAIRFLOW_FRONT_TO_REAR PatchedWritableRackRequestAirflow = "front-to-rear"
	PATCHEDWRITABLERACKREQUESTAIRFLOW_REAR_TO_FRONT PatchedWritableRackRequestAirflow = "rear-to-front"
	PATCHEDWRITABLERACKREQUESTAIRFLOW_EMPTY PatchedWritableRackRequestAirflow = ""
)

// All allowed values of PatchedWritableRackRequestAirflow enum
var AllowedPatchedWritableRackRequestAirflowEnumValues = []PatchedWritableRackRequestAirflow{
	"front-to-rear",
	"rear-to-front",
	"",
}

func (v *PatchedWritableRackRequestAirflow) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PatchedWritableRackRequestAirflow(value)
	for _, existing := range AllowedPatchedWritableRackRequestAirflowEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PatchedWritableRackRequestAirflow", value)
}

// NewPatchedWritableRackRequestAirflowFromValue returns a pointer to a valid PatchedWritableRackRequestAirflow
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPatchedWritableRackRequestAirflowFromValue(v string) (*PatchedWritableRackRequestAirflow, error) {
	ev := PatchedWritableRackRequestAirflow(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PatchedWritableRackRequestAirflow: valid values are %v", v, AllowedPatchedWritableRackRequestAirflowEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PatchedWritableRackRequestAirflow) IsValid() bool {
	for _, existing := range AllowedPatchedWritableRackRequestAirflowEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchedWritableRackRequest_airflow value
func (v PatchedWritableRackRequestAirflow) Ptr() *PatchedWritableRackRequestAirflow {
	return &v
}

type NullablePatchedWritableRackRequestAirflow struct {
	value *PatchedWritableRackRequestAirflow
	isSet bool
}

func (v NullablePatchedWritableRackRequestAirflow) Get() *PatchedWritableRackRequestAirflow {
	return v.value
}

func (v *NullablePatchedWritableRackRequestAirflow) Set(val *PatchedWritableRackRequestAirflow) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableRackRequestAirflow) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableRackRequestAirflow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableRackRequestAirflow(val *PatchedWritableRackRequestAirflow) *NullablePatchedWritableRackRequestAirflow {
	return &NullablePatchedWritableRackRequestAirflow{value: val, isSet: true}
}

func (v NullablePatchedWritableRackRequestAirflow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableRackRequestAirflow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

