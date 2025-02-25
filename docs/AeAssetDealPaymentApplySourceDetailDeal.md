# AeAssetDealPaymentApplySourceDetailDeal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique id for a deal | [optional] 
**Buyer** | [**CommonAmsAccountPointer**](CommonAmsAccountPointer.md) |  | 
**Payment** | [**AeAssetDealPaymentApplySourceDetailDealPayment**](AeAssetDealPaymentApplySourceDetailDealPayment.md) |  | 

## Methods

### NewAeAssetDealPaymentApplySourceDetailDeal

`func NewAeAssetDealPaymentApplySourceDetailDeal(buyer CommonAmsAccountPointer, payment AeAssetDealPaymentApplySourceDetailDealPayment, ) *AeAssetDealPaymentApplySourceDetailDeal`

NewAeAssetDealPaymentApplySourceDetailDeal instantiates a new AeAssetDealPaymentApplySourceDetailDeal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetDealPaymentApplySourceDetailDealWithDefaults

`func NewAeAssetDealPaymentApplySourceDetailDealWithDefaults() *AeAssetDealPaymentApplySourceDetailDeal`

NewAeAssetDealPaymentApplySourceDetailDealWithDefaults instantiates a new AeAssetDealPaymentApplySourceDetailDeal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AeAssetDealPaymentApplySourceDetailDeal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AeAssetDealPaymentApplySourceDetailDeal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AeAssetDealPaymentApplySourceDetailDeal) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AeAssetDealPaymentApplySourceDetailDeal) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBuyer

`func (o *AeAssetDealPaymentApplySourceDetailDeal) GetBuyer() CommonAmsAccountPointer`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *AeAssetDealPaymentApplySourceDetailDeal) GetBuyerOk() (*CommonAmsAccountPointer, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *AeAssetDealPaymentApplySourceDetailDeal) SetBuyer(v CommonAmsAccountPointer)`

SetBuyer sets Buyer field to given value.


### GetPayment

`func (o *AeAssetDealPaymentApplySourceDetailDeal) GetPayment() AeAssetDealPaymentApplySourceDetailDealPayment`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *AeAssetDealPaymentApplySourceDetailDeal) GetPaymentOk() (*AeAssetDealPaymentApplySourceDetailDealPayment, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *AeAssetDealPaymentApplySourceDetailDeal) SetPayment(v AeAssetDealPaymentApplySourceDetailDealPayment)`

SetPayment sets Payment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


