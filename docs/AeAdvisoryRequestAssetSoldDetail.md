# AeAdvisoryRequestAssetSoldDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**EventTypes** | **[]string** | Array of advisory messages that are being requested to be produced | 
**DateRange** | [**DateRange**](DateRange.md) |  | 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeAdvisoryRequestAssetSoldDetail

`func NewAeAdvisoryRequestAssetSoldDetail(auctionId string, eventTypes []string, dateRange DateRange, ) *AeAdvisoryRequestAssetSoldDetail`

NewAeAdvisoryRequestAssetSoldDetail instantiates a new AeAdvisoryRequestAssetSoldDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAdvisoryRequestAssetSoldDetailWithDefaults

`func NewAeAdvisoryRequestAssetSoldDetailWithDefaults() *AeAdvisoryRequestAssetSoldDetail`

NewAeAdvisoryRequestAssetSoldDetailWithDefaults instantiates a new AeAdvisoryRequestAssetSoldDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeAdvisoryRequestAssetSoldDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAdvisoryRequestAssetSoldDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAdvisoryRequestAssetSoldDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetEventTypes

`func (o *AeAdvisoryRequestAssetSoldDetail) GetEventTypes() []string`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *AeAdvisoryRequestAssetSoldDetail) GetEventTypesOk() (*[]string, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *AeAdvisoryRequestAssetSoldDetail) SetEventTypes(v []string)`

SetEventTypes sets EventTypes field to given value.


### GetDateRange

`func (o *AeAdvisoryRequestAssetSoldDetail) GetDateRange() DateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *AeAdvisoryRequestAssetSoldDetail) GetDateRangeOk() (*DateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *AeAdvisoryRequestAssetSoldDetail) SetDateRange(v DateRange)`

SetDateRange sets DateRange field to given value.


### GetInitiator

`func (o *AeAdvisoryRequestAssetSoldDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAdvisoryRequestAssetSoldDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAdvisoryRequestAssetSoldDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAdvisoryRequestAssetSoldDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


