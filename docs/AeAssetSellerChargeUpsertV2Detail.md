# AeAssetSellerChargeUpsertV2Detail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**Asset** | [**CommonAmsAssetPointer**](CommonAmsAssetPointer.md) |  | 
**Seller** | [**CommonAmsAccountPointer**](CommonAmsAccountPointer.md) |  | 
**AddRequestReferenceId** | Pointer to **NullableString** | Unique identifier referencing the add request event that created the charge | [optional] 
**Charge** | [**CommonCharge**](CommonCharge.md) |  | 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeAssetSellerChargeUpsertV2Detail

`func NewAeAssetSellerChargeUpsertV2Detail(auctionId string, asset CommonAmsAssetPointer, seller CommonAmsAccountPointer, charge CommonCharge, ) *AeAssetSellerChargeUpsertV2Detail`

NewAeAssetSellerChargeUpsertV2Detail instantiates a new AeAssetSellerChargeUpsertV2Detail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetSellerChargeUpsertV2DetailWithDefaults

`func NewAeAssetSellerChargeUpsertV2DetailWithDefaults() *AeAssetSellerChargeUpsertV2Detail`

NewAeAssetSellerChargeUpsertV2DetailWithDefaults instantiates a new AeAssetSellerChargeUpsertV2Detail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeAssetSellerChargeUpsertV2Detail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAssetSellerChargeUpsertV2Detail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAssetSellerChargeUpsertV2Detail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetAsset

`func (o *AeAssetSellerChargeUpsertV2Detail) GetAsset() CommonAmsAssetPointer`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *AeAssetSellerChargeUpsertV2Detail) GetAssetOk() (*CommonAmsAssetPointer, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *AeAssetSellerChargeUpsertV2Detail) SetAsset(v CommonAmsAssetPointer)`

SetAsset sets Asset field to given value.


### GetSeller

`func (o *AeAssetSellerChargeUpsertV2Detail) GetSeller() CommonAmsAccountPointer`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *AeAssetSellerChargeUpsertV2Detail) GetSellerOk() (*CommonAmsAccountPointer, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *AeAssetSellerChargeUpsertV2Detail) SetSeller(v CommonAmsAccountPointer)`

SetSeller sets Seller field to given value.


### GetAddRequestReferenceId

`func (o *AeAssetSellerChargeUpsertV2Detail) GetAddRequestReferenceId() string`

GetAddRequestReferenceId returns the AddRequestReferenceId field if non-nil, zero value otherwise.

### GetAddRequestReferenceIdOk

`func (o *AeAssetSellerChargeUpsertV2Detail) GetAddRequestReferenceIdOk() (*string, bool)`

GetAddRequestReferenceIdOk returns a tuple with the AddRequestReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddRequestReferenceId

`func (o *AeAssetSellerChargeUpsertV2Detail) SetAddRequestReferenceId(v string)`

SetAddRequestReferenceId sets AddRequestReferenceId field to given value.

### HasAddRequestReferenceId

`func (o *AeAssetSellerChargeUpsertV2Detail) HasAddRequestReferenceId() bool`

HasAddRequestReferenceId returns a boolean if a field has been set.

### SetAddRequestReferenceIdNil

`func (o *AeAssetSellerChargeUpsertV2Detail) SetAddRequestReferenceIdNil(b bool)`

 SetAddRequestReferenceIdNil sets the value for AddRequestReferenceId to be an explicit nil

### UnsetAddRequestReferenceId
`func (o *AeAssetSellerChargeUpsertV2Detail) UnsetAddRequestReferenceId()`

UnsetAddRequestReferenceId ensures that no value is present for AddRequestReferenceId, not even an explicit nil
### GetCharge

`func (o *AeAssetSellerChargeUpsertV2Detail) GetCharge() CommonCharge`

GetCharge returns the Charge field if non-nil, zero value otherwise.

### GetChargeOk

`func (o *AeAssetSellerChargeUpsertV2Detail) GetChargeOk() (*CommonCharge, bool)`

GetChargeOk returns a tuple with the Charge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharge

`func (o *AeAssetSellerChargeUpsertV2Detail) SetCharge(v CommonCharge)`

SetCharge sets Charge field to given value.


### GetInitiator

`func (o *AeAssetSellerChargeUpsertV2Detail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAssetSellerChargeUpsertV2Detail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAssetSellerChargeUpsertV2Detail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAssetSellerChargeUpsertV2Detail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


