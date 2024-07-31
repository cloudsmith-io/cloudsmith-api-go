# \QuotaApi

All URIs are relative to *https://api.cloudsmith.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QuotaHistoryRead**](QuotaApi.md#QuotaHistoryRead) | **Get** /quota/history/{owner}/ | Quota history for a given namespace.
[**QuotaOssHistoryRead**](QuotaApi.md#QuotaOssHistoryRead) | **Get** /quota/oss/history/{owner}/ | Open-source Quota history for a given namespace.
[**QuotaOssRead**](QuotaApi.md#QuotaOssRead) | **Get** /quota/oss/{owner}/ | Open-source Quota usage for a given namespace.
[**QuotaRead**](QuotaApi.md#QuotaRead) | **Get** /quota/{owner}/ | Quota usage for a given namespace.



## QuotaHistoryRead

> QuotaHistory QuotaHistoryRead(ctx, owner).Execute()

Quota history for a given namespace.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QuotaApi.QuotaHistoryRead(context.Background(), owner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotaApi.QuotaHistoryRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QuotaHistoryRead`: QuotaHistory
    fmt.Fprintf(os.Stdout, "Response from `QuotaApi.QuotaHistoryRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuotaHistoryReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QuotaHistory**](QuotaHistory.md)

### Authorization

[apikey](../README.md#apikey), [basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuotaOssHistoryRead

> QuotaHistory QuotaOssHistoryRead(ctx, owner).Execute()

Open-source Quota history for a given namespace.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QuotaApi.QuotaOssHistoryRead(context.Background(), owner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotaApi.QuotaOssHistoryRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QuotaOssHistoryRead`: QuotaHistory
    fmt.Fprintf(os.Stdout, "Response from `QuotaApi.QuotaOssHistoryRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuotaOssHistoryReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QuotaHistory**](QuotaHistory.md)

### Authorization

[apikey](../README.md#apikey), [basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuotaOssRead

> Quota QuotaOssRead(ctx, owner).Execute()

Open-source Quota usage for a given namespace.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QuotaApi.QuotaOssRead(context.Background(), owner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotaApi.QuotaOssRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QuotaOssRead`: Quota
    fmt.Fprintf(os.Stdout, "Response from `QuotaApi.QuotaOssRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuotaOssReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Quota**](Quota.md)

### Authorization

[apikey](../README.md#apikey), [basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuotaRead

> Quota QuotaRead(ctx, owner).Execute()

Quota usage for a given namespace.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QuotaApi.QuotaRead(context.Background(), owner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotaApi.QuotaRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QuotaRead`: Quota
    fmt.Fprintf(os.Stdout, "Response from `QuotaApi.QuotaRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuotaReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Quota**](Quota.md)

### Authorization

[apikey](../README.md#apikey), [basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

