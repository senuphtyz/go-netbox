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

// PatchedWritableIPSecPolicyRequestPfsGroup Diffie-Hellman group for Perfect Forward Secrecy  * `1` - Group 1 * `2` - Group 2 * `5` - Group 5 * `14` - Group 14 * `15` - Group 15 * `16` - Group 16 * `17` - Group 17 * `18` - Group 18 * `19` - Group 19 * `20` - Group 20 * `21` - Group 21 * `22` - Group 22 * `23` - Group 23 * `24` - Group 24 * `25` - Group 25 * `26` - Group 26 * `27` - Group 27 * `28` - Group 28 * `29` - Group 29 * `30` - Group 30 * `31` - Group 31 * `32` - Group 32 * `33` - Group 33 * `34` - Group 34
type PatchedWritableIPSecPolicyRequestPfsGroup int32

// List of PatchedWritableIPSecPolicyRequest_pfs_group
const (
	PATCHEDWRITABLEIPSECPOLICYREQUESTPFSGROUP__1 PatchedWritableIPSecPolicyRequestPfsGroup = 1
	PATCHEDWRITABLEIPSECPOLICYREQUESTPFSGROUP__2 PatchedWritableIPSecPolicyRequestPfsGroup = 2
	PATCHEDWRITABLEIPSECPOLICYREQUESTPFSGROUP__5 PatchedWritableIPSecPolicyRequestPfsGroup = 5
	PATCHEDWRITABLEIPSECPOLICYREQUESTPFSGROUP__14 PatchedWritableIPSecPolicyRequestPfsGroup = 14
	PATCHEDWRITABLEIPSECPOLICYREQUESTPFSGROUP__15 PatchedWritableIPSecPolicyRequestPfsGroup = 15
	PATCHEDWRITABLEIPSECPOLICYREQUESTPFSGROUP__16 PatchedWritableIPSecPolicyRequestPfsGroup = 16
	PATCHEDWRITABLEIPSECPOLICYREQUESTPFSGROUP__17 PatchedWritableIPSecPolicyRequestPfsGroup = 17
	PATCHEDWRITABLEIPSECPOLICYREQUESTPFSGROUP__18 PatchedWritableIPSecPolicyRequestPfsGroup = 18
	PATCHEDWRITABLEIPSECPOLICYREQUESTPFSGROUP__19 PatchedWritableIPSecPolicyRequestPfsGroup = 19
	PATCHEDWRITABLEIPSECPOLICYREQUESTPFSGROUP__20 PatchedWritableIPSecPolicyRequestPfsGroup = 20
	PATCHEDWRITABLEIPSECPOLICYREQUESTPFSGROUP__21 PatchedWritableIPSecPolicyRequestPfsGroup = 21
	PATCHEDWRITABLEIPSECPOLICYREQUESTPFSGROUP__22 PatchedWritableIPSecPolicyRequestPfsGroup = 22
	PATCHEDWRITABLEIPSECPOLICYREQUESTPFSGROUP__23 PatchedWritableIPSecPolicyRequestPfsGroup = 23
	PATCHEDWRITABLEIPSECPOLICYREQUESTPFSGROUP__24 PatchedWritableIPSecPolicyRequestPfsGroup = 24
	PATCHEDWRITABLEIPSECPOLICYREQUESTPFSGROUP__25 PatchedWritableIPSecPolicyRequestPfsGroup = 25
	PATCHEDWRITABLEIPSECPOLICYREQUESTPFSGROUP__26 PatchedWritableIPSecPolicyRequestPfsGroup = 26
	PATCHEDWRITABLEIPSECPOLICYREQUESTPFSGROUP__27 PatchedWritableIPSecPolicyRequestPfsGroup = 27
	PATCHEDWRITABLEIPSECPOLICYREQUESTPFSGROUP__28 PatchedWritableIPSecPolicyRequestPfsGroup = 28
	PATCHEDWRITABLEIPSECPOLICYREQUESTPFSGROUP__29 PatchedWritableIPSecPolicyRequestPfsGroup = 29
	PATCHEDWRITABLEIPSECPOLICYREQUESTPFSGROUP__30 PatchedWritableIPSecPolicyRequestPfsGroup = 30
	PATCHEDWRITABLEIPSECPOLICYREQUESTPFSGROUP__31 PatchedWritableIPSecPolicyRequestPfsGroup = 31
	PATCHEDWRITABLEIPSECPOLICYREQUESTPFSGROUP__32 PatchedWritableIPSecPolicyRequestPfsGroup = 32
	PATCHEDWRITABLEIPSECPOLICYREQUESTPFSGROUP__33 PatchedWritableIPSecPolicyRequestPfsGroup = 33
	PATCHEDWRITABLEIPSECPOLICYREQUESTPFSGROUP__34 PatchedWritableIPSecPolicyRequestPfsGroup = 34
)

// All allowed values of PatchedWritableIPSecPolicyRequestPfsGroup enum
var AllowedPatchedWritableIPSecPolicyRequestPfsGroupEnumValues = []PatchedWritableIPSecPolicyRequestPfsGroup{
	1,
	2,
	5,
	14,
	15,
	16,
	17,
	18,
	19,
	20,
	21,
	22,
	23,
	24,
	25,
	26,
	27,
	28,
	29,
	30,
	31,
	32,
	33,
	34,
}

func (v *PatchedWritableIPSecPolicyRequestPfsGroup) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PatchedWritableIPSecPolicyRequestPfsGroup(value)
	for _, existing := range AllowedPatchedWritableIPSecPolicyRequestPfsGroupEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PatchedWritableIPSecPolicyRequestPfsGroup", value)
}

// NewPatchedWritableIPSecPolicyRequestPfsGroupFromValue returns a pointer to a valid PatchedWritableIPSecPolicyRequestPfsGroup
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPatchedWritableIPSecPolicyRequestPfsGroupFromValue(v int32) (*PatchedWritableIPSecPolicyRequestPfsGroup, error) {
	ev := PatchedWritableIPSecPolicyRequestPfsGroup(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PatchedWritableIPSecPolicyRequestPfsGroup: valid values are %v", v, AllowedPatchedWritableIPSecPolicyRequestPfsGroupEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PatchedWritableIPSecPolicyRequestPfsGroup) IsValid() bool {
	for _, existing := range AllowedPatchedWritableIPSecPolicyRequestPfsGroupEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchedWritableIPSecPolicyRequest_pfs_group value
func (v PatchedWritableIPSecPolicyRequestPfsGroup) Ptr() *PatchedWritableIPSecPolicyRequestPfsGroup {
	return &v
}

type NullablePatchedWritableIPSecPolicyRequestPfsGroup struct {
	value *PatchedWritableIPSecPolicyRequestPfsGroup
	isSet bool
}

func (v NullablePatchedWritableIPSecPolicyRequestPfsGroup) Get() *PatchedWritableIPSecPolicyRequestPfsGroup {
	return v.value
}

func (v *NullablePatchedWritableIPSecPolicyRequestPfsGroup) Set(val *PatchedWritableIPSecPolicyRequestPfsGroup) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableIPSecPolicyRequestPfsGroup) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableIPSecPolicyRequestPfsGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableIPSecPolicyRequestPfsGroup(val *PatchedWritableIPSecPolicyRequestPfsGroup) *NullablePatchedWritableIPSecPolicyRequestPfsGroup {
	return &NullablePatchedWritableIPSecPolicyRequestPfsGroup{value: val, isSet: true}
}

func (v NullablePatchedWritableIPSecPolicyRequestPfsGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableIPSecPolicyRequestPfsGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

