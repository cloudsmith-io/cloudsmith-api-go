# RawPackageUpload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Architectures** | Pointer to [**[]PackagesOwnerRepoArchitectures**](PackagesOwnerRepoArchitectures.md) |  | [optional] 
**CdnUrl** | Pointer to **string** |  | [optional] 
**ChecksumMd5** | Pointer to **string** |  | [optional] 
**ChecksumSha1** | Pointer to **string** |  | [optional] 
**ChecksumSha256** | Pointer to **string** |  | [optional] 
**ChecksumSha512** | Pointer to **string** |  | [optional] 
**DependenciesChecksumMd5** | Pointer to **string** | A checksum of all of the package&#39;s dependencies. | [optional] 
**DependenciesUrl** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** | A textual description of this package. | [optional] 
**Distro** | Pointer to **map[string]interface{}** |  | [optional] 
**DistroVersion** | Pointer to **map[string]interface{}** |  | [optional] 
**Downloads** | Pointer to **int64** |  | [optional] 
**Epoch** | Pointer to **int64** | The epoch of the package version (if any). | [optional] 
**Extension** | Pointer to **string** |  | [optional] 
**Filename** | Pointer to **string** |  | [optional] 
**Files** | Pointer to [**[]PackagesOwnerRepoFiles**](PackagesOwnerRepoFiles.md) |  | [optional] 
**Format** | Pointer to **string** |  | [optional] 
**FormatUrl** | Pointer to **string** |  | [optional] 
**IdentifierPerm** | Pointer to **string** | Unique and permanent identifier for the package. | [optional] 
**Indexed** | Pointer to **bool** |  | [optional] 
**IsDownloadable** | Pointer to **bool** |  | [optional] 
**IsQuarantined** | Pointer to **bool** |  | [optional] 
**IsSyncAwaiting** | Pointer to **bool** |  | [optional] 
**IsSyncCompleted** | Pointer to **bool** |  | [optional] 
**IsSyncFailed** | Pointer to **bool** |  | [optional] 
**IsSyncInFlight** | Pointer to **bool** |  | [optional] 
**IsSyncInProgress** | Pointer to **bool** |  | [optional] 
**License** | Pointer to **string** | The license of this package. | [optional] 
**Name** | Pointer to **string** | The name of this package. | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**NamespaceUrl** | Pointer to **string** |  | [optional] 
**NumFiles** | Pointer to **int64** |  | [optional] 
**PackageType** | Pointer to **int64** | The type of package contents. | [optional] 
**Release** | Pointer to **string** | The release of the package version (if any). | [optional] 
**Repository** | Pointer to **string** |  | [optional] 
**RepositoryUrl** | Pointer to **string** |  | [optional] 
**SecurityScanCompletedAt** | Pointer to **string** | The datetime the security scanning was completed. | [optional] 
**SecurityScanStartedAt** | Pointer to **string** | The datetime the security scanning was started. | [optional] 
**SecurityScanStatus** | Pointer to **string** |  | [optional] 
**SecurityScanStatusUpdatedAt** | Pointer to **string** | The datetime the security scanning status was updated. | [optional] 
**SelfHtmlUrl** | Pointer to **string** |  | [optional] 
**SelfUrl** | Pointer to **string** |  | [optional] 
**SignatureUrl** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int64** | The calculated size of the package. | [optional] 
**Slug** | Pointer to **string** | The public unique identifier for the package. | [optional] 
**SlugPerm** | Pointer to **string** |  | [optional] 
**Stage** | Pointer to **int64** | The synchronisation (in progress) stage of the package. | [optional] 
**StageStr** | Pointer to **string** |  | [optional] 
**StageUpdatedAt** | Pointer to **string** | The datetime the package stage was updated at. | [optional] 
**Status** | Pointer to **int64** | The synchronisation status of the package. | [optional] 
**StatusReason** | Pointer to **string** | A textual description for the synchronous status reason (if any | [optional] 
**StatusStr** | Pointer to **string** |  | [optional] 
**StatusUpdatedAt** | Pointer to **string** | The datetime the package status was updated at. | [optional] 
**StatusUrl** | Pointer to **string** |  | [optional] 
**Subtype** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** | A one-liner synopsis of this package. | [optional] 
**SyncFinishedAt** | Pointer to **string** | The datetime the package sync was finished at. | [optional] 
**SyncProgress** | Pointer to **int64** | Synchronisation progress (from 0-100) | [optional] 
**TagsImmutable** | Pointer to **map[string]interface{}** | All immutable tags on the package, grouped by tag type. Immutable tags cannot be (easily) deleted. | [optional] 
**TypeDisplay** | Pointer to **string** |  | [optional] 
**UploadedAt** | Pointer to **string** | The date this package was uploaded. | [optional] 
**Uploader** | Pointer to **string** |  | [optional] 
**UploaderUrl** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** | The raw version for this package. | [optional] 
**VersionOrig** | Pointer to **string** |  | [optional] 
**VulnerabilityScanResultsUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewRawPackageUpload

`func NewRawPackageUpload() *RawPackageUpload`

NewRawPackageUpload instantiates a new RawPackageUpload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRawPackageUploadWithDefaults

`func NewRawPackageUploadWithDefaults() *RawPackageUpload`

NewRawPackageUploadWithDefaults instantiates a new RawPackageUpload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchitectures

`func (o *RawPackageUpload) GetArchitectures() []PackagesOwnerRepoArchitectures`

GetArchitectures returns the Architectures field if non-nil, zero value otherwise.

### GetArchitecturesOk

`func (o *RawPackageUpload) GetArchitecturesOk() (*[]PackagesOwnerRepoArchitectures, bool)`

GetArchitecturesOk returns a tuple with the Architectures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitectures

`func (o *RawPackageUpload) SetArchitectures(v []PackagesOwnerRepoArchitectures)`

SetArchitectures sets Architectures field to given value.

### HasArchitectures

`func (o *RawPackageUpload) HasArchitectures() bool`

HasArchitectures returns a boolean if a field has been set.

### GetCdnUrl

`func (o *RawPackageUpload) GetCdnUrl() string`

GetCdnUrl returns the CdnUrl field if non-nil, zero value otherwise.

### GetCdnUrlOk

`func (o *RawPackageUpload) GetCdnUrlOk() (*string, bool)`

GetCdnUrlOk returns a tuple with the CdnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnUrl

`func (o *RawPackageUpload) SetCdnUrl(v string)`

SetCdnUrl sets CdnUrl field to given value.

### HasCdnUrl

`func (o *RawPackageUpload) HasCdnUrl() bool`

HasCdnUrl returns a boolean if a field has been set.

### GetChecksumMd5

`func (o *RawPackageUpload) GetChecksumMd5() string`

GetChecksumMd5 returns the ChecksumMd5 field if non-nil, zero value otherwise.

### GetChecksumMd5Ok

`func (o *RawPackageUpload) GetChecksumMd5Ok() (*string, bool)`

GetChecksumMd5Ok returns a tuple with the ChecksumMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumMd5

`func (o *RawPackageUpload) SetChecksumMd5(v string)`

SetChecksumMd5 sets ChecksumMd5 field to given value.

### HasChecksumMd5

`func (o *RawPackageUpload) HasChecksumMd5() bool`

HasChecksumMd5 returns a boolean if a field has been set.

### GetChecksumSha1

`func (o *RawPackageUpload) GetChecksumSha1() string`

GetChecksumSha1 returns the ChecksumSha1 field if non-nil, zero value otherwise.

### GetChecksumSha1Ok

`func (o *RawPackageUpload) GetChecksumSha1Ok() (*string, bool)`

GetChecksumSha1Ok returns a tuple with the ChecksumSha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumSha1

`func (o *RawPackageUpload) SetChecksumSha1(v string)`

SetChecksumSha1 sets ChecksumSha1 field to given value.

### HasChecksumSha1

`func (o *RawPackageUpload) HasChecksumSha1() bool`

HasChecksumSha1 returns a boolean if a field has been set.

### GetChecksumSha256

`func (o *RawPackageUpload) GetChecksumSha256() string`

GetChecksumSha256 returns the ChecksumSha256 field if non-nil, zero value otherwise.

### GetChecksumSha256Ok

`func (o *RawPackageUpload) GetChecksumSha256Ok() (*string, bool)`

GetChecksumSha256Ok returns a tuple with the ChecksumSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumSha256

`func (o *RawPackageUpload) SetChecksumSha256(v string)`

SetChecksumSha256 sets ChecksumSha256 field to given value.

### HasChecksumSha256

`func (o *RawPackageUpload) HasChecksumSha256() bool`

HasChecksumSha256 returns a boolean if a field has been set.

### GetChecksumSha512

`func (o *RawPackageUpload) GetChecksumSha512() string`

GetChecksumSha512 returns the ChecksumSha512 field if non-nil, zero value otherwise.

### GetChecksumSha512Ok

`func (o *RawPackageUpload) GetChecksumSha512Ok() (*string, bool)`

GetChecksumSha512Ok returns a tuple with the ChecksumSha512 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumSha512

`func (o *RawPackageUpload) SetChecksumSha512(v string)`

SetChecksumSha512 sets ChecksumSha512 field to given value.

### HasChecksumSha512

`func (o *RawPackageUpload) HasChecksumSha512() bool`

HasChecksumSha512 returns a boolean if a field has been set.

### GetDependenciesChecksumMd5

`func (o *RawPackageUpload) GetDependenciesChecksumMd5() string`

GetDependenciesChecksumMd5 returns the DependenciesChecksumMd5 field if non-nil, zero value otherwise.

### GetDependenciesChecksumMd5Ok

`func (o *RawPackageUpload) GetDependenciesChecksumMd5Ok() (*string, bool)`

GetDependenciesChecksumMd5Ok returns a tuple with the DependenciesChecksumMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependenciesChecksumMd5

`func (o *RawPackageUpload) SetDependenciesChecksumMd5(v string)`

SetDependenciesChecksumMd5 sets DependenciesChecksumMd5 field to given value.

### HasDependenciesChecksumMd5

`func (o *RawPackageUpload) HasDependenciesChecksumMd5() bool`

HasDependenciesChecksumMd5 returns a boolean if a field has been set.

### GetDependenciesUrl

`func (o *RawPackageUpload) GetDependenciesUrl() string`

GetDependenciesUrl returns the DependenciesUrl field if non-nil, zero value otherwise.

### GetDependenciesUrlOk

`func (o *RawPackageUpload) GetDependenciesUrlOk() (*string, bool)`

GetDependenciesUrlOk returns a tuple with the DependenciesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependenciesUrl

`func (o *RawPackageUpload) SetDependenciesUrl(v string)`

SetDependenciesUrl sets DependenciesUrl field to given value.

### HasDependenciesUrl

`func (o *RawPackageUpload) HasDependenciesUrl() bool`

HasDependenciesUrl returns a boolean if a field has been set.

### GetDescription

`func (o *RawPackageUpload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RawPackageUpload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RawPackageUpload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RawPackageUpload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDistro

`func (o *RawPackageUpload) GetDistro() map[string]interface{}`

GetDistro returns the Distro field if non-nil, zero value otherwise.

### GetDistroOk

`func (o *RawPackageUpload) GetDistroOk() (*map[string]interface{}, bool)`

GetDistroOk returns a tuple with the Distro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistro

`func (o *RawPackageUpload) SetDistro(v map[string]interface{})`

SetDistro sets Distro field to given value.

### HasDistro

`func (o *RawPackageUpload) HasDistro() bool`

HasDistro returns a boolean if a field has been set.

### GetDistroVersion

`func (o *RawPackageUpload) GetDistroVersion() map[string]interface{}`

GetDistroVersion returns the DistroVersion field if non-nil, zero value otherwise.

### GetDistroVersionOk

`func (o *RawPackageUpload) GetDistroVersionOk() (*map[string]interface{}, bool)`

GetDistroVersionOk returns a tuple with the DistroVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistroVersion

`func (o *RawPackageUpload) SetDistroVersion(v map[string]interface{})`

SetDistroVersion sets DistroVersion field to given value.

### HasDistroVersion

`func (o *RawPackageUpload) HasDistroVersion() bool`

HasDistroVersion returns a boolean if a field has been set.

### GetDownloads

`func (o *RawPackageUpload) GetDownloads() int64`

GetDownloads returns the Downloads field if non-nil, zero value otherwise.

### GetDownloadsOk

`func (o *RawPackageUpload) GetDownloadsOk() (*int64, bool)`

GetDownloadsOk returns a tuple with the Downloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloads

`func (o *RawPackageUpload) SetDownloads(v int64)`

SetDownloads sets Downloads field to given value.

### HasDownloads

`func (o *RawPackageUpload) HasDownloads() bool`

HasDownloads returns a boolean if a field has been set.

### GetEpoch

`func (o *RawPackageUpload) GetEpoch() int64`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *RawPackageUpload) GetEpochOk() (*int64, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *RawPackageUpload) SetEpoch(v int64)`

SetEpoch sets Epoch field to given value.

### HasEpoch

`func (o *RawPackageUpload) HasEpoch() bool`

HasEpoch returns a boolean if a field has been set.

### GetExtension

`func (o *RawPackageUpload) GetExtension() string`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *RawPackageUpload) GetExtensionOk() (*string, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *RawPackageUpload) SetExtension(v string)`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *RawPackageUpload) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### GetFilename

`func (o *RawPackageUpload) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *RawPackageUpload) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *RawPackageUpload) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *RawPackageUpload) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetFiles

`func (o *RawPackageUpload) GetFiles() []PackagesOwnerRepoFiles`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *RawPackageUpload) GetFilesOk() (*[]PackagesOwnerRepoFiles, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *RawPackageUpload) SetFiles(v []PackagesOwnerRepoFiles)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *RawPackageUpload) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetFormat

`func (o *RawPackageUpload) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *RawPackageUpload) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *RawPackageUpload) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *RawPackageUpload) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetFormatUrl

`func (o *RawPackageUpload) GetFormatUrl() string`

GetFormatUrl returns the FormatUrl field if non-nil, zero value otherwise.

### GetFormatUrlOk

`func (o *RawPackageUpload) GetFormatUrlOk() (*string, bool)`

GetFormatUrlOk returns a tuple with the FormatUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatUrl

`func (o *RawPackageUpload) SetFormatUrl(v string)`

SetFormatUrl sets FormatUrl field to given value.

### HasFormatUrl

`func (o *RawPackageUpload) HasFormatUrl() bool`

HasFormatUrl returns a boolean if a field has been set.

### GetIdentifierPerm

`func (o *RawPackageUpload) GetIdentifierPerm() string`

GetIdentifierPerm returns the IdentifierPerm field if non-nil, zero value otherwise.

### GetIdentifierPermOk

`func (o *RawPackageUpload) GetIdentifierPermOk() (*string, bool)`

GetIdentifierPermOk returns a tuple with the IdentifierPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierPerm

`func (o *RawPackageUpload) SetIdentifierPerm(v string)`

SetIdentifierPerm sets IdentifierPerm field to given value.

### HasIdentifierPerm

`func (o *RawPackageUpload) HasIdentifierPerm() bool`

HasIdentifierPerm returns a boolean if a field has been set.

### GetIndexed

`func (o *RawPackageUpload) GetIndexed() bool`

GetIndexed returns the Indexed field if non-nil, zero value otherwise.

### GetIndexedOk

`func (o *RawPackageUpload) GetIndexedOk() (*bool, bool)`

GetIndexedOk returns a tuple with the Indexed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexed

`func (o *RawPackageUpload) SetIndexed(v bool)`

SetIndexed sets Indexed field to given value.

### HasIndexed

`func (o *RawPackageUpload) HasIndexed() bool`

HasIndexed returns a boolean if a field has been set.

### GetIsDownloadable

`func (o *RawPackageUpload) GetIsDownloadable() bool`

GetIsDownloadable returns the IsDownloadable field if non-nil, zero value otherwise.

### GetIsDownloadableOk

`func (o *RawPackageUpload) GetIsDownloadableOk() (*bool, bool)`

GetIsDownloadableOk returns a tuple with the IsDownloadable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDownloadable

`func (o *RawPackageUpload) SetIsDownloadable(v bool)`

SetIsDownloadable sets IsDownloadable field to given value.

### HasIsDownloadable

`func (o *RawPackageUpload) HasIsDownloadable() bool`

HasIsDownloadable returns a boolean if a field has been set.

### GetIsQuarantined

`func (o *RawPackageUpload) GetIsQuarantined() bool`

GetIsQuarantined returns the IsQuarantined field if non-nil, zero value otherwise.

### GetIsQuarantinedOk

`func (o *RawPackageUpload) GetIsQuarantinedOk() (*bool, bool)`

GetIsQuarantinedOk returns a tuple with the IsQuarantined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQuarantined

`func (o *RawPackageUpload) SetIsQuarantined(v bool)`

SetIsQuarantined sets IsQuarantined field to given value.

### HasIsQuarantined

`func (o *RawPackageUpload) HasIsQuarantined() bool`

HasIsQuarantined returns a boolean if a field has been set.

### GetIsSyncAwaiting

`func (o *RawPackageUpload) GetIsSyncAwaiting() bool`

GetIsSyncAwaiting returns the IsSyncAwaiting field if non-nil, zero value otherwise.

### GetIsSyncAwaitingOk

`func (o *RawPackageUpload) GetIsSyncAwaitingOk() (*bool, bool)`

GetIsSyncAwaitingOk returns a tuple with the IsSyncAwaiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncAwaiting

`func (o *RawPackageUpload) SetIsSyncAwaiting(v bool)`

SetIsSyncAwaiting sets IsSyncAwaiting field to given value.

### HasIsSyncAwaiting

`func (o *RawPackageUpload) HasIsSyncAwaiting() bool`

HasIsSyncAwaiting returns a boolean if a field has been set.

### GetIsSyncCompleted

`func (o *RawPackageUpload) GetIsSyncCompleted() bool`

GetIsSyncCompleted returns the IsSyncCompleted field if non-nil, zero value otherwise.

### GetIsSyncCompletedOk

`func (o *RawPackageUpload) GetIsSyncCompletedOk() (*bool, bool)`

GetIsSyncCompletedOk returns a tuple with the IsSyncCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncCompleted

`func (o *RawPackageUpload) SetIsSyncCompleted(v bool)`

SetIsSyncCompleted sets IsSyncCompleted field to given value.

### HasIsSyncCompleted

`func (o *RawPackageUpload) HasIsSyncCompleted() bool`

HasIsSyncCompleted returns a boolean if a field has been set.

### GetIsSyncFailed

`func (o *RawPackageUpload) GetIsSyncFailed() bool`

GetIsSyncFailed returns the IsSyncFailed field if non-nil, zero value otherwise.

### GetIsSyncFailedOk

`func (o *RawPackageUpload) GetIsSyncFailedOk() (*bool, bool)`

GetIsSyncFailedOk returns a tuple with the IsSyncFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncFailed

`func (o *RawPackageUpload) SetIsSyncFailed(v bool)`

SetIsSyncFailed sets IsSyncFailed field to given value.

### HasIsSyncFailed

`func (o *RawPackageUpload) HasIsSyncFailed() bool`

HasIsSyncFailed returns a boolean if a field has been set.

### GetIsSyncInFlight

`func (o *RawPackageUpload) GetIsSyncInFlight() bool`

GetIsSyncInFlight returns the IsSyncInFlight field if non-nil, zero value otherwise.

### GetIsSyncInFlightOk

`func (o *RawPackageUpload) GetIsSyncInFlightOk() (*bool, bool)`

GetIsSyncInFlightOk returns a tuple with the IsSyncInFlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncInFlight

`func (o *RawPackageUpload) SetIsSyncInFlight(v bool)`

SetIsSyncInFlight sets IsSyncInFlight field to given value.

### HasIsSyncInFlight

`func (o *RawPackageUpload) HasIsSyncInFlight() bool`

HasIsSyncInFlight returns a boolean if a field has been set.

### GetIsSyncInProgress

`func (o *RawPackageUpload) GetIsSyncInProgress() bool`

GetIsSyncInProgress returns the IsSyncInProgress field if non-nil, zero value otherwise.

### GetIsSyncInProgressOk

`func (o *RawPackageUpload) GetIsSyncInProgressOk() (*bool, bool)`

GetIsSyncInProgressOk returns a tuple with the IsSyncInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncInProgress

`func (o *RawPackageUpload) SetIsSyncInProgress(v bool)`

SetIsSyncInProgress sets IsSyncInProgress field to given value.

### HasIsSyncInProgress

`func (o *RawPackageUpload) HasIsSyncInProgress() bool`

HasIsSyncInProgress returns a boolean if a field has been set.

### GetLicense

`func (o *RawPackageUpload) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *RawPackageUpload) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *RawPackageUpload) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *RawPackageUpload) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetName

