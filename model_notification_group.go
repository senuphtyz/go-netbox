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

// checks if the NotificationGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationGroup{}

// NotificationGroup Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type NotificationGroup struct {
	Id                   int32   `json:"id"`
	Url                  string  `json:"url"`
	Display              string  `json:"display"`
	DisplayUrl           string  `json:"display_url"`
	Name                 string  `json:"name"`
	Description          *string `json:"description,omitempty"`
	Groups               []Group `json:"groups,omitempty"`
	Users                []User  `json:"users,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NotificationGroup NotificationGroup

// NewNotificationGroup instantiates a new NotificationGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationGroup(id int32, url string, display string, displayUrl string, name string) *NotificationGroup {
	this := NotificationGroup{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.DisplayUrl = displayUrl
	this.Name = name
	return &this
}

// NewNotificationGroupWithDefaults instantiates a new NotificationGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationGroupWithDefaults() *NotificationGroup {
	this := NotificationGroup{}
	return &this
}

// GetId returns the Id field value
func (o *NotificationGroup) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NotificationGroup) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NotificationGroup) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NotificationGroup) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NotificationGroup) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NotificationGroup) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *NotificationGroup) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NotificationGroup) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NotificationGroup) SetDisplay(v string) {
	o.Display = v
}

// GetDisplayUrl returns the DisplayUrl field value
func (o *NotificationGroup) GetDisplayUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayUrl
}

// GetDisplayUrlOk returns a tuple with the DisplayUrl field value
// and a boolean to check if the value has been set.
func (o *NotificationGroup) GetDisplayUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayUrl, true
}

// SetDisplayUrl sets field value
func (o *NotificationGroup) SetDisplayUrl(v string) {
	o.DisplayUrl = v
}

// GetName returns the Name field value
func (o *NotificationGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NotificationGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NotificationGroup) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NotificationGroup) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGroup) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NotificationGroup) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NotificationGroup) SetDescription(v string) {
	o.Description = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *NotificationGroup) GetGroups() []Group {
	if o == nil || IsNil(o.Groups) {
		var ret []Group
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGroup) GetGroupsOk() ([]Group, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *NotificationGroup) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []Group and assigns it to the Groups field.
func (o *NotificationGroup) SetGroups(v []Group) {
	o.Groups = v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *NotificationGroup) GetUsers() []User {
	if o == nil || IsNil(o.Users) {
		var ret []User
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationGroup) GetUsersOk() ([]User, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *NotificationGroup) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []User and assigns it to the Users field.
func (o *NotificationGroup) SetUsers(v []User) {
	o.Users = v
}

func (o NotificationGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["display_url"] = o.DisplayUrl
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NotificationGroup) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"display_url",
		"name",
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

	varNotificationGroup := _NotificationGroup{}

	err = json.Unmarshal(data, &varNotificationGroup)

	if err != nil {
		return err
	}

	*o = NotificationGroup(varNotificationGroup)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "display_url")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "groups")
		delete(additionalProperties, "users")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNotificationGroup struct {
	value *NotificationGroup
	isSet bool
}

func (v NullableNotificationGroup) Get() *NotificationGroup {
	return v.value
}

func (v *NullableNotificationGroup) Set(val *NotificationGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationGroup(val *NotificationGroup) *NullableNotificationGroup {
	return &NullableNotificationGroup{value: val, isSet: true}
}

func (v NullableNotificationGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
