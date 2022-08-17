# CondaPackageUpload

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

### NewCondaPackageUpload

`func NewCondaPackageUpload() *CondaPackageUpload`

NewCondaPackageUpload instantiates a new CondaPackageUpload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCondaPackageUploadWithDefaults

`func NewCondaPackageUploadWithDefaults() *CondaPackageUpload`

NewCondaPackageUploadWithDefaults instantiates a new CondaPackageUpload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchitectures

`func (o *CondaPackageUpload) GetArchitectures() []PackagesOwnerRepoArchitectures`

GetArchitectures returns the Architectures field if non-nil, zero value otherwise.

### GetArchitecturesOk

`func (o *CondaPackageUpload) GetArchitecturesOk() (*[]PackagesOwnerRepoArchitectures, bool)`

GetArchitecturesOk returns a tuple with the Architectures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitectures

`func (o *CondaPackageUpload) SetArchitectures(v []PackagesOwnerRepoArchitectures)`

SetArchitectures sets Architectures field to given value.

### HasArchitectures

`func (o *CondaPackageUpload) HasArchitectures() bool`

HasArchitectures returns a boolean if a field has been set.

### GetCdnUrl

`func (o *CondaPackageUpload) GetCdnUrl() string`

GetCdnUrl returns the CdnUrl field if non-nil, zero value otherwise.

### GetCdnUrlOk

`func (o *CondaPackageUpload) GetCdnUrlOk() (*string, bool)`

GetCdnUrlOk returns a tuple with the CdnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnUrl

`func (o *CondaPackageUpload) SetCdnUrl(v string)`

SetCdnUrl sets CdnUrl field to given value.

### HasCdnUrl

`func (o *CondaPackageUpload) HasCdnUrl() bool`

HasCdnUrl returns a boolean if a field has been set.

### GetChecksumMd5

`func (o *CondaPackageUpload) GetChecksumMd5() string`

GetChecksumMd5 returns the ChecksumMd5 field if non-nil, zero value otherwise.

### GetChecksumMd5Ok

`func (o *CondaPackageUpload) GetChecksumMd5Ok() (*string, bool)`

GetChecksumMd5Ok returns a tuple with the ChecksumMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumMd5

`func (o *CondaPackageUpload) SetChecksumMd5(v string)`

SetChecksumMd5 sets ChecksumMd5 field to given value.

### HasChecksumMd5

`func (o *CondaPackageUpload) HasChecksumMd5() bool`

HasChecksumMd5 returns a boolean if a field has been set.

### GetChecksumSha1

`func (o *CondaPackageUpload) GetChecksumSha1() string`

GetChecksumSha1 returns the ChecksumSha1 field if non-nil, zero value otherwise.

### GetChecksumSha1Ok

`func (o *CondaPackageUpload) GetChecksumSha1Ok() (*string, bool)`

GetChecksumSha1Ok returns a tuple with the ChecksumSha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumSha1

`func (o *CondaPackageUpload) SetChecksumSha1(v string)`

SetChecksumSha1 sets ChecksumSha1 field to given value.

### HasChecksumSha1

`func (o *CondaPackageUpload) HasChecksumSha1() bool`

HasChecksumSha1 returns a boolean if a field has been set.

### GetChecksumSha256

`func (o *CondaPackageUpload) GetChecksumSha256() string`

GetChecksumSha256 returns the ChecksumSha256 field if non-nil, zero value otherwise.

### GetChecksumSha256Ok

`func (o *CondaPackageUpload) GetChecksumSha256Ok() (*string, bool)`

GetChecksumSha256Ok returns a tuple with the ChecksumSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumSha256

`func (o *CondaPackageUpload) SetChecksumSha256(v string)`

SetChecksumSha256 sets ChecksumSha256 field to given value.

### HasChecksumSha256

`func (o *CondaPackageUpload) HasChecksumSha256() bool`

HasChecksumSha256 returns a boolean if a field has been set.

### GetChecksumSha512

`func (o *CondaPackageUpload) GetChecksumSha512() string`

GetChecksumSha512 returns the ChecksumSha512 field if non-nil, zero value otherwise.

### GetChecksumSha512Ok

`func (o *CondaPackageUpload) GetChecksumSha512Ok() (*string, bool)`

GetChecksumSha512Ok returns a tuple with the ChecksumSha512 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumSha512

`func (o *CondaPackageUpload) SetChecksumSha512(v string)`

SetChecksumSha512 sets ChecksumSha512 field to given value.

### HasChecksumSha512

`func (o *CondaPackageUpload) HasChecksumSha512() bool`

HasChecksumSha512 returns a boolean if a field has been set.