`func (o *RawPackageUpload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RawPackageUpload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RawPackageUpload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RawPackageUpload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *RawPackageUpload) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *RawPackageUpload) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *RawPackageUpload) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *RawPackageUpload) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetNamespaceUrl

`func (o *RawPackageUpload) GetNamespaceUrl() string`

GetNamespaceUrl returns the NamespaceUrl field if non-nil, zero value otherwise.

### GetNamespaceUrlOk

`func (o *RawPackageUpload) GetNamespaceUrlOk() (*string, bool)`

GetNamespaceUrlOk returns a tuple with the NamespaceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceUrl

`func (o *RawPackageUpload) SetNamespaceUrl(v string)`

SetNamespaceUrl sets NamespaceUrl field to given value.

### HasNamespaceUrl

`func (o *RawPackageUpload) HasNamespaceUrl() bool`

HasNamespaceUrl returns a boolean if a field has been set.

### GetNumFiles

`func (o *RawPackageUpload) GetNumFiles() int64`

GetNumFiles returns the NumFiles field if non-nil, zero value otherwise.

### GetNumFilesOk

`func (o *RawPackageUpload) GetNumFilesOk() (*int64, bool)`

GetNumFilesOk returns a tuple with the NumFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFiles

`func (o *RawPackageUpload) SetNumFiles(v int64)`

SetNumFiles sets NumFiles field to given value.

### HasNumFiles

`func (o *RawPackageUpload) HasNumFiles() bool`

HasNumFiles returns a boolean if a field has been set.

### GetPackageType

`func (o *RawPackageUpload) GetPackageType() int64`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *RawPackageUpload) GetPackageTypeOk() (*int64, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *RawPackageUpload) SetPackageType(v int64)`

