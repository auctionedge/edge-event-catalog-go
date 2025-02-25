# AeAssetNegotiationUpsertedDetailConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Seller** | [**CommonAmsAccountPointer**](CommonAmsAccountPointer.md) |  | 
**Buyer** | [**CommonAmsAccountPointer**](CommonAmsAccountPointer.md) |  | 
**Lights** | **string** | Light that was on the vehicle when negotiations started | 
**Announcements** | **[]string** |  | 

## Methods

### NewAeAssetNegotiationUpsertedDetailConditions

`func NewAeAssetNegotiationUpsertedDetailConditions(seller CommonAmsAccountPointer, buyer CommonAmsAccountPointer, lights string, announcements []string, ) *AeAssetNegotiationUpsertedDetailConditions`

NewAeAssetNegotiationUpsertedDetailConditions instantiates a new AeAssetNegotiationUpsertedDetailConditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetNegotiationUpsertedDetailConditionsWithDefaults

`func NewAeAssetNegotiationUpsertedDetailConditionsWithDefaults() *AeAssetNegotiationUpsertedDetailConditions`

NewAeAssetNegotiationUpsertedDetailConditionsWithDefaults instantiates a new AeAssetNegotiationUpsertedDetailConditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeller

`func (o *AeAssetNegotiationUpsertedDetailConditions) GetSeller() CommonAmsAccountPointer`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *AeAssetNegotiationUpsertedDetailConditions) GetSellerOk() (*CommonAmsAccountPointer, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *AeAssetNegotiationUpsertedDetailConditions) SetSeller(v CommonAmsAccountPointer)`

SetSeller sets Seller field to given value.


### GetBuyer

`func (o *AeAssetNegotiationUpsertedDetailConditions) GetBuyer() CommonAmsAccountPointer`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *AeAssetNegotiationUpsertedDetailConditions) GetBuyerOk() (*CommonAmsAccountPointer, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *AeAssetNegotiationUpsertedDetailConditions) SetBuyer(v CommonAmsAccountPointer)`

SetBuyer sets Buyer field to given value.


### GetLights

`func (o *AeAssetNegotiationUpsertedDetailConditions) GetLights() string`

GetLights returns the Lights field if non-nil, zero value otherwise.

### GetLightsOk

`func (o *AeAssetNegotiationUpsertedDetailConditions) GetLightsOk() (*string, bool)`

GetLightsOk returns a tuple with the Lights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLights

`func (o *AeAssetNegotiationUpsertedDetailConditions) SetLights(v string)`

SetLights sets Lights field to given value.


### GetAnnouncements

`func (o *AeAssetNegotiationUpsertedDetailConditions) GetAnnouncements() []string`

GetAnnouncements returns the Announcements field if non-nil, zero value otherwise.

### GetAnnouncementsOk

`func (o *AeAssetNegotiationUpsertedDetailConditions) GetAnnouncementsOk() (*[]string, bool)`

GetAnnouncementsOk returns a tuple with the Announcements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncements

`func (o *AeAssetNegotiationUpsertedDetailConditions) SetAnnouncements(v []string)`

SetAnnouncements sets Announcements field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


