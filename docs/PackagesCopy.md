# PackagesCopy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | **string** | None | 
**Republish** | Pointer to **bool** | If true, the package will overwrite any others with the same attributes (e.g. same version); otherwise, it will be flagged as a duplicate. | [optional] 

## Methods

### NewPackagesCopy

`func NewPackagesCopy(destination string, ) *PackagesCopy`

NewPackagesCopy instantiates a new PackagesCopy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackagesCopyWithDefaults

`func NewPackagesCopyWithDefaults() *PackagesCopy`

NewPackagesCopyWithDefaults instantiates a new PackagesCopy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *PackagesCopy) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *PackagesCopy) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *PackagesCopy) SetDestination(v string)`

SetDestination sets Destination field to given value.


### GetRepublish

`func (o *PackagesCopy) GetRepublish() bool`

GetRepublish returns the Republish field if non-nil, zero value otherwise.

### GetRepublishOk

`func (o *PackagesCopy) GetRepublishOk() (*bool, bool)`

GetRepublishOk returns a tuple with the Republish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepublish

`func (o *PackagesCopy) SetRepublish(v bool)`

SetRepublish sets Republish field to given value.

### HasRepublish

`func (o *PackagesCopy) HasRepublish() bool`

HasRepublish returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


