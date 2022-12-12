# RepositoryTokenSync

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tokens** | Pointer to [**[]RepositoryToken**](RepositoryToken.md) | The entitlements that have been synchronised. | [optional] [readonly] 

## Methods

### NewRepositoryTokenSync

`func NewRepositoryTokenSync() *RepositoryTokenSync`

NewRepositoryTokenSync instantiates a new RepositoryTokenSync object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryTokenSyncWithDefaults

`func NewRepositoryTokenSyncWithDefaults() *RepositoryTokenSync`

NewRepositoryTokenSyncWithDefaults instantiates a new RepositoryTokenSync object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokens

`func (o *RepositoryTokenSync) GetTokens() []RepositoryToken`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *RepositoryTokenSync) GetTokensOk() (*[]RepositoryToken, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *RepositoryTokenSync) SetTokens(v []RepositoryToken)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *RepositoryTokenSync) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


