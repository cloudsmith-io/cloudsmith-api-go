# PackagesValidateUploadRaw

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | **string** | A custom content/media (also known as MIME) type to be sent when downloading this file. By default Cloudsmith will attempt to detect the type, but if you need to override it, you can specify it here. | [optional] 
**Description** | **string** | A textual description of this package. | [optional] 
**Name** | **string** | The name of this package. | [optional] 
**PackageFile** | **string** | The primary file for the package. | 
**Republish** | **bool** | If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate. | [optional] 
**Summary** | **string** | A one-liner synopsis of this package. | [optional] 
**Version** | **string** | The raw version for this package. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


