# AeAdvisoryAssetAms

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | [**AeAdvisoryAssetAmsDetail**](AeAdvisoryAssetAmsDetail.md) |  | 
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

### NewAeAdvisoryAssetAms

`func NewAeAdvisoryAssetAms(detail AeAdvisoryAssetAmsDetail, detailType string, id string, source string, ) *AeAdvisoryAssetAms`

NewAeAdvisoryAssetAms instantiates a new AeAdvisoryAssetAms object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAdvisoryAssetAmsWithDefaults

`func NewAeAdvisoryAssetAmsWithDefaults() *AeAdvisoryAssetAms`

NewAeAdvisoryAssetAmsWithDefaults instantiates a new AeAdvisoryAssetAms object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *AeAdvisoryAssetAms) GetDetail() AeAdvisoryAssetAmsDetail`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *AeAdvisoryAssetAms) GetDetailOk() (*AeAdvisoryAssetAmsDetail, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *AeAdvisoryAssetAms) SetDetail(v AeAdvisoryAssetAmsDetail)`

SetDetail sets Detail field to given value.


### GetVersion

`func (o *AeAdvisoryAssetAms) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AeAdvisoryAssetAms) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AeAdvisoryAssetAms) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AeAdvisoryAssetAms) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAccount

`func (o *AeAdvisoryAssetAms) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AeAdvisoryAssetAms) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AeAdvisoryAssetAms) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AeAdvisoryAssetAms) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetDetailType

`func (o *AeAdvisoryAssetAms) GetDetailType() string`

GetDetailType returns the DetailType field if non-nil, zero value otherwise.

### GetDetailTypeOk

`func (o *AeAdvisoryAssetAms) GetDetailTypeOk() (*string, bool)`

GetDetailTypeOk returns a tuple with the DetailType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailType

`func (o *AeAdvisoryAssetAms) SetDetailType(v string)`

SetDetailType sets DetailType field to given value.


### GetId

`func (o *AeAdvisoryAssetAms) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AeAdvisoryAssetAms) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AeAdvisoryAssetAms) SetId(v string)`

SetId sets Id field to given value.


### GetRegion

`func (o *AeAdvisoryAssetAms) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AeAdvisoryAssetAms) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AeAdvisoryAssetAms) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AeAdvisoryAssetAms) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetResources

`func (o *AeAdvisoryAssetAms) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *AeAdvisoryAssetAms) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *AeAdvisoryAssetAms) SetResources(v []string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *AeAdvisoryAssetAms) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetSource

`func (o *AeAdvisoryAssetAms) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AeAdvisoryAssetAms) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AeAdvisoryAssetAms) SetSource(v string)`

SetSource sets Source field to given value.


### GetTime

`func (o *AeAdvisoryAssetAms) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *AeAdvisoryAssetAms) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *AeAdvisoryAssetAms) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *AeAdvisoryAssetAms) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetAuctionId

`func (o *AeAdvisoryAssetAms) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAdvisoryAssetAms) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAdvisoryAssetAms) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.

### HasAuctionId

`func (o *AeAdvisoryAssetAms) HasAuctionId() bool`

HasAuctionId returns a boolean if a field has been set.

### GetExpiresAt

`func (o *AeAdvisoryAssetAms) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *AeAdvisoryAssetAms) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *AeAdvisoryAssetAms) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *AeAdvisoryAssetAms) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetDeprecated

`func (o *AeAdvisoryAssetAms) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *AeAdvisoryAssetAms) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *AeAdvisoryAssetAms) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *AeAdvisoryAssetAms) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *AeAdvisoryAssetAms) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *AeAdvisoryAssetAms) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *AeAdvisoryAssetAms) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *AeAdvisoryAssetAms) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


