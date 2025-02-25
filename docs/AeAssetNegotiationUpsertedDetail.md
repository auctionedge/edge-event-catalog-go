# AeAssetNegotiationUpsertedDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**Id** | **string** | Uniquie identifier for a specific negotiation. | 
**Asset** | [**CommonAmsAssetPointer**](CommonAmsAssetPointer.md) |  | 
**Conditions** | [**AeAssetNegotiationUpsertedDetailConditions**](AeAssetNegotiationUpsertedDetailConditions.md) |  | 
**InitialFloorAmount** | Pointer to [**CommonCurrency**](CommonCurrency.md) |  | [optional] 
**OpenedAt** | **time.Time** |  | 
**ClosedAt** | **NullableTime** |  | 
**Successful** | **NullableBool** |  | 
**FailureReason** | Pointer to **string** |  | [optional] 
**BallControl** | **string** |  | 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeAssetNegotiationUpsertedDetail

`func NewAeAssetNegotiationUpsertedDetail(auctionId string, id string, asset CommonAmsAssetPointer, conditions AeAssetNegotiationUpsertedDetailConditions, openedAt time.Time, closedAt NullableTime, successful NullableBool, ballControl string, ) *AeAssetNegotiationUpsertedDetail`

NewAeAssetNegotiationUpsertedDetail instantiates a new AeAssetNegotiationUpsertedDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetNegotiationUpsertedDetailWithDefaults

`func NewAeAssetNegotiationUpsertedDetailWithDefaults() *AeAssetNegotiationUpsertedDetail`

NewAeAssetNegotiationUpsertedDetailWithDefaults instantiates a new AeAssetNegotiationUpsertedDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeAssetNegotiationUpsertedDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAssetNegotiationUpsertedDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAssetNegotiationUpsertedDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetId

`func (o *AeAssetNegotiationUpsertedDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AeAssetNegotiationUpsertedDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AeAssetNegotiationUpsertedDetail) SetId(v string)`

SetId sets Id field to given value.


### GetAsset

`func (o *AeAssetNegotiationUpsertedDetail) GetAsset() CommonAmsAssetPointer`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *AeAssetNegotiationUpsertedDetail) GetAssetOk() (*CommonAmsAssetPointer, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *AeAssetNegotiationUpsertedDetail) SetAsset(v CommonAmsAssetPointer)`

SetAsset sets Asset field to given value.


### GetConditions

`func (o *AeAssetNegotiationUpsertedDetail) GetConditions() AeAssetNegotiationUpsertedDetailConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *AeAssetNegotiationUpsertedDetail) GetConditionsOk() (*AeAssetNegotiationUpsertedDetailConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *AeAssetNegotiationUpsertedDetail) SetConditions(v AeAssetNegotiationUpsertedDetailConditions)`

SetConditions sets Conditions field to given value.


### GetInitialFloorAmount

`func (o *AeAssetNegotiationUpsertedDetail) GetInitialFloorAmount() CommonCurrency`

GetInitialFloorAmount returns the InitialFloorAmount field if non-nil, zero value otherwise.

### GetInitialFloorAmountOk

`func (o *AeAssetNegotiationUpsertedDetail) GetInitialFloorAmountOk() (*CommonCurrency, bool)`

GetInitialFloorAmountOk returns a tuple with the InitialFloorAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialFloorAmount

`func (o *AeAssetNegotiationUpsertedDetail) SetInitialFloorAmount(v CommonCurrency)`

SetInitialFloorAmount sets InitialFloorAmount field to given value.

### HasInitialFloorAmount

`func (o *AeAssetNegotiationUpsertedDetail) HasInitialFloorAmount() bool`

HasInitialFloorAmount returns a boolean if a field has been set.

### GetOpenedAt

`func (o *AeAssetNegotiationUpsertedDetail) GetOpenedAt() time.Time`

GetOpenedAt returns the OpenedAt field if non-nil, zero value otherwise.

### GetOpenedAtOk

`func (o *AeAssetNegotiationUpsertedDetail) GetOpenedAtOk() (*time.Time, bool)`

GetOpenedAtOk returns a tuple with the OpenedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenedAt

`func (o *AeAssetNegotiationUpsertedDetail) SetOpenedAt(v time.Time)`

SetOpenedAt sets OpenedAt field to given value.


### GetClosedAt

`func (o *AeAssetNegotiationUpsertedDetail) GetClosedAt() time.Time`

GetClosedAt returns the ClosedAt field if non-nil, zero value otherwise.

### GetClosedAtOk

`func (o *AeAssetNegotiationUpsertedDetail) GetClosedAtOk() (*time.Time, bool)`

GetClosedAtOk returns a tuple with the ClosedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedAt

`func (o *AeAssetNegotiationUpsertedDetail) SetClosedAt(v time.Time)`

SetClosedAt sets ClosedAt field to given value.


### SetClosedAtNil

`func (o *AeAssetNegotiationUpsertedDetail) SetClosedAtNil(b bool)`

 SetClosedAtNil sets the value for ClosedAt to be an explicit nil

### UnsetClosedAt
`func (o *AeAssetNegotiationUpsertedDetail) UnsetClosedAt()`

UnsetClosedAt ensures that no value is present for ClosedAt, not even an explicit nil
### GetSuccessful

`func (o *AeAssetNegotiationUpsertedDetail) GetSuccessful() bool`

GetSuccessful returns the Successful field if non-nil, zero value otherwise.

### GetSuccessfulOk

`func (o *AeAssetNegotiationUpsertedDetail) GetSuccessfulOk() (*bool, bool)`

GetSuccessfulOk returns a tuple with the Successful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessful

`func (o *AeAssetNegotiationUpsertedDetail) SetSuccessful(v bool)`

SetSuccessful sets Successful field to given value.


### SetSuccessfulNil

`func (o *AeAssetNegotiationUpsertedDetail) SetSuccessfulNil(b bool)`

 SetSuccessfulNil sets the value for Successful to be an explicit nil

### UnsetSuccessful
`func (o *AeAssetNegotiationUpsertedDetail) UnsetSuccessful()`

UnsetSuccessful ensures that no value is present for Successful, not even an explicit nil
### GetFailureReason

`func (o *AeAssetNegotiationUpsertedDetail) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *AeAssetNegotiationUpsertedDetail) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *AeAssetNegotiationUpsertedDetail) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *AeAssetNegotiationUpsertedDetail) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetBallControl

`func (o *AeAssetNegotiationUpsertedDetail) GetBallControl() string`

GetBallControl returns the BallControl field if non-nil, zero value otherwise.

### GetBallControlOk

`func (o *AeAssetNegotiationUpsertedDetail) GetBallControlOk() (*string, bool)`

GetBallControlOk returns a tuple with the BallControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBallControl

`func (o *AeAssetNegotiationUpsertedDetail) SetBallControl(v string)`

SetBallControl sets BallControl field to given value.


### GetInitiator

`func (o *AeAssetNegotiationUpsertedDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAssetNegotiationUpsertedDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAssetNegotiationUpsertedDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAssetNegotiationUpsertedDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


