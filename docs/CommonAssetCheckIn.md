# CommonAssetCheckIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordedAt** | **time.Time** | Date and time that the Asset first got a stock number in the AMS. | 
**RecordedBy** | Pointer to **string** | The user responsible for checking in the asset. | [optional] 

## Methods

### NewCommonAssetCheckIn

`func NewCommonAssetCheckIn(recordedAt time.Time, ) *CommonAssetCheckIn`

NewCommonAssetCheckIn instantiates a new CommonAssetCheckIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonAssetCheckInWithDefaults

`func NewCommonAssetCheckInWithDefaults() *CommonAssetCheckIn`

NewCommonAssetCheckInWithDefaults instantiates a new CommonAssetCheckIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordedAt

`func (o *CommonAssetCheckIn) GetRecordedAt() time.Time`

GetRecordedAt returns the RecordedAt field if non-nil, zero value otherwise.

### GetRecordedAtOk

`func (o *CommonAssetCheckIn) GetRecordedAtOk() (*time.Time, bool)`

GetRecordedAtOk returns a tuple with the RecordedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordedAt

`func (o *CommonAssetCheckIn) SetRecordedAt(v time.Time)`

SetRecordedAt sets RecordedAt field to given value.


### GetRecordedBy

`func (o *CommonAssetCheckIn) GetRecordedBy() string`

GetRecordedBy returns the RecordedBy field if non-nil, zero value otherwise.

### GetRecordedByOk

`func (o *CommonAssetCheckIn) GetRecordedByOk() (*string, bool)`

GetRecordedByOk returns a tuple with the RecordedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordedBy

`func (o *CommonAssetCheckIn) SetRecordedBy(v string)`

SetRecordedBy sets RecordedBy field to given value.

### HasRecordedBy

`func (o *CommonAssetCheckIn) HasRecordedBy() bool`

HasRecordedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


