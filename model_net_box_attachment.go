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

// checks if the NetBoxAttachment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetBoxAttachment{}

// NetBoxAttachment Adds support for custom fields and tags.
type NetBoxAttachment struct {
	Id int32 `json:"id"`
	Url string `json:"url"`
	Display string `json:"display"`
	ObjectType string `json:"object_type"`
	ObjectId int64 `json:"object_id"`
	Parent string `json:"parent"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	File string `json:"file"`
	Created NullableTime `json:"created"`
	LastUpdated NullableTime `json:"last_updated"`
	Comments *string `json:"comments,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetBoxAttachment NetBoxAttachment

// NewNetBoxAttachment instantiates a new NetBoxAttachment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetBoxAttachment(id int32, url string, display string, objectType string, objectId int64, parent string, file string, created NullableTime, lastUpdated NullableTime) *NetBoxAttachment {
	this := NetBoxAttachment{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.ObjectType = objectType
	this.ObjectId = objectId
	this.Parent = parent
	this.File = file
	this.Created = created
	this.LastUpdated = lastUpdated
	return &this
}

// NewNetBoxAttachmentWithDefaults instantiates a new NetBoxAttachment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetBoxAttachmentWithDefaults() *NetBoxAttachment {
	this := NetBoxAttachment{}
	return &this
}

// GetId returns the Id field value
func (o *NetBoxAttachment) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NetBoxAttachment) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NetBoxAttachment) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NetBoxAttachment) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NetBoxAttachment) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NetBoxAttachment) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *NetBoxAttachment) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NetBoxAttachment) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NetBoxAttachment) SetDisplay(v string) {
	o.Display = v
}

// GetObjectType returns the ObjectType field value
func (o *NetBoxAttachment) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NetBoxAttachment) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NetBoxAttachment) SetObjectType(v string) {
	o.ObjectType = v
}

// GetObjectId returns the ObjectId field value
func (o *NetBoxAttachment) GetObjectId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value
// and a boolean to check if the value has been set.
func (o *NetBoxAttachment) GetObjectIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectId, true
}

// SetObjectId sets field value
func (o *NetBoxAttachment) SetObjectId(v int64) {
	o.ObjectId = v
}

// GetParent returns the Parent field value
func (o *NetBoxAttachment) GetParent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Parent
}

// GetParentOk returns a tuple with the Parent field value
// and a boolean to check if the value has been set.
func (o *NetBoxAttachment) GetParentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parent, true
}

// SetParent sets field value
func (o *NetBoxAttachment) SetParent(v string) {
	o.Parent = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetBoxAttachment) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetBoxAttachment) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetBoxAttachment) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetBoxAttachment) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NetBoxAttachment) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetBoxAttachment) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NetBoxAttachment) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NetBoxAttachment) SetDescription(v string) {
	o.Description = &v
}

// GetFile returns the File field value
func (o *NetBoxAttachment) GetFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.File
}

// GetFileOk returns a tuple with the File field value
// and a boolean to check if the value has been set.
func (o *NetBoxAttachment) GetFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.File, true
}

// SetFile sets field value
func (o *NetBoxAttachment) SetFile(v string) {
	o.File = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *NetBoxAttachment) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetBoxAttachment) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *NetBoxAttachment) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *NetBoxAttachment) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetBoxAttachment) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *NetBoxAttachment) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *NetBoxAttachment) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetBoxAttachment) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *NetBoxAttachment) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *NetBoxAttachment) SetComments(v string) {
	o.Comments = &v
}

func (o NetBoxAttachment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetBoxAttachment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["object_type"] = o.ObjectType
	toSerialize["object_id"] = o.ObjectId
	toSerialize["parent"] = o.Parent
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["file"] = o.File
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NetBoxAttachment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"object_type",
		"object_id",
		"parent",
		"file",
		"created",
		"last_updated",
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

	varNetBoxAttachment := _NetBoxAttachment{}

	err = json.Unmarshal(data, &varNetBoxAttachment)

	if err != nil {
		return err
	}

	*o = NetBoxAttachment(varNetBoxAttachment)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "object_id")
		delete(additionalProperties, "parent")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "file")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		delete(additionalProperties, "comments")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetBoxAttachment struct {
	value *NetBoxAttachment
	isSet bool
}

func (v NullableNetBoxAttachment) Get() *NetBoxAttachment {
	return v.value
}

func (v *NullableNetBoxAttachment) Set(val *NetBoxAttachment) {
	v.value = val
	v.isSet = true
}

func (v NullableNetBoxAttachment) IsSet() bool {
	return v.isSet
}

func (v *NullableNetBoxAttachment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetBoxAttachment(val *NetBoxAttachment) *NullableNetBoxAttachment {
	return &NullableNetBoxAttachment{value: val, isSet: true}
}

func (v NullableNetBoxAttachment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetBoxAttachment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