### GetDependenciesChecksumMd5

`func (o *CondaPackageUpload) GetDependenciesChecksumMd5() string`

GetDependenciesChecksumMd5 returns the DependenciesChecksumMd5 field if non-nil, zero value otherwise.

### GetDependenciesChecksumMd5Ok

`func (o *CondaPackageUpload) GetDependenciesChecksumMd5Ok() (*string, bool)`

GetDependenciesChecksumMd5Ok returns a tuple with the DependenciesChecksumMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependenciesChecksumMd5

`func (o *CondaPackageUpload) SetDependenciesChecksumMd5(v string)`

SetDependenciesChecksumMd5 sets DependenciesChecksumMd5 field to given value.

### HasDependenciesChecksumMd5

`func (o *CondaPackageUpload) HasDependenciesChecksumMd5() bool`

HasDependenciesChecksumMd5 returns a boolean if a field has been set.

### GetDependenciesUrl

`func (o *CondaPackageUpload) GetDependenciesUrl() string`

GetDependenciesUrl returns the DependenciesUrl field if non-nil, zero value otherwise.

### GetDependenciesUrlOk

`func (o *CondaPackageUpload) GetDependenciesUrlOk() (*string, bool)`

GetDependenciesUrlOk returns a tuple with the DependenciesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependenciesUrl

`func (o *CondaPackageUpload) SetDependenciesUrl(v string)`

SetDependenciesUrl sets DependenciesUrl field to given value.

### HasDependenciesUrl

`func (o *CondaPackageUpload) HasDependenciesUrl() bool`

HasDependenciesUrl returns a boolean if a field has been set.

### GetDescription

`func (o *CondaPackageUpload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CondaPackageUpload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CondaPackageUpload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CondaPackageUpload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDistro

`func (o *CondaPackageUpload) GetDistro() map[string]interface{}`

GetDistro returns the Distro field if non-nil, zero value otherwise.

### GetDistroOk

`func (o *CondaPackageUpload) GetDistroOk() (*map[string]interface{}, bool)`

GetDistroOk returns a tuple with the Distro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistro

`func (o *CondaPackageUpload) SetDistro(v map[string]interface{})`

SetDistro sets Distro field to given value.

### HasDistro

`func (o *CondaPackageUpload) HasDistro() bool`

HasDistro returns a boolean if a field has been set.

### GetDistroVersion

`func (o *CondaPackageUpload) GetDistroVersion() map[string]interface{}`

GetDistroVersion returns the DistroVersion field if non-nil, zero value otherwise.

### GetDistroVersionOk

`func (o *CondaPackageUpload) GetDistroVersionOk() (*map[string]interface{}, bool)`

GetDistroVersionOk returns a tuple with the DistroVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistroVersion

`func (o *CondaPackageUpload) SetDistroVersion(v map[string]interface{})`

SetDistroVersion sets DistroVersion field to given value.

### HasDistroVersion

`func (o *CondaPackageUpload) HasDistroVersion() bool`

HasDistroVersion returns a boolean if a field has been set.

### GetDownloads

`func (o *CondaPackageUpload) GetDownloads() int64`

GetDownloads returns the Downloads field if non-nil, zero value otherwise.

### GetDownloadsOk

`func (o *CondaPackageUpload) GetDownloadsOk() (*int64, bool)`

GetDownloadsOk returns a tuple with the Downloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloads

`func (o *CondaPackageUpload) SetDownloads(v int64)`

SetDownloads sets Downloads field to given value.

### HasDownloads

`func (o *CondaPackageUpload) HasDownloads() bool`

HasDownloads returns a boolean if a field has been set.

### GetEpoch

`func (o *CondaPackageUpload) GetEpoch() int64`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *CondaPackageUpload) GetEpochOk() (*int64, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *CondaPackageUpload) SetEpoch(v int64)`

SetEpoch sets Epoch field to given value.

### HasEpoch

`func (o *CondaPackageUpload) HasEpoch() bool`

HasEpoch returns a boolean if a field has been set.

### GetExtension

`func (o *CondaPackageUpload) GetExtension() string`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *CondaPackageUpload) GetExtensionOk() (*string, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *CondaPackageUpload) SetExtension(v string)`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *CondaPackageUpload) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### GetFilename

`func (o *CondaPackageUpload) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *CondaPackageUpload) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *CondaPackageUpload) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *CondaPackageUpload) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetFiles

`func (o *CondaPackageUpload) GetFiles() []PackagesOwnerRepoFiles`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *CondaPackageUpload) GetFilesOk() (*[]PackagesOwnerRepoFiles, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *CondaPackageUpload) SetFiles(v []PackagesOwnerRepoFiles)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *CondaPackageUpload) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetFormat

