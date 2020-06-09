# \BadgesApi

All URIs are relative to *https://api.cloudsmith.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BadgesVersionList**](BadgesApi.md#BadgesVersionList) | **Get** /badges/version/{owner}/{repo}/{package_format}/{package_name}/{package_version}/{package_identifiers}/ | Get latest package version for a package or package group.



## BadgesVersionList

> map[string]interface{} BadgesVersionList(ctx, owner, repo, packageFormat, packageName, packageVersion, packageIdentifiers, optional)

Get latest package version for a package or package group.

Get latest package version for a package or package group.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
**packageFormat** | **string**|  | 
**packageName** | **string**|  | 
**packageVersion** | **string**|  | 
**packageIdentifiers** | **string**|  | 
 **optional** | ***BadgesVersionListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BadgesVersionListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **badgeToken** | **optional.String**| Badge token to authenticate for private packages | 
 **render** | **optional.Bool**| If true, badge will be rendered | 

### Return type

**map[string]interface{}**

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

