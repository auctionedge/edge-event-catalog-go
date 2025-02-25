# NotificationTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeliveryMethod** | **string** | The method to deliver the notification | 
**DeliveryAddress** | **string** | The address to deliver the notification to | 

## Methods

### NewNotificationTarget

`func NewNotificationTarget(deliveryMethod string, deliveryAddress string, ) *NotificationTarget`

NewNotificationTarget instantiates a new NotificationTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationTargetWithDefaults

`func NewNotificationTargetWithDefaults() *NotificationTarget`

NewNotificationTargetWithDefaults instantiates a new NotificationTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeliveryMethod

`func (o *NotificationTarget) GetDeliveryMethod() string`

GetDeliveryMethod returns the DeliveryMethod field if non-nil, zero value otherwise.

### GetDeliveryMethodOk

`func (o *NotificationTarget) GetDeliveryMethodOk() (*string, bool)`

GetDeliveryMethodOk returns a tuple with the DeliveryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryMethod

`func (o *NotificationTarget) SetDeliveryMethod(v string)`

SetDeliveryMethod sets DeliveryMethod field to given value.


### GetDeliveryAddress

`func (o *NotificationTarget) GetDeliveryAddress() string`

GetDeliveryAddress returns the DeliveryAddress field if non-nil, zero value otherwise.

### GetDeliveryAddressOk

`func (o *NotificationTarget) GetDeliveryAddressOk() (*string, bool)`

GetDeliveryAddressOk returns a tuple with the DeliveryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryAddress

`func (o *NotificationTarget) SetDeliveryAddress(v string)`

SetDeliveryAddress sets DeliveryAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