SetPackageType sets PackageType field to given value.

### HasPackageType

`func (o *RawPackageUpload) HasPackageType() bool`

HasPackageType returns a boolean if a field has been set.

### GetRelease

`func (o *RawPackageUpload) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *RawPackageUpload) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *RawPackageUpload) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *RawPackageUpload) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetRepository

`func (o *RawPackageUpload) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *RawPackageUpload) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *RawPackageUpload) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *RawPackageUpload) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetRepositoryUrl

`func (o *RawPackageUpload) GetRepositoryUrl() string`

GetRepositoryUrl returns the RepositoryUrl field if non-nil, zero value otherwise.

### GetRepositoryUrlOk

`func (o *RawPackageUpload) GetRepositoryUrlOk() (*string, bool)`

GetRepositoryUrlOk returns a tuple with the RepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryUrl

`func (o *RawPackageUpload) SetRepositoryUrl(v string)`

SetRepositoryUrl sets RepositoryUrl field to given value.

### HasRepositoryUrl

`func (o *RawPackageUpload) HasRepositoryUrl() bool`

HasRepositoryUrl returns a boolean if a field has been set.

### GetSecurityScanCompletedAt

`func (o *RawPackageUpload) GetSecurityScanCompletedAt() string`

GetSecurityScanCompletedAt returns the SecurityScanCompletedAt field if non-nil, zero value otherwise.

