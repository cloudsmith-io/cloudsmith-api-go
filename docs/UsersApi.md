# \UsersApi

All URIs are relative to *https://api.cloudsmith.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersProfileRead**](UsersApi.md#UsersProfileRead) | **Get** /users/profile/{slug}/ | Provide a brief for the specified user (if any).



## UsersProfileRead

> UserProfile UsersProfileRead(ctx, slug).Execute()

Provide a brief for the specified user (if any).



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
	resp, r, err := apiClient.UsersApi.UsersProfileRead(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersProfileRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersProfileRead`: UserProfile
	fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersProfileRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersProfileReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserProfile**](UserProfile.md)

### Authorization

[apikey](../README.md#apikey), [basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

