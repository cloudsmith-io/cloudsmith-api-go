# OrganizationCustomDomains

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackendKind** | Pointer to **NullableInt64** | The domain for a specific package format. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**CreatedBy** | Pointer to **NullableInt64** |  | [optional] [readonly] 
**DnsAliasValue** | Pointer to **NullableString** | The CNAME value to use to publish this domain publicly. | [optional] [readonly] 
**DnsCertName** | Pointer to **NullableString** | The CNAME name to create to allow us to generate a TLS cert. | [optional] [readonly] 
**DnsCertValue** | Pointer to **NullableString** | The CNAME value to use to allow us to generate a TLS cert. | [optional] [readonly] 
**DomainType** | Pointer to **int64** | The type for the custom domain. | [optional] [readonly] 
**Enabled** | Pointer to **bool** | If checked, the domain is enabled. | [optional] [readonly] 
**Host** | Pointer to **string** |  | [optional] [readonly] 
**Namespace** | Pointer to **int64** |  | [optional] [readonly] 
**Primary** | Pointer to **bool** | If checked, this domain is the preferred/primary domain in the case that there are other overlapping domains (e.g. for the same repository, package format, etc.) | [optional] [readonly] [default to true]
**RedirectRoot** | Pointer to **bool** | If checked, the domain root redirects to the repository. | [optional] [readonly] 
**RedirectRootUrl** | Pointer to **NullableString** | Where root requests should be redirected to if redirect_root is enabled. | [optional] [readonly] 
**Repository** | [**OrganizationCustomDomainNestedRepo**](OrganizationCustomDomainNestedRepo.md) |  | 
**RepositoryOnly** | Pointer to **bool** | If checked, the domain applies to a specific repository only. | [optional] [readonly] 
**SlugPerm** | Pointer to **string** |  | [optional] [readonly] 
**Validated** | Pointer to **bool** | If validated, the domain is ready for requests. | [optional] [readonly] 
**ValidatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewOrganizationCustomDomains

`func NewOrganizationCustomDomains(repository OrganizationCustomDomainNestedRepo, ) *OrganizationCustomDomains`

NewOrganizationCustomDomains instantiates a new OrganizationCustomDomains object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationCustomDomainsWithDefaults

`func NewOrganizationCustomDomainsWithDefaults() *OrganizationCustomDomains`

NewOrganizationCustomDomainsWithDefaults instantiates a new OrganizationCustomDomains object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackendKind

`func (o *OrganizationCustomDomains) GetBackendKind() int64`

GetBackendKind returns the BackendKind field if non-nil, zero value otherwise.

### GetBackendKindOk

`func (o *OrganizationCustomDomains) GetBackendKindOk() (*int64, bool)`

GetBackendKindOk returns a tuple with the BackendKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendKind

`func (o *OrganizationCustomDomains) SetBackendKind(v int64)`

SetBackendKind sets BackendKind field to given value.

### HasBackendKind

`func (o *OrganizationCustomDomains) HasBackendKind() bool`

HasBackendKind returns a boolean if a field has been set.

### SetBackendKindNil

`func (o *OrganizationCustomDomains) SetBackendKindNil(b bool)`

 SetBackendKindNil sets the value for BackendKind to be an explicit nil

### UnsetBackendKind
`func (o *OrganizationCustomDomains) UnsetBackendKind()`

UnsetBackendKind ensures that no value is present for BackendKind, not even an explicit nil
### GetCreatedAt

`func (o *OrganizationCustomDomains) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationCustomDomains) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationCustomDomains) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OrganizationCustomDomains) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *OrganizationCustomDomains) GetCreatedBy() int64`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *OrganizationCustomDomains) GetCreatedByOk() (*int64, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *OrganizationCustomDomains) SetCreatedBy(v int64)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *OrganizationCustomDomains) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *OrganizationCustomDomains) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *OrganizationCustomDomains) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetDnsAliasValue

`func (o *OrganizationCustomDomains) GetDnsAliasValue() string`

GetDnsAliasValue returns the DnsAliasValue field if non-nil, zero value otherwise.

### GetDnsAliasValueOk

`func (o *OrganizationCustomDomains) GetDnsAliasValueOk() (*string, bool)`

GetDnsAliasValueOk returns a tuple with the DnsAliasValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsAliasValue

`func (o *OrganizationCustomDomains) SetDnsAliasValue(v string)`

SetDnsAliasValue sets DnsAliasValue field to given value.

### HasDnsAliasValue