### GetSecurityScanCompletedAtOk

`func (o *RawPackageUpload) GetSecurityScanCompletedAtOk() (*string, bool)`

GetSecurityScanCompletedAtOk returns a tuple with the SecurityScanCompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityScanCompletedAt

`func (o *RawPackageUpload) SetSecurityScanCompletedAt(v string)`

SetSecurityScanCompletedAt sets SecurityScanCompletedAt field to given value.

### HasSecurityScanCompletedAt

`func (o *RawPackageUpload) HasSecurityScanCompletedAt() bool`

HasSecurityScanCompletedAt returns a boolean if a field has been set.

### GetSecurityScanStartedAt

`func (o *RawPackageUpload) GetSecurityScanStartedAt() string`

GetSecurityScanStartedAt returns the SecurityScanStartedAt field if non-nil, zero value otherwise.

### GetSecurityScanStartedAtOk

`func (o *RawPackageUpload) GetSecurityScanStartedAtOk() (*string, bool)`

GetSecurityScanStartedAtOk returns a tuple with the SecurityScanStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityScanStartedAt

`func (o *RawPackageUpload) SetSecurityScanStartedAt(v string)`

SetSecurityScanStartedAt sets SecurityScanStartedAt field to given value.

### HasSecurityScanStartedAt

