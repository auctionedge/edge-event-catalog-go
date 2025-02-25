# AeAssetNegotiationOfferAcceptedDetailOffer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | unique identifier of the offer | 
**Source** | **string** | the source system that created the offer | 
**AuthorizedBy** | **string** | user that entered the offer | 
**Rep** | [**CommonAmsRepPointer**](CommonAmsRepPointer.md) |  | 
**Amount** | [**CommonCurrency**](CommonCurrency.md) |  | 
**Note** | **string** | a noted added at the time of the offfer | 
**Notification** | [**CommonAssetNegotiationBaseOfferNotification**](CommonAssetNegotiationBaseOfferNotification.md) |  | 
**AcceptedBy** | [**CommonAssetNegotiationOfferBy**](CommonAssetNegotiationOfferBy.md) |  | 

## Methods

### NewAeAssetNegotiationOfferAcceptedDetailOffer

`func NewAeAssetNegotiationOfferAcceptedDetailOffer(id string, source string, authorizedBy string, rep CommonAmsRepPointer, amount CommonCurrency, note string, notification CommonAssetNegotiationBaseOfferNotification, acceptedBy CommonAssetNegotiationOfferBy, ) *AeAssetNegotiationOfferAcceptedDetailOffer`

NewAeAssetNegotiationOfferAcceptedDetailOffer instantiates a new AeAssetNegotiationOfferAcceptedDetailOffer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetNegotiationOfferAcceptedDetailOfferWithDefaults

`func NewAeAssetNegotiationOfferAcceptedDetailOfferWithDefaults() *AeAssetNegotiationOfferAcceptedDetailOffer`

NewAeAssetNegotiationOfferAcceptedDetailOfferWithDefaults instantiates a new AeAssetNegotiationOfferAcceptedDetailOffer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AeAssetNegotiationOfferAcceptedDetailOffer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AeAssetNegotiationOfferAcceptedDetailOffer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AeAssetNegotiationOfferAcceptedDetailOffer) SetId(v string)`

SetId sets Id field to given value.


### GetSource

`func (o *AeAssetNegotiationOfferAcceptedDetailOffer) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AeAssetNegotiationOfferAcceptedDetailOffer) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AeAssetNegotiationOfferAcceptedDetailOffer) SetSource(v string)`

SetSource sets Source field to given value.


### GetAuthorizedBy

`func (o *AeAssetNegotiationOfferAcceptedDetailOffer) GetAuthorizedBy() string`

GetAuthorizedBy returns the AuthorizedBy field if non-nil, zero value otherwise.

### GetAuthorizedByOk

`func (o *AeAssetNegotiationOfferAcceptedDetailOffer) GetAuthorizedByOk() (*string, bool)`

GetAuthorizedByOk returns a tuple with the AuthorizedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedBy

`func (o *AeAssetNegotiationOfferAcceptedDetailOffer) SetAuthorizedBy(v string)`

SetAuthorizedBy sets AuthorizedBy field to given value.


### GetRep

`func (o *AeAssetNegotiationOfferAcceptedDetailOffer) GetRep() CommonAmsRepPointer`

GetRep returns the Rep field if non-nil, zero value otherwise.

### GetRepOk

`func (o *AeAssetNegotiationOfferAcceptedDetailOffer) GetRepOk() (*CommonAmsRepPointer, bool)`

GetRepOk returns a tuple with the Rep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRep

`func (o *AeAssetNegotiationOfferAcceptedDetailOffer) SetRep(v CommonAmsRepPointer)`

SetRep sets Rep field to given value.


### GetAmount

`func (o *AeAssetNegotiationOfferAcceptedDetailOffer) GetAmount() CommonCurrency`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AeAssetNegotiationOfferAcceptedDetailOffer) GetAmountOk() (*CommonCurrency, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AeAssetNegotiationOfferAcceptedDetailOffer) SetAmount(v CommonCurrency)`

SetAmount sets Amount field to given value.


### GetNote

`func (o *AeAssetNegotiationOfferAcceptedDetailOffer) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *AeAssetNegotiationOfferAcceptedDetailOffer) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *AeAssetNegotiationOfferAcceptedDetailOffer) SetNote(v string)`

SetNote sets Note field to given value.


### GetNotification

`func (o *AeAssetNegotiationOfferAcceptedDetailOffer) GetNotification() CommonAssetNegotiationBaseOfferNotification`

GetNotification returns the Notification field if non-nil, zero value otherwise.

### GetNotificationOk

`func (o *AeAssetNegotiationOfferAcceptedDetailOffer) GetNotificationOk() (*CommonAssetNegotiationBaseOfferNotification, bool)`

GetNotificationOk returns a tuple with the Notification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotification

`func (o *AeAssetNegotiationOfferAcceptedDetailOffer) SetNotification(v CommonAssetNegotiationBaseOfferNotification)`

SetNotification sets Notification field to given value.


### GetAcceptedBy

`func (o *AeAssetNegotiationOfferAcceptedDetailOffer) GetAcceptedBy() CommonAssetNegotiationOfferBy`

GetAcceptedBy returns the AcceptedBy field if non-nil, zero value otherwise.

### GetAcceptedByOk

`func (o *AeAssetNegotiationOfferAcceptedDetailOffer) GetAcceptedByOk() (*CommonAssetNegotiationOfferBy, bool)`

GetAcceptedByOk returns a tuple with the AcceptedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedBy

`func (o *AeAssetNegotiationOfferAcceptedDetailOffer) SetAcceptedBy(v CommonAssetNegotiationOfferBy)`

SetAcceptedBy sets AcceptedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


