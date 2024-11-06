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

// checks if the DeviceStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceStatus{}

// DeviceStatus struct for DeviceStatus
type DeviceStatus struct {
	Value                *DeviceStatusValue `json:"value,omitempty"`
	Label                *DeviceStatusLabel `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceStatus DeviceStatus

// NewDeviceStatus instantiates a new DeviceStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceStatus() *DeviceStatus {
	this := DeviceStatus{}
	return &this
}

// NewDeviceStatusWithDefaults instantiates a new DeviceStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceStatusWithDefaults() *DeviceStatus {
	this := DeviceStatus{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DeviceStatus) GetValue() DeviceStatusValue {
	if o == nil || IsNil(o.Value) {
		var ret DeviceStatusValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceStatus) GetValueOk() (*DeviceStatusValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DeviceStatus) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given DeviceStatusValue and assigns it to the Value field.
func (o *DeviceStatus) SetValue(v DeviceStatusValue) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *DeviceStatus) GetLabel() DeviceStatusLabel {
	if o == nil || IsNil(o.Label) {
		var ret DeviceStatusLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceStatus) GetLabelOk() (*DeviceStatusLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *DeviceStatus) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given DeviceStatusLabel and assigns it to the Label field.
func (o *DeviceStatus) SetLabel(v DeviceStatusLabel) {
	o.Label = &v
}

func (o DeviceStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceStatus) ToMap() (map[string]interface{}, error) {
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

func (o *DeviceStatus) UnmarshalJSON(data []byte) (err error) {
	varDeviceStatus := _DeviceStatus{}

	err = json.Unmarshal(data, &varDeviceStatus)

	if err != nil {
		return err
	}

	*o = DeviceStatus(varDeviceStatus)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceStatus struct {
	value *DeviceStatus
	isSet bool
}

func (v NullableDeviceStatus) Get() *DeviceStatus {
	return v.value
}

func (v *NullableDeviceStatus) Set(val *DeviceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceStatus(val *DeviceStatus) *NullableDeviceStatus {
	return &NullableDeviceStatus{value: val, isSet: true}
}

func (v NullableDeviceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
