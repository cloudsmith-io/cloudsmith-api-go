# FormatSupportUpstream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthModes** | **[]string** | The authentication modes supported by the upstream format | 
**Caching** | **bool** | If true the upstream format supports caching | 
**Indexing** | **bool** | If true the upstream format supports indexing | 
**IndexingBehavior** | Pointer to **string** | The behavior of the upstream when indexing | [optional] [default to "Unsupported"]
**Proxying** | **bool** | If true the upstream format supports proxying | 
**SignatureVerification** | Pointer to **string** | The signature verification supported by the upstream format | [optional] [default to "Unsupported"]

## Methods

### NewFormatSupportUpstream

`func NewFormatSupportUpstream(authModes []string, caching bool, indexing bool, proxying bool, ) *FormatSupportUpstream`

NewFormatSupportUpstream instantiates a new FormatSupportUpstream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormatSupportUpstreamWithDefaults

`func NewFormatSupportUpstreamWithDefaults() *FormatSupportUpstream`

NewFormatSupportUpstreamWithDefaults instantiates a new FormatSupportUpstream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthModes

`func (o *FormatSupportUpstream) GetAuthModes() []string`

GetAuthModes returns the AuthModes field if non-nil, zero value otherwise.

### GetAuthModesOk

`func (o *FormatSupportUpstream) GetAuthModesOk() (*[]string, bool)`

GetAuthModesOk returns a tuple with the AuthModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthModes

`func (o *FormatSupportUpstream) SetAuthModes(v []string)`

SetAuthModes sets AuthModes field to given value.


### GetCaching

`func (o *FormatSupportUpstream) GetCaching() bool`

GetCaching returns the Caching field if non-nil, zero value otherwise.

### GetCachingOk

`func (o *FormatSupportUpstream) GetCachingOk() (*bool, bool)`

GetCachingOk returns a tuple with the Caching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaching

`func (o *FormatSupportUpstream) SetCaching(v bool)`

SetCaching sets Caching field to given value.


### GetIndexing

`func (o *FormatSupportUpstream) GetIndexing() bool`

GetIndexing returns the Indexing field if non-nil, zero value otherwise.

### GetIndexingOk

`func (o *FormatSupportUpstream) GetIndexingOk() (*bool, bool)`

GetIndexingOk returns a tuple with the Indexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexing

`func (o *FormatSupportUpstream) SetIndexing(v bool)`

SetIndexing sets Indexing field to given value.


### GetIndexingBehavior

`func (o *FormatSupportUpstream) GetIndexingBehavior() string`

GetIndexingBehavior returns the IndexingBehavior field if non-nil, zero value otherwise.

### GetIndexingBehaviorOk

`func (o *FormatSupportUpstream) GetIndexingBehaviorOk() (*string, bool)`

GetIndexingBehaviorOk returns a tuple with the IndexingBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingBehavior

`func (o *FormatSupportUpstream) SetIndexingBehavior(v string)`

SetIndexingBehavior sets IndexingBehavior field to given value.

### HasIndexingBehavior

`func (o *FormatSupportUpstream) HasIndexingBehavior() bool`

HasIndexingBehavior returns a boolean if a field has been set.

### GetProxying

`func (o *FormatSupportUpstream) GetProxying() bool`

GetProxying returns the Proxying field if non-nil, zero value otherwise.

### GetProxyingOk

`func (o *FormatSupportUpstream) GetProxyingOk() (*bool, bool)`

GetProxyingOk returns a tuple with the Proxying field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxying

`func (o *FormatSupportUpstream) SetProxying(v bool)`

SetProxying sets Proxying field to given value.


### GetSignatureVerification

`func (o *FormatSupportUpstream) GetSignatureVerification() string`

GetSignatureVerification returns the SignatureVerification field if non-nil, zero value otherwise.

### GetSignatureVerificationOk

`func (o *FormatSupportUpstream) GetSignatureVerificationOk() (*string, bool)`

GetSignatureVerificationOk returns a tuple with the SignatureVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureVerification

`func (o *FormatSupportUpstream) SetSignatureVerification(v string)`

SetSignatureVerification sets SignatureVerification field to given value.

### HasSignatureVerification

`func (o *FormatSupportUpstream) HasSignatureVerification() bool`

HasSignatureVerification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


