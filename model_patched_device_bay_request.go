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

// checks if the PatchedDeviceBayRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedDeviceBayRequest{}

// PatchedDeviceBayRequest Adds support for custom fields and tags.
type PatchedDeviceBayRequest struct {
	Device *BriefDeviceRequest `json:"device,omitempty"`
	Name *string `json:"name,omitempty"`
	// Physical label
	Label *string `json:"label,omitempty"`
	Description *string `json:"description,omitempty"`
	InstalledDevice NullableBriefDeviceRequest `json:"installed_device,omitempty"`
	Tags []NestedTagRequest `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedDeviceBayRequest PatchedDeviceBayRequest

// NewPatchedDeviceBayRequest instantiates a new PatchedDeviceBayRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedDeviceBayRequest() *PatchedDeviceBayRequest {
	this := PatchedDeviceBayRequest{}
	return &this
}

// NewPatchedDeviceBayRequestWithDefaults instantiates a new PatchedDeviceBayRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedDeviceBayRequestWithDefaults() *PatchedDeviceBayRequest {
	this := PatchedDeviceBayRequest{}
	return &this
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *PatchedDeviceBayRequest) GetDevice() BriefDeviceRequest {
	if o == nil || IsNil(o.Device) {
		var ret BriefDeviceRequest
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDeviceBayRequest) GetDeviceOk() (*BriefDeviceRequest, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *PatchedDeviceBayRequest) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given BriefDeviceRequest and assigns it to the Device field.
func (o *PatchedDeviceBayRequest) SetDevice(v BriefDeviceRequest) {
	o.Device = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedDeviceBayRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDeviceBayRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedDeviceBayRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedDeviceBayRequest) SetName(v string) {
	o.Name = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PatchedDeviceBayRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDeviceBayRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PatchedDeviceBayRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PatchedDeviceBayRequest) SetLabel(v string) {
	o.Label = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedDeviceBayRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDeviceBayRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedDeviceBayRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedDeviceBayRequest) SetDescription(v string) {
	o.Description = &v
}

// GetInstalledDevice returns the InstalledDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedDeviceBayRequest) GetInstalledDevice() BriefDeviceRequest {
	if o == nil || IsNil(o.InstalledDevice.Get()) {
		var ret BriefDeviceRequest
		return ret
	}
	return *o.InstalledDevice.Get()
}

// GetInstalledDeviceOk returns a tuple with the InstalledDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedDeviceBayRequest) GetInstalledDeviceOk() (*BriefDeviceRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstalledDevice.Get(), o.InstalledDevice.IsSet()
}

// HasInstalledDevice returns a boolean if a field has been set.
func (o *PatchedDeviceBayRequest) HasInstalledDevice() bool {
	if o != nil && o.InstalledDevice.IsSet() {
		return true
	}

	return false
}

// SetInstalledDevice gets a reference to the given NullableBriefDeviceRequest and assigns it to the InstalledDevice field.
func (o *PatchedDeviceBayRequest) SetInstalledDevice(v BriefDeviceRequest) {
	o.InstalledDevice.Set(&v)
}
// SetInstalledDeviceNil sets the value for InstalledDevice to be an explicit nil
func (o *PatchedDeviceBayRequest) SetInstalledDeviceNil() {
	o.InstalledDevice.Set(nil)
}

// UnsetInstalledDevice ensures that no value is present for InstalledDevice, not even an explicit nil
func (o *PatchedDeviceBayRequest) UnsetInstalledDevice() {
	o.InstalledDevice.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedDeviceBayRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDeviceBayRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedDeviceBayRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedDeviceBayRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedDeviceBayRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDeviceBayRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedDeviceBayRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedDeviceBayRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PatchedDeviceBayRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedDeviceBayRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.InstalledDevice.IsSet() {
		toSerialize["installed_device"] = o.InstalledDevice.Get()
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

func (o *PatchedDeviceBayRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedDeviceBayRequest := _PatchedDeviceBayRequest{}

	err = json.Unmarshal(data, &varPatchedDeviceBayRequest)

	if err != nil {
		return err
	}

	*o = PatchedDeviceBayRequest(varPatchedDeviceBayRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "description")
		delete(additionalProperties, "installed_device")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedDeviceBayRequest struct {
	value *PatchedDeviceBayRequest
	isSet bool
}

func (v NullablePatchedDeviceBayRequest) Get() *PatchedDeviceBayRequest {
	return v.value
}

func (v *NullablePatchedDeviceBayRequest) Set(val *PatchedDeviceBayRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedDeviceBayRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedDeviceBayRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedDeviceBayRequest(val *PatchedDeviceBayRequest) *NullablePatchedDeviceBayRequest {
	return &NullablePatchedDeviceBayRequest{value: val, isSet: true}
}

func (v NullablePatchedDeviceBayRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedDeviceBayRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


