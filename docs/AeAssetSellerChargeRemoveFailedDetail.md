# AeAssetSellerChargeRemoveFailedDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**Asset** | [**CommonAmsAssetPointer**](CommonAmsAssetPointer.md) |  | 
**Seller** | [**CommonAmsAccountPointer**](CommonAmsAccountPointer.md) |  | 
**ChargeId** | **string** | Unique identifier of the charge to be removed. | 
**RemoveRequestReferenceId** | Pointer to **NullableString** | Unique identifier referencing a remove request event that requested the charge be removed | [optional] 
**Reason** | **string** | An explaination for why the charge could not be added. | 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeAssetSellerChargeRemoveFailedDetail

`func NewAeAssetSellerChargeRemoveFailedDetail(auctionId string, asset CommonAmsAssetPointer, seller CommonAmsAccountPointer, chargeId string, reason string, ) *AeAssetSellerChargeRemoveFailedDetail`

NewAeAssetSellerChargeRemoveFailedDetail instantiates a new AeAssetSellerChargeRemoveFailedDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetSellerChargeRemoveFailedDetailWithDefaults

`func NewAeAssetSellerChargeRemoveFailedDetailWithDefaults() *AeAssetSellerChargeRemoveFailedDetail`

NewAeAssetSellerChargeRemoveFailedDetailWithDefaults instantiates a new AeAssetSellerChargeRemoveFailedDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeAssetSellerChargeRemoveFailedDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAssetSellerChargeRemoveFailedDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAssetSellerChargeRemoveFailedDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetAsset

`func (o *AeAssetSellerChargeRemoveFailedDetail) GetAsset() CommonAmsAssetPointer`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *AeAssetSellerChargeRemoveFailedDetail) GetAssetOk() (*CommonAmsAssetPointer, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *AeAssetSellerChargeRemoveFailedDetail) SetAsset(v CommonAmsAssetPointer)`

SetAsset sets Asset field to given value.


### GetSeller

`func (o *AeAssetSellerChargeRemoveFailedDetail) GetSeller() CommonAmsAccountPointer`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *AeAssetSellerChargeRemoveFailedDetail) GetSellerOk() (*CommonAmsAccountPointer, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *AeAssetSellerChargeRemoveFailedDetail) SetSeller(v CommonAmsAccountPointer)`

SetSeller sets Seller field to given value.


### GetChargeId

`func (o *AeAssetSellerChargeRemoveFailedDetail) GetChargeId() string`

GetChargeId returns the ChargeId field if non-nil, zero value otherwise.

### GetChargeIdOk

`func (o *AeAssetSellerChargeRemoveFailedDetail) GetChargeIdOk() (*string, bool)`

GetChargeIdOk returns a tuple with the ChargeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeId

`func (o *AeAssetSellerChargeRemoveFailedDetail) SetChargeId(v string)`

SetChargeId sets ChargeId field to given value.


### GetRemoveRequestReferenceId

`func (o *AeAssetSellerChargeRemoveFailedDetail) GetRemoveRequestReferenceId() string`

GetRemoveRequestReferenceId returns the RemoveRequestReferenceId field if non-nil, zero value otherwise.

### GetRemoveRequestReferenceIdOk

`func (o *AeAssetSellerChargeRemoveFailedDetail) GetRemoveRequestReferenceIdOk() (*string, bool)`

GetRemoveRequestReferenceIdOk returns a tuple with the RemoveRequestReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveRequestReferenceId

`func (o *AeAssetSellerChargeRemoveFailedDetail) SetRemoveRequestReferenceId(v string)`

SetRemoveRequestReferenceId sets RemoveRequestReferenceId field to given value.

### HasRemoveRequestReferenceId

`func (o *AeAssetSellerChargeRemoveFailedDetail) HasRemoveRequestReferenceId() bool`

HasRemoveRequestReferenceId returns a boolean if a field has been set.

### SetRemoveRequestReferenceIdNil

`func (o *AeAssetSellerChargeRemoveFailedDetail) SetRemoveRequestReferenceIdNil(b bool)`

 SetRemoveRequestReferenceIdNil sets the value for RemoveRequestReferenceId to be an explicit nil

### UnsetRemoveRequestReferenceId
`func (o *AeAssetSellerChargeRemoveFailedDetail) UnsetRemoveRequestReferenceId()`

UnsetRemoveRequestReferenceId ensures that no value is present for RemoveRequestReferenceId, not even an explicit nil
### GetReason

`func (o *AeAssetSellerChargeRemoveFailedDetail) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AeAssetSellerChargeRemoveFailedDetail) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AeAssetSellerChargeRemoveFailedDetail) SetReason(v string)`

SetReason sets Reason field to given value.


### GetInitiator

`func (o *AeAssetSellerChargeRemoveFailedDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAssetSellerChargeRemoveFailedDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAssetSellerChargeRemoveFailedDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAssetSellerChargeRemoveFailedDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


