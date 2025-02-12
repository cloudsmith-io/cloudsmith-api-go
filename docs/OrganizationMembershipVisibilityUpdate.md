# OrganizationMembershipVisibilityUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] [readonly] 
**HasTwoFactor** | Pointer to **bool** |  | [optional] [readonly] 
**JoinedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**LastLoginAt** | Pointer to **NullableTime** |  | [optional] [readonly] 
**LastLoginMethod** | Pointer to **string** |  | [optional] [readonly] [default to "Unknown"]
**Role** | Pointer to **string** |  | [optional] [readonly] [default to "Owner"]
**User** | Pointer to **string** |  | [optional] [readonly] 
**UserId** | Pointer to **string** |  | [optional] [readonly] 
**UserName** | Pointer to **string** |  | [optional] [readonly] 
**UserUrl** | Pointer to **string** |  | [optional] [readonly] 
**Visibility** | Pointer to **string** |  | [optional] [default to "Public"]

## Methods

### NewOrganizationMembershipVisibilityUpdate

`func NewOrganizationMembershipVisibilityUpdate() *OrganizationMembershipVisibilityUpdate`

NewOrganizationMembershipVisibilityUpdate instantiates a new OrganizationMembershipVisibilityUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationMembershipVisibilityUpdateWithDefaults

`func NewOrganizationMembershipVisibilityUpdateWithDefaults() *OrganizationMembershipVisibilityUpdate`

NewOrganizationMembershipVisibilityUpdateWithDefaults instantiates a new OrganizationMembershipVisibilityUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *OrganizationMembershipVisibilityUpdate) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrganizationMembershipVisibilityUpdate) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrganizationMembershipVisibilityUpdate) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrganizationMembershipVisibilityUpdate) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetHasTwoFactor

`func (o *OrganizationMembershipVisibilityUpdate) GetHasTwoFactor() bool`

GetHasTwoFactor returns the HasTwoFactor field if non-nil, zero value otherwise.

### GetHasTwoFactorOk

`func (o *OrganizationMembershipVisibilityUpdate) GetHasTwoFactorOk() (*bool, bool)`

GetHasTwoFactorOk returns a tuple with the HasTwoFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTwoFactor

`func (o *OrganizationMembershipVisibilityUpdate) SetHasTwoFactor(v bool)`

SetHasTwoFactor sets HasTwoFactor field to given value.

### HasHasTwoFactor

`func (o *OrganizationMembershipVisibilityUpdate) HasHasTwoFactor() bool`

HasHasTwoFactor returns a boolean if a field has been set.

### GetJoinedAt

`func (o *OrganizationMembershipVisibilityUpdate) GetJoinedAt() time.Time`

GetJoinedAt returns the JoinedAt field if non-nil, zero value otherwise.

### GetJoinedAtOk

`func (o *OrganizationMembershipVisibilityUpdate) GetJoinedAtOk() (*time.Time, bool)`

GetJoinedAtOk returns a tuple with the JoinedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinedAt

`func (o *OrganizationMembershipVisibilityUpdate) SetJoinedAt(v time.Time)`

SetJoinedAt sets JoinedAt field to given value.

### HasJoinedAt

`func (o *OrganizationMembershipVisibilityUpdate) HasJoinedAt() bool`

HasJoinedAt returns a boolean if a field has been set.

### GetLastLoginAt

`func (o *OrganizationMembershipVisibilityUpdate) GetLastLoginAt() time.Time`

GetLastLoginAt returns the LastLoginAt field if non-nil, zero value otherwise.

### GetLastLoginAtOk

`func (o *OrganizationMembershipVisibilityUpdate) GetLastLoginAtOk() (*time.Time, bool)`

GetLastLoginAtOk returns a tuple with the LastLoginAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginAt

`func (o *OrganizationMembershipVisibilityUpdate) SetLastLoginAt(v time.Time)`

SetLastLoginAt sets LastLoginAt field to given value.

### HasLastLoginAt

`func (o *OrganizationMembershipVisibilityUpdate) HasLastLoginAt() bool`

HasLastLoginAt returns a boolean if a field has been set.

### SetLastLoginAtNil

`func (o *OrganizationMembershipVisibilityUpdate) SetLastLoginAtNil(b bool)`

 SetLastLoginAtNil sets the value for LastLoginAt to be an explicit nil

### UnsetLastLoginAt
`func (o *OrganizationMembershipVisibilityUpdate) UnsetLastLoginAt()`

UnsetLastLoginAt ensures that no value is present for LastLoginAt, not even an explicit nil
### GetLastLoginMethod

`func (o *OrganizationMembershipVisibilityUpdate) GetLastLoginMethod() string`

GetLastLoginMethod returns the LastLoginMethod field if non-nil, zero value otherwise.

### GetLastLoginMethodOk

`func (o *OrganizationMembershipVisibilityUpdate) GetLastLoginMethodOk() (*string, bool)`

GetLastLoginMethodOk returns a tuple with the LastLoginMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginMethod

`func (o *OrganizationMembershipVisibilityUpdate) SetLastLoginMethod(v string)`

SetLastLoginMethod sets LastLoginMethod field to given value.

### HasLastLoginMethod

`func (o *OrganizationMembershipVisibilityUpdate) HasLastLoginMethod() bool`

HasLastLoginMethod returns a boolean if a field has been set.

### GetRole

`func (o *OrganizationMembershipVisibilityUpdate) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OrganizationMembershipVisibilityUpdate) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OrganizationMembershipVisibilityUpdate) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *OrganizationMembershipVisibilityUpdate) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetUser

`func (o *OrganizationMembershipVisibilityUpdate) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *OrganizationMembershipVisibilityUpdate) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *OrganizationMembershipVisibilityUpdate) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *OrganizationMembershipVisibilityUpdate) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserId

`func (o *OrganizationMembershipVisibilityUpdate) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *OrganizationMembershipVisibilityUpdate) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *OrganizationMembershipVisibilityUpdate) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *OrganizationMembershipVisibilityUpdate) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserName

`func (o *OrganizationMembershipVisibilityUpdate) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *OrganizationMembershipVisibilityUpdate) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *OrganizationMembershipVisibilityUpdate) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *OrganizationMembershipVisibilityUpdate) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetUserUrl

`func (o *OrganizationMembershipVisibilityUpdate) GetUserUrl() string`

GetUserUrl returns the UserUrl field if non-nil, zero value otherwise.

### GetUserUrlOk

`func (o *OrganizationMembershipVisibilityUpdate) GetUserUrlOk() (*string, bool)`

GetUserUrlOk returns a tuple with the UserUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUrl

`func (o *OrganizationMembershipVisibilityUpdate) SetUserUrl(v string)`

SetUserUrl sets UserUrl field to given value.

### HasUserUrl

`func (o *OrganizationMembershipVisibilityUpdate) HasUserUrl() bool`

HasUserUrl returns a boolean if a field has been set.

### GetVisibility

`func (o *OrganizationMembershipVisibilityUpdate) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *OrganizationMembershipVisibilityUpdate) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *OrganizationMembershipVisibilityUpdate) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *OrganizationMembershipVisibilityUpdate) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


