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

// checks if the ClusterRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterRequest{}

// ClusterRequest Adds support for custom fields and tags.
type ClusterRequest struct {
	Name string `json:"name"`
	Type BriefClusterTypeRequest `json:"type"`
	Group NullableBriefClusterGroupRequest `json:"group,omitempty"`
	Status *ClusterStatusValue `json:"status,omitempty"`
	Tenant NullableBriefTenantRequest `json:"tenant,omitempty"`
	Site NullableBriefSiteRequest `json:"site,omitempty"`
	Description *string `json:"description,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Tags []NestedTagRequest `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ClusterRequest ClusterRequest

// NewClusterRequest instantiates a new ClusterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterRequest(name string, type_ BriefClusterTypeRequest) *ClusterRequest {
	this := ClusterRequest{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewClusterRequestWithDefaults instantiates a new ClusterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterRequestWithDefaults() *ClusterRequest {
	this := ClusterRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ClusterRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ClusterRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ClusterRequest) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *ClusterRequest) GetType() BriefClusterTypeRequest {
	if o == nil {
		var ret BriefClusterTypeRequest
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ClusterRequest) GetTypeOk() (*BriefClusterTypeRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ClusterRequest) SetType(v BriefClusterTypeRequest) {
	o.Type = v
}

// GetGroup returns the Group field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterRequest) GetGroup() BriefClusterGroupRequest {
	if o == nil || IsNil(o.Group.Get()) {
		var ret BriefClusterGroupRequest
		return ret
	}
	return *o.Group.Get()
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterRequest) GetGroupOk() (*BriefClusterGroupRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Group.Get(), o.Group.IsSet()
}

// HasGroup returns a boolean if a field has been set.
func (o *ClusterRequest) HasGroup() bool {
	if o != nil && o.Group.IsSet() {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NullableBriefClusterGroupRequest and assigns it to the Group field.
func (o *ClusterRequest) SetGroup(v BriefClusterGroupRequest) {
	o.Group.Set(&v)
}
// SetGroupNil sets the value for Group to be an explicit nil
func (o *ClusterRequest) SetGroupNil() {
	o.Group.Set(nil)
}

// UnsetGroup ensures that no value is present for Group, not even an explicit nil
func (o *ClusterRequest) UnsetGroup() {
	o.Group.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ClusterRequest) GetStatus() ClusterStatusValue {
	if o == nil || IsNil(o.Status) {
		var ret ClusterStatusValue
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRequest) GetStatusOk() (*ClusterStatusValue, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ClusterRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ClusterStatusValue and assigns it to the Status field.
func (o *ClusterRequest) SetStatus(v ClusterStatusValue) {
	o.Status = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterRequest) GetTenant() BriefTenantRequest {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BriefTenantRequest
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterRequest) GetTenantOk() (*BriefTenantRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *ClusterRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBriefTenantRequest and assigns it to the Tenant field.
func (o *ClusterRequest) SetTenant(v BriefTenantRequest) {
	o.Tenant.Set(&v)
}
// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *ClusterRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *ClusterRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetSite returns the Site field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterRequest) GetSite() BriefSiteRequest {
	if o == nil || IsNil(o.Site.Get()) {
		var ret BriefSiteRequest
		return ret
	}
	return *o.Site.Get()
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterRequest) GetSiteOk() (*BriefSiteRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Site.Get(), o.Site.IsSet()
}

// HasSite returns a boolean if a field has been set.
func (o *ClusterRequest) HasSite() bool {
	if o != nil && o.Site.IsSet() {
		return true
	}

	return false
}

// SetSite gets a reference to the given NullableBriefSiteRequest and assigns it to the Site field.
func (o *ClusterRequest) SetSite(v BriefSiteRequest) {
	o.Site.Set(&v)
}
// SetSiteNil sets the value for Site to be an explicit nil
func (o *ClusterRequest) SetSiteNil() {
	o.Site.Set(nil)
}

// UnsetSite ensures that no value is present for Site, not even an explicit nil
func (o *ClusterRequest) UnsetSite() {
	o.Site.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ClusterRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ClusterRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ClusterRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *ClusterRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *ClusterRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *ClusterRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ClusterRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ClusterRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *ClusterRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *ClusterRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *ClusterRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *ClusterRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o ClusterRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	if o.Group.IsSet() {
		toSerialize["group"] = o.Group.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if o.Site.IsSet() {
		toSerialize["site"] = o.Site.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
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

func (o *ClusterRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
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

	varClusterRequest := _ClusterRequest{}

	err = json.Unmarshal(data, &varClusterRequest)

	if err != nil {
		return err
	}

	*o = ClusterRequest(varClusterRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "group")
		delete(additionalProperties, "status")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "site")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClusterRequest struct {
	value *ClusterRequest
	isSet bool
}

func (v NullableClusterRequest) Get() *ClusterRequest {
	return v.value
}

func (v *NullableClusterRequest) Set(val *ClusterRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterRequest(val *ClusterRequest) *NullableClusterRequest {
	return &NullableClusterRequest{value: val, isSet: true}
}

func (v NullableClusterRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


