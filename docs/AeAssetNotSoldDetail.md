# AeAssetNotSoldDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**Asset** | [**CommonAmsAssetPointer**](CommonAmsAssetPointer.md) |  | 
**SaleListingId** | **string** | Uniquie identifier for a sale listing. | 
**UpdatedAt** | **time.Time** | The updated date time of the gatepass status | 
**Initiator** | [**CommonInitiator**](CommonInitiator.md) |  | 

## Methods

### NewAeAssetNotSoldDetail

`func NewAeAssetNotSoldDetail(auctionId string, asset CommonAmsAssetPointer, saleListingId string, updatedAt time.Time, initiator CommonInitiator, ) *AeAssetNotSoldDetail`

NewAeAssetNotSoldDetail instantiates a new AeAssetNotSoldDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetNotSoldDetailWithDefaults

`func NewAeAssetNotSoldDetailWithDefaults() *AeAssetNotSoldDetail`

NewAeAssetNotSoldDetailWithDefaults instantiates a new AeAssetNotSoldDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeAssetNotSoldDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAssetNotSoldDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAssetNotSoldDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetAsset

`func (o *AeAssetNotSoldDetail) GetAsset() CommonAmsAssetPointer`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *AeAssetNotSoldDetail) GetAssetOk() (*CommonAmsAssetPointer, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *AeAssetNotSoldDetail) SetAsset(v CommonAmsAssetPointer)`

SetAsset sets Asset field to given value.


### GetSaleListingId

`func (o *AeAssetNotSoldDetail) GetSaleListingId() string`

GetSaleListingId returns the SaleListingId field if non-nil, zero value otherwise.

### GetSaleListingIdOk

`func (o *AeAssetNotSoldDetail) GetSaleListingIdOk() (*string, bool)`

GetSaleListingIdOk returns a tuple with the SaleListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleListingId

`func (o *AeAssetNotSoldDetail) SetSaleListingId(v string)`

SetSaleListingId sets SaleListingId field to given value.


### GetUpdatedAt

`func (o *AeAssetNotSoldDetail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AeAssetNotSoldDetail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AeAssetNotSoldDetail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetInitiator

`func (o *AeAssetNotSoldDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAssetNotSoldDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAssetNotSoldDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


