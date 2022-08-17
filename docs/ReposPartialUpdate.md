# ReposPartialUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextualAuthRealm** | Pointer to **bool** | If checked, missing credentials for this repository where basic authentication is required shall present an enriched value in the &#39;WWW-Authenticate&#39; header containing the namespace and repository. This can be useful for tooling such as SBT where the authentication realm is used to distinguish and disambiguate credentials. | [optional] 
**CopyOwn** | Pointer to **bool** | If checked, users can copy any of their own packages that they have uploaded, assuming that they still have write privilege for the repository. This takes precedence over privileges configured in the &#39;Access Controls&#39; section of the repository, and any inherited from the org. | [optional] 
**CopyPackages** | Pointer to **string** | This defines the minimum level of privilege required for a user to copy packages. Unless the package was uploaded by that user, in which the permission may be overridden by the user-specific copy setting. | [optional] 
**DefaultPrivilege** | Pointer to **string** | This defines the default level of privilege that all of your organization members have for this repository. This does not include collaborators, but applies to any member of the org regardless of their own membership role (i.e. it applies to owners, managers and members). Be careful if setting this to admin, because any member will be able to change settings. | [optional] 
**DeleteOwn** | Pointer to **bool** | If checked, users can delete any of their own packages that they have uploaded, assuming that they still have write privilege for the repository. This takes precedence over privileges configured in the &#39;Access Controls&#39; section of the repository, and any inherited from the org. | [optional] 
**DeletePackages** | Pointer to **string** | This defines the minimum level of privilege required for a user to delete packages. Unless the package was uploaded by that user, in which the permission may be overridden by the user-specific delete setting. | [optional] 
**Description** | Pointer to **string** | A description of the repository&#39;s purpose/contents. | [optional] 
**DockerRefreshTokensEnabled** | Pointer to **bool** | If checked, refresh tokens will be issued in addition to access tokens for Docker authentication. This allows unlimited extension of the lifetime of access tokens. | [optional] 
**IndexFiles** | Pointer to **bool** | If checked, files contained in packages will be indexed, which increase the synchronisation time required for packages. Note that it is recommended you keep this enabled unless the synchronisation time is significantly impacted. | [optional] 
**MoveOwn** | Pointer to **bool** | If checked, users can move any of their own packages that they have uploaded, assuming that they still have write privilege for the repository. This takes precedence over privileges configured in the &#39;Access Controls&#39; section of the repository, and any inherited from the org. | [optional] 
**MovePackages** | Pointer to **string** | This defines the minimum level of privilege required for a user to move packages. Unless the package was uploaded by that user, in which the permission may be overridden by the user-specific move setting. | [optional] 
**Name** | Pointer to **string** | A descriptive name for the repository. | [optional] 
**ProxyNpmjs** | Pointer to **bool** | If checked, Npm packages that are not in the repository when requested by clients will automatically be proxied from the public npmjs.org registry. If there is at least one version for a package, others will not be proxied. | [optional] 
**ProxyPypi** | Pointer to **bool** | If checked, Python packages that are not in the repository when requested by clients will automatically be proxied from the public pypi.python.org registry. If there is at least one version for a package, others will not be proxied. | [optional] 
**RawPackageIndexEnabled** | Pointer to **bool** | If checked, HTML and JSON indexes will be generated that list all available raw packages in the repository. | [optional] 
**RawPackageIndexSignaturesEnabled** | Pointer to **bool** | If checked, the HTML and JSON indexes will display raw package GPG signatures alongside the index packages. | [optional] 
**ReplacePackages** | Pointer to **string** | This defines the minimum level of privilege required for a user to republish packages. Unless the package was uploaded by that user, in which the permission may be overridden by the user-specific republish setting. Please note that the user still requires the privilege to delete packages that will be replaced by the new package; otherwise the republish will fail. | [optional] 
**ReplacePackagesByDefault** | Pointer to **bool** | If checked, uploaded packages will overwrite/replace any others with the same attributes (e.g. same version) by default. This only applies if the user has the required privilege for the republishing AND has the required privilege to delete existing packages that they don&#39;t own. | [optional] 
**RepositoryTypeStr** | Pointer to **string** | The repository type changes how it is accessed and billed. Private repositories can only be used on paid plans, but are visible only to you or authorised delegates. Public repositories are free to use on all plans and visible to all Cloudsmith users. | [optional] 
**ResyncOwn** | Pointer to **bool** | If checked, users can resync any of their own packages that they have uploaded, assuming that they still have write privilege for the repository. This takes precedence over privileges configured in the &#39;Access Controls&#39; section of the repository, and any inherited from the org. | [optional] 
**ResyncPackages** | Pointer to **string** | This defines the minimum level of privilege required for a user to resync packages. Unless the package was uploaded by that user, in which the permission may be overridden by the user-specific resync setting. | [optional] 
**ScanOwn** | Pointer to **bool** | If checked, users can scan any of their own packages that they have uploaded, assuming that they still have write privilege for the repository. This takes precedence over privileges configured in the &#39;Access Controls&#39; section of the repository, and any inherited from the org. | [optional] 
**ScanPackages** | Pointer to **string** | This defines the minimum level of privilege required for a user to scan packages. Unless the package was uploaded by that user, in which the permission may be overridden by the user-specific scan setting. | [optional] 
**ShowSetupAll** | Pointer to **bool** | If checked, the Set Me Up help for all formats will always be shown, even if you don&#39;t have packages of that type uploaded. Otherwise, help will only be shown for packages that are in the repository. For example, if you have uploaded only NuGet packages, then the Set Me Up help for NuGet packages will be shown only. | [optional] 
**Slug** | Pointer to **string** | The slug identifies the repository in URIs. | [optional] 
**StrictNpmValidation** | Pointer to **bool** | If checked, npm packages will be validated strictly to ensure the package matches specifcation. You can turn this off if you have packages that are old or otherwise mildly off-spec, but we can&#39;t guarantee the packages will work with npm-cli or other tooling correctly. Turn off at your own risk! | [optional] 
**UseDebianLabels** | Pointer to **bool** | If checked, a &#39;Label&#39; field will be present in Debian-based repositories. It will contain a string that identifies the entitlement token used to authenticate the repository, in the form of &#39;source&#x3D;t-&lt;identifier&gt;&#39;; or &#39;source&#x3D;none&#39; if no token was used. You can use this to help with pinning. | [optional] 
**UseDefaultCargoUpstream** | Pointer to **bool** | If checked, dependencies of uploaded Cargo crates which do not set an explicit value for \&quot;registry\&quot; will be assumed to be available from crates.io. If unchecked, dependencies with unspecified \&quot;registry\&quot; values will be assumed to be available in the registry being uploaded to. Uncheck this if you want to ensure that dependencies are only ever installed from Cloudsmith unless explicitly specified as belong to another registry. | [optional] 
**UseNoarchPackages** | Pointer to **bool** | If checked, noarch packages (if supported) are enabled in installations/configurations. A noarch package is one that is not tied to specific system architecture (like i686). | [optional] 
**UseSourcePackages** | Pointer to **bool** | If checked, source packages (if supported) are enabled in installations/configurations. A source package is one that contains source code rather than built binaries. | [optional] 
**UseVulnerabilityScanning** | Pointer to **bool** | If checked, vulnerability scanning will be enabled for all supported packages within this repository. | [optional] 
**UserEntitlementsEnabled** | Pointer to **bool** | If checked, users can use and manage their own user-specific entitlement token for the repository (if private). Otherwise, user-specific entitlements are disabled for all users. | [optional] 
**ViewStatistics** | Pointer to **string** | This defines the minimum level of privilege required for a user to view repository statistics, to include entitlement-based usage, if applciable. If a user does not have the permission, they won&#39;t be able to view any statistics, either via the UI, API or CLI. | [optional] 

