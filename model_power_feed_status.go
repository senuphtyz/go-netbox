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

// checks if the PowerFeedStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PowerFeedStatus{}

// PowerFeedStatus struct for PowerFeedStatus
type PowerFeedStatus struct {
	Value                *PatchedWritablePowerFeedRequestStatus `json:"value,omitempty"`
	Label                *PowerFeedStatusLabel                  `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PowerFeedStatus PowerFeedStatus

// NewPowerFeedStatus instantiates a new PowerFeedStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPowerFeedStatus() *PowerFeedStatus {
	this := PowerFeedStatus{}
	return &this
}

// NewPowerFeedStatusWithDefaults instantiates a new PowerFeedStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPowerFeedStatusWithDefaults() *PowerFeedStatus {
	this := PowerFeedStatus{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PowerFeedStatus) GetValue() PatchedWritablePowerFeedRequestStatus {
	if o == nil || IsNil(o.Value) {
		var ret PatchedWritablePowerFeedRequestStatus
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerFeedStatus) GetValueOk() (*PatchedWritablePowerFeedRequestStatus, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PowerFeedStatus) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given PatchedWritablePowerFeedRequestStatus and assigns it to the Value field.
func (o *PowerFeedStatus) SetValue(v PatchedWritablePowerFeedRequestStatus) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PowerFeedStatus) GetLabel() PowerFeedStatusLabel {
	if o == nil || IsNil(o.Label) {
		var ret PowerFeedStatusLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerFeedStatus) GetLabelOk() (*PowerFeedStatusLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PowerFeedStatus) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given PowerFeedStatusLabel and assigns it to the Label field.
func (o *PowerFeedStatus) SetLabel(v PowerFeedStatusLabel) {
	o.Label = &v
}

func (o PowerFeedStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PowerFeedStatus) ToMap() (map[string]interface{}, error) {
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

func (o *PowerFeedStatus) UnmarshalJSON(data []byte) (err error) {
	varPowerFeedStatus := _PowerFeedStatus{}

	err = json.Unmarshal(data, &varPowerFeedStatus)

	if err != nil {
		return err
	}

	*o = PowerFeedStatus(varPowerFeedStatus)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePowerFeedStatus struct {
	value *PowerFeedStatus
	isSet bool
}

func (v NullablePowerFeedStatus) Get() *PowerFeedStatus {
	return v.value
}

func (v *NullablePowerFeedStatus) Set(val *PowerFeedStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerFeedStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerFeedStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerFeedStatus(val *PowerFeedStatus) *NullablePowerFeedStatus {
	return &NullablePowerFeedStatus{value: val, isSet: true}
}

func (v NullablePowerFeedStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerFeedStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
