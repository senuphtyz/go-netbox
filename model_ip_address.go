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

// checks if the IPAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IPAddress{}

// IPAddress Adds support for custom fields and tags.
type IPAddress struct {
	Id                 int32                   `json:"id"`
	Url                string                  `json:"url"`
	DisplayUrl         string                  `json:"display_url"`
	Display            string                  `json:"display"`
	Family             AggregateFamily         `json:"family"`
	Address            string                  `json:"address"`
	Vrf                NullableBriefVRF        `json:"vrf,omitempty"`
	Tenant             NullableBriefTenant     `json:"tenant,omitempty"`
	Status             *IPAddressStatus        `json:"status,omitempty"`
	Role               *IPAddressRole          `json:"role,omitempty"`
	AssignedObjectType NullableString          `json:"assigned_object_type,omitempty"`
	AssignedObjectId   NullableInt64           `json:"assigned_object_id,omitempty"`
	AssignedObject     interface{}             `json:"assigned_object"`
	NatInside          NullableNestedIPAddress `json:"nat_inside,omitempty"`
	NatOutside         []NestedIPAddress       `json:"nat_outside"`
	// Hostname or FQDN (not case-sensitive)
	DnsName              *string                `json:"dns_name,omitempty" validate:"regexp=^([0-9A-Za-z_-]+|\\\\*)(\\\\.[0-9A-Za-z_-]+)*\\\\.?$"`
	Description          *string                `json:"description,omitempty"`
	Comments             *string                `json:"comments,omitempty"`
	Tags                 []NestedTag            `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	Created              NullableTime           `json:"created"`
	LastUpdated          NullableTime           `json:"last_updated"`
	AdditionalProperties map[string]interface{}
}

type _IPAddress IPAddress

// NewIPAddress instantiates a new IPAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIPAddress(id int32, url string, displayUrl string, display string, family AggregateFamily, address string, assignedObject interface{}, natOutside []NestedIPAddress, created NullableTime, lastUpdated NullableTime) *IPAddress {
	this := IPAddress{}
	this.Id = id
	this.Url = url
	this.DisplayUrl = displayUrl
	this.Display = display
	this.Family = family
	this.Address = address
	this.AssignedObject = assignedObject
	this.NatOutside = natOutside
	this.Created = created
	this.LastUpdated = lastUpdated
	return &this
}

// NewIPAddressWithDefaults instantiates a new IPAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIPAddressWithDefaults() *IPAddress {
	this := IPAddress{}
	return &this
}

// GetId returns the Id field value
func (o *IPAddress) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IPAddress) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IPAddress) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *IPAddress) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *IPAddress) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *IPAddress) SetUrl(v string) {
	o.Url = v
}

// GetDisplayUrl returns the DisplayUrl field value
func (o *IPAddress) GetDisplayUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayUrl
}

// GetDisplayUrlOk returns a tuple with the DisplayUrl field value
// and a boolean to check if the value has been set.
func (o *IPAddress) GetDisplayUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayUrl, true
}

// SetDisplayUrl sets field value
func (o *IPAddress) SetDisplayUrl(v string) {
	o.DisplayUrl = v
}

// GetDisplay returns the Display field value
func (o *IPAddress) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *IPAddress) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *IPAddress) SetDisplay(v string) {
	o.Display = v
}

// GetFamily returns the Family field value
func (o *IPAddress) GetFamily() AggregateFamily {
	if o == nil {
		var ret AggregateFamily
		return ret
	}

	return o.Family
}

// GetFamilyOk returns a tuple with the Family field value
// and a boolean to check if the value has been set.
func (o *IPAddress) GetFamilyOk() (*AggregateFamily, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Family, true
}

// SetFamily sets field value
func (o *IPAddress) SetFamily(v AggregateFamily) {
	o.Family = v
}

// GetAddress returns the Address field value
func (o *IPAddress) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *IPAddress) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *IPAddress) SetAddress(v string) {
	o.Address = v
}

// GetVrf returns the Vrf field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IPAddress) GetVrf() BriefVRF {
	if o == nil || IsNil(o.Vrf.Get()) {
		var ret BriefVRF
		return ret
	}
	return *o.Vrf.Get()
}

// GetVrfOk returns a tuple with the Vrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPAddress) GetVrfOk() (*BriefVRF, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vrf.Get(), o.Vrf.IsSet()
}

// HasVrf returns a boolean if a field has been set.
func (o *IPAddress) HasVrf() bool {
	if o != nil && o.Vrf.IsSet() {
		return true
	}

	return false
}

// SetVrf gets a reference to the given NullableBriefVRF and assigns it to the Vrf field.
func (o *IPAddress) SetVrf(v BriefVRF) {
	o.Vrf.Set(&v)
}

// SetVrfNil sets the value for Vrf to be an explicit nil
func (o *IPAddress) SetVrfNil() {
	o.Vrf.Set(nil)
}

// UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
func (o *IPAddress) UnsetVrf() {
	o.Vrf.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IPAddress) GetTenant() BriefTenant {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BriefTenant
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPAddress) GetTenantOk() (*BriefTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *IPAddress) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBriefTenant and assigns it to the Tenant field.
func (o *IPAddress) SetTenant(v BriefTenant) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *IPAddress) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *IPAddress) UnsetTenant() {
	o.Tenant.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IPAddress) GetStatus() IPAddressStatus {
	if o == nil || IsNil(o.Status) {
		var ret IPAddressStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPAddress) GetStatusOk() (*IPAddressStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IPAddress) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given IPAddressStatus and assigns it to the Status field.
func (o *IPAddress) SetStatus(v IPAddressStatus) {
	o.Status = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *IPAddress) GetRole() IPAddressRole {
	if o == nil || IsNil(o.Role) {
		var ret IPAddressRole
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPAddress) GetRoleOk() (*IPAddressRole, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *IPAddress) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given IPAddressRole and assigns it to the Role field.
func (o *IPAddress) SetRole(v IPAddressRole) {
	o.Role = &v
}

// GetAssignedObjectType returns the AssignedObjectType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IPAddress) GetAssignedObjectType() string {
	if o == nil || IsNil(o.AssignedObjectType.Get()) {
		var ret string
		return ret
	}
	return *o.AssignedObjectType.Get()
}

// GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPAddress) GetAssignedObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignedObjectType.Get(), o.AssignedObjectType.IsSet()
}

// HasAssignedObjectType returns a boolean if a field has been set.
func (o *IPAddress) HasAssignedObjectType() bool {
	if o != nil && o.AssignedObjectType.IsSet() {
		return true
	}

	return false
}

// SetAssignedObjectType gets a reference to the given NullableString and assigns it to the AssignedObjectType field.
func (o *IPAddress) SetAssignedObjectType(v string) {
	o.AssignedObjectType.Set(&v)
}

// SetAssignedObjectTypeNil sets the value for AssignedObjectType to be an explicit nil
func (o *IPAddress) SetAssignedObjectTypeNil() {
	o.AssignedObjectType.Set(nil)
}

// UnsetAssignedObjectType ensures that no value is present for AssignedObjectType, not even an explicit nil
func (o *IPAddress) UnsetAssignedObjectType() {
	o.AssignedObjectType.Unset()
}

// GetAssignedObjectId returns the AssignedObjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IPAddress) GetAssignedObjectId() int64 {
	if o == nil || IsNil(o.AssignedObjectId.Get()) {
		var ret int64
		return ret
	}
	return *o.AssignedObjectId.Get()
}

// GetAssignedObjectIdOk returns a tuple with the AssignedObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPAddress) GetAssignedObjectIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignedObjectId.Get(), o.AssignedObjectId.IsSet()
}

// HasAssignedObjectId returns a boolean if a field has been set.
func (o *IPAddress) HasAssignedObjectId() bool {
	if o != nil && o.AssignedObjectId.IsSet() {
		return true
	}

	return false
}

// SetAssignedObjectId gets a reference to the given NullableInt64 and assigns it to the AssignedObjectId field.
func (o *IPAddress) SetAssignedObjectId(v int64) {
	o.AssignedObjectId.Set(&v)
}

// SetAssignedObjectIdNil sets the value for AssignedObjectId to be an explicit nil
func (o *IPAddress) SetAssignedObjectIdNil() {
	o.AssignedObjectId.Set(nil)
}

// UnsetAssignedObjectId ensures that no value is present for AssignedObjectId, not even an explicit nil
func (o *IPAddress) UnsetAssignedObjectId() {
	o.AssignedObjectId.Unset()
}

// GetAssignedObject returns the AssignedObject field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *IPAddress) GetAssignedObject() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.AssignedObject
}

// GetAssignedObjectOk returns a tuple with the AssignedObject field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPAddress) GetAssignedObjectOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AssignedObject) {
		return nil, false
	}
	return &o.AssignedObject, true
}

// SetAssignedObject sets field value
func (o *IPAddress) SetAssignedObject(v interface{}) {
	o.AssignedObject = v
}

// GetNatInside returns the NatInside field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IPAddress) GetNatInside() NestedIPAddress {
	if o == nil || IsNil(o.NatInside.Get()) {
		var ret NestedIPAddress
		return ret
	}
	return *o.NatInside.Get()
}

// GetNatInsideOk returns a tuple with the NatInside field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPAddress) GetNatInsideOk() (*NestedIPAddress, bool) {
	if o == nil {
		return nil, false
	}
	return o.NatInside.Get(), o.NatInside.IsSet()
}

// HasNatInside returns a boolean if a field has been set.
func (o *IPAddress) HasNatInside() bool {
	if o != nil && o.NatInside.IsSet() {
		return true
	}

	return false
}

// SetNatInside gets a reference to the given NullableNestedIPAddress and assigns it to the NatInside field.
func (o *IPAddress) SetNatInside(v NestedIPAddress) {
	o.NatInside.Set(&v)
}

// SetNatInsideNil sets the value for NatInside to be an explicit nil
func (o *IPAddress) SetNatInsideNil() {
	o.NatInside.Set(nil)
}

// UnsetNatInside ensures that no value is present for NatInside, not even an explicit nil
func (o *IPAddress) UnsetNatInside() {
	o.NatInside.Unset()
}

// GetNatOutside returns the NatOutside field value
func (o *IPAddress) GetNatOutside() []NestedIPAddress {
	if o == nil {
		var ret []NestedIPAddress
		return ret
	}

	return o.NatOutside
}

// GetNatOutsideOk returns a tuple with the NatOutside field value
// and a boolean to check if the value has been set.
func (o *IPAddress) GetNatOutsideOk() ([]NestedIPAddress, bool) {
	if o == nil {
		return nil, false
	}
	return o.NatOutside, true
}

// SetNatOutside sets field value
func (o *IPAddress) SetNatOutside(v []NestedIPAddress) {
	o.NatOutside = v
}

// GetDnsName returns the DnsName field value if set, zero value otherwise.
func (o *IPAddress) GetDnsName() string {
	if o == nil || IsNil(o.DnsName) {
		var ret string
		return ret
	}
	return *o.DnsName
}

// GetDnsNameOk returns a tuple with the DnsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPAddress) GetDnsNameOk() (*string, bool) {
	if o == nil || IsNil(o.DnsName) {
		return nil, false
	}
	return o.DnsName, true
}

// HasDnsName returns a boolean if a field has been set.
func (o *IPAddress) HasDnsName() bool {
	if o != nil && !IsNil(o.DnsName) {
		return true
	}

	return false
}

// SetDnsName gets a reference to the given string and assigns it to the DnsName field.
func (o *IPAddress) SetDnsName(v string) {
	o.DnsName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IPAddress) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPAddress) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IPAddress) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IPAddress) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *IPAddress) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPAddress) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *IPAddress) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *IPAddress) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *IPAddress) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPAddress) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *IPAddress) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *IPAddress) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *IPAddress) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPAddress) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *IPAddress) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *IPAddress) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *IPAddress) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPAddress) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *IPAddress) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *IPAddress) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPAddress) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *IPAddress) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

func (o IPAddress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IPAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display_url"] = o.DisplayUrl
	toSerialize["display"] = o.Display
	toSerialize["family"] = o.Family
	toSerialize["address"] = o.Address
	if o.Vrf.IsSet() {
		toSerialize["vrf"] = o.Vrf.Get()
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if o.AssignedObjectType.IsSet() {
		toSerialize["assigned_object_type"] = o.AssignedObjectType.Get()
	}
	if o.AssignedObjectId.IsSet() {
		toSerialize["assigned_object_id"] = o.AssignedObjectId.Get()
	}
	if o.AssignedObject != nil {
		toSerialize["assigned_object"] = o.AssignedObject
	}
	if o.NatInside.IsSet() {
		toSerialize["nat_inside"] = o.NatInside.Get()
	}
	toSerialize["nat_outside"] = o.NatOutside
	if !IsNil(o.DnsName) {
		toSerialize["dns_name"] = o.DnsName
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IPAddress) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display_url",
		"display",
		"family",
		"address",
		"assigned_object",
		"nat_outside",
		"created",
		"last_updated",
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

	varIPAddress := _IPAddress{}

	err = json.Unmarshal(data, &varIPAddress)

	if err != nil {
		return err
	}

	*o = IPAddress(varIPAddress)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display_url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "family")
		delete(additionalProperties, "address")
		delete(additionalProperties, "vrf")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "status")
		delete(additionalProperties, "role")
		delete(additionalProperties, "assigned_object_type")
		delete(additionalProperties, "assigned_object_id")
		delete(additionalProperties, "assigned_object")
		delete(additionalProperties, "nat_inside")
		delete(additionalProperties, "nat_outside")
		delete(additionalProperties, "dns_name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIPAddress struct {
	value *IPAddress
	isSet bool
}

func (v NullableIPAddress) Get() *IPAddress {
	return v.value
}

func (v *NullableIPAddress) Set(val *IPAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableIPAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableIPAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPAddress(val *IPAddress) *NullableIPAddress {
	return &NullableIPAddress{value: val, isSet: true}
}

func (v NullableIPAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
