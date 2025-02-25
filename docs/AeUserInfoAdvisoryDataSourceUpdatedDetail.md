# AeUserInfoAdvisoryDataSourceUpdatedDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** |  | 
**Bucket** | [**AeUserInfoAdvisoryDataSourceUpdatedDetailBucket**](AeUserInfoAdvisoryDataSourceUpdatedDetailBucket.md) |  | 
**Object** | [**AeUserInfoAdvisoryDataSourceUpdatedDetailObject**](AeUserInfoAdvisoryDataSourceUpdatedDetailObject.md) |  | 
**RequestId** | **string** |  | 
**Requester** | **string** |  | 
**SourceIpAddress** | **string** |  | 
**Reason** | **string** |  | 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeUserInfoAdvisoryDataSourceUpdatedDetail

`func NewAeUserInfoAdvisoryDataSourceUpdatedDetail(version string, bucket AeUserInfoAdvisoryDataSourceUpdatedDetailBucket, object AeUserInfoAdvisoryDataSourceUpdatedDetailObject, requestId string, requester string, sourceIpAddress string, reason string, ) *AeUserInfoAdvisoryDataSourceUpdatedDetail`

NewAeUserInfoAdvisoryDataSourceUpdatedDetail instantiates a new AeUserInfoAdvisoryDataSourceUpdatedDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeUserInfoAdvisoryDataSourceUpdatedDetailWithDefaults

`func NewAeUserInfoAdvisoryDataSourceUpdatedDetailWithDefaults() *AeUserInfoAdvisoryDataSourceUpdatedDetail`

NewAeUserInfoAdvisoryDataSourceUpdatedDetailWithDefaults instantiates a new AeUserInfoAdvisoryDataSourceUpdatedDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *AeUserInfoAdvisoryDataSourceUpdatedDetail) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AeUserInfoAdvisoryDataSourceUpdatedDetail) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AeUserInfoAdvisoryDataSourceUpdatedDetail) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetBucket

`func (o *AeUserInfoAdvisoryDataSourceUpdatedDetail) GetBucket() AeUserInfoAdvisoryDataSourceUpdatedDetailBucket`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *AeUserInfoAdvisoryDataSourceUpdatedDetail) GetBucketOk() (*AeUserInfoAdvisoryDataSourceUpdatedDetailBucket, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *AeUserInfoAdvisoryDataSourceUpdatedDetail) SetBucket(v AeUserInfoAdvisoryDataSourceUpdatedDetailBucket)`

SetBucket sets Bucket field to given value.


### GetObject

`func (o *AeUserInfoAdvisoryDataSourceUpdatedDetail) GetObject() AeUserInfoAdvisoryDataSourceUpdatedDetailObject`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *AeUserInfoAdvisoryDataSourceUpdatedDetail) GetObjectOk() (*AeUserInfoAdvisoryDataSourceUpdatedDetailObject, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *AeUserInfoAdvisoryDataSourceUpdatedDetail) SetObject(v AeUserInfoAdvisoryDataSourceUpdatedDetailObject)`

SetObject sets Object field to given value.


### GetRequestId

`func (o *AeUserInfoAdvisoryDataSourceUpdatedDetail) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *AeUserInfoAdvisoryDataSourceUpdatedDetail) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *AeUserInfoAdvisoryDataSourceUpdatedDetail) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetRequester

`func (o *AeUserInfoAdvisoryDataSourceUpdatedDetail) GetRequester() string`

GetRequester returns the Requester field if non-nil, zero value otherwise.

### GetRequesterOk

`func (o *AeUserInfoAdvisoryDataSourceUpdatedDetail) GetRequesterOk() (*string, bool)`

GetRequesterOk returns a tuple with the Requester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequester

`func (o *AeUserInfoAdvisoryDataSourceUpdatedDetail) SetRequester(v string)`

SetRequester sets Requester field to given value.


### GetSourceIpAddress

`func (o *AeUserInfoAdvisoryDataSourceUpdatedDetail) GetSourceIpAddress() string`

GetSourceIpAddress returns the SourceIpAddress field if non-nil, zero value otherwise.

### GetSourceIpAddressOk

`func (o *AeUserInfoAdvisoryDataSourceUpdatedDetail) GetSourceIpAddressOk() (*string, bool)`

GetSourceIpAddressOk returns a tuple with the SourceIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIpAddress

`func (o *AeUserInfoAdvisoryDataSourceUpdatedDetail) SetSourceIpAddress(v string)`

SetSourceIpAddress sets SourceIpAddress field to given value.


### GetReason

`func (o *AeUserInfoAdvisoryDataSourceUpdatedDetail) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AeUserInfoAdvisoryDataSourceUpdatedDetail) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AeUserInfoAdvisoryDataSourceUpdatedDetail) SetReason(v string)`

SetReason sets Reason field to given value.


### GetInitiator

`func (o *AeUserInfoAdvisoryDataSourceUpdatedDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeUserInfoAdvisoryDataSourceUpdatedDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeUserInfoAdvisoryDataSourceUpdatedDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeUserInfoAdvisoryDataSourceUpdatedDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


