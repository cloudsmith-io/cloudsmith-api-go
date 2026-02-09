# PackageBulkActionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The action that was performed. | [optional] [readonly] 
**PackagesActioned** | Pointer to **[]string** | List of package identifiers that were successfully actioned. | [optional] [readonly] 
**PackagesFailedToAction** | Pointer to **map[string]map[string]string** | Dictionary of package identifiers that failed with their error details. | [optional] [readonly] 

## Methods

### NewPackageBulkActionResponse

`func NewPackageBulkActionResponse() *PackageBulkActionResponse`

NewPackageBulkActionResponse instantiates a new PackageBulkActionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageBulkActionResponseWithDefaults

`func NewPackageBulkActionResponseWithDefaults() *PackageBulkActionResponse`

NewPackageBulkActionResponseWithDefaults instantiates a new PackageBulkActionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *PackageBulkActionResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PackageBulkActionResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PackageBulkActionResponse) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *PackageBulkActionResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetPackagesActioned

`func (o *PackageBulkActionResponse) GetPackagesActioned() []string`

GetPackagesActioned returns the PackagesActioned field if non-nil, zero value otherwise.

### GetPackagesActionedOk

`func (o *PackageBulkActionResponse) GetPackagesActionedOk() (*[]string, bool)`

GetPackagesActionedOk returns a tuple with the PackagesActioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackagesActioned

`func (o *PackageBulkActionResponse) SetPackagesActioned(v []string)`

SetPackagesActioned sets PackagesActioned field to given value.

### HasPackagesActioned

`func (o *PackageBulkActionResponse) HasPackagesActioned() bool`

HasPackagesActioned returns a boolean if a field has been set.

### GetPackagesFailedToAction

`func (o *PackageBulkActionResponse) GetPackagesFailedToAction() map[string]map[string]string`

GetPackagesFailedToAction returns the PackagesFailedToAction field if non-nil, zero value otherwise.

### GetPackagesFailedToActionOk

`func (o *PackageBulkActionResponse) GetPackagesFailedToActionOk() (*map[string]map[string]string, bool)`

GetPackagesFailedToActionOk returns a tuple with the PackagesFailedToAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackagesFailedToAction

`func (o *PackageBulkActionResponse) SetPackagesFailedToAction(v map[string]map[string]string)`

SetPackagesFailedToAction sets PackagesFailedToAction field to given value.

### HasPackagesFailedToAction

`func (o *PackageBulkActionResponse) HasPackagesFailedToAction() bool`

HasPackagesFailedToAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


