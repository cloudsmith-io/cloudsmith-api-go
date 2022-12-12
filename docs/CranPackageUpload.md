# CranPackageUpload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Architectures** | Pointer to [**[]Architecture**](Architecture.md) |  | [optional] [readonly] 
**CdnUrl** | Pointer to **NullableString** |  | [optional] [readonly] 
**ChecksumMd5** | Pointer to **string** |  | [optional] [readonly] 
**ChecksumSha1** | Pointer to **string** |  | [optional] [readonly] 
**ChecksumSha256** | Pointer to **string** |  | [optional] [readonly] 
**ChecksumSha512** | Pointer to **string** |  | [optional] [readonly] 
**DependenciesChecksumMd5** | Pointer to **string** | A checksum of all of the package&#39;s dependencies. | [optional] [readonly] 
**DependenciesUrl** | Pointer to **string** |  | [optional] [readonly] 
**Description** | Pointer to **string** | A textual description of this package. | [optional] [readonly] 
**Distro** | Pointer to [**NullableDistribution**](Distribution.md) |  | [optional] 
**DistroVersion** | Pointer to [**DistributionVersion**](DistributionVersion.md) |  | [optional] 
**Downloads** | Pointer to **int64** |  | [optional] [readonly] 
**Epoch** | Pointer to **int64** | The epoch of the package version (if any). | [optional] [readonly] 
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
**License** | Pointer to **string** | The license of this package. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of this package. | [optional] [readonly] 
**Namespace** | Pointer to **string** |  | [optional] [readonly] 
**NamespaceUrl** | Pointer to **string** |  | [optional] [readonly] 
**NumFiles** | Pointer to **int64** |  | [optional] [readonly] 
**OriginRepository** | Pointer to **string** |  | [optional] [readonly] 
**OriginRepositoryUrl** | Pointer to **string** |  | [optional] [readonly] 
**PackageType** | Pointer to **int64** | The type of package contents. | [optional] [readonly] 
**Release** | Pointer to **string** | The release of the package version (if any). | [optional] [readonly] 
**Repository** | Pointer to **string** |  | [optional] [readonly] 
**RepositoryUrl** | Pointer to **string** |  | [optional] [readonly] 
**SecurityScanCompletedAt** | Pointer to **time.Time** | The datetime the security scanning was completed. | [optional] [readonly] 
**SecurityScanStartedAt** | Pointer to **time.Time** | The datetime the security scanning was started. | [optional] [readonly] 
**SecurityScanStatus** | Pointer to **NullableString** |  | [optional] [readonly] [default to "Awaiting Security Scan"]
**SecurityScanStatusUpdatedAt** | Pointer to **time.Time** | The datetime the security scanning status was updated. | [optional] [readonly] 
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
**StatusReason** | Pointer to **string** | A textual description for the synchronous status reason (if any | [optional] [readonly] 
**StatusStr** | Pointer to **string** |  | [optional] [readonly] 
**StatusUpdatedAt** | Pointer to **time.Time** | The datetime the package status was updated at. | [optional] [readonly] 
**StatusUrl** | Pointer to **string** |  | [optional] [readonly] 
**Subtype** | Pointer to **string** |  | [optional] [readonly] 
**Summary** | Pointer to **string** | A one-liner synopsis of this package. | [optional] [readonly] 
**SyncFinishedAt** | Pointer to **time.Time** | The datetime the package sync was finished at. | [optional] [readonly] 
**SyncProgress** | Pointer to **int64** | Synchronisation progress (from 0-100) | [optional] [readonly] 
**TagsImmutable** | Pointer to **map[string]interface{}** | All tags on the package, grouped by tag type. This includes immutable tags, but doesn&#39;t distinguish them from mutable. To see which tags are immutable specifically, see the tags_immutable field. | [optional] 
**TypeDisplay** | Pointer to **string** |  | [optional] [readonly] 
**UploadedAt** | Pointer to **time.Time** | The date this package was uploaded. | [optional] [readonly] 
**Uploader** | Pointer to **string** |  | [optional] [readonly] 
**UploaderUrl** | Pointer to **string** |  | [optional] [readonly] 
**Version** | Pointer to **string** | The raw version for this package. | [optional] [readonly] 
**VersionOrig** | Pointer to **string** |  | [optional] [readonly] 
**VulnerabilityScanResultsUrl** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewCranPackageUpload

`func NewCranPackageUpload() *CranPackageUpload`

NewCranPackageUpload instantiates a new CranPackageUpload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCranPackageUploadWithDefaults

`func NewCranPackageUploadWithDefaults() *CranPackageUpload`

NewCranPackageUploadWithDefaults instantiates a new CranPackageUpload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchitectures

`func (o *CranPackageUpload) GetArchitectures() []Architecture`

GetArchitectures returns the Architectures field if non-nil, zero value otherwise.

### GetArchitecturesOk

`func (o *CranPackageUpload) GetArchitecturesOk() (*[]Architecture, bool)`

GetArchitecturesOk returns a tuple with the Architectures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitectures

`func (o *CranPackageUpload) SetArchitectures(v []Architecture)`

SetArchitectures sets Architectures field to given value.

### HasArchitectures

`func (o *CranPackageUpload) HasArchitectures() bool`

HasArchitectures returns a boolean if a field has been set.

### GetCdnUrl

`func (o *CranPackageUpload) GetCdnUrl() string`

GetCdnUrl returns the CdnUrl field if non-nil, zero value otherwise.

### GetCdnUrlOk

`func (o *CranPackageUpload) GetCdnUrlOk() (*string, bool)`

GetCdnUrlOk returns a tuple with the CdnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnUrl

`func (o *CranPackageUpload) SetCdnUrl(v string)`

SetCdnUrl sets CdnUrl field to given value.

### HasCdnUrl

`func (o *CranPackageUpload) HasCdnUrl() bool`

HasCdnUrl returns a boolean if a field has been set.

### SetCdnUrlNil

`func (o *CranPackageUpload) SetCdnUrlNil(b bool)`

 SetCdnUrlNil sets the value for CdnUrl to be an explicit nil

### UnsetCdnUrl
`func (o *CranPackageUpload) UnsetCdnUrl()`

UnsetCdnUrl ensures that no value is present for CdnUrl, not even an explicit nil
### GetChecksumMd5

`func (o *CranPackageUpload) GetChecksumMd5() string`

GetChecksumMd5 returns the ChecksumMd5 field if non-nil, zero value otherwise.

### GetChecksumMd5Ok

`func (o *CranPackageUpload) GetChecksumMd5Ok() (*string, bool)`

GetChecksumMd5Ok returns a tuple with the ChecksumMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumMd5

`func (o *CranPackageUpload) SetChecksumMd5(v string)`

SetChecksumMd5 sets ChecksumMd5 field to given value.

### HasChecksumMd5

`func (o *CranPackageUpload) HasChecksumMd5() bool`

HasChecksumMd5 returns a boolean if a field has been set.

### GetChecksumSha1

`func (o *CranPackageUpload) GetChecksumSha1() string`

GetChecksumSha1 returns the ChecksumSha1 field if non-nil, zero value otherwise.

### GetChecksumSha1Ok

`func (o *CranPackageUpload) GetChecksumSha1Ok() (*string, bool)`

GetChecksumSha1Ok returns a tuple with the ChecksumSha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumSha1

`func (o *CranPackageUpload) SetChecksumSha1(v string)`

SetChecksumSha1 sets ChecksumSha1 field to given value.

### HasChecksumSha1

`func (o *CranPackageUpload) HasChecksumSha1() bool`

HasChecksumSha1 returns a boolean if a field has been set.

### GetChecksumSha256

`func (o *CranPackageUpload) GetChecksumSha256() string`

GetChecksumSha256 returns the ChecksumSha256 field if non-nil, zero value otherwise.

### GetChecksumSha256Ok

`func (o *CranPackageUpload) GetChecksumSha256Ok() (*string, bool)`

GetChecksumSha256Ok returns a tuple with the ChecksumSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumSha256

`func (o *CranPackageUpload) SetChecksumSha256(v string)`

SetChecksumSha256 sets ChecksumSha256 field to given value.

### HasChecksumSha256

`func (o *CranPackageUpload) HasChecksumSha256() bool`

HasChecksumSha256 returns a boolean if a field has been set.

### GetChecksumSha512

`func (o *CranPackageUpload) GetChecksumSha512() string`

GetChecksumSha512 returns the ChecksumSha512 field if non-nil, zero value otherwise.

### GetChecksumSha512Ok

`func (o *CranPackageUpload) GetChecksumSha512Ok() (*string, bool)`

GetChecksumSha512Ok returns a tuple with the ChecksumSha512 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumSha512

`func (o *CranPackageUpload) SetChecksumSha512(v string)`

SetChecksumSha512 sets ChecksumSha512 field to given value.

### HasChecksumSha512

`func (o *CranPackageUpload) HasChecksumSha512() bool`

HasChecksumSha512 returns a boolean if a field has been set.

### GetDependenciesChecksumMd5

`func (o *CranPackageUpload) GetDependenciesChecksumMd5() string`

GetDependenciesChecksumMd5 returns the DependenciesChecksumMd5 field if non-nil, zero value otherwise.

### GetDependenciesChecksumMd5Ok

`func (o *CranPackageUpload) GetDependenciesChecksumMd5Ok() (*string, bool)`

GetDependenciesChecksumMd5Ok returns a tuple with the DependenciesChecksumMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependenciesChecksumMd5

`func (o *CranPackageUpload) SetDependenciesChecksumMd5(v string)`

SetDependenciesChecksumMd5 sets DependenciesChecksumMd5 field to given value.

### HasDependenciesChecksumMd5

`func (o *CranPackageUpload) HasDependenciesChecksumMd5() bool`

HasDependenciesChecksumMd5 returns a boolean if a field has been set.

### GetDependenciesUrl

`func (o *CranPackageUpload) GetDependenciesUrl() string`

GetDependenciesUrl returns the DependenciesUrl field if non-nil, zero value otherwise.

### GetDependenciesUrlOk

`func (o *CranPackageUpload) GetDependenciesUrlOk() (*string, bool)`

GetDependenciesUrlOk returns a tuple with the DependenciesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependenciesUrl

`func (o *CranPackageUpload) SetDependenciesUrl(v string)`

SetDependenciesUrl sets DependenciesUrl field to given value.

### HasDependenciesUrl

`func (o *CranPackageUpload) HasDependenciesUrl() bool`

HasDependenciesUrl returns a boolean if a field has been set.

### GetDescription

`func (o *CranPackageUpload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CranPackageUpload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CranPackageUpload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CranPackageUpload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDistro

`func (o *CranPackageUpload) GetDistro() Distribution`

GetDistro returns the Distro field if non-nil, zero value otherwise.

### GetDistroOk

`func (o *CranPackageUpload) GetDistroOk() (*Distribution, bool)`

GetDistroOk returns a tuple with the Distro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistro

`func (o *CranPackageUpload) SetDistro(v Distribution)`

SetDistro sets Distro field to given value.

### HasDistro

`func (o *CranPackageUpload) HasDistro() bool`

HasDistro returns a boolean if a field has been set.

### SetDistroNil

`func (o *CranPackageUpload) SetDistroNil(b bool)`

 SetDistroNil sets the value for Distro to be an explicit nil

### UnsetDistro
`func (o *CranPackageUpload) UnsetDistro()`

UnsetDistro ensures that no value is present for Distro, not even an explicit nil
### GetDistroVersion

`func (o *CranPackageUpload) GetDistroVersion() DistributionVersion`

GetDistroVersion returns the DistroVersion field if non-nil, zero value otherwise.

### GetDistroVersionOk

`func (o *CranPackageUpload) GetDistroVersionOk() (*DistributionVersion, bool)`

GetDistroVersionOk returns a tuple with the DistroVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistroVersion

`func (o *CranPackageUpload) SetDistroVersion(v DistributionVersion)`

SetDistroVersion sets DistroVersion field to given value.

### HasDistroVersion

`func (o *CranPackageUpload) HasDistroVersion() bool`

HasDistroVersion returns a boolean if a field has been set.

### GetDownloads

`func (o *CranPackageUpload) GetDownloads() int64`

GetDownloads returns the Downloads field if non-nil, zero value otherwise.

### GetDownloadsOk

`func (o *CranPackageUpload) GetDownloadsOk() (*int64, bool)`

GetDownloadsOk returns a tuple with the Downloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloads

`func (o *CranPackageUpload) SetDownloads(v int64)`

SetDownloads sets Downloads field to given value.

### HasDownloads

`func (o *CranPackageUpload) HasDownloads() bool`

HasDownloads returns a boolean if a field has been set.

### GetEpoch

`func (o *CranPackageUpload) GetEpoch() int64`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *CranPackageUpload) GetEpochOk() (*int64, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *CranPackageUpload) SetEpoch(v int64)`

SetEpoch sets Epoch field to given value.

### HasEpoch

`func (o *CranPackageUpload) HasEpoch() bool`

HasEpoch returns a boolean if a field has been set.

### GetExtension

`func (o *CranPackageUpload) GetExtension() string`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *CranPackageUpload) GetExtensionOk() (*string, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *CranPackageUpload) SetExtension(v string)`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *CranPackageUpload) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### GetFilename

`func (o *CranPackageUpload) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *CranPackageUpload) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *CranPackageUpload) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *CranPackageUpload) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetFiles

