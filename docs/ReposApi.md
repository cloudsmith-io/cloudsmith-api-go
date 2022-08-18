# \ReposApi

All URIs are relative to *https://api.cloudsmith.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReposAllList**](ReposApi.md#ReposAllList) | **Get** /repos/ | Get a list of all repositories associated with current user.
[**ReposCreate**](ReposApi.md#ReposCreate) | **Post** /repos/{owner}/ | Create a new repository in a given namespace.
[**ReposDelete**](ReposApi.md#ReposDelete) | **Delete** /repos/{owner}/{identifier}/ | Delete a repository in a given namespace.
[**ReposGpgCreate**](ReposApi.md#ReposGpgCreate) | **Post** /repos/{owner}/{identifier}/gpg/ | Set the active GPG key for the Repository.
[**ReposGpgList**](ReposApi.md#ReposGpgList) | **Get** /repos/{owner}/{identifier}/gpg/ | Retrieve the active GPG key for the Repository.
[**ReposGpgRegenerate**](ReposApi.md#ReposGpgRegenerate) | **Post** /repos/{owner}/{identifier}/gpg/regenerate/ | Regenerate GPG Key for the Repository.
[**ReposList**](ReposApi.md#ReposList) | **Get** /repos/{owner}/ | Get a list of all repositories within a namespace.
[**ReposPartialUpdate**](ReposApi.md#ReposPartialUpdate) | **Patch** /repos/{owner}/{identifier}/ | Update details about a repository in a given namespace.
[**ReposPrivilegesDelete**](ReposApi.md#ReposPrivilegesDelete) | **Delete** /repos/{owner}/{identifier}/privileges | Remove the specified repository privileges.
[**ReposPrivilegesList**](ReposApi.md#ReposPrivilegesList) | **Get** /repos/{owner}/{identifier}/privileges | List all explicity created privileges for the repository.
[**ReposPrivilegesPartialUpdate**](ReposApi.md#ReposPrivilegesPartialUpdate) | **Patch** /repos/{owner}/{identifier}/privileges | Update the specified repository privileges.
[**ReposPrivilegesUpdate**](ReposApi.md#ReposPrivilegesUpdate) | **Put** /repos/{owner}/{identifier}/privileges | Replace all existing repository privileges with those specified.
[**ReposRead**](ReposApi.md#ReposRead) | **Get** /repos/{owner}/{identifier}/ | Get a specific repository.
[**ReposRsaCreate**](ReposApi.md#ReposRsaCreate) | **Post** /repos/{owner}/{identifier}/rsa/ | Set the active RSA key for the Repository.
[**ReposRsaList**](ReposApi.md#ReposRsaList) | **Get** /repos/{owner}/{identifier}/rsa/ | Retrieve the active RSA key for the Repository.
[**ReposRsaRegenerate**](ReposApi.md#ReposRsaRegenerate) | **Post** /repos/{owner}/{identifier}/rsa/regenerate/ | Regenerate RSA Key for the Repository.



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


## ReposGpgCreate

> RepositoryGpgKey ReposGpgCreate(ctx, owner, identifier).Data(data).Execute()

Set the active GPG key for the Repository.



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
    data := *openapiclient.NewReposGpgCreate("GpgPrivateKey_example") // ReposGpgCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGpgCreate(context.Background(), owner, identifier).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGpgCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGpgCreate`: RepositoryGpgKey
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGpgCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGpgCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**ReposGpgCreate**](ReposGpgCreate.md) |  | 

### Return type

[**RepositoryGpgKey**](RepositoryGpgKey.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGpgList

> RepositoryGpgKey ReposGpgList(ctx, owner, identifier).Execute()

Retrieve the active GPG key for the Repository.



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
    resp, r, err := apiClient.ReposApi.ReposGpgList(context.Background(), owner, identifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGpgList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGpgList`: RepositoryGpgKey
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGpgList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGpgListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RepositoryGpgKey**](RepositoryGpgKey.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGpgRegenerate

> RepositoryGpgKey ReposGpgRegenerate(ctx, owner, identifier).Execute()

Regenerate GPG Key for the Repository.



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
    resp, r, err := apiClient.ReposApi.ReposGpgRegenerate(context.Background(), owner, identifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGpgRegenerate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGpgRegenerate`: RepositoryGpgKey
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGpgRegenerate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGpgRegenerateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RepositoryGpgKey**](RepositoryGpgKey.md)

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


## ReposPrivilegesDelete

> ReposPrivilegesDelete(ctx, owner, identifier).Execute()

Remove the specified repository privileges.



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
    resp, r, err := apiClient.ReposApi.ReposPrivilegesDelete(context.Background(), owner, identifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposPrivilegesDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReposPrivilegesDeleteRequest struct via the builder pattern


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


## ReposPrivilegesList

> []RepositoryPrivilegeList ReposPrivilegesList(ctx, owner, identifier).Execute()

List all explicity created privileges for the repository.



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
    resp, r, err := apiClient.ReposApi.ReposPrivilegesList(context.Background(), owner, identifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposPrivilegesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposPrivilegesList`: []RepositoryPrivilegeList
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposPrivilegesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposPrivilegesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]RepositoryPrivilegeList**](RepositoryPrivilegeList.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposPrivilegesPartialUpdate

> ReposPrivilegesPartialUpdate(ctx, owner, identifier).Execute()

Update the specified repository privileges.



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
    resp, r, err := apiClient.ReposApi.ReposPrivilegesPartialUpdate(context.Background(), owner, identifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposPrivilegesPartialUpdate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReposPrivilegesPartialUpdateRequest struct via the builder pattern


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


## ReposPrivilegesUpdate

> RepositoryPrivilegeList ReposPrivilegesUpdate(ctx, owner, identifier).Execute()

Replace all existing repository privileges with those specified.



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
    resp, r, err := apiClient.ReposApi.ReposPrivilegesUpdate(context.Background(), owner, identifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposPrivilegesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposPrivilegesUpdate`: RepositoryPrivilegeList
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposPrivilegesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposPrivilegesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RepositoryPrivilegeList**](RepositoryPrivilegeList.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
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


## ReposRsaCreate

> RepositoryRsaKey ReposRsaCreate(ctx, owner, identifier).Data(data).Execute()

Set the active RSA key for the Repository.



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
    data := *openapiclient.NewReposRsaCreate("RsaPrivateKey_example") // ReposRsaCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposRsaCreate(context.Background(), owner, identifier).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposRsaCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposRsaCreate`: RepositoryRsaKey
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposRsaCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposRsaCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**ReposRsaCreate**](ReposRsaCreate.md) |  | 

### Return type

[**RepositoryRsaKey**](RepositoryRsaKey.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposRsaList

> RepositoryRsaKey ReposRsaList(ctx, owner, identifier).Execute()

Retrieve the active RSA key for the Repository.



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
    resp, r, err := apiClient.ReposApi.ReposRsaList(context.Background(), owner, identifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposRsaList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposRsaList`: RepositoryRsaKey
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposRsaList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposRsaListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RepositoryRsaKey**](RepositoryRsaKey.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposRsaRegenerate

> RepositoryRsaKey ReposRsaRegenerate(ctx, owner, identifier).Execute()

Regenerate RSA Key for the Repository.



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
    resp, r, err := apiClient.ReposApi.ReposRsaRegenerate(context.Background(), owner, identifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposRsaRegenerate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposRsaRegenerate`: RepositoryRsaKey
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposRsaRegenerate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposRsaRegenerateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RepositoryRsaKey**](RepositoryRsaKey.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

