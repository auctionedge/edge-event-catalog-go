# AeAssetSellerChargeUpsertDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**Asset** | Pointer to [**CommonAmsAssetPointer**](CommonAmsAssetPointer.md) |  | [optional] 
**AssetId** | **string** | Source&#39;s unique identifier for asset | 
**Vin** | **string** | The Vehicle Identification Number(VIN) of the asset. | 
**Stock** | **string** | The stock number of the asset. | 
**Seller** | [**CommonAmsAccountPointer**](CommonAmsAccountPointer.md) |  | 
**AddRequestReferenceId** | Pointer to **NullableString** | Unique identifier referencing the add request event that created the charge | [optional] 
**Charge** | [**AeAssetSellerChargeUpsertDetailCharge**](AeAssetSellerChargeUpsertDetailCharge.md) |  | 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeAssetSellerChargeUpsertDetail

`func NewAeAssetSellerChargeUpsertDetail(auctionId string, assetId string, vin string, stock string, seller CommonAmsAccountPointer, charge AeAssetSellerChargeUpsertDetailCharge, ) *AeAssetSellerChargeUpsertDetail`

NewAeAssetSellerChargeUpsertDetail instantiates a new AeAssetSellerChargeUpsertDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetSellerChargeUpsertDetailWithDefaults

`func NewAeAssetSellerChargeUpsertDetailWithDefaults() *AeAssetSellerChargeUpsertDetail`

NewAeAssetSellerChargeUpsertDetailWithDefaults instantiates a new AeAssetSellerChargeUpsertDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeAssetSellerChargeUpsertDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAssetSellerChargeUpsertDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAssetSellerChargeUpsertDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetAsset

`func (o *AeAssetSellerChargeUpsertDetail) GetAsset() CommonAmsAssetPointer`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *AeAssetSellerChargeUpsertDetail) GetAssetOk() (*CommonAmsAssetPointer, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *AeAssetSellerChargeUpsertDetail) SetAsset(v CommonAmsAssetPointer)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *AeAssetSellerChargeUpsertDetail) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetAssetId

`func (o *AeAssetSellerChargeUpsertDetail) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *AeAssetSellerChargeUpsertDetail) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *AeAssetSellerChargeUpsertDetail) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetVin

`func (o *AeAssetSellerChargeUpsertDetail) GetVin() string`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *AeAssetSellerChargeUpsertDetail) GetVinOk() (*string, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *AeAssetSellerChargeUpsertDetail) SetVin(v string)`

SetVin sets Vin field to given value.


### GetStock

`func (o *AeAssetSellerChargeUpsertDetail) GetStock() string`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *AeAssetSellerChargeUpsertDetail) GetStockOk() (*string, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *AeAssetSellerChargeUpsertDetail) SetStock(v string)`

SetStock sets Stock field to given value.


### GetSeller

`func (o *AeAssetSellerChargeUpsertDetail) GetSeller() CommonAmsAccountPointer`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *AeAssetSellerChargeUpsertDetail) GetSellerOk() (*CommonAmsAccountPointer, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *AeAssetSellerChargeUpsertDetail) SetSeller(v CommonAmsAccountPointer)`

SetSeller sets Seller field to given value.


### GetAddRequestReferenceId

`func (o *AeAssetSellerChargeUpsertDetail) GetAddRequestReferenceId() string`

GetAddRequestReferenceId returns the AddRequestReferenceId field if non-nil, zero value otherwise.

### GetAddRequestReferenceIdOk

`func (o *AeAssetSellerChargeUpsertDetail) GetAddRequestReferenceIdOk() (*string, bool)`

GetAddRequestReferenceIdOk returns a tuple with the AddRequestReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddRequestReferenceId

`func (o *AeAssetSellerChargeUpsertDetail) SetAddRequestReferenceId(v string)`

SetAddRequestReferenceId sets AddRequestReferenceId field to given value.

### HasAddRequestReferenceId

`func (o *AeAssetSellerChargeUpsertDetail) HasAddRequestReferenceId() bool`

HasAddRequestReferenceId returns a boolean if a field has been set.

### SetAddRequestReferenceIdNil

`func (o *AeAssetSellerChargeUpsertDetail) SetAddRequestReferenceIdNil(b bool)`

 SetAddRequestReferenceIdNil sets the value for AddRequestReferenceId to be an explicit nil

### UnsetAddRequestReferenceId
`func (o *AeAssetSellerChargeUpsertDetail) UnsetAddRequestReferenceId()`

UnsetAddRequestReferenceId ensures that no value is present for AddRequestReferenceId, not even an explicit nil
### GetCharge

`func (o *AeAssetSellerChargeUpsertDetail) GetCharge() AeAssetSellerChargeUpsertDetailCharge`

GetCharge returns the Charge field if non-nil, zero value otherwise.

### GetChargeOk

`func (o *AeAssetSellerChargeUpsertDetail) GetChargeOk() (*AeAssetSellerChargeUpsertDetailCharge, bool)`

GetChargeOk returns a tuple with the Charge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharge

`func (o *AeAssetSellerChargeUpsertDetail) SetCharge(v AeAssetSellerChargeUpsertDetailCharge)`

SetCharge sets Charge field to given value.


### GetInitiator

`func (o *AeAssetSellerChargeUpsertDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAssetSellerChargeUpsertDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAssetSellerChargeUpsertDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAssetSellerChargeUpsertDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


