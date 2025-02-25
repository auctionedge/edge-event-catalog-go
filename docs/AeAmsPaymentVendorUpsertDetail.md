# AeAmsPaymentVendorUpsertDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**PaymentVendorId** | **string** | auction unique id for payment vendor in AMS | 
**DisplayName** | **string** | display name of the payment vendor | 
**AuctionAccessId** | Pointer to **NullableString** | AuctionACCESS id for the vendor payment is made with. | [optional] 
**VendorType** | **string** |  | 
**TransactionFee** | [**AeAmsPaymentVendorUpsertDetailTransactionFee**](AeAmsPaymentVendorUpsertDetailTransactionFee.md) |  | 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeAmsPaymentVendorUpsertDetail

`func NewAeAmsPaymentVendorUpsertDetail(auctionId string, paymentVendorId string, displayName string, vendorType string, transactionFee AeAmsPaymentVendorUpsertDetailTransactionFee, ) *AeAmsPaymentVendorUpsertDetail`

NewAeAmsPaymentVendorUpsertDetail instantiates a new AeAmsPaymentVendorUpsertDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAmsPaymentVendorUpsertDetailWithDefaults

`func NewAeAmsPaymentVendorUpsertDetailWithDefaults() *AeAmsPaymentVendorUpsertDetail`

NewAeAmsPaymentVendorUpsertDetailWithDefaults instantiates a new AeAmsPaymentVendorUpsertDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeAmsPaymentVendorUpsertDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAmsPaymentVendorUpsertDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAmsPaymentVendorUpsertDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetPaymentVendorId

`func (o *AeAmsPaymentVendorUpsertDetail) GetPaymentVendorId() string`

GetPaymentVendorId returns the PaymentVendorId field if non-nil, zero value otherwise.

### GetPaymentVendorIdOk

`func (o *AeAmsPaymentVendorUpsertDetail) GetPaymentVendorIdOk() (*string, bool)`

GetPaymentVendorIdOk returns a tuple with the PaymentVendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentVendorId

`func (o *AeAmsPaymentVendorUpsertDetail) SetPaymentVendorId(v string)`

SetPaymentVendorId sets PaymentVendorId field to given value.


### GetDisplayName

`func (o *AeAmsPaymentVendorUpsertDetail) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AeAmsPaymentVendorUpsertDetail) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AeAmsPaymentVendorUpsertDetail) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetAuctionAccessId

`func (o *AeAmsPaymentVendorUpsertDetail) GetAuctionAccessId() string`

GetAuctionAccessId returns the AuctionAccessId field if non-nil, zero value otherwise.

### GetAuctionAccessIdOk

`func (o *AeAmsPaymentVendorUpsertDetail) GetAuctionAccessIdOk() (*string, bool)`

GetAuctionAccessIdOk returns a tuple with the AuctionAccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionAccessId

`func (o *AeAmsPaymentVendorUpsertDetail) SetAuctionAccessId(v string)`

SetAuctionAccessId sets AuctionAccessId field to given value.

### HasAuctionAccessId

`func (o *AeAmsPaymentVendorUpsertDetail) HasAuctionAccessId() bool`

HasAuctionAccessId returns a boolean if a field has been set.

### SetAuctionAccessIdNil

`func (o *AeAmsPaymentVendorUpsertDetail) SetAuctionAccessIdNil(b bool)`

 SetAuctionAccessIdNil sets the value for AuctionAccessId to be an explicit nil

### UnsetAuctionAccessId
`func (o *AeAmsPaymentVendorUpsertDetail) UnsetAuctionAccessId()`

UnsetAuctionAccessId ensures that no value is present for AuctionAccessId, not even an explicit nil
### GetVendorType

`func (o *AeAmsPaymentVendorUpsertDetail) GetVendorType() string`

GetVendorType returns the VendorType field if non-nil, zero value otherwise.

### GetVendorTypeOk

`func (o *AeAmsPaymentVendorUpsertDetail) GetVendorTypeOk() (*string, bool)`

GetVendorTypeOk returns a tuple with the VendorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorType

`func (o *AeAmsPaymentVendorUpsertDetail) SetVendorType(v string)`

SetVendorType sets VendorType field to given value.


### GetTransactionFee

`func (o *AeAmsPaymentVendorUpsertDetail) GetTransactionFee() AeAmsPaymentVendorUpsertDetailTransactionFee`

GetTransactionFee returns the TransactionFee field if non-nil, zero value otherwise.

### GetTransactionFeeOk

`func (o *AeAmsPaymentVendorUpsertDetail) GetTransactionFeeOk() (*AeAmsPaymentVendorUpsertDetailTransactionFee, bool)`

GetTransactionFeeOk returns a tuple with the TransactionFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionFee

`func (o *AeAmsPaymentVendorUpsertDetail) SetTransactionFee(v AeAmsPaymentVendorUpsertDetailTransactionFee)`

SetTransactionFee sets TransactionFee field to given value.


### GetInitiator

`func (o *AeAmsPaymentVendorUpsertDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAmsPaymentVendorUpsertDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAmsPaymentVendorUpsertDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAmsPaymentVendorUpsertDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


