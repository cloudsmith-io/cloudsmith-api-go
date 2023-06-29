# PackageLicensePolicyEvaluationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**EvaluationCount** | Pointer to **int64** |  | [optional] [readonly] 
**Policy** | Pointer to [**NestedLicensePolicy**](NestedLicensePolicy.md) |  | [optional] 
**SlugPerm** | Pointer to **string** |  | [optional] [readonly] 
**Status** | Pointer to **string** |  | [optional] [readonly] [default to "Pending"]
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**ViolationCount** | Pointer to **int64** |  | [optional] [readonly] 

## Methods

### NewPackageLicensePolicyEvaluationRequest

`func NewPackageLicensePolicyEvaluationRequest() *PackageLicensePolicyEvaluationRequest`

NewPackageLicensePolicyEvaluationRequest instantiates a new PackageLicensePolicyEvaluationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageLicensePolicyEvaluationRequestWithDefaults

`func NewPackageLicensePolicyEvaluationRequestWithDefaults() *PackageLicensePolicyEvaluationRequest`

NewPackageLicensePolicyEvaluationRequestWithDefaults instantiates a new PackageLicensePolicyEvaluationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PackageLicensePolicyEvaluationRequest) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PackageLicensePolicyEvaluationRequest) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PackageLicensePolicyEvaluationRequest) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PackageLicensePolicyEvaluationRequest) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEvaluationCount

`func (o *PackageLicensePolicyEvaluationRequest) GetEvaluationCount() int64`

GetEvaluationCount returns the EvaluationCount field if non-nil, zero value otherwise.

### GetEvaluationCountOk

`func (o *PackageLicensePolicyEvaluationRequest) GetEvaluationCountOk() (*int64, bool)`

GetEvaluationCountOk returns a tuple with the EvaluationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationCount

`func (o *PackageLicensePolicyEvaluationRequest) SetEvaluationCount(v int64)`

SetEvaluationCount sets EvaluationCount field to given value.

### HasEvaluationCount

`func (o *PackageLicensePolicyEvaluationRequest) HasEvaluationCount() bool`

HasEvaluationCount returns a boolean if a field has been set.

### GetPolicy

`func (o *PackageLicensePolicyEvaluationRequest) GetPolicy() NestedLicensePolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *PackageLicensePolicyEvaluationRequest) GetPolicyOk() (*NestedLicensePolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *PackageLicensePolicyEvaluationRequest) SetPolicy(v NestedLicensePolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *PackageLicensePolicyEvaluationRequest) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetSlugPerm

`func (o *PackageLicensePolicyEvaluationRequest) GetSlugPerm() string`

GetSlugPerm returns the SlugPerm field if non-nil, zero value otherwise.

### GetSlugPermOk

`func (o *PackageLicensePolicyEvaluationRequest) GetSlugPermOk() (*string, bool)`

GetSlugPermOk returns a tuple with the SlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugPerm

`func (o *PackageLicensePolicyEvaluationRequest) SetSlugPerm(v string)`

SetSlugPerm sets SlugPerm field to given value.

### HasSlugPerm

`func (o *PackageLicensePolicyEvaluationRequest) HasSlugPerm() bool`

HasSlugPerm returns a boolean if a field has been set.

### GetStatus

`func (o *PackageLicensePolicyEvaluationRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PackageLicensePolicyEvaluationRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PackageLicensePolicyEvaluationRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PackageLicensePolicyEvaluationRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PackageLicensePolicyEvaluationRequest) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PackageLicensePolicyEvaluationRequest) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PackageLicensePolicyEvaluationRequest) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PackageLicensePolicyEvaluationRequest) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetViolationCount

`func (o *PackageLicensePolicyEvaluationRequest) GetViolationCount() int64`

GetViolationCount returns the ViolationCount field if non-nil, zero value otherwise.

### GetViolationCountOk

`func (o *PackageLicensePolicyEvaluationRequest) GetViolationCountOk() (*int64, bool)`

GetViolationCountOk returns a tuple with the ViolationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolationCount

`func (o *PackageLicensePolicyEvaluationRequest) SetViolationCount(v int64)`

SetViolationCount sets ViolationCount field to given value.

### HasViolationCount

`func (o *PackageLicensePolicyEvaluationRequest) HasViolationCount() bool`

HasViolationCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


