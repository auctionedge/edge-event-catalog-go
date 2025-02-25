# AeAdvisoryRequestAccountAmsDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**EventTypes** | **[]string** | Array of advisory messages that are being requested to be produced | 
**Ids** | **[]string** |  | 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeAdvisoryRequestAccountAmsDetail

`func NewAeAdvisoryRequestAccountAmsDetail(auctionId string, eventTypes []string, ids []string, ) *AeAdvisoryRequestAccountAmsDetail`

NewAeAdvisoryRequestAccountAmsDetail instantiates a new AeAdvisoryRequestAccountAmsDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAdvisoryRequestAccountAmsDetailWithDefaults

`func NewAeAdvisoryRequestAccountAmsDetailWithDefaults() *AeAdvisoryRequestAccountAmsDetail`

NewAeAdvisoryRequestAccountAmsDetailWithDefaults instantiates a new AeAdvisoryRequestAccountAmsDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeAdvisoryRequestAccountAmsDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAdvisoryRequestAccountAmsDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAdvisoryRequestAccountAmsDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetEventTypes

`func (o *AeAdvisoryRequestAccountAmsDetail) GetEventTypes() []string`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *AeAdvisoryRequestAccountAmsDetail) GetEventTypesOk() (*[]string, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *AeAdvisoryRequestAccountAmsDetail) SetEventTypes(v []string)`

SetEventTypes sets EventTypes field to given value.


### GetIds

`func (o *AeAdvisoryRequestAccountAmsDetail) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *AeAdvisoryRequestAccountAmsDetail) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *AeAdvisoryRequestAccountAmsDetail) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetInitiator

`func (o *AeAdvisoryRequestAccountAmsDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAdvisoryRequestAccountAmsDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAdvisoryRequestAccountAmsDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAdvisoryRequestAccountAmsDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


