# AeDocumentAssetRemoved

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | [**AeDocumentAssetRemovedDetail**](AeDocumentAssetRemovedDetail.md) |  | 
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

### NewAeDocumentAssetRemoved

`func NewAeDocumentAssetRemoved(detail AeDocumentAssetRemovedDetail, detailType string, id string, source string, ) *AeDocumentAssetRemoved`

NewAeDocumentAssetRemoved instantiates a new AeDocumentAssetRemoved object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeDocumentAssetRemovedWithDefaults

`func NewAeDocumentAssetRemovedWithDefaults() *AeDocumentAssetRemoved`

NewAeDocumentAssetRemovedWithDefaults instantiates a new AeDocumentAssetRemoved object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *AeDocumentAssetRemoved) GetDetail() AeDocumentAssetRemovedDetail`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *AeDocumentAssetRemoved) GetDetailOk() (*AeDocumentAssetRemovedDetail, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *AeDocumentAssetRemoved) SetDetail(v AeDocumentAssetRemovedDetail)`

SetDetail sets Detail field to given value.


### GetVersion

`func (o *AeDocumentAssetRemoved) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AeDocumentAssetRemoved) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AeDocumentAssetRemoved) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AeDocumentAssetRemoved) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAccount

`func (o *AeDocumentAssetRemoved) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AeDocumentAssetRemoved) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AeDocumentAssetRemoved) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AeDocumentAssetRemoved) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetDetailType

`func (o *AeDocumentAssetRemoved) GetDetailType() string`

GetDetailType returns the DetailType field if non-nil, zero value otherwise.

### GetDetailTypeOk

`func (o *AeDocumentAssetRemoved) GetDetailTypeOk() (*string, bool)`

GetDetailTypeOk returns a tuple with the DetailType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailType

`func (o *AeDocumentAssetRemoved) SetDetailType(v string)`

SetDetailType sets DetailType field to given value.


### GetId

`func (o *AeDocumentAssetRemoved) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AeDocumentAssetRemoved) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AeDocumentAssetRemoved) SetId(v string)`

SetId sets Id field to given value.


### GetRegion

`func (o *AeDocumentAssetRemoved) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AeDocumentAssetRemoved) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AeDocumentAssetRemoved) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AeDocumentAssetRemoved) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetResources

`func (o *AeDocumentAssetRemoved) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *AeDocumentAssetRemoved) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *AeDocumentAssetRemoved) SetResources(v []string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *AeDocumentAssetRemoved) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetSource

`func (o *AeDocumentAssetRemoved) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AeDocumentAssetRemoved) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AeDocumentAssetRemoved) SetSource(v string)`

SetSource sets Source field to given value.


### GetTime

`func (o *AeDocumentAssetRemoved) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *AeDocumentAssetRemoved) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *AeDocumentAssetRemoved) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *AeDocumentAssetRemoved) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetAuctionId

`func (o *AeDocumentAssetRemoved) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeDocumentAssetRemoved) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeDocumentAssetRemoved) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.

### HasAuctionId

`func (o *AeDocumentAssetRemoved) HasAuctionId() bool`

HasAuctionId returns a boolean if a field has been set.

### GetExpiresAt

`func (o *AeDocumentAssetRemoved) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *AeDocumentAssetRemoved) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *AeDocumentAssetRemoved) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *AeDocumentAssetRemoved) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetDeprecated

`func (o *AeDocumentAssetRemoved) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *AeDocumentAssetRemoved) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *AeDocumentAssetRemoved) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *AeDocumentAssetRemoved) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *AeDocumentAssetRemoved) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *AeDocumentAssetRemoved) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *AeDocumentAssetRemoved) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *AeDocumentAssetRemoved) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


