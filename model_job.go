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

// checks if the Job type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Job{}

// Job struct for Job
type Job struct {
	Id         int32          `json:"id"`
	Url        string         `json:"url"`
	DisplayUrl string         `json:"display_url"`
	Display    string         `json:"display"`
	ObjectType string         `json:"object_type"`
	ObjectId   NullableInt64  `json:"object_id,omitempty"`
	Name       string         `json:"name"`
	Status     BriefJobStatus `json:"status"`
	Created    time.Time      `json:"created"`
	Scheduled  NullableTime   `json:"scheduled,omitempty"`
	// Recurrence interval (in minutes)
	Interval             NullableInt32 `json:"interval,omitempty"`
	Started              NullableTime  `json:"started,omitempty"`
	Completed            NullableTime  `json:"completed,omitempty"`
	User                 BriefUser     `json:"user"`
	Data                 interface{}   `json:"data,omitempty"`
	Error                string        `json:"error"`
	JobId                string        `json:"job_id"`
	AdditionalProperties map[string]interface{}
}

type _Job Job

// NewJob instantiates a new Job object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJob(id int32, url string, displayUrl string, display string, objectType string, name string, status BriefJobStatus, created time.Time, user BriefUser, error_ string, jobId string) *Job {
	this := Job{}
	this.Id = id
	this.Url = url
	this.DisplayUrl = displayUrl
	this.Display = display
	this.ObjectType = objectType
	this.Name = name
	this.Status = status
	this.Created = created
	this.User = user
	this.Error = error_
	this.JobId = jobId
	return &this
}

// NewJobWithDefaults instantiates a new Job object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobWithDefaults() *Job {
	this := Job{}
	return &this
}

// GetId returns the Id field value
func (o *Job) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Job) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Job) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *Job) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Job) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Job) SetUrl(v string) {
	o.Url = v
}

// GetDisplayUrl returns the DisplayUrl field value
func (o *Job) GetDisplayUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayUrl
}

// GetDisplayUrlOk returns a tuple with the DisplayUrl field value
// and a boolean to check if the value has been set.
func (o *Job) GetDisplayUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayUrl, true
}

// SetDisplayUrl sets field value
func (o *Job) SetDisplayUrl(v string) {
	o.DisplayUrl = v
}

// GetDisplay returns the Display field value
func (o *Job) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *Job) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *Job) SetDisplay(v string) {
	o.Display = v
}

// GetObjectType returns the ObjectType field value
func (o *Job) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *Job) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *Job) SetObjectType(v string) {
	o.ObjectType = v
}

// GetObjectId returns the ObjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Job) GetObjectId() int64 {
	if o == nil || IsNil(o.ObjectId.Get()) {
		var ret int64
		return ret
	}
	return *o.ObjectId.Get()
}

// GetObjectIdOk returns a tuple with the ObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Job) GetObjectIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObjectId.Get(), o.ObjectId.IsSet()
}

// HasObjectId returns a boolean if a field has been set.
func (o *Job) HasObjectId() bool {
	if o != nil && o.ObjectId.IsSet() {
		return true
	}

	return false
}

// SetObjectId gets a reference to the given NullableInt64 and assigns it to the ObjectId field.
func (o *Job) SetObjectId(v int64) {
	o.ObjectId.Set(&v)
}

// SetObjectIdNil sets the value for ObjectId to be an explicit nil
func (o *Job) SetObjectIdNil() {
	o.ObjectId.Set(nil)
}

// UnsetObjectId ensures that no value is present for ObjectId, not even an explicit nil
func (o *Job) UnsetObjectId() {
	o.ObjectId.Unset()
}

// GetName returns the Name field value
func (o *Job) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Job) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Job) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value
func (o *Job) GetStatus() BriefJobStatus {
	if o == nil {
		var ret BriefJobStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Job) GetStatusOk() (*BriefJobStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Job) SetStatus(v BriefJobStatus) {
	o.Status = v
}

// GetCreated returns the Created field value
func (o *Job) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *Job) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *Job) SetCreated(v time.Time) {
	o.Created = v
}

// GetScheduled returns the Scheduled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Job) GetScheduled() time.Time {
	if o == nil || IsNil(o.Scheduled.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Scheduled.Get()
}

// GetScheduledOk returns a tuple with the Scheduled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Job) GetScheduledOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scheduled.Get(), o.Scheduled.IsSet()
}

// HasScheduled returns a boolean if a field has been set.
func (o *Job) HasScheduled() bool {
	if o != nil && o.Scheduled.IsSet() {
		return true
	}

	return false
}

// SetScheduled gets a reference to the given NullableTime and assigns it to the Scheduled field.
func (o *Job) SetScheduled(v time.Time) {
	o.Scheduled.Set(&v)
}

// SetScheduledNil sets the value for Scheduled to be an explicit nil
func (o *Job) SetScheduledNil() {
	o.Scheduled.Set(nil)
}

// UnsetScheduled ensures that no value is present for Scheduled, not even an explicit nil
func (o *Job) UnsetScheduled() {
	o.Scheduled.Unset()
}

// GetInterval returns the Interval field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Job) GetInterval() int32 {
	if o == nil || IsNil(o.Interval.Get()) {
		var ret int32
		return ret
	}
	return *o.Interval.Get()
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Job) GetIntervalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Interval.Get(), o.Interval.IsSet()
}