`func (o *CranPackageUpload) GetFiles() []PackageFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *CranPackageUpload) GetFilesOk() (*[]PackageFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *CranPackageUpload) SetFiles(v []PackageFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *CranPackageUpload) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetFormat

`func (o *CranPackageUpload) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CranPackageUpload) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CranPackageUpload) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CranPackageUpload) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetFormatUrl

`func (o *CranPackageUpload) GetFormatUrl() string`

GetFormatUrl returns the FormatUrl field if non-nil, zero value otherwise.

### GetFormatUrlOk

`func (o *CranPackageUpload) GetFormatUrlOk() (*string, bool)`

GetFormatUrlOk returns a tuple with the FormatUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatUrl

`func (o *CranPackageUpload) SetFormatUrl(v string)`

SetFormatUrl sets FormatUrl field to given value.

### HasFormatUrl

`func (o *CranPackageUpload) HasFormatUrl() bool`

HasFormatUrl returns a boolean if a field has been set.

### GetIdentifierPerm

`func (o *CranPackageUpload) GetIdentifierPerm() string`

GetIdentifierPerm returns the IdentifierPerm field if non-nil, zero value otherwise.

### GetIdentifierPermOk

`func (o *CranPackageUpload) GetIdentifierPermOk() (*string, bool)`

