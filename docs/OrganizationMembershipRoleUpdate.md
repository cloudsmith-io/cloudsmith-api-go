# OrganizationMembershipRoleUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] [readonly] 
**HasTwoFactor** | Pointer to **bool** |  | [optional] [readonly] 
**JoinedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**LastLoginAt** | Pointer to **NullableTime** |  | [optional] [readonly] 
**LastLoginMethod** | Pointer to **string** |  | [optional] [readonly] [default to "Unknown"]
**Role** | Pointer to **string** |  | [optional] [default to "Owner"]
**User** | Pointer to **string** |  | [optional] [readonly] 
**UserId** | Pointer to **string** |  | [optional] [readonly] 
**UserName** | Pointer to **string** |  | [optional] [readonly] 
**UserUrl** | Pointer to **string** |  | [optional] [readonly] 
**Visibility** | Pointer to **string** |  | [optional] [readonly] [default to "Public"]

## Methods

### NewOrganizationMembershipRoleUpdate

`func NewOrganizationMembershipRoleUpdate() *OrganizationMembershipRoleUpdate`

NewOrganizationMembershipRoleUpdate instantiates a new OrganizationMembershipRoleUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationMembershipRoleUpdateWithDefaults

`func NewOrganizationMembershipRoleUpdateWithDefaults() *OrganizationMembershipRoleUpdate`

NewOrganizationMembershipRoleUpdateWithDefaults instantiates a new OrganizationMembershipRoleUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *OrganizationMembershipRoleUpdate) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrganizationMembershipRoleUpdate) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrganizationMembershipRoleUpdate) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrganizationMembershipRoleUpdate) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetHasTwoFactor

`func (o *OrganizationMembershipRoleUpdate) GetHasTwoFactor() bool`

GetHasTwoFactor returns the HasTwoFactor field if non-nil, zero value otherwise.

### GetHasTwoFactorOk

`func (o *OrganizationMembershipRoleUpdate) GetHasTwoFactorOk() (*bool, bool)`

GetHasTwoFactorOk returns a tuple with the HasTwoFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTwoFactor

`func (o *OrganizationMembershipRoleUpdate) SetHasTwoFactor(v bool)`

SetHasTwoFactor sets HasTwoFactor field to given value.

### HasHasTwoFactor

`func (o *OrganizationMembershipRoleUpdate) HasHasTwoFactor() bool`

HasHasTwoFactor returns a boolean if a field has been set.

### GetJoinedAt

`func (o *OrganizationMembershipRoleUpdate) GetJoinedAt() time.Time`

GetJoinedAt returns the JoinedAt field if non-nil, zero value otherwise.

### GetJoinedAtOk

`func (o *OrganizationMembershipRoleUpdate) GetJoinedAtOk() (*time.Time, bool)`

GetJoinedAtOk returns a tuple with the JoinedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinedAt

`func (o *OrganizationMembershipRoleUpdate) SetJoinedAt(v time.Time)`

SetJoinedAt sets JoinedAt field to given value.

### HasJoinedAt

`func (o *OrganizationMembershipRoleUpdate) HasJoinedAt() bool`

HasJoinedAt returns a boolean if a field has been set.

### GetLastLoginAt

`func (o *OrganizationMembershipRoleUpdate) GetLastLoginAt() time.Time`

GetLastLoginAt returns the LastLoginAt field if non-nil, zero value otherwise.

### GetLastLoginAtOk

`func (o *OrganizationMembershipRoleUpdate) GetLastLoginAtOk() (*time.Time, bool)`

GetLastLoginAtOk returns a tuple with the LastLoginAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginAt

`func (o *OrganizationMembershipRoleUpdate) SetLastLoginAt(v time.Time)`

SetLastLoginAt sets LastLoginAt field to given value.

### HasLastLoginAt

`func (o *OrganizationMembershipRoleUpdate) HasLastLoginAt() bool`

HasLastLoginAt returns a boolean if a field has been set.

### SetLastLoginAtNil

`func (o *OrganizationMembershipRoleUpdate) SetLastLoginAtNil(b bool)`

 SetLastLoginAtNil sets the value for LastLoginAt to be an explicit nil

### UnsetLastLoginAt
`func (o *OrganizationMembershipRoleUpdate) UnsetLastLoginAt()`

UnsetLastLoginAt ensures that no value is present for LastLoginAt, not even an explicit nil
### GetLastLoginMethod

`func (o *OrganizationMembershipRoleUpdate) GetLastLoginMethod() string`

GetLastLoginMethod returns the LastLoginMethod field if non-nil, zero value otherwise.

### GetLastLoginMethodOk

`func (o *OrganizationMembershipRoleUpdate) GetLastLoginMethodOk() (*string, bool)`

GetLastLoginMethodOk returns a tuple with the LastLoginMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginMethod

`func (o *OrganizationMembershipRoleUpdate) SetLastLoginMethod(v string)`

SetLastLoginMethod sets LastLoginMethod field to given value.

### HasLastLoginMethod

`func (o *OrganizationMembershipRoleUpdate) HasLastLoginMethod() bool`

HasLastLoginMethod returns a boolean if a field has been set.

### GetRole

`func (o *OrganizationMembershipRoleUpdate) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OrganizationMembershipRoleUpdate) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OrganizationMembershipRoleUpdate) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *OrganizationMembershipRoleUpdate) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetUser

`func (o *OrganizationMembershipRoleUpdate) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *OrganizationMembershipRoleUpdate) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *OrganizationMembershipRoleUpdate) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *OrganizationMembershipRoleUpdate) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserId

`func (o *OrganizationMembershipRoleUpdate) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *OrganizationMembershipRoleUpdate) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *OrganizationMembershipRoleUpdate) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *OrganizationMembershipRoleUpdate) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserName

`func (o *OrganizationMembershipRoleUpdate) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *OrganizationMembershipRoleUpdate) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *OrganizationMembershipRoleUpdate) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *OrganizationMembershipRoleUpdate) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetUserUrl

`func (o *OrganizationMembershipRoleUpdate) GetUserUrl() string`

GetUserUrl returns the UserUrl field if non-nil, zero value otherwise.

### GetUserUrlOk

`func (o *OrganizationMembershipRoleUpdate) GetUserUrlOk() (*string, bool)`

GetUserUrlOk returns a tuple with the UserUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUrl

`func (o *OrganizationMembershipRoleUpdate) SetUserUrl(v string)`

SetUserUrl sets UserUrl field to given value.

### HasUserUrl

`func (o *OrganizationMembershipRoleUpdate) HasUserUrl() bool`

HasUserUrl returns a boolean if a field has been set.

### GetVisibility

`func (o *OrganizationMembershipRoleUpdate) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *OrganizationMembershipRoleUpdate) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *OrganizationMembershipRoleUpdate) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *OrganizationMembershipRoleUpdate) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


