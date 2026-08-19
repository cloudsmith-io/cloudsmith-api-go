# \OpenidApi

All URIs are relative to *https://api.cloudsmith.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpenidCreate**](OpenidApi.md#OpenidCreate) | **Post** /openid/{owner}/ | Get a JWT token for a configured service account belonging to the requesting org.



## OpenidCreate

> Oidc1 OpenidCreate(ctx, owner).Data(data).Execute()

Get a JWT token for a configured service account belonging to the requesting org.



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
	data := *openapiclient.NewOidcRequest("OidcToken_example", "ServiceSlug_example") // OidcRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenidApi.OpenidCreate(context.Background(), owner).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenidApi.OpenidCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpenidCreate`: Oidc1
	fmt.Fprintf(os.Stdout, "Response from `OpenidApi.OpenidCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpenidCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**OidcRequest**](OidcRequest.md) |  | 

### Return type

[**Oidc1**](Oidc1.md)

### Authorization

[apikey](../README.md#apikey), [basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

