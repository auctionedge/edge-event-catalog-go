# AeAssetDealChargeRemovedDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**Asset** | Pointer to [**CommonAmsAssetPointer**](CommonAmsAssetPointer.md) |  | [optional] 
**DealId** | Pointer to **string** | Unique id for a deal | [optional] 
**AssetId** | **string** | Source&#39;s unique identifier for asset | 
**ChargeId** | **string** | Charge ID. | 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeAssetDealChargeRemovedDetail

`func NewAeAssetDealChargeRemovedDetail(auctionId string, assetId string, chargeId string, ) *AeAssetDealChargeRemovedDetail`

NewAeAssetDealChargeRemovedDetail instantiates a new AeAssetDealChargeRemovedDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetDealChargeRemovedDetailWithDefaults

`func NewAeAssetDealChargeRemovedDetailWithDefaults() *AeAssetDealChargeRemovedDetail`

NewAeAssetDealChargeRemovedDetailWithDefaults instantiates a new AeAssetDealChargeRemovedDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeAssetDealChargeRemovedDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAssetDealChargeRemovedDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAssetDealChargeRemovedDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetAsset

`func (o *AeAssetDealChargeRemovedDetail) GetAsset() CommonAmsAssetPointer`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *AeAssetDealChargeRemovedDetail) GetAssetOk() (*CommonAmsAssetPointer, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *AeAssetDealChargeRemovedDetail) SetAsset(v CommonAmsAssetPointer)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *AeAssetDealChargeRemovedDetail) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetDealId

`func (o *AeAssetDealChargeRemovedDetail) GetDealId() string`

GetDealId returns the DealId field if non-nil, zero value otherwise.

### GetDealIdOk

`func (o *AeAssetDealChargeRemovedDetail) GetDealIdOk() (*string, bool)`

GetDealIdOk returns a tuple with the DealId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealId

`func (o *AeAssetDealChargeRemovedDetail) SetDealId(v string)`

SetDealId sets DealId field to given value.

### HasDealId

`func (o *AeAssetDealChargeRemovedDetail) HasDealId() bool`

HasDealId returns a boolean if a field has been set.

### GetAssetId

`func (o *AeAssetDealChargeRemovedDetail) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *AeAssetDealChargeRemovedDetail) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *AeAssetDealChargeRemovedDetail) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetChargeId

`func (o *AeAssetDealChargeRemovedDetail) GetChargeId() string`

GetChargeId returns the ChargeId field if non-nil, zero value otherwise.

### GetChargeIdOk

`func (o *AeAssetDealChargeRemovedDetail) GetChargeIdOk() (*string, bool)`

GetChargeIdOk returns a tuple with the ChargeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeId

`func (o *AeAssetDealChargeRemovedDetail) SetChargeId(v string)`

SetChargeId sets ChargeId field to given value.


### GetInitiator

`func (o *AeAssetDealChargeRemovedDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAssetDealChargeRemovedDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAssetDealChargeRemovedDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAssetDealChargeRemovedDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


