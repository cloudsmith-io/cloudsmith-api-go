# WebhookTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | **string** |  | 
**Template** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewWebhookTemplate

`func NewWebhookTemplate(event string, ) *WebhookTemplate`

NewWebhookTemplate instantiates a new WebhookTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookTemplateWithDefaults

`func NewWebhookTemplateWithDefaults() *WebhookTemplate`

NewWebhookTemplateWithDefaults instantiates a new WebhookTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *WebhookTemplate) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WebhookTemplate) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WebhookTemplate) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetTemplate

`func (o *WebhookTemplate) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *WebhookTemplate) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *WebhookTemplate) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *WebhookTemplate) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### SetTemplateNil

`func (o *WebhookTemplate) SetTemplateNil(b bool)`

 SetTemplateNil sets the value for Template to be an explicit nil

### UnsetTemplate
`func (o *WebhookTemplate) UnsetTemplate()`

UnsetTemplate ensures that no value is present for Template, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


