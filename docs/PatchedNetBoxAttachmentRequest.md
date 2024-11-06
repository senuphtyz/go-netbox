# PatchedNetBoxAttachmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | Pointer to **string** |  | [optional] 
**ObjectId** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**File** | Pointer to ***os.File** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedNetBoxAttachmentRequest

`func NewPatchedNetBoxAttachmentRequest() *PatchedNetBoxAttachmentRequest`

NewPatchedNetBoxAttachmentRequest instantiates a new PatchedNetBoxAttachmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedNetBoxAttachmentRequestWithDefaults

`func NewPatchedNetBoxAttachmentRequestWithDefaults() *PatchedNetBoxAttachmentRequest`

NewPatchedNetBoxAttachmentRequestWithDefaults instantiates a new PatchedNetBoxAttachmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *PatchedNetBoxAttachmentRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PatchedNetBoxAttachmentRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PatchedNetBoxAttachmentRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *PatchedNetBoxAttachmentRequest) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetObjectId

`func (o *PatchedNetBoxAttachmentRequest) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *PatchedNetBoxAttachmentRequest) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *PatchedNetBoxAttachmentRequest) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *PatchedNetBoxAttachmentRequest) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetName

`func (o *PatchedNetBoxAttachmentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedNetBoxAttachmentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedNetBoxAttachmentRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedNetBoxAttachmentRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedNetBoxAttachmentRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedNetBoxAttachmentRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedNetBoxAttachmentRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedNetBoxAttachmentRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFile

`func (o *PatchedNetBoxAttachmentRequest) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *PatchedNetBoxAttachmentRequest) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *PatchedNetBoxAttachmentRequest) SetFile(v *os.File)`

SetFile sets File field to given value.

### HasFile

`func (o *PatchedNetBoxAttachmentRequest) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetComments

`func (o *PatchedNetBoxAttachmentRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedNetBoxAttachmentRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedNetBoxAttachmentRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedNetBoxAttachmentRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


