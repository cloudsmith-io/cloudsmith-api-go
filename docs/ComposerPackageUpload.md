# ComposerPackageUpload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Architectures** | Pointer to [**[]Architecture**](Architecture.md) |  | [optional] [readonly] 
**CdnUrl** | Pointer to **NullableString** |  | [optional] [readonly] 
**ChecksumMd5** | Pointer to **string** |  | [optional] [readonly] 
**ChecksumSha1** | Pointer to **string** |  | [optional] [readonly] 
**ChecksumSha256** | Pointer to **string** |  | [optional] [readonly] 
**ChecksumSha512** | Pointer to **string** |  | [optional] [readonly] 
**DependenciesChecksumMd5** | Pointer to **NullableString** | A checksum of all of the package&#39;s dependencies. | [optional] [readonly] 
**DependenciesUrl** | Pointer to **string** |  | [optional] [readonly] 
**Description** | Pointer to **NullableString** | A textual description of this package. | [optional] [readonly] 
**Distro** | Pointer to [**NullableDistribution**](Distribution.md) |  | [optional] 
**DistroVersion** | Pointer to [**DistributionVersion**](DistributionVersion.md) |  | [optional] 
**Downloads** | Pointer to **int64** |  | [optional] [readonly] 
**Epoch** | Pointer to **NullableInt64** | The epoch of the package version (if any). | [optional] [readonly] 
**Extension** | Pointer to **string** |  | [optional] [readonly] 
**Filename** | Pointer to **string** |  | [optional] [readonly] 
**Files** | Pointer to [**[]PackageFile**](PackageFile.md) |  | [optional] [readonly] 
**Format** | Pointer to **string** |  | [optional] [readonly] 
**FormatUrl** | Pointer to **string** |  | [optional] [readonly] 
**IdentifierPerm** | Pointer to **string** | Unique and permanent identifier for the package. | [optional] [readonly] 
**Indexed** | Pointer to **bool** |  | [optional] [readonly] 
**IsDownloadable** | Pointer to **bool** |  | [optional] [readonly] 
**IsQuarantined** | Pointer to **bool** |  | [optional] [readonly] 
**IsSyncAwaiting** | Pointer to **bool** |  | [optional] [readonly] 
**IsSyncCompleted** | Pointer to **bool** |  | [optional] [readonly] 
**IsSyncFailed** | Pointer to **bool** |  | [optional] [readonly] 
**IsSyncInFlight** | Pointer to **bool** |  | [optional] [readonly] 
**IsSyncInProgress** | Pointer to **bool** |  | [optional] [readonly] 
**License** | Pointer to **NullableString** | The license of this package. | [optional] [readonly] 
**Name** | Pointer to **NullableString** | The name of this package. | [optional] [readonly] 
**Namespace** | Pointer to **string** |  | [optional] [readonly] 
**NamespaceUrl** | Pointer to **string** |  | [optional] [readonly] 
**NumFiles** | Pointer to **int64** |  | [optional] [readonly] 
**OriginRepository** | Pointer to **string** |  | [optional] [readonly] 
**OriginRepositoryUrl** | Pointer to **string** |  | [optional] [readonly] 
**PackageType** | Pointer to **int64** | The type of package contents. | [optional] [readonly] 
**Release** | Pointer to **NullableString** | The release of the package version (if any). | [optional] [readonly] 
**Repository** | Pointer to **string** |  | [optional] [readonly] 
**RepositoryUrl** | Pointer to **string** |  | [optional] [readonly] 
**SecurityScanCompletedAt** | Pointer to **NullableTime** | The datetime the security scanning was completed. | [optional] [readonly] 
**SecurityScanStartedAt** | Pointer to **NullableTime** | The datetime the security scanning was started. | [optional] [readonly] 
**SecurityScanStatus** | Pointer to **NullableString** |  | [optional] [readonly] [default to "Awaiting Security Scan"]
**SecurityScanStatusUpdatedAt** | Pointer to **NullableTime** | The datetime the security scanning status was updated. | [optional] [readonly] 
**SelfHtmlUrl** | Pointer to **string** |  | [optional] [readonly] 
**SelfUrl** | Pointer to **string** |  | [optional] [readonly] 
**SignatureUrl** | Pointer to **NullableString** |  | [optional] [readonly] 
**Size** | Pointer to **int64** | The calculated size of the package. | [optional] [readonly] 
**Slug** | Pointer to **string** | The public unique identifier for the package. | [optional] [readonly] 
**SlugPerm** | Pointer to **string** |  | [optional] [readonly] 
**Stage** | Pointer to **int64** | The synchronisation (in progress) stage of the package. | [optional] [readonly] 
**StageStr** | Pointer to **string** |  | [optional] [readonly] 
**StageUpdatedAt** | Pointer to **time.Time** | The datetime the package stage was updated at. | [optional] [readonly] 
**Status** | Pointer to **int64** | The synchronisation status of the package. | [optional] [readonly] 
**StatusReason** | Pointer to **NullableString** | A textual description for the synchronous status reason (if any | [optional] [readonly] 
**StatusStr** | Pointer to **string** |  | [optional] [readonly] 
**StatusUpdatedAt** | Pointer to **time.Time** | The datetime the package status was updated at. | [optional] [readonly] 
**StatusUrl** | Pointer to **string** |  | [optional] [readonly] 
**Subtype** | Pointer to **string** |  | [optional] [readonly] 
**Summary** | Pointer to **NullableString** | A one-liner synopsis of this package. | [optional] [readonly] 
**SyncFinishedAt** | Pointer to **NullableTime** | The datetime the package sync was finished at. | [optional] [readonly] 
**SyncProgress** | Pointer to **int64** | Synchronisation progress (from 0-100) | [optional] [readonly] 
**TagsImmutable** | Pointer to **map[string]interface{}** | All tags on the package, grouped by tag type. This includes immutable tags, but doesn&#39;t distinguish them from mutable. To see which tags are immutable specifically, see the tags_immutable field. | [optional] 
**TypeDisplay** | Pointer to **string** |  | [optional] [readonly] 
**UploadedAt** | Pointer to **time.Time** | The date this package was uploaded. | [optional] [readonly] 
**Uploader** | Pointer to **string** |  | [optional] [readonly] 
**UploaderUrl** | Pointer to **string** |  | [optional] [readonly] 
**Version** | Pointer to **NullableString** | The raw version for this package. | [optional] [readonly] 
**VersionOrig** | Pointer to **string** |  | [optional] [readonly] 
**VulnerabilityScanResultsUrl** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewComposerPackageUpload

`func NewComposerPackageUpload() *ComposerPackageUpload`

NewComposerPackageUpload instantiates a new ComposerPackageUpload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComposerPackageUploadWithDefaults

`func NewComposerPackageUploadWithDefaults() *ComposerPackageUpload`

NewComposerPackageUploadWithDefaults instantiates a new ComposerPackageUpload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchitectures

`func (o *ComposerPackageUpload) GetArchitectures() []Architecture`

GetArchitectures returns the Architectures field if non-nil, zero value otherwise.

### GetArchitecturesOk

`func (o *ComposerPackageUpload) GetArchitecturesOk() (*[]Architecture, bool)`

GetArchitecturesOk returns a tuple with the Architectures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitectures

`func (o *ComposerPackageUpload) SetArchitectures(v []Architecture)`

SetArchitectures sets Architectures field to given value.

### HasArchitectures

`func (o *ComposerPackageUpload) HasArchitectures() bool`

HasArchitectures returns a boolean if a field has been set.

### GetCdnUrl

`func (o *ComposerPackageUpload) GetCdnUrl() string`

GetCdnUrl returns the CdnUrl field if non-nil, zero value otherwise.

### GetCdnUrlOk

`func (o *ComposerPackageUpload) GetCdnUrlOk() (*string, bool)`

GetCdnUrlOk returns a tuple with the CdnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnUrl

`func (o *ComposerPackageUpload) SetCdnUrl(v string)`

SetCdnUrl sets CdnUrl field to given value.

### HasCdnUrl

`func (o *ComposerPackageUpload) HasCdnUrl() bool`

HasCdnUrl returns a boolean if a field has been set.

### SetCdnUrlNil

`func (o *ComposerPackageUpload) SetCdnUrlNil(b bool)`

 SetCdnUrlNil sets the value for CdnUrl to be an explicit nil

### UnsetCdnUrl
`func (o *ComposerPackageUpload) UnsetCdnUrl()`

UnsetCdnUrl ensures that no value is present for CdnUrl, not even an explicit nil
### GetChecksumMd5

`func (o *ComposerPackageUpload) GetChecksumMd5() string`

GetChecksumMd5 returns the ChecksumMd5 field if non-nil, zero value otherwise.

### GetChecksumMd5Ok

`func (o *ComposerPackageUpload) GetChecksumMd5Ok() (*string, bool)`

GetChecksumMd5Ok returns a tuple with the ChecksumMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumMd5

`func (o *ComposerPackageUpload) SetChecksumMd5(v string)`

SetChecksumMd5 sets ChecksumMd5 field to given value.

### HasChecksumMd5

`func (o *ComposerPackageUpload) HasChecksumMd5() bool`

HasChecksumMd5 returns a boolean if a field has been set.

### GetChecksumSha1

`func (o *ComposerPackageUpload) GetChecksumSha1() string`

GetChecksumSha1 returns the ChecksumSha1 field if non-nil, zero value otherwise.

### GetChecksumSha1Ok

`func (o *ComposerPackageUpload) GetChecksumSha1Ok() (*string, bool)`

GetChecksumSha1Ok returns a tuple with the ChecksumSha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumSha1

`func (o *ComposerPackageUpload) SetChecksumSha1(v string)`

SetChecksumSha1 sets ChecksumSha1 field to given value.

### HasChecksumSha1

`func (o *ComposerPackageUpload) HasChecksumSha1() bool`

HasChecksumSha1 returns a boolean if a field has been set.

### GetChecksumSha256

`func (o *ComposerPackageUpload) GetChecksumSha256() string`

GetChecksumSha256 returns the ChecksumSha256 field if non-nil, zero value otherwise.

### GetChecksumSha256Ok

`func (o *ComposerPackageUpload) GetChecksumSha256Ok() (*string, bool)`

GetChecksumSha256Ok returns a tuple with the ChecksumSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumSha256

`func (o *ComposerPackageUpload) SetChecksumSha256(v string)`

SetChecksumSha256 sets ChecksumSha256 field to given value.

### HasChecksumSha256

`func (o *ComposerPackageUpload) HasChecksumSha256() bool`

HasChecksumSha256 returns a boolean if a field has been set.

### GetChecksumSha512

`func (o *ComposerPackageUpload) GetChecksumSha512() string`

GetChecksumSha512 returns the ChecksumSha512 field if non-nil, zero value otherwise.

### GetChecksumSha512Ok

`func (o *ComposerPackageUpload) GetChecksumSha512Ok() (*string, bool)`

GetChecksumSha512Ok returns a tuple with the ChecksumSha512 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumSha512

`func (o *ComposerPackageUpload) SetChecksumSha512(v string)`

SetChecksumSha512 sets ChecksumSha512 field to given value.

### HasChecksumSha512

`func (o *ComposerPackageUpload) HasChecksumSha512() bool`

HasChecksumSha512 returns a boolean if a field has been set.

### GetDependenciesChecksumMd5

`func (o *ComposerPackageUpload) GetDependenciesChecksumMd5() string`

GetDependenciesChecksumMd5 returns the DependenciesChecksumMd5 field if non-nil, zero value otherwise.

### GetDependenciesChecksumMd5Ok

`func (o *ComposerPackageUpload) GetDependenciesChecksumMd5Ok() (*string, bool)`

GetDependenciesChecksumMd5Ok returns a tuple with the DependenciesChecksumMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependenciesChecksumMd5

`func (o *ComposerPackageUpload) SetDependenciesChecksumMd5(v string)`

SetDependenciesChecksumMd5 sets DependenciesChecksumMd5 field to given value.

### HasDependenciesChecksumMd5

`func (o *ComposerPackageUpload) HasDependenciesChecksumMd5() bool`

HasDependenciesChecksumMd5 returns a boolean if a field has been set.

### SetDependenciesChecksumMd5Nil

`func (o *ComposerPackageUpload) SetDependenciesChecksumMd5Nil(b bool)`

 SetDependenciesChecksumMd5Nil sets the value for DependenciesChecksumMd5 to be an explicit nil

### UnsetDependenciesChecksumMd5
`func (o *ComposerPackageUpload) UnsetDependenciesChecksumMd5()`

UnsetDependenciesChecksumMd5 ensures that no value is present for DependenciesChecksumMd5, not even an explicit nil
### GetDependenciesUrl

`func (o *ComposerPackageUpload) GetDependenciesUrl() string`

GetDependenciesUrl returns the DependenciesUrl field if non-nil, zero value otherwise.

### GetDependenciesUrlOk

`func (o *ComposerPackageUpload) GetDependenciesUrlOk() (*string, bool)`

GetDependenciesUrlOk returns a tuple with the DependenciesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependenciesUrl

`func (o *ComposerPackageUpload) SetDependenciesUrl(v string)`

SetDependenciesUrl sets DependenciesUrl field to given value.

### HasDependenciesUrl

`func (o *ComposerPackageUpload) HasDependenciesUrl() bool`

HasDependenciesUrl returns a boolean if a field has been set.

### GetDescription

`func (o *ComposerPackageUpload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ComposerPackageUpload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ComposerPackageUpload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ComposerPackageUpload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ComposerPackageUpload) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ComposerPackageUpload) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDistro

`func (o *ComposerPackageUpload) GetDistro() Distribution`

GetDistro returns the Distro field if non-nil, zero value otherwise.

### GetDistroOk

`func (o *ComposerPackageUpload) GetDistroOk() (*Distribution, bool)`

GetDistroOk returns a tuple with the Distro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistro

`func (o *ComposerPackageUpload) SetDistro(v Distribution)`

SetDistro sets Distro field to given value.

### HasDistro

`func (o *ComposerPackageUpload) HasDistro() bool`

HasDistro returns a boolean if a field has been set.

### SetDistroNil

`func (o *ComposerPackageUpload) SetDistroNil(b bool)`

 SetDistroNil sets the value for Distro to be an explicit nil

### UnsetDistro
`func (o *ComposerPackageUpload) UnsetDistro()`

UnsetDistro ensures that no value is present for Distro, not even an explicit nil
### GetDistroVersion

`func (o *ComposerPackageUpload) GetDistroVersion() DistributionVersion`

GetDistroVersion returns the DistroVersion field if non-nil, zero value otherwise.

### GetDistroVersionOk

`func (o *ComposerPackageUpload) GetDistroVersionOk() (*DistributionVersion, bool)`

GetDistroVersionOk returns a tuple with the DistroVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistroVersion

`func (o *ComposerPackageUpload) SetDistroVersion(v DistributionVersion)`

SetDistroVersion sets DistroVersion field to given value.

### HasDistroVersion

`func (o *ComposerPackageUpload) HasDistroVersion() bool`

HasDistroVersion returns a boolean if a field has been set.

### GetDownloads

`func (o *ComposerPackageUpload) GetDownloads() int64`

GetDownloads returns the Downloads field if non-nil, zero value otherwise.

### GetDownloadsOk

`func (o *ComposerPackageUpload) GetDownloadsOk() (*int64, bool)`

GetDownloadsOk returns a tuple with the Downloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloads

`func (o *ComposerPackageUpload) SetDownloads(v int64)`

SetDownloads sets Downloads field to given value.

### HasDownloads

`func (o *ComposerPackageUpload) HasDownloads() bool`

HasDownloads returns a boolean if a field has been set.

### GetEpoch

`func (o *ComposerPackageUpload) GetEpoch() int64`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *ComposerPackageUpload) GetEpochOk() (*int64, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *ComposerPackageUpload) SetEpoch(v int64)`

SetEpoch sets Epoch field to given value.

### HasEpoch

`func (o *ComposerPackageUpload) HasEpoch() bool`

HasEpoch returns a boolean if a field has been set.

### SetEpochNil

`func (o *ComposerPackageUpload) SetEpochNil(b bool)`

 SetEpochNil sets the value for Epoch to be an explicit nil

### UnsetEpoch
`func (o *ComposerPackageUpload) UnsetEpoch()`

UnsetEpoch ensures that no value is present for Epoch, not even an explicit nil
### GetExtension

`func (o *ComposerPackageUpload) GetExtension() string`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *ComposerPackageUpload) GetExtensionOk() (*string, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *ComposerPackageUpload) SetExtension(v string)`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *ComposerPackageUpload) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### GetFilename

`func (o *ComposerPackageUpload) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *ComposerPackageUpload) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *ComposerPackageUpload) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *ComposerPackageUpload) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetFiles

