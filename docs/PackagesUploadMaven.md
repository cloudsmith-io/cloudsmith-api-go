# PackagesUploadMaven

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

### NewPackagesUploadMaven

`func NewPackagesUploadMaven(packageFile string, ) *PackagesUploadMaven`

NewPackagesUploadMaven instantiates a new PackagesUploadMaven object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackagesUploadMavenWithDefaults

`func NewPackagesUploadMavenWithDefaults() *PackagesUploadMaven`

NewPackagesUploadMavenWithDefaults instantiates a new PackagesUploadMaven object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactId

`func (o *PackagesUploadMaven) GetArtifactId() string`

GetArtifactId returns the ArtifactId field if non-nil, zero value otherwise.

### GetArtifactIdOk

`func (o *PackagesUploadMaven) GetArtifactIdOk() (*string, bool)`

GetArtifactIdOk returns a tuple with the ArtifactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactId

`func (o *PackagesUploadMaven) SetArtifactId(v string)`

SetArtifactId sets ArtifactId field to given value.

### HasArtifactId

`func (o *PackagesUploadMaven) HasArtifactId() bool`

HasArtifactId returns a boolean if a field has been set.

### GetGroupId

`func (o *PackagesUploadMaven) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *PackagesUploadMaven) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *PackagesUploadMaven) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *PackagesUploadMaven) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetJavadocFile

`func (o *PackagesUploadMaven) GetJavadocFile() string`

GetJavadocFile returns the JavadocFile field if non-nil, zero value otherwise.

### GetJavadocFileOk

`func (o *PackagesUploadMaven) GetJavadocFileOk() (*string, bool)`

GetJavadocFileOk returns a tuple with the JavadocFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavadocFile

`func (o *PackagesUploadMaven) SetJavadocFile(v string)`

SetJavadocFile sets JavadocFile field to given value.

### HasJavadocFile

`func (o *PackagesUploadMaven) HasJavadocFile() bool`

HasJavadocFile returns a boolean if a field has been set.

### GetPackageFile

`func (o *PackagesUploadMaven) GetPackageFile() string`

GetPackageFile returns the PackageFile field if non-nil, zero value otherwise.

### GetPackageFileOk

`func (o *PackagesUploadMaven) GetPackageFileOk() (*string, bool)`

GetPackageFileOk returns a tuple with the PackageFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageFile

`func (o *PackagesUploadMaven) SetPackageFile(v string)`

SetPackageFile sets PackageFile field to given value.


### GetPackaging

`func (o *PackagesUploadMaven) GetPackaging() string`

GetPackaging returns the Packaging field if non-nil, zero value otherwise.

### GetPackagingOk

`func (o *PackagesUploadMaven) GetPackagingOk() (*string, bool)`

GetPackagingOk returns a tuple with the Packaging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackaging

`func (o *PackagesUploadMaven) SetPackaging(v string)`

SetPackaging sets Packaging field to given value.

### HasPackaging

`func (o *PackagesUploadMaven) HasPackaging() bool`

HasPackaging returns a boolean if a field has been set.

### GetPomFile

`func (o *PackagesUploadMaven) GetPomFile() string`

GetPomFile returns the PomFile field if non-nil, zero value otherwise.

### GetPomFileOk

`func (o *PackagesUploadMaven) GetPomFileOk() (*string, bool)`

GetPomFileOk returns a tuple with the PomFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPomFile

`func (o *PackagesUploadMaven) SetPomFile(v string)`

SetPomFile sets PomFile field to given value.

### HasPomFile

`func (o *PackagesUploadMaven) HasPomFile() bool`

HasPomFile returns a boolean if a field has been set.

### GetRepublish

`func (o *PackagesUploadMaven) GetRepublish() bool`

GetRepublish returns the Republish field if non-nil, zero value otherwise.

### GetRepublishOk

`func (o *PackagesUploadMaven) GetRepublishOk() (*bool, bool)`

GetRepublishOk returns a tuple with the Republish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepublish

`func (o *PackagesUploadMaven) SetRepublish(v bool)`

SetRepublish sets Republish field to given value.

### HasRepublish

`func (o *PackagesUploadMaven) HasRepublish() bool`

HasRepublish returns a boolean if a field has been set.

### GetSourcesFile

`func (o *PackagesUploadMaven) GetSourcesFile() string`

GetSourcesFile returns the SourcesFile field if non-nil, zero value otherwise.

### GetSourcesFileOk

`func (o *PackagesUploadMaven) GetSourcesFileOk() (*string, bool)`

GetSourcesFileOk returns a tuple with the SourcesFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcesFile

`func (o *PackagesUploadMaven) SetSourcesFile(v string)`

SetSourcesFile sets SourcesFile field to given value.

### HasSourcesFile

`func (o *PackagesUploadMaven) HasSourcesFile() bool`

HasSourcesFile returns a boolean if a field has been set.

### GetTags

`func (o *PackagesUploadMaven) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PackagesUploadMaven) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PackagesUploadMaven) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PackagesUploadMaven) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTestsFile

`func (o *PackagesUploadMaven) GetTestsFile() string`

GetTestsFile returns the TestsFile field if non-nil, zero value otherwise.

### GetTestsFileOk

`func (o *PackagesUploadMaven) GetTestsFileOk() (*string, bool)`

GetTestsFileOk returns a tuple with the TestsFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestsFile

`func (o *PackagesUploadMaven) SetTestsFile(v string)`

SetTestsFile sets TestsFile field to given value.

### HasTestsFile

`func (o *PackagesUploadMaven) HasTestsFile() bool`

HasTestsFile returns a boolean if a field has been set.

### GetVersion

`func (o *PackagesUploadMaven) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PackagesUploadMaven) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PackagesUploadMaven) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PackagesUploadMaven) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


