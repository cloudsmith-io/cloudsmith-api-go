# \MetricsApi

All URIs are relative to *https://api.cloudsmith.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MetricsEntitlementsList**](MetricsApi.md#MetricsEntitlementsList) | **Get** /metrics/entitlements/{owner}/ | View for listing entitlement token metrics, across an account.
[**MetricsEntitlementsList0**](MetricsApi.md#MetricsEntitlementsList0) | **Get** /metrics/entitlements/{owner}/{repo}/ | View for listing entitlement token metrics, for a repository.
[**MetricsPackagesList**](MetricsApi.md#MetricsPackagesList) | **Get** /metrics/packages/{owner}/{repo}/ | View for listing package usage metrics, for a repository.



## MetricsEntitlementsList

> EntitlementUsageMetrics MetricsEntitlementsList(ctx, owner).Page(page).PageSize(pageSize).Finish(finish).Start(start).Tokens(tokens).Execute()

View for listing entitlement token metrics, across an account.



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
    finish := "finish_example" // string | Include metrics upto and including this UTC date or UTC datetime. For example '2020-12-31' or '2021-12-13T00:00:00Z'. (optional)
    start := "start_example" // string | Include metrics from and including this UTC date or UTC datetime. For example '2020-12-31' or '2021-12-13T00:00:00Z'. (optional)
    tokens := "tokens_example" // string | A comma seperated list of tokens (slug perm) to include in the results. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsApi.MetricsEntitlementsList(context.Background(), owner).Page(page).PageSize(pageSize).Finish(finish).Start(start).Tokens(tokens).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.MetricsEntitlementsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MetricsEntitlementsList`: EntitlementUsageMetrics
    fmt.Fprintf(os.Stdout, "Response from `MetricsApi.MetricsEntitlementsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMetricsEntitlementsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 
 **finish** | **string** | Include metrics upto and including this UTC date or UTC datetime. For example &#39;2020-12-31&#39; or &#39;2021-12-13T00:00:00Z&#39;. | 
 **start** | **string** | Include metrics from and including this UTC date or UTC datetime. For example &#39;2020-12-31&#39; or &#39;2021-12-13T00:00:00Z&#39;. | 
 **tokens** | **string** | A comma seperated list of tokens (slug perm) to include in the results. | 

### Return type

[**EntitlementUsageMetrics**](EntitlementUsageMetrics.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsEntitlementsList0

> EntitlementUsageMetrics MetricsEntitlementsList0(ctx, owner, repo).Page(page).PageSize(pageSize).Finish(finish).Start(start).Tokens(tokens).Execute()

View for listing entitlement token metrics, for a repository.



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
    finish := "finish_example" // string | Include metrics upto and including this UTC date or UTC datetime. For example '2020-12-31' or '2021-12-13T00:00:00Z'. (optional)
    start := "start_example" // string | Include metrics from and including this UTC date or UTC datetime. For example '2020-12-31' or '2021-12-13T00:00:00Z'. (optional)
    tokens := "tokens_example" // string | A comma seperated list of tokens (slug perm) to include in the results. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsApi.MetricsEntitlementsList0(context.Background(), owner, repo).Page(page).PageSize(pageSize).Finish(finish).Start(start).Tokens(tokens).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.MetricsEntitlementsList0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MetricsEntitlementsList0`: EntitlementUsageMetrics
    fmt.Fprintf(os.Stdout, "Response from `MetricsApi.MetricsEntitlementsList0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMetricsEntitlementsList0Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 
 **finish** | **string** | Include metrics upto and including this UTC date or UTC datetime. For example &#39;2020-12-31&#39; or &#39;2021-12-13T00:00:00Z&#39;. | 
 **start** | **string** | Include metrics from and including this UTC date or UTC datetime. For example &#39;2020-12-31&#39; or &#39;2021-12-13T00:00:00Z&#39;. | 
 **tokens** | **string** | A comma seperated list of tokens (slug perm) to include in the results. | 

### Return type

[**EntitlementUsageMetrics**](EntitlementUsageMetrics.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricsPackagesList

> PackageUsageMetrics MetricsPackagesList(ctx, owner, repo).Page(page).PageSize(pageSize).Finish(finish).Packages(packages).Start(start).Execute()

View for listing package usage metrics, for a repository.



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
    finish := "finish_example" // string | Include metrics upto and including this UTC date or UTC datetime. For example '2020-12-31' or '2021-12-13T00:00:00Z'. (optional)
    packages := "packages_example" // string | A comma seperated list of packages (slug perm) to include in the results. (optional)
    start := "start_example" // string | Include metrics from and including this UTC date or UTC datetime. For example '2020-12-31' or '2021-12-13T00:00:00Z'. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsApi.MetricsPackagesList(context.Background(), owner, repo).Page(page).PageSize(pageSize).Finish(finish).Packages(packages).Start(start).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.MetricsPackagesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MetricsPackagesList`: PackageUsageMetrics
    fmt.Fprintf(os.Stdout, "Response from `MetricsApi.MetricsPackagesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMetricsPackagesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 
 **finish** | **string** | Include metrics upto and including this UTC date or UTC datetime. For example &#39;2020-12-31&#39; or &#39;2021-12-13T00:00:00Z&#39;. | 
 **packages** | **string** | A comma seperated list of packages (slug perm) to include in the results. | 
 **start** | **string** | Include metrics from and including this UTC date or UTC datetime. For example &#39;2020-12-31&#39; or &#39;2021-12-13T00:00:00Z&#39;. | 

### Return type

[**PackageUsageMetrics**](PackageUsageMetrics.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

