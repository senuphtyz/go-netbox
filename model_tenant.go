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

// checks if the Tenant type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Tenant{}

// Tenant Adds support for custom fields and tags.
type Tenant struct {
	Id int32 `json:"id"`
	Url string `json:"url"`
	DisplayUrl string `json:"display_url"`
	Display string `json:"display"`
	Name string `json:"name"`
	Slug string `json:"slug" validate:"regexp=^[-a-zA-Z0-9_]+$"`
	Group NullableBriefTenantGroup `json:"group,omitempty"`
	Description *string `json:"description,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Tags []NestedTag `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Created NullableTime `json:"created"`
	LastUpdated NullableTime `json:"last_updated"`
	CircuitCount int64 `json:"circuit_count"`
	DeviceCount int64 `json:"device_count"`
	IpaddressCount int64 `json:"ipaddress_count"`
	PrefixCount int64 `json:"prefix_count"`
	RackCount int64 `json:"rack_count"`
	SiteCount int64 `json:"site_count"`
	VirtualmachineCount int64 `json:"virtualmachine_count"`
	VlanCount int64 `json:"vlan_count"`
	VrfCount int64 `json:"vrf_count"`
	ClusterCount int64 `json:"cluster_count"`
	AdditionalProperties map[string]interface{}
}

type _Tenant Tenant

// NewTenant instantiates a new Tenant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenant(id int32, url string, displayUrl string, display string, name string, slug string, created NullableTime, lastUpdated NullableTime, circuitCount int64, deviceCount int64, ipaddressCount int64, prefixCount int64, rackCount int64, siteCount int64, virtualmachineCount int64, vlanCount int64, vrfCount int64, clusterCount int64) *Tenant {
	this := Tenant{}
	this.Id = id
	this.Url = url
	this.DisplayUrl = displayUrl
	this.Display = display
	this.Name = name
	this.Slug = slug
	this.Created = created
	this.LastUpdated = lastUpdated
	this.CircuitCount = circuitCount
	this.DeviceCount = deviceCount
	this.IpaddressCount = ipaddressCount
	this.PrefixCount = prefixCount
	this.RackCount = rackCount
	this.SiteCount = siteCount
	this.VirtualmachineCount = virtualmachineCount
	this.VlanCount = vlanCount
	this.VrfCount = vrfCount
	this.ClusterCount = clusterCount
	return &this
}

// NewTenantWithDefaults instantiates a new Tenant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantWithDefaults() *Tenant {
	this := Tenant{}
	return &this
}

// GetId returns the Id field value
func (o *Tenant) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Tenant) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *Tenant) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Tenant) SetUrl(v string) {
	o.Url = v
}

// GetDisplayUrl returns the DisplayUrl field value
func (o *Tenant) GetDisplayUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayUrl
}

// GetDisplayUrlOk returns a tuple with the DisplayUrl field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetDisplayUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayUrl, true
}

// SetDisplayUrl sets field value
func (o *Tenant) SetDisplayUrl(v string) {
	o.DisplayUrl = v
}

// GetDisplay returns the Display field value
func (o *Tenant) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *Tenant) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *Tenant) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Tenant) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *Tenant) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *Tenant) SetSlug(v string) {
	o.Slug = v
}

// GetGroup returns the Group field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tenant) GetGroup() BriefTenantGroup {
	if o == nil || IsNil(o.Group.Get()) {
		var ret BriefTenantGroup
		return ret
	}
	return *o.Group.Get()
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tenant) GetGroupOk() (*BriefTenantGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.Group.Get(), o.Group.IsSet()
}

// HasGroup returns a boolean if a field has been set.
func (o *Tenant) HasGroup() bool {
	if o != nil && o.Group.IsSet() {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NullableBriefTenantGroup and assigns it to the Group field.
func (o *Tenant) SetGroup(v BriefTenantGroup) {
	o.Group.Set(&v)
}
// SetGroupNil sets the value for Group to be an explicit nil
func (o *Tenant) SetGroupNil() {
	o.Group.Set(nil)
}

// UnsetGroup ensures that no value is present for Group, not even an explicit nil
func (o *Tenant) UnsetGroup() {
	o.Group.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Tenant) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Tenant) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Tenant) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *Tenant) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *Tenant) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *Tenant) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Tenant) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Tenant) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *Tenant) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *Tenant) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *Tenant) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *Tenant) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Tenant) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tenant) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *Tenant) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Tenant) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tenant) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *Tenant) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// GetCircuitCount returns the CircuitCount field value
func (o *Tenant) GetCircuitCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CircuitCount
}

// GetCircuitCountOk returns a tuple with the CircuitCount field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetCircuitCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CircuitCount, true
}

// SetCircuitCount sets field value
func (o *Tenant) SetCircuitCount(v int64) {
	o.CircuitCount = v
}

// GetDeviceCount returns the DeviceCount field value
func (o *Tenant) GetDeviceCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DeviceCount
}

// GetDeviceCountOk returns a tuple with the DeviceCount field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetDeviceCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceCount, true
}

