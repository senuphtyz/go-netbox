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

// checks if the RackOuterUnit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RackOuterUnit{}

// RackOuterUnit struct for RackOuterUnit
type RackOuterUnit struct {
	Value *PatchedWritableRackRequestOuterUnit `json:"value,omitempty"`
	Label *RackOuterUnitLabel `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RackOuterUnit RackOuterUnit

// NewRackOuterUnit instantiates a new RackOuterUnit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRackOuterUnit() *RackOuterUnit {
	this := RackOuterUnit{}
	return &this
}

// NewRackOuterUnitWithDefaults instantiates a new RackOuterUnit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRackOuterUnitWithDefaults() *RackOuterUnit {
	this := RackOuterUnit{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RackOuterUnit) GetValue() PatchedWritableRackRequestOuterUnit {
	if o == nil || IsNil(o.Value) {
		var ret PatchedWritableRackRequestOuterUnit
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackOuterUnit) GetValueOk() (*PatchedWritableRackRequestOuterUnit, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RackOuterUnit) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given PatchedWritableRackRequestOuterUnit and assigns it to the Value field.
func (o *RackOuterUnit) SetValue(v PatchedWritableRackRequestOuterUnit) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *RackOuterUnit) GetLabel() RackOuterUnitLabel {
	if o == nil || IsNil(o.Label) {
		var ret RackOuterUnitLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackOuterUnit) GetLabelOk() (*RackOuterUnitLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *RackOuterUnit) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given RackOuterUnitLabel and assigns it to the Label field.
func (o *RackOuterUnit) SetLabel(v RackOuterUnitLabel) {
	o.Label = &v
}

func (o RackOuterUnit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RackOuterUnit) ToMap() (map[string]interface{}, error) {
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

func (o *RackOuterUnit) UnmarshalJSON(data []byte) (err error) {
	varRackOuterUnit := _RackOuterUnit{}

	err = json.Unmarshal(data, &varRackOuterUnit)

	if err != nil {
		return err
	}

	*o = RackOuterUnit(varRackOuterUnit)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRackOuterUnit struct {
	value *RackOuterUnit
	isSet bool
}

func (v NullableRackOuterUnit) Get() *RackOuterUnit {
	return v.value
}

func (v *NullableRackOuterUnit) Set(val *RackOuterUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableRackOuterUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableRackOuterUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRackOuterUnit(val *RackOuterUnit) *NullableRackOuterUnit {
	return &NullableRackOuterUnit{value: val, isSet: true}
}

func (v NullableRackOuterUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRackOuterUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