GetIdentifierPermOk returns a tuple with the IdentifierPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierPerm

`func (o *CranPackageUpload) SetIdentifierPerm(v string)`

SetIdentifierPerm sets IdentifierPerm field to given value.

### HasIdentifierPerm

`func (o *CranPackageUpload) HasIdentifierPerm() bool`

HasIdentifierPerm returns a boolean if a field has been set.

### GetIndexed

`func (o *CranPackageUpload) GetIndexed() bool`

GetIndexed returns the Indexed field if non-nil, zero value otherwise.

### GetIndexedOk

`func (o *CranPackageUpload) GetIndexedOk() (*bool, bool)`

GetIndexedOk returns a tuple with the Indexed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexed

`func (o *CranPackageUpload) SetIndexed(v bool)`

SetIndexed sets Indexed field to given value.

### HasIndexed

`func (o *CranPackageUpload) HasIndexed() bool`

HasIndexed returns a boolean if a field has been set.

### GetIsDownloadable

`func (o *CranPackageUpload) GetIsDownloadable() bool`

GetIsDownloadable returns the IsDownloadable field if non-nil, zero value otherwise.

### GetIsDownloadableOk

`func (o *CranPackageUpload) GetIsDownloadableOk() (*bool, bool)`

GetIsDownloadableOk returns a tuple with the IsDownloadable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDownloadable

