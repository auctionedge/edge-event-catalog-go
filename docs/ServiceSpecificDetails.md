# ServiceSpecificDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnableWithinDays** | Pointer to **int32** | The number of days available to be returned. | [optional] 
**PricePass** | Pointer to [**CommonCurrency**](CommonCurrency.md) |  | [optional] 
**PriceFail** | Pointer to [**CommonCurrency**](CommonCurrency.md) |  | [optional] 

## Methods

### NewServiceSpecificDetails

`func NewServiceSpecificDetails() *ServiceSpecificDetails`

NewServiceSpecificDetails instantiates a new ServiceSpecificDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceSpecificDetailsWithDefaults

`func NewServiceSpecificDetailsWithDefaults() *ServiceSpecificDetails`

NewServiceSpecificDetailsWithDefaults instantiates a new ServiceSpecificDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnableWithinDays

`func (o *ServiceSpecificDetails) GetReturnableWithinDays() int32`

GetReturnableWithinDays returns the ReturnableWithinDays field if non-nil, zero value otherwise.

### GetReturnableWithinDaysOk

`func (o *ServiceSpecificDetails) GetReturnableWithinDaysOk() (*int32, bool)`

GetReturnableWithinDaysOk returns a tuple with the ReturnableWithinDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnableWithinDays

`func (o *ServiceSpecificDetails) SetReturnableWithinDays(v int32)`

SetReturnableWithinDays sets ReturnableWithinDays field to given value.

### HasReturnableWithinDays

`func (o *ServiceSpecificDetails) HasReturnableWithinDays() bool`

HasReturnableWithinDays returns a boolean if a field has been set.

### GetPricePass

`func (o *ServiceSpecificDetails) GetPricePass() CommonCurrency`

GetPricePass returns the PricePass field if non-nil, zero value otherwise.

### GetPricePassOk

`func (o *ServiceSpecificDetails) GetPricePassOk() (*CommonCurrency, bool)`

GetPricePassOk returns a tuple with the PricePass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePass

`func (o *ServiceSpecificDetails) SetPricePass(v CommonCurrency)`

SetPricePass sets PricePass field to given value.

### HasPricePass

`func (o *ServiceSpecificDetails) HasPricePass() bool`

HasPricePass returns a boolean if a field has been set.

### GetPriceFail

`func (o *ServiceSpecificDetails) GetPriceFail() CommonCurrency`

GetPriceFail returns the PriceFail field if non-nil, zero value otherwise.

### GetPriceFailOk

`func (o *ServiceSpecificDetails) GetPriceFailOk() (*CommonCurrency, bool)`

GetPriceFailOk returns a tuple with the PriceFail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceFail

`func (o *ServiceSpecificDetails) SetPriceFail(v CommonCurrency)`

SetPriceFail sets PriceFail field to given value.

### HasPriceFail

`func (o *ServiceSpecificDetails) HasPriceFail() bool`

HasPriceFail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


