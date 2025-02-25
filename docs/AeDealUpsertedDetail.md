# AeDealUpsertedDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique id for a deal | 
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**Asset** | [**CommonAmsAssetPointer**](CommonAmsAssetPointer.md) |  | 
**Seller** | **string** | Source&#39;s unique identifier for account | 
**SellerAgent** | Pointer to **string** | The representative id an AMS account | [optional] 
**SaleListing** | **string** | Uniquie identifier for a sale listing. | 
**SoldAt** | **time.Time** | The time in ISO-8601 formatted at which the asset was sold | 
**Buyer** | **string** | Source&#39;s unique identifier for account | 
**BuyerAgent** | Pointer to **string** | The representative id an AMS account | [optional] 
**Amount** | [**CommonCurrency**](CommonCurrency.md) |  | 
**SoldMethod** | Pointer to **string** | The type of sale the deal was made in. | [optional] 
**UpdatedAt** | **time.Time** | The time in ISO-8601 format in UTC timezone at which the record was generated | 
**Initiator** | [**CommonInitiator**](CommonInitiator.md) |  | 

## Methods

### NewAeDealUpsertedDetail

`func NewAeDealUpsertedDetail(id string, auctionId string, asset CommonAmsAssetPointer, seller string, saleListing string, soldAt time.Time, buyer string, amount CommonCurrency, updatedAt time.Time, initiator CommonInitiator, ) *AeDealUpsertedDetail`

NewAeDealUpsertedDetail instantiates a new AeDealUpsertedDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeDealUpsertedDetailWithDefaults

`func NewAeDealUpsertedDetailWithDefaults() *AeDealUpsertedDetail`

NewAeDealUpsertedDetailWithDefaults instantiates a new AeDealUpsertedDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AeDealUpsertedDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AeDealUpsertedDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AeDealUpsertedDetail) SetId(v string)`

SetId sets Id field to given value.


### GetAuctionId

`func (o *AeDealUpsertedDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeDealUpsertedDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeDealUpsertedDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetAsset

`func (o *AeDealUpsertedDetail) GetAsset() CommonAmsAssetPointer`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *AeDealUpsertedDetail) GetAssetOk() (*CommonAmsAssetPointer, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *AeDealUpsertedDetail) SetAsset(v CommonAmsAssetPointer)`

SetAsset sets Asset field to given value.


### GetSeller

`func (o *AeDealUpsertedDetail) GetSeller() string`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *AeDealUpsertedDetail) GetSellerOk() (*string, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *AeDealUpsertedDetail) SetSeller(v string)`

SetSeller sets Seller field to given value.


### GetSellerAgent

`func (o *AeDealUpsertedDetail) GetSellerAgent() string`

GetSellerAgent returns the SellerAgent field if non-nil, zero value otherwise.

### GetSellerAgentOk

`func (o *AeDealUpsertedDetail) GetSellerAgentOk() (*string, bool)`

GetSellerAgentOk returns a tuple with the SellerAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerAgent

`func (o *AeDealUpsertedDetail) SetSellerAgent(v string)`

SetSellerAgent sets SellerAgent field to given value.

### HasSellerAgent

`func (o *AeDealUpsertedDetail) HasSellerAgent() bool`

HasSellerAgent returns a boolean if a field has been set.

### GetSaleListing

`func (o *AeDealUpsertedDetail) GetSaleListing() string`

GetSaleListing returns the SaleListing field if non-nil, zero value otherwise.

### GetSaleListingOk

`func (o *AeDealUpsertedDetail) GetSaleListingOk() (*string, bool)`

GetSaleListingOk returns a tuple with the SaleListing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleListing

`func (o *AeDealUpsertedDetail) SetSaleListing(v string)`

SetSaleListing sets SaleListing field to given value.


### GetSoldAt

`func (o *AeDealUpsertedDetail) GetSoldAt() time.Time`

GetSoldAt returns the SoldAt field if non-nil, zero value otherwise.

### GetSoldAtOk

`func (o *AeDealUpsertedDetail) GetSoldAtOk() (*time.Time, bool)`

GetSoldAtOk returns a tuple with the SoldAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoldAt

`func (o *AeDealUpsertedDetail) SetSoldAt(v time.Time)`

SetSoldAt sets SoldAt field to given value.


### GetBuyer

`func (o *AeDealUpsertedDetail) GetBuyer() string`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *AeDealUpsertedDetail) GetBuyerOk() (*string, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *AeDealUpsertedDetail) SetBuyer(v string)`

SetBuyer sets Buyer field to given value.


### GetBuyerAgent

`func (o *AeDealUpsertedDetail) GetBuyerAgent() string`

GetBuyerAgent returns the BuyerAgent field if non-nil, zero value otherwise.

### GetBuyerAgentOk

`func (o *AeDealUpsertedDetail) GetBuyerAgentOk() (*string, bool)`

GetBuyerAgentOk returns a tuple with the BuyerAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerAgent

`func (o *AeDealUpsertedDetail) SetBuyerAgent(v string)`

SetBuyerAgent sets BuyerAgent field to given value.

### HasBuyerAgent

`func (o *AeDealUpsertedDetail) HasBuyerAgent() bool`

HasBuyerAgent returns a boolean if a field has been set.

### GetAmount

`func (o *AeDealUpsertedDetail) GetAmount() CommonCurrency`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AeDealUpsertedDetail) GetAmountOk() (*CommonCurrency, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AeDealUpsertedDetail) SetAmount(v CommonCurrency)`

SetAmount sets Amount field to given value.


### GetSoldMethod

`func (o *AeDealUpsertedDetail) GetSoldMethod() string`

GetSoldMethod returns the SoldMethod field if non-nil, zero value otherwise.

### GetSoldMethodOk

`func (o *AeDealUpsertedDetail) GetSoldMethodOk() (*string, bool)`

GetSoldMethodOk returns a tuple with the SoldMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoldMethod

`func (o *AeDealUpsertedDetail) SetSoldMethod(v string)`

SetSoldMethod sets SoldMethod field to given value.

### HasSoldMethod

`func (o *AeDealUpsertedDetail) HasSoldMethod() bool`

HasSoldMethod returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AeDealUpsertedDetail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AeDealUpsertedDetail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AeDealUpsertedDetail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetInitiator

`func (o *AeDealUpsertedDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeDealUpsertedDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeDealUpsertedDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


