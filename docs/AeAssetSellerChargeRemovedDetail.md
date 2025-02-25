# AeAssetSellerChargeRemovedDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**ChargeId** | **string** | Unique identifier of the charge being removed. | 
**Seller** | [**CommonAmsAccountPointer**](CommonAmsAccountPointer.md) |  | 
**Asset** | [**CommonAmsAssetPointer**](CommonAmsAssetPointer.md) |  | 
**RemoveRequestReferenceId** | Pointer to **NullableString** | Unique identifier referencing the remove request event that removed the charge | [optional] 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeAssetSellerChargeRemovedDetail

`func NewAeAssetSellerChargeRemovedDetail(auctionId string, chargeId string, seller CommonAmsAccountPointer, asset CommonAmsAssetPointer, ) *AeAssetSellerChargeRemovedDetail`

NewAeAssetSellerChargeRemovedDetail instantiates a new AeAssetSellerChargeRemovedDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetSellerChargeRemovedDetailWithDefaults

`func NewAeAssetSellerChargeRemovedDetailWithDefaults() *AeAssetSellerChargeRemovedDetail`

NewAeAssetSellerChargeRemovedDetailWithDefaults instantiates a new AeAssetSellerChargeRemovedDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeAssetSellerChargeRemovedDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAssetSellerChargeRemovedDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAssetSellerChargeRemovedDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetChargeId

`func (o *AeAssetSellerChargeRemovedDetail) GetChargeId() string`

GetChargeId returns the ChargeId field if non-nil, zero value otherwise.

### GetChargeIdOk

`func (o *AeAssetSellerChargeRemovedDetail) GetChargeIdOk() (*string, bool)`

GetChargeIdOk returns a tuple with the ChargeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeId

`func (o *AeAssetSellerChargeRemovedDetail) SetChargeId(v string)`

SetChargeId sets ChargeId field to given value.


### GetSeller

`func (o *AeAssetSellerChargeRemovedDetail) GetSeller() CommonAmsAccountPointer`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *AeAssetSellerChargeRemovedDetail) GetSellerOk() (*CommonAmsAccountPointer, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *AeAssetSellerChargeRemovedDetail) SetSeller(v CommonAmsAccountPointer)`

SetSeller sets Seller field to given value.


### GetAsset

`func (o *AeAssetSellerChargeRemovedDetail) GetAsset() CommonAmsAssetPointer`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *AeAssetSellerChargeRemovedDetail) GetAssetOk() (*CommonAmsAssetPointer, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *AeAssetSellerChargeRemovedDetail) SetAsset(v CommonAmsAssetPointer)`

SetAsset sets Asset field to given value.


### GetRemoveRequestReferenceId

`func (o *AeAssetSellerChargeRemovedDetail) GetRemoveRequestReferenceId() string`

GetRemoveRequestReferenceId returns the RemoveRequestReferenceId field if non-nil, zero value otherwise.

### GetRemoveRequestReferenceIdOk

`func (o *AeAssetSellerChargeRemovedDetail) GetRemoveRequestReferenceIdOk() (*string, bool)`

GetRemoveRequestReferenceIdOk returns a tuple with the RemoveRequestReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveRequestReferenceId

`func (o *AeAssetSellerChargeRemovedDetail) SetRemoveRequestReferenceId(v string)`

SetRemoveRequestReferenceId sets RemoveRequestReferenceId field to given value.

### HasRemoveRequestReferenceId

`func (o *AeAssetSellerChargeRemovedDetail) HasRemoveRequestReferenceId() bool`

HasRemoveRequestReferenceId returns a boolean if a field has been set.

### SetRemoveRequestReferenceIdNil

`func (o *AeAssetSellerChargeRemovedDetail) SetRemoveRequestReferenceIdNil(b bool)`

 SetRemoveRequestReferenceIdNil sets the value for RemoveRequestReferenceId to be an explicit nil

### UnsetRemoveRequestReferenceId
`func (o *AeAssetSellerChargeRemovedDetail) UnsetRemoveRequestReferenceId()`

UnsetRemoveRequestReferenceId ensures that no value is present for RemoveRequestReferenceId, not even an explicit nil
### GetInitiator

`func (o *AeAssetSellerChargeRemovedDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAssetSellerChargeRemovedDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAssetSellerChargeRemovedDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAssetSellerChargeRemovedDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


