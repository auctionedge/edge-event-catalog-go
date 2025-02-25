# AeAssetLabelGeneratedDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**RequestId** | **string** | Unique identifier for the request | 
**Asset** | Pointer to [**CommonAmsAssetPointer**](CommonAmsAssetPointer.md) |  | [optional] 
**AssetId** | **string** | Source&#39;s unique identifier for asset | 
**Stock** | **string** | The stock number of the asset. | 
**Vin** | **string** | The Vehicle Identification Number(VIN) of the asset. | 
**Quantity** | **float32** | Number of copies of the label to generate for asset | 
**LabelType** | [**LabelType**](LabelType.md) |  | 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeAssetLabelGeneratedDetail

`func NewAeAssetLabelGeneratedDetail(auctionId string, requestId string, assetId string, stock string, vin string, quantity float32, labelType LabelType, ) *AeAssetLabelGeneratedDetail`

NewAeAssetLabelGeneratedDetail instantiates a new AeAssetLabelGeneratedDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetLabelGeneratedDetailWithDefaults

`func NewAeAssetLabelGeneratedDetailWithDefaults() *AeAssetLabelGeneratedDetail`

NewAeAssetLabelGeneratedDetailWithDefaults instantiates a new AeAssetLabelGeneratedDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeAssetLabelGeneratedDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAssetLabelGeneratedDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAssetLabelGeneratedDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetRequestId

`func (o *AeAssetLabelGeneratedDetail) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *AeAssetLabelGeneratedDetail) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *AeAssetLabelGeneratedDetail) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetAsset

`func (o *AeAssetLabelGeneratedDetail) GetAsset() CommonAmsAssetPointer`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *AeAssetLabelGeneratedDetail) GetAssetOk() (*CommonAmsAssetPointer, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *AeAssetLabelGeneratedDetail) SetAsset(v CommonAmsAssetPointer)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *AeAssetLabelGeneratedDetail) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetAssetId

`func (o *AeAssetLabelGeneratedDetail) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *AeAssetLabelGeneratedDetail) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *AeAssetLabelGeneratedDetail) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetStock

`func (o *AeAssetLabelGeneratedDetail) GetStock() string`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *AeAssetLabelGeneratedDetail) GetStockOk() (*string, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *AeAssetLabelGeneratedDetail) SetStock(v string)`

SetStock sets Stock field to given value.


### GetVin

`func (o *AeAssetLabelGeneratedDetail) GetVin() string`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *AeAssetLabelGeneratedDetail) GetVinOk() (*string, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *AeAssetLabelGeneratedDetail) SetVin(v string)`

SetVin sets Vin field to given value.


### GetQuantity

`func (o *AeAssetLabelGeneratedDetail) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *AeAssetLabelGeneratedDetail) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *AeAssetLabelGeneratedDetail) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetLabelType

`func (o *AeAssetLabelGeneratedDetail) GetLabelType() LabelType`

GetLabelType returns the LabelType field if non-nil, zero value otherwise.

### GetLabelTypeOk

`func (o *AeAssetLabelGeneratedDetail) GetLabelTypeOk() (*LabelType, bool)`

GetLabelTypeOk returns a tuple with the LabelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelType

`func (o *AeAssetLabelGeneratedDetail) SetLabelType(v LabelType)`

SetLabelType sets LabelType field to given value.


### GetInitiator

`func (o *AeAssetLabelGeneratedDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAssetLabelGeneratedDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAssetLabelGeneratedDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAssetLabelGeneratedDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


