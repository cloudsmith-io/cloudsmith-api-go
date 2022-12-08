# RepositoryAuditLogResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actor** | **NullableString** |  | 
**ActorIpAddress** | **NullableString** |  | 
**ActorKind** | Pointer to **string** |  | [optional] [readonly] 
**ActorLocation** | [**GeoIpLocation**](GeoIpLocation.md) |  | 
**ActorSlugPerm** | **NullableString** |  | 
**ActorUrl** | Pointer to **NullableString** |  | [optional] [readonly] 
**Context** | **string** |  | 
**Event** | **string** |  | 
**EventAt** | **time.Time** |  | 
**Object** | **string** |  | 
**ObjectKind** | **string** |  | 
**ObjectSlugPerm** | **string** |  | 
**Uuid** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewRepositoryAuditLogResponse

`func NewRepositoryAuditLogResponse(actor NullableString, actorIpAddress NullableString, actorLocation GeoIpLocation, actorSlugPerm NullableString, context string, event string, eventAt time.Time, object string, objectKind string, objectSlugPerm string, ) *RepositoryAuditLogResponse`

NewRepositoryAuditLogResponse instantiates a new RepositoryAuditLogResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryAuditLogResponseWithDefaults

`func NewRepositoryAuditLogResponseWithDefaults() *RepositoryAuditLogResponse`

NewRepositoryAuditLogResponseWithDefaults instantiates a new RepositoryAuditLogResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActor

`func (o *RepositoryAuditLogResponse) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *RepositoryAuditLogResponse) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *RepositoryAuditLogResponse) SetActor(v string)`

SetActor sets Actor field to given value.


### SetActorNil

`func (o *RepositoryAuditLogResponse) SetActorNil(b bool)`

 SetActorNil sets the value for Actor to be an explicit nil

### UnsetActor
`func (o *RepositoryAuditLogResponse) UnsetActor()`

UnsetActor ensures that no value is present for Actor, not even an explicit nil
### GetActorIpAddress

`func (o *RepositoryAuditLogResponse) GetActorIpAddress() string`

GetActorIpAddress returns the ActorIpAddress field if non-nil, zero value otherwise.

### GetActorIpAddressOk

`func (o *RepositoryAuditLogResponse) GetActorIpAddressOk() (*string, bool)`

GetActorIpAddressOk returns a tuple with the ActorIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorIpAddress

`func (o *RepositoryAuditLogResponse) SetActorIpAddress(v string)`

SetActorIpAddress sets ActorIpAddress field to given value.


### SetActorIpAddressNil

`func (o *RepositoryAuditLogResponse) SetActorIpAddressNil(b bool)`

 SetActorIpAddressNil sets the value for ActorIpAddress to be an explicit nil

### UnsetActorIpAddress
`func (o *RepositoryAuditLogResponse) UnsetActorIpAddress()`

UnsetActorIpAddress ensures that no value is present for ActorIpAddress, not even an explicit nil
### GetActorKind

`func (o *RepositoryAuditLogResponse) GetActorKind() string`

GetActorKind returns the ActorKind field if non-nil, zero value otherwise.

### GetActorKindOk

`func (o *RepositoryAuditLogResponse) GetActorKindOk() (*string, bool)`

GetActorKindOk returns a tuple with the ActorKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorKind

`func (o *RepositoryAuditLogResponse) SetActorKind(v string)`

SetActorKind sets ActorKind field to given value.

### HasActorKind

`func (o *RepositoryAuditLogResponse) HasActorKind() bool`

HasActorKind returns a boolean if a field has been set.

### GetActorLocation

`func (o *RepositoryAuditLogResponse) GetActorLocation() GeoIpLocation`

GetActorLocation returns the ActorLocation field if non-nil, zero value otherwise.

### GetActorLocationOk

`func (o *RepositoryAuditLogResponse) GetActorLocationOk() (*GeoIpLocation, bool)`

GetActorLocationOk returns a tuple with the ActorLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorLocation

`func (o *RepositoryAuditLogResponse) SetActorLocation(v GeoIpLocation)`

SetActorLocation sets ActorLocation field to given value.


### GetActorSlugPerm

`func (o *RepositoryAuditLogResponse) GetActorSlugPerm() string`

