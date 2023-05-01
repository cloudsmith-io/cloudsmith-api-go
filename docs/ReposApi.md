# \ReposApi

All URIs are relative to *https://api.cloudsmith.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReposCreate**](ReposApi.md#ReposCreate) | **Post** /repos/{owner}/ | Create a new repository in a given namespace.
[**ReposDelete**](ReposApi.md#ReposDelete) | **Delete** /repos/{owner}/{identifier}/ | Delete a repository in a given namespace.
[**ReposGeoipDisable**](ReposApi.md#ReposGeoipDisable) | **Post** /repos/{owner}/{identifier}/geoip/disable/ | Disable GeoIP for this repository.
[**ReposGeoipEnable**](ReposApi.md#ReposGeoipEnable) | **Post** /repos/{owner}/{identifier}/geoip/enable/ | Enable GeoIP for this repository.
[**ReposGeoipPartialUpdate**](ReposApi.md#ReposGeoipPartialUpdate) | **Patch** /repos/{owner}/{identifier}/geoip | Partially update repository geoip rules.
[**ReposGeoipRead**](ReposApi.md#ReposGeoipRead) | **Get** /repos/{owner}/{identifier}/geoip | List all repository geoip rules.
[**ReposGeoipTest**](ReposApi.md#ReposGeoipTest) | **Post** /repos/{owner}/{identifier}/geoip/test/ | Test a list of IP addresses against the repository&#39;s current GeoIP rules.
[**ReposGeoipUpdate**](ReposApi.md#ReposGeoipUpdate) | **Put** /repos/{owner}/{identifier}/geoip | Replace repository geoip rules.
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
    data := *openapiclient.NewRepositoryCreateRequest("Name_example") // RepositoryCreateRequest |  (optional)

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

 **data** | [**RepositoryCreateRequest**](RepositoryCreateRequest.md) |  | 

### Return type

[**RepositoryCreate**](RepositoryCreate.md)

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


## ReposGeoipDisable

> ReposGeoipDisable(ctx, owner, identifier).Data(data).Execute()

Disable GeoIP for this repository.



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
    data := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGeoipDisable(context.Background(), owner, identifier).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGeoipDisable``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReposGeoipDisableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | **map[string]interface{}** |  | 

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


## ReposGeoipEnable

> ReposGeoipEnable(ctx, owner, identifier).Data(data).Execute()

Enable GeoIP for this repository.



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
    data := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGeoipEnable(context.Background(), owner, identifier).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGeoipEnable``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReposGeoipEnableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | **map[string]interface{}** |  | 

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


## ReposGeoipPartialUpdate

> RepositoryGeoIpRules ReposGeoipPartialUpdate(ctx, owner, identifier).Data(data).Execute()

Partially update repository geoip rules.



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
    data := *openapiclient.NewRepositoryGeoIpRulesRequestPatch() // RepositoryGeoIpRulesRequestPatch |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGeoipPartialUpdate(context.Background(), owner, identifier).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGeoipPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGeoipPartialUpdate`: RepositoryGeoIpRules
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGeoipPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGeoipPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**RepositoryGeoIpRulesRequestPatch**](RepositoryGeoIpRulesRequestPatch.md) |  | 

### Return type

[**RepositoryGeoIpRules**](RepositoryGeoIpRules.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGeoipRead

> RepositoryGeoIpRules ReposGeoipRead(ctx, owner, identifier).Execute()

List all repository geoip rules.



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
    resp, r, err := apiClient.ReposApi.ReposGeoipRead(context.Background(), owner, identifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGeoipRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGeoipRead`: RepositoryGeoIpRules
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGeoipRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGeoipReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RepositoryGeoIpRules**](RepositoryGeoIpRules.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGeoipTest

> RepositoryGeoIpTestAddressResponse ReposGeoipTest(ctx, owner, identifier).Data(data).Execute()

Test a list of IP addresses against the repository's current GeoIP rules.



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
    data := *openapiclient.NewRepositoryGeoIpTestAddress([]string{"Addresses_example"}) // RepositoryGeoIpTestAddress |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGeoipTest(context.Background(), owner, identifier).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGeoipTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGeoipTest`: RepositoryGeoIpTestAddressResponse
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGeoipTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGeoipTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**RepositoryGeoIpTestAddress**](RepositoryGeoIpTestAddress.md) |  | 

### Return type

[**RepositoryGeoIpTestAddressResponse**](RepositoryGeoIpTestAddressResponse.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposGeoipUpdate

> RepositoryGeoIpRules ReposGeoipUpdate(ctx, owner, identifier).Data(data).Execute()

Replace repository geoip rules.



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
    data := *openapiclient.NewRepositoryGeoIpRulesRequest(*openapiclient.NewRepositoryGeoIpCidr([]string{"Allow_example"}, []string{"Deny_example"}), *openapiclient.NewRepositoryGeoIpCountryCode([]string{"Allow_example"}, []string{"Deny_example"})) // RepositoryGeoIpRulesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReposApi.ReposGeoipUpdate(context.Background(), owner, identifier).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReposApi.ReposGeoipUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReposGeoipUpdate`: RepositoryGeoIpRules
    fmt.Fprintf(os.Stdout, "Response from `ReposApi.ReposGeoipUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReposGeoipUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**RepositoryGeoIpRulesRequest**](RepositoryGeoIpRulesRequest.md) |  | 

### Return type

[**RepositoryGeoIpRules**](RepositoryGeoIpRules.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

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
    data := *openapiclient.NewRepositoryGpgKeyCreate("GpgPrivateKey_example") // RepositoryGpgKeyCreate |  (optional)

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


 **data** | [**RepositoryGpgKeyCreate**](RepositoryGpgKeyCreate.md) |  | 

### Return type

[**RepositoryGpgKey**](RepositoryGpgKey.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

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
- **Accept**: application/json

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposNamespaceList

> []Repository ReposNamespaceList(ctx, owner).Page(page).PageSize(pageSize).Execute()

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
    // response from `ReposNamespaceList`: []Repository
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

[**[]Repository**](Repository.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
    data := *openapiclient.NewRepositoryRequestPatch() // RepositoryRequestPatch |  (optional)

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


 **data** | [**RepositoryRequestPatch**](RepositoryRequestPatch.md) |  | 

### Return type

[**Repository**](Repository.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposPrivilegesList

> RepositoryPrivilegeInput ReposPrivilegesList(ctx, owner, identifier).Page(page).PageSize(pageSize).Execute()

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
    // response from `ReposPrivilegesList`: RepositoryPrivilegeInput
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

[**RepositoryPrivilegeInput**](RepositoryPrivilegeInput.md)

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
- **Accept**: application/json

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
    data := *openapiclient.NewRepositoryRsaKeyCreate("RsaPrivateKey_example") // RepositoryRsaKeyCreate |  (optional)

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


 **data** | [**RepositoryRsaKeyCreate**](RepositoryRsaKeyCreate.md) |  | 

### Return type

[**RepositoryRsaKey**](RepositoryRsaKey.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

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
- **Accept**: application/json

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReposUserList

> []Repository ReposUserList(ctx).Page(page).PageSize(pageSize).Execute()

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
    // response from `ReposUserList`: []Repository
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

[**[]Repository**](Repository.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