`func (o *ComposerPackageUpload) GetFiles() []PackageFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ComposerPackageUpload) GetFilesOk() (*[]PackageFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ComposerPackageUpload) SetFiles(v []PackageFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ComposerPackageUpload) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetFormat

`func (o *ComposerPackageUpload) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ComposerPackageUpload) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ComposerPackageUpload) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *ComposerPackageUpload) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetFormatUrl

`func (o *ComposerPackageUpload) GetFormatUrl() string`

GetFormatUrl returns the FormatUrl field if non-nil, zero value otherwise.

### GetFormatUrlOk

`func (o *ComposerPackageUpload) GetFormatUrlOk() (*string, bool)`

GetFormatUrlOk returns a tuple with the FormatUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatUrl

`func (o *ComposerPackageUpload) SetFormatUrl(v string)`

SetFormatUrl sets FormatUrl field to given value.

### HasFormatUrl

`func (o *ComposerPackageUpload) HasFormatUrl() bool`

HasFormatUrl returns a boolean if a field has been set.

### GetIdentifierPerm

`func (o *ComposerPackageUpload) GetIdentifierPerm() string`

GetIdentifierPerm returns the IdentifierPerm field if non-nil, zero value otherwise.

### GetIdentifierPermOk

`func (o *ComposerPackageUpload) GetIdentifierPermOk() (*string, bool)`

