# \DistrosApi

All URIs are relative to *https://api.cloudsmith.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DistrosList**](DistrosApi.md#DistrosList) | **Get** /distros/ | Get a list of all supported distributions.
[**DistrosRead**](DistrosApi.md#DistrosRead) | **Get** /distros/{slug}/ | View for viewing/listing distributions.



## DistrosList

> []DistributionFull DistrosList(ctx).Execute()

Get a list of all supported distributions.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DistrosApi.DistrosList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistrosApi.DistrosList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DistrosList`: []DistributionFull
    fmt.Fprintf(os.Stdout, "Response from `DistrosApi.DistrosList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDistrosListRequest struct via the builder pattern


### Return type

[**[]DistributionFull**](DistributionFull.md)

### Authorization

[apikey](../README.md#apikey), [basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DistrosRead

> DistributionFull DistrosRead(ctx, slug).Execute()

View for viewing/listing distributions.



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
    slug := "slug_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DistrosApi.DistrosRead(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistrosApi.DistrosRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DistrosRead`: DistributionFull
    fmt.Fprintf(os.Stdout, "Response from `DistrosApi.DistrosRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistrosReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DistributionFull**](DistributionFull.md)

### Authorization

[apikey](../README.md#apikey), [basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

