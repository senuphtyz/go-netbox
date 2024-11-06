/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.1.4 (4.1)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the CircuitGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CircuitGroup{}

// CircuitGroup Adds support for custom fields and tags.
type CircuitGroup struct {
	Id int32 `json:"id"`
	Url string `json:"url"`
	DisplayUrl string `json:"display_url"`
	Display string `json:"display"`
	Name string `json:"name"`
	Slug string `json:"slug" validate:"regexp=^[-a-zA-Z0-9_]+$"`
	Description *string `json:"description,omitempty"`
	Tenant NullableBriefTenant `json:"tenant,omitempty"`
	Tags []NestedTag `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Created NullableTime `json:"created"`
	LastUpdated NullableTime `json:"last_updated"`
	CircuitCount int64 `json:"circuit_count"`
	AdditionalProperties map[string]interface{}
}

type _CircuitGroup CircuitGroup

// NewCircuitGroup instantiates a new CircuitGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCircuitGroup(id int32, url string, displayUrl string, display string, name string, slug string, created NullableTime, lastUpdated NullableTime, circuitCount int64) *CircuitGroup {
	this := CircuitGroup{}
	this.Id = id
	this.Url = url
	this.DisplayUrl = displayUrl
	this.Display = display
	this.Name = name
	this.Slug = slug
	this.Created = created
	this.LastUpdated = lastUpdated
	this.CircuitCount = circuitCount
	return &this
}

// NewCircuitGroupWithDefaults instantiates a new CircuitGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCircuitGroupWithDefaults() *CircuitGroup {
	this := CircuitGroup{}
	return &this
}

// GetId returns the Id field value
func (o *CircuitGroup) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CircuitGroup) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CircuitGroup) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *CircuitGroup) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *CircuitGroup) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *CircuitGroup) SetUrl(v string) {
	o.Url = v
}

// GetDisplayUrl returns the DisplayUrl field value
func (o *CircuitGroup) GetDisplayUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayUrl
}

// GetDisplayUrlOk returns a tuple with the DisplayUrl field value
// and a boolean to check if the value has been set.
func (o *CircuitGroup) GetDisplayUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayUrl, true
}

// SetDisplayUrl sets field value
func (o *CircuitGroup) SetDisplayUrl(v string) {
	o.DisplayUrl = v
}

// GetDisplay returns the Display field value
func (o *CircuitGroup) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *CircuitGroup) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *CircuitGroup) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *CircuitGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CircuitGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CircuitGroup) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *CircuitGroup) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *CircuitGroup) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *CircuitGroup) SetSlug(v string) {
	o.Slug = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CircuitGroup) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitGroup) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CircuitGroup) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CircuitGroup) SetDescription(v string) {
	o.Description = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CircuitGroup) GetTenant() BriefTenant {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BriefTenant
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CircuitGroup) GetTenantOk() (*BriefTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *CircuitGroup) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBriefTenant and assigns it to the Tenant field.
func (o *CircuitGroup) SetTenant(v BriefTenant) {
	o.Tenant.Set(&v)
}
// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *CircuitGroup) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *CircuitGroup) UnsetTenant() {
	o.Tenant.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CircuitGroup) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitGroup) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CircuitGroup) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *CircuitGroup) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *CircuitGroup) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitGroup) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *CircuitGroup) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *CircuitGroup) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *CircuitGroup) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CircuitGroup) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *CircuitGroup) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *CircuitGroup) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CircuitGroup) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *CircuitGroup) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// GetCircuitCount returns the CircuitCount field value
func (o *CircuitGroup) GetCircuitCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CircuitCount
}

// GetCircuitCountOk returns a tuple with the CircuitCount field value
// and a boolean to check if the value has been set.
func (o *CircuitGroup) GetCircuitCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CircuitCount, true
}

// SetCircuitCount sets field value
func (o *CircuitGroup) SetCircuitCount(v int64) {
	o.CircuitCount = v
}

func (o CircuitGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CircuitGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display_url"] = o.DisplayUrl
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()
	toSerialize["circuit_count"] = o.CircuitCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CircuitGroup) UnmarshalJSON(data []byte) (err error) {
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
		"circuit_count",
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

	varCircuitGroup := _CircuitGroup{}

	err = json.Unmarshal(data, &varCircuitGroup)

	if err != nil {
		return err
	}

	*o = CircuitGroup(varCircuitGroup)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display_url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "description")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		delete(additionalProperties, "circuit_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCircuitGroup struct {
	value *CircuitGroup
	isSet bool
}

func (v NullableCircuitGroup) Get() *CircuitGroup {
	return v.value
}

func (v *NullableCircuitGroup) Set(val *CircuitGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableCircuitGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableCircuitGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCircuitGroup(val *CircuitGroup) *NullableCircuitGroup {
	return &NullableCircuitGroup{value: val, isSet: true}
}

func (v NullableCircuitGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCircuitGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


