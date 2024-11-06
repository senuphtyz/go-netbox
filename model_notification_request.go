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
	"fmt"
)

// checks if the NotificationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationRequest{}

// NotificationRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type NotificationRequest struct {
	ObjectType string `json:"object_type"`
	ObjectId int64 `json:"object_id"`
	User BriefUserRequest `json:"user"`
	Read NullableTime `json:"read,omitempty"`
	EventType Event `json:"event_type"`
	AdditionalProperties map[string]interface{}
}

type _NotificationRequest NotificationRequest

// NewNotificationRequest instantiates a new NotificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationRequest(objectType string, objectId int64, user BriefUserRequest, eventType Event) *NotificationRequest {
	this := NotificationRequest{}
	this.ObjectType = objectType
	this.ObjectId = objectId
	this.User = user
	this.EventType = eventType
	return &this
}

// NewNotificationRequestWithDefaults instantiates a new NotificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationRequestWithDefaults() *NotificationRequest {
	this := NotificationRequest{}
	return &this
}

// GetObjectType returns the ObjectType field value
func (o *NotificationRequest) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NotificationRequest) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NotificationRequest) SetObjectType(v string) {
	o.ObjectType = v
}

// GetObjectId returns the ObjectId field value
func (o *NotificationRequest) GetObjectId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value
// and a boolean to check if the value has been set.
func (o *NotificationRequest) GetObjectIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectId, true
}

// SetObjectId sets field value
func (o *NotificationRequest) SetObjectId(v int64) {
	o.ObjectId = v
}

// GetUser returns the User field value
func (o *NotificationRequest) GetUser() BriefUserRequest {
	if o == nil {
		var ret BriefUserRequest
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *NotificationRequest) GetUserOk() (*BriefUserRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *NotificationRequest) SetUser(v BriefUserRequest) {
	o.User = v
}

// GetRead returns the Read field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationRequest) GetRead() time.Time {
	if o == nil || IsNil(o.Read.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Read.Get()
}

// GetReadOk returns a tuple with the Read field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationRequest) GetReadOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Read.Get(), o.Read.IsSet()
}

// HasRead returns a boolean if a field has been set.
func (o *NotificationRequest) HasRead() bool {
	if o != nil && o.Read.IsSet() {
		return true
	}

	return false
}

// SetRead gets a reference to the given NullableTime and assigns it to the Read field.
func (o *NotificationRequest) SetRead(v time.Time) {
	o.Read.Set(&v)
}
// SetReadNil sets the value for Read to be an explicit nil
func (o *NotificationRequest) SetReadNil() {
	o.Read.Set(nil)
}

// UnsetRead ensures that no value is present for Read, not even an explicit nil
func (o *NotificationRequest) UnsetRead() {
	o.Read.Unset()
}

// GetEventType returns the EventType field value
func (o *NotificationRequest) GetEventType() Event {
	if o == nil {
		var ret Event
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *NotificationRequest) GetEventTypeOk() (*Event, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *NotificationRequest) SetEventType(v Event) {
	o.EventType = v
}

func (o NotificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object_type"] = o.ObjectType
	toSerialize["object_id"] = o.ObjectId
	toSerialize["user"] = o.User
	if o.Read.IsSet() {
		toSerialize["read"] = o.Read.Get()
	}
	toSerialize["event_type"] = o.EventType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NotificationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object_type",
		"object_id",
		"user",
		"event_type",
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

	varNotificationRequest := _NotificationRequest{}

	err = json.Unmarshal(data, &varNotificationRequest)

	if err != nil {
		return err
	}

	*o = NotificationRequest(varNotificationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "object_id")
		delete(additionalProperties, "user")
		delete(additionalProperties, "read")
		delete(additionalProperties, "event_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNotificationRequest struct {
	value *NotificationRequest
	isSet bool
}

func (v NullableNotificationRequest) Get() *NotificationRequest {
	return v.value
}

func (v *NullableNotificationRequest) Set(val *NotificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationRequest(val *NotificationRequest) *NullableNotificationRequest {
	return &NullableNotificationRequest{value: val, isSet: true}
}

func (v NullableNotificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