GetIdentifierPermOk returns a tuple with the IdentifierPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierPerm

`func (o *ComposerPackageUpload) SetIdentifierPerm(v string)`

SetIdentifierPerm sets IdentifierPerm field to given value.

### HasIdentifierPerm

`func (o *ComposerPackageUpload) HasIdentifierPerm() bool`

HasIdentifierPerm returns a boolean if a field has been set.

### GetIndexed

`func (o *ComposerPackageUpload) GetIndexed() bool`

GetIndexed returns the Indexed field if non-nil, zero value otherwise.

### GetIndexedOk

`func (o *ComposerPackageUpload) GetIndexedOk() (*bool, bool)`

GetIndexedOk returns a tuple with the Indexed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexed

`func (o *ComposerPackageUpload) SetIndexed(v bool)`

SetIndexed sets Indexed field to given value.

### HasIndexed

`func (o *ComposerPackageUpload) HasIndexed() bool`

HasIndexed returns a boolean if a field has been set.

### GetIsDownloadable

`func (o *ComposerPackageUpload) GetIsDownloadable() bool`

GetIsDownloadable returns the IsDownloadable field if non-nil, zero value otherwise.

### GetIsDownloadableOk

`func (o *ComposerPackageUpload) GetIsDownloadableOk() (*bool, bool)`

