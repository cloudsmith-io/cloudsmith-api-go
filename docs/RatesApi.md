# \RatesApi

All URIs are relative to *https://api.cloudsmith.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RatesLimitsList**](RatesApi.md#RatesLimitsList) | **Get** /rates/limits/ | Endpoint to check rate limits for current user.



## RatesLimitsList

> ResourcesRateCheck RatesLimitsList(ctx).Execute()

Endpoint to check rate limits for current user.



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
    resp, r, err := apiClient.RatesApi.RatesLimitsList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RatesApi.RatesLimitsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RatesLimitsList`: ResourcesRateCheck
    fmt.Fprintf(os.Stdout, "Response from `RatesApi.RatesLimitsList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRatesLimitsListRequest struct via the builder pattern


### Return type

[**ResourcesRateCheck**](ResourcesRateCheck.md)

### Authorization

[apikey](../README.md#apikey), [basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