`func (o *CondaPackageUpload) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CondaPackageUpload) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CondaPackageUpload) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CondaPackageUpload) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetFormatUrl

`func (o *CondaPackageUpload) GetFormatUrl() string`

GetFormatUrl returns the FormatUrl field if non-nil, zero value otherwise.

### GetFormatUrlOk

`func (o *CondaPackageUpload) GetFormatUrlOk() (*string, bool)`

GetFormatUrlOk returns a tuple with the FormatUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatUrl

`func (o *CondaPackageUpload) SetFormatUrl(v string)`

SetFormatUrl sets FormatUrl field to given value.

### HasFormatUrl

`func (o *CondaPackageUpload) HasFormatUrl() bool`

HasFormatUrl returns a boolean if a field has been set.

### GetIdentifierPerm

`func (o *CondaPackageUpload) GetIdentifierPerm() string`

GetIdentifierPerm returns the IdentifierPerm field if non-nil, zero value otherwise.

### GetIdentifierPermOk

`func (o *CondaPackageUpload) GetIdentifierPermOk() (*string, bool)`

GetIdentifierPermOk returns a tuple with the IdentifierPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierPerm

`func (o *CondaPackageUpload) SetIdentifierPerm(v string)`

SetIdentifierPerm sets IdentifierPerm field to given value.

### HasIdentifierPerm

`func (o *CondaPackageUpload) HasIdentifierPerm() bool`

HasIdentifierPerm returns a boolean if a field has been set.

### GetIndexed

`func (o *CondaPackageUpload) GetIndexed() bool`

GetIndexed returns the Indexed field if non-nil, zero value otherwise.

### GetIndexedOk

`func (o *CondaPackageUpload) GetIndexedOk() (*bool, bool)`

GetIndexedOk returns a tuple with the Indexed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexed

`func (o *CondaPackageUpload) SetIndexed(v bool)`

SetIndexed sets Indexed field to given value.

### HasIndexed

`func (o *CondaPackageUpload) HasIndexed() bool`

HasIndexed returns a boolean if a field has been set.

### GetIsDownloadable

`func (o *CondaPackageUpload) GetIsDownloadable() bool`

GetIsDownloadable returns the IsDownloadable field if non-nil, zero value otherwise.

### GetIsDownloadableOk

`func (o *CondaPackageUpload) GetIsDownloadableOk() (*bool, bool)`

GetIsDownloadableOk returns a tuple with the IsDownloadable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDownloadable

`func (o *CondaPackageUpload) SetIsDownloadable(v bool)`

SetIsDownloadable sets IsDownloadable field to given value.

### HasIsDownloadable

`func (o *CondaPackageUpload) HasIsDownloadable() bool`

HasIsDownloadable returns a boolean if a field has been set.

### GetIsQuarantined

`func (o *CondaPackageUpload) GetIsQuarantined() bool`

GetIsQuarantined returns the IsQuarantined field if non-nil, zero value otherwise.

### GetIsQuarantinedOk

`func (o *CondaPackageUpload) GetIsQuarantinedOk() (*bool, bool)`

GetIsQuarantinedOk returns a tuple with the IsQuarantined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQuarantined

`func (o *CondaPackageUpload) SetIsQuarantined(v bool)`

SetIsQuarantined sets IsQuarantined field to given value.

### HasIsQuarantined

`func (o *CondaPackageUpload) HasIsQuarantined() bool`

HasIsQuarantined returns a boolean if a field has been set.

### GetIsSyncAwaiting

`func (o *CondaPackageUpload) GetIsSyncAwaiting() bool`

GetIsSyncAwaiting returns the IsSyncAwaiting field if non-nil, zero value otherwise.

### GetIsSyncAwaitingOk

`func (o *CondaPackageUpload) GetIsSyncAwaitingOk() (*bool, bool)`

GetIsSyncAwaitingOk returns a tuple with the IsSyncAwaiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncAwaiting

`func (o *CondaPackageUpload) SetIsSyncAwaiting(v bool)`

SetIsSyncAwaiting sets IsSyncAwaiting field to given value.

### HasIsSyncAwaiting

`func (o *CondaPackageUpload) HasIsSyncAwaiting() bool`

HasIsSyncAwaiting returns a boolean if a field has been set.

### GetIsSyncCompleted

`func (o *CondaPackageUpload) GetIsSyncCompleted() bool`

GetIsSyncCompleted returns the IsSyncCompleted field if non-nil, zero value otherwise.

### GetIsSyncCompletedOk

`func (o *CondaPackageUpload) GetIsSyncCompletedOk() (*bool, bool)`

GetIsSyncCompletedOk returns a tuple with the IsSyncCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncCompleted

`func (o *CondaPackageUpload) SetIsSyncCompleted(v bool)`

SetIsSyncCompleted sets IsSyncCompleted field to given value.

### HasIsSyncCompleted

`func (o *CondaPackageUpload) HasIsSyncCompleted() bool`

HasIsSyncCompleted returns a boolean if a field has been set.

### GetIsSyncFailed

`func (o *CondaPackageUpload) GetIsSyncFailed() bool`

GetIsSyncFailed returns the IsSyncFailed field if non-nil, zero value otherwise.

### GetIsSyncFailedOk

`func (o *CondaPackageUpload) GetIsSyncFailedOk() (*bool, bool)`

GetIsSyncFailedOk returns a tuple with the IsSyncFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncFailed

`func (o *CondaPackageUpload) SetIsSyncFailed(v bool)`

SetIsSyncFailed sets IsSyncFailed field to given value.

### HasIsSyncFailed

`func (o *CondaPackageUpload) HasIsSyncFailed() bool`

HasIsSyncFailed returns a boolean if a field has been set.

### GetIsSyncInFlight

`func (o *CondaPackageUpload) GetIsSyncInFlight() bool`

GetIsSyncInFlight returns the IsSyncInFlight field if non-nil, zero value otherwise.

### GetIsSyncInFlightOk

`func (o *CondaPackageUpload) GetIsSyncInFlightOk() (*bool, bool)`

GetIsSyncInFlightOk returns a tuple with the IsSyncInFlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncInFlight

`func (o *CondaPackageUpload) SetIsSyncInFlight(v bool)`

SetIsSyncInFlight sets IsSyncInFlight field to given value.

### HasIsSyncInFlight

`func (o *CondaPackageUpload) HasIsSyncInFlight() bool`

HasIsSyncInFlight returns a boolean if a field has been set.

### GetIsSyncInProgress

`func (o *CondaPackageUpload) GetIsSyncInProgress() bool`

GetIsSyncInProgress returns the IsSyncInProgress field if non-nil, zero value otherwise.

### GetIsSyncInProgressOk

`func (o *CondaPackageUpload) GetIsSyncInProgressOk() (*bool, bool)`

GetIsSyncInProgressOk returns a tuple with the IsSyncInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncInProgress

`func (o *CondaPackageUpload) SetIsSyncInProgress(v bool)`

SetIsSyncInProgress sets IsSyncInProgress field to given value.

### HasIsSyncInProgress

`func (o *CondaPackageUpload) HasIsSyncInProgress() bool`

HasIsSyncInProgress returns a boolean if a field has been set.

### GetLicense

`func (o *CondaPackageUpload) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *CondaPackageUpload) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *CondaPackageUpload) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *CondaPackageUpload) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetName

