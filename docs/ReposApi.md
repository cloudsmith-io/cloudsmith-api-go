# \ReposApi

All URIs are relative to *https://api.cloudsmith.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReposCreate**](ReposApi.md#ReposCreate) | **Post** /repos/{owner}/ | Create a new repository in a given namespace.
[**ReposDelete**](ReposApi.md#ReposDelete) | **Delete** /repos/{owner}/{identifier}/ | Delete a repository in a given namespace.
[**ReposGpgCreate**](ReposApi.md#ReposGpgCreate) | **Post** /repos/{owner}/{identifier}/gpg/ | Set the active GPG key for the Repository.
[**ReposGpgList**](ReposApi.md#ReposGpgList) | **Get** /repos/{owner}/{identifier}/gpg/ | Retrieve the active GPG key for the Repository.
[**ReposGpgRegenerate**](ReposApi.md#ReposGpgRegenerate) | **Post** /repos/{owner}/{identifier}/gpg/regenerate/ | Regenerate GPG Key for the Repository.
[**ReposNamespaceList**](ReposApi.md#ReposNamespaceList) | **Get** /repos/{owner}/ | Get a list of all repositories within a namespace.
[**ReposPartialUpdate**](ReposApi.md#ReposPartialUpdate) | **Patch** /repos/{owner}/{identifier}/ | Update details about a repository in a given namespace.
[**ReposPrivilegesList**](ReposApi.md#ReposPrivilegesList) | **Get** /repos/{owner}/{identifier}/privileges | List all explicity created privileges for the repository.
[**ReposPrivilegesPartialUpdate**](ReposApi.md#ReposPrivilegesPartialUpdate) | **Patch** /repos/{owner}/{identifier}/privileges | Modify privileges for the repository.
[**ReposPrivilegesUpdate**](ReposApi.md#ReposPrivilegesUpdate) | **Put** /repos/{owner}/{identifier}/privileges | Replace all existing repository privileges with those specified.
[**ReposRead**](ReposApi.md#ReposRead) | **Get** /repos/{owner}/{identifier}/ | Get a specific repository.
[**ReposRsaCreate**](ReposApi.md#ReposRsaCreate) | **Post** /repos/{owner}/{identifier}/rsa/ | Set the active RSA key for the Repository.
[**ReposRsaList**](ReposApi.md#ReposRsaList) | **Get** /repos/{owner}/{identifier}/rsa/ | Retrieve the active RSA key for the Repository.
[**ReposRsaRegenerate**](ReposApi.md#ReposRsaRegenerate) | **Post** /repos/{owner}/{identifier}/rsa/regenerate/ | Regenerate RSA Key for the Repository.
[**ReposUserList**](ReposApi.md#ReposUserList) | **Get** /repos/ | Get a list of all repositories associated with current user.



## ReposCreate

> RepositoryCreateResponse ReposCreate(ctx, owner).Data(data).Execute()

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
    data := *openapiclient.NewRepositoryCreateRequest("Name_example") // RepositoryCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposCreate(context.Background(), owner).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposCreate`: RepositoryCreateResponse
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

 **data** | [**RepositoryCreateRequest**](RepositoryCreateRequest.md) |  | 

### Return type

[**RepositoryCreateResponse**](RepositoryCreateResponse.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGpgCreate

> RepositoryGpgKeyResponse ReposGpgCreate(ctx, owner, identifier).Data(data).Execute()

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
    data := *openapiclient.NewRepositoryGpgKeyCreate("GpgPrivateKey_example") // RepositoryGpgKeyCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGpgCreate(context.Background(), owner, identifier).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGpgCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGpgCreate`: RepositoryGpgKeyResponse
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


 **data** | [**RepositoryGpgKeyCreate**](RepositoryGpgKeyCreate.md) |  | 

### Return type

[**RepositoryGpgKeyResponse**](RepositoryGpgKeyResponse.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGpgList

> RepositoryGpgKeyResponse ReposGpgList(ctx, owner, identifier).Execute()

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
    // response from `ReposGpgList`: RepositoryGpgKeyResponse
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

[**RepositoryGpgKeyResponse**](RepositoryGpgKeyResponse.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGpgRegenerate

> RepositoryGpgKeyResponse ReposGpgRegenerate(ctx, owner, identifier).Execute()

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
    // response from `ReposGpgRegenerate`: RepositoryGpgKeyResponse
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

[**RepositoryGpgKeyResponse**](RepositoryGpgKeyResponse.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposNamespaceList

> []RepositoryResponse ReposNamespaceList(ctx, owner).Page(page).PageSize(pageSize).Execute()

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
    resp, r, err := apiClient.ReposApi.ReposNamespaceList(context.Background(), owner).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposNamespaceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposNamespaceList`: []RepositoryResponse
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposNamespaceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposNamespaceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 

### Return type

[**[]RepositoryResponse**](RepositoryResponse.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposPartialUpdate

> RepositoryResponse ReposPartialUpdate(ctx, owner, identifier).Data(data).Execute()

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
    data := *openapiclient.NewRepositoryRequestPatch() // RepositoryRequestPatch |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposPartialUpdate(context.Background(), owner, identifier).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposPartialUpdate`: RepositoryResponse
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


 **data** | [**RepositoryRequestPatch**](RepositoryRequestPatch.md) |  | 

### Return type

[**RepositoryResponse**](RepositoryResponse.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposPrivilegesList

> RepositoryPrivilegeInputResponse ReposPrivilegesList(ctx, owner, identifier).Page(page).PageSize(pageSize).Execute()

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
    page := int64(56) // int64 | A page number within the paginated result set. (optional)
    pageSize := int64(56) // int64 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposPrivilegesList(context.Background(), owner, identifier).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposPrivilegesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposPrivilegesList`: RepositoryPrivilegeInputResponse
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


 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 

### Return type

[**RepositoryPrivilegeInputResponse**](RepositoryPrivilegeInputResponse.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposPrivilegesPartialUpdate

> ReposPrivilegesPartialUpdate(ctx, owner, identifier).Data(data).Execute()

Modify privileges for the repository.



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
    data := *openapiclient.NewRepositoryPrivilegeInputRequestPatch() // RepositoryPrivilegeInputRequestPatch |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposPrivilegesPartialUpdate(context.Background(), owner, identifier).Data(data).Execute()
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


 **data** | [**RepositoryPrivilegeInputRequestPatch**](RepositoryPrivilegeInputRequestPatch.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposPrivilegesUpdate

> ReposPrivilegesUpdate(ctx, owner, identifier).Data(data).Execute()

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
    data := *openapiclient.NewRepositoryPrivilegeInputRequest([]openapiclient.RepositoryPrivilegeDict{*openapiclient.NewRepositoryPrivilegeDict("Privilege_example")}) // RepositoryPrivilegeInputRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposPrivilegesUpdate(context.Background(), owner, identifier).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposPrivilegesUpdate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReposPrivilegesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**RepositoryPrivilegeInputRequest**](RepositoryPrivilegeInputRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposRead

> RepositoryResponse ReposRead(ctx, owner, identifier).Execute()

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
    // response from `ReposRead`: RepositoryResponse
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

[**RepositoryResponse**](RepositoryResponse.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposRsaCreate

> RepositoryRsaKeyResponse ReposRsaCreate(ctx, owner, identifier).Data(data).Execute()

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
    data := *openapiclient.NewRepositoryRsaKeyCreate("RsaPrivateKey_example") // RepositoryRsaKeyCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposRsaCreate(context.Background(), owner, identifier).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposRsaCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposRsaCreate`: RepositoryRsaKeyResponse
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


 **data** | [**RepositoryRsaKeyCreate**](RepositoryRsaKeyCreate.md) |  | 

### Return type

[**RepositoryRsaKeyResponse**](RepositoryRsaKeyResponse.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposRsaList

> RepositoryRsaKeyResponse ReposRsaList(ctx, owner, identifier).Execute()

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
    // response from `ReposRsaList`: RepositoryRsaKeyResponse
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

[**RepositoryRsaKeyResponse**](RepositoryRsaKeyResponse.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposRsaRegenerate

> RepositoryRsaKeyResponse ReposRsaRegenerate(ctx, owner, identifier).Execute()

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
    // response from `ReposRsaRegenerate`: RepositoryRsaKeyResponse
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

[**RepositoryRsaKeyResponse**](RepositoryRsaKeyResponse.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUserList

> []RepositoryResponse ReposUserList(ctx).Page(page).PageSize(pageSize).Execute()

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
    resp, r, err := apiClient.ReposApi.ReposUserList(context.Background()).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposUserList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposUserList`: []RepositoryResponse
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposUserList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReposUserListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 

### Return type

[**[]RepositoryResponse**](RepositoryResponse.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

