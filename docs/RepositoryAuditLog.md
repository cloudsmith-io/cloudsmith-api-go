# RepositoryAuditLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actor** | **string** |  | 
**ActorIpAddress** | **string** |  | 
**ActorKind** | Pointer to **string** |  | [optional] 
**ActorLocation** | **map[string]interface{}** |  | 
**ActorSlugPerm** | **string** |  | 
**ActorUrl** | Pointer to **string** |  | [optional] 
**Context** | **string** |  | 
**Event** | **string** |  | 
**EventAt** | **string** |  | 
**Object** | **string** |  | 
**ObjectKind** | **string** |  | 
**ObjectSlugPerm** | **string** |  | 
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewRepositoryAuditLog

`func NewRepositoryAuditLog(actor string, actorIpAddress string, actorLocation map[string]interface{}, actorSlugPerm string, context string, event string, eventAt string, object string, objectKind string, objectSlugPerm string, ) *RepositoryAuditLog`

NewRepositoryAuditLog instantiates a new RepositoryAuditLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryAuditLogWithDefaults

`func NewRepositoryAuditLogWithDefaults() *RepositoryAuditLog`

NewRepositoryAuditLogWithDefaults instantiates a new RepositoryAuditLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActor

`func (o *RepositoryAuditLog) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *RepositoryAuditLog) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *RepositoryAuditLog) SetActor(v string)`

SetActor sets Actor field to given value.


### GetActorIpAddress

`func (o *RepositoryAuditLog) GetActorIpAddress() string`

GetActorIpAddress returns the ActorIpAddress field if non-nil, zero value otherwise.

### GetActorIpAddressOk

`func (o *RepositoryAuditLog) GetActorIpAddressOk() (*string, bool)`

GetActorIpAddressOk returns a tuple with the ActorIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorIpAddress

`func (o *RepositoryAuditLog) SetActorIpAddress(v string)`

SetActorIpAddress sets ActorIpAddress field to given value.


### GetActorKind

`func (o *RepositoryAuditLog) GetActorKind() string`

GetActorKind returns the ActorKind field if non-nil, zero value otherwise.

### GetActorKindOk

`func (o *RepositoryAuditLog) GetActorKindOk() (*string, bool)`

GetActorKindOk returns a tuple with the ActorKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorKind

`func (o *RepositoryAuditLog) SetActorKind(v string)`

SetActorKind sets ActorKind field to given value.

### HasActorKind

`func (o *RepositoryAuditLog) HasActorKind() bool`

HasActorKind returns a boolean if a field has been set.

### GetActorLocation

`func (o *RepositoryAuditLog) GetActorLocation() map[string]interface{}`

GetActorLocation returns the ActorLocation field if non-nil, zero value otherwise.

### GetActorLocationOk

`func (o *RepositoryAuditLog) GetActorLocationOk() (*map[string]interface{}, bool)`

GetActorLocationOk returns a tuple with the ActorLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorLocation

`func (o *RepositoryAuditLog) SetActorLocation(v map[string]interface{})`

SetActorLocation sets ActorLocation field to given value.


### GetActorSlugPerm

`func (o *RepositoryAuditLog) GetActorSlugPerm() string`

GetActorSlugPerm returns the ActorSlugPerm field if non-nil, zero value otherwise.

### GetActorSlugPermOk

`func (o *RepositoryAuditLog) GetActorSlugPermOk() (*string, bool)`

GetActorSlugPermOk returns a tuple with the ActorSlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorSlugPerm

`func (o *RepositoryAuditLog) SetActorSlugPerm(v string)`

SetActorSlugPerm sets ActorSlugPerm field to given value.


### GetActorUrl

`func (o *RepositoryAuditLog) GetActorUrl() string`

GetActorUrl returns the ActorUrl field if non-nil, zero value otherwise.

### GetActorUrlOk

`func (o *RepositoryAuditLog) GetActorUrlOk() (*string, bool)`

GetActorUrlOk returns a tuple with the ActorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorUrl

`func (o *RepositoryAuditLog) SetActorUrl(v string)`

SetActorUrl sets ActorUrl field to given value.

### HasActorUrl

`func (o *RepositoryAuditLog) HasActorUrl() bool`

HasActorUrl returns a boolean if a field has been set.

### GetContext

`func (o *RepositoryAuditLog) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *RepositoryAuditLog) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *RepositoryAuditLog) SetContext(v string)`

SetContext sets Context field to given value.


### GetEvent

`func (o *RepositoryAuditLog) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *RepositoryAuditLog) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *RepositoryAuditLog) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetEventAt

`func (o *RepositoryAuditLog) GetEventAt() string`

GetEventAt returns the EventAt field if non-nil, zero value otherwise.

### GetEventAtOk

`func (o *RepositoryAuditLog) GetEventAtOk() (*string, bool)`

GetEventAtOk returns a tuple with the EventAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventAt

`func (o *RepositoryAuditLog) SetEventAt(v string)`

SetEventAt sets EventAt field to given value.


### GetObject

`func (o *RepositoryAuditLog) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *RepositoryAuditLog) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *RepositoryAuditLog) SetObject(v string)`

SetObject sets Object field to given value.


### GetObjectKind

`func (o *RepositoryAuditLog) GetObjectKind() string`

GetObjectKind returns the ObjectKind field if non-nil, zero value otherwise.

### GetObjectKindOk

`func (o *RepositoryAuditLog) GetObjectKindOk() (*string, bool)`

GetObjectKindOk returns a tuple with the ObjectKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectKind

`func (o *RepositoryAuditLog) SetObjectKind(v string)`

SetObjectKind sets ObjectKind field to given value.


### GetObjectSlugPerm

`func (o *RepositoryAuditLog) GetObjectSlugPerm() string`

GetObjectSlugPerm returns the ObjectSlugPerm field if non-nil, zero value otherwise.

### GetObjectSlugPermOk

`func (o *RepositoryAuditLog) GetObjectSlugPermOk() (*string, bool)`

GetObjectSlugPermOk returns a tuple with the ObjectSlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSlugPerm

`func (o *RepositoryAuditLog) SetObjectSlugPerm(v string)`

SetObjectSlugPerm sets ObjectSlugPerm field to given value.


### GetUuid

`func (o *RepositoryAuditLog) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RepositoryAuditLog) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RepositoryAuditLog) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RepositoryAuditLog) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


