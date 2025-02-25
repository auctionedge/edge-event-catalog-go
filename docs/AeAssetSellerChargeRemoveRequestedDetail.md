# AeAssetSellerChargeRemoveRequestedDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**Asset** | [**CommonAmsAssetPointer**](CommonAmsAssetPointer.md) |  | 
**Seller** | [**CommonAmsAccountPointer**](CommonAmsAccountPointer.md) |  | 
**ReferenceId** | **string** | Identifier that can be used to identify this event in a response event. | 
**ChargeId** | **string** | Unique identifier of the charge to be removed. | 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeAssetSellerChargeRemoveRequestedDetail

`func NewAeAssetSellerChargeRemoveRequestedDetail(auctionId string, asset CommonAmsAssetPointer, seller CommonAmsAccountPointer, referenceId string, chargeId string, ) *AeAssetSellerChargeRemoveRequestedDetail`

NewAeAssetSellerChargeRemoveRequestedDetail instantiates a new AeAssetSellerChargeRemoveRequestedDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetSellerChargeRemoveRequestedDetailWithDefaults

`func NewAeAssetSellerChargeRemoveRequestedDetailWithDefaults() *AeAssetSellerChargeRemoveRequestedDetail`

NewAeAssetSellerChargeRemoveRequestedDetailWithDefaults instantiates a new AeAssetSellerChargeRemoveRequestedDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeAssetSellerChargeRemoveRequestedDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAssetSellerChargeRemoveRequestedDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAssetSellerChargeRemoveRequestedDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetAsset

`func (o *AeAssetSellerChargeRemoveRequestedDetail) GetAsset() CommonAmsAssetPointer`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *AeAssetSellerChargeRemoveRequestedDetail) GetAssetOk() (*CommonAmsAssetPointer, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *AeAssetSellerChargeRemoveRequestedDetail) SetAsset(v CommonAmsAssetPointer)`

SetAsset sets Asset field to given value.


### GetSeller

`func (o *AeAssetSellerChargeRemoveRequestedDetail) GetSeller() CommonAmsAccountPointer`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *AeAssetSellerChargeRemoveRequestedDetail) GetSellerOk() (*CommonAmsAccountPointer, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *AeAssetSellerChargeRemoveRequestedDetail) SetSeller(v CommonAmsAccountPointer)`

SetSeller sets Seller field to given value.


### GetReferenceId

`func (o *AeAssetSellerChargeRemoveRequestedDetail) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *AeAssetSellerChargeRemoveRequestedDetail) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *AeAssetSellerChargeRemoveRequestedDetail) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetChargeId

`func (o *AeAssetSellerChargeRemoveRequestedDetail) GetChargeId() string`

GetChargeId returns the ChargeId field if non-nil, zero value otherwise.

### GetChargeIdOk

`func (o *AeAssetSellerChargeRemoveRequestedDetail) GetChargeIdOk() (*string, bool)`

GetChargeIdOk returns a tuple with the ChargeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeId

`func (o *AeAssetSellerChargeRemoveRequestedDetail) SetChargeId(v string)`

SetChargeId sets ChargeId field to given value.


### GetInitiator

`func (o *AeAssetSellerChargeRemoveRequestedDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAssetSellerChargeRemoveRequestedDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAssetSellerChargeRemoveRequestedDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAssetSellerChargeRemoveRequestedDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


