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

// checks if the BriefIKEPolicyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefIKEPolicyRequest{}

// BriefIKEPolicyRequest Adds support for custom fields and tags.
type BriefIKEPolicyRequest struct {
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BriefIKEPolicyRequest BriefIKEPolicyRequest

// NewBriefIKEPolicyRequest instantiates a new BriefIKEPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefIKEPolicyRequest(name string) *BriefIKEPolicyRequest {
	this := BriefIKEPolicyRequest{}
	this.Name = name
	return &this
}

// NewBriefIKEPolicyRequestWithDefaults instantiates a new BriefIKEPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefIKEPolicyRequestWithDefaults() *BriefIKEPolicyRequest {
	this := BriefIKEPolicyRequest{}
	return &this
}

// GetName returns the Name field value
func (o *BriefIKEPolicyRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BriefIKEPolicyRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BriefIKEPolicyRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BriefIKEPolicyRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefIKEPolicyRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BriefIKEPolicyRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BriefIKEPolicyRequest) SetDescription(v string) {
	o.Description = &v
}

func (o BriefIKEPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefIKEPolicyRequest) ToMap() (map[string]interface{}, error) {
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

func (o *BriefIKEPolicyRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varBriefIKEPolicyRequest := _BriefIKEPolicyRequest{}

	err = json.Unmarshal(data, &varBriefIKEPolicyRequest)

	if err != nil {
		return err
	}

	*o = BriefIKEPolicyRequest(varBriefIKEPolicyRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefIKEPolicyRequest struct {
	value *BriefIKEPolicyRequest
	isSet bool
}

func (v NullableBriefIKEPolicyRequest) Get() *BriefIKEPolicyRequest {
	return v.value
}

func (v *NullableBriefIKEPolicyRequest) Set(val *BriefIKEPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefIKEPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefIKEPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefIKEPolicyRequest(val *BriefIKEPolicyRequest) *NullableBriefIKEPolicyRequest {
	return &NullableBriefIKEPolicyRequest{value: val, isSet: true}
}

func (v NullableBriefIKEPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefIKEPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


