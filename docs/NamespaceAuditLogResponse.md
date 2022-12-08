# NamespaceAuditLogResponse

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
**Target** | **string** |  | 
**TargetKind** | **string** |  | 
**TargetSlugPerm** | Pointer to **NullableString** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewNamespaceAuditLogResponse

`func NewNamespaceAuditLogResponse(actor NullableString, actorIpAddress NullableString, actorLocation GeoIpLocation, actorSlugPerm NullableString, context string, event string, eventAt time.Time, object string, objectKind string, objectSlugPerm string, target string, targetKind string, ) *NamespaceAuditLogResponse`

NewNamespaceAuditLogResponse instantiates a new NamespaceAuditLogResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamespaceAuditLogResponseWithDefaults

`func NewNamespaceAuditLogResponseWithDefaults() *NamespaceAuditLogResponse`

NewNamespaceAuditLogResponseWithDefaults instantiates a new NamespaceAuditLogResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActor

`func (o *NamespaceAuditLogResponse) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *NamespaceAuditLogResponse) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *NamespaceAuditLogResponse) SetActor(v string)`

SetActor sets Actor field to given value.


### SetActorNil

`func (o *NamespaceAuditLogResponse) SetActorNil(b bool)`

 SetActorNil sets the value for Actor to be an explicit nil

### UnsetActor
`func (o *NamespaceAuditLogResponse) UnsetActor()`

UnsetActor ensures that no value is present for Actor, not even an explicit nil
### GetActorIpAddress

`func (o *NamespaceAuditLogResponse) GetActorIpAddress() string`

GetActorIpAddress returns the ActorIpAddress field if non-nil, zero value otherwise.

### GetActorIpAddressOk

`func (o *NamespaceAuditLogResponse) GetActorIpAddressOk() (*string, bool)`

GetActorIpAddressOk returns a tuple with the ActorIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorIpAddress

`func (o *NamespaceAuditLogResponse) SetActorIpAddress(v string)`

SetActorIpAddress sets ActorIpAddress field to given value.


### SetActorIpAddressNil

`func (o *NamespaceAuditLogResponse) SetActorIpAddressNil(b bool)`

 SetActorIpAddressNil sets the value for ActorIpAddress to be an explicit nil

### UnsetActorIpAddress
`func (o *NamespaceAuditLogResponse) UnsetActorIpAddress()`

UnsetActorIpAddress ensures that no value is present for ActorIpAddress, not even an explicit nil
### GetActorKind

`func (o *NamespaceAuditLogResponse) GetActorKind() string`

GetActorKind returns the ActorKind field if non-nil, zero value otherwise.

### GetActorKindOk

`func (o *NamespaceAuditLogResponse) GetActorKindOk() (*string, bool)`

GetActorKindOk returns a tuple with the ActorKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorKind

`func (o *NamespaceAuditLogResponse) SetActorKind(v string)`

SetActorKind sets ActorKind field to given value.

### HasActorKind

`func (o *NamespaceAuditLogResponse) HasActorKind() bool`

HasActorKind returns a boolean if a field has been set.

### GetActorLocation

`func (o *NamespaceAuditLogResponse) GetActorLocation() GeoIpLocation`

GetActorLocation returns the ActorLocation field if non-nil, zero value otherwise.

### GetActorLocationOk

`func (o *NamespaceAuditLogResponse) GetActorLocationOk() (*GeoIpLocation, bool)`

GetActorLocationOk returns a tuple with the ActorLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorLocation

`func (o *NamespaceAuditLogResponse) SetActorLocation(v GeoIpLocation)`

SetActorLocation sets ActorLocation field to given value.


### GetActorSlugPerm

`func (o *NamespaceAuditLogResponse) GetActorSlugPerm() string`

GetActorSlugPerm returns the ActorSlugPerm field if non-nil, zero value otherwise.

### GetActorSlugPermOk

`func (o *NamespaceAuditLogResponse) GetActorSlugPermOk() (*string, bool)`

GetActorSlugPermOk returns a tuple with the ActorSlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorSlugPerm

`func (o *NamespaceAuditLogResponse) SetActorSlugPerm(v string)`

SetActorSlugPerm sets ActorSlugPerm field to given value.


### SetActorSlugPermNil

`func (o *NamespaceAuditLogResponse) SetActorSlugPermNil(b bool)`

 SetActorSlugPermNil sets the value for ActorSlugPerm to be an explicit nil

### UnsetActorSlugPerm
`func (o *NamespaceAuditLogResponse) UnsetActorSlugPerm()`

UnsetActorSlugPerm ensures that no value is present for ActorSlugPerm, not even an explicit nil
### GetActorUrl

`func (o *NamespaceAuditLogResponse) GetActorUrl() string`

GetActorUrl returns the ActorUrl field if non-nil, zero value otherwise.

### GetActorUrlOk

`func (o *NamespaceAuditLogResponse) GetActorUrlOk() (*string, bool)`

