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

// checks if the BriefVirtualMachineRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefVirtualMachineRequest{}

// BriefVirtualMachineRequest Adds support for custom fields and tags.
type BriefVirtualMachineRequest struct {
	Name                 string  `json:"name"`
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BriefVirtualMachineRequest BriefVirtualMachineRequest

// NewBriefVirtualMachineRequest instantiates a new BriefVirtualMachineRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefVirtualMachineRequest(name string) *BriefVirtualMachineRequest {
	this := BriefVirtualMachineRequest{}
	this.Name = name
	return &this
}

// NewBriefVirtualMachineRequestWithDefaults instantiates a new BriefVirtualMachineRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefVirtualMachineRequestWithDefaults() *BriefVirtualMachineRequest {
	this := BriefVirtualMachineRequest{}
	return &this
}

// GetName returns the Name field value
func (o *BriefVirtualMachineRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BriefVirtualMachineRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BriefVirtualMachineRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BriefVirtualMachineRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefVirtualMachineRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BriefVirtualMachineRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BriefVirtualMachineRequest) SetDescription(v string) {
	o.Description = &v
}

func (o BriefVirtualMachineRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefVirtualMachineRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefVirtualMachineRequest) UnmarshalJSON(data []byte) (err error) {
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

	varBriefVirtualMachineRequest := _BriefVirtualMachineRequest{}

	err = json.Unmarshal(data, &varBriefVirtualMachineRequest)

	if err != nil {
		return err
	}

	*o = BriefVirtualMachineRequest(varBriefVirtualMachineRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefVirtualMachineRequest struct {
	value *BriefVirtualMachineRequest
	isSet bool
}

func (v NullableBriefVirtualMachineRequest) Get() *BriefVirtualMachineRequest {
	return v.value
}

func (v *NullableBriefVirtualMachineRequest) Set(val *BriefVirtualMachineRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefVirtualMachineRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefVirtualMachineRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefVirtualMachineRequest(val *BriefVirtualMachineRequest) *NullableBriefVirtualMachineRequest {
	return &NullableBriefVirtualMachineRequest{value: val, isSet: true}
}

func (v NullableBriefVirtualMachineRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefVirtualMachineRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
