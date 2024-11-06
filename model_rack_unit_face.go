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

// checks if the RackUnitFace type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RackUnitFace{}

// RackUnitFace struct for RackUnitFace
type RackUnitFace struct {
	Value                *RackUnitFaceValue `json:"value,omitempty"`
	Label                *DeviceFaceLabel   `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RackUnitFace RackUnitFace

// NewRackUnitFace instantiates a new RackUnitFace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRackUnitFace() *RackUnitFace {
	this := RackUnitFace{}
	return &this
}

// NewRackUnitFaceWithDefaults instantiates a new RackUnitFace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRackUnitFaceWithDefaults() *RackUnitFace {
	this := RackUnitFace{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RackUnitFace) GetValue() RackUnitFaceValue {
	if o == nil || IsNil(o.Value) {
		var ret RackUnitFaceValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackUnitFace) GetValueOk() (*RackUnitFaceValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RackUnitFace) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given RackUnitFaceValue and assigns it to the Value field.
func (o *RackUnitFace) SetValue(v RackUnitFaceValue) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *RackUnitFace) GetLabel() DeviceFaceLabel {
	if o == nil || IsNil(o.Label) {
		var ret DeviceFaceLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackUnitFace) GetLabelOk() (*DeviceFaceLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *RackUnitFace) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given DeviceFaceLabel and assigns it to the Label field.
func (o *RackUnitFace) SetLabel(v DeviceFaceLabel) {
	o.Label = &v
}

func (o RackUnitFace) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RackUnitFace) ToMap() (map[string]interface{}, error) {
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

func (o *RackUnitFace) UnmarshalJSON(data []byte) (err error) {
	varRackUnitFace := _RackUnitFace{}

	err = json.Unmarshal(data, &varRackUnitFace)

	if err != nil {
		return err
	}

	*o = RackUnitFace(varRackUnitFace)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRackUnitFace struct {
	value *RackUnitFace
	isSet bool
}

func (v NullableRackUnitFace) Get() *RackUnitFace {
	return v.value
}

func (v *NullableRackUnitFace) Set(val *RackUnitFace) {
	v.value = val
	v.isSet = true
}

func (v NullableRackUnitFace) IsSet() bool {
	return v.isSet
}

func (v *NullableRackUnitFace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRackUnitFace(val *RackUnitFace) *NullableRackUnitFace {
	return &NullableRackUnitFace{value: val, isSet: true}
}

func (v NullableRackUnitFace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRackUnitFace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
