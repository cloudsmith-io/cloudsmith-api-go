# OrganizationSAMLAuthRequestPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SamlAuthEnabled** | Pointer to **bool** |  | [optional] 
**SamlAuthEnforced** | Pointer to **bool** |  | [optional] 
**SamlMetadataInline** | Pointer to **string** | If configured, SAML metadata will be used as entered instead of retrieved from a remote URL. | [optional] 
**SamlMetadataInlineWebapp** | Pointer to **NullableString** | When configured, this inline SAML metadata is used instead of the legacy app SAML configuration when signing into the new Cloudsmith web application. | [optional] 
**SamlMetadataUrl** | Pointer to **NullableString** | If configured, SAML metadata be retrieved from a remote URL. | [optional] 
**SamlMetadataUrlWebapp** | Pointer to **NullableString** | When configured, this SAML metadata URL is used instead of the legacy app SAML configuration when signing into the new Cloudsmith web application. | [optional] 

## Methods

### NewOrganizationSAMLAuthRequestPatch

`func NewOrganizationSAMLAuthRequestPatch() *OrganizationSAMLAuthRequestPatch`

NewOrganizationSAMLAuthRequestPatch instantiates a new OrganizationSAMLAuthRequestPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationSAMLAuthRequestPatchWithDefaults

`func NewOrganizationSAMLAuthRequestPatchWithDefaults() *OrganizationSAMLAuthRequestPatch`

NewOrganizationSAMLAuthRequestPatchWithDefaults instantiates a new OrganizationSAMLAuthRequestPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSamlAuthEnabled

`func (o *OrganizationSAMLAuthRequestPatch) GetSamlAuthEnabled() bool`

GetSamlAuthEnabled returns the SamlAuthEnabled field if non-nil, zero value otherwise.

### GetSamlAuthEnabledOk

`func (o *OrganizationSAMLAuthRequestPatch) GetSamlAuthEnabledOk() (*bool, bool)`

GetSamlAuthEnabledOk returns a tuple with the SamlAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAuthEnabled

`func (o *OrganizationSAMLAuthRequestPatch) SetSamlAuthEnabled(v bool)`

SetSamlAuthEnabled sets SamlAuthEnabled field to given value.

### HasSamlAuthEnabled

`func (o *OrganizationSAMLAuthRequestPatch) HasSamlAuthEnabled() bool`

HasSamlAuthEnabled returns a boolean if a field has been set.

### GetSamlAuthEnforced

`func (o *OrganizationSAMLAuthRequestPatch) GetSamlAuthEnforced() bool`

GetSamlAuthEnforced returns the SamlAuthEnforced field if non-nil, zero value otherwise.

### GetSamlAuthEnforcedOk

`func (o *OrganizationSAMLAuthRequestPatch) GetSamlAuthEnforcedOk() (*bool, bool)`

GetSamlAuthEnforcedOk returns a tuple with the SamlAuthEnforced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAuthEnforced

`func (o *OrganizationSAMLAuthRequestPatch) SetSamlAuthEnforced(v bool)`

SetSamlAuthEnforced sets SamlAuthEnforced field to given value.

### HasSamlAuthEnforced

`func (o *OrganizationSAMLAuthRequestPatch) HasSamlAuthEnforced() bool`

HasSamlAuthEnforced returns a boolean if a field has been set.

### GetSamlMetadataInline

`func (o *OrganizationSAMLAuthRequestPatch) GetSamlMetadataInline() string`

GetSamlMetadataInline returns the SamlMetadataInline field if non-nil, zero value otherwise.

### GetSamlMetadataInlineOk

`func (o *OrganizationSAMLAuthRequestPatch) GetSamlMetadataInlineOk() (*string, bool)`

GetSamlMetadataInlineOk returns a tuple with the SamlMetadataInline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlMetadataInline

`func (o *OrganizationSAMLAuthRequestPatch) SetSamlMetadataInline(v string)`

SetSamlMetadataInline sets SamlMetadataInline field to given value.

### HasSamlMetadataInline

`func (o *OrganizationSAMLAuthRequestPatch) HasSamlMetadataInline() bool`

HasSamlMetadataInline returns a boolean if a field has been set.

### GetSamlMetadataInlineWebapp

`func (o *OrganizationSAMLAuthRequestPatch) GetSamlMetadataInlineWebapp() string`

