# MavenPackageUploadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtifactId** | Pointer to **NullableString** | The ID of the artifact. | [optional] 
**GroupId** | Pointer to **NullableString** | Artifact&#39;s group ID. | [optional] 
**JavadocFile** | Pointer to **NullableString** | Adds bundled Java documentation to the Maven package | [optional] 
**PackageFile** | **string** | The primary file for the package. | 
**Packaging** | Pointer to **NullableString** | Artifact&#39;s Maven packaging type. | [optional] 
**PomFile** | Pointer to **NullableString** | The POM file is an XML file containing the Maven coordinates. | [optional] 
**Republish** | Pointer to **bool** | If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate. | [optional] 
**SourcesFile** | Pointer to **NullableString** | Adds bundled Java source code to the Maven package. | [optional] 
**Tags** | Pointer to **NullableString** | A comma-separated values list of tags to add to the package. | [optional] 
**TestsFile** | Pointer to **NullableString** | Adds bundled Java tests to the Maven package. | [optional] 
**Version** | Pointer to **NullableString** | The raw version for this package. | [optional] 

## Methods

### NewMavenPackageUploadRequest

`func NewMavenPackageUploadRequest(packageFile string, ) *MavenPackageUploadRequest`

NewMavenPackageUploadRequest instantiates a new MavenPackageUploadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMavenPackageUploadRequestWithDefaults

`func NewMavenPackageUploadRequestWithDefaults() *MavenPackageUploadRequest`

NewMavenPackageUploadRequestWithDefaults instantiates a new MavenPackageUploadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactId

`func (o *MavenPackageUploadRequest) GetArtifactId() string`

GetArtifactId returns the ArtifactId field if non-nil, zero value otherwise.

### GetArtifactIdOk

`func (o *MavenPackageUploadRequest) GetArtifactIdOk() (*string, bool)`

GetArtifactIdOk returns a tuple with the ArtifactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactId

`func (o *MavenPackageUploadRequest) SetArtifactId(v string)`

SetArtifactId sets ArtifactId field to given value.

### HasArtifactId

`func (o *MavenPackageUploadRequest) HasArtifactId() bool`

HasArtifactId returns a boolean if a field has been set.

### SetArtifactIdNil

`func (o *MavenPackageUploadRequest) SetArtifactIdNil(b bool)`

 SetArtifactIdNil sets the value for ArtifactId to be an explicit nil

### UnsetArtifactId
`func (o *MavenPackageUploadRequest) UnsetArtifactId()`

UnsetArtifactId ensures that no value is present for ArtifactId, not even an explicit nil
### GetGroupId

