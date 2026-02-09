# \BroadcastsApi

All URIs are relative to *https://api.cloudsmith.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BroadcastsCreateBroadcastToken**](BroadcastsApi.md#BroadcastsCreateBroadcastToken) | **Post** /broadcasts/{org}/broadcast-token/ | Create a broadcast token.



## BroadcastsCreateBroadcastToken

> BroadcastToken BroadcastsCreateBroadcastToken(ctx, org).Data(data).Execute()

Create a broadcast token.



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
	org := "org_example" // string | 
	data := *openapiclient.NewBroadcastTokenInput("EntitlementToken_example") // BroadcastTokenInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BroadcastsApi.BroadcastsCreateBroadcastToken(context.Background(), org).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BroadcastsApi.BroadcastsCreateBroadcastToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BroadcastsCreateBroadcastToken`: BroadcastToken
	fmt.Fprintf(os.Stdout, "Response from `BroadcastsApi.BroadcastsCreateBroadcastToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBroadcastsCreateBroadcastTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**BroadcastTokenInput**](BroadcastTokenInput.md) |  | 

### Return type

[**BroadcastToken**](BroadcastToken.md)

### Authorization

[apikey](../README.md#apikey), [basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