`func (o *CondaPackageUpload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CondaPackageUpload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CondaPackageUpload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CondaPackageUpload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *CondaPackageUpload) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CondaPackageUpload) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CondaPackageUpload) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *CondaPackageUpload) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetNamespaceUrl

`func (o *CondaPackageUpload) GetNamespaceUrl() string`

GetNamespaceUrl returns the NamespaceUrl field if non-nil, zero value otherwise.

### GetNamespaceUrlOk

`func (o *CondaPackageUpload) GetNamespaceUrlOk() (*string, bool)`

GetNamespaceUrlOk returns a tuple with the NamespaceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceUrl

`func (o *CondaPackageUpload) SetNamespaceUrl(v string)`

SetNamespaceUrl sets NamespaceUrl field to given value.

### HasNamespaceUrl

`func (o *CondaPackageUpload) HasNamespaceUrl() bool`

HasNamespaceUrl returns a boolean if a field has been set.

### GetNumFiles

`func (o *CondaPackageUpload) GetNumFiles() int64`

GetNumFiles returns the NumFiles field if non-nil, zero value otherwise.

### GetNumFilesOk

`func (o *CondaPackageUpload) GetNumFilesOk() (*int64, bool)`

GetNumFilesOk returns a tuple with the NumFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFiles

`func (o *CondaPackageUpload) SetNumFiles(v int64)`

SetNumFiles sets NumFiles field to given value.

### HasNumFiles

`func (o *CondaPackageUpload) HasNumFiles() bool`

HasNumFiles returns a boolean if a field has been set.

### GetPackageType

`func (o *CondaPackageUpload) GetPackageType() int64`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *CondaPackageUpload) GetPackageTypeOk() (*int64, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *CondaPackageUpload) SetPackageType(v int64)`

