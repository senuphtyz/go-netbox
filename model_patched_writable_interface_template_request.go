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

// checks if the PatchedWritableInterfaceTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableInterfaceTemplateRequest{}

// PatchedWritableInterfaceTemplateRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedWritableInterfaceTemplateRequest struct {
	DeviceType NullableBriefDeviceTypeRequest `json:"device_type,omitempty"`
	ModuleType NullableBriefModuleTypeRequest `json:"module_type,omitempty"`
	// {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name *string `json:"name,omitempty"`
	// Physical label
	Label *string `json:"label,omitempty"`
	Type *InterfaceTypeValue `json:"type,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	MgmtOnly *bool `json:"mgmt_only,omitempty"`
	Description *string `json:"description,omitempty"`
	Bridge NullableInt32 `json:"bridge,omitempty"`
	PoeMode *InterfacePoeModeValue `json:"poe_mode,omitempty"`
	PoeType *InterfacePoeTypeValue `json:"poe_type,omitempty"`
	RfRole *WirelessRole `json:"rf_role,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritableInterfaceTemplateRequest PatchedWritableInterfaceTemplateRequest

// NewPatchedWritableInterfaceTemplateRequest instantiates a new PatchedWritableInterfaceTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableInterfaceTemplateRequest() *PatchedWritableInterfaceTemplateRequest {
	this := PatchedWritableInterfaceTemplateRequest{}
	return &this
}

// NewPatchedWritableInterfaceTemplateRequestWithDefaults instantiates a new PatchedWritableInterfaceTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableInterfaceTemplateRequestWithDefaults() *PatchedWritableInterfaceTemplateRequest {
	this := PatchedWritableInterfaceTemplateRequest{}
	return &this
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableInterfaceTemplateRequest) GetDeviceType() BriefDeviceTypeRequest {
	if o == nil || IsNil(o.DeviceType.Get()) {
		var ret BriefDeviceTypeRequest
		return ret
	}
	return *o.DeviceType.Get()
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableInterfaceTemplateRequest) GetDeviceTypeOk() (*BriefDeviceTypeRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceType.Get(), o.DeviceType.IsSet()
}

// HasDeviceType returns a boolean if a field has been set.
func (o *PatchedWritableInterfaceTemplateRequest) HasDeviceType() bool {
	if o != nil && o.DeviceType.IsSet() {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given NullableBriefDeviceTypeRequest and assigns it to the DeviceType field.
func (o *PatchedWritableInterfaceTemplateRequest) SetDeviceType(v BriefDeviceTypeRequest) {
	o.DeviceType.Set(&v)
}
// SetDeviceTypeNil sets the value for DeviceType to be an explicit nil
func (o *PatchedWritableInterfaceTemplateRequest) SetDeviceTypeNil() {
	o.DeviceType.Set(nil)
}

// UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
func (o *PatchedWritableInterfaceTemplateRequest) UnsetDeviceType() {
	o.DeviceType.Unset()
}

// GetModuleType returns the ModuleType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableInterfaceTemplateRequest) GetModuleType() BriefModuleTypeRequest {
	if o == nil || IsNil(o.ModuleType.Get()) {
		var ret BriefModuleTypeRequest
		return ret
	}
	return *o.ModuleType.Get()
}

// GetModuleTypeOk returns a tuple with the ModuleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableInterfaceTemplateRequest) GetModuleTypeOk() (*BriefModuleTypeRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModuleType.Get(), o.ModuleType.IsSet()
}

// HasModuleType returns a boolean if a field has been set.
func (o *PatchedWritableInterfaceTemplateRequest) HasModuleType() bool {
	if o != nil && o.ModuleType.IsSet() {
		return true
	}

	return false
}

// SetModuleType gets a reference to the given NullableBriefModuleTypeRequest and assigns it to the ModuleType field.
func (o *PatchedWritableInterfaceTemplateRequest) SetModuleType(v BriefModuleTypeRequest) {
	o.ModuleType.Set(&v)
}
// SetModuleTypeNil sets the value for ModuleType to be an explicit nil
func (o *PatchedWritableInterfaceTemplateRequest) SetModuleTypeNil() {
	o.ModuleType.Set(nil)
}

// UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
func (o *PatchedWritableInterfaceTemplateRequest) UnsetModuleType() {
	o.ModuleType.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWritableInterfaceTemplateRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableInterfaceTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWritableInterfaceTemplateRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWritableInterfaceTemplateRequest) SetName(v string) {
	o.Name = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PatchedWritableInterfaceTemplateRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableInterfaceTemplateRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PatchedWritableInterfaceTemplateRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PatchedWritableInterfaceTemplateRequest) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PatchedWritableInterfaceTemplateRequest) GetType() InterfaceTypeValue {
	if o == nil || IsNil(o.Type) {
		var ret InterfaceTypeValue
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableInterfaceTemplateRequest) GetTypeOk() (*InterfaceTypeValue, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PatchedWritableInterfaceTemplateRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given InterfaceTypeValue and assigns it to the Type field.
func (o *PatchedWritableInterfaceTemplateRequest) SetType(v InterfaceTypeValue) {
	o.Type = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PatchedWritableInterfaceTemplateRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableInterfaceTemplateRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PatchedWritableInterfaceTemplateRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PatchedWritableInterfaceTemplateRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMgmtOnly returns the MgmtOnly field value if set, zero value otherwise.
func (o *PatchedWritableInterfaceTemplateRequest) GetMgmtOnly() bool {
	if o == nil || IsNil(o.MgmtOnly) {
		var ret bool
		return ret
	}
	return *o.MgmtOnly
}

// GetMgmtOnlyOk returns a tuple with the MgmtOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableInterfaceTemplateRequest) GetMgmtOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.MgmtOnly) {
		return nil, false
	}
	return o.MgmtOnly, true
}

// HasMgmtOnly returns a boolean if a field has been set.
func (o *PatchedWritableInterfaceTemplateRequest) HasMgmtOnly() bool {
	if o != nil && !IsNil(o.MgmtOnly) {
		return true
	}

	return false
}

// SetMgmtOnly gets a reference to the given bool and assigns it to the MgmtOnly field.
func (o *PatchedWritableInterfaceTemplateRequest) SetMgmtOnly(v bool) {
	o.MgmtOnly = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableInterfaceTemplateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableInterfaceTemplateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableInterfaceTemplateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableInterfaceTemplateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetBridge returns the Bridge field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableInterfaceTemplateRequest) GetBridge() int32 {
	if o == nil || IsNil(o.Bridge.Get()) {
		var ret int32
		return ret
	}
	return *o.Bridge.Get()
}

// GetBridgeOk returns a tuple with the Bridge field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableInterfaceTemplateRequest) GetBridgeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bridge.Get(), o.Bridge.IsSet()
}

// HasBridge returns a boolean if a field has been set.
func (o *PatchedWritableInterfaceTemplateRequest) HasBridge() bool {
	if o != nil && o.Bridge.IsSet() {
		return true
	}

	return false
}

// SetBridge gets a reference to the given NullableInt32 and assigns it to the Bridge field.
func (o *PatchedWritableInterfaceTemplateRequest) SetBridge(v int32) {
	o.Bridge.Set(&v)
}
// SetBridgeNil sets the value for Bridge to be an explicit nil
func (o *PatchedWritableInterfaceTemplateRequest) SetBridgeNil() {
	o.Bridge.Set(nil)
}

// UnsetBridge ensures that no value is present for Bridge, not even an explicit nil
func (o *PatchedWritableInterfaceTemplateRequest) UnsetBridge() {
	o.Bridge.Unset()
}

// GetPoeMode returns the PoeMode field value if set, zero value otherwise.
func (o *PatchedWritableInterfaceTemplateRequest) GetPoeMode() InterfacePoeModeValue {
	if o == nil || IsNil(o.PoeMode) {
		var ret InterfacePoeModeValue
		return ret
	}
	return *o.PoeMode
}

// GetPoeModeOk returns a tuple with the PoeMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableInterfaceTemplateRequest) GetPoeModeOk() (*InterfacePoeModeValue, bool) {
	if o == nil || IsNil(o.PoeMode) {
		return nil, false
	}
	return o.PoeMode, true
}