`func (o *CranPackageUpload) SetIsDownloadable(v bool)`

SetIsDownloadable sets IsDownloadable field to given value.

### HasIsDownloadable

`func (o *CranPackageUpload) HasIsDownloadable() bool`

HasIsDownloadable returns a boolean if a field has been set.

### GetIsQuarantined

`func (o *CranPackageUpload) GetIsQuarantined() bool`

GetIsQuarantined returns the IsQuarantined field if non-nil, zero value otherwise.

### GetIsQuarantinedOk

`func (o *CranPackageUpload) GetIsQuarantinedOk() (*bool, bool)`

GetIsQuarantinedOk returns a tuple with the IsQuarantined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQuarantined

`func (o *CranPackageUpload) SetIsQuarantined(v bool)`

SetIsQuarantined sets IsQuarantined field to given value.

### HasIsQuarantined

`func (o *CranPackageUpload) HasIsQuarantined() bool`

HasIsQuarantined returns a boolean if a field has been set.

### GetIsSyncAwaiting

`func (o *CranPackageUpload) GetIsSyncAwaiting() bool`

GetIsSyncAwaiting returns the IsSyncAwaiting field if non-nil, zero value otherwise.

### GetIsSyncAwaitingOk

`func (o *CranPackageUpload) GetIsSyncAwaitingOk() (*bool, bool)`

GetIsSyncAwaitingOk returns a tuple with the IsSyncAwaiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncAwaiting

`func (o *CranPackageUpload) SetIsSyncAwaiting(v bool)`

SetIsSyncAwaiting sets IsSyncAwaiting field to given value.

### HasIsSyncAwaiting

`func (o *CranPackageUpload) HasIsSyncAwaiting() bool`

HasIsSyncAwaiting returns a boolean if a field has been set.

### GetIsSyncCompleted

`func (o *CranPackageUpload) GetIsSyncCompleted() bool`

GetIsSyncCompleted returns the IsSyncCompleted field if non-nil, zero value otherwise.

### GetIsSyncCompletedOk

`func (o *CranPackageUpload) GetIsSyncCompletedOk() (*bool, bool)`

GetIsSyncCompletedOk returns a tuple with the IsSyncCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncCompleted

`func (o *CranPackageUpload) SetIsSyncCompleted(v bool)`

SetIsSyncCompleted sets IsSyncCompleted field to given value.

### HasIsSyncCompleted

`func (o *CranPackageUpload) HasIsSyncCompleted() bool`

HasIsSyncCompleted returns a boolean if a field has been set.

### GetIsSyncFailed

`func (o *CranPackageUpload) GetIsSyncFailed() bool`

GetIsSyncFailed returns the IsSyncFailed field if non-nil, zero value otherwise.

### GetIsSyncFailedOk

`func (o *CranPackageUpload) GetIsSyncFailedOk() (*bool, bool)`

GetIsSyncFailedOk returns a tuple with the IsSyncFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncFailed

`func (o *CranPackageUpload) SetIsSyncFailed(v bool)`

SetIsSyncFailed sets IsSyncFailed field to given value.

### HasIsSyncFailed

`func (o *CranPackageUpload) HasIsSyncFailed() bool`

HasIsSyncFailed returns a boolean if a field has been set.

### GetIsSyncInFlight

`func (o *CranPackageUpload) GetIsSyncInFlight() bool`

GetIsSyncInFlight returns the IsSyncInFlight field if non-nil, zero value otherwise.

### GetIsSyncInFlightOk

`func (o *CranPackageUpload) GetIsSyncInFlightOk() (*bool, bool)`

GetIsSyncInFlightOk returns a tuple with the IsSyncInFlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncInFlight

`func (o *CranPackageUpload) SetIsSyncInFlight(v bool)`

SetIsSyncInFlight sets IsSyncInFlight field to given value.

### HasIsSyncInFlight

`func (o *CranPackageUpload) HasIsSyncInFlight() bool`

HasIsSyncInFlight returns a boolean if a field has been set.

### GetIsSyncInProgress

`func (o *CranPackageUpload) GetIsSyncInProgress() bool`

GetIsSyncInProgress returns the IsSyncInProgress field if non-nil, zero value otherwise.

### GetIsSyncInProgressOk

`func (o *CranPackageUpload) GetIsSyncInProgressOk() (*bool, bool)`

GetIsSyncInProgressOk returns a tuple with the IsSyncInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncInProgress

`func (o *CranPackageUpload) SetIsSyncInProgress(v bool)`

SetIsSyncInProgress sets IsSyncInProgress field to given value.

### HasIsSyncInProgress

`func (o *CranPackageUpload) HasIsSyncInProgress() bool`

HasIsSyncInProgress returns a boolean if a field has been set.

### GetLicense

`func (o *CranPackageUpload) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *CranPackageUpload) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *CranPackageUpload) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *CranPackageUpload) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetName

`func (o *CranPackageUpload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CranPackageUpload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CranPackageUpload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CranPackageUpload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *CranPackageUpload) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CranPackageUpload) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CranPackageUpload) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *CranPackageUpload) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetNamespaceUrl

`func (o *CranPackageUpload) GetNamespaceUrl() string`

GetNamespaceUrl returns the NamespaceUrl field if non-nil, zero value otherwise.

### GetNamespaceUrlOk

`func (o *CranPackageUpload) GetNamespaceUrlOk() (*string, bool)`

