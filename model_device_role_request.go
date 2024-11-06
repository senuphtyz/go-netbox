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

// checks if the DeviceRoleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceRoleRequest{}

// DeviceRoleRequest Adds support for custom fields and tags.
type DeviceRoleRequest struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
	Color *string `json:"color,omitempty"`
	// Virtual machines may be assigned to this role
	VmRole *bool `json:"vm_role,omitempty"`
	ConfigTemplate NullableBriefConfigTemplateRequest `json:"config_template,omitempty"`
	Description *string `json:"description,omitempty"`
	Tags []NestedTagRequest `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceRoleRequest DeviceRoleRequest

// NewDeviceRoleRequest instantiates a new DeviceRoleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceRoleRequest(name string, slug string) *DeviceRoleRequest {
	this := DeviceRoleRequest{}
	this.Name = name
	this.Slug = slug
	return &this
}

// NewDeviceRoleRequestWithDefaults instantiates a new DeviceRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceRoleRequestWithDefaults() *DeviceRoleRequest {
	this := DeviceRoleRequest{}
	return &this
}

// GetName returns the Name field value
func (o *DeviceRoleRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DeviceRoleRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DeviceRoleRequest) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *DeviceRoleRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *DeviceRoleRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *DeviceRoleRequest) SetSlug(v string) {
	o.Slug = v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *DeviceRoleRequest) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRoleRequest) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *DeviceRoleRequest) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *DeviceRoleRequest) SetColor(v string) {
	o.Color = &v
}

// GetVmRole returns the VmRole field value if set, zero value otherwise.
func (o *DeviceRoleRequest) GetVmRole() bool {
	if o == nil || IsNil(o.VmRole) {
		var ret bool
		return ret
	}
	return *o.VmRole
}

// GetVmRoleOk returns a tuple with the VmRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRoleRequest) GetVmRoleOk() (*bool, bool) {
	if o == nil || IsNil(o.VmRole) {
		return nil, false
	}
	return o.VmRole, true
}

// HasVmRole returns a boolean if a field has been set.
func (o *DeviceRoleRequest) HasVmRole() bool {
	if o != nil && !IsNil(o.VmRole) {
		return true
	}

	return false
}

// SetVmRole gets a reference to the given bool and assigns it to the VmRole field.
func (o *DeviceRoleRequest) SetVmRole(v bool) {
	o.VmRole = &v
}

// GetConfigTemplate returns the ConfigTemplate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceRoleRequest) GetConfigTemplate() BriefConfigTemplateRequest {
	if o == nil || IsNil(o.ConfigTemplate.Get()) {
		var ret BriefConfigTemplateRequest
		return ret
	}
	return *o.ConfigTemplate.Get()
}

// GetConfigTemplateOk returns a tuple with the ConfigTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceRoleRequest) GetConfigTemplateOk() (*BriefConfigTemplateRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigTemplate.Get(), o.ConfigTemplate.IsSet()
}

// HasConfigTemplate returns a boolean if a field has been set.
func (o *DeviceRoleRequest) HasConfigTemplate() bool {
	if o != nil && o.ConfigTemplate.IsSet() {
		return true
	}

	return false
}

// SetConfigTemplate gets a reference to the given NullableBriefConfigTemplateRequest and assigns it to the ConfigTemplate field.
func (o *DeviceRoleRequest) SetConfigTemplate(v BriefConfigTemplateRequest) {
	o.ConfigTemplate.Set(&v)
}
// SetConfigTemplateNil sets the value for ConfigTemplate to be an explicit nil
func (o *DeviceRoleRequest) SetConfigTemplateNil() {
	o.ConfigTemplate.Set(nil)
}

// UnsetConfigTemplate ensures that no value is present for ConfigTemplate, not even an explicit nil
func (o *DeviceRoleRequest) UnsetConfigTemplate() {
	o.ConfigTemplate.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DeviceRoleRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRoleRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DeviceRoleRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DeviceRoleRequest) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DeviceRoleRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRoleRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DeviceRoleRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *DeviceRoleRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *DeviceRoleRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRoleRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *DeviceRoleRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *DeviceRoleRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o DeviceRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceRoleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if !IsNil(o.VmRole) {
		toSerialize["vm_role"] = o.VmRole
	}
	if o.ConfigTemplate.IsSet() {
		toSerialize["config_template"] = o.ConfigTemplate.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceRoleRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"slug",
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

	varDeviceRoleRequest := _DeviceRoleRequest{}

	err = json.Unmarshal(data, &varDeviceRoleRequest)

	if err != nil {
		return err
	}

	*o = DeviceRoleRequest(varDeviceRoleRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "color")
		delete(additionalProperties, "vm_role")
		delete(additionalProperties, "config_template")
		delete(additionalProperties, "description")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceRoleRequest struct {
	value *DeviceRoleRequest
	isSet bool
}

func (v NullableDeviceRoleRequest) Get() *DeviceRoleRequest {
	return v.value
}

func (v *NullableDeviceRoleRequest) Set(val *DeviceRoleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceRoleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceRoleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceRoleRequest(val *DeviceRoleRequest) *NullableDeviceRoleRequest {
	return &NullableDeviceRoleRequest{value: val, isSet: true}
}

func (v NullableDeviceRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceRoleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


