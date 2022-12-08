# \StatusApi

All URIs are relative to *https://api.cloudsmith.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StatusCheckBasic**](StatusApi.md#StatusCheckBasic) | **Get** /status/check/basic/ | Endpoint to check basic API connectivity.



## StatusCheckBasic

> StatusBasicResponse StatusCheckBasic(ctx).Execute()

Endpoint to check basic API connectivity.



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
    resp, r, err := apiClient.StatusApi.StatusCheckBasic(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatusApi.StatusCheckBasic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StatusCheckBasic`: StatusBasicResponse
    fmt.Fprintf(os.Stdout, "Response from `StatusApi.StatusCheckBasic`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStatusCheckBasicRequest struct via the builder pattern


### Return type

[**StatusBasicResponse**](StatusBasicResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