GetNamespaceUrlOk returns a tuple with the NamespaceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceUrl

`func (o *CranPackageUpload) SetNamespaceUrl(v string)`

SetNamespaceUrl sets NamespaceUrl field to given value.

### HasNamespaceUrl

`func (o *CranPackageUpload) HasNamespaceUrl() bool`

HasNamespaceUrl returns a boolean if a field has been set.

### GetNumFiles

`func (o *CranPackageUpload) GetNumFiles() int64`

GetNumFiles returns the NumFiles field if non-nil, zero value otherwise.

### GetNumFilesOk

`func (o *CranPackageUpload) GetNumFilesOk() (*int64, bool)`

GetNumFilesOk returns a tuple with the NumFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFiles

`func (o *CranPackageUpload) SetNumFiles(v int64)`

SetNumFiles sets NumFiles field to given value.

### HasNumFiles

`func (o *CranPackageUpload) HasNumFiles() bool`

HasNumFiles returns a boolean if a field has been set.

### GetOriginRepository

`func (o *CranPackageUpload) GetOriginRepository() string`

GetOriginRepository returns the OriginRepository field if non-nil, zero value otherwise.

### GetOriginRepositoryOk

`func (o *CranPackageUpload) GetOriginRepositoryOk() (*string, bool)`

GetOriginRepositoryOk returns a tuple with the OriginRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginRepository

`func (o *CranPackageUpload) SetOriginRepository(v string)`

SetOriginRepository sets OriginRepository field to given value.

### HasOriginRepository

`func (o *CranPackageUpload) HasOriginRepository() bool`

HasOriginRepository returns a boolean if a field has been set.

### GetOriginRepositoryUrl

`func (o *CranPackageUpload) GetOriginRepositoryUrl() string`

GetOriginRepositoryUrl returns the OriginRepositoryUrl field if non-nil, zero value otherwise.

### GetOriginRepositoryUrlOk

`func (o *CranPackageUpload) GetOriginRepositoryUrlOk() (*string, bool)`

GetOriginRepositoryUrlOk returns a tuple with the OriginRepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginRepositoryUrl

`func (o *CranPackageUpload) SetOriginRepositoryUrl(v string)`

SetOriginRepositoryUrl sets OriginRepositoryUrl field to given value.

### HasOriginRepositoryUrl

`func (o *CranPackageUpload) HasOriginRepositoryUrl() bool`

HasOriginRepositoryUrl returns a boolean if a field has been set.

### GetPackageType

`func (o *CranPackageUpload) GetPackageType() int64`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *CranPackageUpload) GetPackageTypeOk() (*int64, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *CranPackageUpload) SetPackageType(v int64)`

SetPackageType sets PackageType field to given value.

### HasPackageType

`func (o *CranPackageUpload) HasPackageType() bool`

HasPackageType returns a boolean if a field has been set.

### GetRelease

`func (o *CranPackageUpload) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *CranPackageUpload) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *CranPackageUpload) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *CranPackageUpload) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetRepository

