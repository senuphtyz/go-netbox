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

// PatchedWritablePowerFeedRequestStatus * `offline` - Offline * `active` - Active * `planned` - Planned * `failed` - Failed
type PatchedWritablePowerFeedRequestStatus string

// List of PatchedWritablePowerFeedRequest_status
const (
	PATCHEDWRITABLEPOWERFEEDREQUESTSTATUS_OFFLINE PatchedWritablePowerFeedRequestStatus = "offline"
	PATCHEDWRITABLEPOWERFEEDREQUESTSTATUS_ACTIVE PatchedWritablePowerFeedRequestStatus = "active"
	PATCHEDWRITABLEPOWERFEEDREQUESTSTATUS_PLANNED PatchedWritablePowerFeedRequestStatus = "planned"
	PATCHEDWRITABLEPOWERFEEDREQUESTSTATUS_FAILED PatchedWritablePowerFeedRequestStatus = "failed"
)

// All allowed values of PatchedWritablePowerFeedRequestStatus enum
var AllowedPatchedWritablePowerFeedRequestStatusEnumValues = []PatchedWritablePowerFeedRequestStatus{
	"offline",
	"active",
	"planned",
	"failed",
}

func (v *PatchedWritablePowerFeedRequestStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PatchedWritablePowerFeedRequestStatus(value)
	for _, existing := range AllowedPatchedWritablePowerFeedRequestStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PatchedWritablePowerFeedRequestStatus", value)
}

// NewPatchedWritablePowerFeedRequestStatusFromValue returns a pointer to a valid PatchedWritablePowerFeedRequestStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPatchedWritablePowerFeedRequestStatusFromValue(v string) (*PatchedWritablePowerFeedRequestStatus, error) {
	ev := PatchedWritablePowerFeedRequestStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PatchedWritablePowerFeedRequestStatus: valid values are %v", v, AllowedPatchedWritablePowerFeedRequestStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PatchedWritablePowerFeedRequestStatus) IsValid() bool {
	for _, existing := range AllowedPatchedWritablePowerFeedRequestStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchedWritablePowerFeedRequest_status value
func (v PatchedWritablePowerFeedRequestStatus) Ptr() *PatchedWritablePowerFeedRequestStatus {
	return &v
}

type NullablePatchedWritablePowerFeedRequestStatus struct {
	value *PatchedWritablePowerFeedRequestStatus
	isSet bool
}

func (v NullablePatchedWritablePowerFeedRequestStatus) Get() *PatchedWritablePowerFeedRequestStatus {
	return v.value
}

func (v *NullablePatchedWritablePowerFeedRequestStatus) Set(val *PatchedWritablePowerFeedRequestStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritablePowerFeedRequestStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritablePowerFeedRequestStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritablePowerFeedRequestStatus(val *PatchedWritablePowerFeedRequestStatus) *NullablePatchedWritablePowerFeedRequestStatus {
	return &NullablePatchedWritablePowerFeedRequestStatus{value: val, isSet: true}
}

func (v NullablePatchedWritablePowerFeedRequestStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritablePowerFeedRequestStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

