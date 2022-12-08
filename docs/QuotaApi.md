# \QuotaApi

All URIs are relative to *https://api.cloudsmith.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QuotaHistoryRead**](QuotaApi.md#QuotaHistoryRead) | **Get** /quota/history/{owner}/ | Quota history for a given namespace.
[**QuotaOssHistoryRead**](QuotaApi.md#QuotaOssHistoryRead) | **Get** /quota/oss/history/{owner}/ | Open-source Quota history for a given namespace.
[**QuotaOssRead**](QuotaApi.md#QuotaOssRead) | **Get** /quota/oss/{owner}/ | Open-source Quota usage for a given namespace.
[**QuotaRead**](QuotaApi.md#QuotaRead) | **Get** /quota/{owner}/ | Quota usage for a given namespace.



## QuotaHistoryRead

> QuotaHistoryResponse QuotaHistoryRead(ctx, owner).Execute()

Quota history for a given namespace.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QuotaApi.QuotaHistoryRead(context.Background(), owner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotaApi.QuotaHistoryRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QuotaHistoryRead`: QuotaHistoryResponse
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

[**QuotaHistoryResponse**](QuotaHistoryResponse.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuotaOssHistoryRead

> QuotaHistoryResponse QuotaOssHistoryRead(ctx, owner).Execute()

Open-source Quota history for a given namespace.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QuotaApi.QuotaOssHistoryRead(context.Background(), owner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotaApi.QuotaOssHistoryRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QuotaOssHistoryRead`: QuotaHistoryResponse
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

[**QuotaHistoryResponse**](QuotaHistoryResponse.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuotaOssRead

> QuotaResponse QuotaOssRead(ctx, owner).Execute()

Open-source Quota usage for a given namespace.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QuotaApi.QuotaOssRead(context.Background(), owner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotaApi.QuotaOssRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QuotaOssRead`: QuotaResponse
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

[**QuotaResponse**](QuotaResponse.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuotaRead

> QuotaResponse QuotaRead(ctx, owner).Execute()

Quota usage for a given namespace.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QuotaApi.QuotaRead(context.Background(), owner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotaApi.QuotaRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QuotaRead`: QuotaResponse
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

[**QuotaResponse**](QuotaResponse.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