`func (o *CranPackageUpload) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *CranPackageUpload) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *CranPackageUpload) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *CranPackageUpload) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetRepositoryUrl

`func (o *CranPackageUpload) GetRepositoryUrl() string`

GetRepositoryUrl returns the RepositoryUrl field if non-nil, zero value otherwise.

### GetRepositoryUrlOk

`func (o *CranPackageUpload) GetRepositoryUrlOk() (*string, bool)`

GetRepositoryUrlOk returns a tuple with the RepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryUrl

`func (o *CranPackageUpload) SetRepositoryUrl(v string)`

SetRepositoryUrl sets RepositoryUrl field to given value.

### HasRepositoryUrl

`func (o *CranPackageUpload) HasRepositoryUrl() bool`

HasRepositoryUrl returns a boolean if a field has been set.

### GetSecurityScanCompletedAt

`func (o *CranPackageUpload) GetSecurityScanCompletedAt() time.Time`

GetSecurityScanCompletedAt returns the SecurityScanCompletedAt field if non-nil, zero value otherwise.

### GetSecurityScanCompletedAtOk

`func (o *CranPackageUpload) GetSecurityScanCompletedAtOk() (*time.Time, bool)`

GetSecurityScanCompletedAtOk returns a tuple with the SecurityScanCompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityScanCompletedAt

`func (o *CranPackageUpload) SetSecurityScanCompletedAt(v time.Time)`

SetSecurityScanCompletedAt sets SecurityScanCompletedAt field to given value.

### HasSecurityScanCompletedAt

`func (o *CranPackageUpload) HasSecurityScanCompletedAt() bool`

HasSecurityScanCompletedAt returns a boolean if a field has been set.

### GetSecurityScanStartedAt

`func (o *CranPackageUpload) GetSecurityScanStartedAt() time.Time`

GetSecurityScanStartedAt returns the SecurityScanStartedAt field if non-nil, zero value otherwise.

### GetSecurityScanStartedAtOk

`func (o *CranPackageUpload) GetSecurityScanStartedAtOk() (*time.Time, bool)`

GetSecurityScanStartedAtOk returns a tuple with the SecurityScanStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityScanStartedAt

`func (o *CranPackageUpload) SetSecurityScanStartedAt(v time.Time)`

SetSecurityScanStartedAt sets SecurityScanStartedAt field to given value.

### HasSecurityScanStartedAt

`func (o *CranPackageUpload) HasSecurityScanStartedAt() bool`

HasSecurityScanStartedAt returns a boolean if a field has been set.

### GetSecurityScanStatus

`func (o *CranPackageUpload) GetSecurityScanStatus() string`

GetSecurityScanStatus returns the SecurityScanStatus field if non-nil, zero value otherwise.

### GetSecurityScanStatusOk

`func (o *CranPackageUpload) GetSecurityScanStatusOk() (*string, bool)`

GetSecurityScanStatusOk returns a tuple with the SecurityScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityScanStatus

`func (o *CranPackageUpload) SetSecurityScanStatus(v string)`

SetSecurityScanStatus sets SecurityScanStatus field to given value.

### HasSecurityScanStatus

`func (o *CranPackageUpload) HasSecurityScanStatus() bool`

HasSecurityScanStatus returns a boolean if a field has been set.

### SetSecurityScanStatusNil

`func (o *CranPackageUpload) SetSecurityScanStatusNil(b bool)`

 SetSecurityScanStatusNil sets the value for SecurityScanStatus to be an explicit nil

### UnsetSecurityScanStatus
`func (o *CranPackageUpload) UnsetSecurityScanStatus()`

UnsetSecurityScanStatus ensures that no value is present for SecurityScanStatus, not even an explicit nil
### GetSecurityScanStatusUpdatedAt

`func (o *CranPackageUpload) GetSecurityScanStatusUpdatedAt() time.Time`

GetSecurityScanStatusUpdatedAt returns the SecurityScanStatusUpdatedAt field if non-nil, zero value otherwise.

### GetSecurityScanStatusUpdatedAtOk

`func (o *CranPackageUpload) GetSecurityScanStatusUpdatedAtOk() (*time.Time, bool)`

GetSecurityScanStatusUpdatedAtOk returns a tuple with the SecurityScanStatusUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityScanStatusUpdatedAt

`func (o *CranPackageUpload) SetSecurityScanStatusUpdatedAt(v time.Time)`

SetSecurityScanStatusUpdatedAt sets SecurityScanStatusUpdatedAt field to given value.

### HasSecurityScanStatusUpdatedAt

`func (o *CranPackageUpload) HasSecurityScanStatusUpdatedAt() bool`

HasSecurityScanStatusUpdatedAt returns a boolean if a field has been set.

### GetSelfHtmlUrl

`func (o *CranPackageUpload) GetSelfHtmlUrl() string`

GetSelfHtmlUrl returns the SelfHtmlUrl field if non-nil, zero value otherwise.

### GetSelfHtmlUrlOk

`func (o *CranPackageUpload) GetSelfHtmlUrlOk() (*string, bool)`

GetSelfHtmlUrlOk returns a tuple with the SelfHtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfHtmlUrl

`func (o *CranPackageUpload) SetSelfHtmlUrl(v string)`

SetSelfHtmlUrl sets SelfHtmlUrl field to given value.

### HasSelfHtmlUrl

`func (o *CranPackageUpload) HasSelfHtmlUrl() bool`

HasSelfHtmlUrl returns a boolean if a field has been set.

### GetSelfUrl

`func (o *CranPackageUpload) GetSelfUrl() string`

GetSelfUrl returns the SelfUrl field if non-nil, zero value otherwise.

### GetSelfUrlOk

`func (o *CranPackageUpload) GetSelfUrlOk() (*string, bool)`

GetSelfUrlOk returns a tuple with the SelfUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfUrl

`func (o *CranPackageUpload) SetSelfUrl(v string)`

SetSelfUrl sets SelfUrl field to given value.

### HasSelfUrl

`func (o *CranPackageUpload) HasSelfUrl() bool`

HasSelfUrl returns a boolean if a field has been set.

### GetSignatureUrl

`func (o *CranPackageUpload) GetSignatureUrl() string`

GetSignatureUrl returns the SignatureUrl field if non-nil, zero value otherwise.

### GetSignatureUrlOk

`func (o *CranPackageUpload) GetSignatureUrlOk() (*string, bool)`

GetSignatureUrlOk returns a tuple with the SignatureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureUrl

`func (o *CranPackageUpload) SetSignatureUrl(v string)`

SetSignatureUrl sets SignatureUrl field to given value.

### HasSignatureUrl

`func (o *CranPackageUpload) HasSignatureUrl() bool`

HasSignatureUrl returns a boolean if a field has been set.

### SetSignatureUrlNil

`func (o *CranPackageUpload) SetSignatureUrlNil(b bool)`

 SetSignatureUrlNil sets the value for SignatureUrl to be an explicit nil

### UnsetSignatureUrl
`func (o *CranPackageUpload) UnsetSignatureUrl()`

UnsetSignatureUrl ensures that no value is present for SignatureUrl, not even an explicit nil
### GetSize

`func (o *CranPackageUpload) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CranPackageUpload) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CranPackageUpload) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *CranPackageUpload) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSlug

`func (o *CranPackageUpload) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *CranPackageUpload) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *CranPackageUpload) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *CranPackageUpload) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSlugPerm

`func (o *CranPackageUpload) GetSlugPerm() string`

GetSlugPerm returns the SlugPerm field if non-nil, zero value otherwise.

### GetSlugPermOk

`func (o *CranPackageUpload) GetSlugPermOk() (*string, bool)`

GetSlugPermOk returns a tuple with the SlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugPerm

`func (o *CranPackageUpload) SetSlugPerm(v string)`

SetSlugPerm sets SlugPerm field to given value.

### HasSlugPerm

`func (o *CranPackageUpload) HasSlugPerm() bool`

HasSlugPerm returns a boolean if a field has been set.

### GetStage

`func (o *CranPackageUpload) GetStage() int64`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *CranPackageUpload) GetStageOk() (*int64, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *CranPackageUpload) SetStage(v int64)`

