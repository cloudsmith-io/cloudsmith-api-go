# \VulnerabilitiesApi

All URIs are relative to *https://api.cloudsmith.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VulnerabilitiesNamespaceList**](VulnerabilitiesApi.md#VulnerabilitiesNamespaceList) | **Get** /vulnerabilities/{owner}/ | Lists scan results for a specific namespace.
[**VulnerabilitiesPackageList**](VulnerabilitiesApi.md#VulnerabilitiesPackageList) | **Get** /vulnerabilities/{owner}/{repo}/{package}/ | Lists scan results for a specific package.
[**VulnerabilitiesRead**](VulnerabilitiesApi.md#VulnerabilitiesRead) | **Get** /vulnerabilities/{owner}/{repo}/{package}/{identifier}/ | Get a scan result.
[**VulnerabilitiesRepoList**](VulnerabilitiesApi.md#VulnerabilitiesRepoList) | **Get** /vulnerabilities/{owner}/{repo}/ | Lists scan results for a specific repository.



## VulnerabilitiesNamespaceList

> []VulnerabilityScanResultsList VulnerabilitiesNamespaceList(ctx, owner).Page(page).PageSize(pageSize).Execute()

Lists scan results for a specific namespace.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VulnerabilitiesApi.VulnerabilitiesNamespaceList(context.Background(), owner).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VulnerabilitiesApi.VulnerabilitiesNamespaceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VulnerabilitiesNamespaceList`: []VulnerabilityScanResultsList
	fmt.Fprintf(os.Stdout, "Response from `VulnerabilitiesApi.VulnerabilitiesNamespaceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVulnerabilitiesNamespaceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 

### Return type

[**[]VulnerabilityScanResultsList**](VulnerabilityScanResultsList.md)

### Authorization

[apikey](../README.md#apikey), [basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VulnerabilitiesPackageList

> []VulnerabilityScanResultsList VulnerabilitiesPackageList(ctx, owner, repo, package_).Page(page).PageSize(pageSize).Execute()

Lists scan results for a specific package.



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
	repo := "repo_example" // string | 
	package_ := "package__example" // string | 
	page := int64(56) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(56) // int64 | Number of results to return per page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VulnerabilitiesApi.VulnerabilitiesPackageList(context.Background(), owner, repo, package_).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VulnerabilitiesApi.VulnerabilitiesPackageList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VulnerabilitiesPackageList`: []VulnerabilityScanResultsList
	fmt.Fprintf(os.Stdout, "Response from `VulnerabilitiesApi.VulnerabilitiesPackageList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 
**package_** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVulnerabilitiesPackageListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 

### Return type

[**[]VulnerabilityScanResultsList**](VulnerabilityScanResultsList.md)

### Authorization

[apikey](../README.md#apikey), [basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VulnerabilitiesRead

> VulnerabilityScanResults VulnerabilitiesRead(ctx, owner, repo, package_, identifier).Execute()

Get a scan result.



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
	repo := "repo_example" // string | 
	package_ := "package__example" // string | 
	identifier := "identifier_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VulnerabilitiesApi.VulnerabilitiesRead(context.Background(), owner, repo, package_, identifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VulnerabilitiesApi.VulnerabilitiesRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VulnerabilitiesRead`: VulnerabilityScanResults
	fmt.Fprintf(os.Stdout, "Response from `VulnerabilitiesApi.VulnerabilitiesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 
**package_** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVulnerabilitiesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**VulnerabilityScanResults**](VulnerabilityScanResults.md)

### Authorization

[apikey](../README.md#apikey), [basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VulnerabilitiesRepoList

> []VulnerabilityScanResultsList VulnerabilitiesRepoList(ctx, owner, repo).Page(page).PageSize(pageSize).Execute()

Lists scan results for a specific repository.



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
	repo := "repo_example" // string | 
	page := int64(56) // int64 | A page number within the paginated result set. (optional)
	pageSize := int64(56) // int64 | Number of results to return per page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VulnerabilitiesApi.VulnerabilitiesRepoList(context.Background(), owner, repo).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VulnerabilitiesApi.VulnerabilitiesRepoList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VulnerabilitiesRepoList`: []VulnerabilityScanResultsList
	fmt.Fprintf(os.Stdout, "Response from `VulnerabilitiesApi.VulnerabilitiesRepoList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVulnerabilitiesRepoListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 

### Return type

[**[]VulnerabilityScanResultsList**](VulnerabilityScanResultsList.md)

### Authorization

[apikey](../README.md#apikey), [basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