GetIsDownloadableOk returns a tuple with the IsDownloadable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDownloadable

`func (o *ComposerPackageUpload) SetIsDownloadable(v bool)`

SetIsDownloadable sets IsDownloadable field to given value.

### HasIsDownloadable

`func (o *ComposerPackageUpload) HasIsDownloadable() bool`

HasIsDownloadable returns a boolean if a field has been set.

### GetIsQuarantined

`func (o *ComposerPackageUpload) GetIsQuarantined() bool`

GetIsQuarantined returns the IsQuarantined field if non-nil, zero value otherwise.

### GetIsQuarantinedOk

`func (o *ComposerPackageUpload) GetIsQuarantinedOk() (*bool, bool)`

GetIsQuarantinedOk returns a tuple with the IsQuarantined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQuarantined

`func (o *ComposerPackageUpload) SetIsQuarantined(v bool)`

SetIsQuarantined sets IsQuarantined field to given value.

### HasIsQuarantined

`func (o *ComposerPackageUpload) HasIsQuarantined() bool`

HasIsQuarantined returns a boolean if a field has been set.

### GetIsSyncAwaiting

`func (o *ComposerPackageUpload) GetIsSyncAwaiting() bool`

GetIsSyncAwaiting returns the IsSyncAwaiting field if non-nil, zero value otherwise.

### GetIsSyncAwaitingOk

`func (o *ComposerPackageUpload) GetIsSyncAwaitingOk() (*bool, bool)`

GetIsSyncAwaitingOk returns a tuple with the IsSyncAwaiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncAwaiting

`func (o *ComposerPackageUpload) SetIsSyncAwaiting(v bool)`

SetIsSyncAwaiting sets IsSyncAwaiting field to given value.

### HasIsSyncAwaiting

`func (o *ComposerPackageUpload) HasIsSyncAwaiting() bool`

HasIsSyncAwaiting returns a boolean if a field has been set.

### GetIsSyncCompleted

`func (o *ComposerPackageUpload) GetIsSyncCompleted() bool`

GetIsSyncCompleted returns the IsSyncCompleted field if non-nil, zero value otherwise.

### GetIsSyncCompletedOk

`func (o *ComposerPackageUpload) GetIsSyncCompletedOk() (*bool, bool)`

GetIsSyncCompletedOk returns a tuple with the IsSyncCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncCompleted

`func (o *ComposerPackageUpload) SetIsSyncCompleted(v bool)`

SetIsSyncCompleted sets IsSyncCompleted field to given value.

### HasIsSyncCompleted

`func (o *ComposerPackageUpload) HasIsSyncCompleted() bool`

HasIsSyncCompleted returns a boolean if a field has been set.

### GetIsSyncFailed

`func (o *ComposerPackageUpload) GetIsSyncFailed() bool`

GetIsSyncFailed returns the IsSyncFailed field if non-nil, zero value otherwise.

### GetIsSyncFailedOk

`func (o *ComposerPackageUpload) GetIsSyncFailedOk() (*bool, bool)`

GetIsSyncFailedOk returns a tuple with the IsSyncFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncFailed

`func (o *ComposerPackageUpload) SetIsSyncFailed(v bool)`

SetIsSyncFailed sets IsSyncFailed field to given value.

### HasIsSyncFailed

`func (o *ComposerPackageUpload) HasIsSyncFailed() bool`

HasIsSyncFailed returns a boolean if a field has been set.

### GetIsSyncInFlight

`func (o *ComposerPackageUpload) GetIsSyncInFlight() bool`

GetIsSyncInFlight returns the IsSyncInFlight field if non-nil, zero value otherwise.

### GetIsSyncInFlightOk

`func (o *ComposerPackageUpload) GetIsSyncInFlightOk() (*bool, bool)`

GetIsSyncInFlightOk returns a tuple with the IsSyncInFlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncInFlight

`func (o *ComposerPackageUpload) SetIsSyncInFlight(v bool)`

SetIsSyncInFlight sets IsSyncInFlight field to given value.

### HasIsSyncInFlight

`func (o *ComposerPackageUpload) HasIsSyncInFlight() bool`

HasIsSyncInFlight returns a boolean if a field has been set.

### GetIsSyncInProgress

`func (o *ComposerPackageUpload) GetIsSyncInProgress() bool`

GetIsSyncInProgress returns the IsSyncInProgress field if non-nil, zero value otherwise.

### GetIsSyncInProgressOk

`func (o *ComposerPackageUpload) GetIsSyncInProgressOk() (*bool, bool)`

GetIsSyncInProgressOk returns a tuple with the IsSyncInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncInProgress

`func (o *ComposerPackageUpload) SetIsSyncInProgress(v bool)`

SetIsSyncInProgress sets IsSyncInProgress field to given value.

### HasIsSyncInProgress

`func (o *ComposerPackageUpload) HasIsSyncInProgress() bool`

HasIsSyncInProgress returns a boolean if a field has been set.

### GetLicense

`func (o *ComposerPackageUpload) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *ComposerPackageUpload) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *ComposerPackageUpload) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *ComposerPackageUpload) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### SetLicenseNil

`func (o *ComposerPackageUpload) SetLicenseNil(b bool)`

 SetLicenseNil sets the value for License to be an explicit nil

### UnsetLicense
`func (o *ComposerPackageUpload) UnsetLicense()`

UnsetLicense ensures that no value is present for License, not even an explicit nil
### GetName

`func (o *ComposerPackageUpload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComposerPackageUpload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComposerPackageUpload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComposerPackageUpload) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ComposerPackageUpload) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ComposerPackageUpload) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNamespace

