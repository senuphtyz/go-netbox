/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0.3 (4.0)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the WirelessLinkStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WirelessLinkStatus{}

// WirelessLinkStatus struct for WirelessLinkStatus
type WirelessLinkStatus struct {
	Value                *PatchedWritableCableRequestStatus `json:"value,omitempty"`
	Label                *WirelessLinkStatusLabel           `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WirelessLinkStatus WirelessLinkStatus

// NewWirelessLinkStatus instantiates a new WirelessLinkStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWirelessLinkStatus() *WirelessLinkStatus {
	this := WirelessLinkStatus{}
	return &this
}

// NewWirelessLinkStatusWithDefaults instantiates a new WirelessLinkStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWirelessLinkStatusWithDefaults() *WirelessLinkStatus {
	this := WirelessLinkStatus{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *WirelessLinkStatus) GetValue() PatchedWritableCableRequestStatus {
	if o == nil || IsNil(o.Value) {
		var ret PatchedWritableCableRequestStatus
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WirelessLinkStatus) GetValueOk() (*PatchedWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *WirelessLinkStatus) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given PatchedWritableCableRequestStatus and assigns it to the Value field.
func (o *WirelessLinkStatus) SetValue(v PatchedWritableCableRequestStatus) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WirelessLinkStatus) GetLabel() WirelessLinkStatusLabel {
	if o == nil || IsNil(o.Label) {
		var ret WirelessLinkStatusLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WirelessLinkStatus) GetLabelOk() (*WirelessLinkStatusLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WirelessLinkStatus) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given WirelessLinkStatusLabel and assigns it to the Label field.
func (o *WirelessLinkStatus) SetLabel(v WirelessLinkStatusLabel) {
	o.Label = &v
}

func (o WirelessLinkStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WirelessLinkStatus) ToMap() (map[string]interface{}, error) {
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

func (o *WirelessLinkStatus) UnmarshalJSON(data []byte) (err error) {
	varWirelessLinkStatus := _WirelessLinkStatus{}

	err = json.Unmarshal(data, &varWirelessLinkStatus)

	if err != nil {
		return err
	}

	*o = WirelessLinkStatus(varWirelessLinkStatus)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWirelessLinkStatus struct {
	value *WirelessLinkStatus
	isSet bool
}

func (v NullableWirelessLinkStatus) Get() *WirelessLinkStatus {
	return v.value
}

func (v *NullableWirelessLinkStatus) Set(val *WirelessLinkStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableWirelessLinkStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableWirelessLinkStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWirelessLinkStatus(val *WirelessLinkStatus) *NullableWirelessLinkStatus {
	return &NullableWirelessLinkStatus{value: val, isSet: true}
}

func (v NullableWirelessLinkStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWirelessLinkStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}