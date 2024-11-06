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

// checks if the TopologyDummy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TopologyDummy{}

// TopologyDummy struct for TopologyDummy
type TopologyDummy struct {
	Id                   int32          `json:"id"`
	Name                 NullableString `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TopologyDummy TopologyDummy

// NewTopologyDummy instantiates a new TopologyDummy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopologyDummy(id int32) *TopologyDummy {
	this := TopologyDummy{}
	this.Id = id
	return &this
}

// NewTopologyDummyWithDefaults instantiates a new TopologyDummy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopologyDummyWithDefaults() *TopologyDummy {
	this := TopologyDummy{}
	return &this
}

// GetId returns the Id field value
func (o *TopologyDummy) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TopologyDummy) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TopologyDummy) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TopologyDummy) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TopologyDummy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *TopologyDummy) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *TopologyDummy) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *TopologyDummy) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *TopologyDummy) UnsetName() {
	o.Name.Unset()
}

func (o TopologyDummy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TopologyDummy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TopologyDummy) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTopologyDummy := _TopologyDummy{}

	err = json.Unmarshal(data, &varTopologyDummy)

	if err != nil {
		return err
	}

	*o = TopologyDummy(varTopologyDummy)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTopologyDummy struct {
	value *TopologyDummy
	isSet bool
}

func (v NullableTopologyDummy) Get() *TopologyDummy {
	return v.value
}

func (v *NullableTopologyDummy) Set(val *TopologyDummy) {
	v.value = val
	v.isSet = true
}

func (v NullableTopologyDummy) IsSet() bool {
	return v.isSet
}

func (v *NullableTopologyDummy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopologyDummy(val *TopologyDummy) *NullableTopologyDummy {
	return &NullableTopologyDummy{value: val, isSet: true}
}

func (v NullableTopologyDummy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopologyDummy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
