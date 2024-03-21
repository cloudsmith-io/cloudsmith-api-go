# OrganizationMembership

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] [readonly] 
**HasTwoFactor** | Pointer to **bool** |  | [optional] [readonly] 
**IsActive** | Pointer to **string** |  | [optional] [readonly] 
**JoinedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**LastLoginAt** | Pointer to **NullableTime** |  | [optional] [readonly] 
**LastLoginMethod** | Pointer to **string** |  | [optional] [readonly] [default to "Unknown"]
**Role** | Pointer to **string** |  | [optional] [readonly] [default to "Owner"]
**User** | Pointer to **string** |  | [optional] [readonly] 
**UserId** | Pointer to **string** |  | [optional] [readonly] 
**UserName** | Pointer to **string** |  | [optional] [readonly] 
**UserUrl** | Pointer to **string** |  | [optional] [readonly] 
**Visibility** | Pointer to **string** |  | [optional] [readonly] [default to "Public"]

## Methods

### NewOrganizationMembership

`func NewOrganizationMembership() *OrganizationMembership`

NewOrganizationMembership instantiates a new OrganizationMembership object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationMembershipWithDefaults

`func NewOrganizationMembershipWithDefaults() *OrganizationMembership`

NewOrganizationMembershipWithDefaults instantiates a new OrganizationMembership object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *OrganizationMembership) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrganizationMembership) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrganizationMembership) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrganizationMembership) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetHasTwoFactor

`func (o *OrganizationMembership) GetHasTwoFactor() bool`

GetHasTwoFactor returns the HasTwoFactor field if non-nil, zero value otherwise.

### GetHasTwoFactorOk

`func (o *OrganizationMembership) GetHasTwoFactorOk() (*bool, bool)`

GetHasTwoFactorOk returns a tuple with the HasTwoFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTwoFactor

`func (o *OrganizationMembership) SetHasTwoFactor(v bool)`

SetHasTwoFactor sets HasTwoFactor field to given value.

### HasHasTwoFactor

`func (o *OrganizationMembership) HasHasTwoFactor() bool`

HasHasTwoFactor returns a boolean if a field has been set.

### GetIsActive

`func (o *OrganizationMembership) GetIsActive() string`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *OrganizationMembership) GetIsActiveOk() (*string, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *OrganizationMembership) SetIsActive(v string)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *OrganizationMembership) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetJoinedAt

`func (o *OrganizationMembership) GetJoinedAt() time.Time`

GetJoinedAt returns the JoinedAt field if non-nil, zero value otherwise.

### GetJoinedAtOk

`func (o *OrganizationMembership) GetJoinedAtOk() (*time.Time, bool)`

GetJoinedAtOk returns a tuple with the JoinedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinedAt

`func (o *OrganizationMembership) SetJoinedAt(v time.Time)`

SetJoinedAt sets JoinedAt field to given value.

### HasJoinedAt

`func (o *OrganizationMembership) HasJoinedAt() bool`

HasJoinedAt returns a boolean if a field has been set.

### GetLastLoginAt

`func (o *OrganizationMembership) GetLastLoginAt() time.Time`

GetLastLoginAt returns the LastLoginAt field if non-nil, zero value otherwise.

### GetLastLoginAtOk

`func (o *OrganizationMembership) GetLastLoginAtOk() (*time.Time, bool)`

GetLastLoginAtOk returns a tuple with the LastLoginAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginAt

`func (o *OrganizationMembership) SetLastLoginAt(v time.Time)`

SetLastLoginAt sets LastLoginAt field to given value.

### HasLastLoginAt

`func (o *OrganizationMembership) HasLastLoginAt() bool`

HasLastLoginAt returns a boolean if a field has been set.

### SetLastLoginAtNil

`func (o *OrganizationMembership) SetLastLoginAtNil(b bool)`

 SetLastLoginAtNil sets the value for LastLoginAt to be an explicit nil

### UnsetLastLoginAt
`func (o *OrganizationMembership) UnsetLastLoginAt()`

UnsetLastLoginAt ensures that no value is present for LastLoginAt, not even an explicit nil
### GetLastLoginMethod

`func (o *OrganizationMembership) GetLastLoginMethod() string`

GetLastLoginMethod returns the LastLoginMethod field if non-nil, zero value otherwise.

### GetLastLoginMethodOk

`func (o *OrganizationMembership) GetLastLoginMethodOk() (*string, bool)`

GetLastLoginMethodOk returns a tuple with the LastLoginMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginMethod

`func (o *OrganizationMembership) SetLastLoginMethod(v string)`

SetLastLoginMethod sets LastLoginMethod field to given value.

### HasLastLoginMethod

`func (o *OrganizationMembership) HasLastLoginMethod() bool`

HasLastLoginMethod returns a boolean if a field has been set.

### GetRole

`func (o *OrganizationMembership) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OrganizationMembership) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OrganizationMembership) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *OrganizationMembership) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetUser

`func (o *OrganizationMembership) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *OrganizationMembership) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *OrganizationMembership) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *OrganizationMembership) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserId

`func (o *OrganizationMembership) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *OrganizationMembership) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *OrganizationMembership) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *OrganizationMembership) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserName

`func (o *OrganizationMembership) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *OrganizationMembership) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *OrganizationMembership) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *OrganizationMembership) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetUserUrl

`func (o *OrganizationMembership) GetUserUrl() string`

GetUserUrl returns the UserUrl field if non-nil, zero value otherwise.

### GetUserUrlOk

`func (o *OrganizationMembership) GetUserUrlOk() (*string, bool)`

GetUserUrlOk returns a tuple with the UserUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUrl

`func (o *OrganizationMembership) SetUserUrl(v string)`

SetUserUrl sets UserUrl field to given value.

### HasUserUrl

`func (o *OrganizationMembership) HasUserUrl() bool`

HasUserUrl returns a boolean if a field has been set.

### GetVisibility

`func (o *OrganizationMembership) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *OrganizationMembership) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *OrganizationMembership) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *OrganizationMembership) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


