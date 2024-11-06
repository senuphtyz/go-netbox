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

// checks if the PatchedWritableConsolePortRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableConsolePortRequest{}

// PatchedWritableConsolePortRequest Adds support for custom fields and tags.
type PatchedWritableConsolePortRequest struct {
	Device *BriefDeviceRequest `json:"device,omitempty"`
	Module NullableBriefModuleRequest `json:"module,omitempty"`
	Name *string `json:"name,omitempty"`
	// Physical label
	Label *string `json:"label,omitempty"`
	Type *PatchedWritableConsolePortRequestType `json:"type,omitempty"`
	Speed NullablePatchedWritableConsolePortRequestSpeed `json:"speed,omitempty"`
	Description *string `json:"description,omitempty"`
	// Treat as if a cable is connected
	MarkConnected *bool `json:"mark_connected,omitempty"`
	Tags []NestedTagRequest `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritableConsolePortRequest PatchedWritableConsolePortRequest

// NewPatchedWritableConsolePortRequest instantiates a new PatchedWritableConsolePortRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableConsolePortRequest() *PatchedWritableConsolePortRequest {
	this := PatchedWritableConsolePortRequest{}
	return &this
}

// NewPatchedWritableConsolePortRequestWithDefaults instantiates a new PatchedWritableConsolePortRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableConsolePortRequestWithDefaults() *PatchedWritableConsolePortRequest {
	this := PatchedWritableConsolePortRequest{}
	return &this
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *PatchedWritableConsolePortRequest) GetDevice() BriefDeviceRequest {
	if o == nil || IsNil(o.Device) {
		var ret BriefDeviceRequest
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConsolePortRequest) GetDeviceOk() (*BriefDeviceRequest, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *PatchedWritableConsolePortRequest) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given BriefDeviceRequest and assigns it to the Device field.
func (o *PatchedWritableConsolePortRequest) SetDevice(v BriefDeviceRequest) {
	o.Device = &v
}

// GetModule returns the Module field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableConsolePortRequest) GetModule() BriefModuleRequest {
	if o == nil || IsNil(o.Module.Get()) {
		var ret BriefModuleRequest
		return ret
	}
	return *o.Module.Get()
}

// GetModuleOk returns a tuple with the Module field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableConsolePortRequest) GetModuleOk() (*BriefModuleRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Module.Get(), o.Module.IsSet()
}

// HasModule returns a boolean if a field has been set.
func (o *PatchedWritableConsolePortRequest) HasModule() bool {
	if o != nil && o.Module.IsSet() {
		return true
	}

	return false
}

// SetModule gets a reference to the given NullableBriefModuleRequest and assigns it to the Module field.
func (o *PatchedWritableConsolePortRequest) SetModule(v BriefModuleRequest) {
	o.Module.Set(&v)
}
// SetModuleNil sets the value for Module to be an explicit nil
func (o *PatchedWritableConsolePortRequest) SetModuleNil() {
	o.Module.Set(nil)
}

// UnsetModule ensures that no value is present for Module, not even an explicit nil
func (o *PatchedWritableConsolePortRequest) UnsetModule() {
	o.Module.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWritableConsolePortRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConsolePortRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWritableConsolePortRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWritableConsolePortRequest) SetName(v string) {
	o.Name = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PatchedWritableConsolePortRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConsolePortRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PatchedWritableConsolePortRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PatchedWritableConsolePortRequest) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PatchedWritableConsolePortRequest) GetType() PatchedWritableConsolePortRequestType {
	if o == nil || IsNil(o.Type) {
		var ret PatchedWritableConsolePortRequestType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConsolePortRequest) GetTypeOk() (*PatchedWritableConsolePortRequestType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PatchedWritableConsolePortRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given PatchedWritableConsolePortRequestType and assigns it to the Type field.
func (o *PatchedWritableConsolePortRequest) SetType(v PatchedWritableConsolePortRequestType) {
	o.Type = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableConsolePortRequest) GetSpeed() PatchedWritableConsolePortRequestSpeed {
	if o == nil || IsNil(o.Speed.Get()) {
		var ret PatchedWritableConsolePortRequestSpeed
		return ret
	}
	return *o.Speed.Get()
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableConsolePortRequest) GetSpeedOk() (*PatchedWritableConsolePortRequestSpeed, bool) {
	if o == nil {
		return nil, false
	}
	return o.Speed.Get(), o.Speed.IsSet()
}

// HasSpeed returns a boolean if a field has been set.
func (o *PatchedWritableConsolePortRequest) HasSpeed() bool {
	if o != nil && o.Speed.IsSet() {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given NullablePatchedWritableConsolePortRequestSpeed and assigns it to the Speed field.
func (o *PatchedWritableConsolePortRequest) SetSpeed(v PatchedWritableConsolePortRequestSpeed) {
	o.Speed.Set(&v)
}
// SetSpeedNil sets the value for Speed to be an explicit nil
func (o *PatchedWritableConsolePortRequest) SetSpeedNil() {
	o.Speed.Set(nil)
}

// UnsetSpeed ensures that no value is present for Speed, not even an explicit nil
func (o *PatchedWritableConsolePortRequest) UnsetSpeed() {
	o.Speed.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableConsolePortRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConsolePortRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableConsolePortRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableConsolePortRequest) SetDescription(v string) {
	o.Description = &v
}

// GetMarkConnected returns the MarkConnected field value if set, zero value otherwise.
func (o *PatchedWritableConsolePortRequest) GetMarkConnected() bool {
	if o == nil || IsNil(o.MarkConnected) {
		var ret bool
		return ret
	}
	return *o.MarkConnected
}

// GetMarkConnectedOk returns a tuple with the MarkConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConsolePortRequest) GetMarkConnectedOk() (*bool, bool) {
	if o == nil || IsNil(o.MarkConnected) {
		return nil, false
	}
	return o.MarkConnected, true
}

// HasMarkConnected returns a boolean if a field has been set.
func (o *PatchedWritableConsolePortRequest) HasMarkConnected() bool {
	if o != nil && !IsNil(o.MarkConnected) {
		return true
	}

	return false
}

// SetMarkConnected gets a reference to the given bool and assigns it to the MarkConnected field.
func (o *PatchedWritableConsolePortRequest) SetMarkConnected(v bool) {
	o.MarkConnected = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritableConsolePortRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConsolePortRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritableConsolePortRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedWritableConsolePortRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritableConsolePortRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConsolePortRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritableConsolePortRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedWritableConsolePortRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PatchedWritableConsolePortRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableConsolePortRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if o.Module.IsSet() {
		toSerialize["module"] = o.Module.Get()
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Speed.IsSet() {
		toSerialize["speed"] = o.Speed.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.MarkConnected) {
		toSerialize["mark_connected"] = o.MarkConnected
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

func (o *PatchedWritableConsolePortRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedWritableConsolePortRequest := _PatchedWritableConsolePortRequest{}

	err = json.Unmarshal(data, &varPatchedWritableConsolePortRequest)

	if err != nil {
		return err
	}

	*o = PatchedWritableConsolePortRequest(varPatchedWritableConsolePortRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device")
		delete(additionalProperties, "module")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "type")
		delete(additionalProperties, "speed")
		delete(additionalProperties, "description")
		delete(additionalProperties, "mark_connected")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritableConsolePortRequest struct {
	value *PatchedWritableConsolePortRequest
	isSet bool
}

func (v NullablePatchedWritableConsolePortRequest) Get() *PatchedWritableConsolePortRequest {
	return v.value
}

func (v *NullablePatchedWritableConsolePortRequest) Set(val *PatchedWritableConsolePortRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableConsolePortRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableConsolePortRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableConsolePortRequest(val *PatchedWritableConsolePortRequest) *NullablePatchedWritableConsolePortRequest {
	return &NullablePatchedWritableConsolePortRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableConsolePortRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableConsolePortRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


