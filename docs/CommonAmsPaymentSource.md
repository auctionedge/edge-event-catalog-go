# CommonAmsPaymentSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentSourceId** | **string** | auction unique id for payment source in AMS | 
**PaymentVendorId** | **string** | auction unique id for payment vendor in AMS | 
**DisplayName** | **string** | The display name of the payment source. | 

## Methods

### NewCommonAmsPaymentSource

`func NewCommonAmsPaymentSource(paymentSourceId string, paymentVendorId string, displayName string, ) *CommonAmsPaymentSource`

NewCommonAmsPaymentSource instantiates a new CommonAmsPaymentSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonAmsPaymentSourceWithDefaults

`func NewCommonAmsPaymentSourceWithDefaults() *CommonAmsPaymentSource`

NewCommonAmsPaymentSourceWithDefaults instantiates a new CommonAmsPaymentSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentSourceId

`func (o *CommonAmsPaymentSource) GetPaymentSourceId() string`

GetPaymentSourceId returns the PaymentSourceId field if non-nil, zero value otherwise.

### GetPaymentSourceIdOk

`func (o *CommonAmsPaymentSource) GetPaymentSourceIdOk() (*string, bool)`

GetPaymentSourceIdOk returns a tuple with the PaymentSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSourceId

`func (o *CommonAmsPaymentSource) SetPaymentSourceId(v string)`

SetPaymentSourceId sets PaymentSourceId field to given value.


### GetPaymentVendorId

`func (o *CommonAmsPaymentSource) GetPaymentVendorId() string`

GetPaymentVendorId returns the PaymentVendorId field if non-nil, zero value otherwise.

### GetPaymentVendorIdOk

`func (o *CommonAmsPaymentSource) GetPaymentVendorIdOk() (*string, bool)`

GetPaymentVendorIdOk returns a tuple with the PaymentVendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentVendorId

`func (o *CommonAmsPaymentSource) SetPaymentVendorId(v string)`

SetPaymentVendorId sets PaymentVendorId field to given value.


### GetDisplayName

`func (o *CommonAmsPaymentSource) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CommonAmsPaymentSource) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CommonAmsPaymentSource) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


