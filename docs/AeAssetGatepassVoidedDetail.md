# AeAssetGatepassVoidedDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**Asset** | Pointer to [**CommonAmsAssetPointer**](CommonAmsAssetPointer.md) |  | [optional] 
**Vin** | **string** | The Vehicle Identification Number(VIN) of the asset. | 
**Stock** | **string** | The stock number of the asset. | 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeAssetGatepassVoidedDetail

`func NewAeAssetGatepassVoidedDetail(auctionId string, vin string, stock string, ) *AeAssetGatepassVoidedDetail`

NewAeAssetGatepassVoidedDetail instantiates a new AeAssetGatepassVoidedDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetGatepassVoidedDetailWithDefaults

`func NewAeAssetGatepassVoidedDetailWithDefaults() *AeAssetGatepassVoidedDetail`

NewAeAssetGatepassVoidedDetailWithDefaults instantiates a new AeAssetGatepassVoidedDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeAssetGatepassVoidedDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAssetGatepassVoidedDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAssetGatepassVoidedDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetAsset

`func (o *AeAssetGatepassVoidedDetail) GetAsset() CommonAmsAssetPointer`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *AeAssetGatepassVoidedDetail) GetAssetOk() (*CommonAmsAssetPointer, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *AeAssetGatepassVoidedDetail) SetAsset(v CommonAmsAssetPointer)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *AeAssetGatepassVoidedDetail) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetVin

`func (o *AeAssetGatepassVoidedDetail) GetVin() string`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *AeAssetGatepassVoidedDetail) GetVinOk() (*string, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *AeAssetGatepassVoidedDetail) SetVin(v string)`

SetVin sets Vin field to given value.


### GetStock

`func (o *AeAssetGatepassVoidedDetail) GetStock() string`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *AeAssetGatepassVoidedDetail) GetStockOk() (*string, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *AeAssetGatepassVoidedDetail) SetStock(v string)`

SetStock sets Stock field to given value.


### GetInitiator

`func (o *AeAssetGatepassVoidedDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAssetGatepassVoidedDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAssetGatepassVoidedDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAssetGatepassVoidedDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