// HasInterval returns a boolean if a field has been set.
func (o *Job) HasInterval() bool {
	if o != nil && o.Interval.IsSet() {
		return true
	}

	return false
}

// SetInterval gets a reference to the given NullableInt32 and assigns it to the Interval field.
func (o *Job) SetInterval(v int32) {
	o.Interval.Set(&v)
}

// SetIntervalNil sets the value for Interval to be an explicit nil
func (o *Job) SetIntervalNil() {
	o.Interval.Set(nil)
}

// UnsetInterval ensures that no value is present for Interval, not even an explicit nil
func (o *Job) UnsetInterval() {
	o.Interval.Unset()
}

// GetStarted returns the Started field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Job) GetStarted() time.Time {
	if o == nil || IsNil(o.Started.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Started.Get()
}

// GetStartedOk returns a tuple with the Started field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Job) GetStartedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Started.Get(), o.Started.IsSet()
}

// HasStarted returns a boolean if a field has been set.
func (o *Job) HasStarted() bool {
	if o != nil && o.Started.IsSet() {
		return true
	}

	return false
}

// SetStarted gets a reference to the given NullableTime and assigns it to the Started field.
func (o *Job) SetStarted(v time.Time) {
	o.Started.Set(&v)
}

// SetStartedNil sets the value for Started to be an explicit nil
func (o *Job) SetStartedNil() {
	o.Started.Set(nil)
}

// UnsetStarted ensures that no value is present for Started, not even an explicit nil
func (o *Job) UnsetStarted() {
	o.Started.Unset()
}

// GetCompleted returns the Completed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Job) GetCompleted() time.Time {
	if o == nil || IsNil(o.Completed.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Completed.Get()
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Job) GetCompletedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Completed.Get(), o.Completed.IsSet()
}

// HasCompleted returns a boolean if a field has been set.
func (o *Job) HasCompleted() bool {
	if o != nil && o.Completed.IsSet() {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given NullableTime and assigns it to the Completed field.
func (o *Job) SetCompleted(v time.Time) {
	o.Completed.Set(&v)
}

// SetCompletedNil sets the value for Completed to be an explicit nil
func (o *Job) SetCompletedNil() {
	o.Completed.Set(nil)
}

// UnsetCompleted ensures that no value is present for Completed, not even an explicit nil
func (o *Job) UnsetCompleted() {
	o.Completed.Unset()
}

// GetUser returns the User field value
func (o *Job) GetUser() BriefUser {
	if o == nil {
		var ret BriefUser
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *Job) GetUserOk() (*BriefUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *Job) SetUser(v BriefUser) {
	o.User = v
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Job) GetData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Job) GetDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Job) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given interface{} and assigns it to the Data field.
func (o *Job) SetData(v interface{}) {
	o.Data = v
}

// GetError returns the Error field value
func (o *Job) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *Job) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *Job) SetError(v string) {
	o.Error = v
}

// GetJobId returns the JobId field value
func (o *Job) GetJobId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value
// and a boolean to check if the value has been set.
func (o *Job) GetJobIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobId, true
}

// SetJobId sets field value
func (o *Job) SetJobId(v string) {
	o.JobId = v
}

func (o Job) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Job) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display_url"] = o.DisplayUrl
	toSerialize["display"] = o.Display
	toSerialize["object_type"] = o.ObjectType
	if o.ObjectId.IsSet() {
		toSerialize["object_id"] = o.ObjectId.Get()
	}
	toSerialize["name"] = o.Name
	toSerialize["status"] = o.Status
	toSerialize["created"] = o.Created
	if o.Scheduled.IsSet() {
		toSerialize["scheduled"] = o.Scheduled.Get()
	}
	if o.Interval.IsSet() {
		toSerialize["interval"] = o.Interval.Get()
	}
	if o.Started.IsSet() {
		toSerialize["started"] = o.Started.Get()
	}
	if o.Completed.IsSet() {
		toSerialize["completed"] = o.Completed.Get()
	}
	toSerialize["user"] = o.User
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	toSerialize["error"] = o.Error
	toSerialize["job_id"] = o.JobId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Job) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display_url",
		"display",
		"object_type",
		"name",
		"status",
		"created",
		"user",
		"error",
		"job_id",
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

	varJob := _Job{}

	err = json.Unmarshal(data, &varJob)

	if err != nil {
		return err
	}

	*o = Job(varJob)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display_url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "object_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "status")
		delete(additionalProperties, "created")
		delete(additionalProperties, "scheduled")
		delete(additionalProperties, "interval")
		delete(additionalProperties, "started")
		delete(additionalProperties, "completed")
		delete(additionalProperties, "user")
		delete(additionalProperties, "data")
		delete(additionalProperties, "error")
		delete(additionalProperties, "job_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableJob struct {
	value *Job
	isSet bool
}

func (v NullableJob) Get() *Job {
	return v.value
}

func (v *NullableJob) Set(val *Job) {
	v.value = val
	v.isSet = true
}

func (v NullableJob) IsSet() bool {
	return v.isSet
}

func (v *NullableJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJob(val *Job) *NullableJob {
	return &NullableJob{value: val, isSet: true}
}

func (v NullableJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
