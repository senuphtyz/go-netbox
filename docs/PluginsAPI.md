# \PluginsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PluginsNetboxAttachmentsNetboxAttachmentsBulkDestroy**](PluginsAPI.md#PluginsNetboxAttachmentsNetboxAttachmentsBulkDestroy) | **Delete** /api/plugins/netbox-attachments/netbox-attachments/ | 
[**PluginsNetboxAttachmentsNetboxAttachmentsBulkPartialUpdate**](PluginsAPI.md#PluginsNetboxAttachmentsNetboxAttachmentsBulkPartialUpdate) | **Patch** /api/plugins/netbox-attachments/netbox-attachments/ | 
[**PluginsNetboxAttachmentsNetboxAttachmentsBulkUpdate**](PluginsAPI.md#PluginsNetboxAttachmentsNetboxAttachmentsBulkUpdate) | **Put** /api/plugins/netbox-attachments/netbox-attachments/ | 
[**PluginsNetboxAttachmentsNetboxAttachmentsCreate**](PluginsAPI.md#PluginsNetboxAttachmentsNetboxAttachmentsCreate) | **Post** /api/plugins/netbox-attachments/netbox-attachments/ | 
[**PluginsNetboxAttachmentsNetboxAttachmentsDestroy**](PluginsAPI.md#PluginsNetboxAttachmentsNetboxAttachmentsDestroy) | **Delete** /api/plugins/netbox-attachments/netbox-attachments/{id}/ | 
[**PluginsNetboxAttachmentsNetboxAttachmentsList**](PluginsAPI.md#PluginsNetboxAttachmentsNetboxAttachmentsList) | **Get** /api/plugins/netbox-attachments/netbox-attachments/ | 
[**PluginsNetboxAttachmentsNetboxAttachmentsPartialUpdate**](PluginsAPI.md#PluginsNetboxAttachmentsNetboxAttachmentsPartialUpdate) | **Patch** /api/plugins/netbox-attachments/netbox-attachments/{id}/ | 
[**PluginsNetboxAttachmentsNetboxAttachmentsRetrieve**](PluginsAPI.md#PluginsNetboxAttachmentsNetboxAttachmentsRetrieve) | **Get** /api/plugins/netbox-attachments/netbox-attachments/{id}/ | 
[**PluginsNetboxAttachmentsNetboxAttachmentsUpdate**](PluginsAPI.md#PluginsNetboxAttachmentsNetboxAttachmentsUpdate) | **Put** /api/plugins/netbox-attachments/netbox-attachments/{id}/ | 
[**PluginsNetboxTopologyViewsImagesList**](PluginsAPI.md#PluginsNetboxTopologyViewsImagesList) | **Get** /api/plugins/netbox_topology_views/images/ | 
[**PluginsNetboxTopologyViewsImagesRetrieve**](PluginsAPI.md#PluginsNetboxTopologyViewsImagesRetrieve) | **Get** /api/plugins/netbox_topology_views/images/{id}/ | 
[**PluginsNetboxTopologyViewsImagesSaveCreate**](PluginsAPI.md#PluginsNetboxTopologyViewsImagesSaveCreate) | **Post** /api/plugins/netbox_topology_views/images/save/ | 
[**PluginsNetboxTopologyViewsSaveCoordsList**](PluginsAPI.md#PluginsNetboxTopologyViewsSaveCoordsList) | **Get** /api/plugins/netbox_topology_views/save-coords/ | 
[**PluginsNetboxTopologyViewsSaveCoordsRetrieve**](PluginsAPI.md#PluginsNetboxTopologyViewsSaveCoordsRetrieve) | **Get** /api/plugins/netbox_topology_views/save-coords/{id}/ | 
[**PluginsNetboxTopologyViewsSaveCoordsSaveCoordsPartialUpdate**](PluginsAPI.md#PluginsNetboxTopologyViewsSaveCoordsSaveCoordsPartialUpdate) | **Patch** /api/plugins/netbox_topology_views/save-coords/save_coords/ | 
[**PluginsNetboxTopologyViewsXmlExportList**](PluginsAPI.md#PluginsNetboxTopologyViewsXmlExportList) | **Get** /api/plugins/netbox_topology_views/xml-export/ | 



## PluginsNetboxAttachmentsNetboxAttachmentsBulkDestroy

> PluginsNetboxAttachmentsNetboxAttachmentsBulkDestroy(ctx).NetBoxAttachmentRequest(netBoxAttachmentRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	netBoxAttachmentRequest := []openapiclient.NetBoxAttachmentRequest{*openapiclient.NewNetBoxAttachmentRequest("ObjectType_example", int64(123), "TODO")} // []NetBoxAttachmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PluginsAPI.PluginsNetboxAttachmentsNetboxAttachmentsBulkDestroy(context.Background()).NetBoxAttachmentRequest(netBoxAttachmentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.PluginsNetboxAttachmentsNetboxAttachmentsBulkDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPluginsNetboxAttachmentsNetboxAttachmentsBulkDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **netBoxAttachmentRequest** | [**[]NetBoxAttachmentRequest**](NetBoxAttachmentRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PluginsNetboxAttachmentsNetboxAttachmentsBulkPartialUpdate

> []NetBoxAttachment PluginsNetboxAttachmentsNetboxAttachmentsBulkPartialUpdate(ctx).NetBoxAttachmentRequest(netBoxAttachmentRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	netBoxAttachmentRequest := []openapiclient.NetBoxAttachmentRequest{*openapiclient.NewNetBoxAttachmentRequest("ObjectType_example", int64(123), "TODO")} // []NetBoxAttachmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginsAPI.PluginsNetboxAttachmentsNetboxAttachmentsBulkPartialUpdate(context.Background()).NetBoxAttachmentRequest(netBoxAttachmentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.PluginsNetboxAttachmentsNetboxAttachmentsBulkPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PluginsNetboxAttachmentsNetboxAttachmentsBulkPartialUpdate`: []NetBoxAttachment
	fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.PluginsNetboxAttachmentsNetboxAttachmentsBulkPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPluginsNetboxAttachmentsNetboxAttachmentsBulkPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **netBoxAttachmentRequest** | [**[]NetBoxAttachmentRequest**](NetBoxAttachmentRequest.md) |  | 

### Return type

[**[]NetBoxAttachment**](NetBoxAttachment.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PluginsNetboxAttachmentsNetboxAttachmentsBulkUpdate

> []NetBoxAttachment PluginsNetboxAttachmentsNetboxAttachmentsBulkUpdate(ctx).NetBoxAttachmentRequest(netBoxAttachmentRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	netBoxAttachmentRequest := []openapiclient.NetBoxAttachmentRequest{*openapiclient.NewNetBoxAttachmentRequest("ObjectType_example", int64(123), "TODO")} // []NetBoxAttachmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginsAPI.PluginsNetboxAttachmentsNetboxAttachmentsBulkUpdate(context.Background()).NetBoxAttachmentRequest(netBoxAttachmentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.PluginsNetboxAttachmentsNetboxAttachmentsBulkUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PluginsNetboxAttachmentsNetboxAttachmentsBulkUpdate`: []NetBoxAttachment
	fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.PluginsNetboxAttachmentsNetboxAttachmentsBulkUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPluginsNetboxAttachmentsNetboxAttachmentsBulkUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **netBoxAttachmentRequest** | [**[]NetBoxAttachmentRequest**](NetBoxAttachmentRequest.md) |  | 

### Return type

[**[]NetBoxAttachment**](NetBoxAttachment.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PluginsNetboxAttachmentsNetboxAttachmentsCreate

> NetBoxAttachment PluginsNetboxAttachmentsNetboxAttachmentsCreate(ctx).NetBoxAttachmentRequest(netBoxAttachmentRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	netBoxAttachmentRequest := *openapiclient.NewNetBoxAttachmentRequest("ObjectType_example", int64(123), "TODO") // NetBoxAttachmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginsAPI.PluginsNetboxAttachmentsNetboxAttachmentsCreate(context.Background()).NetBoxAttachmentRequest(netBoxAttachmentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.PluginsNetboxAttachmentsNetboxAttachmentsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PluginsNetboxAttachmentsNetboxAttachmentsCreate`: NetBoxAttachment
	fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.PluginsNetboxAttachmentsNetboxAttachmentsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPluginsNetboxAttachmentsNetboxAttachmentsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **netBoxAttachmentRequest** | [**NetBoxAttachmentRequest**](NetBoxAttachmentRequest.md) |  | 

### Return type

[**NetBoxAttachment**](NetBoxAttachment.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PluginsNetboxAttachmentsNetboxAttachmentsDestroy

> PluginsNetboxAttachmentsNetboxAttachmentsDestroy(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this NetBox Attachment.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PluginsAPI.PluginsNetboxAttachmentsNetboxAttachmentsDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.PluginsNetboxAttachmentsNetboxAttachmentsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this NetBox Attachment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPluginsNetboxAttachmentsNetboxAttachmentsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PluginsNetboxAttachmentsNetboxAttachmentsList

> PaginatedNetBoxAttachmentList PluginsNetboxAttachmentsNetboxAttachmentsList(ctx).Created(created).CreatedByRequest(createdByRequest).Description(description).Id(id).IdEmpty(idEmpty).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdN(idN).LastUpdated(lastUpdated).LastUpdatedEmpty(lastUpdatedEmpty).LastUpdatedGt(lastUpdatedGt).LastUpdatedGte(lastUpdatedGte).LastUpdatedLt(lastUpdatedLt).LastUpdatedLte(lastUpdatedLte).LastUpdatedN(lastUpdatedN).Limit(limit).ModifiedByRequest(modifiedByRequest).Name(name).ObjectId(objectId).ObjectIdEmpty(objectIdEmpty).ObjectIdGt(objectIdGt).ObjectIdGte(objectIdGte).ObjectIdLt(objectIdLt).ObjectIdLte(objectIdLte).ObjectIdN(objectIdN).ObjectType(objectType).ObjectTypeN(objectTypeN).ObjectTypeId(objectTypeId).ObjectTypeIdN(objectTypeIdN).Offset(offset).Ordering(ordering).Q(q).Tag(tag).TagN(tagN).UpdatedByRequest(updatedByRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	created := time.Now() // time.Time |  (optional)
	createdByRequest := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	description := "description_example" // string |  (optional)
	id := []int32{int32(123)} // []int32 |  (optional)
	idEmpty := true // bool |  (optional)
	idGt := []int32{int32(123)} // []int32 |  (optional)
	idGte := []int32{int32(123)} // []int32 |  (optional)
	idLt := []int32{int32(123)} // []int32 |  (optional)
	idLte := []int32{int32(123)} // []int32 |  (optional)
	idN := []int32{int32(123)} // []int32 |  (optional)
	lastUpdated := []time.Time{time.Now()} // []time.Time |  (optional)
	lastUpdatedEmpty := []time.Time{time.Now()} // []time.Time |  (optional)
	lastUpdatedGt := []time.Time{time.Now()} // []time.Time |  (optional)
	lastUpdatedGte := []time.Time{time.Now()} // []time.Time |  (optional)
	lastUpdatedLt := []time.Time{time.Now()} // []time.Time |  (optional)
	lastUpdatedLte := []time.Time{time.Now()} // []time.Time |  (optional)
	lastUpdatedN := []time.Time{time.Now()} // []time.Time |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	modifiedByRequest := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	name := "name_example" // string |  (optional)
	objectId := []int32{int32(123)} // []int32 |  (optional)
	objectIdEmpty := true // bool |  (optional)
	objectIdGt := []int32{int32(123)} // []int32 |  (optional)
	objectIdGte := []int32{int32(123)} // []int32 |  (optional)
	objectIdLt := []int32{int32(123)} // []int32 |  (optional)
	objectIdLte := []int32{int32(123)} // []int32 |  (optional)
	objectIdN := []int32{int32(123)} // []int32 |  (optional)
	objectType := "objectType_example" // string |  (optional)
	objectTypeN := "objectTypeN_example" // string |  (optional)
	objectTypeId := int32(56) // int32 |  (optional)
	objectTypeIdN := int32(56) // int32 |  (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	q := "q_example" // string | Search (optional)
	tag := []string{"Inner_example"} // []string |  (optional)
	tagN := []string{"Inner_example"} // []string |  (optional)
	updatedByRequest := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginsAPI.PluginsNetboxAttachmentsNetboxAttachmentsList(context.Background()).Created(created).CreatedByRequest(createdByRequest).Description(description).Id(id).IdEmpty(idEmpty).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdN(idN).LastUpdated(lastUpdated).LastUpdatedEmpty(lastUpdatedEmpty).LastUpdatedGt(lastUpdatedGt).LastUpdatedGte(lastUpdatedGte).LastUpdatedLt(lastUpdatedLt).LastUpdatedLte(lastUpdatedLte).LastUpdatedN(lastUpdatedN).Limit(limit).ModifiedByRequest(modifiedByRequest).Name(name).ObjectId(objectId).ObjectIdEmpty(objectIdEmpty).ObjectIdGt(objectIdGt).ObjectIdGte(objectIdGte).ObjectIdLt(objectIdLt).ObjectIdLte(objectIdLte).ObjectIdN(objectIdN).ObjectType(objectType).ObjectTypeN(objectTypeN).ObjectTypeId(objectTypeId).ObjectTypeIdN(objectTypeIdN).Offset(offset).Ordering(ordering).Q(q).Tag(tag).TagN(tagN).UpdatedByRequest(updatedByRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.PluginsNetboxAttachmentsNetboxAttachmentsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PluginsNetboxAttachmentsNetboxAttachmentsList`: PaginatedNetBoxAttachmentList
	fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.PluginsNetboxAttachmentsNetboxAttachmentsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPluginsNetboxAttachmentsNetboxAttachmentsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **created** | **time.Time** |  | 
 **createdByRequest** | **string** |  | 
 **description** | **string** |  | 
 **id** | **[]int32** |  | 
 **idEmpty** | **bool** |  | 
 **idGt** | **[]int32** |  | 
 **idGte** | **[]int32** |  | 
 **idLt** | **[]int32** |  | 
 **idLte** | **[]int32** |  | 
 **idN** | **[]int32** |  | 
 **lastUpdated** | [**[]time.Time**](time.Time.md) |  | 
 **lastUpdatedEmpty** | [**[]time.Time**](time.Time.md) |  | 
 **lastUpdatedGt** | [**[]time.Time**](time.Time.md) |  | 
 **lastUpdatedGte** | [**[]time.Time**](time.Time.md) |  | 
 **lastUpdatedLt** | [**[]time.Time**](time.Time.md) |  | 
 **lastUpdatedLte** | [**[]time.Time**](time.Time.md) |  | 
 **lastUpdatedN** | [**[]time.Time**](time.Time.md) |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **modifiedByRequest** | **string** |  | 
 **name** | **string** |  | 
 **objectId** | **[]int32** |  | 
 **objectIdEmpty** | **bool** |  | 
 **objectIdGt** | **[]int32** |  | 
 **objectIdGte** | **[]int32** |  | 
 **objectIdLt** | **[]int32** |  | 
 **objectIdLte** | **[]int32** |  | 
 **objectIdN** | **[]int32** |  | 
 **objectType** | **string** |  | 
 **objectTypeN** | **string** |  | 
 **objectTypeId** | **int32** |  | 
 **objectTypeIdN** | **int32** |  | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **q** | **string** | Search | 
 **tag** | **[]string** |  | 
 **tagN** | **[]string** |  | 
 **updatedByRequest** | **string** |  | 

### Return type

[**PaginatedNetBoxAttachmentList**](PaginatedNetBoxAttachmentList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PluginsNetboxAttachmentsNetboxAttachmentsPartialUpdate

> NetBoxAttachment PluginsNetboxAttachmentsNetboxAttachmentsPartialUpdate(ctx, id).PatchedNetBoxAttachmentRequest(patchedNetBoxAttachmentRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this NetBox Attachment.
	patchedNetBoxAttachmentRequest := *openapiclient.NewPatchedNetBoxAttachmentRequest() // PatchedNetBoxAttachmentRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginsAPI.PluginsNetboxAttachmentsNetboxAttachmentsPartialUpdate(context.Background(), id).PatchedNetBoxAttachmentRequest(patchedNetBoxAttachmentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.PluginsNetboxAttachmentsNetboxAttachmentsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PluginsNetboxAttachmentsNetboxAttachmentsPartialUpdate`: NetBoxAttachment
	fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.PluginsNetboxAttachmentsNetboxAttachmentsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this NetBox Attachment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPluginsNetboxAttachmentsNetboxAttachmentsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedNetBoxAttachmentRequest** | [**PatchedNetBoxAttachmentRequest**](PatchedNetBoxAttachmentRequest.md) |  | 

### Return type

[**NetBoxAttachment**](NetBoxAttachment.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PluginsNetboxAttachmentsNetboxAttachmentsRetrieve

> NetBoxAttachment PluginsNetboxAttachmentsNetboxAttachmentsRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this NetBox Attachment.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginsAPI.PluginsNetboxAttachmentsNetboxAttachmentsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.PluginsNetboxAttachmentsNetboxAttachmentsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PluginsNetboxAttachmentsNetboxAttachmentsRetrieve`: NetBoxAttachment
	fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.PluginsNetboxAttachmentsNetboxAttachmentsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this NetBox Attachment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPluginsNetboxAttachmentsNetboxAttachmentsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetBoxAttachment**](NetBoxAttachment.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PluginsNetboxAttachmentsNetboxAttachmentsUpdate

> NetBoxAttachment PluginsNetboxAttachmentsNetboxAttachmentsUpdate(ctx, id).NetBoxAttachmentRequest(netBoxAttachmentRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this NetBox Attachment.
	netBoxAttachmentRequest := *openapiclient.NewNetBoxAttachmentRequest("ObjectType_example", int64(123), "TODO") // NetBoxAttachmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginsAPI.PluginsNetboxAttachmentsNetboxAttachmentsUpdate(context.Background(), id).NetBoxAttachmentRequest(netBoxAttachmentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.PluginsNetboxAttachmentsNetboxAttachmentsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PluginsNetboxAttachmentsNetboxAttachmentsUpdate`: NetBoxAttachment
	fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.PluginsNetboxAttachmentsNetboxAttachmentsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this NetBox Attachment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPluginsNetboxAttachmentsNetboxAttachmentsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **netBoxAttachmentRequest** | [**NetBoxAttachmentRequest**](NetBoxAttachmentRequest.md) |  | 

### Return type

[**NetBoxAttachment**](NetBoxAttachment.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PluginsNetboxTopologyViewsImagesList

> PaginatedRoleImageList PluginsNetboxTopologyViewsImagesList(ctx).Limit(limit).Offset(offset).Ordering(ordering).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginsAPI.PluginsNetboxTopologyViewsImagesList(context.Background()).Limit(limit).Offset(offset).Ordering(ordering).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.PluginsNetboxTopologyViewsImagesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PluginsNetboxTopologyViewsImagesList`: PaginatedRoleImageList
	fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.PluginsNetboxTopologyViewsImagesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPluginsNetboxTopologyViewsImagesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **string** | Which field to use when ordering the results. | 

### Return type

[**PaginatedRoleImageList**](PaginatedRoleImageList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PluginsNetboxTopologyViewsImagesRetrieve

> RoleImage PluginsNetboxTopologyViewsImagesRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this device role.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginsAPI.PluginsNetboxTopologyViewsImagesRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.PluginsNetboxTopologyViewsImagesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PluginsNetboxTopologyViewsImagesRetrieve`: RoleImage
	fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.PluginsNetboxTopologyViewsImagesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this device role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPluginsNetboxTopologyViewsImagesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RoleImage**](RoleImage.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PluginsNetboxTopologyViewsImagesSaveCreate

> RoleImage PluginsNetboxTopologyViewsImagesSaveCreate(ctx).RoleImageRequest(roleImageRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	roleImageRequest := *openapiclient.NewRoleImageRequest("Image_example") // RoleImageRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginsAPI.PluginsNetboxTopologyViewsImagesSaveCreate(context.Background()).RoleImageRequest(roleImageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.PluginsNetboxTopologyViewsImagesSaveCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PluginsNetboxTopologyViewsImagesSaveCreate`: RoleImage
	fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.PluginsNetboxTopologyViewsImagesSaveCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPluginsNetboxTopologyViewsImagesSaveCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roleImageRequest** | [**RoleImageRequest**](RoleImageRequest.md) |  | 

### Return type

[**RoleImage**](RoleImage.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PluginsNetboxTopologyViewsSaveCoordsList

> PaginatedTopologyDummyList PluginsNetboxTopologyViewsSaveCoordsList(ctx).Limit(limit).Offset(offset).Ordering(ordering).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginsAPI.PluginsNetboxTopologyViewsSaveCoordsList(context.Background()).Limit(limit).Offset(offset).Ordering(ordering).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.PluginsNetboxTopologyViewsSaveCoordsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PluginsNetboxTopologyViewsSaveCoordsList`: PaginatedTopologyDummyList
	fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.PluginsNetboxTopologyViewsSaveCoordsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPluginsNetboxTopologyViewsSaveCoordsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **string** | Which field to use when ordering the results. | 

### Return type

[**PaginatedTopologyDummyList**](PaginatedTopologyDummyList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PluginsNetboxTopologyViewsSaveCoordsRetrieve

> TopologyDummy PluginsNetboxTopologyViewsSaveCoordsRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this device.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginsAPI.PluginsNetboxTopologyViewsSaveCoordsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.PluginsNetboxTopologyViewsSaveCoordsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PluginsNetboxTopologyViewsSaveCoordsRetrieve`: TopologyDummy
	fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.PluginsNetboxTopologyViewsSaveCoordsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPluginsNetboxTopologyViewsSaveCoordsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TopologyDummy**](TopologyDummy.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PluginsNetboxTopologyViewsSaveCoordsSaveCoordsPartialUpdate

> TopologyDummy PluginsNetboxTopologyViewsSaveCoordsSaveCoordsPartialUpdate(ctx).PatchedTopologyDummyRequest(patchedTopologyDummyRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	patchedTopologyDummyRequest := *openapiclient.NewPatchedTopologyDummyRequest() // PatchedTopologyDummyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginsAPI.PluginsNetboxTopologyViewsSaveCoordsSaveCoordsPartialUpdate(context.Background()).PatchedTopologyDummyRequest(patchedTopologyDummyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.PluginsNetboxTopologyViewsSaveCoordsSaveCoordsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PluginsNetboxTopologyViewsSaveCoordsSaveCoordsPartialUpdate`: TopologyDummy
	fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.PluginsNetboxTopologyViewsSaveCoordsSaveCoordsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPluginsNetboxTopologyViewsSaveCoordsSaveCoordsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedTopologyDummyRequest** | [**PatchedTopologyDummyRequest**](PatchedTopologyDummyRequest.md) |  | 

### Return type

[**TopologyDummy**](TopologyDummy.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PluginsNetboxTopologyViewsXmlExportList

> PaginatedTopologyDummyList PluginsNetboxTopologyViewsXmlExportList(ctx).Limit(limit).Offset(offset).Ordering(ordering).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginsAPI.PluginsNetboxTopologyViewsXmlExportList(context.Background()).Limit(limit).Offset(offset).Ordering(ordering).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.PluginsNetboxTopologyViewsXmlExportList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PluginsNetboxTopologyViewsXmlExportList`: PaginatedTopologyDummyList
	fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.PluginsNetboxTopologyViewsXmlExportList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPluginsNetboxTopologyViewsXmlExportListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **string** | Which field to use when ordering the results. | 

### Return type

[**PaginatedTopologyDummyList**](PaginatedTopologyDummyList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