GetSamlMetadataInlineWebapp returns the SamlMetadataInlineWebapp field if non-nil, zero value otherwise.

### GetSamlMetadataInlineWebappOk

`func (o *OrganizationSAMLAuthRequestPatch) GetSamlMetadataInlineWebappOk() (*string, bool)`

GetSamlMetadataInlineWebappOk returns a tuple with the SamlMetadataInlineWebapp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlMetadataInlineWebapp

`func (o *OrganizationSAMLAuthRequestPatch) SetSamlMetadataInlineWebapp(v string)`

SetSamlMetadataInlineWebapp sets SamlMetadataInlineWebapp field to given value.

### HasSamlMetadataInlineWebapp

`func (o *OrganizationSAMLAuthRequestPatch) HasSamlMetadataInlineWebapp() bool`

HasSamlMetadataInlineWebapp returns a boolean if a field has been set.

### SetSamlMetadataInlineWebappNil

`func (o *OrganizationSAMLAuthRequestPatch) SetSamlMetadataInlineWebappNil(b bool)`

 SetSamlMetadataInlineWebappNil sets the value for SamlMetadataInlineWebapp to be an explicit nil

### UnsetSamlMetadataInlineWebapp
`func (o *OrganizationSAMLAuthRequestPatch) UnsetSamlMetadataInlineWebapp()`

UnsetSamlMetadataInlineWebapp ensures that no value is present for SamlMetadataInlineWebapp, not even an explicit nil
### GetSamlMetadataUrl

`func (o *OrganizationSAMLAuthRequestPatch) GetSamlMetadataUrl() string`

GetSamlMetadataUrl returns the SamlMetadataUrl field if non-nil, zero value otherwise.

### GetSamlMetadataUrlOk

`func (o *OrganizationSAMLAuthRequestPatch) GetSamlMetadataUrlOk() (*string, bool)`

GetSamlMetadataUrlOk returns a tuple with the SamlMetadataUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlMetadataUrl

`func (o *OrganizationSAMLAuthRequestPatch) SetSamlMetadataUrl(v string)`

SetSamlMetadataUrl sets SamlMetadataUrl field to given value.

### HasSamlMetadataUrl

`func (o *OrganizationSAMLAuthRequestPatch) HasSamlMetadataUrl() bool`

HasSamlMetadataUrl returns a boolean if a field has been set.

### SetSamlMetadataUrlNil

`func (o *OrganizationSAMLAuthRequestPatch) SetSamlMetadataUrlNil(b bool)`

 SetSamlMetadataUrlNil sets the value for SamlMetadataUrl to be an explicit nil

### UnsetSamlMetadataUrl
`func (o *OrganizationSAMLAuthRequestPatch) UnsetSamlMetadataUrl()`

UnsetSamlMetadataUrl ensures that no value is present for SamlMetadataUrl, not even an explicit nil
### GetSamlMetadataUrlWebapp

`func (o *OrganizationSAMLAuthRequestPatch) GetSamlMetadataUrlWebapp() string`

GetSamlMetadataUrlWebapp returns the SamlMetadataUrlWebapp field if non-nil, zero value otherwise.

### GetSamlMetadataUrlWebappOk

`func (o *OrganizationSAMLAuthRequestPatch) GetSamlMetadataUrlWebappOk() (*string, bool)`

GetSamlMetadataUrlWebappOk returns a tuple with the SamlMetadataUrlWebapp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlMetadataUrlWebapp

`func (o *OrganizationSAMLAuthRequestPatch) SetSamlMetadataUrlWebapp(v string)`

SetSamlMetadataUrlWebapp sets SamlMetadataUrlWebapp field to given value.

### HasSamlMetadataUrlWebapp

`func (o *OrganizationSAMLAuthRequestPatch) HasSamlMetadataUrlWebapp() bool`

HasSamlMetadataUrlWebapp returns a boolean if a field has been set.

### SetSamlMetadataUrlWebappNil

`func (o *OrganizationSAMLAuthRequestPatch) SetSamlMetadataUrlWebappNil(b bool)`

 SetSamlMetadataUrlWebappNil sets the value for SamlMetadataUrlWebapp to be an explicit nil

### UnsetSamlMetadataUrlWebapp
`func (o *OrganizationSAMLAuthRequestPatch) UnsetSamlMetadataUrlWebapp()`

UnsetSamlMetadataUrlWebapp ensures that no value is present for SamlMetadataUrlWebapp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