`func (o *ComposerPackageUpload) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ComposerPackageUpload) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ComposerPackageUpload) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ComposerPackageUpload) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetNamespaceUrl

`func (o *ComposerPackageUpload) GetNamespaceUrl() string`

GetNamespaceUrl returns the NamespaceUrl field if non-nil, zero value otherwise.

### GetNamespaceUrlOk

`func (o *ComposerPackageUpload) GetNamespaceUrlOk() (*string, bool)`

GetNamespaceUrlOk returns a tuple with the NamespaceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceUrl

`func (o *ComposerPackageUpload) SetNamespaceUrl(v string)`

SetNamespaceUrl sets NamespaceUrl field to given value.

### HasNamespaceUrl

`func (o *ComposerPackageUpload) HasNamespaceUrl() bool`

HasNamespaceUrl returns a boolean if a field has been set.

### GetNumFiles

`func (o *ComposerPackageUpload) GetNumFiles() int64`

GetNumFiles returns the NumFiles field if non-nil, zero value otherwise.

### GetNumFilesOk

`func (o *ComposerPackageUpload) GetNumFilesOk() (*int64, bool)`

GetNumFilesOk returns a tuple with the NumFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFiles

`func (o *ComposerPackageUpload) SetNumFiles(v int64)`

SetNumFiles sets NumFiles field to given value.

### HasNumFiles

`func (o *ComposerPackageUpload) HasNumFiles() bool`

HasNumFiles returns a boolean if a field has been set.

### GetOriginRepository

`func (o *ComposerPackageUpload) GetOriginRepository() string`

GetOriginRepository returns the OriginRepository field if non-nil, zero value otherwise.

### GetOriginRepositoryOk

`func (o *ComposerPackageUpload) GetOriginRepositoryOk() (*string, bool)`

GetOriginRepositoryOk returns a tuple with the OriginRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginRepository

`func (o *ComposerPackageUpload) SetOriginRepository(v string)`

SetOriginRepository sets OriginRepository field to given value.

### HasOriginRepository

`func (o *ComposerPackageUpload) HasOriginRepository() bool`

HasOriginRepository returns a boolean if a field has been set.

### GetOriginRepositoryUrl

`func (o *ComposerPackageUpload) GetOriginRepositoryUrl() string`

GetOriginRepositoryUrl returns the OriginRepositoryUrl field if non-nil, zero value otherwise.

### GetOriginRepositoryUrlOk

`func (o *ComposerPackageUpload) GetOriginRepositoryUrlOk() (*string, bool)`

GetOriginRepositoryUrlOk returns a tuple with the OriginRepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginRepositoryUrl

`func (o *ComposerPackageUpload) SetOriginRepositoryUrl(v string)`

SetOriginRepositoryUrl sets OriginRepositoryUrl field to given value.

### HasOriginRepositoryUrl

`func (o *ComposerPackageUpload) HasOriginRepositoryUrl() bool`

HasOriginRepositoryUrl returns a boolean if a field has been set.

### GetPackageType

`func (o *ComposerPackageUpload) GetPackageType() int64`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *ComposerPackageUpload) GetPackageTypeOk() (*int64, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *ComposerPackageUpload) SetPackageType(v int64)`

SetPackageType sets PackageType field to given value.

### HasPackageType

`func (o *ComposerPackageUpload) HasPackageType() bool`

HasPackageType returns a boolean if a field has been set.

### GetRelease

`func (o *ComposerPackageUpload) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *ComposerPackageUpload) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *ComposerPackageUpload) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *ComposerPackageUpload) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### SetReleaseNil

`func (o *ComposerPackageUpload) SetReleaseNil(b bool)`

 SetReleaseNil sets the value for Release to be an explicit nil

### UnsetRelease
`func (o *ComposerPackageUpload) UnsetRelease()`

UnsetRelease ensures that no value is present for Release, not even an explicit nil
### GetRepository

`func (o *ComposerPackageUpload) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *ComposerPackageUpload) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *ComposerPackageUpload) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *ComposerPackageUpload) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetRepositoryUrl

`func (o *ComposerPackageUpload) GetRepositoryUrl() string`

GetRepositoryUrl returns the RepositoryUrl field if non-nil, zero value otherwise.

### GetRepositoryUrlOk

`func (o *ComposerPackageUpload) GetRepositoryUrlOk() (*string, bool)`

GetRepositoryUrlOk returns a tuple with the RepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryUrl

`func (o *ComposerPackageUpload) SetRepositoryUrl(v string)`

SetRepositoryUrl sets RepositoryUrl field to given value.

### HasRepositoryUrl

`func (o *ComposerPackageUpload) HasRepositoryUrl() bool`

HasRepositoryUrl returns a boolean if a field has been set.

### GetSecurityScanCompletedAt

`func (o *ComposerPackageUpload) GetSecurityScanCompletedAt() time.Time`

GetSecurityScanCompletedAt returns the SecurityScanCompletedAt field if non-nil, zero value otherwise.

### GetSecurityScanCompletedAtOk

`func (o *ComposerPackageUpload) GetSecurityScanCompletedAtOk() (*time.Time, bool)`

GetSecurityScanCompletedAtOk returns a tuple with the SecurityScanCompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityScanCompletedAt

`func (o *ComposerPackageUpload) SetSecurityScanCompletedAt(v time.Time)`

SetSecurityScanCompletedAt sets SecurityScanCompletedAt field to given value.

### HasSecurityScanCompletedAt

`func (o *ComposerPackageUpload) HasSecurityScanCompletedAt() bool`

HasSecurityScanCompletedAt returns a boolean if a field has been set.

### SetSecurityScanCompletedAtNil

`func (o *ComposerPackageUpload) SetSecurityScanCompletedAtNil(b bool)`

 SetSecurityScanCompletedAtNil sets the value for SecurityScanCompletedAt to be an explicit nil

### UnsetSecurityScanCompletedAt
`func (o *ComposerPackageUpload) UnsetSecurityScanCompletedAt()`

UnsetSecurityScanCompletedAt ensures that no value is present for SecurityScanCompletedAt, not even an explicit nil
### GetSecurityScanStartedAt

`func (o *ComposerPackageUpload) GetSecurityScanStartedAt() time.Time`

GetSecurityScanStartedAt returns the SecurityScanStartedAt field if non-nil, zero value otherwise.

### GetSecurityScanStartedAtOk

`func (o *ComposerPackageUpload) GetSecurityScanStartedAtOk() (*time.Time, bool)`

GetSecurityScanStartedAtOk returns a tuple with the SecurityScanStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityScanStartedAt

`func (o *ComposerPackageUpload) SetSecurityScanStartedAt(v time.Time)`

