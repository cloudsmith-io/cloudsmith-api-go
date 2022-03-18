# PackagesTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | None | [optional] 
**IsImmutable** | Pointer to **bool** | If true, created tags will be immutable. An immutable flag is a tag that cannot be removed from a package. | [optional] 
**Tags** | Pointer to **[]string** | A list of tags to apply the action to. Not required for clears. | [optional] 

## Methods

### NewPackagesTag

`func NewPackagesTag() *PackagesTag`

NewPackagesTag instantiates a new PackagesTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackagesTagWithDefaults

`func NewPackagesTagWithDefaults() *PackagesTag`

NewPackagesTagWithDefaults instantiates a new PackagesTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *PackagesTag) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PackagesTag) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PackagesTag) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *PackagesTag) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetIsImmutable

`func (o *PackagesTag) GetIsImmutable() bool`

GetIsImmutable returns the IsImmutable field if non-nil, zero value otherwise.

### GetIsImmutableOk

`func (o *PackagesTag) GetIsImmutableOk() (*bool, bool)`

GetIsImmutableOk returns a tuple with the IsImmutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsImmutable

`func (o *PackagesTag) SetIsImmutable(v bool)`

SetIsImmutable sets IsImmutable field to given value.

### HasIsImmutable

`func (o *PackagesTag) HasIsImmutable() bool`

HasIsImmutable returns a boolean if a field has been set.

### GetTags

`func (o *PackagesTag) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PackagesTag) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PackagesTag) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PackagesTag) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


