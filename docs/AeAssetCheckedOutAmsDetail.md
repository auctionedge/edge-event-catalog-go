# AeAssetCheckedOutAmsDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**Asset** | [**CommonAmsAssetPointer**](CommonAmsAssetPointer.md) |  | 
**UpdatedAt** | **time.Time** | The updated date time of the gatepass status | 
**CheckedOutBy** | **string** | The auction personnel that checked out the vehicle | 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeAssetCheckedOutAmsDetail

`func NewAeAssetCheckedOutAmsDetail(auctionId string, asset CommonAmsAssetPointer, updatedAt time.Time, checkedOutBy string, ) *AeAssetCheckedOutAmsDetail`

NewAeAssetCheckedOutAmsDetail instantiates a new AeAssetCheckedOutAmsDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetCheckedOutAmsDetailWithDefaults

`func NewAeAssetCheckedOutAmsDetailWithDefaults() *AeAssetCheckedOutAmsDetail`

NewAeAssetCheckedOutAmsDetailWithDefaults instantiates a new AeAssetCheckedOutAmsDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeAssetCheckedOutAmsDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAssetCheckedOutAmsDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAssetCheckedOutAmsDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetAsset

`func (o *AeAssetCheckedOutAmsDetail) GetAsset() CommonAmsAssetPointer`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *AeAssetCheckedOutAmsDetail) GetAssetOk() (*CommonAmsAssetPointer, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *AeAssetCheckedOutAmsDetail) SetAsset(v CommonAmsAssetPointer)`

SetAsset sets Asset field to given value.


### GetUpdatedAt

`func (o *AeAssetCheckedOutAmsDetail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AeAssetCheckedOutAmsDetail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AeAssetCheckedOutAmsDetail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCheckedOutBy

`func (o *AeAssetCheckedOutAmsDetail) GetCheckedOutBy() string`

GetCheckedOutBy returns the CheckedOutBy field if non-nil, zero value otherwise.

### GetCheckedOutByOk

`func (o *AeAssetCheckedOutAmsDetail) GetCheckedOutByOk() (*string, bool)`

GetCheckedOutByOk returns a tuple with the CheckedOutBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckedOutBy

`func (o *AeAssetCheckedOutAmsDetail) SetCheckedOutBy(v string)`

SetCheckedOutBy sets CheckedOutBy field to given value.


### GetInitiator

`func (o *AeAssetCheckedOutAmsDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAssetCheckedOutAmsDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAssetCheckedOutAmsDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAssetCheckedOutAmsDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