// SetDeviceCount sets field value
func (o *Tenant) SetDeviceCount(v int64) {
	o.DeviceCount = v
}

// GetIpaddressCount returns the IpaddressCount field value
func (o *Tenant) GetIpaddressCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.IpaddressCount
}

// GetIpaddressCountOk returns a tuple with the IpaddressCount field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetIpaddressCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IpaddressCount, true
}

// SetIpaddressCount sets field value
func (o *Tenant) SetIpaddressCount(v int64) {
	o.IpaddressCount = v
}

// GetPrefixCount returns the PrefixCount field value
func (o *Tenant) GetPrefixCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PrefixCount
}

// GetPrefixCountOk returns a tuple with the PrefixCount field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetPrefixCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrefixCount, true
}

// SetPrefixCount sets field value
func (o *Tenant) SetPrefixCount(v int64) {
	o.PrefixCount = v
}

// GetRackCount returns the RackCount field value
func (o *Tenant) GetRackCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RackCount
}

// GetRackCountOk returns a tuple with the RackCount field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetRackCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RackCount, true
}

// SetRackCount sets field value
func (o *Tenant) SetRackCount(v int64) {
	o.RackCount = v
}

// GetSiteCount returns the SiteCount field value
func (o *Tenant) GetSiteCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SiteCount
}

// GetSiteCountOk returns a tuple with the SiteCount field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetSiteCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiteCount, true
}

// SetSiteCount sets field value
func (o *Tenant) SetSiteCount(v int64) {
	o.SiteCount = v
}

// GetVirtualmachineCount returns the VirtualmachineCount field value
func (o *Tenant) GetVirtualmachineCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.VirtualmachineCount
}

// GetVirtualmachineCountOk returns a tuple with the VirtualmachineCount field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetVirtualmachineCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VirtualmachineCount, true
}

// SetVirtualmachineCount sets field value
func (o *Tenant) SetVirtualmachineCount(v int64) {
	o.VirtualmachineCount = v
}

// GetVlanCount returns the VlanCount field value
func (o *Tenant) GetVlanCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.VlanCount
}

// GetVlanCountOk returns a tuple with the VlanCount field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetVlanCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VlanCount, true
}

// SetVlanCount sets field value
func (o *Tenant) SetVlanCount(v int64) {
	o.VlanCount = v
}

// GetVrfCount returns the VrfCount field value
func (o *Tenant) GetVrfCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.VrfCount
}

// GetVrfCountOk returns a tuple with the VrfCount field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetVrfCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VrfCount, true
}

// SetVrfCount sets field value
func (o *Tenant) SetVrfCount(v int64) {
	o.VrfCount = v
}

// GetClusterCount returns the ClusterCount field value
func (o *Tenant) GetClusterCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ClusterCount
}

// GetClusterCountOk returns a tuple with the ClusterCount field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetClusterCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterCount, true
}

// SetClusterCount sets field value
func (o *Tenant) SetClusterCount(v int64) {
	o.ClusterCount = v
}

func (o Tenant) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Tenant) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display_url"] = o.DisplayUrl
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	if o.Group.IsSet() {
		toSerialize["group"] = o.Group.Get()
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
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()
	toSerialize["circuit_count"] = o.CircuitCount
	toSerialize["device_count"] = o.DeviceCount
	toSerialize["ipaddress_count"] = o.IpaddressCount
	toSerialize["prefix_count"] = o.PrefixCount
	toSerialize["rack_count"] = o.RackCount
	toSerialize["site_count"] = o.SiteCount
	toSerialize["virtualmachine_count"] = o.VirtualmachineCount
	toSerialize["vlan_count"] = o.VlanCount
	toSerialize["vrf_count"] = o.VrfCount
	toSerialize["cluster_count"] = o.ClusterCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Tenant) UnmarshalJSON(data []byte) (err error) {
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
		"device_count",
		"ipaddress_count",
		"prefix_count",
		"rack_count",
		"site_count",
		"virtualmachine_count",
		"vlan_count",
		"vrf_count",
		"cluster_count",
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

	varTenant := _Tenant{}

	err = json.Unmarshal(data, &varTenant)

	if err != nil {
		return err
	}

	*o = Tenant(varTenant)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display_url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "group")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		delete(additionalProperties, "circuit_count")
		delete(additionalProperties, "device_count")
		delete(additionalProperties, "ipaddress_count")
		delete(additionalProperties, "prefix_count")
		delete(additionalProperties, "rack_count")
		delete(additionalProperties, "site_count")
		delete(additionalProperties, "virtualmachine_count")
		delete(additionalProperties, "vlan_count")
		delete(additionalProperties, "vrf_count")
		delete(additionalProperties, "cluster_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTenant struct {
	value *Tenant
	isSet bool
}

func (v NullableTenant) Get() *Tenant {
	return v.value
}

func (v *NullableTenant) Set(val *Tenant) {
	v.value = val
	v.isSet = true
}

func (v NullableTenant) IsSet() bool {
	return v.isSet
}

func (v *NullableTenant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenant(val *Tenant) *NullableTenant {
	return &NullableTenant{value: val, isSet: true}
}

func (v NullableTenant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


