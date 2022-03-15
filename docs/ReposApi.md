# \ReposApi

All URIs are relative to *https://api.cloudsmith.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReposAllList**](ReposApi.md#ReposAllList) | **Get** /repos/ | Get a list of all repositories associated with current user.
[**ReposCreate**](ReposApi.md#ReposCreate) | **Post** /repos/{owner}/ | Create a new repository in a given namespace.
[**ReposDelete**](ReposApi.md#ReposDelete) | **Delete** /repos/{owner}/{identifier}/ | Delete a repository in a given namespace.
[**ReposList**](ReposApi.md#ReposList) | **Get** /repos/{owner}/ | Get a list of all repositories within a namespace.
[**ReposPartialUpdate**](ReposApi.md#ReposPartialUpdate) | **Patch** /repos/{owner}/{identifier}/ | Update details about a repository in a given namespace.
[**ReposRead**](ReposApi.md#ReposRead) | **Get** /repos/{owner}/{identifier}/ | Get a specific repository.



## ReposAllList

> []Repository ReposAllList(ctx).Page(page).PageSize(pageSize).Execute()

Get a list of all repositories associated with current user.



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
    page := int64(56) // int64 | A page number within the paginated result set. (optional)
    pageSize := int64(56) // int64 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposAllList(context.Background()).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposAllList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposAllList`: []Repository
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposAllList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReposAllListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 

### Return type

[**[]Repository**](Repository.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposCreate

> RepositoryCreate ReposCreate(ctx, owner).Data(data).Execute()

Create a new repository in a given namespace.



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
    data := *openapiclient.NewReposCreate("Name_example") // ReposCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposCreate(context.Background(), owner).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposCreate`: RepositoryCreate
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**ReposCreate**](ReposCreate.md) |  | 

### Return type

[**RepositoryCreate**](RepositoryCreate.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposDelete

> ReposDelete(ctx, owner, identifier).Execute()

Delete a repository in a given namespace.



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
    identifier := "identifier_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposDelete(context.Background(), owner, identifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposList

> []Repository ReposList(ctx, owner).Page(page).PageSize(pageSize).Execute()

Get a list of all repositories within a namespace.



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
    resp, r, err := apiClient.ReposApi.ReposList(context.Background(), owner).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposList`: []Repository
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 

### Return type

[**[]Repository**](Repository.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposPartialUpdate

> Repository ReposPartialUpdate(ctx, owner, identifier).Data(data).Execute()

Update details about a repository in a given namespace.



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
    identifier := "identifier_example" // string | 
    data := *openapiclient.NewReposPartialUpdate() // ReposPartialUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposPartialUpdate(context.Background(), owner, identifier).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposPartialUpdate`: Repository
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**ReposPartialUpdate**](ReposPartialUpdate.md) |  | 

### Return type

[**Repository**](Repository.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposRead

> Repository ReposRead(ctx, owner, identifier).Execute()

Get a specific repository.



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
    identifier := "identifier_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposRead(context.Background(), owner, identifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposRead`: Repository
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Repository**](Repository.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

