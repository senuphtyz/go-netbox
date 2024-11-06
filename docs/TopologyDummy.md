# TopologyDummy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTopologyDummy

`func NewTopologyDummy(id int32, ) *TopologyDummy`

NewTopologyDummy instantiates a new TopologyDummy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyDummyWithDefaults

`func NewTopologyDummyWithDefaults() *TopologyDummy`

NewTopologyDummyWithDefaults instantiates a new TopologyDummy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TopologyDummy) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TopologyDummy) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TopologyDummy) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *TopologyDummy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TopologyDummy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TopologyDummy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TopologyDummy) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TopologyDummy) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TopologyDummy) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


