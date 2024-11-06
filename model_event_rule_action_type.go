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

// checks if the EventRuleActionType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventRuleActionType{}

// EventRuleActionType struct for EventRuleActionType
type EventRuleActionType struct {
	Value                *EventRuleActionTypeValue `json:"value,omitempty"`
	Label                *EventRuleActionTypeLabel `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EventRuleActionType EventRuleActionType

// NewEventRuleActionType instantiates a new EventRuleActionType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventRuleActionType() *EventRuleActionType {
	this := EventRuleActionType{}
	return &this
}

// NewEventRuleActionTypeWithDefaults instantiates a new EventRuleActionType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventRuleActionTypeWithDefaults() *EventRuleActionType {
	this := EventRuleActionType{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *EventRuleActionType) GetValue() EventRuleActionTypeValue {
	if o == nil || IsNil(o.Value) {
		var ret EventRuleActionTypeValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventRuleActionType) GetValueOk() (*EventRuleActionTypeValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *EventRuleActionType) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given EventRuleActionTypeValue and assigns it to the Value field.
func (o *EventRuleActionType) SetValue(v EventRuleActionTypeValue) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *EventRuleActionType) GetLabel() EventRuleActionTypeLabel {
	if o == nil || IsNil(o.Label) {
		var ret EventRuleActionTypeLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventRuleActionType) GetLabelOk() (*EventRuleActionTypeLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *EventRuleActionType) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given EventRuleActionTypeLabel and assigns it to the Label field.
func (o *EventRuleActionType) SetLabel(v EventRuleActionTypeLabel) {
	o.Label = &v
}

func (o EventRuleActionType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventRuleActionType) ToMap() (map[string]interface{}, error) {
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

func (o *EventRuleActionType) UnmarshalJSON(data []byte) (err error) {
	varEventRuleActionType := _EventRuleActionType{}

	err = json.Unmarshal(data, &varEventRuleActionType)

	if err != nil {
		return err
	}

	*o = EventRuleActionType(varEventRuleActionType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEventRuleActionType struct {
	value *EventRuleActionType
	isSet bool
}

func (v NullableEventRuleActionType) Get() *EventRuleActionType {
	return v.value
}

func (v *NullableEventRuleActionType) Set(val *EventRuleActionType) {
	v.value = val
	v.isSet = true
}

func (v NullableEventRuleActionType) IsSet() bool {
	return v.isSet
}

func (v *NullableEventRuleActionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventRuleActionType(val *EventRuleActionType) *NullableEventRuleActionType {
	return &NullableEventRuleActionType{value: val, isSet: true}
}

func (v NullableEventRuleActionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventRuleActionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
