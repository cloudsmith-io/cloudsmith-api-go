# RepositoryWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** |  | [optional] 
**CreatedBy** | **string** |  | [optional] 
**CreatedByUrl** | **string** |  | [optional] 
**DisableReason** | **string** |  | [optional] 
**DisableReasonStr** | **string** |  | [optional] 
**Events** | **[]string** |  | 
**Identifier** | **int64** |  | [optional] 
**IsActive** | **bool** | If enabled, the webhook will trigger on events and send payloads to the configured target URL. | [optional] 
**IsLastResponseBad** | **bool** |  | [optional] 
**LastResponseStatus** | **int64** |  | [optional] 
**LastResponseStatusStr** | **string** |  | [optional] 
**NumSent** | **int64** |  | [optional] 
**PackageQuery** | **string** | The package-based search query for webhooks to fire. This uses the same syntax as the standard search used for repositories, and also supports boolean logic operators such as OR/AND/NOT and parentheses for grouping. If a package does not match, the webhook will not fire. | [optional] 
**RequestBodyFormat** | **string** | The format of the payloads for webhook requests. | [optional] 
**RequestBodyFormatStr** | **string** |  | [optional] 
**RequestBodyTemplateFormat** | **string** | The format of the payloads for webhook requests. | [optional] 
**RequestBodyTemplateFormatStr** | **string** |  | [optional] 
**RequestContentType** | **string** | The value that will be sent for the &#39;Content Type&#39; header.  | [optional] 
**SecretHeader** | **string** | The header to send the predefined secret in. This must be unique from existing headers or it won&#39;t be sent. You can use this as a form of authentication on the endpoint side. | [optional] 
**SelfUrl** | **string** |  | [optional] 
**SlugPerm** | **string** |  | [optional] 
**TargetUrl** | **string** | The destination URL that webhook payloads will be POST&#39;ed to. | 
**Templates** | [**[]WebhooksOwnerRepoTemplates**](_webhooks__owner___repo___templates.md) |  | 
**UpdatedAt** | **string** |  | [optional] 
**UpdatedBy** | **string** |  | [optional] 
**UpdatedByUrl** | **string** |  | [optional] 
**VerifySsl** | **bool** | If enabled, SSL certificates is verified when webhooks are sent. It&#39;s recommended to leave this enabled as not verifying the integrity of SSL certificates leaves you susceptible to Man-in-the-Middle (MITM) attacks. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


