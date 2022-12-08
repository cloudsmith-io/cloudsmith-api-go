# FormatResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Description of the package format | 
**Distributions** | Pointer to [**[]Distribution**](Distribution.md) | The distributions supported by this package format | [optional] 
**Extensions** | **[]string** | A non-exhaustive list of extensions supported | 
**Name** | **string** | Name for the package format | 
**Premium** | **bool** | If true the package format is a premium-only feature | 
**PremiumPlanId** | Pointer to **NullableString** | The minimum plan id required for this package format | [optional] 
**PremiumPlanName** | Pointer to **NullableString** | The minimum plan name required for this package format | [optional] 
**Slug** | **string** | Slug for the package format | 
**Supports** | [**FormatSupport**](FormatSupport.md) |  | 

## Methods

### NewFormatResponse

`func NewFormatResponse(description string, extensions []string, name string, premium bool, slug string, supports FormatSupport, ) *FormatResponse`

NewFormatResponse instantiates a new FormatResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormatResponseWithDefaults

`func NewFormatResponseWithDefaults() *FormatResponse`

NewFormatResponseWithDefaults instantiates a new FormatResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *FormatResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FormatResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FormatResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDistributions

`func (o *FormatResponse) GetDistributions() []Distribution`

GetDistributions returns the Distributions field if non-nil, zero value otherwise.

### GetDistributionsOk

`func (o *FormatResponse) GetDistributionsOk() (*[]Distribution, bool)`

GetDistributionsOk returns a tuple with the Distributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributions

`func (o *FormatResponse) SetDistributions(v []Distribution)`

SetDistributions sets Distributions field to given value.

### HasDistributions

`func (o *FormatResponse) HasDistributions() bool`

HasDistributions returns a boolean if a field has been set.

### SetDistributionsNil

`func (o *FormatResponse) SetDistributionsNil(b bool)`

 SetDistributionsNil sets the value for Distributions to be an explicit nil

### UnsetDistributions
`func (o *FormatResponse) UnsetDistributions()`

UnsetDistributions ensures that no value is present for Distributions, not even an explicit nil
### GetExtensions

`func (o *FormatResponse) GetExtensions() []string`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *FormatResponse) GetExtensionsOk() (*[]string, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *FormatResponse) SetExtensions(v []string)`

SetExtensions sets Extensions field to given value.


### GetName

`func (o *FormatResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FormatResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FormatResponse) SetName(v string)`

SetName sets Name field to given value.


### GetPremium

`func (o *FormatResponse) GetPremium() bool`

GetPremium returns the Premium field if non-nil, zero value otherwise.

### GetPremiumOk

`func (o *FormatResponse) GetPremiumOk() (*bool, bool)`

GetPremiumOk returns a tuple with the Premium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremium

`func (o *FormatResponse) SetPremium(v bool)`

SetPremium sets Premium field to given value.


### GetPremiumPlanId

`func (o *FormatResponse) GetPremiumPlanId() string`

GetPremiumPlanId returns the PremiumPlanId field if non-nil, zero value otherwise.

### GetPremiumPlanIdOk

`func (o *FormatResponse) GetPremiumPlanIdOk() (*string, bool)`

GetPremiumPlanIdOk returns a tuple with the PremiumPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiumPlanId

`func (o *FormatResponse) SetPremiumPlanId(v string)`

SetPremiumPlanId sets PremiumPlanId field to given value.

### HasPremiumPlanId

`func (o *FormatResponse) HasPremiumPlanId() bool`

HasPremiumPlanId returns a boolean if a field has been set.

### SetPremiumPlanIdNil

`func (o *FormatResponse) SetPremiumPlanIdNil(b bool)`

 SetPremiumPlanIdNil sets the value for PremiumPlanId to be an explicit nil

### UnsetPremiumPlanId
`func (o *FormatResponse) UnsetPremiumPlanId()`

UnsetPremiumPlanId ensures that no value is present for PremiumPlanId, not even an explicit nil
### GetPremiumPlanName

`func (o *FormatResponse) GetPremiumPlanName() string`

GetPremiumPlanName returns the PremiumPlanName field if non-nil, zero value otherwise.

### GetPremiumPlanNameOk

`func (o *FormatResponse) GetPremiumPlanNameOk() (*string, bool)`

GetPremiumPlanNameOk returns a tuple with the PremiumPlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiumPlanName

`func (o *FormatResponse) SetPremiumPlanName(v string)`

SetPremiumPlanName sets PremiumPlanName field to given value.

### HasPremiumPlanName

`func (o *FormatResponse) HasPremiumPlanName() bool`

HasPremiumPlanName returns a boolean if a field has been set.

### SetPremiumPlanNameNil

`func (o *FormatResponse) SetPremiumPlanNameNil(b bool)`

 SetPremiumPlanNameNil sets the value for PremiumPlanName to be an explicit nil

### UnsetPremiumPlanName
`func (o *FormatResponse) UnsetPremiumPlanName()`

UnsetPremiumPlanName ensures that no value is present for PremiumPlanName, not even an explicit nil
### GetSlug

`func (o *FormatResponse) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *FormatResponse) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *FormatResponse) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetSupports

`func (o *FormatResponse) GetSupports() FormatSupport`

GetSupports returns the Supports field if non-nil, zero value otherwise.

### GetSupportsOk

`func (o *FormatResponse) GetSupportsOk() (*FormatSupport, bool)`

GetSupportsOk returns a tuple with the Supports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupports

`func (o *FormatResponse) SetSupports(v FormatSupport)`

SetSupports sets Supports field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


