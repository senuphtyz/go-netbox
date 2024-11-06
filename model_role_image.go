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

// checks if the RoleImage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleImage{}

// RoleImage struct for RoleImage
type RoleImage struct {
	Role                 string `json:"role"`
	Image                string `json:"image"`
	AdditionalProperties map[string]interface{}
}

type _RoleImage RoleImage

// NewRoleImage instantiates a new RoleImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleImage(role string, image string) *RoleImage {
	this := RoleImage{}
	this.Role = role
	this.Image = image
	return &this
}

// NewRoleImageWithDefaults instantiates a new RoleImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleImageWithDefaults() *RoleImage {
	this := RoleImage{}
	return &this
}

// GetRole returns the Role field value
func (o *RoleImage) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *RoleImage) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *RoleImage) SetRole(v string) {
	o.Role = v
}

// GetImage returns the Image field value
func (o *RoleImage) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *RoleImage) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *RoleImage) SetImage(v string) {
	o.Image = v
}

func (o RoleImage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleImage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["role"] = o.Role
	toSerialize["image"] = o.Image

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoleImage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"role",
		"image",
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

	varRoleImage := _RoleImage{}

	err = json.Unmarshal(data, &varRoleImage)

	if err != nil {
		return err
	}

	*o = RoleImage(varRoleImage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "role")
		delete(additionalProperties, "image")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoleImage struct {
	value *RoleImage
	isSet bool
}

func (v NullableRoleImage) Get() *RoleImage {
	return v.value
}

func (v *NullableRoleImage) Set(val *RoleImage) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleImage) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleImage(val *RoleImage) *NullableRoleImage {
	return &NullableRoleImage{value: val, isSet: true}
}

func (v NullableRoleImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
