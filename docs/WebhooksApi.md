# \WebhooksApi

All URIs are relative to *https://api.cloudsmith.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WebhooksCreate**](WebhooksApi.md#WebhooksCreate) | **Post** /webhooks/{owner}/{repo}/ | Create a specific webhook in a repository.
[**WebhooksDelete**](WebhooksApi.md#WebhooksDelete) | **Delete** /webhooks/{owner}/{repo}/{identifier}/ | Delete a specific webhook in a repository.
[**WebhooksList**](WebhooksApi.md#WebhooksList) | **Get** /webhooks/{owner}/{repo}/ | Get a list of all webhooks in a repository.
[**WebhooksPartialUpdate**](WebhooksApi.md#WebhooksPartialUpdate) | **Patch** /webhooks/{owner}/{repo}/{identifier}/ | Update a specific webhook in a repository.
[**WebhooksRead**](WebhooksApi.md#WebhooksRead) | **Get** /webhooks/{owner}/{repo}/{identifier}/ | Views for working with repository webhooks.



## WebhooksCreate

> RepositoryWebhook WebhooksCreate(ctx, owner, repo, optional)

Create a specific webhook in a repository.

Create a specific webhook in a repository.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***WebhooksCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WebhooksCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of WebhooksCreate**](WebhooksCreate.md)|  | 

### Return type

[**RepositoryWebhook**](RepositoryWebhook.md)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksDelete

> WebhooksDelete(ctx, owner, repo, identifier)

Delete a specific webhook in a repository.

Delete a specific webhook in a repository.

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


## WebhooksList

> []RepositoryWebhook WebhooksList(ctx, owner, repo, optional)

Get a list of all webhooks in a repository.

Get a list of all webhooks in a repository.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
 **optional** | ***WebhooksListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WebhooksListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| A page number within the paginated result set. | 
 **pageSize** | **optional.Int32**| Number of results to return per page. | 

### Return type

[**[]RepositoryWebhook**](RepositoryWebhook.md)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksPartialUpdate

> RepositoryWebhook WebhooksPartialUpdate(ctx, owner, repo, identifier, optional)

Update a specific webhook in a repository.

Update a specific webhook in a repository.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
**identifier** | **string**|  | 
 **optional** | ***WebhooksPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WebhooksPartialUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **data** | [**optional.Interface of WebhooksPartialUpdate**](WebhooksPartialUpdate.md)|  | 

### Return type

[**RepositoryWebhook**](RepositoryWebhook.md)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksRead

> RepositoryWebhook WebhooksRead(ctx, owner, repo, identifier)

Views for working with repository webhooks.

Views for working with repository webhooks.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string**|  | 
**repo** | **string**|  | 
**identifier** | **string**|  | 

### Return type

[**RepositoryWebhook**](RepositoryWebhook.md)

### Authorization

[apikey](../README.md#apikey), [csrf_token](../README.md#csrf_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