## Methods

### NewReposPartialUpdate

`func NewReposPartialUpdate() *ReposPartialUpdate`

NewReposPartialUpdate instantiates a new ReposPartialUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposPartialUpdateWithDefaults

`func NewReposPartialUpdateWithDefaults() *ReposPartialUpdate`

NewReposPartialUpdateWithDefaults instantiates a new ReposPartialUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContextualAuthRealm

`func (o *ReposPartialUpdate) GetContextualAuthRealm() bool`

GetContextualAuthRealm returns the ContextualAuthRealm field if non-nil, zero value otherwise.

### GetContextualAuthRealmOk

`func (o *ReposPartialUpdate) GetContextualAuthRealmOk() (*bool, bool)`

GetContextualAuthRealmOk returns a tuple with the ContextualAuthRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextualAuthRealm

`func (o *ReposPartialUpdate) SetContextualAuthRealm(v bool)`

SetContextualAuthRealm sets ContextualAuthRealm field to given value.

### HasContextualAuthRealm

`func (o *ReposPartialUpdate) HasContextualAuthRealm() bool`

HasContextualAuthRealm returns a boolean if a field has been set.

### GetCopyOwn

`func (o *ReposPartialUpdate) GetCopyOwn() bool`

GetCopyOwn returns the CopyOwn field if non-nil, zero value otherwise.

