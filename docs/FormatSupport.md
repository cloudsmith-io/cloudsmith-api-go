# FormatSupport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dependencies** | **bool** | If true the package format supports dependencies | 
**Distributions** | **bool** | If true the package format supports distributions | 
**FileLists** | **bool** | If true the package format supports file lists | 
**Metadata** | **bool** | If true the package format supports metadata | 
**Versioning** | **bool** | If true the package format supports versioning | 

## Methods

### NewFormatSupport

`func NewFormatSupport(dependencies bool, distributions bool, fileLists bool, metadata bool, versioning bool, ) *FormatSupport`

NewFormatSupport instantiates a new FormatSupport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormatSupportWithDefaults

`func NewFormatSupportWithDefaults() *FormatSupport`

NewFormatSupportWithDefaults instantiates a new FormatSupport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDependencies

`func (o *FormatSupport) GetDependencies() bool`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *FormatSupport) GetDependenciesOk() (*bool, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *FormatSupport) SetDependencies(v bool)`

SetDependencies sets Dependencies field to given value.


### GetDistributions

`func (o *FormatSupport) GetDistributions() bool`

GetDistributions returns the Distributions field if non-nil, zero value otherwise.

### GetDistributionsOk

`func (o *FormatSupport) GetDistributionsOk() (*bool, bool)`

GetDistributionsOk returns a tuple with the Distributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributions

`func (o *FormatSupport) SetDistributions(v bool)`

SetDistributions sets Distributions field to given value.


### GetFileLists

`func (o *FormatSupport) GetFileLists() bool`

GetFileLists returns the FileLists field if non-nil, zero value otherwise.

### GetFileListsOk

`func (o *FormatSupport) GetFileListsOk() (*bool, bool)`

GetFileListsOk returns a tuple with the FileLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileLists

`func (o *FormatSupport) SetFileLists(v bool)`

SetFileLists sets FileLists field to given value.


### GetMetadata

`func (o *FormatSupport) GetMetadata() bool`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FormatSupport) GetMetadataOk() (*bool, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FormatSupport) SetMetadata(v bool)`

SetMetadata sets Metadata field to given value.


### GetVersioning

`func (o *FormatSupport) GetVersioning() bool`

GetVersioning returns the Versioning field if non-nil, zero value otherwise.

### GetVersioningOk

`func (o *FormatSupport) GetVersioningOk() (*bool, bool)`

GetVersioningOk returns a tuple with the Versioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersioning

`func (o *FormatSupport) SetVersioning(v bool)`

SetVersioning sets Versioning field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


