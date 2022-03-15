# PackagesValidateUploadMaven

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtifactId** | Pointer to **string** | The ID of the artifact. | [optional] 
**GroupId** | Pointer to **string** | Artifact&#39;s group ID. | [optional] 
**JavadocFile** | Pointer to **string** | Adds bundled Java documentation to the Maven package | [optional] 
**PackageFile** | **string** | The primary file for the package. | 
**Packaging** | Pointer to **string** | Artifact&#39;s Maven packaging type. | [optional] 
**PomFile** | Pointer to **string** | The POM file is an XML file containing the Maven coordinates. | [optional] 
**Republish** | Pointer to **bool** | If true, the uploaded package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate. | [optional] 
**SourcesFile** | Pointer to **string** | Adds bundled Java source code to the Maven package. | [optional] 
**Tags** | Pointer to **string** | A comma-separated values list of tags to add to the package. | [optional] 
**TestsFile** | Pointer to **string** | Adds bundled Java tests to the Maven package. | [optional] 
**Version** | Pointer to **string** | The raw version for this package. | [optional] 

## Methods

### NewPackagesValidateUploadMaven

`func NewPackagesValidateUploadMaven(packageFile string, ) *PackagesValidateUploadMaven`

NewPackagesValidateUploadMaven instantiates a new PackagesValidateUploadMaven object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackagesValidateUploadMavenWithDefaults

`func NewPackagesValidateUploadMavenWithDefaults() *PackagesValidateUploadMaven`

NewPackagesValidateUploadMavenWithDefaults instantiates a new PackagesValidateUploadMaven object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactId

`func (o *PackagesValidateUploadMaven) GetArtifactId() string`

GetArtifactId returns the ArtifactId field if non-nil, zero value otherwise.

### GetArtifactIdOk

`func (o *PackagesValidateUploadMaven) GetArtifactIdOk() (*string, bool)`

GetArtifactIdOk returns a tuple with the ArtifactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactId

`func (o *PackagesValidateUploadMaven) SetArtifactId(v string)`

SetArtifactId sets ArtifactId field to given value.

### HasArtifactId

`func (o *PackagesValidateUploadMaven) HasArtifactId() bool`

HasArtifactId returns a boolean if a field has been set.

### GetGroupId

`func (o *PackagesValidateUploadMaven) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *PackagesValidateUploadMaven) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *PackagesValidateUploadMaven) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *PackagesValidateUploadMaven) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetJavadocFile

`func (o *PackagesValidateUploadMaven) GetJavadocFile() string`

GetJavadocFile returns the JavadocFile field if non-nil, zero value otherwise.

### GetJavadocFileOk

`func (o *PackagesValidateUploadMaven) GetJavadocFileOk() (*string, bool)`

GetJavadocFileOk returns a tuple with the JavadocFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavadocFile

`func (o *PackagesValidateUploadMaven) SetJavadocFile(v string)`

SetJavadocFile sets JavadocFile field to given value.

### HasJavadocFile

`func (o *PackagesValidateUploadMaven) HasJavadocFile() bool`

HasJavadocFile returns a boolean if a field has been set.

### GetPackageFile

`func (o *PackagesValidateUploadMaven) GetPackageFile() string`

GetPackageFile returns the PackageFile field if non-nil, zero value otherwise.

### GetPackageFileOk

`func (o *PackagesValidateUploadMaven) GetPackageFileOk() (*string, bool)`

GetPackageFileOk returns a tuple with the PackageFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageFile

`func (o *PackagesValidateUploadMaven) SetPackageFile(v string)`

SetPackageFile sets PackageFile field to given value.


### GetPackaging

`func (o *PackagesValidateUploadMaven) GetPackaging() string`

GetPackaging returns the Packaging field if non-nil, zero value otherwise.

### GetPackagingOk

`func (o *PackagesValidateUploadMaven) GetPackagingOk() (*string, bool)`

GetPackagingOk returns a tuple with the Packaging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackaging

`func (o *PackagesValidateUploadMaven) SetPackaging(v string)`

SetPackaging sets Packaging field to given value.

### HasPackaging

`func (o *PackagesValidateUploadMaven) HasPackaging() bool`

HasPackaging returns a boolean if a field has been set.

### GetPomFile

`func (o *PackagesValidateUploadMaven) GetPomFile() string`

GetPomFile returns the PomFile field if non-nil, zero value otherwise.

### GetPomFileOk

`func (o *PackagesValidateUploadMaven) GetPomFileOk() (*string, bool)`

GetPomFileOk returns a tuple with the PomFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPomFile

`func (o *PackagesValidateUploadMaven) SetPomFile(v string)`

SetPomFile sets PomFile field to given value.

### HasPomFile

`func (o *PackagesValidateUploadMaven) HasPomFile() bool`

HasPomFile returns a boolean if a field has been set.

### GetRepublish

`func (o *PackagesValidateUploadMaven) GetRepublish() bool`

GetRepublish returns the Republish field if non-nil, zero value otherwise.

### GetRepublishOk

`func (o *PackagesValidateUploadMaven) GetRepublishOk() (*bool, bool)`

GetRepublishOk returns a tuple with the Republish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepublish

`func (o *PackagesValidateUploadMaven) SetRepublish(v bool)`

SetRepublish sets Republish field to given value.

### HasRepublish

`func (o *PackagesValidateUploadMaven) HasRepublish() bool`

HasRepublish returns a boolean if a field has been set.

### GetSourcesFile

`func (o *PackagesValidateUploadMaven) GetSourcesFile() string`

GetSourcesFile returns the SourcesFile field if non-nil, zero value otherwise.

### GetSourcesFileOk

`func (o *PackagesValidateUploadMaven) GetSourcesFileOk() (*string, bool)`

GetSourcesFileOk returns a tuple with the SourcesFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcesFile

`func (o *PackagesValidateUploadMaven) SetSourcesFile(v string)`

SetSourcesFile sets SourcesFile field to given value.

### HasSourcesFile

`func (o *PackagesValidateUploadMaven) HasSourcesFile() bool`

HasSourcesFile returns a boolean if a field has been set.

### GetTags

`func (o *PackagesValidateUploadMaven) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PackagesValidateUploadMaven) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PackagesValidateUploadMaven) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PackagesValidateUploadMaven) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTestsFile

`func (o *PackagesValidateUploadMaven) GetTestsFile() string`

GetTestsFile returns the TestsFile field if non-nil, zero value otherwise.

### GetTestsFileOk

`func (o *PackagesValidateUploadMaven) GetTestsFileOk() (*string, bool)`

GetTestsFileOk returns a tuple with the TestsFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestsFile

`func (o *PackagesValidateUploadMaven) SetTestsFile(v string)`

SetTestsFile sets TestsFile field to given value.

### HasTestsFile

`func (o *PackagesValidateUploadMaven) HasTestsFile() bool`

HasTestsFile returns a boolean if a field has been set.

### GetVersion

`func (o *PackagesValidateUploadMaven) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PackagesValidateUploadMaven) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PackagesValidateUploadMaven) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PackagesValidateUploadMaven) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