`func (o *OrganizationCustomDomains) HasDnsAliasValue() bool`

HasDnsAliasValue returns a boolean if a field has been set.

### SetDnsAliasValueNil

`func (o *OrganizationCustomDomains) SetDnsAliasValueNil(b bool)`

 SetDnsAliasValueNil sets the value for DnsAliasValue to be an explicit nil

### UnsetDnsAliasValue
`func (o *OrganizationCustomDomains) UnsetDnsAliasValue()`

UnsetDnsAliasValue ensures that no value is present for DnsAliasValue, not even an explicit nil
### GetDnsCertName

`func (o *OrganizationCustomDomains) GetDnsCertName() string`

GetDnsCertName returns the DnsCertName field if non-nil, zero value otherwise.

### GetDnsCertNameOk

`func (o *OrganizationCustomDomains) GetDnsCertNameOk() (*string, bool)`

GetDnsCertNameOk returns a tuple with the DnsCertName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsCertName

`func (o *OrganizationCustomDomains) SetDnsCertName(v string)`

SetDnsCertName sets DnsCertName field to given value.

### HasDnsCertName

`func (o *OrganizationCustomDomains) HasDnsCertName() bool`

HasDnsCertName returns a boolean if a field has been set.

### SetDnsCertNameNil

`func (o *OrganizationCustomDomains) SetDnsCertNameNil(b bool)`

 SetDnsCertNameNil sets the value for DnsCertName to be an explicit nil

### UnsetDnsCertName
`func (o *OrganizationCustomDomains) UnsetDnsCertName()`

UnsetDnsCertName ensures that no value is present for DnsCertName, not even an explicit nil
### GetDnsCertValue

`func (o *OrganizationCustomDomains) GetDnsCertValue() string`

GetDnsCertValue returns the DnsCertValue field if non-nil, zero value otherwise.

### GetDnsCertValueOk

`func (o *OrganizationCustomDomains) GetDnsCertValueOk() (*string, bool)`

GetDnsCertValueOk returns a tuple with the DnsCertValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsCertValue

`func (o *OrganizationCustomDomains) SetDnsCertValue(v string)`

SetDnsCertValue sets DnsCertValue field to given value.

### HasDnsCertValue

`func (o *OrganizationCustomDomains) HasDnsCertValue() bool`

HasDnsCertValue returns a boolean if a field has been set.

### SetDnsCertValueNil

`func (o *OrganizationCustomDomains) SetDnsCertValueNil(b bool)`

 SetDnsCertValueNil sets the value for DnsCertValue to be an explicit nil

### UnsetDnsCertValue
`func (o *OrganizationCustomDomains) UnsetDnsCertValue()`

UnsetDnsCertValue ensures that no value is present for DnsCertValue, not even an explicit nil
### GetDomainType

`func (o *OrganizationCustomDomains) GetDomainType() int64`

GetDomainType returns the DomainType field if non-nil, zero value otherwise.

### GetDomainTypeOk

`func (o *OrganizationCustomDomains) GetDomainTypeOk() (*int64, bool)`

GetDomainTypeOk returns a tuple with the DomainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainType

`func (o *OrganizationCustomDomains) SetDomainType(v int64)`

SetDomainType sets DomainType field to given value.

### HasDomainType

`func (o *OrganizationCustomDomains) HasDomainType() bool`

HasDomainType returns a boolean if a field has been set.

### GetEnabled

`func (o *OrganizationCustomDomains) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OrganizationCustomDomains) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OrganizationCustomDomains) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OrganizationCustomDomains) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHost

`func (o *OrganizationCustomDomains) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *OrganizationCustomDomains) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *OrganizationCustomDomains) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *OrganizationCustomDomains) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetNamespace

`func (o *OrganizationCustomDomains) GetNamespace() int64`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *OrganizationCustomDomains) GetNamespaceOk() (*int64, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *OrganizationCustomDomains) SetNamespace(v int64)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *OrganizationCustomDomains) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetPrimary

