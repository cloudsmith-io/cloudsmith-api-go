# \EntitlementsApi

All URIs are relative to *https://api.cloudsmith.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EntitlementsCreate**](EntitlementsApi.md#EntitlementsCreate) | **Post** /entitlements/{owner}/{repo}/ | Create a specific entitlement in a repository.
[**EntitlementsDelete**](EntitlementsApi.md#EntitlementsDelete) | **Delete** /entitlements/{owner}/{repo}/{identifier}/ | Delete a specific entitlement in a repository.
[**EntitlementsDisable**](EntitlementsApi.md#EntitlementsDisable) | **Post** /entitlements/{owner}/{repo}/{identifier}/disable/ | Disable an entitlement token in a repository.
[**EntitlementsEnable**](EntitlementsApi.md#EntitlementsEnable) | **Post** /entitlements/{owner}/{repo}/{identifier}/enable/ | Enable an entitlement token in a repository.
[**EntitlementsList**](EntitlementsApi.md#EntitlementsList) | **Get** /entitlements/{owner}/{repo}/ | Get a list of all entitlements in a repository.
[**EntitlementsPartialUpdate**](EntitlementsApi.md#EntitlementsPartialUpdate) | **Patch** /entitlements/{owner}/{repo}/{identifier}/ | Update a specific entitlement in a repository.
[**EntitlementsRead**](EntitlementsApi.md#EntitlementsRead) | **Get** /entitlements/{owner}/{repo}/{identifier}/ | Get a specific entitlement in a repository.
[**EntitlementsRefresh**](EntitlementsApi.md#EntitlementsRefresh) | **Post** /entitlements/{owner}/{repo}/{identifier}/refresh/ | Refresh an entitlement token in a repository.
[**EntitlementsReset**](EntitlementsApi.md#EntitlementsReset) | **Post** /entitlements/{owner}/{repo}/{identifier}/reset/ | Reset the statistics for an entitlement token in a repository.
[**EntitlementsSync**](EntitlementsApi.md#EntitlementsSync) | **Post** /entitlements/{owner}/{repo}/sync/ | Synchronise tokens from a source repository.



## EntitlementsCreate

> RepositoryToken EntitlementsCreate(ctx, owner, repo).ShowTokens(showTokens).Data(data).Execute()

Create a specific entitlement in a repository.



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
    showTokens := true // bool | Show entitlement token strings in results (optional) (default to false)
    data := *openapiclient.NewRepositoryTokenRequest("Name_example") // RepositoryTokenRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntitlementsApi.EntitlementsCreate(context.Background(), owner, repo).ShowTokens(showTokens).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.EntitlementsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EntitlementsCreate`: RepositoryToken
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsApi.EntitlementsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEntitlementsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **showTokens** | **bool** | Show entitlement token strings in results | [default to false]
 **data** | [**RepositoryTokenRequest**](RepositoryTokenRequest.md) |  | 

### Return type

[**RepositoryToken**](RepositoryToken.md)

### Authorization

[apikey](../README.md#apikey), [basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntitlementsDelete

> EntitlementsDelete(ctx, owner, repo, identifier).Execute()

Delete a specific entitlement in a repository.



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
    identifier := "identifier_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EntitlementsApi.EntitlementsDelete(context.Background(), owner, repo, identifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.EntitlementsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEntitlementsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntitlementsDisable

> EntitlementsDisable(ctx, owner, repo, identifier).Execute()

Disable an entitlement token in a repository.



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
    identifier := "identifier_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EntitlementsApi.EntitlementsDisable(context.Background(), owner, repo, identifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.EntitlementsDisable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEntitlementsDisableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntitlementsEnable

> EntitlementsEnable(ctx, owner, repo, identifier).Execute()

Enable an entitlement token in a repository.



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
    identifier := "identifier_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EntitlementsApi.EntitlementsEnable(context.Background(), owner, repo, identifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.EntitlementsEnable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEntitlementsEnableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntitlementsList

> []RepositoryToken EntitlementsList(ctx, owner, repo).Page(page).PageSize(pageSize).ShowTokens(showTokens).Query(query).Active(active).Execute()

Get a list of all entitlements in a repository.



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
    showTokens := true // bool | Show entitlement token strings in results (optional) (default to false)
    query := "query_example" // string | A search term for querying names of entitlements. (optional)
    active := true // bool | If true, only include active tokens (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntitlementsApi.EntitlementsList(context.Background(), owner, repo).Page(page).PageSize(pageSize).ShowTokens(showTokens).Query(query).Active(active).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.EntitlementsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EntitlementsList`: []RepositoryToken
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsApi.EntitlementsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEntitlementsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int64** | A page number within the paginated result set. | 
 **pageSize** | **int64** | Number of results to return per page. | 
 **showTokens** | **bool** | Show entitlement token strings in results | [default to false]
 **query** | **string** | A search term for querying names of entitlements. | 
 **active** | **bool** | If true, only include active tokens | [default to false]

### Return type

[**[]RepositoryToken**](RepositoryToken.md)

### Authorization

[apikey](../README.md#apikey), [basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntitlementsPartialUpdate

> RepositoryToken EntitlementsPartialUpdate(ctx, owner, repo, identifier).ShowTokens(showTokens).Data(data).Execute()

Update a specific entitlement in a repository.



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
    identifier := "identifier_example" // string | 
    showTokens := true // bool | Show entitlement token strings in results (optional) (default to false)
    data := *openapiclient.NewRepositoryTokenRequestPatch() // RepositoryTokenRequestPatch |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntitlementsApi.EntitlementsPartialUpdate(context.Background(), owner, repo, identifier).ShowTokens(showTokens).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.EntitlementsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EntitlementsPartialUpdate`: RepositoryToken
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsApi.EntitlementsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEntitlementsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **showTokens** | **bool** | Show entitlement token strings in results | [default to false]
 **data** | [**RepositoryTokenRequestPatch**](RepositoryTokenRequestPatch.md) |  | 

### Return type

[**RepositoryToken**](RepositoryToken.md)

### Authorization

[apikey](../README.md#apikey), [basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntitlementsRead

> RepositoryToken EntitlementsRead(ctx, owner, repo, identifier).Fuzzy(fuzzy).ShowTokens(showTokens).Execute()

Get a specific entitlement in a repository.



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
    identifier := "identifier_example" // string | 
    fuzzy := true // bool | If true, entitlement identifiers including name will be fuzzy matched. (optional) (default to false)
    showTokens := true // bool | Show entitlement token strings in results (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntitlementsApi.EntitlementsRead(context.Background(), owner, repo, identifier).Fuzzy(fuzzy).ShowTokens(showTokens).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.EntitlementsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EntitlementsRead`: RepositoryToken
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsApi.EntitlementsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEntitlementsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fuzzy** | **bool** | If true, entitlement identifiers including name will be fuzzy matched. | [default to false]
 **showTokens** | **bool** | Show entitlement token strings in results | [default to false]

### Return type

[**RepositoryToken**](RepositoryToken.md)

### Authorization

[apikey](../README.md#apikey), [basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntitlementsRefresh

> RepositoryTokenRefresh EntitlementsRefresh(ctx, owner, repo, identifier).ShowTokens(showTokens).Data(data).Execute()

Refresh an entitlement token in a repository.



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
    identifier := "identifier_example" // string | 
    showTokens := true // bool | Show entitlement token strings in results (optional) (default to false)
    data := *openapiclient.NewRepositoryTokenRefreshRequest() // RepositoryTokenRefreshRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntitlementsApi.EntitlementsRefresh(context.Background(), owner, repo, identifier).ShowTokens(showTokens).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.EntitlementsRefresh``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EntitlementsRefresh`: RepositoryTokenRefresh
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsApi.EntitlementsRefresh`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEntitlementsRefreshRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **showTokens** | **bool** | Show entitlement token strings in results | [default to false]
 **data** | [**RepositoryTokenRefreshRequest**](RepositoryTokenRefreshRequest.md) |  | 

### Return type

[**RepositoryTokenRefresh**](RepositoryTokenRefresh.md)

### Authorization

[apikey](../README.md#apikey), [basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntitlementsReset

> EntitlementsReset(ctx, owner, repo, identifier).ShowTokens(showTokens).Execute()

Reset the statistics for an entitlement token in a repository.



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
    identifier := "identifier_example" // string | 
    showTokens := true // bool | Show entitlement token strings in results (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EntitlementsApi.EntitlementsReset(context.Background(), owner, repo, identifier).ShowTokens(showTokens).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.EntitlementsReset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEntitlementsResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **showTokens** | **bool** | Show entitlement token strings in results | [default to false]

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntitlementsSync

> RepositoryTokenSync EntitlementsSync(ctx, owner, repo).ShowTokens(showTokens).Data(data).Execute()

Synchronise tokens from a source repository.



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
    showTokens := true // bool | Show entitlement token strings in results (optional) (default to false)
    data := *openapiclient.NewRepositoryTokenSyncRequest("Source_example") // RepositoryTokenSyncRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntitlementsApi.EntitlementsSync(context.Background(), owner, repo).ShowTokens(showTokens).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.EntitlementsSync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EntitlementsSync`: RepositoryTokenSync
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsApi.EntitlementsSync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEntitlementsSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **showTokens** | **bool** | Show entitlement token strings in results | [default to false]
 **data** | [**RepositoryTokenSyncRequest**](RepositoryTokenSyncRequest.md) |  | 

### Return type

[**RepositoryTokenSync**](RepositoryTokenSync.md)

### Authorization

[apikey](../README.md#apikey), [basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

