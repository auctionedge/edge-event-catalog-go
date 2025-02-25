# Deal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmsSaleListing** | [**DealAmsSaleListing**](DealAmsSaleListing.md) |  | 
**Buyer** | [**CommonAmsAccountPointer**](CommonAmsAccountPointer.md) |  | 
**SoldAt** | Pointer to **NullableTime** | The date time that asset was sold at. | [optional] 
**SoldAmountUsd** | **float32** | The amount that the asset was sold for. | 
**OnOffer** | Pointer to **bool** | Whether or not the current deal is \&quot;on offer\&quot; / \&quot;on IF\&quot; | [optional] [default to false]
**InArbitration** | Pointer to **bool** | Whether or not the current deal is being arbitrated. | [optional] [default to false]

## Methods

### NewDeal

`func NewDeal(amsSaleListing DealAmsSaleListing, buyer CommonAmsAccountPointer, soldAmountUsd float32, ) *Deal`

NewDeal instantiates a new Deal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDealWithDefaults

`func NewDealWithDefaults() *Deal`

NewDealWithDefaults instantiates a new Deal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmsSaleListing

`func (o *Deal) GetAmsSaleListing() DealAmsSaleListing`

GetAmsSaleListing returns the AmsSaleListing field if non-nil, zero value otherwise.

### GetAmsSaleListingOk

`func (o *Deal) GetAmsSaleListingOk() (*DealAmsSaleListing, bool)`

GetAmsSaleListingOk returns a tuple with the AmsSaleListing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmsSaleListing

`func (o *Deal) SetAmsSaleListing(v DealAmsSaleListing)`

SetAmsSaleListing sets AmsSaleListing field to given value.


### GetBuyer

`func (o *Deal) GetBuyer() CommonAmsAccountPointer`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *Deal) GetBuyerOk() (*CommonAmsAccountPointer, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *Deal) SetBuyer(v CommonAmsAccountPointer)`

SetBuyer sets Buyer field to given value.


### GetSoldAt

`func (o *Deal) GetSoldAt() time.Time`

GetSoldAt returns the SoldAt field if non-nil, zero value otherwise.

### GetSoldAtOk

`func (o *Deal) GetSoldAtOk() (*time.Time, bool)`

GetSoldAtOk returns a tuple with the SoldAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoldAt

`func (o *Deal) SetSoldAt(v time.Time)`

SetSoldAt sets SoldAt field to given value.

### HasSoldAt

`func (o *Deal) HasSoldAt() bool`

HasSoldAt returns a boolean if a field has been set.

### SetSoldAtNil

`func (o *Deal) SetSoldAtNil(b bool)`

 SetSoldAtNil sets the value for SoldAt to be an explicit nil

### UnsetSoldAt
`func (o *Deal) UnsetSoldAt()`

UnsetSoldAt ensures that no value is present for SoldAt, not even an explicit nil
### GetSoldAmountUsd

`func (o *Deal) GetSoldAmountUsd() float32`

GetSoldAmountUsd returns the SoldAmountUsd field if non-nil, zero value otherwise.

### GetSoldAmountUsdOk

`func (o *Deal) GetSoldAmountUsdOk() (*float32, bool)`

GetSoldAmountUsdOk returns a tuple with the SoldAmountUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoldAmountUsd

`func (o *Deal) SetSoldAmountUsd(v float32)`

SetSoldAmountUsd sets SoldAmountUsd field to given value.


### GetOnOffer

`func (o *Deal) GetOnOffer() bool`

GetOnOffer returns the OnOffer field if non-nil, zero value otherwise.

### GetOnOfferOk

`func (o *Deal) GetOnOfferOk() (*bool, bool)`

GetOnOfferOk returns a tuple with the OnOffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnOffer

`func (o *Deal) SetOnOffer(v bool)`

SetOnOffer sets OnOffer field to given value.

### HasOnOffer

`func (o *Deal) HasOnOffer() bool`

HasOnOffer returns a boolean if a field has been set.

### GetInArbitration

`func (o *Deal) GetInArbitration() bool`

GetInArbitration returns the InArbitration field if non-nil, zero value otherwise.

### GetInArbitrationOk

`func (o *Deal) GetInArbitrationOk() (*bool, bool)`

GetInArbitrationOk returns a tuple with the InArbitration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInArbitration

`func (o *Deal) SetInArbitration(v bool)`

SetInArbitration sets InArbitration field to given value.

### HasInArbitration

`func (o *Deal) HasInArbitration() bool`

HasInArbitration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


