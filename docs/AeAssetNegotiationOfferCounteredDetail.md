# AeAssetNegotiationOfferCounteredDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**NegotiationId** | **string** | Uniquie identifier for a specific negotiation. | 
**CounteredOfferId** | **string** | unique identifier of the offer that is being countered | 
**Asset** | [**CommonAmsAssetPointer**](CommonAmsAssetPointer.md) |  | 
**Offer** | [**AeAssetNegotiationOfferCounteredDetailOffer**](AeAssetNegotiationOfferCounteredDetailOffer.md) |  | 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeAssetNegotiationOfferCounteredDetail

`func NewAeAssetNegotiationOfferCounteredDetail(auctionId string, negotiationId string, counteredOfferId string, asset CommonAmsAssetPointer, offer AeAssetNegotiationOfferCounteredDetailOffer, ) *AeAssetNegotiationOfferCounteredDetail`

NewAeAssetNegotiationOfferCounteredDetail instantiates a new AeAssetNegotiationOfferCounteredDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetNegotiationOfferCounteredDetailWithDefaults

`func NewAeAssetNegotiationOfferCounteredDetailWithDefaults() *AeAssetNegotiationOfferCounteredDetail`

NewAeAssetNegotiationOfferCounteredDetailWithDefaults instantiates a new AeAssetNegotiationOfferCounteredDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeAssetNegotiationOfferCounteredDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAssetNegotiationOfferCounteredDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAssetNegotiationOfferCounteredDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetNegotiationId

`func (o *AeAssetNegotiationOfferCounteredDetail) GetNegotiationId() string`

GetNegotiationId returns the NegotiationId field if non-nil, zero value otherwise.

### GetNegotiationIdOk

`func (o *AeAssetNegotiationOfferCounteredDetail) GetNegotiationIdOk() (*string, bool)`

GetNegotiationIdOk returns a tuple with the NegotiationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegotiationId

`func (o *AeAssetNegotiationOfferCounteredDetail) SetNegotiationId(v string)`

SetNegotiationId sets NegotiationId field to given value.


### GetCounteredOfferId

`func (o *AeAssetNegotiationOfferCounteredDetail) GetCounteredOfferId() string`

GetCounteredOfferId returns the CounteredOfferId field if non-nil, zero value otherwise.

### GetCounteredOfferIdOk

`func (o *AeAssetNegotiationOfferCounteredDetail) GetCounteredOfferIdOk() (*string, bool)`

GetCounteredOfferIdOk returns a tuple with the CounteredOfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounteredOfferId

`func (o *AeAssetNegotiationOfferCounteredDetail) SetCounteredOfferId(v string)`

SetCounteredOfferId sets CounteredOfferId field to given value.


### GetAsset

`func (o *AeAssetNegotiationOfferCounteredDetail) GetAsset() CommonAmsAssetPointer`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *AeAssetNegotiationOfferCounteredDetail) GetAssetOk() (*CommonAmsAssetPointer, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *AeAssetNegotiationOfferCounteredDetail) SetAsset(v CommonAmsAssetPointer)`

SetAsset sets Asset field to given value.


### GetOffer

`func (o *AeAssetNegotiationOfferCounteredDetail) GetOffer() AeAssetNegotiationOfferCounteredDetailOffer`

GetOffer returns the Offer field if non-nil, zero value otherwise.

### GetOfferOk

`func (o *AeAssetNegotiationOfferCounteredDetail) GetOfferOk() (*AeAssetNegotiationOfferCounteredDetailOffer, bool)`

GetOfferOk returns a tuple with the Offer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffer

`func (o *AeAssetNegotiationOfferCounteredDetail) SetOffer(v AeAssetNegotiationOfferCounteredDetailOffer)`

SetOffer sets Offer field to given value.


### GetInitiator

`func (o *AeAssetNegotiationOfferCounteredDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAssetNegotiationOfferCounteredDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAssetNegotiationOfferCounteredDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAssetNegotiationOfferCounteredDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


