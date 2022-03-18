# Distribution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to **string** |  | [optional] 
**FormatUrl** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**SelfUrl** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** | The slug identifier for this distribution | [optional] 
**Variants** | Pointer to **string** |  | [optional] 
**Versions** | Pointer to [**[]DistrosVersions**](DistrosVersions.md) | A list of the versions for this distribution | [optional] 

## Methods

### NewDistribution

`func NewDistribution(name string, ) *Distribution`

NewDistribution instantiates a new Distribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDistributionWithDefaults

`func NewDistributionWithDefaults() *Distribution`

NewDistributionWithDefaults instantiates a new Distribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *Distribution) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *Distribution) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *Distribution) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *Distribution) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetFormatUrl

`func (o *Distribution) GetFormatUrl() string`

GetFormatUrl returns the FormatUrl field if non-nil, zero value otherwise.

### GetFormatUrlOk

`func (o *Distribution) GetFormatUrlOk() (*string, bool)`

GetFormatUrlOk returns a tuple with the FormatUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatUrl

`func (o *Distribution) SetFormatUrl(v string)`

SetFormatUrl sets FormatUrl field to given value.

### HasFormatUrl

`func (o *Distribution) HasFormatUrl() bool`

HasFormatUrl returns a boolean if a field has been set.

### GetName

`func (o *Distribution) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Distribution) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Distribution) SetName(v string)`

SetName sets Name field to given value.


### GetSelfUrl

`func (o *Distribution) GetSelfUrl() string`

GetSelfUrl returns the SelfUrl field if non-nil, zero value otherwise.

### GetSelfUrlOk

`func (o *Distribution) GetSelfUrlOk() (*string, bool)`

GetSelfUrlOk returns a tuple with the SelfUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfUrl

`func (o *Distribution) SetSelfUrl(v string)`

SetSelfUrl sets SelfUrl field to given value.

### HasSelfUrl

`func (o *Distribution) HasSelfUrl() bool`

HasSelfUrl returns a boolean if a field has been set.

### GetSlug

`func (o *Distribution) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Distribution) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Distribution) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *Distribution) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetVariants

`func (o *Distribution) GetVariants() string`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *Distribution) GetVariantsOk() (*string, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *Distribution) SetVariants(v string)`

SetVariants sets Variants field to given value.

### HasVariants

`func (o *Distribution) HasVariants() bool`

HasVariants returns a boolean if a field has been set.

### GetVersions

`func (o *Distribution) GetVersions() []DistrosVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *Distribution) GetVersionsOk() (*[]DistrosVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *Distribution) SetVersions(v []DistrosVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *Distribution) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


