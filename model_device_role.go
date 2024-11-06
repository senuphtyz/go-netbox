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
	"time"
)

// checks if the DeviceRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceRole{}

// DeviceRole Adds support for custom fields and tags.
type DeviceRole struct {
	Id         int32   `json:"id"`
	Url        string  `json:"url"`
	DisplayUrl string  `json:"display_url"`
	Display    string  `json:"display"`
	Name       string  `json:"name"`
	Slug       string  `json:"slug" validate:"regexp=^[-a-zA-Z0-9_]+$"`
	Color      *string `json:"color,omitempty" validate:"regexp=^[0-9a-f]{6}$"`
	// Virtual machines may be assigned to this role
	VmRole               *bool                       `json:"vm_role,omitempty"`
	ConfigTemplate       NullableBriefConfigTemplate `json:"config_template,omitempty"`
	Description          *string                     `json:"description,omitempty"`
	Tags                 []NestedTag                 `json:"tags,omitempty"`
	CustomFields         map[string]interface{}      `json:"custom_fields,omitempty"`
	Created              NullableTime                `json:"created"`
	LastUpdated          NullableTime                `json:"last_updated"`
	DeviceCount          int64                       `json:"device_count"`
	VirtualmachineCount  int64                       `json:"virtualmachine_count"`
	AdditionalProperties map[string]interface{}
}

type _DeviceRole DeviceRole

// NewDeviceRole instantiates a new DeviceRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceRole(id int32, url string, displayUrl string, display string, name string, slug string, created NullableTime, lastUpdated NullableTime, deviceCount int64, virtualmachineCount int64) *DeviceRole {
	this := DeviceRole{}
	this.Id = id
	this.Url = url
	this.DisplayUrl = displayUrl
	this.Display = display
	this.Name = name
	this.Slug = slug
	this.Created = created
	this.LastUpdated = lastUpdated
	this.DeviceCount = deviceCount
	this.VirtualmachineCount = virtualmachineCount
	return &this
}

// NewDeviceRoleWithDefaults instantiates a new DeviceRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceRoleWithDefaults() *DeviceRole {
	this := DeviceRole{}
	return &this
}

// GetId returns the Id field value
func (o *DeviceRole) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeviceRole) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeviceRole) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *DeviceRole) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *DeviceRole) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *DeviceRole) SetUrl(v string) {
	o.Url = v
}

// GetDisplayUrl returns the DisplayUrl field value
func (o *DeviceRole) GetDisplayUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayUrl
}

// GetDisplayUrlOk returns a tuple with the DisplayUrl field value
// and a boolean to check if the value has been set.
func (o *DeviceRole) GetDisplayUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayUrl, true
}

// SetDisplayUrl sets field value
func (o *DeviceRole) SetDisplayUrl(v string) {
	o.DisplayUrl = v
}

// GetDisplay returns the Display field value
func (o *DeviceRole) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *DeviceRole) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *DeviceRole) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *DeviceRole) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DeviceRole) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DeviceRole) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *DeviceRole) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *DeviceRole) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *DeviceRole) SetSlug(v string) {
	o.Slug = v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *DeviceRole) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRole) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *DeviceRole) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *DeviceRole) SetColor(v string) {
	o.Color = &v
}

// GetVmRole returns the VmRole field value if set, zero value otherwise.
func (o *DeviceRole) GetVmRole() bool {
	if o == nil || IsNil(o.VmRole) {
		var ret bool
		return ret
	}
	return *o.VmRole
}

// GetVmRoleOk returns a tuple with the VmRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRole) GetVmRoleOk() (*bool, bool) {
	if o == nil || IsNil(o.VmRole) {
		return nil, false
	}
	return o.VmRole, true
}

// HasVmRole returns a boolean if a field has been set.
func (o *DeviceRole) HasVmRole() bool {
	if o != nil && !IsNil(o.VmRole) {
		return true
	}

	return false
}

// SetVmRole gets a reference to the given bool and assigns it to the VmRole field.
func (o *DeviceRole) SetVmRole(v bool) {
	o.VmRole = &v
}

