# AeAssetNegotiationOfferCounteredDetailOffer

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
**CounteredBy** | [**CommonAssetNegotiationOfferBy**](CommonAssetNegotiationOfferBy.md) |  | 

## Methods

### NewAeAssetNegotiationOfferCounteredDetailOffer

`func NewAeAssetNegotiationOfferCounteredDetailOffer(id string, source string, authorizedBy string, rep CommonAmsRepPointer, amount CommonCurrency, note string, notification CommonAssetNegotiationBaseOfferNotification, counteredBy CommonAssetNegotiationOfferBy, ) *AeAssetNegotiationOfferCounteredDetailOffer`

NewAeAssetNegotiationOfferCounteredDetailOffer instantiates a new AeAssetNegotiationOfferCounteredDetailOffer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetNegotiationOfferCounteredDetailOfferWithDefaults

`func NewAeAssetNegotiationOfferCounteredDetailOfferWithDefaults() *AeAssetNegotiationOfferCounteredDetailOffer`

NewAeAssetNegotiationOfferCounteredDetailOfferWithDefaults instantiates a new AeAssetNegotiationOfferCounteredDetailOffer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AeAssetNegotiationOfferCounteredDetailOffer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AeAssetNegotiationOfferCounteredDetailOffer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AeAssetNegotiationOfferCounteredDetailOffer) SetId(v string)`

SetId sets Id field to given value.


### GetSource

`func (o *AeAssetNegotiationOfferCounteredDetailOffer) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AeAssetNegotiationOfferCounteredDetailOffer) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AeAssetNegotiationOfferCounteredDetailOffer) SetSource(v string)`

SetSource sets Source field to given value.


### GetAuthorizedBy

`func (o *AeAssetNegotiationOfferCounteredDetailOffer) GetAuthorizedBy() string`

GetAuthorizedBy returns the AuthorizedBy field if non-nil, zero value otherwise.

### GetAuthorizedByOk

`func (o *AeAssetNegotiationOfferCounteredDetailOffer) GetAuthorizedByOk() (*string, bool)`

GetAuthorizedByOk returns a tuple with the AuthorizedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedBy

`func (o *AeAssetNegotiationOfferCounteredDetailOffer) SetAuthorizedBy(v string)`

SetAuthorizedBy sets AuthorizedBy field to given value.


### GetRep

`func (o *AeAssetNegotiationOfferCounteredDetailOffer) GetRep() CommonAmsRepPointer`

GetRep returns the Rep field if non-nil, zero value otherwise.

### GetRepOk

`func (o *AeAssetNegotiationOfferCounteredDetailOffer) GetRepOk() (*CommonAmsRepPointer, bool)`

GetRepOk returns a tuple with the Rep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRep

`func (o *AeAssetNegotiationOfferCounteredDetailOffer) SetRep(v CommonAmsRepPointer)`

SetRep sets Rep field to given value.


### GetAmount

`func (o *AeAssetNegotiationOfferCounteredDetailOffer) GetAmount() CommonCurrency`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AeAssetNegotiationOfferCounteredDetailOffer) GetAmountOk() (*CommonCurrency, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AeAssetNegotiationOfferCounteredDetailOffer) SetAmount(v CommonCurrency)`

SetAmount sets Amount field to given value.


### GetNote

`func (o *AeAssetNegotiationOfferCounteredDetailOffer) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *AeAssetNegotiationOfferCounteredDetailOffer) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *AeAssetNegotiationOfferCounteredDetailOffer) SetNote(v string)`

SetNote sets Note field to given value.


### GetNotification

`func (o *AeAssetNegotiationOfferCounteredDetailOffer) GetNotification() CommonAssetNegotiationBaseOfferNotification`

GetNotification returns the Notification field if non-nil, zero value otherwise.

### GetNotificationOk

`func (o *AeAssetNegotiationOfferCounteredDetailOffer) GetNotificationOk() (*CommonAssetNegotiationBaseOfferNotification, bool)`

GetNotificationOk returns a tuple with the Notification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotification

`func (o *AeAssetNegotiationOfferCounteredDetailOffer) SetNotification(v CommonAssetNegotiationBaseOfferNotification)`

SetNotification sets Notification field to given value.


### GetCounteredBy

`func (o *AeAssetNegotiationOfferCounteredDetailOffer) GetCounteredBy() CommonAssetNegotiationOfferBy`

GetCounteredBy returns the CounteredBy field if non-nil, zero value otherwise.

### GetCounteredByOk

`func (o *AeAssetNegotiationOfferCounteredDetailOffer) GetCounteredByOk() (*CommonAssetNegotiationOfferBy, bool)`

GetCounteredByOk returns a tuple with the CounteredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounteredBy

`func (o *AeAssetNegotiationOfferCounteredDetailOffer) SetCounteredBy(v CommonAssetNegotiationOfferBy)`

SetCounteredBy sets CounteredBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


