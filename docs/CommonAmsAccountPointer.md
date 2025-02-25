# CommonAmsAccountPointer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountKey** | **string** | The account key of an AMS account (i.e. the account or dealer number) | 
**AaId** | **NullableString** | The Auction Access ID of the AMS account. | 
**AccountId** | Pointer to **string** | Source&#39;s unique identifier for account | [optional] 

## Methods

### NewCommonAmsAccountPointer

`func NewCommonAmsAccountPointer(accountKey string, aaId NullableString, ) *CommonAmsAccountPointer`

NewCommonAmsAccountPointer instantiates a new CommonAmsAccountPointer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonAmsAccountPointerWithDefaults

`func NewCommonAmsAccountPointerWithDefaults() *CommonAmsAccountPointer`

NewCommonAmsAccountPointerWithDefaults instantiates a new CommonAmsAccountPointer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountKey

`func (o *CommonAmsAccountPointer) GetAccountKey() string`

GetAccountKey returns the AccountKey field if non-nil, zero value otherwise.

### GetAccountKeyOk

`func (o *CommonAmsAccountPointer) GetAccountKeyOk() (*string, bool)`

GetAccountKeyOk returns a tuple with the AccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountKey

`func (o *CommonAmsAccountPointer) SetAccountKey(v string)`

SetAccountKey sets AccountKey field to given value.


### GetAaId

`func (o *CommonAmsAccountPointer) GetAaId() string`

GetAaId returns the AaId field if non-nil, zero value otherwise.

### GetAaIdOk

`func (o *CommonAmsAccountPointer) GetAaIdOk() (*string, bool)`

GetAaIdOk returns a tuple with the AaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAaId

`func (o *CommonAmsAccountPointer) SetAaId(v string)`

SetAaId sets AaId field to given value.


### SetAaIdNil

`func (o *CommonAmsAccountPointer) SetAaIdNil(b bool)`

 SetAaIdNil sets the value for AaId to be an explicit nil

### UnsetAaId
`func (o *CommonAmsAccountPointer) UnsetAaId()`

UnsetAaId ensures that no value is present for AaId, not even an explicit nil
### GetAccountId

`func (o *CommonAmsAccountPointer) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CommonAmsAccountPointer) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CommonAmsAccountPointer) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CommonAmsAccountPointer) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


