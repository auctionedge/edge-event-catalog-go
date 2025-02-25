# CommonInitiator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InitiatorType** | **string** | Identify if the event was initiated by an application or a user | 
**InitiatorKey** | **string** | Identifier (or URI) of a system/service component (e.g., code module) or user (username) | 

## Methods

### NewCommonInitiator

`func NewCommonInitiator(initiatorType string, initiatorKey string, ) *CommonInitiator`

NewCommonInitiator instantiates a new CommonInitiator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonInitiatorWithDefaults

`func NewCommonInitiatorWithDefaults() *CommonInitiator`

NewCommonInitiatorWithDefaults instantiates a new CommonInitiator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInitiatorType

`func (o *CommonInitiator) GetInitiatorType() string`

GetInitiatorType returns the InitiatorType field if non-nil, zero value otherwise.

### GetInitiatorTypeOk

`func (o *CommonInitiator) GetInitiatorTypeOk() (*string, bool)`

GetInitiatorTypeOk returns a tuple with the InitiatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorType

`func (o *CommonInitiator) SetInitiatorType(v string)`

SetInitiatorType sets InitiatorType field to given value.


### GetInitiatorKey

`func (o *CommonInitiator) GetInitiatorKey() string`

GetInitiatorKey returns the InitiatorKey field if non-nil, zero value otherwise.

### GetInitiatorKeyOk

`func (o *CommonInitiator) GetInitiatorKeyOk() (*string, bool)`

GetInitiatorKeyOk returns a tuple with the InitiatorKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorKey

`func (o *CommonInitiator) SetInitiatorKey(v string)`

SetInitiatorKey sets InitiatorKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