`func (o *RawPackageUpload) HasSecurityScanStartedAt() bool`

HasSecurityScanStartedAt returns a boolean if a field has been set.

### GetSecurityScanStatus

`func (o *RawPackageUpload) GetSecurityScanStatus() string`

GetSecurityScanStatus returns the SecurityScanStatus field if non-nil, zero value otherwise.

### GetSecurityScanStatusOk

`func (o *RawPackageUpload) GetSecurityScanStatusOk() (*string, bool)`

GetSecurityScanStatusOk returns a tuple with the SecurityScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityScanStatus

`func (o *RawPackageUpload) SetSecurityScanStatus(v string)`

SetSecurityScanStatus sets SecurityScanStatus field to given value.

### HasSecurityScanStatus

`func (o *RawPackageUpload) HasSecurityScanStatus() bool`

HasSecurityScanStatus returns a boolean if a field has been set.

### GetSecurityScanStatusUpdatedAt

`func (o *RawPackageUpload) GetSecurityScanStatusUpdatedAt() string`

GetSecurityScanStatusUpdatedAt returns the SecurityScanStatusUpdatedAt field if non-nil, zero value otherwise.

### GetSecurityScanStatusUpdatedAtOk

`func (o *RawPackageUpload) GetSecurityScanStatusUpdatedAtOk() (*string, bool)`

GetSecurityScanStatusUpdatedAtOk returns a tuple with the SecurityScanStatusUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityScanStatusUpdatedAt

`func (o *RawPackageUpload) SetSecurityScanStatusUpdatedAt(v string)`

SetSecurityScanStatusUpdatedAt sets SecurityScanStatusUpdatedAt field to given value.

### HasSecurityScanStatusUpdatedAt

`func (o *RawPackageUpload) HasSecurityScanStatusUpdatedAt() bool`

HasSecurityScanStatusUpdatedAt returns a boolean if a field has been set.

### GetSelfHtmlUrl

`func (o *RawPackageUpload) GetSelfHtmlUrl() string`

GetSelfHtmlUrl returns the SelfHtmlUrl field if non-nil, zero value otherwise.

### GetSelfHtmlUrlOk

`func (o *RawPackageUpload) GetSelfHtmlUrlOk() (*string, bool)`

GetSelfHtmlUrlOk returns a tuple with the SelfHtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfHtmlUrl

`func (o *RawPackageUpload) SetSelfHtmlUrl(v string)`

SetSelfHtmlUrl sets SelfHtmlUrl field to given value.

### HasSelfHtmlUrl

`func (o *RawPackageUpload) HasSelfHtmlUrl() bool`

HasSelfHtmlUrl returns a boolean if a field has been set.

### GetSelfUrl

`func (o *RawPackageUpload) GetSelfUrl() string`

GetSelfUrl returns the SelfUrl field if non-nil, zero value otherwise.

### GetSelfUrlOk

`func (o *RawPackageUpload) GetSelfUrlOk() (*string, bool)`

GetSelfUrlOk returns a tuple with the SelfUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfUrl

`func (o *RawPackageUpload) SetSelfUrl(v string)`

SetSelfUrl sets SelfUrl field to given value.

### HasSelfUrl

`func (o *RawPackageUpload) HasSelfUrl() bool`

HasSelfUrl returns a boolean if a field has been set.

### GetSignatureUrl

`func (o *RawPackageUpload) GetSignatureUrl() string`

GetSignatureUrl returns the SignatureUrl field if non-nil, zero value otherwise.

### GetSignatureUrlOk

`func (o *RawPackageUpload) GetSignatureUrlOk() (*string, bool)`

GetSignatureUrlOk returns a tuple with the SignatureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureUrl

`func (o *RawPackageUpload) SetSignatureUrl(v string)`

SetSignatureUrl sets SignatureUrl field to given value.

### HasSignatureUrl

`func (o *RawPackageUpload) HasSignatureUrl() bool`

HasSignatureUrl returns a boolean if a field has been set.

### GetSize

`func (o *RawPackageUpload) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *RawPackageUpload) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *RawPackageUpload) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *RawPackageUpload) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSlug

`func (o *RawPackageUpload) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *RawPackageUpload) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *RawPackageUpload) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *RawPackageUpload) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSlugPerm

`func (o *RawPackageUpload) GetSlugPerm() string`

GetSlugPerm returns the SlugPerm field if non-nil, zero value otherwise.

### GetSlugPermOk

`func (o *RawPackageUpload) GetSlugPermOk() (*string, bool)`

GetSlugPermOk returns a tuple with the SlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugPerm

`func (o *RawPackageUpload) SetSlugPerm(v string)`

SetSlugPerm sets SlugPerm field to given value.

### HasSlugPerm

`func (o *RawPackageUpload) HasSlugPerm() bool`

HasSlugPerm returns a boolean if a field has been set.

### GetStage

`func (o *RawPackageUpload) GetStage() int64`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *RawPackageUpload) GetStageOk() (*int64, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *RawPackageUpload) SetStage(v int64)`

SetStage sets Stage field to given value.

### HasStage

`func (o *RawPackageUpload) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetStageStr

