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

// checks if the InterfaceTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InterfaceTemplateRequest{}

// InterfaceTemplateRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type InterfaceTemplateRequest struct {
	DeviceType NullableBriefDeviceTypeRequest `json:"device_type,omitempty"`
	ModuleType NullableBriefModuleTypeRequest `json:"module_type,omitempty"`
	// {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name string `json:"name"`
	// Physical label
	Label                *string                                 `json:"label,omitempty"`
	Type                 InterfaceTypeValue                      `json:"type"`
	Enabled              *bool                                   `json:"enabled,omitempty"`
	MgmtOnly             *bool                                   `json:"mgmt_only,omitempty"`
	Description          *string                                 `json:"description,omitempty"`
	Bridge               NullableNestedInterfaceTemplateRequest  `json:"bridge,omitempty"`
	PoeMode              NullableInterfaceTemplateRequestPoeMode `json:"poe_mode,omitempty"`
	PoeType              NullableInterfaceTemplateRequestPoeType `json:"poe_type,omitempty"`
	RfRole               NullableInterfaceTemplateRequestRfRole  `json:"rf_role,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InterfaceTemplateRequest InterfaceTemplateRequest

// NewInterfaceTemplateRequest instantiates a new InterfaceTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterfaceTemplateRequest(name string, type_ InterfaceTypeValue) *InterfaceTemplateRequest {
	this := InterfaceTemplateRequest{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewInterfaceTemplateRequestWithDefaults instantiates a new InterfaceTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterfaceTemplateRequestWithDefaults() *InterfaceTemplateRequest {
	this := InterfaceTemplateRequest{}
	return &this
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InterfaceTemplateRequest) GetDeviceType() BriefDeviceTypeRequest {
	if o == nil || IsNil(o.DeviceType.Get()) {
		var ret BriefDeviceTypeRequest
		return ret
	}
	return *o.DeviceType.Get()
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InterfaceTemplateRequest) GetDeviceTypeOk() (*BriefDeviceTypeRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceType.Get(), o.DeviceType.IsSet()
}

// HasDeviceType returns a boolean if a field has been set.
func (o *InterfaceTemplateRequest) HasDeviceType() bool {
	if o != nil && o.DeviceType.IsSet() {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given NullableBriefDeviceTypeRequest and assigns it to the DeviceType field.
func (o *InterfaceTemplateRequest) SetDeviceType(v BriefDeviceTypeRequest) {
	o.DeviceType.Set(&v)
}

// SetDeviceTypeNil sets the value for DeviceType to be an explicit nil
func (o *InterfaceTemplateRequest) SetDeviceTypeNil() {
	o.DeviceType.Set(nil)
}

// UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
func (o *InterfaceTemplateRequest) UnsetDeviceType() {
	o.DeviceType.Unset()
}

// GetModuleType returns the ModuleType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InterfaceTemplateRequest) GetModuleType() BriefModuleTypeRequest {
	if o == nil || IsNil(o.ModuleType.Get()) {
		var ret BriefModuleTypeRequest
		return ret
	}
	return *o.ModuleType.Get()
}

// GetModuleTypeOk returns a tuple with the ModuleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InterfaceTemplateRequest) GetModuleTypeOk() (*BriefModuleTypeRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModuleType.Get(), o.ModuleType.IsSet()
}

// HasModuleType returns a boolean if a field has been set.
func (o *InterfaceTemplateRequest) HasModuleType() bool {
	if o != nil && o.ModuleType.IsSet() {
		return true
	}

	return false
}

// SetModuleType gets a reference to the given NullableBriefModuleTypeRequest and assigns it to the ModuleType field.
func (o *InterfaceTemplateRequest) SetModuleType(v BriefModuleTypeRequest) {
	o.ModuleType.Set(&v)
}

// SetModuleTypeNil sets the value for ModuleType to be an explicit nil
func (o *InterfaceTemplateRequest) SetModuleTypeNil() {
	o.ModuleType.Set(nil)
}

// UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
func (o *InterfaceTemplateRequest) UnsetModuleType() {
	o.ModuleType.Unset()
}

// GetName returns the Name field value
func (o *InterfaceTemplateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InterfaceTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InterfaceTemplateRequest) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *InterfaceTemplateRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceTemplateRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *InterfaceTemplateRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *InterfaceTemplateRequest) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value
func (o *InterfaceTemplateRequest) GetType() InterfaceTypeValue {
	if o == nil {
		var ret InterfaceTypeValue
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InterfaceTemplateRequest) GetTypeOk() (*InterfaceTypeValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InterfaceTemplateRequest) SetType(v InterfaceTypeValue) {
	o.Type = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InterfaceTemplateRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceTemplateRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InterfaceTemplateRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InterfaceTemplateRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMgmtOnly returns the MgmtOnly field value if set, zero value otherwise.
func (o *InterfaceTemplateRequest) GetMgmtOnly() bool {
	if o == nil || IsNil(o.MgmtOnly) {
		var ret bool
		return ret
	}
	return *o.MgmtOnly
}

// GetMgmtOnlyOk returns a tuple with the MgmtOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceTemplateRequest) GetMgmtOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.MgmtOnly) {
		return nil, false
	}
	return o.MgmtOnly, true
}

// HasMgmtOnly returns a boolean if a field has been set.
func (o *InterfaceTemplateRequest) HasMgmtOnly() bool {
	if o != nil && !IsNil(o.MgmtOnly) {
		return true
	}

	return false
}

// SetMgmtOnly gets a reference to the given bool and assigns it to the MgmtOnly field.
func (o *InterfaceTemplateRequest) SetMgmtOnly(v bool) {
	o.MgmtOnly = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InterfaceTemplateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceTemplateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InterfaceTemplateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InterfaceTemplateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetBridge returns the Bridge field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InterfaceTemplateRequest) GetBridge() NestedInterfaceTemplateRequest {
	if o == nil || IsNil(o.Bridge.Get()) {
		var ret NestedInterfaceTemplateRequest
		return ret
	}
	return *o.Bridge.Get()
}

// GetBridgeOk returns a tuple with the Bridge field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InterfaceTemplateRequest) GetBridgeOk() (*NestedInterfaceTemplateRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bridge.Get(), o.Bridge.IsSet()
}

// HasBridge returns a boolean if a field has been set.
func (o *InterfaceTemplateRequest) HasBridge() bool {
	if o != nil && o.Bridge.IsSet() {
		return true
	}

	return false
}

// SetBridge gets a reference to the given NullableNestedInterfaceTemplateRequest and assigns it to the Bridge field.
func (o *InterfaceTemplateRequest) SetBridge(v NestedInterfaceTemplateRequest) {
	o.Bridge.Set(&v)
}

// SetBridgeNil sets the value for Bridge to be an explicit nil
func (o *InterfaceTemplateRequest) SetBridgeNil() {
	o.Bridge.Set(nil)
}

// UnsetBridge ensures that no value is present for Bridge, not even an explicit nil
func (o *InterfaceTemplateRequest) UnsetBridge() {
	o.Bridge.Unset()
}

// GetPoeMode returns the PoeMode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InterfaceTemplateRequest) GetPoeMode() InterfaceTemplateRequestPoeMode {
	if o == nil || IsNil(o.PoeMode.Get()) {
		var ret InterfaceTemplateRequestPoeMode
		return ret
	}
	return *o.PoeMode.Get()
}

// GetPoeModeOk returns a tuple with the PoeMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InterfaceTemplateRequest) GetPoeModeOk() (*InterfaceTemplateRequestPoeMode, bool) {
	if o == nil {
		return nil, false
	}
	return o.PoeMode.Get(), o.PoeMode.IsSet()
}

// HasPoeMode returns a boolean if a field has been set.
func (o *InterfaceTemplateRequest) HasPoeMode() bool {
	if o != nil && o.PoeMode.IsSet() {
		return true
	}

	return false
}

// SetPoeMode gets a reference to the given NullableInterfaceTemplateRequestPoeMode and assigns it to the PoeMode field.
func (o *InterfaceTemplateRequest) SetPoeMode(v InterfaceTemplateRequestPoeMode) {
	o.PoeMode.Set(&v)
}

// SetPoeModeNil sets the value for PoeMode to be an explicit nil
func (o *InterfaceTemplateRequest) SetPoeModeNil() {
	o.PoeMode.Set(nil)
}

// UnsetPoeMode ensures that no value is present for PoeMode, not even an explicit nil
func (o *InterfaceTemplateRequest) UnsetPoeMode() {
	o.PoeMode.Unset()
}

// GetPoeType returns the PoeType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InterfaceTemplateRequest) GetPoeType() InterfaceTemplateRequestPoeType {
	if o == nil || IsNil(o.PoeType.Get()) {
		var ret InterfaceTemplateRequestPoeType
		return ret
	}
	return *o.PoeType.Get()
}

// GetPoeTypeOk returns a tuple with the PoeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InterfaceTemplateRequest) GetPoeTypeOk() (*InterfaceTemplateRequestPoeType, bool) {
	if o == nil {
		return nil, false
	}
	return o.PoeType.Get(), o.PoeType.IsSet()
}

// HasPoeType returns a boolean if a field has been set.
func (o *InterfaceTemplateRequest) HasPoeType() bool {
	if o != nil && o.PoeType.IsSet() {
		return true
	}

	return false
}

// SetPoeType gets a reference to the given NullableInterfaceTemplateRequestPoeType and assigns it to the PoeType field.
func (o *InterfaceTemplateRequest) SetPoeType(v InterfaceTemplateRequestPoeType) {
	o.PoeType.Set(&v)
}

// SetPoeTypeNil sets the value for PoeType to be an explicit nil
func (o *InterfaceTemplateRequest) SetPoeTypeNil() {
	o.PoeType.Set(nil)
}

// UnsetPoeType ensures that no value is present for PoeType, not even an explicit nil
func (o *InterfaceTemplateRequest) UnsetPoeType() {
	o.PoeType.Unset()
}

// GetRfRole returns the RfRole field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InterfaceTemplateRequest) GetRfRole() InterfaceTemplateRequestRfRole {
	if o == nil || IsNil(o.RfRole.Get()) {
		var ret InterfaceTemplateRequestRfRole
		return ret
	}
	return *o.RfRole.Get()
}

// GetRfRoleOk returns a tuple with the RfRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InterfaceTemplateRequest) GetRfRoleOk() (*InterfaceTemplateRequestRfRole, bool) {
	if o == nil {
		return nil, false
	}
	return o.RfRole.Get(), o.RfRole.IsSet()
}

// HasRfRole returns a boolean if a field has been set.
func (o *InterfaceTemplateRequest) HasRfRole() bool {
	if o != nil && o.RfRole.IsSet() {
		return true
	}

	return false
}

// SetRfRole gets a reference to the given NullableInterfaceTemplateRequestRfRole and assigns it to the RfRole field.
func (o *InterfaceTemplateRequest) SetRfRole(v InterfaceTemplateRequestRfRole) {
	o.RfRole.Set(&v)
}

// SetRfRoleNil sets the value for RfRole to be an explicit nil
func (o *InterfaceTemplateRequest) SetRfRoleNil() {
	o.RfRole.Set(nil)
}

// UnsetRfRole ensures that no value is present for RfRole, not even an explicit nil
func (o *InterfaceTemplateRequest) UnsetRfRole() {
	o.RfRole.Unset()
}

func (o InterfaceTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InterfaceTemplateRequest) ToMap() (map[string]interface{}, error) {
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
	toSerialize["type"] = o.Type
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.MgmtOnly) {
		toSerialize["mgmt_only"] = o.MgmtOnly
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.Bridge.IsSet() {
		toSerialize["bridge"] = o.Bridge.Get()
	}
	if o.PoeMode.IsSet() {
		toSerialize["poe_mode"] = o.PoeMode.Get()
	}
	if o.PoeType.IsSet() {
		toSerialize["poe_type"] = o.PoeType.Get()
	}
	if o.RfRole.IsSet() {
		toSerialize["rf_role"] = o.RfRole.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InterfaceTemplateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
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

	varInterfaceTemplateRequest := _InterfaceTemplateRequest{}

	err = json.Unmarshal(data, &varInterfaceTemplateRequest)

	if err != nil {
		return err
	}

	*o = InterfaceTemplateRequest(varInterfaceTemplateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device_type")
		delete(additionalProperties, "module_type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "type")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "mgmt_only")
		delete(additionalProperties, "description")
		delete(additionalProperties, "bridge")
		delete(additionalProperties, "poe_mode")
		delete(additionalProperties, "poe_type")
		delete(additionalProperties, "rf_role")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInterfaceTemplateRequest struct {
	value *InterfaceTemplateRequest
	isSet bool
}

func (v NullableInterfaceTemplateRequest) Get() *InterfaceTemplateRequest {
	return v.value
}

func (v *NullableInterfaceTemplateRequest) Set(val *InterfaceTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceTemplateRequest(val *InterfaceTemplateRequest) *NullableInterfaceTemplateRequest {
	return &NullableInterfaceTemplateRequest{value: val, isSet: true}
}

func (v NullableInterfaceTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
