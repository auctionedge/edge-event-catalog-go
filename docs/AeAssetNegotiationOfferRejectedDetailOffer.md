# AeAssetNegotiationOfferRejectedDetailOffer

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
**RejectedBy** | [**CommonAssetNegotiationOfferBy**](CommonAssetNegotiationOfferBy.md) |  | 

## Methods

### NewAeAssetNegotiationOfferRejectedDetailOffer

`func NewAeAssetNegotiationOfferRejectedDetailOffer(id string, source string, authorizedBy string, rep CommonAmsRepPointer, amount CommonCurrency, note string, notification CommonAssetNegotiationBaseOfferNotification, rejectedBy CommonAssetNegotiationOfferBy, ) *AeAssetNegotiationOfferRejectedDetailOffer`

NewAeAssetNegotiationOfferRejectedDetailOffer instantiates a new AeAssetNegotiationOfferRejectedDetailOffer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetNegotiationOfferRejectedDetailOfferWithDefaults

`func NewAeAssetNegotiationOfferRejectedDetailOfferWithDefaults() *AeAssetNegotiationOfferRejectedDetailOffer`

NewAeAssetNegotiationOfferRejectedDetailOfferWithDefaults instantiates a new AeAssetNegotiationOfferRejectedDetailOffer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AeAssetNegotiationOfferRejectedDetailOffer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AeAssetNegotiationOfferRejectedDetailOffer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AeAssetNegotiationOfferRejectedDetailOffer) SetId(v string)`

SetId sets Id field to given value.


### GetSource

`func (o *AeAssetNegotiationOfferRejectedDetailOffer) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AeAssetNegotiationOfferRejectedDetailOffer) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AeAssetNegotiationOfferRejectedDetailOffer) SetSource(v string)`

SetSource sets Source field to given value.


### GetAuthorizedBy

`func (o *AeAssetNegotiationOfferRejectedDetailOffer) GetAuthorizedBy() string`

GetAuthorizedBy returns the AuthorizedBy field if non-nil, zero value otherwise.

### GetAuthorizedByOk

`func (o *AeAssetNegotiationOfferRejectedDetailOffer) GetAuthorizedByOk() (*string, bool)`

GetAuthorizedByOk returns a tuple with the AuthorizedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedBy

`func (o *AeAssetNegotiationOfferRejectedDetailOffer) SetAuthorizedBy(v string)`

SetAuthorizedBy sets AuthorizedBy field to given value.


### GetRep

`func (o *AeAssetNegotiationOfferRejectedDetailOffer) GetRep() CommonAmsRepPointer`

GetRep returns the Rep field if non-nil, zero value otherwise.

### GetRepOk

`func (o *AeAssetNegotiationOfferRejectedDetailOffer) GetRepOk() (*CommonAmsRepPointer, bool)`

GetRepOk returns a tuple with the Rep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRep

`func (o *AeAssetNegotiationOfferRejectedDetailOffer) SetRep(v CommonAmsRepPointer)`

SetRep sets Rep field to given value.


### GetAmount

`func (o *AeAssetNegotiationOfferRejectedDetailOffer) GetAmount() CommonCurrency`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AeAssetNegotiationOfferRejectedDetailOffer) GetAmountOk() (*CommonCurrency, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AeAssetNegotiationOfferRejectedDetailOffer) SetAmount(v CommonCurrency)`

SetAmount sets Amount field to given value.


### GetNote

`func (o *AeAssetNegotiationOfferRejectedDetailOffer) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *AeAssetNegotiationOfferRejectedDetailOffer) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *AeAssetNegotiationOfferRejectedDetailOffer) SetNote(v string)`

SetNote sets Note field to given value.


### GetNotification

`func (o *AeAssetNegotiationOfferRejectedDetailOffer) GetNotification() CommonAssetNegotiationBaseOfferNotification`

GetNotification returns the Notification field if non-nil, zero value otherwise.

### GetNotificationOk

`func (o *AeAssetNegotiationOfferRejectedDetailOffer) GetNotificationOk() (*CommonAssetNegotiationBaseOfferNotification, bool)`

GetNotificationOk returns a tuple with the Notification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotification

`func (o *AeAssetNegotiationOfferRejectedDetailOffer) SetNotification(v CommonAssetNegotiationBaseOfferNotification)`

SetNotification sets Notification field to given value.


### GetRejectedBy

`func (o *AeAssetNegotiationOfferRejectedDetailOffer) GetRejectedBy() CommonAssetNegotiationOfferBy`

GetRejectedBy returns the RejectedBy field if non-nil, zero value otherwise.

### GetRejectedByOk

`func (o *AeAssetNegotiationOfferRejectedDetailOffer) GetRejectedByOk() (*CommonAssetNegotiationOfferBy, bool)`

GetRejectedByOk returns a tuple with the RejectedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedBy

`func (o *AeAssetNegotiationOfferRejectedDetailOffer) SetRejectedBy(v CommonAssetNegotiationOfferBy)`

SetRejectedBy sets RejectedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


