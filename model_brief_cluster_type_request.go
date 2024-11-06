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

// checks if the BriefClusterTypeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefClusterTypeRequest{}

// BriefClusterTypeRequest Adds support for custom fields and tags.
type BriefClusterTypeRequest struct {
	Name string `json:"name"`
	Slug string `json:"slug" validate:"regexp=^[-a-zA-Z0-9_]+$"`
	Description *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BriefClusterTypeRequest BriefClusterTypeRequest

// NewBriefClusterTypeRequest instantiates a new BriefClusterTypeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefClusterTypeRequest(name string, slug string) *BriefClusterTypeRequest {
	this := BriefClusterTypeRequest{}
	this.Name = name
	this.Slug = slug
	return &this
}

// NewBriefClusterTypeRequestWithDefaults instantiates a new BriefClusterTypeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefClusterTypeRequestWithDefaults() *BriefClusterTypeRequest {
	this := BriefClusterTypeRequest{}
	return &this
}

// GetName returns the Name field value
func (o *BriefClusterTypeRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BriefClusterTypeRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BriefClusterTypeRequest) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *BriefClusterTypeRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *BriefClusterTypeRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *BriefClusterTypeRequest) SetSlug(v string) {
	o.Slug = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BriefClusterTypeRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefClusterTypeRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BriefClusterTypeRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BriefClusterTypeRequest) SetDescription(v string) {
	o.Description = &v
}

func (o BriefClusterTypeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefClusterTypeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefClusterTypeRequest) UnmarshalJSON(data []byte) (err error) {
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

	varBriefClusterTypeRequest := _BriefClusterTypeRequest{}

	err = json.Unmarshal(data, &varBriefClusterTypeRequest)

	if err != nil {
		return err
	}

	*o = BriefClusterTypeRequest(varBriefClusterTypeRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefClusterTypeRequest struct {
	value *BriefClusterTypeRequest
	isSet bool
}

func (v NullableBriefClusterTypeRequest) Get() *BriefClusterTypeRequest {
	return v.value
}

func (v *NullableBriefClusterTypeRequest) Set(val *BriefClusterTypeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefClusterTypeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefClusterTypeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefClusterTypeRequest(val *BriefClusterTypeRequest) *NullableBriefClusterTypeRequest {
	return &NullableBriefClusterTypeRequest{value: val, isSet: true}
}

func (v NullableBriefClusterTypeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefClusterTypeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


