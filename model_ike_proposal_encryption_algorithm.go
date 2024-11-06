/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.1.4 (4.1)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the IKEProposalEncryptionAlgorithm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IKEProposalEncryptionAlgorithm{}

// IKEProposalEncryptionAlgorithm struct for IKEProposalEncryptionAlgorithm
type IKEProposalEncryptionAlgorithm struct {
	Value                *IKEProposalEncryptionAlgorithmValue `json:"value,omitempty"`
	Label                *IKEProposalEncryptionAlgorithmLabel `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IKEProposalEncryptionAlgorithm IKEProposalEncryptionAlgorithm

// NewIKEProposalEncryptionAlgorithm instantiates a new IKEProposalEncryptionAlgorithm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIKEProposalEncryptionAlgorithm() *IKEProposalEncryptionAlgorithm {
	this := IKEProposalEncryptionAlgorithm{}
	return &this
}

// NewIKEProposalEncryptionAlgorithmWithDefaults instantiates a new IKEProposalEncryptionAlgorithm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIKEProposalEncryptionAlgorithmWithDefaults() *IKEProposalEncryptionAlgorithm {
	this := IKEProposalEncryptionAlgorithm{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *IKEProposalEncryptionAlgorithm) GetValue() IKEProposalEncryptionAlgorithmValue {
	if o == nil || IsNil(o.Value) {
		var ret IKEProposalEncryptionAlgorithmValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IKEProposalEncryptionAlgorithm) GetValueOk() (*IKEProposalEncryptionAlgorithmValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *IKEProposalEncryptionAlgorithm) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given IKEProposalEncryptionAlgorithmValue and assigns it to the Value field.
func (o *IKEProposalEncryptionAlgorithm) SetValue(v IKEProposalEncryptionAlgorithmValue) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *IKEProposalEncryptionAlgorithm) GetLabel() IKEProposalEncryptionAlgorithmLabel {
	if o == nil || IsNil(o.Label) {
		var ret IKEProposalEncryptionAlgorithmLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IKEProposalEncryptionAlgorithm) GetLabelOk() (*IKEProposalEncryptionAlgorithmLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *IKEProposalEncryptionAlgorithm) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given IKEProposalEncryptionAlgorithmLabel and assigns it to the Label field.
func (o *IKEProposalEncryptionAlgorithm) SetLabel(v IKEProposalEncryptionAlgorithmLabel) {
	o.Label = &v
}

func (o IKEProposalEncryptionAlgorithm) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IKEProposalEncryptionAlgorithm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IKEProposalEncryptionAlgorithm) UnmarshalJSON(data []byte) (err error) {
	varIKEProposalEncryptionAlgorithm := _IKEProposalEncryptionAlgorithm{}

	err = json.Unmarshal(data, &varIKEProposalEncryptionAlgorithm)

	if err != nil {
		return err
	}

	*o = IKEProposalEncryptionAlgorithm(varIKEProposalEncryptionAlgorithm)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIKEProposalEncryptionAlgorithm struct {
	value *IKEProposalEncryptionAlgorithm
	isSet bool
}

func (v NullableIKEProposalEncryptionAlgorithm) Get() *IKEProposalEncryptionAlgorithm {
	return v.value
}

func (v *NullableIKEProposalEncryptionAlgorithm) Set(val *IKEProposalEncryptionAlgorithm) {
	v.value = val
	v.isSet = true
}

func (v NullableIKEProposalEncryptionAlgorithm) IsSet() bool {
	return v.isSet
}

func (v *NullableIKEProposalEncryptionAlgorithm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIKEProposalEncryptionAlgorithm(val *IKEProposalEncryptionAlgorithm) *NullableIKEProposalEncryptionAlgorithm {
	return &NullableIKEProposalEncryptionAlgorithm{value: val, isSet: true}
}

func (v NullableIKEProposalEncryptionAlgorithm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIKEProposalEncryptionAlgorithm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