### GetCopyOwnOk

`func (o *ReposPartialUpdate) GetCopyOwnOk() (*bool, bool)`

GetCopyOwnOk returns a tuple with the CopyOwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyOwn

`func (o *ReposPartialUpdate) SetCopyOwn(v bool)`

SetCopyOwn sets CopyOwn field to given value.

### HasCopyOwn

`func (o *ReposPartialUpdate) HasCopyOwn() bool`

HasCopyOwn returns a boolean if a field has been set.

### GetCopyPackages

`func (o *ReposPartialUpdate) GetCopyPackages() string`

GetCopyPackages returns the CopyPackages field if non-nil, zero value otherwise.

### GetCopyPackagesOk

`func (o *ReposPartialUpdate) GetCopyPackagesOk() (*string, bool)`

GetCopyPackagesOk returns a tuple with the CopyPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyPackages

`func (o *ReposPartialUpdate) SetCopyPackages(v string)`

SetCopyPackages sets CopyPackages field to given value.

### HasCopyPackages

`func (o *ReposPartialUpdate) HasCopyPackages() bool`

HasCopyPackages returns a boolean if a field has been set.

### GetDefaultPrivilege

`func (o *ReposPartialUpdate) GetDefaultPrivilege() string`

GetDefaultPrivilege returns the DefaultPrivilege field if non-nil, zero value otherwise.

### GetDefaultPrivilegeOk

`func (o *ReposPartialUpdate) GetDefaultPrivilegeOk() (*string, bool)`

GetDefaultPrivilegeOk returns a tuple with the DefaultPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPrivilege

`func (o *ReposPartialUpdate) SetDefaultPrivilege(v string)`

SetDefaultPrivilege sets DefaultPrivilege field to given value.

### HasDefaultPrivilege

`func (o *ReposPartialUpdate) HasDefaultPrivilege() bool`

HasDefaultPrivilege returns a boolean if a field has been set.

### GetDeleteOwn

`func (o *ReposPartialUpdate) GetDeleteOwn() bool`

GetDeleteOwn returns the DeleteOwn field if non-nil, zero value otherwise.

### GetDeleteOwnOk

`func (o *ReposPartialUpdate) GetDeleteOwnOk() (*bool, bool)`

GetDeleteOwnOk returns a tuple with the DeleteOwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOwn

`func (o *ReposPartialUpdate) SetDeleteOwn(v bool)`

SetDeleteOwn sets DeleteOwn field to given value.

### HasDeleteOwn