`func (o *RawPackageUpload) GetStageStr() string`

GetStageStr returns the StageStr field if non-nil, zero value otherwise.

### GetStageStrOk

`func (o *RawPackageUpload) GetStageStrOk() (*string, bool)`

GetStageStrOk returns a tuple with the StageStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageStr

`func (o *RawPackageUpload) SetStageStr(v string)`

SetStageStr sets StageStr field to given value.

### HasStageStr

`func (o *RawPackageUpload) HasStageStr() bool`

HasStageStr returns a boolean if a field has been set.

### GetStageUpdatedAt

`func (o *RawPackageUpload) GetStageUpdatedAt() string`

GetStageUpdatedAt returns the StageUpdatedAt field if non-nil, zero value otherwise.

### GetStageUpdatedAtOk

`func (o *RawPackageUpload) GetStageUpdatedAtOk() (*string, bool)`

GetStageUpdatedAtOk returns a tuple with the StageUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageUpdatedAt

`func (o *RawPackageUpload) SetStageUpdatedAt(v string)`

SetStageUpdatedAt sets StageUpdatedAt field to given value.

### HasStageUpdatedAt

`func (o *RawPackageUpload) HasStageUpdatedAt() bool`

HasStageUpdatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *RawPackageUpload) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RawPackageUpload) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RawPackageUpload) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RawPackageUpload) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusReason

`func (o *RawPackageUpload) GetStatusReason() string`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *RawPackageUpload) GetStatusReasonOk() (*string, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *RawPackageUpload) SetStatusReason(v string)`

SetStatusReason sets StatusReason field to given value.

### HasStatusReason

`func (o *RawPackageUpload) HasStatusReason() bool`

HasStatusReason returns a boolean if a field has been set.

### GetStatusStr

`func (o *RawPackageUpload) GetStatusStr() string`

GetStatusStr returns the StatusStr field if non-nil, zero value otherwise.

### GetStatusStrOk

`func (o *RawPackageUpload) GetStatusStrOk() (*string, bool)`

GetStatusStrOk returns a tuple with the StatusStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusStr

`func (o *RawPackageUpload) SetStatusStr(v string)`

SetStatusStr sets StatusStr field to given value.

### HasStatusStr

`func (o *RawPackageUpload) HasStatusStr() bool`

HasStatusStr returns a boolean if a field has been set.

### GetStatusUpdatedAt

`func (o *RawPackageUpload) GetStatusUpdatedAt() string`

GetStatusUpdatedAt returns the StatusUpdatedAt field if non-nil, zero value otherwise.

### GetStatusUpdatedAtOk

`func (o *RawPackageUpload) GetStatusUpdatedAtOk() (*string, bool)`

GetStatusUpdatedAtOk returns a tuple with the StatusUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusUpdatedAt

`func (o *RawPackageUpload) SetStatusUpdatedAt(v string)`

SetStatusUpdatedAt sets StatusUpdatedAt field to given value.

### HasStatusUpdatedAt

`func (o *RawPackageUpload) HasStatusUpdatedAt() bool`

HasStatusUpdatedAt returns a boolean if a field has been set.

### GetStatusUrl

`func (o *RawPackageUpload) GetStatusUrl() string`

GetStatusUrl returns the StatusUrl field if non-nil, zero value otherwise.

### GetStatusUrlOk

`func (o *RawPackageUpload) GetStatusUrlOk() (*string, bool)`

GetStatusUrlOk returns a tuple with the StatusUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusUrl

`func (o *RawPackageUpload) SetStatusUrl(v string)`

SetStatusUrl sets StatusUrl field to given value.

### HasStatusUrl

`func (o *RawPackageUpload) HasStatusUrl() bool`

HasStatusUrl returns a boolean if a field has been set.

### GetSubtype

`func (o *RawPackageUpload) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *RawPackageUpload) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *RawPackageUpload) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *RawPackageUpload) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetSummary

