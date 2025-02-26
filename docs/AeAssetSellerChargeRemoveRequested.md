# AeAssetSellerChargeRemoveRequested

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | [**AeAssetSellerChargeRemoveRequestedDetail**](AeAssetSellerChargeRemoveRequestedDetail.md) |  | 
**Version** | Pointer to **string** | Optional version | [optional] 
**Account** | Pointer to **string** | The publisher AWS account number | [optional] 
**DetailType** | **string** | Identifies, in combination with the source field, the fields and values that appear in the detail field. | 
**Id** | **string** | A Version 4 UUID that&#39;s generated for every event. You can use id to trace events as they move through rules to targets. | 
**Region** | Pointer to **string** | Identifies the AWS Region where the event originated | [optional] 
**Resources** | Pointer to **[]string** | Identifiers for any resources involved in the event | [optional] 
**Source** | **string** | Identifies the service that generated the event. May not begin with &#39;aws.&#39; | 
**Time** | Pointer to **time.Time** | The event timestamp, which can be specified by the service originating the event. May be before the actual publication of the event. ISO-8601 format | [optional] 
**AuctionId** | Pointer to **string** | UA id of the auction. Not all events will need the identification of an auction. | [optional] 
**ExpiresAt** | Pointer to **time.Time** | Timestamp after which the consumer should no longer process the message. ISO-8601 format | [optional] 
**Deprecated** | Pointer to **bool** | True if this version has been deprecated. | [optional] [default to false]
**IdempotencyKey** | Pointer to **string** | MD5 content hash on the entire event content (Library generated) | [optional] 

## Methods

### NewAeAssetSellerChargeRemoveRequested

`func NewAeAssetSellerChargeRemoveRequested(detail AeAssetSellerChargeRemoveRequestedDetail, detailType string, id string, source string, ) *AeAssetSellerChargeRemoveRequested`

NewAeAssetSellerChargeRemoveRequested instantiates a new AeAssetSellerChargeRemoveRequested object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetSellerChargeRemoveRequestedWithDefaults

`func NewAeAssetSellerChargeRemoveRequestedWithDefaults() *AeAssetSellerChargeRemoveRequested`

NewAeAssetSellerChargeRemoveRequestedWithDefaults instantiates a new AeAssetSellerChargeRemoveRequested object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *AeAssetSellerChargeRemoveRequested) GetDetail() AeAssetSellerChargeRemoveRequestedDetail`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *AeAssetSellerChargeRemoveRequested) GetDetailOk() (*AeAssetSellerChargeRemoveRequestedDetail, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *AeAssetSellerChargeRemoveRequested) SetDetail(v AeAssetSellerChargeRemoveRequestedDetail)`

SetDetail sets Detail field to given value.


### GetVersion

`func (o *AeAssetSellerChargeRemoveRequested) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AeAssetSellerChargeRemoveRequested) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AeAssetSellerChargeRemoveRequested) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AeAssetSellerChargeRemoveRequested) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAccount

`func (o *AeAssetSellerChargeRemoveRequested) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AeAssetSellerChargeRemoveRequested) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AeAssetSellerChargeRemoveRequested) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AeAssetSellerChargeRemoveRequested) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetDetailType

`func (o *AeAssetSellerChargeRemoveRequested) GetDetailType() string`

GetDetailType returns the DetailType field if non-nil, zero value otherwise.

### GetDetailTypeOk

`func (o *AeAssetSellerChargeRemoveRequested) GetDetailTypeOk() (*string, bool)`

GetDetailTypeOk returns a tuple with the DetailType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailType

`func (o *AeAssetSellerChargeRemoveRequested) SetDetailType(v string)`

SetDetailType sets DetailType field to given value.


### GetId

`func (o *AeAssetSellerChargeRemoveRequested) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AeAssetSellerChargeRemoveRequested) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AeAssetSellerChargeRemoveRequested) SetId(v string)`

SetId sets Id field to given value.


### GetRegion

`func (o *AeAssetSellerChargeRemoveRequested) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AeAssetSellerChargeRemoveRequested) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AeAssetSellerChargeRemoveRequested) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AeAssetSellerChargeRemoveRequested) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetResources

`func (o *AeAssetSellerChargeRemoveRequested) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *AeAssetSellerChargeRemoveRequested) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *AeAssetSellerChargeRemoveRequested) SetResources(v []string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *AeAssetSellerChargeRemoveRequested) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetSource

`func (o *AeAssetSellerChargeRemoveRequested) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AeAssetSellerChargeRemoveRequested) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AeAssetSellerChargeRemoveRequested) SetSource(v string)`

SetSource sets Source field to given value.


### GetTime

`func (o *AeAssetSellerChargeRemoveRequested) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *AeAssetSellerChargeRemoveRequested) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *AeAssetSellerChargeRemoveRequested) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *AeAssetSellerChargeRemoveRequested) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetAuctionId

`func (o *AeAssetSellerChargeRemoveRequested) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAssetSellerChargeRemoveRequested) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAssetSellerChargeRemoveRequested) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.

### HasAuctionId

`func (o *AeAssetSellerChargeRemoveRequested) HasAuctionId() bool`

HasAuctionId returns a boolean if a field has been set.

### GetExpiresAt

`func (o *AeAssetSellerChargeRemoveRequested) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *AeAssetSellerChargeRemoveRequested) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *AeAssetSellerChargeRemoveRequested) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *AeAssetSellerChargeRemoveRequested) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetDeprecated

`func (o *AeAssetSellerChargeRemoveRequested) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *AeAssetSellerChargeRemoveRequested) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *AeAssetSellerChargeRemoveRequested) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *AeAssetSellerChargeRemoveRequested) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *AeAssetSellerChargeRemoveRequested) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *AeAssetSellerChargeRemoveRequested) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *AeAssetSellerChargeRemoveRequested) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *AeAssetSellerChargeRemoveRequested) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