SetPackageType sets PackageType field to given value.

### HasPackageType

`func (o *CondaPackageUpload) HasPackageType() bool`

HasPackageType returns a boolean if a field has been set.

### GetRelease

`func (o *CondaPackageUpload) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *CondaPackageUpload) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *CondaPackageUpload) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *CondaPackageUpload) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetRepository

`func (o *CondaPackageUpload) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *CondaPackageUpload) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *CondaPackageUpload) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *CondaPackageUpload) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetRepositoryUrl

`func (o *CondaPackageUpload) GetRepositoryUrl() string`

GetRepositoryUrl returns the RepositoryUrl field if non-nil, zero value otherwise.

### GetRepositoryUrlOk

`func (o *CondaPackageUpload) GetRepositoryUrlOk() (*string, bool)`

GetRepositoryUrlOk returns a tuple with the RepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryUrl

`func (o *CondaPackageUpload) SetRepositoryUrl(v string)`

SetRepositoryUrl sets RepositoryUrl field to given value.

### HasRepositoryUrl

`func (o *CondaPackageUpload) HasRepositoryUrl() bool`

HasRepositoryUrl returns a boolean if a field has been set.

### GetSecurityScanCompletedAt

`func (o *CondaPackageUpload) GetSecurityScanCompletedAt() string`

GetSecurityScanCompletedAt returns the SecurityScanCompletedAt field if non-nil, zero value otherwise.

### GetSecurityScanCompletedAtOk

`func (o *CondaPackageUpload) GetSecurityScanCompletedAtOk() (*string, bool)`

GetSecurityScanCompletedAtOk returns a tuple with the SecurityScanCompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityScanCompletedAt

`func (o *CondaPackageUpload) SetSecurityScanCompletedAt(v string)`

SetSecurityScanCompletedAt sets SecurityScanCompletedAt field to given value.

### HasSecurityScanCompletedAt

`func (o *CondaPackageUpload) HasSecurityScanCompletedAt() bool`

HasSecurityScanCompletedAt returns a boolean if a field has been set.

### GetSecurityScanStartedAt

`func (o *CondaPackageUpload) GetSecurityScanStartedAt() string`

GetSecurityScanStartedAt returns the SecurityScanStartedAt field if non-nil, zero value otherwise.

### GetSecurityScanStartedAtOk

`func (o *CondaPackageUpload) GetSecurityScanStartedAtOk() (*string, bool)`

GetSecurityScanStartedAtOk returns a tuple with the SecurityScanStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityScanStartedAt

`func (o *CondaPackageUpload) SetSecurityScanStartedAt(v string)`

SetSecurityScanStartedAt sets SecurityScanStartedAt field to given value.

### HasSecurityScanStartedAt

`func (o *CondaPackageUpload) HasSecurityScanStartedAt() bool`

HasSecurityScanStartedAt returns a boolean if a field has been set.

### GetSecurityScanStatus

`func (o *CondaPackageUpload) GetSecurityScanStatus() string`

GetSecurityScanStatus returns the SecurityScanStatus field if non-nil, zero value otherwise.

### GetSecurityScanStatusOk

`func (o *CondaPackageUpload) GetSecurityScanStatusOk() (*string, bool)`

GetSecurityScanStatusOk returns a tuple with the SecurityScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityScanStatus

`func (o *CondaPackageUpload) SetSecurityScanStatus(v string)`

SetSecurityScanStatus sets SecurityScanStatus field to given value.

### HasSecurityScanStatus

`func (o *CondaPackageUpload) HasSecurityScanStatus() bool`

HasSecurityScanStatus returns a boolean if a field has been set.

### GetSecurityScanStatusUpdatedAt

`func (o *CondaPackageUpload) GetSecurityScanStatusUpdatedAt() string`

GetSecurityScanStatusUpdatedAt returns the SecurityScanStatusUpdatedAt field if non-nil, zero value otherwise.

### GetSecurityScanStatusUpdatedAtOk

`func (o *CondaPackageUpload) GetSecurityScanStatusUpdatedAtOk() (*string, bool)`

GetSecurityScanStatusUpdatedAtOk returns a tuple with the SecurityScanStatusUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityScanStatusUpdatedAt

`func (o *CondaPackageUpload) SetSecurityScanStatusUpdatedAt(v string)`

SetSecurityScanStatusUpdatedAt sets SecurityScanStatusUpdatedAt field to given value.

### HasSecurityScanStatusUpdatedAt

`func (o *CondaPackageUpload) HasSecurityScanStatusUpdatedAt() bool`

HasSecurityScanStatusUpdatedAt returns a boolean if a field has been set.

