# AeAssetSaleListingUpsertedAmsDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**Asset** | [**CommonAmsAssetPointer**](CommonAmsAssetPointer.md) |  | 
**SaleListing** | [**CommonAssetSaleListing**](CommonAssetSaleListing.md) |  | 
**UpdatedAt** | **time.Time** | The updated date time of the gatepass status | 
**Initiator** | [**CommonInitiator**](CommonInitiator.md) |  | 

## Methods

### NewAeAssetSaleListingUpsertedAmsDetail

`func NewAeAssetSaleListingUpsertedAmsDetail(auctionId string, asset CommonAmsAssetPointer, saleListing CommonAssetSaleListing, updatedAt time.Time, initiator CommonInitiator, ) *AeAssetSaleListingUpsertedAmsDetail`

NewAeAssetSaleListingUpsertedAmsDetail instantiates a new AeAssetSaleListingUpsertedAmsDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetSaleListingUpsertedAmsDetailWithDefaults

`func NewAeAssetSaleListingUpsertedAmsDetailWithDefaults() *AeAssetSaleListingUpsertedAmsDetail`

NewAeAssetSaleListingUpsertedAmsDetailWithDefaults instantiates a new AeAssetSaleListingUpsertedAmsDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeAssetSaleListingUpsertedAmsDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAssetSaleListingUpsertedAmsDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAssetSaleListingUpsertedAmsDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetAsset

`func (o *AeAssetSaleListingUpsertedAmsDetail) GetAsset() CommonAmsAssetPointer`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *AeAssetSaleListingUpsertedAmsDetail) GetAssetOk() (*CommonAmsAssetPointer, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *AeAssetSaleListingUpsertedAmsDetail) SetAsset(v CommonAmsAssetPointer)`

SetAsset sets Asset field to given value.


### GetSaleListing

`func (o *AeAssetSaleListingUpsertedAmsDetail) GetSaleListing() CommonAssetSaleListing`

GetSaleListing returns the SaleListing field if non-nil, zero value otherwise.

### GetSaleListingOk

`func (o *AeAssetSaleListingUpsertedAmsDetail) GetSaleListingOk() (*CommonAssetSaleListing, bool)`

GetSaleListingOk returns a tuple with the SaleListing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleListing

`func (o *AeAssetSaleListingUpsertedAmsDetail) SetSaleListing(v CommonAssetSaleListing)`

SetSaleListing sets SaleListing field to given value.


### GetUpdatedAt

`func (o *AeAssetSaleListingUpsertedAmsDetail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AeAssetSaleListingUpsertedAmsDetail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AeAssetSaleListingUpsertedAmsDetail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetInitiator

`func (o *AeAssetSaleListingUpsertedAmsDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAssetSaleListingUpsertedAmsDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAssetSaleListingUpsertedAmsDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


