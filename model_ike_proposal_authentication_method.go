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

// checks if the IKEProposalAuthenticationMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IKEProposalAuthenticationMethod{}

// IKEProposalAuthenticationMethod struct for IKEProposalAuthenticationMethod
type IKEProposalAuthenticationMethod struct {
	Value                *IKEProposalAuthenticationMethodValue `json:"value,omitempty"`
	Label                *IKEProposalAuthenticationMethodLabel `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IKEProposalAuthenticationMethod IKEProposalAuthenticationMethod

// NewIKEProposalAuthenticationMethod instantiates a new IKEProposalAuthenticationMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIKEProposalAuthenticationMethod() *IKEProposalAuthenticationMethod {
	this := IKEProposalAuthenticationMethod{}
	return &this
}

// NewIKEProposalAuthenticationMethodWithDefaults instantiates a new IKEProposalAuthenticationMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIKEProposalAuthenticationMethodWithDefaults() *IKEProposalAuthenticationMethod {
	this := IKEProposalAuthenticationMethod{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *IKEProposalAuthenticationMethod) GetValue() IKEProposalAuthenticationMethodValue {
	if o == nil || IsNil(o.Value) {
		var ret IKEProposalAuthenticationMethodValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IKEProposalAuthenticationMethod) GetValueOk() (*IKEProposalAuthenticationMethodValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *IKEProposalAuthenticationMethod) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given IKEProposalAuthenticationMethodValue and assigns it to the Value field.
func (o *IKEProposalAuthenticationMethod) SetValue(v IKEProposalAuthenticationMethodValue) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *IKEProposalAuthenticationMethod) GetLabel() IKEProposalAuthenticationMethodLabel {
	if o == nil || IsNil(o.Label) {
		var ret IKEProposalAuthenticationMethodLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IKEProposalAuthenticationMethod) GetLabelOk() (*IKEProposalAuthenticationMethodLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *IKEProposalAuthenticationMethod) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given IKEProposalAuthenticationMethodLabel and assigns it to the Label field.
func (o *IKEProposalAuthenticationMethod) SetLabel(v IKEProposalAuthenticationMethodLabel) {
	o.Label = &v
}

func (o IKEProposalAuthenticationMethod) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IKEProposalAuthenticationMethod) ToMap() (map[string]interface{}, error) {
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

func (o *IKEProposalAuthenticationMethod) UnmarshalJSON(data []byte) (err error) {
	varIKEProposalAuthenticationMethod := _IKEProposalAuthenticationMethod{}

	err = json.Unmarshal(data, &varIKEProposalAuthenticationMethod)

	if err != nil {
		return err
	}

	*o = IKEProposalAuthenticationMethod(varIKEProposalAuthenticationMethod)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIKEProposalAuthenticationMethod struct {
	value *IKEProposalAuthenticationMethod
	isSet bool
}

func (v NullableIKEProposalAuthenticationMethod) Get() *IKEProposalAuthenticationMethod {
	return v.value
}

func (v *NullableIKEProposalAuthenticationMethod) Set(val *IKEProposalAuthenticationMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableIKEProposalAuthenticationMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableIKEProposalAuthenticationMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIKEProposalAuthenticationMethod(val *IKEProposalAuthenticationMethod) *NullableIKEProposalAuthenticationMethod {
	return &NullableIKEProposalAuthenticationMethod{value: val, isSet: true}
}

func (v NullableIKEProposalAuthenticationMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIKEProposalAuthenticationMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
