# \UsersApi

All URIs are relative to *https://api.cloudsmith.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersProfileRead**](UsersApi.md#UsersProfileRead) | **Get** /users/profile/{slug}/ | Provide a brief for the specified user (if any).



## UsersProfileRead

> UserProfile UsersProfileRead(ctx, slug)

Provide a brief for the specified user (if any).

Provide a brief for the specified user (if any).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string**|  | 

### Return type

[**UserProfile**](UserProfile.md)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