### GetSelfHtmlUrl

`func (o *CondaPackageUpload) GetSelfHtmlUrl() string`

GetSelfHtmlUrl returns the SelfHtmlUrl field if non-nil, zero value otherwise.

### GetSelfHtmlUrlOk

`func (o *CondaPackageUpload) GetSelfHtmlUrlOk() (*string, bool)`

GetSelfHtmlUrlOk returns a tuple with the SelfHtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfHtmlUrl

`func (o *CondaPackageUpload) SetSelfHtmlUrl(v string)`

SetSelfHtmlUrl sets SelfHtmlUrl field to given value.

### HasSelfHtmlUrl

`func (o *CondaPackageUpload) HasSelfHtmlUrl() bool`

HasSelfHtmlUrl returns a boolean if a field has been set.

### GetSelfUrl

`func (o *CondaPackageUpload) GetSelfUrl() string`

GetSelfUrl returns the SelfUrl field if non-nil, zero value otherwise.

### GetSelfUrlOk

`func (o *CondaPackageUpload) GetSelfUrlOk() (*string, bool)`

GetSelfUrlOk returns a tuple with the SelfUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfUrl

`func (o *CondaPackageUpload) SetSelfUrl(v string)`

SetSelfUrl sets SelfUrl field to given value.

### HasSelfUrl

`func (o *CondaPackageUpload) HasSelfUrl() bool`

HasSelfUrl returns a boolean if a field has been set.

### GetSignatureUrl

`func (o *CondaPackageUpload) GetSignatureUrl() string`

GetSignatureUrl returns the SignatureUrl field if non-nil, zero value otherwise.

### GetSignatureUrlOk

`func (o *CondaPackageUpload) GetSignatureUrlOk() (*string, bool)`

GetSignatureUrlOk returns a tuple with the SignatureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureUrl

`func (o *CondaPackageUpload) SetSignatureUrl(v string)`

SetSignatureUrl sets SignatureUrl field to given value.

### HasSignatureUrl

`func (o *CondaPackageUpload) HasSignatureUrl() bool`

HasSignatureUrl returns a boolean if a field has been set.

### GetSize

`func (o *CondaPackageUpload) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CondaPackageUpload) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CondaPackageUpload) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *CondaPackageUpload) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSlug

`func (o *CondaPackageUpload) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *CondaPackageUpload) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *CondaPackageUpload) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *CondaPackageUpload) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSlugPerm

`func (o *CondaPackageUpload) GetSlugPerm() string`

GetSlugPerm returns the SlugPerm field if non-nil, zero value otherwise.

### GetSlugPermOk

`func (o *CondaPackageUpload) GetSlugPermOk() (*string, bool)`

GetSlugPermOk returns a tuple with the SlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugPerm

`func (o *CondaPackageUpload) SetSlugPerm(v string)`

SetSlugPerm sets SlugPerm field to given value.

### HasSlugPerm

`func (o *CondaPackageUpload) HasSlugPerm() bool`

HasSlugPerm returns a boolean if a field has been set.

### GetStage

`func (o *CondaPackageUpload) GetStage() int64`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *CondaPackageUpload) GetStageOk() (*int64, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *CondaPackageUpload) SetStage(v int64)`

SetStage sets Stage field to given value.

### HasStage

`func (o *CondaPackageUpload) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetStageStr

`func (o *CondaPackageUpload) GetStageStr() string`

GetStageStr returns the StageStr field if non-nil, zero value otherwise.

### GetStageStrOk

`func (o *CondaPackageUpload) GetStageStrOk() (*string, bool)`

GetStageStrOk returns a tuple with the StageStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageStr

`func (o *CondaPackageUpload) SetStageStr(v string)`

SetStageStr sets StageStr field to given value.

### HasStageStr

`func (o *CondaPackageUpload) HasStageStr() bool`

HasStageStr returns a boolean if a field has been set.

### GetStageUpdatedAt

`func (o *CondaPackageUpload) GetStageUpdatedAt() string`

GetStageUpdatedAt returns the StageUpdatedAt field if non-nil, zero value otherwise.

### GetStageUpdatedAtOk

`func (o *CondaPackageUpload) GetStageUpdatedAtOk() (*string, bool)`

GetStageUpdatedAtOk returns a tuple with the StageUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageUpdatedAt

`func (o *CondaPackageUpload) SetStageUpdatedAt(v string)`

SetStageUpdatedAt sets StageUpdatedAt field to given value.

### HasStageUpdatedAt

`func (o *CondaPackageUpload) HasStageUpdatedAt() bool`

HasStageUpdatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *CondaPackageUpload) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CondaPackageUpload) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CondaPackageUpload) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CondaPackageUpload) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusReason

`func (o *CondaPackageUpload) GetStatusReason() string`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *CondaPackageUpload) GetStatusReasonOk() (*string, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *CondaPackageUpload) SetStatusReason(v string)`

SetStatusReason sets StatusReason field to given value.

### HasStatusReason

`func (o *CondaPackageUpload) HasStatusReason() bool`

HasStatusReason returns a boolean if a field has been set.

### GetStatusStr

`func (o *CondaPackageUpload) GetStatusStr() string`

GetStatusStr returns the StatusStr field if non-nil, zero value otherwise.

### GetStatusStrOk

`func (o *CondaPackageUpload) GetStatusStrOk() (*string, bool)`

GetStatusStrOk returns a tuple with the StatusStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusStr

`func (o *CondaPackageUpload) SetStatusStr(v string)`

SetStatusStr sets StatusStr field to given value.

### HasStatusStr

`func (o *CondaPackageUpload) HasStatusStr() bool`

HasStatusStr returns a boolean if a field has been set.

### GetStatusUpdatedAt

`func (o *CondaPackageUpload) GetStatusUpdatedAt() string`

GetStatusUpdatedAt returns the StatusUpdatedAt field if non-nil, zero value otherwise.

### GetStatusUpdatedAtOk

`func (o *CondaPackageUpload) GetStatusUpdatedAtOk() (*string, bool)`

GetStatusUpdatedAtOk returns a tuple with the StatusUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusUpdatedAt

`func (o *CondaPackageUpload) SetStatusUpdatedAt(v string)`

SetStatusUpdatedAt sets StatusUpdatedAt field to given value.

### HasStatusUpdatedAt

`func (o *CondaPackageUpload) HasStatusUpdatedAt() bool`

HasStatusUpdatedAt returns a boolean if a field has been set.

### GetStatusUrl

`func (o *CondaPackageUpload) GetStatusUrl() string`

GetStatusUrl returns the StatusUrl field if non-nil, zero value otherwise.

### GetStatusUrlOk

`func (o *CondaPackageUpload) GetStatusUrlOk() (*string, bool)`

GetStatusUrlOk returns a tuple with the StatusUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusUrl

`func (o *CondaPackageUpload) SetStatusUrl(v string)`

SetStatusUrl sets StatusUrl field to given value.

### HasStatusUrl

`func (o *CondaPackageUpload) HasStatusUrl() bool`

HasStatusUrl returns a boolean if a field has been set.

### GetSubtype

`func (o *CondaPackageUpload) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *CondaPackageUpload) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *CondaPackageUpload) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *CondaPackageUpload) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetSummary

