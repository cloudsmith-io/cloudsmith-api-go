# \VulnerabilitiesApi

All URIs are relative to *https://api.cloudsmith.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VulnerabilitiesList**](VulnerabilitiesApi.md#VulnerabilitiesList) | **Get** /vulnerabilities/{owner}/ | Lists scan results for a specific namespace.
[**VulnerabilitiesList0**](VulnerabilitiesApi.md#VulnerabilitiesList0) | **Get** /vulnerabilities/{owner}/{repo}/ | Lists scan results for a specific repository.
[**VulnerabilitiesList1**](VulnerabilitiesApi.md#VulnerabilitiesList1) | **Get** /vulnerabilities/{owner}/{repo}/{package}/ | Lists scan results for a specific package.
[**VulnerabilitiesRead**](VulnerabilitiesApi.md#VulnerabilitiesRead) | **Get** /vulnerabilities/{owner}/{repo}/{package}/{scan_id}/ | Returns a Scan Result.



## VulnerabilitiesList

> []VulnerabilityScanResultsList VulnerabilitiesList(ctx, owner).Page(page).PageSize(pageSize).Execute()

Lists scan results for a specific namespace.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | 
    page := int64(56) // int64 | A page number within the paginated result set. (optional)
    pageSize := int64(56) // int64 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VulnerabilitiesApi.VulnerabilitiesList(context.Background(), owner).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VulnerabilitiesApi.VulnerabilitiesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VulnerabilitiesList`: []VulnerabilityScanResultsList
    fmt.Fprintf(os.Stdout, "Response from `VulnerabilitiesApi.VulnerabilitiesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVulnerabilitiesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 

### Return type

[**[]VulnerabilityScanResultsList**](VulnerabilityScanResultsList.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VulnerabilitiesList0

> []VulnerabilityScanResultsList VulnerabilitiesList0(ctx, owner, repo).Page(page).PageSize(pageSize).Execute()

Lists scan results for a specific repository.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    page := int64(56) // int64 | A page number within the paginated result set. (optional)
    pageSize := int64(56) // int64 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VulnerabilitiesApi.VulnerabilitiesList0(context.Background(), owner, repo).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VulnerabilitiesApi.VulnerabilitiesList0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VulnerabilitiesList0`: []VulnerabilityScanResultsList
    fmt.Fprintf(os.Stdout, "Response from `VulnerabilitiesApi.VulnerabilitiesList0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVulnerabilitiesList0Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 

### Return type

[**[]VulnerabilityScanResultsList**](VulnerabilityScanResultsList.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VulnerabilitiesList1

> []VulnerabilityScanResultsList VulnerabilitiesList1(ctx, owner, repo, package_).Page(page).PageSize(pageSize).Execute()

Lists scan results for a specific package.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    package_ := "package__example" // string | 
    page := int64(56) // int64 | A page number within the paginated result set. (optional)
    pageSize := int64(56) // int64 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VulnerabilitiesApi.VulnerabilitiesList1(context.Background(), owner, repo, package_).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VulnerabilitiesApi.VulnerabilitiesList1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VulnerabilitiesList1`: []VulnerabilityScanResultsList
    fmt.Fprintf(os.Stdout, "Response from `VulnerabilitiesApi.VulnerabilitiesList1`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiVulnerabilitiesList1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 

### Return type

[**[]VulnerabilityScanResultsList**](VulnerabilityScanResultsList.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VulnerabilitiesRead

> VulnerabilityScanResults VulnerabilitiesRead(ctx, owner, repo, package_, scanId).Execute()

Returns a Scan Result.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | 
    repo := "repo_example" // string | 
    package_ := "package__example" // string | 
    scanId := "scanId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VulnerabilitiesApi.VulnerabilitiesRead(context.Background(), owner, repo, package_, scanId).Execute()
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
**scanId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVulnerabilitiesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**VulnerabilityScanResults**](VulnerabilityScanResults.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

