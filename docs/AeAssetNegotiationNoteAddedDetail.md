# AeAssetNegotiationNoteAddedDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**NegotiationId** | **string** | the unique identifier of the negotiation | 
**Note** | [**AeAssetNegotiationNoteAddedDetailNote**](AeAssetNegotiationNoteAddedDetailNote.md) |  | 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeAssetNegotiationNoteAddedDetail

`func NewAeAssetNegotiationNoteAddedDetail(auctionId string, negotiationId string, note AeAssetNegotiationNoteAddedDetailNote, ) *AeAssetNegotiationNoteAddedDetail`

NewAeAssetNegotiationNoteAddedDetail instantiates a new AeAssetNegotiationNoteAddedDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetNegotiationNoteAddedDetailWithDefaults

`func NewAeAssetNegotiationNoteAddedDetailWithDefaults() *AeAssetNegotiationNoteAddedDetail`

NewAeAssetNegotiationNoteAddedDetailWithDefaults instantiates a new AeAssetNegotiationNoteAddedDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeAssetNegotiationNoteAddedDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAssetNegotiationNoteAddedDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAssetNegotiationNoteAddedDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetNegotiationId

`func (o *AeAssetNegotiationNoteAddedDetail) GetNegotiationId() string`

GetNegotiationId returns the NegotiationId field if non-nil, zero value otherwise.

### GetNegotiationIdOk

`func (o *AeAssetNegotiationNoteAddedDetail) GetNegotiationIdOk() (*string, bool)`

GetNegotiationIdOk returns a tuple with the NegotiationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegotiationId

`func (o *AeAssetNegotiationNoteAddedDetail) SetNegotiationId(v string)`

SetNegotiationId sets NegotiationId field to given value.


### GetNote

`func (o *AeAssetNegotiationNoteAddedDetail) GetNote() AeAssetNegotiationNoteAddedDetailNote`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *AeAssetNegotiationNoteAddedDetail) GetNoteOk() (*AeAssetNegotiationNoteAddedDetailNote, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *AeAssetNegotiationNoteAddedDetail) SetNote(v AeAssetNegotiationNoteAddedDetailNote)`

SetNote sets Note field to given value.


### GetInitiator

`func (o *AeAssetNegotiationNoteAddedDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAssetNegotiationNoteAddedDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAssetNegotiationNoteAddedDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAssetNegotiationNoteAddedDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


