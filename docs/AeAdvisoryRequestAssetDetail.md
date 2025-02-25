# AeAdvisoryRequestAssetDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**EventTypes** | **[]string** | Array of advisory messages that are being requested to be produced | 
**Ids** | **[]string** |  | 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeAdvisoryRequestAssetDetail

`func NewAeAdvisoryRequestAssetDetail(auctionId string, eventTypes []string, ids []string, ) *AeAdvisoryRequestAssetDetail`

NewAeAdvisoryRequestAssetDetail instantiates a new AeAdvisoryRequestAssetDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAdvisoryRequestAssetDetailWithDefaults

`func NewAeAdvisoryRequestAssetDetailWithDefaults() *AeAdvisoryRequestAssetDetail`

NewAeAdvisoryRequestAssetDetailWithDefaults instantiates a new AeAdvisoryRequestAssetDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeAdvisoryRequestAssetDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAdvisoryRequestAssetDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAdvisoryRequestAssetDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetEventTypes

`func (o *AeAdvisoryRequestAssetDetail) GetEventTypes() []string`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *AeAdvisoryRequestAssetDetail) GetEventTypesOk() (*[]string, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *AeAdvisoryRequestAssetDetail) SetEventTypes(v []string)`

SetEventTypes sets EventTypes field to given value.


### GetIds

`func (o *AeAdvisoryRequestAssetDetail) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *AeAdvisoryRequestAssetDetail) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *AeAdvisoryRequestAssetDetail) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetInitiator

`func (o *AeAdvisoryRequestAssetDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAdvisoryRequestAssetDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAdvisoryRequestAssetDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAdvisoryRequestAssetDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