SetStage sets Stage field to given value.

### HasStage

`func (o *CranPackageUpload) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetStageStr

`func (o *CranPackageUpload) GetStageStr() string`

GetStageStr returns the StageStr field if non-nil, zero value otherwise.

### GetStageStrOk

`func (o *CranPackageUpload) GetStageStrOk() (*string, bool)`

GetStageStrOk returns a tuple with the StageStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageStr

`func (o *CranPackageUpload) SetStageStr(v string)`

SetStageStr sets StageStr field to given value.

### HasStageStr

`func (o *CranPackageUpload) HasStageStr() bool`

HasStageStr returns a boolean if a field has been set.

### GetStageUpdatedAt

`func (o *CranPackageUpload) GetStageUpdatedAt() time.Time`

GetStageUpdatedAt returns the StageUpdatedAt field if non-nil, zero value otherwise.

### GetStageUpdatedAtOk

`func (o *CranPackageUpload) GetStageUpdatedAtOk() (*time.Time, bool)`

GetStageUpdatedAtOk returns a tuple with the StageUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageUpdatedAt

`func (o *CranPackageUpload) SetStageUpdatedAt(v time.Time)`

SetStageUpdatedAt sets StageUpdatedAt field to given value.

### HasStageUpdatedAt

`func (o *CranPackageUpload) HasStageUpdatedAt() bool`

HasStageUpdatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *CranPackageUpload) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CranPackageUpload) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CranPackageUpload) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CranPackageUpload) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusReason

`func (o *CranPackageUpload) GetStatusReason() string`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *CranPackageUpload) GetStatusReasonOk() (*string, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *CranPackageUpload) SetStatusReason(v string)`

SetStatusReason sets StatusReason field to given value.

### HasStatusReason

`func (o *CranPackageUpload) HasStatusReason() bool`

HasStatusReason returns a boolean if a field has been set.

### GetStatusStr

`func (o *CranPackageUpload) GetStatusStr() string`

GetStatusStr returns the StatusStr field if non-nil, zero value otherwise.

### GetStatusStrOk

`func (o *CranPackageUpload) GetStatusStrOk() (*string, bool)`

GetStatusStrOk returns a tuple with the StatusStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusStr

`func (o *CranPackageUpload) SetStatusStr(v string)`

SetStatusStr sets StatusStr field to given value.

### HasStatusStr

`func (o *CranPackageUpload) HasStatusStr() bool`

HasStatusStr returns a boolean if a field has been set.

### GetStatusUpdatedAt

`func (o *CranPackageUpload) GetStatusUpdatedAt() time.Time`

GetStatusUpdatedAt returns the StatusUpdatedAt field if non-nil, zero value otherwise.

### GetStatusUpdatedAtOk

`func (o *CranPackageUpload) GetStatusUpdatedAtOk() (*time.Time, bool)`

GetStatusUpdatedAtOk returns a tuple with the StatusUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusUpdatedAt

`func (o *CranPackageUpload) SetStatusUpdatedAt(v time.Time)`

SetStatusUpdatedAt sets StatusUpdatedAt field to given value.

### HasStatusUpdatedAt

`func (o *CranPackageUpload) HasStatusUpdatedAt() bool`

HasStatusUpdatedAt returns a boolean if a field has been set.

### GetStatusUrl

`func (o *CranPackageUpload) GetStatusUrl() string`

GetStatusUrl returns the StatusUrl field if non-nil, zero value otherwise.

### GetStatusUrlOk

`func (o *CranPackageUpload) GetStatusUrlOk() (*string, bool)`

GetStatusUrlOk returns a tuple with the StatusUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusUrl

`func (o *CranPackageUpload) SetStatusUrl(v string)`

SetStatusUrl sets StatusUrl field to given value.

### HasStatusUrl

`func (o *CranPackageUpload) HasStatusUrl() bool`

HasStatusUrl returns a boolean if a field has been set.

### GetSubtype

`func (o *CranPackageUpload) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *CranPackageUpload) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *CranPackageUpload) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *CranPackageUpload) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetSummary

`func (o *CranPackageUpload) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *CranPackageUpload) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *CranPackageUpload) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *CranPackageUpload) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetSyncFinishedAt

`func (o *CranPackageUpload) GetSyncFinishedAt() time.Time`

GetSyncFinishedAt returns the SyncFinishedAt field if non-nil, zero value otherwise.

### GetSyncFinishedAtOk

`func (o *CranPackageUpload) GetSyncFinishedAtOk() (*time.Time, bool)`

GetSyncFinishedAtOk returns a tuple with the SyncFinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncFinishedAt

`func (o *CranPackageUpload) SetSyncFinishedAt(v time.Time)`

SetSyncFinishedAt sets SyncFinishedAt field to given value.

### HasSyncFinishedAt

`func (o *CranPackageUpload) HasSyncFinishedAt() bool`

HasSyncFinishedAt returns a boolean if a field has been set.

### GetSyncProgress

`func (o *CranPackageUpload) GetSyncProgress() int64`

GetSyncProgress returns the SyncProgress field if non-nil, zero value otherwise.