// GetConfigTemplate returns the ConfigTemplate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceRole) GetConfigTemplate() BriefConfigTemplate {
	if o == nil || IsNil(o.ConfigTemplate.Get()) {
		var ret BriefConfigTemplate
		return ret
	}
	return *o.ConfigTemplate.Get()
}

// GetConfigTemplateOk returns a tuple with the ConfigTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceRole) GetConfigTemplateOk() (*BriefConfigTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigTemplate.Get(), o.ConfigTemplate.IsSet()
}

// HasConfigTemplate returns a boolean if a field has been set.
func (o *DeviceRole) HasConfigTemplate() bool {
	if o != nil && o.ConfigTemplate.IsSet() {
		return true
	}

	return false
}

// SetConfigTemplate gets a reference to the given NullableBriefConfigTemplate and assigns it to the ConfigTemplate field.
func (o *DeviceRole) SetConfigTemplate(v BriefConfigTemplate) {
	o.ConfigTemplate.Set(&v)
}

// SetConfigTemplateNil sets the value for ConfigTemplate to be an explicit nil
func (o *DeviceRole) SetConfigTemplateNil() {
	o.ConfigTemplate.Set(nil)
}

// UnsetConfigTemplate ensures that no value is present for ConfigTemplate, not even an explicit nil
func (o *DeviceRole) UnsetConfigTemplate() {
	o.ConfigTemplate.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DeviceRole) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRole) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DeviceRole) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DeviceRole) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DeviceRole) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRole) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DeviceRole) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *DeviceRole) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *DeviceRole) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRole) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *DeviceRole) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *DeviceRole) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *DeviceRole) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceRole) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *DeviceRole) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *DeviceRole) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceRole) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *DeviceRole) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// GetDeviceCount returns the DeviceCount field value
func (o *DeviceRole) GetDeviceCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DeviceCount
}

// GetDeviceCountOk returns a tuple with the DeviceCount field value
// and a boolean to check if the value has been set.
func (o *DeviceRole) GetDeviceCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceCount, true
}

// SetDeviceCount sets field value
func (o *DeviceRole) SetDeviceCount(v int64) {
	o.DeviceCount = v
}

// GetVirtualmachineCount returns the VirtualmachineCount field value
func (o *DeviceRole) GetVirtualmachineCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.VirtualmachineCount
}

// GetVirtualmachineCountOk returns a tuple with the VirtualmachineCount field value
// and a boolean to check if the value has been set.
func (o *DeviceRole) GetVirtualmachineCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VirtualmachineCount, true
}

// SetVirtualmachineCount sets field value
func (o *DeviceRole) SetVirtualmachineCount(v int64) {
	o.VirtualmachineCount = v
}

func (o DeviceRole) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display_url"] = o.DisplayUrl
	toSerialize["display"] = o.Display
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
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()
	toSerialize["device_count"] = o.DeviceCount
	toSerialize["virtualmachine_count"] = o.VirtualmachineCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceRole) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display_url",
		"display",
		"name",
		"slug",
		"created",
		"last_updated",
		"device_count",
		"virtualmachine_count",
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

	varDeviceRole := _DeviceRole{}

	err = json.Unmarshal(data, &varDeviceRole)

	if err != nil {
		return err
	}

	*o = DeviceRole(varDeviceRole)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display_url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "color")
		delete(additionalProperties, "vm_role")
		delete(additionalProperties, "config_template")
		delete(additionalProperties, "description")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		delete(additionalProperties, "device_count")
		delete(additionalProperties, "virtualmachine_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceRole struct {
	value *DeviceRole
	isSet bool
}

func (v NullableDeviceRole) Get() *DeviceRole {
	return v.value
}

func (v *NullableDeviceRole) Set(val *DeviceRole) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceRole) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceRole(val *DeviceRole) *NullableDeviceRole {
	return &NullableDeviceRole{value: val, isSet: true}
}

func (v NullableDeviceRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
