# \RecycleBinApi

All URIs are relative to *https://api.cloudsmith.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RecycleBinAction**](RecycleBinApi.md#RecycleBinAction) | **Post** /recycle-bin/{owner}/action/ | 
[**RecycleBinList**](RecycleBinApi.md#RecycleBinList) | **Get** /recycle-bin/{owner}/ | List soft-deleted packages in recycle bin



## RecycleBinAction

> PackageBulkActionResponse RecycleBinAction(ctx, owner).Data(data).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
	owner := "owner_example" // string | 
	data := *openapiclient.NewPackageRecycleBin("Action_example", []string{"Identifiers_example"}) // PackageRecycleBin |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecycleBinApi.RecycleBinAction(context.Background(), owner).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecycleBinApi.RecycleBinAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecycleBinAction`: PackageBulkActionResponse
	fmt.Fprintf(os.Stdout, "Response from `RecycleBinApi.RecycleBinAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecycleBinActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**PackageRecycleBin**](PackageRecycleBin.md) |  | 

### Return type

[**PackageBulkActionResponse**](PackageBulkActionResponse.md)

### Authorization

[apikey](../README.md#apikey), [basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecycleBinList

> []RecycleBinPackage RecycleBinList(ctx, owner).Page(page).PageSize(pageSize).Repository(repository).Execute()

List soft-deleted packages in recycle bin



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cloudsmith-io/cloudsmith-api-go"
)

func main() {
	owner := "owner_example" // string | 
	page := int64(56) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(56) // int64 | Number of results to return per page. (optional)
	repository := "repository_example" // string | Filter packages by repository slug (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecycleBinApi.RecycleBinList(context.Background(), owner).Page(page).PageSize(pageSize).Repository(repository).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecycleBinApi.RecycleBinList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecycleBinList`: []RecycleBinPackage
	fmt.Fprintf(os.Stdout, "Response from `RecycleBinApi.RecycleBinList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecycleBinListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 
 **repository** | **string** | Filter packages by repository slug | 

### Return type

[**[]RecycleBinPackage**](RecycleBinPackage.md)

### Authorization

[apikey](../README.md#apikey), [basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