### GetSyncProgressOk

`func (o *CranPackageUpload) GetSyncProgressOk() (*int64, bool)`

GetSyncProgressOk returns a tuple with the SyncProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncProgress

`func (o *CranPackageUpload) SetSyncProgress(v int64)`

SetSyncProgress sets SyncProgress field to given value.

### HasSyncProgress

`func (o *CranPackageUpload) HasSyncProgress() bool`

HasSyncProgress returns a boolean if a field has been set.

### GetTagsImmutable

`func (o *CranPackageUpload) GetTagsImmutable() map[string]interface{}`

GetTagsImmutable returns the TagsImmutable field if non-nil, zero value otherwise.

### GetTagsImmutableOk

`func (o *CranPackageUpload) GetTagsImmutableOk() (*map[string]interface{}, bool)`

GetTagsImmutableOk returns a tuple with the TagsImmutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsImmutable

`func (o *CranPackageUpload) SetTagsImmutable(v map[string]interface{})`

SetTagsImmutable sets TagsImmutable field to given value.

### HasTagsImmutable

`func (o *CranPackageUpload) HasTagsImmutable() bool`

HasTagsImmutable returns a boolean if a field has been set.

### GetTypeDisplay

`func (o *CranPackageUpload) GetTypeDisplay() string`

GetTypeDisplay returns the TypeDisplay field if non-nil, zero value otherwise.

### GetTypeDisplayOk

`func (o *CranPackageUpload) GetTypeDisplayOk() (*string, bool)`

GetTypeDisplayOk returns a tuple with the TypeDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeDisplay

`func (o *CranPackageUpload) SetTypeDisplay(v string)`

SetTypeDisplay sets TypeDisplay field to given value.

### HasTypeDisplay

`func (o *CranPackageUpload) HasTypeDisplay() bool`

HasTypeDisplay returns a boolean if a field has been set.

### GetUploadedAt

`func (o *CranPackageUpload) GetUploadedAt() time.Time`

GetUploadedAt returns the UploadedAt field if non-nil, zero value otherwise.

### GetUploadedAtOk

`func (o *CranPackageUpload) GetUploadedAtOk() (*time.Time, bool)`

GetUploadedAtOk returns a tuple with the UploadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedAt

`func (o *CranPackageUpload) SetUploadedAt(v time.Time)`

SetUploadedAt sets UploadedAt field to given value.

### HasUploadedAt

`func (o *CranPackageUpload) HasUploadedAt() bool`

HasUploadedAt returns a boolean if a field has been set.

### GetUploader

`func (o *CranPackageUpload) GetUploader() string`

GetUploader returns the Uploader field if non-nil, zero value otherwise.

### GetUploaderOk

`func (o *CranPackageUpload) GetUploaderOk() (*string, bool)`

GetUploaderOk returns a tuple with the Uploader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploader

`func (o *CranPackageUpload) SetUploader(v string)`

SetUploader sets Uploader field to given value.

### HasUploader

`func (o *CranPackageUpload) HasUploader() bool`

HasUploader returns a boolean if a field has been set.

### GetUploaderUrl

`func (o *CranPackageUpload) GetUploaderUrl() string`

GetUploaderUrl returns the UploaderUrl field if non-nil, zero value otherwise.

### GetUploaderUrlOk

`func (o *CranPackageUpload) GetUploaderUrlOk() (*string, bool)`

GetUploaderUrlOk returns a tuple with the UploaderUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploaderUrl

`func (o *CranPackageUpload) SetUploaderUrl(v string)`

SetUploaderUrl sets UploaderUrl field to given value.

### HasUploaderUrl

`func (o *CranPackageUpload) HasUploaderUrl() bool`

HasUploaderUrl returns a boolean if a field has been set.

### GetVersion

`func (o *CranPackageUpload) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CranPackageUpload) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CranPackageUpload) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CranPackageUpload) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVersionOrig

`func (o *CranPackageUpload) GetVersionOrig() string`

GetVersionOrig returns the VersionOrig field if non-nil, zero value otherwise.

### GetVersionOrigOk

`func (o *CranPackageUpload) GetVersionOrigOk() (*string, bool)`

GetVersionOrigOk returns a tuple with the VersionOrig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionOrig

`func (o *CranPackageUpload) SetVersionOrig(v string)`

SetVersionOrig sets VersionOrig field to given value.

### HasVersionOrig

`func (o *CranPackageUpload) HasVersionOrig() bool`

HasVersionOrig returns a boolean if a field has been set.

### GetVulnerabilityScanResultsUrl

`func (o *CranPackageUpload) GetVulnerabilityScanResultsUrl() string`

GetVulnerabilityScanResultsUrl returns the VulnerabilityScanResultsUrl field if non-nil, zero value otherwise.

### GetVulnerabilityScanResultsUrlOk

`func (o *CranPackageUpload) GetVulnerabilityScanResultsUrlOk() (*string, bool)`

GetVulnerabilityScanResultsUrlOk returns a tuple with the VulnerabilityScanResultsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityScanResultsUrl

`func (o *CranPackageUpload) SetVulnerabilityScanResultsUrl(v string)`

SetVulnerabilityScanResultsUrl sets VulnerabilityScanResultsUrl field to given value.

### HasVulnerabilityScanResultsUrl

`func (o *CranPackageUpload) HasVulnerabilityScanResultsUrl() bool`

HasVulnerabilityScanResultsUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