`func (o *ReposPartialUpdate) HasDeleteOwn() bool`

HasDeleteOwn returns a boolean if a field has been set.

### GetDeletePackages

`func (o *ReposPartialUpdate) GetDeletePackages() string`

GetDeletePackages returns the DeletePackages field if non-nil, zero value otherwise.

### GetDeletePackagesOk

`func (o *ReposPartialUpdate) GetDeletePackagesOk() (*string, bool)`

GetDeletePackagesOk returns a tuple with the DeletePackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletePackages

`func (o *ReposPartialUpdate) SetDeletePackages(v string)`

SetDeletePackages sets DeletePackages field to given value.

### HasDeletePackages

`func (o *ReposPartialUpdate) HasDeletePackages() bool`

HasDeletePackages returns a boolean if a field has been set.

### GetDescription

`func (o *ReposPartialUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReposPartialUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReposPartialUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReposPartialUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDockerRefreshTokensEnabled

`func (o *ReposPartialUpdate) GetDockerRefreshTokensEnabled() bool`

GetDockerRefreshTokensEnabled returns the DockerRefreshTokensEnabled field if non-nil, zero value otherwise.

### GetDockerRefreshTokensEnabledOk

`func (o *ReposPartialUpdate) GetDockerRefreshTokensEnabledOk() (*bool, bool)`

GetDockerRefreshTokensEnabledOk returns a tuple with the DockerRefreshTokensEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerRefreshTokensEnabled

`func (o *ReposPartialUpdate) SetDockerRefreshTokensEnabled(v bool)`

SetDockerRefreshTokensEnabled sets DockerRefreshTokensEnabled field to given value.

### HasDockerRefreshTokensEnabled

`func (o *ReposPartialUpdate) HasDockerRefreshTokensEnabled() bool`

HasDockerRefreshTokensEnabled returns a boolean if a field has been set.

### GetIndexFiles

`func (o *ReposPartialUpdate) GetIndexFiles() bool`

GetIndexFiles returns the IndexFiles field if non-nil, zero value otherwise.

### GetIndexFilesOk

`func (o *ReposPartialUpdate) GetIndexFilesOk() (*bool, bool)`

GetIndexFilesOk returns a tuple with the IndexFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexFiles

`func (o *ReposPartialUpdate) SetIndexFiles(v bool)`

SetIndexFiles sets IndexFiles field to given value.

### HasIndexFiles

`func (o *ReposPartialUpdate) HasIndexFiles() bool`

HasIndexFiles returns a boolean if a field has been set.

### GetMoveOwn

`func (o *ReposPartialUpdate) GetMoveOwn() bool`

GetMoveOwn returns the MoveOwn field if non-nil, zero value otherwise.

### GetMoveOwnOk

`func (o *ReposPartialUpdate) GetMoveOwnOk() (*bool, bool)`

GetMoveOwnOk returns a tuple with the MoveOwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveOwn

`func (o *ReposPartialUpdate) SetMoveOwn(v bool)`

SetMoveOwn sets MoveOwn field to given value.

### HasMoveOwn

`func (o *ReposPartialUpdate) HasMoveOwn() bool`

HasMoveOwn returns a boolean if a field has been set.

### GetMovePackages

`func (o *ReposPartialUpdate) GetMovePackages() string`

GetMovePackages returns the MovePackages field if non-nil, zero value otherwise.

### GetMovePackagesOk

`func (o *ReposPartialUpdate) GetMovePackagesOk() (*string, bool)`

GetMovePackagesOk returns a tuple with the MovePackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovePackages

`func (o *ReposPartialUpdate) SetMovePackages(v string)`

SetMovePackages sets MovePackages field to given value.

### HasMovePackages

`func (o *ReposPartialUpdate) HasMovePackages() bool`

HasMovePackages returns a boolean if a field has been set.

### GetName

`func (o *ReposPartialUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReposPartialUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReposPartialUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReposPartialUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProxyNpmjs

`func (o *ReposPartialUpdate) GetProxyNpmjs() bool`

GetProxyNpmjs returns the ProxyNpmjs field if non-nil, zero value otherwise.

### GetProxyNpmjsOk

`func (o *ReposPartialUpdate) GetProxyNpmjsOk() (*bool, bool)`

GetProxyNpmjsOk returns a tuple with the ProxyNpmjs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyNpmjs

`func (o *ReposPartialUpdate) SetProxyNpmjs(v bool)`

SetProxyNpmjs sets ProxyNpmjs field to given value.

### HasProxyNpmjs

`func (o *ReposPartialUpdate) HasProxyNpmjs() bool`

HasProxyNpmjs returns a boolean if a field has been set.

### GetProxyPypi

`func (o *ReposPartialUpdate) GetProxyPypi() bool`

GetProxyPypi returns the ProxyPypi field if non-nil, zero value otherwise.

### GetProxyPypiOk

`func (o *ReposPartialUpdate) GetProxyPypiOk() (*bool, bool)`

GetProxyPypiOk returns a tuple with the ProxyPypi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyPypi

`func (o *ReposPartialUpdate) SetProxyPypi(v bool)`

SetProxyPypi sets ProxyPypi field to given value.

### HasProxyPypi

`func (o *ReposPartialUpdate) HasProxyPypi() bool`

HasProxyPypi returns a boolean if a field has been set.

### GetRawPackageIndexEnabled

`func (o *ReposPartialUpdate) GetRawPackageIndexEnabled() bool`

GetRawPackageIndexEnabled returns the RawPackageIndexEnabled field if non-nil, zero value otherwise.

### GetRawPackageIndexEnabledOk

`func (o *ReposPartialUpdate) GetRawPackageIndexEnabledOk() (*bool, bool)`

GetRawPackageIndexEnabledOk returns a tuple with the RawPackageIndexEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawPackageIndexEnabled

`func (o *ReposPartialUpdate) SetRawPackageIndexEnabled(v bool)`

SetRawPackageIndexEnabled sets RawPackageIndexEnabled field to given value.

### HasRawPackageIndexEnabled

`func (o *ReposPartialUpdate) HasRawPackageIndexEnabled() bool`

HasRawPackageIndexEnabled returns a boolean if a field has been set.

### GetRawPackageIndexSignaturesEnabled

`func (o *ReposPartialUpdate) GetRawPackageIndexSignaturesEnabled() bool`

GetRawPackageIndexSignaturesEnabled returns the RawPackageIndexSignaturesEnabled field if non-nil, zero value otherwise.

### GetRawPackageIndexSignaturesEnabledOk

`func (o *ReposPartialUpdate) GetRawPackageIndexSignaturesEnabledOk() (*bool, bool)`

GetRawPackageIndexSignaturesEnabledOk returns a tuple with the RawPackageIndexSignaturesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawPackageIndexSignaturesEnabled

`func (o *ReposPartialUpdate) SetRawPackageIndexSignaturesEnabled(v bool)`

SetRawPackageIndexSignaturesEnabled sets RawPackageIndexSignaturesEnabled field to given value.

### HasRawPackageIndexSignaturesEnabled

`func (o *ReposPartialUpdate) HasRawPackageIndexSignaturesEnabled() bool`

HasRawPackageIndexSignaturesEnabled returns a boolean if a field has been set.

### GetReplacePackages

`func (o *ReposPartialUpdate) GetReplacePackages() string`

GetReplacePackages returns the ReplacePackages field if non-nil, zero value otherwise.

### GetReplacePackagesOk

`func (o *ReposPartialUpdate) GetReplacePackagesOk() (*string, bool)`

GetReplacePackagesOk returns a tuple with the ReplacePackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacePackages

`func (o *ReposPartialUpdate) SetReplacePackages(v string)`

SetReplacePackages sets ReplacePackages field to given value.

### HasReplacePackages

`func (o *ReposPartialUpdate) HasReplacePackages() bool`

HasReplacePackages returns a boolean if a field has been set.

### GetReplacePackagesByDefault

`func (o *ReposPartialUpdate) GetReplacePackagesByDefault() bool`

GetReplacePackagesByDefault returns the ReplacePackagesByDefault field if non-nil, zero value otherwise.

### GetReplacePackagesByDefaultOk

`func (o *ReposPartialUpdate) GetReplacePackagesByDefaultOk() (*bool, bool)`

GetReplacePackagesByDefaultOk returns a tuple with the ReplacePackagesByDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacePackagesByDefault

`func (o *ReposPartialUpdate) SetReplacePackagesByDefault(v bool)`

SetReplacePackagesByDefault sets ReplacePackagesByDefault field to given value.

### HasReplacePackagesByDefault

`func (o *ReposPartialUpdate) HasReplacePackagesByDefault() bool`

HasReplacePackagesByDefault returns a boolean if a field has been set.

### GetRepositoryTypeStr

`func (o *ReposPartialUpdate) GetRepositoryTypeStr() string`

GetRepositoryTypeStr returns the RepositoryTypeStr field if non-nil, zero value otherwise.

### GetRepositoryTypeStrOk

`func (o *ReposPartialUpdate) GetRepositoryTypeStrOk() (*string, bool)`

GetRepositoryTypeStrOk returns a tuple with the RepositoryTypeStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryTypeStr

`func (o *ReposPartialUpdate) SetRepositoryTypeStr(v string)`

SetRepositoryTypeStr sets RepositoryTypeStr field to given value.

### HasRepositoryTypeStr

`func (o *ReposPartialUpdate) HasRepositoryTypeStr() bool`

HasRepositoryTypeStr returns a boolean if a field has been set.

### GetResyncOwn

`func (o *ReposPartialUpdate) GetResyncOwn() bool`

GetResyncOwn returns the ResyncOwn field if non-nil, zero value otherwise.

### GetResyncOwnOk

`func (o *ReposPartialUpdate) GetResyncOwnOk() (*bool, bool)`

GetResyncOwnOk returns a tuple with the ResyncOwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResyncOwn

`func (o *ReposPartialUpdate) SetResyncOwn(v bool)`

SetResyncOwn sets ResyncOwn field to given value.

### HasResyncOwn

`func (o *ReposPartialUpdate) HasResyncOwn() bool`

HasResyncOwn returns a boolean if a field has been set.

### GetResyncPackages

`func (o *ReposPartialUpdate) GetResyncPackages() string`

GetResyncPackages returns the ResyncPackages field if non-nil, zero value otherwise.

### GetResyncPackagesOk

`func (o *ReposPartialUpdate) GetResyncPackagesOk() (*string, bool)`

GetResyncPackagesOk returns a tuple with the ResyncPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResyncPackages

`func (o *ReposPartialUpdate) SetResyncPackages(v string)`

SetResyncPackages sets ResyncPackages field to given value.

### HasResyncPackages

`func (o *ReposPartialUpdate) HasResyncPackages() bool`

HasResyncPackages returns a boolean if a field has been set.

### GetScanOwn

`func (o *ReposPartialUpdate) GetScanOwn() bool`

GetScanOwn returns the ScanOwn field if non-nil, zero value otherwise.

### GetScanOwnOk

`func (o *ReposPartialUpdate) GetScanOwnOk() (*bool, bool)`

GetScanOwnOk returns a tuple with the ScanOwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanOwn

`func (o *ReposPartialUpdate) SetScanOwn(v bool)`

SetScanOwn sets ScanOwn field to given value.

### HasScanOwn

`func (o *ReposPartialUpdate) HasScanOwn() bool`

HasScanOwn returns a boolean if a field has been set.

### GetScanPackages

`func (o *ReposPartialUpdate) GetScanPackages() string`

GetScanPackages returns the ScanPackages field if non-nil, zero value otherwise.

### GetScanPackagesOk

`func (o *ReposPartialUpdate) GetScanPackagesOk() (*string, bool)`

GetScanPackagesOk returns a tuple with the ScanPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanPackages

`func (o *ReposPartialUpdate) SetScanPackages(v string)`

SetScanPackages sets ScanPackages field to given value.

### HasScanPackages

`func (o *ReposPartialUpdate) HasScanPackages() bool`

HasScanPackages returns a boolean if a field has been set.

### GetShowSetupAll

`func (o *ReposPartialUpdate) GetShowSetupAll() bool`

GetShowSetupAll returns the ShowSetupAll field if non-nil, zero value otherwise.

### GetShowSetupAllOk

`func (o *ReposPartialUpdate) GetShowSetupAllOk() (*bool, bool)`

GetShowSetupAllOk returns a tuple with the ShowSetupAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowSetupAll

`func (o *ReposPartialUpdate) SetShowSetupAll(v bool)`

SetShowSetupAll sets ShowSetupAll field to given value.

### HasShowSetupAll

`func (o *ReposPartialUpdate) HasShowSetupAll() bool`

HasShowSetupAll returns a boolean if a field has been set.

### GetSlug

`func (o *ReposPartialUpdate) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ReposPartialUpdate) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ReposPartialUpdate) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ReposPartialUpdate) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetStrictNpmValidation

`func (o *ReposPartialUpdate) GetStrictNpmValidation() bool`

GetStrictNpmValidation returns the StrictNpmValidation field if non-nil, zero value otherwise.

### GetStrictNpmValidationOk

`func (o *ReposPartialUpdate) GetStrictNpmValidationOk() (*bool, bool)`

GetStrictNpmValidationOk returns a tuple with the StrictNpmValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrictNpmValidation

`func (o *ReposPartialUpdate) SetStrictNpmValidation(v bool)`

SetStrictNpmValidation sets StrictNpmValidation field to given value.

### HasStrictNpmValidation

`func (o *ReposPartialUpdate) HasStrictNpmValidation() bool`

HasStrictNpmValidation returns a boolean if a field has been set.

### GetUseDebianLabels

`func (o *ReposPartialUpdate) GetUseDebianLabels() bool`

GetUseDebianLabels returns the UseDebianLabels field if non-nil, zero value otherwise.

### GetUseDebianLabelsOk

`func (o *ReposPartialUpdate) GetUseDebianLabelsOk() (*bool, bool)`

GetUseDebianLabelsOk returns a tuple with the UseDebianLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDebianLabels

`func (o *ReposPartialUpdate) SetUseDebianLabels(v bool)`

SetUseDebianLabels sets UseDebianLabels field to given value.

### HasUseDebianLabels

`func (o *ReposPartialUpdate) HasUseDebianLabels() bool`

HasUseDebianLabels returns a boolean if a field has been set.

### GetUseDefaultCargoUpstream

`func (o *ReposPartialUpdate) GetUseDefaultCargoUpstream() bool`

GetUseDefaultCargoUpstream returns the UseDefaultCargoUpstream field if non-nil, zero value otherwise.

### GetUseDefaultCargoUpstreamOk

`func (o *ReposPartialUpdate) GetUseDefaultCargoUpstreamOk() (*bool, bool)`

GetUseDefaultCargoUpstreamOk returns a tuple with the UseDefaultCargoUpstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDefaultCargoUpstream

`func (o *ReposPartialUpdate) SetUseDefaultCargoUpstream(v bool)`

SetUseDefaultCargoUpstream sets UseDefaultCargoUpstream field to given value.

### HasUseDefaultCargoUpstream

`func (o *ReposPartialUpdate) HasUseDefaultCargoUpstream() bool`

HasUseDefaultCargoUpstream returns a boolean if a field has been set.

### GetUseNoarchPackages

`func (o *ReposPartialUpdate) GetUseNoarchPackages() bool`

GetUseNoarchPackages returns the UseNoarchPackages field if non-nil, zero value otherwise.

### GetUseNoarchPackagesOk

`func (o *ReposPartialUpdate) GetUseNoarchPackagesOk() (*bool, bool)`

GetUseNoarchPackagesOk returns a tuple with the UseNoarchPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNoarchPackages

`func (o *ReposPartialUpdate) SetUseNoarchPackages(v bool)`

SetUseNoarchPackages sets UseNoarchPackages field to given value.

### HasUseNoarchPackages

`func (o *ReposPartialUpdate) HasUseNoarchPackages() bool`

HasUseNoarchPackages returns a boolean if a field has been set.

### GetUseSourcePackages

`func (o *ReposPartialUpdate) GetUseSourcePackages() bool`

GetUseSourcePackages returns the UseSourcePackages field if non-nil, zero value otherwise.

### GetUseSourcePackagesOk

`func (o *ReposPartialUpdate) GetUseSourcePackagesOk() (*bool, bool)`

GetUseSourcePackagesOk returns a tuple with the UseSourcePackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSourcePackages

`func (o *ReposPartialUpdate) SetUseSourcePackages(v bool)`

SetUseSourcePackages sets UseSourcePackages field to given value.

### HasUseSourcePackages

`func (o *ReposPartialUpdate) HasUseSourcePackages() bool`

HasUseSourcePackages returns a boolean if a field has been set.

### GetUseVulnerabilityScanning

`func (o *ReposPartialUpdate) GetUseVulnerabilityScanning() bool`

GetUseVulnerabilityScanning returns the UseVulnerabilityScanning field if non-nil, zero value otherwise.

### GetUseVulnerabilityScanningOk

`func (o *ReposPartialUpdate) GetUseVulnerabilityScanningOk() (*bool, bool)`

GetUseVulnerabilityScanningOk returns a tuple with the UseVulnerabilityScanning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseVulnerabilityScanning

`func (o *ReposPartialUpdate) SetUseVulnerabilityScanning(v bool)`

SetUseVulnerabilityScanning sets UseVulnerabilityScanning field to given value.

### HasUseVulnerabilityScanning

`func (o *ReposPartialUpdate) HasUseVulnerabilityScanning() bool`

HasUseVulnerabilityScanning returns a boolean if a field has been set.

### GetUserEntitlementsEnabled

`func (o *ReposPartialUpdate) GetUserEntitlementsEnabled() bool`

GetUserEntitlementsEnabled returns the UserEntitlementsEnabled field if non-nil, zero value otherwise.

### GetUserEntitlementsEnabledOk

`func (o *ReposPartialUpdate) GetUserEntitlementsEnabledOk() (*bool, bool)`

GetUserEntitlementsEnabledOk returns a tuple with the UserEntitlementsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEntitlementsEnabled

`func (o *ReposPartialUpdate) SetUserEntitlementsEnabled(v bool)`

SetUserEntitlementsEnabled sets UserEntitlementsEnabled field to given value.

### HasUserEntitlementsEnabled

`func (o *ReposPartialUpdate) HasUserEntitlementsEnabled() bool`

HasUserEntitlementsEnabled returns a boolean if a field has been set.

### GetViewStatistics

`func (o *ReposPartialUpdate) GetViewStatistics() string`

GetViewStatistics returns the ViewStatistics field if non-nil, zero value otherwise.

### GetViewStatisticsOk

`func (o *ReposPartialUpdate) GetViewStatisticsOk() (*string, bool)`

GetViewStatisticsOk returns a tuple with the ViewStatistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewStatistics

`func (o *ReposPartialUpdate) SetViewStatistics(v string)`

SetViewStatistics sets ViewStatistics field to given value.

### HasViewStatistics

`func (o *ReposPartialUpdate) HasViewStatistics() bool`

HasViewStatistics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


