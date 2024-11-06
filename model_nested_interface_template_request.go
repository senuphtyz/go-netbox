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

// checks if the NestedInterfaceTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedInterfaceTemplateRequest{}

// NestedInterfaceTemplateRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedInterfaceTemplateRequest struct {
	// {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name                 string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _NestedInterfaceTemplateRequest NestedInterfaceTemplateRequest

// NewNestedInterfaceTemplateRequest instantiates a new NestedInterfaceTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedInterfaceTemplateRequest(name string) *NestedInterfaceTemplateRequest {
	this := NestedInterfaceTemplateRequest{}
	this.Name = name
	return &this
}

// NewNestedInterfaceTemplateRequestWithDefaults instantiates a new NestedInterfaceTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedInterfaceTemplateRequestWithDefaults() *NestedInterfaceTemplateRequest {
	this := NestedInterfaceTemplateRequest{}
	return &this
}

// GetName returns the Name field value
func (o *NestedInterfaceTemplateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedInterfaceTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedInterfaceTemplateRequest) SetName(v string) {
	o.Name = v
}

func (o NestedInterfaceTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedInterfaceTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedInterfaceTemplateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varNestedInterfaceTemplateRequest := _NestedInterfaceTemplateRequest{}

	err = json.Unmarshal(data, &varNestedInterfaceTemplateRequest)

	if err != nil {
		return err
	}

	*o = NestedInterfaceTemplateRequest(varNestedInterfaceTemplateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedInterfaceTemplateRequest struct {
	value *NestedInterfaceTemplateRequest
	isSet bool
}

func (v NullableNestedInterfaceTemplateRequest) Get() *NestedInterfaceTemplateRequest {
	return v.value
}

func (v *NullableNestedInterfaceTemplateRequest) Set(val *NestedInterfaceTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedInterfaceTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedInterfaceTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedInterfaceTemplateRequest(val *NestedInterfaceTemplateRequest) *NullableNestedInterfaceTemplateRequest {
	return &NullableNestedInterfaceTemplateRequest{value: val, isSet: true}
}

func (v NullableNestedInterfaceTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedInterfaceTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
