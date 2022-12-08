# DistributionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to **string** |  | [optional] [readonly] 
**FormatUrl** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**SelfUrl** | Pointer to **string** |  | [optional] [readonly] 
**Slug** | Pointer to **string** | The slug identifier for this distribution | [optional] [readonly] 
**Variants** | Pointer to **NullableString** |  | [optional] 
**Versions** | Pointer to [**[]DistributionVersion**](DistributionVersion.md) | A list of the versions for this distribution | [optional] [readonly] 

## Methods

### NewDistributionResponse

`func NewDistributionResponse(name string, ) *DistributionResponse`

NewDistributionResponse instantiates a new DistributionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDistributionResponseWithDefaults

`func NewDistributionResponseWithDefaults() *DistributionResponse`

NewDistributionResponseWithDefaults instantiates a new DistributionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *DistributionResponse) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *DistributionResponse) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *DistributionResponse) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *DistributionResponse) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetFormatUrl

`func (o *DistributionResponse) GetFormatUrl() string`

GetFormatUrl returns the FormatUrl field if non-nil, zero value otherwise.

### GetFormatUrlOk

`func (o *DistributionResponse) GetFormatUrlOk() (*string, bool)`

GetFormatUrlOk returns a tuple with the FormatUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatUrl

`func (o *DistributionResponse) SetFormatUrl(v string)`

SetFormatUrl sets FormatUrl field to given value.

### HasFormatUrl

`func (o *DistributionResponse) HasFormatUrl() bool`

HasFormatUrl returns a boolean if a field has been set.

### GetName

`func (o *DistributionResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DistributionResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DistributionResponse) SetName(v string)`

SetName sets Name field to given value.


### GetSelfUrl

`func (o *DistributionResponse) GetSelfUrl() string`

GetSelfUrl returns the SelfUrl field if non-nil, zero value otherwise.

### GetSelfUrlOk

`func (o *DistributionResponse) GetSelfUrlOk() (*string, bool)`

GetSelfUrlOk returns a tuple with the SelfUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfUrl

`func (o *DistributionResponse) SetSelfUrl(v string)`

SetSelfUrl sets SelfUrl field to given value.

### HasSelfUrl

`func (o *DistributionResponse) HasSelfUrl() bool`

HasSelfUrl returns a boolean if a field has been set.

### GetSlug

`func (o *DistributionResponse) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *DistributionResponse) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *DistributionResponse) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *DistributionResponse) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetVariants

`func (o *DistributionResponse) GetVariants() string`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *DistributionResponse) GetVariantsOk() (*string, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *DistributionResponse) SetVariants(v string)`

SetVariants sets Variants field to given value.

### HasVariants

`func (o *DistributionResponse) HasVariants() bool`

HasVariants returns a boolean if a field has been set.

### SetVariantsNil

`func (o *DistributionResponse) SetVariantsNil(b bool)`

 SetVariantsNil sets the value for Variants to be an explicit nil

### UnsetVariants
`func (o *DistributionResponse) UnsetVariants()`

UnsetVariants ensures that no value is present for Variants, not even an explicit nil
### GetVersions

`func (o *DistributionResponse) GetVersions() []DistributionVersion`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *DistributionResponse) GetVersionsOk() (*[]DistributionVersion, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *DistributionResponse) SetVersions(v []DistributionVersion)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *DistributionResponse) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


