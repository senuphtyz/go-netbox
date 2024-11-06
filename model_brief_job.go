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
	"time"
)

// checks if the BriefJob type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefJob{}

// BriefJob struct for BriefJob
type BriefJob struct {
	Url                  string         `json:"url"`
	Status               BriefJobStatus `json:"status"`
	Created              time.Time      `json:"created"`
	Completed            NullableTime   `json:"completed,omitempty"`
	User                 BriefUser      `json:"user"`
	AdditionalProperties map[string]interface{}
}

type _BriefJob BriefJob

// NewBriefJob instantiates a new BriefJob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefJob(url string, status BriefJobStatus, created time.Time, user BriefUser) *BriefJob {
	this := BriefJob{}
	this.Url = url
	this.Status = status
	this.Created = created
	this.User = user
	return &this
}

// NewBriefJobWithDefaults instantiates a new BriefJob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefJobWithDefaults() *BriefJob {
	this := BriefJob{}
	return &this
}

// GetUrl returns the Url field value
func (o *BriefJob) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *BriefJob) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *BriefJob) SetUrl(v string) {
	o.Url = v
}

// GetStatus returns the Status field value
func (o *BriefJob) GetStatus() BriefJobStatus {
	if o == nil {
		var ret BriefJobStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BriefJob) GetStatusOk() (*BriefJobStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BriefJob) SetStatus(v BriefJobStatus) {
	o.Status = v
}

// GetCreated returns the Created field value
func (o *BriefJob) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *BriefJob) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *BriefJob) SetCreated(v time.Time) {
	o.Created = v
}

// GetCompleted returns the Completed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BriefJob) GetCompleted() time.Time {
	if o == nil || IsNil(o.Completed.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Completed.Get()
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BriefJob) GetCompletedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Completed.Get(), o.Completed.IsSet()
}

// HasCompleted returns a boolean if a field has been set.
func (o *BriefJob) HasCompleted() bool {
	if o != nil && o.Completed.IsSet() {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given NullableTime and assigns it to the Completed field.
func (o *BriefJob) SetCompleted(v time.Time) {
	o.Completed.Set(&v)
}

// SetCompletedNil sets the value for Completed to be an explicit nil
func (o *BriefJob) SetCompletedNil() {
	o.Completed.Set(nil)
}

// UnsetCompleted ensures that no value is present for Completed, not even an explicit nil
func (o *BriefJob) UnsetCompleted() {
	o.Completed.Unset()
}

// GetUser returns the User field value
func (o *BriefJob) GetUser() BriefUser {
	if o == nil {
		var ret BriefUser
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *BriefJob) GetUserOk() (*BriefUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *BriefJob) SetUser(v BriefUser) {
	o.User = v
}

func (o BriefJob) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefJob) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	toSerialize["status"] = o.Status
	toSerialize["created"] = o.Created
	if o.Completed.IsSet() {
		toSerialize["completed"] = o.Completed.Get()
	}
	toSerialize["user"] = o.User

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefJob) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
		"status",
		"created",
		"user",
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

	varBriefJob := _BriefJob{}

	err = json.Unmarshal(data, &varBriefJob)

	if err != nil {
		return err
	}

	*o = BriefJob(varBriefJob)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "url")
		delete(additionalProperties, "status")
		delete(additionalProperties, "created")
		delete(additionalProperties, "completed")
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefJob struct {
	value *BriefJob
	isSet bool
}

func (v NullableBriefJob) Get() *BriefJob {
	return v.value
}

func (v *NullableBriefJob) Set(val *BriefJob) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefJob) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefJob(val *BriefJob) *NullableBriefJob {
	return &NullableBriefJob{value: val, isSet: true}
}

func (v NullableBriefJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
