# CommonAssetBuyNow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuyNowPrice** | [**CommonCurrency**](CommonCurrency.md) |  | 
**BuyNowValidAt** | **time.Time** | The time that this buy-now price is available for the asset | 
**BuyNowValidUntil** | **time.Time** | The time that this buy-now price for the asset expires. | 

## Methods

### NewCommonAssetBuyNow

`func NewCommonAssetBuyNow(buyNowPrice CommonCurrency, buyNowValidAt time.Time, buyNowValidUntil time.Time, ) *CommonAssetBuyNow`

NewCommonAssetBuyNow instantiates a new CommonAssetBuyNow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonAssetBuyNowWithDefaults

`func NewCommonAssetBuyNowWithDefaults() *CommonAssetBuyNow`

NewCommonAssetBuyNowWithDefaults instantiates a new CommonAssetBuyNow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuyNowPrice

`func (o *CommonAssetBuyNow) GetBuyNowPrice() CommonCurrency`

GetBuyNowPrice returns the BuyNowPrice field if non-nil, zero value otherwise.

### GetBuyNowPriceOk

`func (o *CommonAssetBuyNow) GetBuyNowPriceOk() (*CommonCurrency, bool)`

GetBuyNowPriceOk returns a tuple with the BuyNowPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyNowPrice

`func (o *CommonAssetBuyNow) SetBuyNowPrice(v CommonCurrency)`

SetBuyNowPrice sets BuyNowPrice field to given value.


### GetBuyNowValidAt

`func (o *CommonAssetBuyNow) GetBuyNowValidAt() time.Time`

GetBuyNowValidAt returns the BuyNowValidAt field if non-nil, zero value otherwise.

### GetBuyNowValidAtOk

`func (o *CommonAssetBuyNow) GetBuyNowValidAtOk() (*time.Time, bool)`

GetBuyNowValidAtOk returns a tuple with the BuyNowValidAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyNowValidAt

`func (o *CommonAssetBuyNow) SetBuyNowValidAt(v time.Time)`

SetBuyNowValidAt sets BuyNowValidAt field to given value.


### GetBuyNowValidUntil

`func (o *CommonAssetBuyNow) GetBuyNowValidUntil() time.Time`

GetBuyNowValidUntil returns the BuyNowValidUntil field if non-nil, zero value otherwise.

### GetBuyNowValidUntilOk

`func (o *CommonAssetBuyNow) GetBuyNowValidUntilOk() (*time.Time, bool)`

GetBuyNowValidUntilOk returns a tuple with the BuyNowValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyNowValidUntil

`func (o *CommonAssetBuyNow) SetBuyNowValidUntil(v time.Time)`

SetBuyNowValidUntil sets BuyNowValidUntil field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