// HasPoeMode returns a boolean if a field has been set.
func (o *PatchedWritableInterfaceTemplateRequest) HasPoeMode() bool {
	if o != nil && !IsNil(o.PoeMode) {
		return true
	}

	return false
}

// SetPoeMode gets a reference to the given InterfacePoeModeValue and assigns it to the PoeMode field.
func (o *PatchedWritableInterfaceTemplateRequest) SetPoeMode(v InterfacePoeModeValue) {
	o.PoeMode = &v
}

// GetPoeType returns the PoeType field value if set, zero value otherwise.
func (o *PatchedWritableInterfaceTemplateRequest) GetPoeType() InterfacePoeTypeValue {
	if o == nil || IsNil(o.PoeType) {
		var ret InterfacePoeTypeValue
		return ret
	}
	return *o.PoeType
}

// GetPoeTypeOk returns a tuple with the PoeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableInterfaceTemplateRequest) GetPoeTypeOk() (*InterfacePoeTypeValue, bool) {
	if o == nil || IsNil(o.PoeType) {
		return nil, false
	}
	return o.PoeType, true
}

// HasPoeType returns a boolean if a field has been set.
func (o *PatchedWritableInterfaceTemplateRequest) HasPoeType() bool {
	if o != nil && !IsNil(o.PoeType) {
		return true
	}

	return false
}

// SetPoeType gets a reference to the given InterfacePoeTypeValue and assigns it to the PoeType field.
func (o *PatchedWritableInterfaceTemplateRequest) SetPoeType(v InterfacePoeTypeValue) {
	o.PoeType = &v
}

// GetRfRole returns the RfRole field value if set, zero value otherwise.
func (o *PatchedWritableInterfaceTemplateRequest) GetRfRole() WirelessRole {
	if o == nil || IsNil(o.RfRole) {
		var ret WirelessRole
		return ret
	}
	return *o.RfRole
}

// GetRfRoleOk returns a tuple with the RfRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableInterfaceTemplateRequest) GetRfRoleOk() (*WirelessRole, bool) {
	if o == nil || IsNil(o.RfRole) {
		return nil, false
	}
	return o.RfRole, true
}

// HasRfRole returns a boolean if a field has been set.
func (o *PatchedWritableInterfaceTemplateRequest) HasRfRole() bool {
	if o != nil && !IsNil(o.RfRole) {
		return true
	}

	return false
}

// SetRfRole gets a reference to the given WirelessRole and assigns it to the RfRole field.
func (o *PatchedWritableInterfaceTemplateRequest) SetRfRole(v WirelessRole) {
	o.RfRole = &v
}

func (o PatchedWritableInterfaceTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableInterfaceTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DeviceType.IsSet() {
		toSerialize["device_type"] = o.DeviceType.Get()
	}
	if o.ModuleType.IsSet() {
		toSerialize["module_type"] = o.ModuleType.Get()
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
	if !IsNil(o.PoeMode) {
		toSerialize["poe_mode"] = o.PoeMode
	}
	if !IsNil(o.PoeType) {
		toSerialize["poe_type"] = o.PoeType
	}
	if !IsNil(o.RfRole) {
		toSerialize["rf_role"] = o.RfRole
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedWritableInterfaceTemplateRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedWritableInterfaceTemplateRequest := _PatchedWritableInterfaceTemplateRequest{}

	err = json.Unmarshal(data, &varPatchedWritableInterfaceTemplateRequest)

	if err != nil {
		return err
	}

	*o = PatchedWritableInterfaceTemplateRequest(varPatchedWritableInterfaceTemplateRequest)

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

type NullablePatchedWritableInterfaceTemplateRequest struct {
	value *PatchedWritableInterfaceTemplateRequest
	isSet bool
}

func (v NullablePatchedWritableInterfaceTemplateRequest) Get() *PatchedWritableInterfaceTemplateRequest {
	return v.value
}

func (v *NullablePatchedWritableInterfaceTemplateRequest) Set(val *PatchedWritableInterfaceTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableInterfaceTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableInterfaceTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableInterfaceTemplateRequest(val *PatchedWritableInterfaceTemplateRequest) *NullablePatchedWritableInterfaceTemplateRequest {
	return &NullablePatchedWritableInterfaceTemplateRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableInterfaceTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableInterfaceTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


