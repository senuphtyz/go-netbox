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

// checks if the BriefWirelessLANGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefWirelessLANGroup{}

// BriefWirelessLANGroup Extends PrimaryModelSerializer to include MPTT support.
type BriefWirelessLANGroup struct {
	Id int32 `json:"id"`
	Url string `json:"url"`
	Display string `json:"display"`
	Name string `json:"name"`
	Slug string `json:"slug" validate:"regexp=^[-a-zA-Z0-9_]+$"`
	Description *string `json:"description,omitempty"`
	WirelesslanCount int32 `json:"wirelesslan_count"`
	Depth int32 `json:"_depth"`
	AdditionalProperties map[string]interface{}
}

type _BriefWirelessLANGroup BriefWirelessLANGroup

// NewBriefWirelessLANGroup instantiates a new BriefWirelessLANGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefWirelessLANGroup(id int32, url string, display string, name string, slug string, wirelesslanCount int32, depth int32) *BriefWirelessLANGroup {
	this := BriefWirelessLANGroup{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	this.Slug = slug
	this.WirelesslanCount = wirelesslanCount
	this.Depth = depth
	return &this
}

// NewBriefWirelessLANGroupWithDefaults instantiates a new BriefWirelessLANGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefWirelessLANGroupWithDefaults() *BriefWirelessLANGroup {
	this := BriefWirelessLANGroup{}
	return &this
}

// GetId returns the Id field value
func (o *BriefWirelessLANGroup) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BriefWirelessLANGroup) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BriefWirelessLANGroup) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *BriefWirelessLANGroup) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *BriefWirelessLANGroup) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *BriefWirelessLANGroup) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *BriefWirelessLANGroup) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *BriefWirelessLANGroup) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *BriefWirelessLANGroup) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *BriefWirelessLANGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BriefWirelessLANGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BriefWirelessLANGroup) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *BriefWirelessLANGroup) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *BriefWirelessLANGroup) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *BriefWirelessLANGroup) SetSlug(v string) {
	o.Slug = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BriefWirelessLANGroup) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefWirelessLANGroup) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BriefWirelessLANGroup) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BriefWirelessLANGroup) SetDescription(v string) {
	o.Description = &v
}

// GetWirelesslanCount returns the WirelesslanCount field value
func (o *BriefWirelessLANGroup) GetWirelesslanCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WirelesslanCount
}

// GetWirelesslanCountOk returns a tuple with the WirelesslanCount field value
// and a boolean to check if the value has been set.
func (o *BriefWirelessLANGroup) GetWirelesslanCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WirelesslanCount, true
}

// SetWirelesslanCount sets field value
func (o *BriefWirelessLANGroup) SetWirelesslanCount(v int32) {
	o.WirelesslanCount = v
}

// GetDepth returns the Depth field value
func (o *BriefWirelessLANGroup) GetDepth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Depth
}

// GetDepthOk returns a tuple with the Depth field value
// and a boolean to check if the value has been set.
func (o *BriefWirelessLANGroup) GetDepthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Depth, true
}

// SetDepth sets field value
func (o *BriefWirelessLANGroup) SetDepth(v int32) {
	o.Depth = v
}

func (o BriefWirelessLANGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefWirelessLANGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["wirelesslan_count"] = o.WirelesslanCount
	toSerialize["_depth"] = o.Depth

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefWirelessLANGroup) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"name",
		"slug",
		"wirelesslan_count",
		"_depth",
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

	varBriefWirelessLANGroup := _BriefWirelessLANGroup{}

	err = json.Unmarshal(data, &varBriefWirelessLANGroup)

	if err != nil {
		return err
	}

	*o = BriefWirelessLANGroup(varBriefWirelessLANGroup)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "description")
		delete(additionalProperties, "wirelesslan_count")
		delete(additionalProperties, "_depth")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefWirelessLANGroup struct {
	value *BriefWirelessLANGroup
	isSet bool
}

func (v NullableBriefWirelessLANGroup) Get() *BriefWirelessLANGroup {
	return v.value
}

func (v *NullableBriefWirelessLANGroup) Set(val *BriefWirelessLANGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefWirelessLANGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefWirelessLANGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefWirelessLANGroup(val *BriefWirelessLANGroup) *NullableBriefWirelessLANGroup {
	return &NullableBriefWirelessLANGroup{value: val, isSet: true}
}

func (v NullableBriefWirelessLANGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefWirelessLANGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