SetSecurityScanStartedAt sets SecurityScanStartedAt field to given value.

### HasSecurityScanStartedAt

`func (o *ComposerPackageUpload) HasSecurityScanStartedAt() bool`

HasSecurityScanStartedAt returns a boolean if a field has been set.

### SetSecurityScanStartedAtNil

`func (o *ComposerPackageUpload) SetSecurityScanStartedAtNil(b bool)`

 SetSecurityScanStartedAtNil sets the value for SecurityScanStartedAt to be an explicit nil

### UnsetSecurityScanStartedAt
`func (o *ComposerPackageUpload) UnsetSecurityScanStartedAt()`

UnsetSecurityScanStartedAt ensures that no value is present for SecurityScanStartedAt, not even an explicit nil
### GetSecurityScanStatus

`func (o *ComposerPackageUpload) GetSecurityScanStatus() string`

GetSecurityScanStatus returns the SecurityScanStatus field if non-nil, zero value otherwise.

### GetSecurityScanStatusOk

`func (o *ComposerPackageUpload) GetSecurityScanStatusOk() (*string, bool)`

GetSecurityScanStatusOk returns a tuple with the SecurityScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityScanStatus

`func (o *ComposerPackageUpload) SetSecurityScanStatus(v string)`

SetSecurityScanStatus sets SecurityScanStatus field to given value.

### HasSecurityScanStatus

`func (o *ComposerPackageUpload) HasSecurityScanStatus() bool`

HasSecurityScanStatus returns a boolean if a field has been set.

### SetSecurityScanStatusNil

`func (o *ComposerPackageUpload) SetSecurityScanStatusNil(b bool)`

 SetSecurityScanStatusNil sets the value for SecurityScanStatus to be an explicit nil

### UnsetSecurityScanStatus
`func (o *ComposerPackageUpload) UnsetSecurityScanStatus()`

UnsetSecurityScanStatus ensures that no value is present for SecurityScanStatus, not even an explicit nil
### GetSecurityScanStatusUpdatedAt

`func (o *ComposerPackageUpload) GetSecurityScanStatusUpdatedAt() time.Time`

GetSecurityScanStatusUpdatedAt returns the SecurityScanStatusUpdatedAt field if non-nil, zero value otherwise.

### GetSecurityScanStatusUpdatedAtOk

`func (o *ComposerPackageUpload) GetSecurityScanStatusUpdatedAtOk() (*time.Time, bool)`

GetSecurityScanStatusUpdatedAtOk returns a tuple with the SecurityScanStatusUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityScanStatusUpdatedAt

`func (o *ComposerPackageUpload) SetSecurityScanStatusUpdatedAt(v time.Time)`

SetSecurityScanStatusUpdatedAt sets SecurityScanStatusUpdatedAt field to given value.

### HasSecurityScanStatusUpdatedAt

`func (o *ComposerPackageUpload) HasSecurityScanStatusUpdatedAt() bool`

HasSecurityScanStatusUpdatedAt returns a boolean if a field has been set.

### SetSecurityScanStatusUpdatedAtNil

`func (o *ComposerPackageUpload) SetSecurityScanStatusUpdatedAtNil(b bool)`

 SetSecurityScanStatusUpdatedAtNil sets the value for SecurityScanStatusUpdatedAt to be an explicit nil

### UnsetSecurityScanStatusUpdatedAt
`func (o *ComposerPackageUpload) UnsetSecurityScanStatusUpdatedAt()`

UnsetSecurityScanStatusUpdatedAt ensures that no value is present for SecurityScanStatusUpdatedAt, not even an explicit nil
### GetSelfHtmlUrl

`func (o *ComposerPackageUpload) GetSelfHtmlUrl() string`

GetSelfHtmlUrl returns the SelfHtmlUrl field if non-nil, zero value otherwise.

### GetSelfHtmlUrlOk

`func (o *ComposerPackageUpload) GetSelfHtmlUrlOk() (*string, bool)`

GetSelfHtmlUrlOk returns a tuple with the SelfHtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfHtmlUrl

`func (o *ComposerPackageUpload) SetSelfHtmlUrl(v string)`

SetSelfHtmlUrl sets SelfHtmlUrl field to given value.

### HasSelfHtmlUrl

`func (o *ComposerPackageUpload) HasSelfHtmlUrl() bool`

HasSelfHtmlUrl returns a boolean if a field has been set.

### GetSelfUrl

`func (o *ComposerPackageUpload) GetSelfUrl() string`

GetSelfUrl returns the SelfUrl field if non-nil, zero value otherwise.

### GetSelfUrlOk

`func (o *ComposerPackageUpload) GetSelfUrlOk() (*string, bool)`

GetSelfUrlOk returns a tuple with the SelfUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfUrl

`func (o *ComposerPackageUpload) SetSelfUrl(v string)`

SetSelfUrl sets SelfUrl field to given value.

### HasSelfUrl

`func (o *ComposerPackageUpload) HasSelfUrl() bool`

HasSelfUrl returns a boolean if a field has been set.

### GetSignatureUrl

`func (o *ComposerPackageUpload) GetSignatureUrl() string`

GetSignatureUrl returns the SignatureUrl field if non-nil, zero value otherwise.

### GetSignatureUrlOk

`func (o *ComposerPackageUpload) GetSignatureUrlOk() (*string, bool)`

GetSignatureUrlOk returns a tuple with the SignatureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureUrl

`func (o *ComposerPackageUpload) SetSignatureUrl(v string)`

SetSignatureUrl sets SignatureUrl field to given value.

### HasSignatureUrl

`func (o *ComposerPackageUpload) HasSignatureUrl() bool`

HasSignatureUrl returns a boolean if a field has been set.

### SetSignatureUrlNil

`func (o *ComposerPackageUpload) SetSignatureUrlNil(b bool)`

 SetSignatureUrlNil sets the value for SignatureUrl to be an explicit nil

### UnsetSignatureUrl
`func (o *ComposerPackageUpload) UnsetSignatureUrl()`

UnsetSignatureUrl ensures that no value is present for SignatureUrl, not even an explicit nil
### GetSize

`func (o *ComposerPackageUpload) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ComposerPackageUpload) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ComposerPackageUpload) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *ComposerPackageUpload) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSlug

`func (o *ComposerPackageUpload) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ComposerPackageUpload) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ComposerPackageUpload) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ComposerPackageUpload) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSlugPerm

`func (o *ComposerPackageUpload) GetSlugPerm() string`

