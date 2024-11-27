# OrganizationSAMLAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SamlAuthEnabled** | **bool** |  | 
**SamlAuthEnforced** | **bool** |  | 
**SamlMetadataInline** | Pointer to **string** | If configured, SAML metadata will be used as entered instead of retrieved from a remote URL. | [optional] 
**SamlMetadataUrl** | Pointer to **NullableString** | If configured, SAML metadata be retrieved from a remote URL. | [optional] 

## Methods

### NewOrganizationSAMLAuth

`func NewOrganizationSAMLAuth(samlAuthEnabled bool, samlAuthEnforced bool, ) *OrganizationSAMLAuth`

NewOrganizationSAMLAuth instantiates a new OrganizationSAMLAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationSAMLAuthWithDefaults

`func NewOrganizationSAMLAuthWithDefaults() *OrganizationSAMLAuth`

NewOrganizationSAMLAuthWithDefaults instantiates a new OrganizationSAMLAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSamlAuthEnabled

`func (o *OrganizationSAMLAuth) GetSamlAuthEnabled() bool`

GetSamlAuthEnabled returns the SamlAuthEnabled field if non-nil, zero value otherwise.

### GetSamlAuthEnabledOk

`func (o *OrganizationSAMLAuth) GetSamlAuthEnabledOk() (*bool, bool)`

GetSamlAuthEnabledOk returns a tuple with the SamlAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAuthEnabled

`func (o *OrganizationSAMLAuth) SetSamlAuthEnabled(v bool)`

SetSamlAuthEnabled sets SamlAuthEnabled field to given value.


### GetSamlAuthEnforced

`func (o *OrganizationSAMLAuth) GetSamlAuthEnforced() bool`

GetSamlAuthEnforced returns the SamlAuthEnforced field if non-nil, zero value otherwise.

### GetSamlAuthEnforcedOk

`func (o *OrganizationSAMLAuth) GetSamlAuthEnforcedOk() (*bool, bool)`

GetSamlAuthEnforcedOk returns a tuple with the SamlAuthEnforced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAuthEnforced

`func (o *OrganizationSAMLAuth) SetSamlAuthEnforced(v bool)`

SetSamlAuthEnforced sets SamlAuthEnforced field to given value.


### GetSamlMetadataInline

`func (o *OrganizationSAMLAuth) GetSamlMetadataInline() string`

GetSamlMetadataInline returns the SamlMetadataInline field if non-nil, zero value otherwise.

### GetSamlMetadataInlineOk

`func (o *OrganizationSAMLAuth) GetSamlMetadataInlineOk() (*string, bool)`

GetSamlMetadataInlineOk returns a tuple with the SamlMetadataInline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlMetadataInline

`func (o *OrganizationSAMLAuth) SetSamlMetadataInline(v string)`

SetSamlMetadataInline sets SamlMetadataInline field to given value.

### HasSamlMetadataInline

`func (o *OrganizationSAMLAuth) HasSamlMetadataInline() bool`

HasSamlMetadataInline returns a boolean if a field has been set.

### GetSamlMetadataUrl

`func (o *OrganizationSAMLAuth) GetSamlMetadataUrl() string`

GetSamlMetadataUrl returns the SamlMetadataUrl field if non-nil, zero value otherwise.

### GetSamlMetadataUrlOk

`func (o *OrganizationSAMLAuth) GetSamlMetadataUrlOk() (*string, bool)`

GetSamlMetadataUrlOk returns a tuple with the SamlMetadataUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlMetadataUrl

`func (o *OrganizationSAMLAuth) SetSamlMetadataUrl(v string)`

SetSamlMetadataUrl sets SamlMetadataUrl field to given value.

### HasSamlMetadataUrl

`func (o *OrganizationSAMLAuth) HasSamlMetadataUrl() bool`

HasSamlMetadataUrl returns a boolean if a field has been set.

### SetSamlMetadataUrlNil

`func (o *OrganizationSAMLAuth) SetSamlMetadataUrlNil(b bool)`

 SetSamlMetadataUrlNil sets the value for SamlMetadataUrl to be an explicit nil

### UnsetSamlMetadataUrl
`func (o *OrganizationSAMLAuth) UnsetSamlMetadataUrl()`

UnsetSamlMetadataUrl ensures that no value is present for SamlMetadataUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


