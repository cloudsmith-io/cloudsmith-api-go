# Format

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Description of the package format | 
**Distributions** | Pointer to [**[]FormatsDistributions**](FormatsDistributions.md) | The distributions supported by this package format | [optional] 
**Extensions** | **[]string** | A non-exhaustive list of extensions supported | 
**Name** | **string** | Name for the package format | 
**Premium** | **bool** | If true the package format is a premium-only feature | 
**PremiumPlanId** | Pointer to **string** | The minimum plan id required for this package format | [optional] 
**PremiumPlanName** | Pointer to **string** | The minimum plan name required for this package format | [optional] 
**Slug** | **string** | Slug for the package format | 
**Supports** | **map[string]interface{}** | A set of what the package format supports | 

## Methods

### NewFormat

`func NewFormat(description string, extensions []string, name string, premium bool, slug string, supports map[string]interface{}, ) *Format`

NewFormat instantiates a new Format object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormatWithDefaults

`func NewFormatWithDefaults() *Format`

NewFormatWithDefaults instantiates a new Format object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Format) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Format) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Format) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDistributions

`func (o *Format) GetDistributions() []FormatsDistributions`

GetDistributions returns the Distributions field if non-nil, zero value otherwise.

### GetDistributionsOk

`func (o *Format) GetDistributionsOk() (*[]FormatsDistributions, bool)`

GetDistributionsOk returns a tuple with the Distributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributions

`func (o *Format) SetDistributions(v []FormatsDistributions)`

SetDistributions sets Distributions field to given value.

### HasDistributions

`func (o *Format) HasDistributions() bool`

HasDistributions returns a boolean if a field has been set.

### GetExtensions

`func (o *Format) GetExtensions() []string`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Format) GetExtensionsOk() (*[]string, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *Format) SetExtensions(v []string)`

SetExtensions sets Extensions field to given value.


### GetName

`func (o *Format) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Format) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Format) SetName(v string)`

SetName sets Name field to given value.


### GetPremium

`func (o *Format) GetPremium() bool`

GetPremium returns the Premium field if non-nil, zero value otherwise.

### GetPremiumOk

`func (o *Format) GetPremiumOk() (*bool, bool)`

GetPremiumOk returns a tuple with the Premium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremium

`func (o *Format) SetPremium(v bool)`

SetPremium sets Premium field to given value.


### GetPremiumPlanId

`func (o *Format) GetPremiumPlanId() string`

GetPremiumPlanId returns the PremiumPlanId field if non-nil, zero value otherwise.

### GetPremiumPlanIdOk

`func (o *Format) GetPremiumPlanIdOk() (*string, bool)`

GetPremiumPlanIdOk returns a tuple with the PremiumPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiumPlanId

`func (o *Format) SetPremiumPlanId(v string)`

SetPremiumPlanId sets PremiumPlanId field to given value.

### HasPremiumPlanId

`func (o *Format) HasPremiumPlanId() bool`

HasPremiumPlanId returns a boolean if a field has been set.

### GetPremiumPlanName

`func (o *Format) GetPremiumPlanName() string`

GetPremiumPlanName returns the PremiumPlanName field if non-nil, zero value otherwise.

### GetPremiumPlanNameOk

`func (o *Format) GetPremiumPlanNameOk() (*string, bool)`

GetPremiumPlanNameOk returns a tuple with the PremiumPlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiumPlanName

`func (o *Format) SetPremiumPlanName(v string)`

SetPremiumPlanName sets PremiumPlanName field to given value.

### HasPremiumPlanName

`func (o *Format) HasPremiumPlanName() bool`

HasPremiumPlanName returns a boolean if a field has been set.

### GetSlug

`func (o *Format) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Format) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Format) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetSupports

`func (o *Format) GetSupports() map[string]interface{}`

GetSupports returns the Supports field if non-nil, zero value otherwise.

### GetSupportsOk

`func (o *Format) GetSupportsOk() (*map[string]interface{}, bool)`

GetSupportsOk returns a tuple with the Supports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupports

`func (o *Format) SetSupports(v map[string]interface{})`

SetSupports sets Supports field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