GetSlugPerm returns the SlugPerm field if non-nil, zero value otherwise.

### GetSlugPermOk

`func (o *ComposerPackageUpload) GetSlugPermOk() (*string, bool)`

GetSlugPermOk returns a tuple with the SlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugPerm

`func (o *ComposerPackageUpload) SetSlugPerm(v string)`

SetSlugPerm sets SlugPerm field to given value.

### HasSlugPerm

`func (o *ComposerPackageUpload) HasSlugPerm() bool`

HasSlugPerm returns a boolean if a field has been set.

### GetStage

`func (o *ComposerPackageUpload) GetStage() int64`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *ComposerPackageUpload) GetStageOk() (*int64, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *ComposerPackageUpload) SetStage(v int64)`

SetStage sets Stage field to given value.

### HasStage

`func (o *ComposerPackageUpload) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetStageStr

`func (o *ComposerPackageUpload) GetStageStr() string`

GetStageStr returns the StageStr field if non-nil, zero value otherwise.

### GetStageStrOk

`func (o *ComposerPackageUpload) GetStageStrOk() (*string, bool)`

GetStageStrOk returns a tuple with the StageStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageStr

`func (o *ComposerPackageUpload) SetStageStr(v string)`

SetStageStr sets StageStr field to given value.

### HasStageStr

`func (o *ComposerPackageUpload) HasStageStr() bool`

HasStageStr returns a boolean if a field has been set.

### GetStageUpdatedAt

`func (o *ComposerPackageUpload) GetStageUpdatedAt() time.Time`

GetStageUpdatedAt returns the StageUpdatedAt field if non-nil, zero value otherwise.

### GetStageUpdatedAtOk

`func (o *ComposerPackageUpload) GetStageUpdatedAtOk() (*time.Time, bool)`

GetStageUpdatedAtOk returns a tuple with the StageUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageUpdatedAt

`func (o *ComposerPackageUpload) SetStageUpdatedAt(v time.Time)`

SetStageUpdatedAt sets StageUpdatedAt field to given value.

### HasStageUpdatedAt

`func (o *ComposerPackageUpload) HasStageUpdatedAt() bool`

HasStageUpdatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *ComposerPackageUpload) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ComposerPackageUpload) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ComposerPackageUpload) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ComposerPackageUpload) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusReason

`func (o *ComposerPackageUpload) GetStatusReason() string`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *ComposerPackageUpload) GetStatusReasonOk() (*string, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *ComposerPackageUpload) SetStatusReason(v string)`

SetStatusReason sets StatusReason field to given value.

### HasStatusReason

`func (o *ComposerPackageUpload) HasStatusReason() bool`

HasStatusReason returns a boolean if a field has been set.

### SetStatusReasonNil

`func (o *ComposerPackageUpload) SetStatusReasonNil(b bool)`

 SetStatusReasonNil sets the value for StatusReason to be an explicit nil

### UnsetStatusReason
`func (o *ComposerPackageUpload) UnsetStatusReason()`

UnsetStatusReason ensures that no value is present for StatusReason, not even an explicit nil
### GetStatusStr

`func (o *ComposerPackageUpload) GetStatusStr() string`

GetStatusStr returns the StatusStr field if non-nil, zero value otherwise.

### GetStatusStrOk

`func (o *ComposerPackageUpload) GetStatusStrOk() (*string, bool)`

GetStatusStrOk returns a tuple with the StatusStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusStr

`func (o *ComposerPackageUpload) SetStatusStr(v string)`

SetStatusStr sets StatusStr field to given value.

### HasStatusStr

`func (o *ComposerPackageUpload) HasStatusStr() bool`

HasStatusStr returns a boolean if a field has been set.

### GetStatusUpdatedAt

`func (o *ComposerPackageUpload) GetStatusUpdatedAt() time.Time`

GetStatusUpdatedAt returns the StatusUpdatedAt field if non-nil, zero value otherwise.

### GetStatusUpdatedAtOk

`func (o *ComposerPackageUpload) GetStatusUpdatedAtOk() (*time.Time, bool)`

GetStatusUpdatedAtOk returns a tuple with the StatusUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusUpdatedAt

`func (o *ComposerPackageUpload) SetStatusUpdatedAt(v time.Time)`

SetStatusUpdatedAt sets StatusUpdatedAt field to given value.

### HasStatusUpdatedAt

`func (o *ComposerPackageUpload) HasStatusUpdatedAt() bool`

HasStatusUpdatedAt returns a boolean if a field has been set.

### GetStatusUrl

`func (o *ComposerPackageUpload) GetStatusUrl() string`

GetStatusUrl returns the StatusUrl field if non-nil, zero value otherwise.

### GetStatusUrlOk

`func (o *ComposerPackageUpload) GetStatusUrlOk() (*string, bool)`

GetStatusUrlOk returns a tuple with the StatusUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusUrl

`func (o *ComposerPackageUpload) SetStatusUrl(v string)`

SetStatusUrl sets StatusUrl field to given value.

### HasStatusUrl

`func (o *ComposerPackageUpload) HasStatusUrl() bool`

HasStatusUrl returns a boolean if a field has been set.

### GetSubtype

`func (o *ComposerPackageUpload) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *ComposerPackageUpload) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *ComposerPackageUpload) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *ComposerPackageUpload) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetSummary

`func (o *ComposerPackageUpload) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ComposerPackageUpload) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ComposerPackageUpload) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *ComposerPackageUpload) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### SetSummaryNil

`func (o *ComposerPackageUpload) SetSummaryNil(b bool)`

 SetSummaryNil sets the value for Summary to be an explicit nil

### UnsetSummary
`func (o *ComposerPackageUpload) UnsetSummary()`

UnsetSummary ensures that no value is present for Summary, not even an explicit nil
### GetSyncFinishedAt

`func (o *ComposerPackageUpload) GetSyncFinishedAt() time.Time`

GetSyncFinishedAt returns the SyncFinishedAt field if non-nil, zero value otherwise.

### GetSyncFinishedAtOk

`func (o *ComposerPackageUpload) GetSyncFinishedAtOk() (*time.Time, bool)`

GetSyncFinishedAtOk returns a tuple with the SyncFinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncFinishedAt

`func (o *ComposerPackageUpload) SetSyncFinishedAt(v time.Time)`

SetSyncFinishedAt sets SyncFinishedAt field to given value.

### HasSyncFinishedAt

`func (o *ComposerPackageUpload) HasSyncFinishedAt() bool`

HasSyncFinishedAt returns a boolean if a field has been set.

### SetSyncFinishedAtNil

`func (o *ComposerPackageUpload) SetSyncFinishedAtNil(b bool)`

 SetSyncFinishedAtNil sets the value for SyncFinishedAt to be an explicit nil

### UnsetSyncFinishedAt
`func (o *ComposerPackageUpload) UnsetSyncFinishedAt()`

UnsetSyncFinishedAt ensures that no value is present for SyncFinishedAt, not even an explicit nil
### GetSyncProgress

`func (o *ComposerPackageUpload) GetSyncProgress() int64`

GetSyncProgress returns the SyncProgress field if non-nil, zero value otherwise.

### GetSyncProgressOk

`func (o *ComposerPackageUpload) GetSyncProgressOk() (*int64, bool)`

GetSyncProgressOk returns a tuple with the SyncProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncProgress

`func (o *ComposerPackageUpload) SetSyncProgress(v int64)`

SetSyncProgress sets SyncProgress field to given value.

### HasSyncProgress

`func (o *ComposerPackageUpload) HasSyncProgress() bool`

HasSyncProgress returns a boolean if a field has been set.

### GetTagsImmutable

`func (o *ComposerPackageUpload) GetTagsImmutable() map[string]interface{}`

GetTagsImmutable returns the TagsImmutable field if non-nil, zero value otherwise.

### GetTagsImmutableOk

`func (o *ComposerPackageUpload) GetTagsImmutableOk() (*map[string]interface{}, bool)`

GetTagsImmutableOk returns a tuple with the TagsImmutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsImmutable

`func (o *ComposerPackageUpload) SetTagsImmutable(v map[string]interface{})`

SetTagsImmutable sets TagsImmutable field to given value.

### HasTagsImmutable

`func (o *ComposerPackageUpload) HasTagsImmutable() bool`

HasTagsImmutable returns a boolean if a field has been set.

### GetTypeDisplay

`func (o *ComposerPackageUpload) GetTypeDisplay() string`

GetTypeDisplay returns the TypeDisplay field if non-nil, zero value otherwise.

### GetTypeDisplayOk

`func (o *ComposerPackageUpload) GetTypeDisplayOk() (*string, bool)`

GetTypeDisplayOk returns a tuple with the TypeDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeDisplay

`func (o *ComposerPackageUpload) SetTypeDisplay(v string)`

SetTypeDisplay sets TypeDisplay field to given value.

### HasTypeDisplay

`func (o *ComposerPackageUpload) HasTypeDisplay() bool`

HasTypeDisplay returns a boolean if a field has been set.

### GetUploadedAt

`func (o *ComposerPackageUpload) GetUploadedAt() time.Time`

GetUploadedAt returns the UploadedAt field if non-nil, zero value otherwise.

### GetUploadedAtOk

`func (o *ComposerPackageUpload) GetUploadedAtOk() (*time.Time, bool)`

GetUploadedAtOk returns a tuple with the UploadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedAt

`func (o *ComposerPackageUpload) SetUploadedAt(v time.Time)`

SetUploadedAt sets UploadedAt field to given value.

### HasUploadedAt

`func (o *ComposerPackageUpload) HasUploadedAt() bool`

HasUploadedAt returns a boolean if a field has been set.

### GetUploader

`func (o *ComposerPackageUpload) GetUploader() string`

GetUploader returns the Uploader field if non-nil, zero value otherwise.

### GetUploaderOk

`func (o *ComposerPackageUpload) GetUploaderOk() (*string, bool)`

GetUploaderOk returns a tuple with the Uploader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploader

`func (o *ComposerPackageUpload) SetUploader(v string)`

SetUploader sets Uploader field to given value.

### HasUploader

`func (o *ComposerPackageUpload) HasUploader() bool`

HasUploader returns a boolean if a field has been set.

### GetUploaderUrl

`func (o *ComposerPackageUpload) GetUploaderUrl() string`

GetUploaderUrl returns the UploaderUrl field if non-nil, zero value otherwise.

### GetUploaderUrlOk

`func (o *ComposerPackageUpload) GetUploaderUrlOk() (*string, bool)`

GetUploaderUrlOk returns a tuple with the UploaderUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploaderUrl

`func (o *ComposerPackageUpload) SetUploaderUrl(v string)`

SetUploaderUrl sets UploaderUrl field to given value.

### HasUploaderUrl

`func (o *ComposerPackageUpload) HasUploaderUrl() bool`

HasUploaderUrl returns a boolean if a field has been set.

### GetVersion

`func (o *ComposerPackageUpload) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ComposerPackageUpload) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ComposerPackageUpload) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ComposerPackageUpload) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ComposerPackageUpload) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ComposerPackageUpload) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetVersionOrig

`func (o *ComposerPackageUpload) GetVersionOrig() string`

GetVersionOrig returns the VersionOrig field if non-nil, zero value otherwise.

### GetVersionOrigOk

`func (o *ComposerPackageUpload) GetVersionOrigOk() (*string, bool)`

GetVersionOrigOk returns a tuple with the VersionOrig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionOrig

`func (o *ComposerPackageUpload) SetVersionOrig(v string)`

SetVersionOrig sets VersionOrig field to given value.

### HasVersionOrig

`func (o *ComposerPackageUpload) HasVersionOrig() bool`

HasVersionOrig returns a boolean if a field has been set.

### GetVulnerabilityScanResultsUrl

`func (o *ComposerPackageUpload) GetVulnerabilityScanResultsUrl() string`

GetVulnerabilityScanResultsUrl returns the VulnerabilityScanResultsUrl field if non-nil, zero value otherwise.

### GetVulnerabilityScanResultsUrlOk

`func (o *ComposerPackageUpload) GetVulnerabilityScanResultsUrlOk() (*string, bool)`

GetVulnerabilityScanResultsUrlOk returns a tuple with the VulnerabilityScanResultsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityScanResultsUrl

`func (o *ComposerPackageUpload) SetVulnerabilityScanResultsUrl(v string)`

SetVulnerabilityScanResultsUrl sets VulnerabilityScanResultsUrl field to given value.

### HasVulnerabilityScanResultsUrl

`func (o *ComposerPackageUpload) HasVulnerabilityScanResultsUrl() bool`

HasVulnerabilityScanResultsUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


