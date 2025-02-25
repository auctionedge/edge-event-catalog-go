# AeNotificationRequestDeliveryDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | The message to be delivered | 
**TargetList** | [**[]NotificationTarget**](NotificationTarget.md) |  | 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeNotificationRequestDeliveryDetail

`func NewAeNotificationRequestDeliveryDetail(message string, targetList []NotificationTarget, ) *AeNotificationRequestDeliveryDetail`

NewAeNotificationRequestDeliveryDetail instantiates a new AeNotificationRequestDeliveryDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeNotificationRequestDeliveryDetailWithDefaults

`func NewAeNotificationRequestDeliveryDetailWithDefaults() *AeNotificationRequestDeliveryDetail`

NewAeNotificationRequestDeliveryDetailWithDefaults instantiates a new AeNotificationRequestDeliveryDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *AeNotificationRequestDeliveryDetail) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AeNotificationRequestDeliveryDetail) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AeNotificationRequestDeliveryDetail) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTargetList

`func (o *AeNotificationRequestDeliveryDetail) GetTargetList() []NotificationTarget`

GetTargetList returns the TargetList field if non-nil, zero value otherwise.

### GetTargetListOk

`func (o *AeNotificationRequestDeliveryDetail) GetTargetListOk() (*[]NotificationTarget, bool)`

GetTargetListOk returns a tuple with the TargetList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetList

`func (o *AeNotificationRequestDeliveryDetail) SetTargetList(v []NotificationTarget)`

SetTargetList sets TargetList field to given value.


### GetInitiator

`func (o *AeNotificationRequestDeliveryDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeNotificationRequestDeliveryDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeNotificationRequestDeliveryDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeNotificationRequestDeliveryDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


