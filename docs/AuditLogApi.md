# \AuditLogApi

All URIs are relative to *https://api.cloudsmith.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuditLogNamespaceList**](AuditLogApi.md#AuditLogNamespaceList) | **Get** /audit-log/{owner}/ | Lists audit log entries for a specific namespace.
[**AuditLogRepoList**](AuditLogApi.md#AuditLogRepoList) | **Get** /audit-log/{owner}/{repo}/ | Lists audit log entries for a specific repository.



## AuditLogNamespaceList

> []NamespaceAuditLog AuditLogNamespaceList(ctx, owner).Page(page).PageSize(pageSize).Query(query).Execute()

Lists audit log entries for a specific namespace.



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
    query := "query_example" // string | A search term for querying events, actors, or timestamps of log records. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuditLogApi.AuditLogNamespaceList(context.Background(), owner).Page(page).PageSize(pageSize).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLogApi.AuditLogNamespaceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuditLogNamespaceList`: []NamespaceAuditLog
    fmt.Fprintf(os.Stdout, "Response from `AuditLogApi.AuditLogNamespaceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditLogNamespaceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 
 **query** | **string** | A search term for querying events, actors, or timestamps of log records. | 

### Return type

[**[]NamespaceAuditLog**](NamespaceAuditLog.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditLogRepoList

> []RepositoryAuditLog AuditLogRepoList(ctx, owner, repo).Page(page).PageSize(pageSize).Query(query).Execute()

Lists audit log entries for a specific repository.



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
    query := "query_example" // string | A search term for querying events, actors, or timestamps of log records. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuditLogApi.AuditLogRepoList(context.Background(), owner, repo).Page(page).PageSize(pageSize).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLogApi.AuditLogRepoList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuditLogRepoList`: []RepositoryAuditLog
    fmt.Fprintf(os.Stdout, "Response from `AuditLogApi.AuditLogRepoList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditLogRepoListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 
 **query** | **string** | A search term for querying events, actors, or timestamps of log records. | 

### Return type

[**[]RepositoryAuditLog**](RepositoryAuditLog.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

