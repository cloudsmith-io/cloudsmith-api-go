# OrganizationMembershipResponse

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
**Visibility** | Pointer to **string** |  | [optional] [readonly] [default to "Public"]

## Methods

### NewOrganizationMembershipResponse

`func NewOrganizationMembershipResponse() *OrganizationMembershipResponse`

NewOrganizationMembershipResponse instantiates a new OrganizationMembershipResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationMembershipResponseWithDefaults

`func NewOrganizationMembershipResponseWithDefaults() *OrganizationMembershipResponse`

NewOrganizationMembershipResponseWithDefaults instantiates a new OrganizationMembershipResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *OrganizationMembershipResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrganizationMembershipResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrganizationMembershipResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrganizationMembershipResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetHasTwoFactor

`func (o *OrganizationMembershipResponse) GetHasTwoFactor() bool`

GetHasTwoFactor returns the HasTwoFactor field if non-nil, zero value otherwise.

### GetHasTwoFactorOk

`func (o *OrganizationMembershipResponse) GetHasTwoFactorOk() (*bool, bool)`

GetHasTwoFactorOk returns a tuple with the HasTwoFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTwoFactor

`func (o *OrganizationMembershipResponse) SetHasTwoFactor(v bool)`

SetHasTwoFactor sets HasTwoFactor field to given value.

### HasHasTwoFactor

`func (o *OrganizationMembershipResponse) HasHasTwoFactor() bool`

HasHasTwoFactor returns a boolean if a field has been set.

### GetJoinedAt

`func (o *OrganizationMembershipResponse) GetJoinedAt() time.Time`

GetJoinedAt returns the JoinedAt field if non-nil, zero value otherwise.

### GetJoinedAtOk

`func (o *OrganizationMembershipResponse) GetJoinedAtOk() (*time.Time, bool)`

GetJoinedAtOk returns a tuple with the JoinedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinedAt

`func (o *OrganizationMembershipResponse) SetJoinedAt(v time.Time)`

SetJoinedAt sets JoinedAt field to given value.

### HasJoinedAt

`func (o *OrganizationMembershipResponse) HasJoinedAt() bool`

HasJoinedAt returns a boolean if a field has been set.

### GetLastLoginAt

`func (o *OrganizationMembershipResponse) GetLastLoginAt() time.Time`

GetLastLoginAt returns the LastLoginAt field if non-nil, zero value otherwise.

### GetLastLoginAtOk

`func (o *OrganizationMembershipResponse) GetLastLoginAtOk() (*time.Time, bool)`

GetLastLoginAtOk returns a tuple with the LastLoginAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginAt

`func (o *OrganizationMembershipResponse) SetLastLoginAt(v time.Time)`

SetLastLoginAt sets LastLoginAt field to given value.

### HasLastLoginAt

`func (o *OrganizationMembershipResponse) HasLastLoginAt() bool`

HasLastLoginAt returns a boolean if a field has been set.

### SetLastLoginAtNil

`func (o *OrganizationMembershipResponse) SetLastLoginAtNil(b bool)`

 SetLastLoginAtNil sets the value for LastLoginAt to be an explicit nil

### UnsetLastLoginAt
`func (o *OrganizationMembershipResponse) UnsetLastLoginAt()`

UnsetLastLoginAt ensures that no value is present for LastLoginAt, not even an explicit nil
### GetLastLoginMethod

`func (o *OrganizationMembershipResponse) GetLastLoginMethod() string`

GetLastLoginMethod returns the LastLoginMethod field if non-nil, zero value otherwise.

### GetLastLoginMethodOk

`func (o *OrganizationMembershipResponse) GetLastLoginMethodOk() (*string, bool)`

GetLastLoginMethodOk returns a tuple with the LastLoginMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginMethod

`func (o *OrganizationMembershipResponse) SetLastLoginMethod(v string)`

SetLastLoginMethod sets LastLoginMethod field to given value.

### HasLastLoginMethod

`func (o *OrganizationMembershipResponse) HasLastLoginMethod() bool`

HasLastLoginMethod returns a boolean if a field has been set.

### GetRole

`func (o *OrganizationMembershipResponse) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OrganizationMembershipResponse) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OrganizationMembershipResponse) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *OrganizationMembershipResponse) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetUser

`func (o *OrganizationMembershipResponse) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *OrganizationMembershipResponse) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *OrganizationMembershipResponse) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *OrganizationMembershipResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserId

`func (o *OrganizationMembershipResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *OrganizationMembershipResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *OrganizationMembershipResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *OrganizationMembershipResponse) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserName

`func (o *OrganizationMembershipResponse) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *OrganizationMembershipResponse) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *OrganizationMembershipResponse) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *OrganizationMembershipResponse) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetUserUrl

`func (o *OrganizationMembershipResponse) GetUserUrl() string`

GetUserUrl returns the UserUrl field if non-nil, zero value otherwise.

### GetUserUrlOk

`func (o *OrganizationMembershipResponse) GetUserUrlOk() (*string, bool)`

GetUserUrlOk returns a tuple with the UserUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUrl

`func (o *OrganizationMembershipResponse) SetUserUrl(v string)`

SetUserUrl sets UserUrl field to given value.

### HasUserUrl

`func (o *OrganizationMembershipResponse) HasUserUrl() bool`

HasUserUrl returns a boolean if a field has been set.

### GetVisibility

`func (o *OrganizationMembershipResponse) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *OrganizationMembershipResponse) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *OrganizationMembershipResponse) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *OrganizationMembershipResponse) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


