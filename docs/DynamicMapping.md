# DynamicMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClaimValue** | **string** | The OIDC token claim value that must be present in the token for it to successfully authenticate as the mapped &#x60;service_account&#x60;.  Note: This field and the dynamic mappings feature are still in early access. Breaking changes are possible as we receive feedback on this feature. | 
**ServiceAccount** | **string** | The service account associated with the provider setting and &#x60;claim_value&#x60;  Note: This field and the dynamic mappings feature are still in early access. Breaking changes are possible as we receive feedback on this feature. | 

## Methods

### NewDynamicMapping

`func NewDynamicMapping(claimValue string, serviceAccount string, ) *DynamicMapping`

NewDynamicMapping instantiates a new DynamicMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicMappingWithDefaults

`func NewDynamicMappingWithDefaults() *DynamicMapping`

NewDynamicMappingWithDefaults instantiates a new DynamicMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaimValue

`func (o *DynamicMapping) GetClaimValue() string`

GetClaimValue returns the ClaimValue field if non-nil, zero value otherwise.

### GetClaimValueOk

`func (o *DynamicMapping) GetClaimValueOk() (*string, bool)`

GetClaimValueOk returns a tuple with the ClaimValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimValue

`func (o *DynamicMapping) SetClaimValue(v string)`

SetClaimValue sets ClaimValue field to given value.


### GetServiceAccount

`func (o *DynamicMapping) GetServiceAccount() string`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *DynamicMapping) GetServiceAccountOk() (*string, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *DynamicMapping) SetServiceAccount(v string)`

SetServiceAccount sets ServiceAccount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


