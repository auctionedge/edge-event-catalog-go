# CommonAssetNegotiationBaseOffer

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

## Methods

### NewCommonAssetNegotiationBaseOffer

`func NewCommonAssetNegotiationBaseOffer(id string, source string, authorizedBy string, rep CommonAmsRepPointer, amount CommonCurrency, note string, notification CommonAssetNegotiationBaseOfferNotification, ) *CommonAssetNegotiationBaseOffer`

NewCommonAssetNegotiationBaseOffer instantiates a new CommonAssetNegotiationBaseOffer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonAssetNegotiationBaseOfferWithDefaults

`func NewCommonAssetNegotiationBaseOfferWithDefaults() *CommonAssetNegotiationBaseOffer`

NewCommonAssetNegotiationBaseOfferWithDefaults instantiates a new CommonAssetNegotiationBaseOffer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommonAssetNegotiationBaseOffer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommonAssetNegotiationBaseOffer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommonAssetNegotiationBaseOffer) SetId(v string)`

SetId sets Id field to given value.


### GetSource

`func (o *CommonAssetNegotiationBaseOffer) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CommonAssetNegotiationBaseOffer) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CommonAssetNegotiationBaseOffer) SetSource(v string)`

SetSource sets Source field to given value.


### GetAuthorizedBy

`func (o *CommonAssetNegotiationBaseOffer) GetAuthorizedBy() string`

GetAuthorizedBy returns the AuthorizedBy field if non-nil, zero value otherwise.

### GetAuthorizedByOk

`func (o *CommonAssetNegotiationBaseOffer) GetAuthorizedByOk() (*string, bool)`

GetAuthorizedByOk returns a tuple with the AuthorizedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedBy

`func (o *CommonAssetNegotiationBaseOffer) SetAuthorizedBy(v string)`

SetAuthorizedBy sets AuthorizedBy field to given value.


### GetRep

`func (o *CommonAssetNegotiationBaseOffer) GetRep() CommonAmsRepPointer`

GetRep returns the Rep field if non-nil, zero value otherwise.

### GetRepOk

`func (o *CommonAssetNegotiationBaseOffer) GetRepOk() (*CommonAmsRepPointer, bool)`

GetRepOk returns a tuple with the Rep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRep

`func (o *CommonAssetNegotiationBaseOffer) SetRep(v CommonAmsRepPointer)`

SetRep sets Rep field to given value.


### GetAmount

`func (o *CommonAssetNegotiationBaseOffer) GetAmount() CommonCurrency`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CommonAssetNegotiationBaseOffer) GetAmountOk() (*CommonCurrency, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CommonAssetNegotiationBaseOffer) SetAmount(v CommonCurrency)`

SetAmount sets Amount field to given value.


### GetNote

`func (o *CommonAssetNegotiationBaseOffer) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *CommonAssetNegotiationBaseOffer) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *CommonAssetNegotiationBaseOffer) SetNote(v string)`

SetNote sets Note field to given value.


### GetNotification

`func (o *CommonAssetNegotiationBaseOffer) GetNotification() CommonAssetNegotiationBaseOfferNotification`

GetNotification returns the Notification field if non-nil, zero value otherwise.

### GetNotificationOk

`func (o *CommonAssetNegotiationBaseOffer) GetNotificationOk() (*CommonAssetNegotiationBaseOfferNotification, bool)`

GetNotificationOk returns a tuple with the Notification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotification

`func (o *CommonAssetNegotiationBaseOffer) SetNotification(v CommonAssetNegotiationBaseOfferNotification)`

SetNotification sets Notification field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


