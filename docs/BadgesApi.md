# \BadgesApi

All URIs are relative to *https://api.cloudsmith.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BadgesVersionList**](BadgesApi.md#BadgesVersionList) | **Get** /badges/version/{owner}/{repo}/{package_format}/{package_name}/{package_version}/{package_identifiers}/ | Get latest package version for a package or package group.



## BadgesVersionList

> map[string]interface{} BadgesVersionList(ctx, owner, repo, packageFormat, packageName, packageVersion, packageIdentifiers).BadgeToken(badgeToken).CacheSeconds(cacheSeconds).Color(color).Label(label).LabelColor(labelColor).LogoColor(logoColor).LogoWidth(logoWidth).Render(render).Shields(shields).ShowLatest(showLatest).Style(style).Execute()

Get latest package version for a package or package group.



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
    repo := "repo_example" // string | 
    packageFormat := "packageFormat_example" // string | 
    packageName := "packageName_example" // string | 
    packageVersion := "packageVersion_example" // string | 
    packageIdentifiers := "packageIdentifiers_example" // string | 
    badgeToken := "badgeToken_example" // string | Badge token to authenticate for private packages (optional)
    cacheSeconds := "cacheSeconds_example" // string | Override the shields.io badge cacheSeconds value. (optional)
    color := "color_example" // string | Override the shields.io badge color value. (optional)
    label := "label_example" // string | Override the shields.io badge label value. (optional)
    labelColor := "labelColor_example" // string | Override the shields.io badge labelColor value. (optional)
    logoColor := "logoColor_example" // string | Override the shields.io badge logoColor value. (optional)
    logoWidth := "logoWidth_example" // string | Override the shields.io badge logoWidth value. (optional)
    render := true // bool | If true, badge will be rendered (optional)
    shields := true // bool | If true, a shields response will be generated (optional)
    showLatest := true // bool | If true, for latest version badges a '(latest)' suffix is added (optional)
    style := "style_example" // string | Override the shields.io badge style value. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BadgesApi.BadgesVersionList(context.Background(), owner, repo, packageFormat, packageName, packageVersion, packageIdentifiers).BadgeToken(badgeToken).CacheSeconds(cacheSeconds).Color(color).Label(label).LabelColor(labelColor).LogoColor(logoColor).LogoWidth(logoWidth).Render(render).Shields(shields).ShowLatest(showLatest).Style(style).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BadgesApi.BadgesVersionList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BadgesVersionList`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BadgesApi.BadgesVersionList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** |  | 
**repo** | **string** |  | 
**packageFormat** | **string** |  | 
**packageName** | **string** |  | 
**packageVersion** | **string** |  | 
**packageIdentifiers** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBadgesVersionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **badgeToken** | **string** | Badge token to authenticate for private packages | 
 **cacheSeconds** | **string** | Override the shields.io badge cacheSeconds value. | 
 **color** | **string** | Override the shields.io badge color value. | 
 **label** | **string** | Override the shields.io badge label value. | 
 **labelColor** | **string** | Override the shields.io badge labelColor value. | 
 **logoColor** | **string** | Override the shields.io badge logoColor value. | 
 **logoWidth** | **string** | Override the shields.io badge logoWidth value. | 
 **render** | **bool** | If true, badge will be rendered | 
 **shields** | **bool** | If true, a shields response will be generated | 
 **showLatest** | **bool** | If true, for latest version badges a &#39;(latest)&#39; suffix is added | 
 **style** | **string** | Override the shields.io badge style value. | 

### Return type

**map[string]interface{}**

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

