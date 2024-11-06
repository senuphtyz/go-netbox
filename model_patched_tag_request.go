/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.1.4 (4.1)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the PatchedTagRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedTagRequest{}

// PatchedTagRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedTagRequest struct {
	Name *string `json:"name,omitempty"`
	Slug *string `json:"slug,omitempty"`
	Color *string `json:"color,omitempty"`
	Description *string `json:"description,omitempty"`
	ObjectTypes []string `json:"object_types,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedTagRequest PatchedTagRequest

// NewPatchedTagRequest instantiates a new PatchedTagRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedTagRequest() *PatchedTagRequest {
	this := PatchedTagRequest{}
	return &this
}

// NewPatchedTagRequestWithDefaults instantiates a new PatchedTagRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedTagRequestWithDefaults() *PatchedTagRequest {
	this := PatchedTagRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedTagRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedTagRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedTagRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedTagRequest) SetName(v string) {
	o.Name = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *PatchedTagRequest) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedTagRequest) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *PatchedTagRequest) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *PatchedTagRequest) SetSlug(v string) {
	o.Slug = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *PatchedTagRequest) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedTagRequest) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *PatchedTagRequest) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *PatchedTagRequest) SetColor(v string) {
	o.Color = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedTagRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedTagRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedTagRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedTagRequest) SetDescription(v string) {
	o.Description = &v
}

// GetObjectTypes returns the ObjectTypes field value if set, zero value otherwise.
func (o *PatchedTagRequest) GetObjectTypes() []string {
	if o == nil || IsNil(o.ObjectTypes) {
		var ret []string
		return ret
	}
	return o.ObjectTypes
}

// GetObjectTypesOk returns a tuple with the ObjectTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedTagRequest) GetObjectTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.ObjectTypes) {
		return nil, false
	}
	return o.ObjectTypes, true
}

// HasObjectTypes returns a boolean if a field has been set.
func (o *PatchedTagRequest) HasObjectTypes() bool {
	if o != nil && !IsNil(o.ObjectTypes) {
		return true
	}

	return false
}

// SetObjectTypes gets a reference to the given []string and assigns it to the ObjectTypes field.
func (o *PatchedTagRequest) SetObjectTypes(v []string) {
	o.ObjectTypes = v
}

func (o PatchedTagRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedTagRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ObjectTypes) {
		toSerialize["object_types"] = o.ObjectTypes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedTagRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedTagRequest := _PatchedTagRequest{}

	err = json.Unmarshal(data, &varPatchedTagRequest)

	if err != nil {
		return err
	}

	*o = PatchedTagRequest(varPatchedTagRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "color")
		delete(additionalProperties, "description")
		delete(additionalProperties, "object_types")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedTagRequest struct {
	value *PatchedTagRequest
	isSet bool
}

func (v NullablePatchedTagRequest) Get() *PatchedTagRequest {
	return v.value
}

func (v *NullablePatchedTagRequest) Set(val *PatchedTagRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedTagRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedTagRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedTagRequest(val *PatchedTagRequest) *NullablePatchedTagRequest {
	return &NullablePatchedTagRequest{value: val, isSet: true}
}

func (v NullablePatchedTagRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedTagRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