`func (o *CondaPackageUpload) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *CondaPackageUpload) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *CondaPackageUpload) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *CondaPackageUpload) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetSyncFinishedAt

`func (o *CondaPackageUpload) GetSyncFinishedAt() string`

GetSyncFinishedAt returns the SyncFinishedAt field if non-nil, zero value otherwise.

### GetSyncFinishedAtOk

`func (o *CondaPackageUpload) GetSyncFinishedAtOk() (*string, bool)`

GetSyncFinishedAtOk returns a tuple with the SyncFinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncFinishedAt

`func (o *CondaPackageUpload) SetSyncFinishedAt(v string)`

SetSyncFinishedAt sets SyncFinishedAt field to given value.

### HasSyncFinishedAt

`func (o *CondaPackageUpload) HasSyncFinishedAt() bool`

HasSyncFinishedAt returns a boolean if a field has been set.

### GetSyncProgress

`func (o *CondaPackageUpload) GetSyncProgress() int64`

GetSyncProgress returns the SyncProgress field if non-nil, zero value otherwise.

### GetSyncProgressOk

`func (o *CondaPackageUpload) GetSyncProgressOk() (*int64, bool)`

GetSyncProgressOk returns a tuple with the SyncProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncProgress

`func (o *CondaPackageUpload) SetSyncProgress(v int64)`

SetSyncProgress sets SyncProgress field to given value.

### HasSyncProgress

`func (o *CondaPackageUpload) HasSyncProgress() bool`

HasSyncProgress returns a boolean if a field has been set.

### GetTagsImmutable

`func (o *CondaPackageUpload) GetTagsImmutable() map[string]interface{}`

GetTagsImmutable returns the TagsImmutable field if non-nil, zero value otherwise.

### GetTagsImmutableOk

`func (o *CondaPackageUpload) GetTagsImmutableOk() (*map[string]interface{}, bool)`

GetTagsImmutableOk returns a tuple with the TagsImmutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsImmutable

`func (o *CondaPackageUpload) SetTagsImmutable(v map[string]interface{})`

SetTagsImmutable sets TagsImmutable field to given value.

### HasTagsImmutable

`func (o *CondaPackageUpload) HasTagsImmutable() bool`

HasTagsImmutable returns a boolean if a field has been set.

### GetTypeDisplay

`func (o *CondaPackageUpload) GetTypeDisplay() string`

GetTypeDisplay returns the TypeDisplay field if non-nil, zero value otherwise.

### GetTypeDisplayOk

`func (o *CondaPackageUpload) GetTypeDisplayOk() (*string, bool)`

GetTypeDisplayOk returns a tuple with the TypeDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeDisplay

`func (o *CondaPackageUpload) SetTypeDisplay(v string)`

SetTypeDisplay sets TypeDisplay field to given value.

### HasTypeDisplay

`func (o *CondaPackageUpload) HasTypeDisplay() bool`

HasTypeDisplay returns a boolean if a field has been set.

### GetUploadedAt

`func (o *CondaPackageUpload) GetUploadedAt() string`

GetUploadedAt returns the UploadedAt field if non-nil, zero value otherwise.

### GetUploadedAtOk

`func (o *CondaPackageUpload) GetUploadedAtOk() (*string, bool)`

GetUploadedAtOk returns a tuple with the UploadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedAt

`func (o *CondaPackageUpload) SetUploadedAt(v string)`

SetUploadedAt sets UploadedAt field to given value.

### HasUploadedAt

`func (o *CondaPackageUpload) HasUploadedAt() bool`

HasUploadedAt returns a boolean if a field has been set.

### GetUploader

`func (o *CondaPackageUpload) GetUploader() string`

GetUploader returns the Uploader field if non-nil, zero value otherwise.

### GetUploaderOk

`func (o *CondaPackageUpload) GetUploaderOk() (*string, bool)`

GetUploaderOk returns a tuple with the Uploader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploader

`func (o *CondaPackageUpload) SetUploader(v string)`

SetUploader sets Uploader field to given value.

### HasUploader

`func (o *CondaPackageUpload) HasUploader() bool`

HasUploader returns a boolean if a field has been set.

### GetUploaderUrl

`func (o *CondaPackageUpload) GetUploaderUrl() string`

GetUploaderUrl returns the UploaderUrl field if non-nil, zero value otherwise.

### GetUploaderUrlOk

`func (o *CondaPackageUpload) GetUploaderUrlOk() (*string, bool)`

GetUploaderUrlOk returns a tuple with the UploaderUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploaderUrl

`func (o *CondaPackageUpload) SetUploaderUrl(v string)`

SetUploaderUrl sets UploaderUrl field to given value.

### HasUploaderUrl

`func (o *CondaPackageUpload) HasUploaderUrl() bool`

HasUploaderUrl returns a boolean if a field has been set.

### GetVersion

`func (o *CondaPackageUpload) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CondaPackageUpload) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CondaPackageUpload) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CondaPackageUpload) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVersionOrig

`func (o *CondaPackageUpload) GetVersionOrig() string`

GetVersionOrig returns the VersionOrig field if non-nil, zero value otherwise.

### GetVersionOrigOk

`func (o *CondaPackageUpload) GetVersionOrigOk() (*string, bool)`

GetVersionOrigOk returns a tuple with the VersionOrig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionOrig

`func (o *CondaPackageUpload) SetVersionOrig(v string)`

SetVersionOrig sets VersionOrig field to given value.

### HasVersionOrig

`func (o *CondaPackageUpload) HasVersionOrig() bool`

HasVersionOrig returns a boolean if a field has been set.

### GetVulnerabilityScanResultsUrl

`func (o *CondaPackageUpload) GetVulnerabilityScanResultsUrl() string`

GetVulnerabilityScanResultsUrl returns the VulnerabilityScanResultsUrl field if non-nil, zero value otherwise.

### GetVulnerabilityScanResultsUrlOk

`func (o *CondaPackageUpload) GetVulnerabilityScanResultsUrlOk() (*string, bool)`

GetVulnerabilityScanResultsUrlOk returns a tuple with the VulnerabilityScanResultsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityScanResultsUrl

`func (o *CondaPackageUpload) SetVulnerabilityScanResultsUrl(v string)`

SetVulnerabilityScanResultsUrl sets VulnerabilityScanResultsUrl field to given value.

### HasVulnerabilityScanResultsUrl

`func (o *CondaPackageUpload) HasVulnerabilityScanResultsUrl() bool`

HasVulnerabilityScanResultsUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


