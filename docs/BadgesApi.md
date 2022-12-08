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
    cacheSeconds := "cacheSeconds_example" // string | Override the shields.io badge cacheSeconds value. (optional) (default to "300")
    color := "color_example" // string | Override the shields.io badge color value. (optional) (default to "12577E")
    label := "label_example" // string | Override the shields.io badge label value. (optional) (default to "cloudsmith")
    labelColor := "labelColor_example" // string | Override the shields.io badge labelColor value. (optional) (default to "021F2F")
    logoColor := "logoColor_example" // string | Override the shields.io badge logoColor value. (optional) (default to "45B6EE")
    logoWidth := "logoWidth_example" // string | Override the shields.io badge logoWidth value. (optional) (default to "10")
    render := true // bool | If true, badge will be rendered (optional) (default to false)
    shields := true // bool | If true, a shields response will be generated (optional) (default to false)
    showLatest := true // bool | If true, for latest version badges a '(latest)' suffix is added (optional) (default to false)
    style := "style_example" // string | Override the shields.io badge style value. (optional) (default to "flat-square")

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
 **cacheSeconds** | **string** | Override the shields.io badge cacheSeconds value. | [default to &quot;300&quot;]
 **color** | **string** | Override the shields.io badge color value. | [default to &quot;12577E&quot;]
 **label** | **string** | Override the shields.io badge label value. | [default to &quot;cloudsmith&quot;]
 **labelColor** | **string** | Override the shields.io badge labelColor value. | [default to &quot;021F2F&quot;]
 **logoColor** | **string** | Override the shields.io badge logoColor value. | [default to &quot;45B6EE&quot;]
 **logoWidth** | **string** | Override the shields.io badge logoWidth value. | [default to &quot;10&quot;]
 **render** | **bool** | If true, badge will be rendered | [default to false]
 **shields** | **bool** | If true, a shields response will be generated | [default to false]
 **showLatest** | **bool** | If true, for latest version badges a &#39;(latest)&#39; suffix is added | [default to false]
 **style** | **string** | Override the shields.io badge style value. | [default to &quot;flat-square&quot;]

### Return type

**map[string]interface{}**

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

