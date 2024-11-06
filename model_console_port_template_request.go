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

// checks if the ConsolePortTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConsolePortTemplateRequest{}

// ConsolePortTemplateRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type ConsolePortTemplateRequest struct {
	DeviceType NullableBriefDeviceTypeRequest `json:"device_type,omitempty"`
	ModuleType NullableBriefModuleTypeRequest `json:"module_type,omitempty"`
	// {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name string `json:"name"`
	// Physical label
	Label                *string               `json:"label,omitempty"`
	Type                 *ConsolePortTypeValue `json:"type,omitempty"`
	Description          *string               `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConsolePortTemplateRequest ConsolePortTemplateRequest

// NewConsolePortTemplateRequest instantiates a new ConsolePortTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsolePortTemplateRequest(name string) *ConsolePortTemplateRequest {
	this := ConsolePortTemplateRequest{}
	this.Name = name
	return &this
}

// NewConsolePortTemplateRequestWithDefaults instantiates a new ConsolePortTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsolePortTemplateRequestWithDefaults() *ConsolePortTemplateRequest {
	this := ConsolePortTemplateRequest{}
	return &this
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConsolePortTemplateRequest) GetDeviceType() BriefDeviceTypeRequest {
	if o == nil || IsNil(o.DeviceType.Get()) {
		var ret BriefDeviceTypeRequest
		return ret
	}
	return *o.DeviceType.Get()
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsolePortTemplateRequest) GetDeviceTypeOk() (*BriefDeviceTypeRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceType.Get(), o.DeviceType.IsSet()
}

// HasDeviceType returns a boolean if a field has been set.
func (o *ConsolePortTemplateRequest) HasDeviceType() bool {
	if o != nil && o.DeviceType.IsSet() {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given NullableBriefDeviceTypeRequest and assigns it to the DeviceType field.
func (o *ConsolePortTemplateRequest) SetDeviceType(v BriefDeviceTypeRequest) {
	o.DeviceType.Set(&v)
}

// SetDeviceTypeNil sets the value for DeviceType to be an explicit nil
func (o *ConsolePortTemplateRequest) SetDeviceTypeNil() {
	o.DeviceType.Set(nil)
}

// UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
func (o *ConsolePortTemplateRequest) UnsetDeviceType() {
	o.DeviceType.Unset()
}

// GetModuleType returns the ModuleType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConsolePortTemplateRequest) GetModuleType() BriefModuleTypeRequest {
	if o == nil || IsNil(o.ModuleType.Get()) {
		var ret BriefModuleTypeRequest
		return ret
	}
	return *o.ModuleType.Get()
}

// GetModuleTypeOk returns a tuple with the ModuleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsolePortTemplateRequest) GetModuleTypeOk() (*BriefModuleTypeRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModuleType.Get(), o.ModuleType.IsSet()
}

// HasModuleType returns a boolean if a field has been set.
func (o *ConsolePortTemplateRequest) HasModuleType() bool {
	if o != nil && o.ModuleType.IsSet() {
		return true
	}

	return false
}

// SetModuleType gets a reference to the given NullableBriefModuleTypeRequest and assigns it to the ModuleType field.
func (o *ConsolePortTemplateRequest) SetModuleType(v BriefModuleTypeRequest) {
	o.ModuleType.Set(&v)
}

// SetModuleTypeNil sets the value for ModuleType to be an explicit nil
func (o *ConsolePortTemplateRequest) SetModuleTypeNil() {
	o.ModuleType.Set(nil)
}

// UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
func (o *ConsolePortTemplateRequest) UnsetModuleType() {
	o.ModuleType.Unset()
}

// GetName returns the Name field value
func (o *ConsolePortTemplateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConsolePortTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConsolePortTemplateRequest) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ConsolePortTemplateRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsolePortTemplateRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ConsolePortTemplateRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ConsolePortTemplateRequest) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ConsolePortTemplateRequest) GetType() ConsolePortTypeValue {
	if o == nil || IsNil(o.Type) {
		var ret ConsolePortTypeValue
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsolePortTemplateRequest) GetTypeOk() (*ConsolePortTypeValue, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ConsolePortTemplateRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ConsolePortTypeValue and assigns it to the Type field.
func (o *ConsolePortTemplateRequest) SetType(v ConsolePortTypeValue) {
	o.Type = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConsolePortTemplateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsolePortTemplateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConsolePortTemplateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConsolePortTemplateRequest) SetDescription(v string) {
	o.Description = &v
}

func (o ConsolePortTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConsolePortTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DeviceType.IsSet() {
		toSerialize["device_type"] = o.DeviceType.Get()
	}
	if o.ModuleType.IsSet() {
		toSerialize["module_type"] = o.ModuleType.Get()
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConsolePortTemplateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varConsolePortTemplateRequest := _ConsolePortTemplateRequest{}

	err = json.Unmarshal(data, &varConsolePortTemplateRequest)

	if err != nil {
		return err
	}

	*o = ConsolePortTemplateRequest(varConsolePortTemplateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device_type")
		delete(additionalProperties, "module_type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "type")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConsolePortTemplateRequest struct {
	value *ConsolePortTemplateRequest
	isSet bool
}

func (v NullableConsolePortTemplateRequest) Get() *ConsolePortTemplateRequest {
	return v.value
}

func (v *NullableConsolePortTemplateRequest) Set(val *ConsolePortTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConsolePortTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConsolePortTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsolePortTemplateRequest(val *ConsolePortTemplateRequest) *NullableConsolePortTemplateRequest {
	return &NullableConsolePortTemplateRequest{value: val, isSet: true}
}

func (v NullableConsolePortTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsolePortTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