`func (o *OrganizationCustomDomains) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *OrganizationCustomDomains) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *OrganizationCustomDomains) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *OrganizationCustomDomains) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetRedirectRoot

`func (o *OrganizationCustomDomains) GetRedirectRoot() bool`

GetRedirectRoot returns the RedirectRoot field if non-nil, zero value otherwise.

### GetRedirectRootOk

`func (o *OrganizationCustomDomains) GetRedirectRootOk() (*bool, bool)`

GetRedirectRootOk returns a tuple with the RedirectRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectRoot

`func (o *OrganizationCustomDomains) SetRedirectRoot(v bool)`

SetRedirectRoot sets RedirectRoot field to given value.

### HasRedirectRoot

`func (o *OrganizationCustomDomains) HasRedirectRoot() bool`

HasRedirectRoot returns a boolean if a field has been set.

### GetRedirectRootUrl

`func (o *OrganizationCustomDomains) GetRedirectRootUrl() string`

GetRedirectRootUrl returns the RedirectRootUrl field if non-nil, zero value otherwise.

### GetRedirectRootUrlOk

`func (o *OrganizationCustomDomains) GetRedirectRootUrlOk() (*string, bool)`

GetRedirectRootUrlOk returns a tuple with the RedirectRootUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectRootUrl

`func (o *OrganizationCustomDomains) SetRedirectRootUrl(v string)`

SetRedirectRootUrl sets RedirectRootUrl field to given value.

### HasRedirectRootUrl

`func (o *OrganizationCustomDomains) HasRedirectRootUrl() bool`

HasRedirectRootUrl returns a boolean if a field has been set.

### SetRedirectRootUrlNil

`func (o *OrganizationCustomDomains) SetRedirectRootUrlNil(b bool)`

 SetRedirectRootUrlNil sets the value for RedirectRootUrl to be an explicit nil

### UnsetRedirectRootUrl
`func (o *OrganizationCustomDomains) UnsetRedirectRootUrl()`

UnsetRedirectRootUrl ensures that no value is present for RedirectRootUrl, not even an explicit nil
### GetRepository

`func (o *OrganizationCustomDomains) GetRepository() OrganizationCustomDomainNestedRepo`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *OrganizationCustomDomains) GetRepositoryOk() (*OrganizationCustomDomainNestedRepo, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *OrganizationCustomDomains) SetRepository(v OrganizationCustomDomainNestedRepo)`

SetRepository sets Repository field to given value.


### GetRepositoryOnly

`func (o *OrganizationCustomDomains) GetRepositoryOnly() bool`

GetRepositoryOnly returns the RepositoryOnly field if non-nil, zero value otherwise.

### GetRepositoryOnlyOk

`func (o *OrganizationCustomDomains) GetRepositoryOnlyOk() (*bool, bool)`

GetRepositoryOnlyOk returns a tuple with the RepositoryOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryOnly

`func (o *OrganizationCustomDomains) SetRepositoryOnly(v bool)`

SetRepositoryOnly sets RepositoryOnly field to given value.

### HasRepositoryOnly

`func (o *OrganizationCustomDomains) HasRepositoryOnly() bool`

HasRepositoryOnly returns a boolean if a field has been set.

### GetSlugPerm

`func (o *OrganizationCustomDomains) GetSlugPerm() string`

GetSlugPerm returns the SlugPerm field if non-nil, zero value otherwise.

### GetSlugPermOk

`func (o *OrganizationCustomDomains) GetSlugPermOk() (*string, bool)`

GetSlugPermOk returns a tuple with the SlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugPerm

`func (o *OrganizationCustomDomains) SetSlugPerm(v string)`

SetSlugPerm sets SlugPerm field to given value.

### HasSlugPerm

`func (o *OrganizationCustomDomains) HasSlugPerm() bool`

HasSlugPerm returns a boolean if a field has been set.

### GetValidated

`func (o *OrganizationCustomDomains) GetValidated() bool`

GetValidated returns the Validated field if non-nil, zero value otherwise.

### GetValidatedOk

`func (o *OrganizationCustomDomains) GetValidatedOk() (*bool, bool)`

GetValidatedOk returns a tuple with the Validated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidated

`func (o *OrganizationCustomDomains) SetValidated(v bool)`

SetValidated sets Validated field to given value.

### HasValidated

`func (o *OrganizationCustomDomains) HasValidated() bool`

HasValidated returns a boolean if a field has been set.

### GetValidatedAt

`func (o *OrganizationCustomDomains) GetValidatedAt() time.Time`

GetValidatedAt returns the ValidatedAt field if non-nil, zero value otherwise.

### GetValidatedAtOk

`func (o *OrganizationCustomDomains) GetValidatedAtOk() (*time.Time, bool)`

GetValidatedAtOk returns a tuple with the ValidatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatedAt

`func (o *OrganizationCustomDomains) SetValidatedAt(v time.Time)`

SetValidatedAt sets ValidatedAt field to given value.

### HasValidatedAt

`func (o *OrganizationCustomDomains) HasValidatedAt() bool`

HasValidatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


