# \EntitlementsApi

All URIs are relative to *https://api.cloudsmith.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EntitlementsCreate**](EntitlementsApi.md#EntitlementsCreate) | **Post** /entitlements/{owner}/{repo}/ | Create a specific entitlement in a repository.
[**EntitlementsDelete**](EntitlementsApi.md#EntitlementsDelete) | **Delete** /entitlements/{owner}/{repo}/{identifier}/ | Delete a specific entitlement in a repository.
[**EntitlementsDisable**](EntitlementsApi.md#EntitlementsDisable) | **Post** /entitlements/{owner}/{repo}/{identifier}/disable/ | Disable an entitlement token in a repository.
[**EntitlementsEnable**](EntitlementsApi.md#EntitlementsEnable) | **Post** /entitlements/{owner}/{repo}/{identifier}/enable/ | Enable an entitlement token in a repository.
[**EntitlementsList**](EntitlementsApi.md#EntitlementsList) | **Get** /entitlements/{owner}/{repo}/ | Get a list of all entitlements in a repository.
[**EntitlementsPartialUpdate**](EntitlementsApi.md#EntitlementsPartialUpdate) | **Patch** /entitlements/{owner}/{repo}/{identifier}/ | Update a specific entitlement in a repository.
[**EntitlementsRead**](EntitlementsApi.md#EntitlementsRead) | **Get** /entitlements/{owner}/{repo}/{identifier}/ | Get a specific entitlement in a repository.
[**EntitlementsRefresh**](EntitlementsApi.md#EntitlementsRefresh) | **Post** /entitlements/{owner}/{repo}/{identifier}/refresh/ | Refresh an entitlement token in a repository.
[**EntitlementsReset**](EntitlementsApi.md#EntitlementsReset) | **Post** /entitlements/{owner}/{repo}/{identifier}/reset/ | Reset the statistics for an entitlement token in a repository.
[**EntitlementsSync**](EntitlementsApi.md#EntitlementsSync) | **Post** /entitlements/{owner}/{repo}/sync/ | Synchronise tokens from a source repository.



## EntitlementsCreate

> RepositoryToken EntitlementsCreate(ctx, owner, repo, optional)

Create a specific entitlement in a repository.

Create a specific entitlement in a repository.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***EntitlementsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EntitlementsCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **showTokens** | **optional.Bool**| Show entitlement token strings in results | 
 **data** | [**optional.Interface of EntitlementsCreate**](EntitlementsCreate.md)|  | 

### Return type

[**RepositoryToken**](RepositoryToken.md)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntitlementsDelete

> EntitlementsDelete(ctx, owner, repo, identifier)

Delete a specific entitlement in a repository.

Delete a specific entitlement in a repository.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
**identifier** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntitlementsDisable

> EntitlementsDisable(ctx, owner, repo, identifier)

Disable an entitlement token in a repository.

Disable an entitlement token in a repository.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
**identifier** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntitlementsEnable

> EntitlementsEnable(ctx, owner, repo, identifier)

Enable an entitlement token in a repository.

Enable an entitlement token in a repository.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
**identifier** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntitlementsList

> []RepositoryToken EntitlementsList(ctx, owner, repo, optional)

Get a list of all entitlements in a repository.

Get a list of all entitlements in a repository.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***EntitlementsListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EntitlementsListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| A page number within the paginated result set. | 
 **pageSize** | **optional.Int32**| Number of results to return per page. | 
 **showTokens** | **optional.Bool**| Show entitlement token strings in results | 

### Return type

[**[]RepositoryToken**](RepositoryToken.md)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntitlementsPartialUpdate

> RepositoryToken EntitlementsPartialUpdate(ctx, owner, repo, identifier, optional)

Update a specific entitlement in a repository.

Update a specific entitlement in a repository.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
**identifier** | **string**|  | 
 **optional** | ***EntitlementsPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EntitlementsPartialUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **showTokens** | **optional.Bool**| Show entitlement token strings in results | 
 **data** | [**optional.Interface of EntitlementsPartialUpdate**](EntitlementsPartialUpdate.md)|  | 

### Return type

[**RepositoryToken**](RepositoryToken.md)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntitlementsRead

> RepositoryToken EntitlementsRead(ctx, owner, repo, identifier, optional)

Get a specific entitlement in a repository.

Get a specific entitlement in a repository.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
**identifier** | **string**|  | 
 **optional** | ***EntitlementsReadOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EntitlementsReadOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **showTokens** | **optional.Bool**| Show entitlement token strings in results | 

### Return type

[**RepositoryToken**](RepositoryToken.md)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntitlementsRefresh

> RepositoryTokenRefresh EntitlementsRefresh(ctx, owner, repo, identifier, optional)

Refresh an entitlement token in a repository.

Refresh an entitlement token in a repository.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
**identifier** | **string**|  | 
 **optional** | ***EntitlementsRefreshOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EntitlementsRefreshOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **showTokens** | **optional.Bool**| Show entitlement token strings in results | 
 **data** | [**optional.Interface of EntitlementsRefresh**](EntitlementsRefresh.md)|  | 

### Return type

[**RepositoryTokenRefresh**](RepositoryTokenRefresh.md)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntitlementsReset

> EntitlementsReset(ctx, owner, repo, identifier)

Reset the statistics for an entitlement token in a repository.

Reset the statistics for an entitlement token in a repository.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
**identifier** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntitlementsSync

> RepositoryTokenSync EntitlementsSync(ctx, owner, repo, optional)

Synchronise tokens from a source repository.

Synchronise tokens from a source repository.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***EntitlementsSyncOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EntitlementsSyncOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **showTokens** | **optional.Bool**| Show entitlement token strings in results | 
 **data** | [**optional.Interface of EntitlementsSync**](EntitlementsSync.md)|  | 

### Return type

[**RepositoryTokenSync**](RepositoryTokenSync.md)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