GetActorUrlOk returns a tuple with the ActorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorUrl

`func (o *NamespaceAuditLogResponse) SetActorUrl(v string)`

SetActorUrl sets ActorUrl field to given value.

### HasActorUrl

`func (o *NamespaceAuditLogResponse) HasActorUrl() bool`

HasActorUrl returns a boolean if a field has been set.

### SetActorUrlNil

`func (o *NamespaceAuditLogResponse) SetActorUrlNil(b bool)`

 SetActorUrlNil sets the value for ActorUrl to be an explicit nil

### UnsetActorUrl
`func (o *NamespaceAuditLogResponse) UnsetActorUrl()`

UnsetActorUrl ensures that no value is present for ActorUrl, not even an explicit nil
### GetContext

`func (o *NamespaceAuditLogResponse) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *NamespaceAuditLogResponse) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *NamespaceAuditLogResponse) SetContext(v string)`

SetContext sets Context field to given value.


### GetEvent

`func (o *NamespaceAuditLogResponse) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *NamespaceAuditLogResponse) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *NamespaceAuditLogResponse) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetEventAt

`func (o *NamespaceAuditLogResponse) GetEventAt() time.Time`

GetEventAt returns the EventAt field if non-nil, zero value otherwise.

### GetEventAtOk

`func (o *NamespaceAuditLogResponse) GetEventAtOk() (*time.Time, bool)`

GetEventAtOk returns a tuple with the EventAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventAt

`func (o *NamespaceAuditLogResponse) SetEventAt(v time.Time)`

SetEventAt sets EventAt field to given value.


### GetObject

`func (o *NamespaceAuditLogResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *NamespaceAuditLogResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *NamespaceAuditLogResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetObjectKind

`func (o *NamespaceAuditLogResponse) GetObjectKind() string`

GetObjectKind returns the ObjectKind field if non-nil, zero value otherwise.

### GetObjectKindOk

`func (o *NamespaceAuditLogResponse) GetObjectKindOk() (*string, bool)`

GetObjectKindOk returns a tuple with the ObjectKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectKind

`func (o *NamespaceAuditLogResponse) SetObjectKind(v string)`

SetObjectKind sets ObjectKind field to given value.


### GetObjectSlugPerm

`func (o *NamespaceAuditLogResponse) GetObjectSlugPerm() string`

GetObjectSlugPerm returns the ObjectSlugPerm field if non-nil, zero value otherwise.

### GetObjectSlugPermOk

`func (o *NamespaceAuditLogResponse) GetObjectSlugPermOk() (*string, bool)`

GetObjectSlugPermOk returns a tuple with the ObjectSlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSlugPerm

`func (o *NamespaceAuditLogResponse) SetObjectSlugPerm(v string)`

SetObjectSlugPerm sets ObjectSlugPerm field to given value.


### GetTarget

`func (o *NamespaceAuditLogResponse) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *NamespaceAuditLogResponse) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *NamespaceAuditLogResponse) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetTargetKind

`func (o *NamespaceAuditLogResponse) GetTargetKind() string`

GetTargetKind returns the TargetKind field if non-nil, zero value otherwise.

### GetTargetKindOk

`func (o *NamespaceAuditLogResponse) GetTargetKindOk() (*string, bool)`

GetTargetKindOk returns a tuple with the TargetKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetKind

`func (o *NamespaceAuditLogResponse) SetTargetKind(v string)`

SetTargetKind sets TargetKind field to given value.


### GetTargetSlugPerm

`func (o *NamespaceAuditLogResponse) GetTargetSlugPerm() string`

GetTargetSlugPerm returns the TargetSlugPerm field if non-nil, zero value otherwise.

### GetTargetSlugPermOk

`func (o *NamespaceAuditLogResponse) GetTargetSlugPermOk() (*string, bool)`

GetTargetSlugPermOk returns a tuple with the TargetSlugPerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSlugPerm

`func (o *NamespaceAuditLogResponse) SetTargetSlugPerm(v string)`

SetTargetSlugPerm sets TargetSlugPerm field to given value.

### HasTargetSlugPerm

`func (o *NamespaceAuditLogResponse) HasTargetSlugPerm() bool`

HasTargetSlugPerm returns a boolean if a field has been set.

### SetTargetSlugPermNil

`func (o *NamespaceAuditLogResponse) SetTargetSlugPermNil(b bool)`

 SetTargetSlugPermNil sets the value for TargetSlugPerm to be an explicit nil

### UnsetTargetSlugPerm
`func (o *NamespaceAuditLogResponse) UnsetTargetSlugPerm()`

UnsetTargetSlugPerm ensures that no value is present for TargetSlugPerm, not even an explicit nil
### GetUuid

`func (o *NamespaceAuditLogResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *NamespaceAuditLogResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *NamespaceAuditLogResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *NamespaceAuditLogResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


