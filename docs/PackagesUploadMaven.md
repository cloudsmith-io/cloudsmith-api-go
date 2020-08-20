# PackagesUploadMaven

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtifactId** | **string** | The ID of the artifact. | [optional] 
**GroupId** | **string** | Artifact&#39;s group ID. | [optional] 
**JavadocFile** | **string** | Adds bundled Java documentation to the Maven package | [optional] 
**PackageFile** | **string** | The primary file for the package. | 
**Packaging** | **string** | Artifact&#39;s Maven packaging type. | [optional] 
**PomFile** | **string** | The POM file is an XML file containing the Maven coordinates. | [optional] 
**Republish** | **bool** | If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate. | [optional] 
**SourcesFile** | **string** | Adds bundled Java source code to the Maven package. | [optional] 
**Tags** | **string** | A comma-separated values list of tags to add to the package. | [optional] 
**TestsFile** | **string** | Adds bundled Java tests to the Maven package. | [optional] 
**Version** | **string** | The raw version for this package. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


