# AeAssetNegotiationOfferRejectedDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**NegotiationId** | **string** | Uniquie identifier for a specific negotiation. | 
**RejectedOfferId** | **string** | unique identifier of the offer that is being rejected | 
**Asset** | [**CommonAmsAssetPointer**](CommonAmsAssetPointer.md) |  | 
**Offer** | [**AeAssetNegotiationOfferRejectedDetailOffer**](AeAssetNegotiationOfferRejectedDetailOffer.md) |  | 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeAssetNegotiationOfferRejectedDetail

`func NewAeAssetNegotiationOfferRejectedDetail(auctionId string, negotiationId string, rejectedOfferId string, asset CommonAmsAssetPointer, offer AeAssetNegotiationOfferRejectedDetailOffer, ) *AeAssetNegotiationOfferRejectedDetail`

NewAeAssetNegotiationOfferRejectedDetail instantiates a new AeAssetNegotiationOfferRejectedDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetNegotiationOfferRejectedDetailWithDefaults

`func NewAeAssetNegotiationOfferRejectedDetailWithDefaults() *AeAssetNegotiationOfferRejectedDetail`

NewAeAssetNegotiationOfferRejectedDetailWithDefaults instantiates a new AeAssetNegotiationOfferRejectedDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeAssetNegotiationOfferRejectedDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAssetNegotiationOfferRejectedDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAssetNegotiationOfferRejectedDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetNegotiationId

`func (o *AeAssetNegotiationOfferRejectedDetail) GetNegotiationId() string`

GetNegotiationId returns the NegotiationId field if non-nil, zero value otherwise.

### GetNegotiationIdOk

`func (o *AeAssetNegotiationOfferRejectedDetail) GetNegotiationIdOk() (*string, bool)`

GetNegotiationIdOk returns a tuple with the NegotiationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegotiationId

`func (o *AeAssetNegotiationOfferRejectedDetail) SetNegotiationId(v string)`

SetNegotiationId sets NegotiationId field to given value.


### GetRejectedOfferId

`func (o *AeAssetNegotiationOfferRejectedDetail) GetRejectedOfferId() string`

GetRejectedOfferId returns the RejectedOfferId field if non-nil, zero value otherwise.

### GetRejectedOfferIdOk

`func (o *AeAssetNegotiationOfferRejectedDetail) GetRejectedOfferIdOk() (*string, bool)`

GetRejectedOfferIdOk returns a tuple with the RejectedOfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedOfferId

`func (o *AeAssetNegotiationOfferRejectedDetail) SetRejectedOfferId(v string)`

SetRejectedOfferId sets RejectedOfferId field to given value.


### GetAsset

`func (o *AeAssetNegotiationOfferRejectedDetail) GetAsset() CommonAmsAssetPointer`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *AeAssetNegotiationOfferRejectedDetail) GetAssetOk() (*CommonAmsAssetPointer, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *AeAssetNegotiationOfferRejectedDetail) SetAsset(v CommonAmsAssetPointer)`

SetAsset sets Asset field to given value.


### GetOffer

`func (o *AeAssetNegotiationOfferRejectedDetail) GetOffer() AeAssetNegotiationOfferRejectedDetailOffer`

GetOffer returns the Offer field if non-nil, zero value otherwise.

### GetOfferOk

`func (o *AeAssetNegotiationOfferRejectedDetail) GetOfferOk() (*AeAssetNegotiationOfferRejectedDetailOffer, bool)`

GetOfferOk returns a tuple with the Offer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffer

`func (o *AeAssetNegotiationOfferRejectedDetail) SetOffer(v AeAssetNegotiationOfferRejectedDetailOffer)`

SetOffer sets Offer field to given value.


### GetInitiator

`func (o *AeAssetNegotiationOfferRejectedDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAssetNegotiationOfferRejectedDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAssetNegotiationOfferRejectedDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAssetNegotiationOfferRejectedDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


