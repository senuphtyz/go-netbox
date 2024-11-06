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

// checks if the BriefIPAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefIPAddress{}

// BriefIPAddress Adds support for custom fields and tags.
type BriefIPAddress struct {
	Id int32 `json:"id"`
	Url string `json:"url"`
	Display string `json:"display"`
	Family AggregateFamily `json:"family"`
	Address string `json:"address"`
	Description *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BriefIPAddress BriefIPAddress

// NewBriefIPAddress instantiates a new BriefIPAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefIPAddress(id int32, url string, display string, family AggregateFamily, address string) *BriefIPAddress {
	this := BriefIPAddress{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Family = family
	this.Address = address
	return &this
}

// NewBriefIPAddressWithDefaults instantiates a new BriefIPAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefIPAddressWithDefaults() *BriefIPAddress {
	this := BriefIPAddress{}
	return &this
}

// GetId returns the Id field value
func (o *BriefIPAddress) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BriefIPAddress) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BriefIPAddress) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *BriefIPAddress) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *BriefIPAddress) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *BriefIPAddress) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *BriefIPAddress) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *BriefIPAddress) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *BriefIPAddress) SetDisplay(v string) {
	o.Display = v
}

// GetFamily returns the Family field value
func (o *BriefIPAddress) GetFamily() AggregateFamily {
	if o == nil {
		var ret AggregateFamily
		return ret
	}

	return o.Family
}

// GetFamilyOk returns a tuple with the Family field value
// and a boolean to check if the value has been set.
func (o *BriefIPAddress) GetFamilyOk() (*AggregateFamily, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Family, true
}

// SetFamily sets field value
func (o *BriefIPAddress) SetFamily(v AggregateFamily) {
	o.Family = v
}

// GetAddress returns the Address field value
func (o *BriefIPAddress) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *BriefIPAddress) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *BriefIPAddress) SetAddress(v string) {
	o.Address = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BriefIPAddress) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefIPAddress) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BriefIPAddress) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BriefIPAddress) SetDescription(v string) {
	o.Description = &v
}

func (o BriefIPAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefIPAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["family"] = o.Family
	toSerialize["address"] = o.Address
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefIPAddress) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"family",
		"address",
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

	varBriefIPAddress := _BriefIPAddress{}

	err = json.Unmarshal(data, &varBriefIPAddress)

	if err != nil {
		return err
	}

	*o = BriefIPAddress(varBriefIPAddress)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "family")
		delete(additionalProperties, "address")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefIPAddress struct {
	value *BriefIPAddress
	isSet bool
}

func (v NullableBriefIPAddress) Get() *BriefIPAddress {
	return v.value
}

func (v *NullableBriefIPAddress) Set(val *BriefIPAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefIPAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefIPAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefIPAddress(val *BriefIPAddress) *NullableBriefIPAddress {
	return &NullableBriefIPAddress{value: val, isSet: true}
}

func (v NullableBriefIPAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefIPAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