`func (o *RawPackageUpload) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *RawPackageUpload) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *RawPackageUpload) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *RawPackageUpload) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetSyncFinishedAt

`func (o *RawPackageUpload) GetSyncFinishedAt() string`

GetSyncFinishedAt returns the SyncFinishedAt field if non-nil, zero value otherwise.

### GetSyncFinishedAtOk

`func (o *RawPackageUpload) GetSyncFinishedAtOk() (*string, bool)`

GetSyncFinishedAtOk returns a tuple with the SyncFinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncFinishedAt

`func (o *RawPackageUpload) SetSyncFinishedAt(v string)`

SetSyncFinishedAt sets SyncFinishedAt field to given value.

### HasSyncFinishedAt

`func (o *RawPackageUpload) HasSyncFinishedAt() bool`

HasSyncFinishedAt returns a boolean if a field has been set.

### GetSyncProgress

`func (o *RawPackageUpload) GetSyncProgress() int64`

GetSyncProgress returns the SyncProgress field if non-nil, zero value otherwise.

### GetSyncProgressOk

`func (o *RawPackageUpload) GetSyncProgressOk() (*int64, bool)`

GetSyncProgressOk returns a tuple with the SyncProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncProgress

`func (o *RawPackageUpload) SetSyncProgress(v int64)`

SetSyncProgress sets SyncProgress field to given value.

### HasSyncProgress

`func (o *RawPackageUpload) HasSyncProgress() bool`

HasSyncProgress returns a boolean if a field has been set.

### GetTagsImmutable

`func (o *RawPackageUpload) GetTagsImmutable() map[string]interface{}`

GetTagsImmutable returns the TagsImmutable field if non-nil, zero value otherwise.

### GetTagsImmutableOk

`func (o *RawPackageUpload) GetTagsImmutableOk() (*map[string]interface{}, bool)`

GetTagsImmutableOk returns a tuple with the TagsImmutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsImmutable

`func (o *RawPackageUpload) SetTagsImmutable(v map[string]interface{})`

SetTagsImmutable sets TagsImmutable field to given value.

### HasTagsImmutable

`func (o *RawPackageUpload) HasTagsImmutable() bool`

HasTagsImmutable returns a boolean if a field has been set.

### GetTypeDisplay

`func (o *RawPackageUpload) GetTypeDisplay() string`

GetTypeDisplay returns the TypeDisplay field if non-nil, zero value otherwise.

### GetTypeDisplayOk

`func (o *RawPackageUpload) GetTypeDisplayOk() (*string, bool)`

GetTypeDisplayOk returns a tuple with the TypeDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeDisplay

`func (o *RawPackageUpload) SetTypeDisplay(v string)`

SetTypeDisplay sets TypeDisplay field to given value.

### HasTypeDisplay

`func (o *RawPackageUpload) HasTypeDisplay() bool`

HasTypeDisplay returns a boolean if a field has been set.

### GetUploadedAt

`func (o *RawPackageUpload) GetUploadedAt() string`

GetUploadedAt returns the UploadedAt field if non-nil, zero value otherwise.

### GetUploadedAtOk

`func (o *RawPackageUpload) GetUploadedAtOk() (*string, bool)`

GetUploadedAtOk returns a tuple with the UploadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedAt

`func (o *RawPackageUpload) SetUploadedAt(v string)`

SetUploadedAt sets UploadedAt field to given value.

### HasUploadedAt

`func (o *RawPackageUpload) HasUploadedAt() bool`

HasUploadedAt returns a boolean if a field has been set.

### GetUploader

`func (o *RawPackageUpload) GetUploader() string`

GetUploader returns the Uploader field if non-nil, zero value otherwise.

### GetUploaderOk

`func (o *RawPackageUpload) GetUploaderOk() (*string, bool)`

GetUploaderOk returns a tuple with the Uploader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploader

`func (o *RawPackageUpload) SetUploader(v string)`

SetUploader sets Uploader field to given value.

### HasUploader

`func (o *RawPackageUpload) HasUploader() bool`

HasUploader returns a boolean if a field has been set.

### GetUploaderUrl

`func (o *RawPackageUpload) GetUploaderUrl() string`

GetUploaderUrl returns the UploaderUrl field if non-nil, zero value otherwise.

### GetUploaderUrlOk

`func (o *RawPackageUpload) GetUploaderUrlOk() (*string, bool)`

GetUploaderUrlOk returns a tuple with the UploaderUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploaderUrl

`func (o *RawPackageUpload) SetUploaderUrl(v string)`

SetUploaderUrl sets UploaderUrl field to given value.

### HasUploaderUrl

`func (o *RawPackageUpload) HasUploaderUrl() bool`

HasUploaderUrl returns a boolean if a field has been set.

### GetVersion

`func (o *RawPackageUpload) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RawPackageUpload) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RawPackageUpload) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RawPackageUpload) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVersionOrig

`func (o *RawPackageUpload) GetVersionOrig() string`

GetVersionOrig returns the VersionOrig field if non-nil, zero value otherwise.

### GetVersionOrigOk

`func (o *RawPackageUpload) GetVersionOrigOk() (*string, bool)`

GetVersionOrigOk returns a tuple with the VersionOrig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionOrig

`func (o *RawPackageUpload) SetVersionOrig(v string)`

SetVersionOrig sets VersionOrig field to given value.

### HasVersionOrig

`func (o *RawPackageUpload) HasVersionOrig() bool`

HasVersionOrig returns a boolean if a field has been set.

### GetVulnerabilityScanResultsUrl

`func (o *RawPackageUpload) GetVulnerabilityScanResultsUrl() string`

GetVulnerabilityScanResultsUrl returns the VulnerabilityScanResultsUrl field if non-nil, zero value otherwise.

### GetVulnerabilityScanResultsUrlOk

`func (o *RawPackageUpload) GetVulnerabilityScanResultsUrlOk() (*string, bool)`

GetVulnerabilityScanResultsUrlOk returns a tuple with the VulnerabilityScanResultsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityScanResultsUrl

`func (o *RawPackageUpload) SetVulnerabilityScanResultsUrl(v string)`

SetVulnerabilityScanResultsUrl sets VulnerabilityScanResultsUrl field to given value.

### HasVulnerabilityScanResultsUrl

`func (o *RawPackageUpload) HasVulnerabilityScanResultsUrl() bool`

HasVulnerabilityScanResultsUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