`func (o *MavenPackageUploadRequest) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *MavenPackageUploadRequest) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *MavenPackageUploadRequest) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *MavenPackageUploadRequest) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *MavenPackageUploadRequest) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *MavenPackageUploadRequest) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetJavadocFile

`func (o *MavenPackageUploadRequest) GetJavadocFile() string`

GetJavadocFile returns the JavadocFile field if non-nil, zero value otherwise.

### GetJavadocFileOk

`func (o *MavenPackageUploadRequest) GetJavadocFileOk() (*string, bool)`

GetJavadocFileOk returns a tuple with the JavadocFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavadocFile

`func (o *MavenPackageUploadRequest) SetJavadocFile(v string)`

SetJavadocFile sets JavadocFile field to given value.

### HasJavadocFile

`func (o *MavenPackageUploadRequest) HasJavadocFile() bool`

HasJavadocFile returns a boolean if a field has been set.

### SetJavadocFileNil

`func (o *MavenPackageUploadRequest) SetJavadocFileNil(b bool)`

 SetJavadocFileNil sets the value for JavadocFile to be an explicit nil

### UnsetJavadocFile
`func (o *MavenPackageUploadRequest) UnsetJavadocFile()`

UnsetJavadocFile ensures that no value is present for JavadocFile, not even an explicit nil
### GetPackageFile

`func (o *MavenPackageUploadRequest) GetPackageFile() string`

GetPackageFile returns the PackageFile field if non-nil, zero value otherwise.

### GetPackageFileOk

`func (o *MavenPackageUploadRequest) GetPackageFileOk() (*string, bool)`

GetPackageFileOk returns a tuple with the PackageFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageFile

`func (o *MavenPackageUploadRequest) SetPackageFile(v string)`

SetPackageFile sets PackageFile field to given value.


### GetPackaging

`func (o *MavenPackageUploadRequest) GetPackaging() string`

GetPackaging returns the Packaging field if non-nil, zero value otherwise.

### GetPackagingOk

`func (o *MavenPackageUploadRequest) GetPackagingOk() (*string, bool)`

GetPackagingOk returns a tuple with the Packaging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackaging

`func (o *MavenPackageUploadRequest) SetPackaging(v string)`

SetPackaging sets Packaging field to given value.

### HasPackaging

`func (o *MavenPackageUploadRequest) HasPackaging() bool`

HasPackaging returns a boolean if a field has been set.

### SetPackagingNil

`func (o *MavenPackageUploadRequest) SetPackagingNil(b bool)`

 SetPackagingNil sets the value for Packaging to be an explicit nil

### UnsetPackaging
`func (o *MavenPackageUploadRequest) UnsetPackaging()`

UnsetPackaging ensures that no value is present for Packaging, not even an explicit nil
### GetPomFile

`func (o *MavenPackageUploadRequest) GetPomFile() string`

GetPomFile returns the PomFile field if non-nil, zero value otherwise.

### GetPomFileOk

`func (o *MavenPackageUploadRequest) GetPomFileOk() (*string, bool)`

GetPomFileOk returns a tuple with the PomFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPomFile

`func (o *MavenPackageUploadRequest) SetPomFile(v string)`

SetPomFile sets PomFile field to given value.

### HasPomFile

`func (o *MavenPackageUploadRequest) HasPomFile() bool`

HasPomFile returns a boolean if a field has been set.

### SetPomFileNil

`func (o *MavenPackageUploadRequest) SetPomFileNil(b bool)`

 SetPomFileNil sets the value for PomFile to be an explicit nil

### UnsetPomFile
`func (o *MavenPackageUploadRequest) UnsetPomFile()`

UnsetPomFile ensures that no value is present for PomFile, not even an explicit nil
### GetRepublish

`func (o *MavenPackageUploadRequest) GetRepublish() bool`

GetRepublish returns the Republish field if non-nil, zero value otherwise.

### GetRepublishOk

`func (o *MavenPackageUploadRequest) GetRepublishOk() (*bool, bool)`

GetRepublishOk returns a tuple with the Republish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepublish

`func (o *MavenPackageUploadRequest) SetRepublish(v bool)`

SetRepublish sets Republish field to given value.

### HasRepublish

`func (o *MavenPackageUploadRequest) HasRepublish() bool`

HasRepublish returns a boolean if a field has been set.

### GetSourcesFile

`func (o *MavenPackageUploadRequest) GetSourcesFile() string`

GetSourcesFile returns the SourcesFile field if non-nil, zero value otherwise.

### GetSourcesFileOk

`func (o *MavenPackageUploadRequest) GetSourcesFileOk() (*string, bool)`

GetSourcesFileOk returns a tuple with the SourcesFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcesFile

`func (o *MavenPackageUploadRequest) SetSourcesFile(v string)`

SetSourcesFile sets SourcesFile field to given value.

### HasSourcesFile

`func (o *MavenPackageUploadRequest) HasSourcesFile() bool`

HasSourcesFile returns a boolean if a field has been set.

### SetSourcesFileNil

`func (o *MavenPackageUploadRequest) SetSourcesFileNil(b bool)`

 SetSourcesFileNil sets the value for SourcesFile to be an explicit nil

### UnsetSourcesFile
`func (o *MavenPackageUploadRequest) UnsetSourcesFile()`

UnsetSourcesFile ensures that no value is present for SourcesFile, not even an explicit nil
### GetTags

`func (o *MavenPackageUploadRequest) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MavenPackageUploadRequest) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MavenPackageUploadRequest) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MavenPackageUploadRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *MavenPackageUploadRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *MavenPackageUploadRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetTestsFile

`func (o *MavenPackageUploadRequest) GetTestsFile() string`

GetTestsFile returns the TestsFile field if non-nil, zero value otherwise.

### GetTestsFileOk

`func (o *MavenPackageUploadRequest) GetTestsFileOk() (*string, bool)`

GetTestsFileOk returns a tuple with the TestsFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestsFile

`func (o *MavenPackageUploadRequest) SetTestsFile(v string)`

SetTestsFile sets TestsFile field to given value.

### HasTestsFile

`func (o *MavenPackageUploadRequest) HasTestsFile() bool`

HasTestsFile returns a boolean if a field has been set.

### SetTestsFileNil

`func (o *MavenPackageUploadRequest) SetTestsFileNil(b bool)`

 SetTestsFileNil sets the value for TestsFile to be an explicit nil

### UnsetTestsFile
`func (o *MavenPackageUploadRequest) UnsetTestsFile()`

UnsetTestsFile ensures that no value is present for TestsFile, not even an explicit nil
### GetVersion

`func (o *MavenPackageUploadRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MavenPackageUploadRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MavenPackageUploadRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MavenPackageUploadRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *MavenPackageUploadRequest) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *MavenPackageUploadRequest) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


