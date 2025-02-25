# AeAssetSellerChargeAddRequestedV2Detail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**Asset** | [**CommonAmsAssetPointer**](CommonAmsAssetPointer.md) |  | 
**Seller** | [**CommonAmsAccountPointer**](CommonAmsAccountPointer.md) |  | 
**ReferenceId** | **string** | Identifier that can be used to identify this event in a response event. | 
**Charge** | [**CommonCharge**](CommonCharge.md) |  | 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeAssetSellerChargeAddRequestedV2Detail

`func NewAeAssetSellerChargeAddRequestedV2Detail(auctionId string, asset CommonAmsAssetPointer, seller CommonAmsAccountPointer, referenceId string, charge CommonCharge, ) *AeAssetSellerChargeAddRequestedV2Detail`

NewAeAssetSellerChargeAddRequestedV2Detail instantiates a new AeAssetSellerChargeAddRequestedV2Detail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetSellerChargeAddRequestedV2DetailWithDefaults

`func NewAeAssetSellerChargeAddRequestedV2DetailWithDefaults() *AeAssetSellerChargeAddRequestedV2Detail`

NewAeAssetSellerChargeAddRequestedV2DetailWithDefaults instantiates a new AeAssetSellerChargeAddRequestedV2Detail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeAssetSellerChargeAddRequestedV2Detail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAssetSellerChargeAddRequestedV2Detail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAssetSellerChargeAddRequestedV2Detail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetAsset

`func (o *AeAssetSellerChargeAddRequestedV2Detail) GetAsset() CommonAmsAssetPointer`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *AeAssetSellerChargeAddRequestedV2Detail) GetAssetOk() (*CommonAmsAssetPointer, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *AeAssetSellerChargeAddRequestedV2Detail) SetAsset(v CommonAmsAssetPointer)`

SetAsset sets Asset field to given value.


### GetSeller

`func (o *AeAssetSellerChargeAddRequestedV2Detail) GetSeller() CommonAmsAccountPointer`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *AeAssetSellerChargeAddRequestedV2Detail) GetSellerOk() (*CommonAmsAccountPointer, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *AeAssetSellerChargeAddRequestedV2Detail) SetSeller(v CommonAmsAccountPointer)`

SetSeller sets Seller field to given value.


### GetReferenceId

`func (o *AeAssetSellerChargeAddRequestedV2Detail) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *AeAssetSellerChargeAddRequestedV2Detail) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *AeAssetSellerChargeAddRequestedV2Detail) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetCharge

`func (o *AeAssetSellerChargeAddRequestedV2Detail) GetCharge() CommonCharge`

GetCharge returns the Charge field if non-nil, zero value otherwise.

### GetChargeOk

`func (o *AeAssetSellerChargeAddRequestedV2Detail) GetChargeOk() (*CommonCharge, bool)`

GetChargeOk returns a tuple with the Charge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharge

`func (o *AeAssetSellerChargeAddRequestedV2Detail) SetCharge(v CommonCharge)`

SetCharge sets Charge field to given value.


### GetInitiator

`func (o *AeAssetSellerChargeAddRequestedV2Detail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAssetSellerChargeAddRequestedV2Detail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAssetSellerChargeAddRequestedV2Detail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAssetSellerChargeAddRequestedV2Detail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


