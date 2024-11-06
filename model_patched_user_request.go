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
)

// checks if the PatchedUserRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedUserRequest{}

// PatchedUserRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedUserRequest struct {
	// Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.
	Username  *string `json:"username,omitempty" validate:"regexp=^[\\\\w.@+-]+$"`
	Password  *string `json:"password,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Email     *string `json:"email,omitempty"`
	// Designates whether the user can log into this admin site.
	IsStaff *bool `json:"is_staff,omitempty"`
	// Designates whether this user should be treated as active. Unselect this instead of deleting accounts.
	IsActive             *bool        `json:"is_active,omitempty"`
	DateJoined           *time.Time   `json:"date_joined,omitempty"`
	LastLogin            NullableTime `json:"last_login,omitempty"`
	Groups               []int32      `json:"groups,omitempty"`
	Permissions          []int32      `json:"permissions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedUserRequest PatchedUserRequest

// NewPatchedUserRequest instantiates a new PatchedUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedUserRequest() *PatchedUserRequest {
	this := PatchedUserRequest{}
	return &this
}

// NewPatchedUserRequestWithDefaults instantiates a new PatchedUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedUserRequestWithDefaults() *PatchedUserRequest {
	this := PatchedUserRequest{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *PatchedUserRequest) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserRequest) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *PatchedUserRequest) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *PatchedUserRequest) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *PatchedUserRequest) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserRequest) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *PatchedUserRequest) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *PatchedUserRequest) SetPassword(v string) {
	o.Password = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *PatchedUserRequest) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserRequest) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *PatchedUserRequest) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *PatchedUserRequest) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *PatchedUserRequest) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserRequest) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *PatchedUserRequest) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *PatchedUserRequest) SetLastName(v string) {
	o.LastName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *PatchedUserRequest) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserRequest) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PatchedUserRequest) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *PatchedUserRequest) SetEmail(v string) {
	o.Email = &v
}

// GetIsStaff returns the IsStaff field value if set, zero value otherwise.
func (o *PatchedUserRequest) GetIsStaff() bool {
	if o == nil || IsNil(o.IsStaff) {
		var ret bool
		return ret
	}
	return *o.IsStaff
}

// GetIsStaffOk returns a tuple with the IsStaff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserRequest) GetIsStaffOk() (*bool, bool) {
	if o == nil || IsNil(o.IsStaff) {
		return nil, false
	}
	return o.IsStaff, true
}

// HasIsStaff returns a boolean if a field has been set.
func (o *PatchedUserRequest) HasIsStaff() bool {
	if o != nil && !IsNil(o.IsStaff) {
		return true
	}

	return false
}

// SetIsStaff gets a reference to the given bool and assigns it to the IsStaff field.
func (o *PatchedUserRequest) SetIsStaff(v bool) {
	o.IsStaff = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *PatchedUserRequest) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserRequest) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *PatchedUserRequest) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *PatchedUserRequest) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetDateJoined returns the DateJoined field value if set, zero value otherwise.
func (o *PatchedUserRequest) GetDateJoined() time.Time {
	if o == nil || IsNil(o.DateJoined) {
		var ret time.Time
		return ret
	}
	return *o.DateJoined
}

// GetDateJoinedOk returns a tuple with the DateJoined field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserRequest) GetDateJoinedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateJoined) {
		return nil, false
	}
	return o.DateJoined, true
}

// HasDateJoined returns a boolean if a field has been set.
func (o *PatchedUserRequest) HasDateJoined() bool {
	if o != nil && !IsNil(o.DateJoined) {
		return true
	}

	return false
}

// SetDateJoined gets a reference to the given time.Time and assigns it to the DateJoined field.
func (o *PatchedUserRequest) SetDateJoined(v time.Time) {
	o.DateJoined = &v
}

// GetLastLogin returns the LastLogin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedUserRequest) GetLastLogin() time.Time {
	if o == nil || IsNil(o.LastLogin.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastLogin.Get()
}

// GetLastLoginOk returns a tuple with the LastLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedUserRequest) GetLastLoginOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastLogin.Get(), o.LastLogin.IsSet()
}

// HasLastLogin returns a boolean if a field has been set.
func (o *PatchedUserRequest) HasLastLogin() bool {
	if o != nil && o.LastLogin.IsSet() {
		return true
	}

	return false
}

// SetLastLogin gets a reference to the given NullableTime and assigns it to the LastLogin field.
func (o *PatchedUserRequest) SetLastLogin(v time.Time) {
	o.LastLogin.Set(&v)
}

// SetLastLoginNil sets the value for LastLogin to be an explicit nil
func (o *PatchedUserRequest) SetLastLoginNil() {
	o.LastLogin.Set(nil)
}

// UnsetLastLogin ensures that no value is present for LastLogin, not even an explicit nil
func (o *PatchedUserRequest) UnsetLastLogin() {
	o.LastLogin.Unset()
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *PatchedUserRequest) GetGroups() []int32 {
	if o == nil || IsNil(o.Groups) {
		var ret []int32
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserRequest) GetGroupsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *PatchedUserRequest) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []int32 and assigns it to the Groups field.
func (o *PatchedUserRequest) SetGroups(v []int32) {
	o.Groups = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *PatchedUserRequest) GetPermissions() []int32 {
	if o == nil || IsNil(o.Permissions) {
		var ret []int32
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserRequest) GetPermissionsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *PatchedUserRequest) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []int32 and assigns it to the Permissions field.
func (o *PatchedUserRequest) SetPermissions(v []int32) {
	o.Permissions = v
}

func (o PatchedUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedUserRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.FirstName) {
		toSerialize["first_name"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.IsStaff) {
		toSerialize["is_staff"] = o.IsStaff
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	if !IsNil(o.DateJoined) {
		toSerialize["date_joined"] = o.DateJoined
	}
	if o.LastLogin.IsSet() {
		toSerialize["last_login"] = o.LastLogin.Get()
	}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedUserRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedUserRequest := _PatchedUserRequest{}

	err = json.Unmarshal(data, &varPatchedUserRequest)

	if err != nil {
		return err
	}

	*o = PatchedUserRequest(varPatchedUserRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "username")
		delete(additionalProperties, "password")
		delete(additionalProperties, "first_name")
		delete(additionalProperties, "last_name")
		delete(additionalProperties, "email")
		delete(additionalProperties, "is_staff")
		delete(additionalProperties, "is_active")
		delete(additionalProperties, "date_joined")
		delete(additionalProperties, "last_login")
		delete(additionalProperties, "groups")
		delete(additionalProperties, "permissions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedUserRequest struct {
	value *PatchedUserRequest
	isSet bool
}

func (v NullablePatchedUserRequest) Get() *PatchedUserRequest {
	return v.value
}

func (v *NullablePatchedUserRequest) Set(val *PatchedUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedUserRequest(val *PatchedUserRequest) *NullablePatchedUserRequest {
	return &NullablePatchedUserRequest{value: val, isSet: true}
}

func (v NullablePatchedUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
