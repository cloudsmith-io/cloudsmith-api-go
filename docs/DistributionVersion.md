# DistributionVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The textual name for this version. | [optional] 
**Slug** | Pointer to **string** | The slug identifier for this version | [optional] [readonly] 

## Methods

### NewDistributionVersion

`func NewDistributionVersion() *DistributionVersion`

NewDistributionVersion instantiates a new DistributionVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDistributionVersionWithDefaults

`func NewDistributionVersionWithDefaults() *DistributionVersion`

NewDistributionVersionWithDefaults instantiates a new DistributionVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DistributionVersion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DistributionVersion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DistributionVersion) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DistributionVersion) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *DistributionVersion) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *DistributionVersion) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *DistributionVersion) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *DistributionVersion) HasSlug() bool`

HasSlug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


