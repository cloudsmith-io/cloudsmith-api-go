# \FormatsApi

All URIs are relative to *https://api.cloudsmith.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FormatsList**](FormatsApi.md#FormatsList) | **Get** /formats/ | Get a list of all supported package formats.
[**FormatsRead**](FormatsApi.md#FormatsRead) | **Get** /formats/{slug}/ | Get a specific supported package format.



## FormatsList

> []FormatResponse FormatsList(ctx).Execute()

Get a list of all supported package formats.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FormatsApi.FormatsList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FormatsApi.FormatsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FormatsList`: []FormatResponse
    fmt.Fprintf(os.Stdout, "Response from `FormatsApi.FormatsList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFormatsListRequest struct via the builder pattern


### Return type

[**[]FormatResponse**](FormatResponse.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FormatsRead

> FormatResponse FormatsRead(ctx, slug).Execute()

Get a specific supported package format.



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
    slug := "slug_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FormatsApi.FormatsRead(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FormatsApi.FormatsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FormatsRead`: FormatResponse
    fmt.Fprintf(os.Stdout, "Response from `FormatsApi.FormatsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFormatsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FormatResponse**](FormatResponse.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