GetActorSlugPerm returns the ActorSlugPerm field if non-nil, zero value otherwise.

### GetActorSlugPermOk

`func (o *RepositoryAuditLogResponse) GetActorSlugPermOk() (*string, bool)`

GetActorSlugPermOk returns a tuple with the ActorSlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorSlugPerm

`func (o *RepositoryAuditLogResponse) SetActorSlugPerm(v string)`

SetActorSlugPerm sets ActorSlugPerm field to given value.


### SetActorSlugPermNil

`func (o *RepositoryAuditLogResponse) SetActorSlugPermNil(b bool)`

 SetActorSlugPermNil sets the value for ActorSlugPerm to be an explicit nil

### UnsetActorSlugPerm
`func (o *RepositoryAuditLogResponse) UnsetActorSlugPerm()`

UnsetActorSlugPerm ensures that no value is present for ActorSlugPerm, not even an explicit nil
### GetActorUrl

`func (o *RepositoryAuditLogResponse) GetActorUrl() string`

GetActorUrl returns the ActorUrl field if non-nil, zero value otherwise.

### GetActorUrlOk

`func (o *RepositoryAuditLogResponse) GetActorUrlOk() (*string, bool)`

GetActorUrlOk returns a tuple with the ActorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorUrl

`func (o *RepositoryAuditLogResponse) SetActorUrl(v string)`

SetActorUrl sets ActorUrl field to given value.

### HasActorUrl

`func (o *RepositoryAuditLogResponse) HasActorUrl() bool`

HasActorUrl returns a boolean if a field has been set.

### SetActorUrlNil

`func (o *RepositoryAuditLogResponse) SetActorUrlNil(b bool)`

 SetActorUrlNil sets the value for ActorUrl to be an explicit nil

### UnsetActorUrl
`func (o *RepositoryAuditLogResponse) UnsetActorUrl()`

UnsetActorUrl ensures that no value is present for ActorUrl, not even an explicit nil
### GetContext

`func (o *RepositoryAuditLogResponse) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *RepositoryAuditLogResponse) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *RepositoryAuditLogResponse) SetContext(v string)`

SetContext sets Context field to given value.


### GetEvent

`func (o *RepositoryAuditLogResponse) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *RepositoryAuditLogResponse) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *RepositoryAuditLogResponse) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetEventAt

`func (o *RepositoryAuditLogResponse) GetEventAt() time.Time`

GetEventAt returns the EventAt field if non-nil, zero value otherwise.

### GetEventAtOk

`func (o *RepositoryAuditLogResponse) GetEventAtOk() (*time.Time, bool)`

GetEventAtOk returns a tuple with the EventAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventAt

`func (o *RepositoryAuditLogResponse) SetEventAt(v time.Time)`

SetEventAt sets EventAt field to given value.


### GetObject

`func (o *RepositoryAuditLogResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *RepositoryAuditLogResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *RepositoryAuditLogResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetObjectKind

`func (o *RepositoryAuditLogResponse) GetObjectKind() string`

GetObjectKind returns the ObjectKind field if non-nil, zero value otherwise.

### GetObjectKindOk

`func (o *RepositoryAuditLogResponse) GetObjectKindOk() (*string, bool)`

GetObjectKindOk returns a tuple with the ObjectKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectKind

`func (o *RepositoryAuditLogResponse) SetObjectKind(v string)`

SetObjectKind sets ObjectKind field to given value.


### GetObjectSlugPerm

`func (o *RepositoryAuditLogResponse) GetObjectSlugPerm() string`

GetObjectSlugPerm returns the ObjectSlugPerm field if non-nil, zero value otherwise.

### GetObjectSlugPermOk

`func (o *RepositoryAuditLogResponse) GetObjectSlugPermOk() (*string, bool)`

GetObjectSlugPermOk returns a tuple with the ObjectSlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSlugPerm

`func (o *RepositoryAuditLogResponse) SetObjectSlugPerm(v string)`

SetObjectSlugPerm sets ObjectSlugPerm field to given value.


### GetUuid

`func (o *RepositoryAuditLogResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RepositoryAuditLogResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RepositoryAuditLogResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RepositoryAuditLogResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


