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

// checks if the BriefVirtualChassis type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefVirtualChassis{}

// BriefVirtualChassis Adds support for custom fields and tags.
type BriefVirtualChassis struct {
	Id                   int32                `json:"id"`
	Url                  string               `json:"url"`
	Display              string               `json:"display"`
	Name                 string               `json:"name"`
	Master               NullableNestedDevice `json:"master,omitempty"`
	Description          *string              `json:"description,omitempty"`
	MemberCount          int32                `json:"member_count"`
	AdditionalProperties map[string]interface{}
}

type _BriefVirtualChassis BriefVirtualChassis

// NewBriefVirtualChassis instantiates a new BriefVirtualChassis object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefVirtualChassis(id int32, url string, display string, name string, memberCount int32) *BriefVirtualChassis {
	this := BriefVirtualChassis{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	this.MemberCount = memberCount
	return &this
}

// NewBriefVirtualChassisWithDefaults instantiates a new BriefVirtualChassis object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefVirtualChassisWithDefaults() *BriefVirtualChassis {
	this := BriefVirtualChassis{}
	return &this
}

// GetId returns the Id field value
func (o *BriefVirtualChassis) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BriefVirtualChassis) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BriefVirtualChassis) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *BriefVirtualChassis) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *BriefVirtualChassis) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *BriefVirtualChassis) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *BriefVirtualChassis) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *BriefVirtualChassis) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *BriefVirtualChassis) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *BriefVirtualChassis) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BriefVirtualChassis) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BriefVirtualChassis) SetName(v string) {
	o.Name = v
}

// GetMaster returns the Master field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BriefVirtualChassis) GetMaster() NestedDevice {
	if o == nil || IsNil(o.Master.Get()) {
		var ret NestedDevice
		return ret
	}
	return *o.Master.Get()
}

// GetMasterOk returns a tuple with the Master field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BriefVirtualChassis) GetMasterOk() (*NestedDevice, bool) {
	if o == nil {
		return nil, false
	}
	return o.Master.Get(), o.Master.IsSet()
}

// HasMaster returns a boolean if a field has been set.
func (o *BriefVirtualChassis) HasMaster() bool {
	if o != nil && o.Master.IsSet() {
		return true
	}

	return false
}

// SetMaster gets a reference to the given NullableNestedDevice and assigns it to the Master field.
func (o *BriefVirtualChassis) SetMaster(v NestedDevice) {
	o.Master.Set(&v)
}

// SetMasterNil sets the value for Master to be an explicit nil
func (o *BriefVirtualChassis) SetMasterNil() {
	o.Master.Set(nil)
}

// UnsetMaster ensures that no value is present for Master, not even an explicit nil
func (o *BriefVirtualChassis) UnsetMaster() {
	o.Master.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BriefVirtualChassis) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefVirtualChassis) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BriefVirtualChassis) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BriefVirtualChassis) SetDescription(v string) {
	o.Description = &v
}

// GetMemberCount returns the MemberCount field value
func (o *BriefVirtualChassis) GetMemberCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MemberCount
}

// GetMemberCountOk returns a tuple with the MemberCount field value
// and a boolean to check if the value has been set.
func (o *BriefVirtualChassis) GetMemberCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MemberCount, true
}

// SetMemberCount sets field value
func (o *BriefVirtualChassis) SetMemberCount(v int32) {
	o.MemberCount = v
}

func (o BriefVirtualChassis) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefVirtualChassis) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name
	if o.Master.IsSet() {
		toSerialize["master"] = o.Master.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["member_count"] = o.MemberCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefVirtualChassis) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"name",
		"member_count",
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

	varBriefVirtualChassis := _BriefVirtualChassis{}

	err = json.Unmarshal(data, &varBriefVirtualChassis)

	if err != nil {
		return err
	}

	*o = BriefVirtualChassis(varBriefVirtualChassis)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "master")
		delete(additionalProperties, "description")
		delete(additionalProperties, "member_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefVirtualChassis struct {
	value *BriefVirtualChassis
	isSet bool
}

func (v NullableBriefVirtualChassis) Get() *BriefVirtualChassis {
	return v.value
}

func (v *NullableBriefVirtualChassis) Set(val *BriefVirtualChassis) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefVirtualChassis) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefVirtualChassis) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefVirtualChassis(val *BriefVirtualChassis) *NullableBriefVirtualChassis {
	return &NullableBriefVirtualChassis{value: val, isSet: true}
}

func (v NullableBriefVirtualChassis) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefVirtualChassis) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
