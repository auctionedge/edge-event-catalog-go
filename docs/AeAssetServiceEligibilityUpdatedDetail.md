# AeAssetServiceEligibilityUpdatedDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**Asset** | [**CommonAmsAssetPointer**](CommonAmsAssetPointer.md) |  | 
**EligibleServices** | Pointer to [**[]Service**](Service.md) |  | [optional] 
**IneligibleServices** | Pointer to [**[]Service**](Service.md) |  | [optional] 
**UpdatedAt** | **time.Time** | The updated date time of the gatepass status | 
**Initiator** | [**CommonInitiator**](CommonInitiator.md) |  | 

## Methods

### NewAeAssetServiceEligibilityUpdatedDetail

`func NewAeAssetServiceEligibilityUpdatedDetail(auctionId string, asset CommonAmsAssetPointer, updatedAt time.Time, initiator CommonInitiator, ) *AeAssetServiceEligibilityUpdatedDetail`

NewAeAssetServiceEligibilityUpdatedDetail instantiates a new AeAssetServiceEligibilityUpdatedDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetServiceEligibilityUpdatedDetailWithDefaults

`func NewAeAssetServiceEligibilityUpdatedDetailWithDefaults() *AeAssetServiceEligibilityUpdatedDetail`

NewAeAssetServiceEligibilityUpdatedDetailWithDefaults instantiates a new AeAssetServiceEligibilityUpdatedDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeAssetServiceEligibilityUpdatedDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAssetServiceEligibilityUpdatedDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAssetServiceEligibilityUpdatedDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetAsset

`func (o *AeAssetServiceEligibilityUpdatedDetail) GetAsset() CommonAmsAssetPointer`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *AeAssetServiceEligibilityUpdatedDetail) GetAssetOk() (*CommonAmsAssetPointer, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *AeAssetServiceEligibilityUpdatedDetail) SetAsset(v CommonAmsAssetPointer)`

SetAsset sets Asset field to given value.


### GetEligibleServices

`func (o *AeAssetServiceEligibilityUpdatedDetail) GetEligibleServices() []Service`

GetEligibleServices returns the EligibleServices field if non-nil, zero value otherwise.

### GetEligibleServicesOk

`func (o *AeAssetServiceEligibilityUpdatedDetail) GetEligibleServicesOk() (*[]Service, bool)`

GetEligibleServicesOk returns a tuple with the EligibleServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibleServices

`func (o *AeAssetServiceEligibilityUpdatedDetail) SetEligibleServices(v []Service)`

SetEligibleServices sets EligibleServices field to given value.

### HasEligibleServices

`func (o *AeAssetServiceEligibilityUpdatedDetail) HasEligibleServices() bool`

HasEligibleServices returns a boolean if a field has been set.

### GetIneligibleServices

`func (o *AeAssetServiceEligibilityUpdatedDetail) GetIneligibleServices() []Service`

GetIneligibleServices returns the IneligibleServices field if non-nil, zero value otherwise.

### GetIneligibleServicesOk

`func (o *AeAssetServiceEligibilityUpdatedDetail) GetIneligibleServicesOk() (*[]Service, bool)`

GetIneligibleServicesOk returns a tuple with the IneligibleServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIneligibleServices

`func (o *AeAssetServiceEligibilityUpdatedDetail) SetIneligibleServices(v []Service)`

SetIneligibleServices sets IneligibleServices field to given value.

### HasIneligibleServices

`func (o *AeAssetServiceEligibilityUpdatedDetail) HasIneligibleServices() bool`

HasIneligibleServices returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AeAssetServiceEligibilityUpdatedDetail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AeAssetServiceEligibilityUpdatedDetail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AeAssetServiceEligibilityUpdatedDetail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetInitiator

`func (o *AeAssetServiceEligibilityUpdatedDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAssetServiceEligibilityUpdatedDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAssetServiceEligibilityUpdatedDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


