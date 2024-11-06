# NetBoxAttachmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | **string** |  | 
**ObjectId** | **int64** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**File** | ***os.File** |  | 
**Comments** | Pointer to **string** |  | [optional] 

## Methods

### NewNetBoxAttachmentRequest

`func NewNetBoxAttachmentRequest(objectType string, objectId int64, file *os.File, ) *NetBoxAttachmentRequest`

NewNetBoxAttachmentRequest instantiates a new NetBoxAttachmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetBoxAttachmentRequestWithDefaults

`func NewNetBoxAttachmentRequestWithDefaults() *NetBoxAttachmentRequest`

NewNetBoxAttachmentRequestWithDefaults instantiates a new NetBoxAttachmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *NetBoxAttachmentRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NetBoxAttachmentRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NetBoxAttachmentRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetObjectId

`func (o *NetBoxAttachmentRequest) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *NetBoxAttachmentRequest) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *NetBoxAttachmentRequest) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.


### GetName

`func (o *NetBoxAttachmentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetBoxAttachmentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetBoxAttachmentRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetBoxAttachmentRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *NetBoxAttachmentRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetBoxAttachmentRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetBoxAttachmentRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetBoxAttachmentRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFile

`func (o *NetBoxAttachmentRequest) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *NetBoxAttachmentRequest) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *NetBoxAttachmentRequest) SetFile(v *os.File)`

SetFile sets File field to given value.


### GetComments

`func (o *NetBoxAttachmentRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *NetBoxAttachmentRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *NetBoxAttachmentRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *NetBoxAttachmentRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


