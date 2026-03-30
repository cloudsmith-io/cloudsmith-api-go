# \BulkActionApi

All URIs are relative to *https://api.cloudsmith.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkAction**](BulkActionApi.md#BulkAction) | **Post** /bulk-action/{owner}/ | 



## BulkAction

> PackageBulkActionResponse BulkAction(ctx, owner).Data(data).Execute()





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
	data := *openapiclient.NewPackageBulkAction("Action_example", []string{"Identifiers_example"}) // PackageBulkAction |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BulkActionApi.BulkAction(context.Background(), owner).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkActionApi.BulkAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkAction`: PackageBulkActionResponse
	fmt.Fprintf(os.Stdout, "Response from `BulkActionApi.BulkAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**PackageBulkAction**](PackageBulkAction.md) |  | 

### Return type

[**PackageBulkActionResponse**](PackageBulkActionResponse.md)

### Authorization

[apikey](../README.md#apikey), [basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

