# CommonCurrency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cents** | **int32** | Amount of money in cents so $175.00 would be 17500 cents. | 
**CurrencyCode** | **string** |  | 

## Methods

### NewCommonCurrency

`func NewCommonCurrency(cents int32, currencyCode string, ) *CommonCurrency`

NewCommonCurrency instantiates a new CommonCurrency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonCurrencyWithDefaults

`func NewCommonCurrencyWithDefaults() *CommonCurrency`

NewCommonCurrencyWithDefaults instantiates a new CommonCurrency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCents

`func (o *CommonCurrency) GetCents() int32`

GetCents returns the Cents field if non-nil, zero value otherwise.

### GetCentsOk

`func (o *CommonCurrency) GetCentsOk() (*int32, bool)`

GetCentsOk returns a tuple with the Cents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCents

`func (o *CommonCurrency) SetCents(v int32)`

SetCents sets Cents field to given value.


### GetCurrencyCode

`func (o *CommonCurrency) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *CommonCurrency) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *CommonCurrency) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


