# AeAssetGatepassBuyerRevokedDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**Asset** | Pointer to [**CommonAmsAssetPointer**](CommonAmsAssetPointer.md) |  | [optional] 
**AssetId** | **string** | Source&#39;s unique identifier for asset | 
**AssetStock** | **string** | The stock number of the asset. | 
**BuyerAccountKey** | **string** | The account key of an AMS account (i.e. the account or dealer number) | 
**UpdatedAt** | Pointer to **time.Time** | The updated date time of the gatepass status | [optional] 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeAssetGatepassBuyerRevokedDetail

`func NewAeAssetGatepassBuyerRevokedDetail(auctionId string, assetId string, assetStock string, buyerAccountKey string, ) *AeAssetGatepassBuyerRevokedDetail`

NewAeAssetGatepassBuyerRevokedDetail instantiates a new AeAssetGatepassBuyerRevokedDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetGatepassBuyerRevokedDetailWithDefaults

`func NewAeAssetGatepassBuyerRevokedDetailWithDefaults() *AeAssetGatepassBuyerRevokedDetail`

NewAeAssetGatepassBuyerRevokedDetailWithDefaults instantiates a new AeAssetGatepassBuyerRevokedDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeAssetGatepassBuyerRevokedDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAssetGatepassBuyerRevokedDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAssetGatepassBuyerRevokedDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetAsset

`func (o *AeAssetGatepassBuyerRevokedDetail) GetAsset() CommonAmsAssetPointer`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *AeAssetGatepassBuyerRevokedDetail) GetAssetOk() (*CommonAmsAssetPointer, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *AeAssetGatepassBuyerRevokedDetail) SetAsset(v CommonAmsAssetPointer)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *AeAssetGatepassBuyerRevokedDetail) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetAssetId

`func (o *AeAssetGatepassBuyerRevokedDetail) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *AeAssetGatepassBuyerRevokedDetail) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *AeAssetGatepassBuyerRevokedDetail) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetAssetStock

`func (o *AeAssetGatepassBuyerRevokedDetail) GetAssetStock() string`

GetAssetStock returns the AssetStock field if non-nil, zero value otherwise.

### GetAssetStockOk

`func (o *AeAssetGatepassBuyerRevokedDetail) GetAssetStockOk() (*string, bool)`

GetAssetStockOk returns a tuple with the AssetStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetStock

`func (o *AeAssetGatepassBuyerRevokedDetail) SetAssetStock(v string)`

SetAssetStock sets AssetStock field to given value.


### GetBuyerAccountKey

`func (o *AeAssetGatepassBuyerRevokedDetail) GetBuyerAccountKey() string`

GetBuyerAccountKey returns the BuyerAccountKey field if non-nil, zero value otherwise.

### GetBuyerAccountKeyOk

`func (o *AeAssetGatepassBuyerRevokedDetail) GetBuyerAccountKeyOk() (*string, bool)`

GetBuyerAccountKeyOk returns a tuple with the BuyerAccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerAccountKey

`func (o *AeAssetGatepassBuyerRevokedDetail) SetBuyerAccountKey(v string)`

SetBuyerAccountKey sets BuyerAccountKey field to given value.


### GetUpdatedAt

`func (o *AeAssetGatepassBuyerRevokedDetail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AeAssetGatepassBuyerRevokedDetail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AeAssetGatepassBuyerRevokedDetail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AeAssetGatepassBuyerRevokedDetail) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetInitiator

`func (o *AeAssetGatepassBuyerRevokedDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAssetGatepassBuyerRevokedDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAssetGatepassBuyerRevokedDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAssetGatepassBuyerRevokedDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


