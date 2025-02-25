# AeAssetSellerChargeUpsertDetailCharge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | AMS&#39;s unique identifier for the charge | 
**DisplayName** | **string** | Description of the charge | 
**ChargeType** | **string** | The type of charge that is being requested to be added. | 
**Amount** | **float32** | Amount to be charged to the seller | 
**Cost** | Pointer to **float32** | Amount of money the auction is out related to the charge. | [optional] 
**Note** | Pointer to **string** | Notes for the auction relating to the charge | [optional] 

## Methods

### NewAeAssetSellerChargeUpsertDetailCharge

`func NewAeAssetSellerChargeUpsertDetailCharge(id string, displayName string, chargeType string, amount float32, ) *AeAssetSellerChargeUpsertDetailCharge`

NewAeAssetSellerChargeUpsertDetailCharge instantiates a new AeAssetSellerChargeUpsertDetailCharge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetSellerChargeUpsertDetailChargeWithDefaults

`func NewAeAssetSellerChargeUpsertDetailChargeWithDefaults() *AeAssetSellerChargeUpsertDetailCharge`

NewAeAssetSellerChargeUpsertDetailChargeWithDefaults instantiates a new AeAssetSellerChargeUpsertDetailCharge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AeAssetSellerChargeUpsertDetailCharge) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AeAssetSellerChargeUpsertDetailCharge) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AeAssetSellerChargeUpsertDetailCharge) SetId(v string)`

SetId sets Id field to given value.


### GetDisplayName

`func (o *AeAssetSellerChargeUpsertDetailCharge) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AeAssetSellerChargeUpsertDetailCharge) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AeAssetSellerChargeUpsertDetailCharge) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetChargeType

`func (o *AeAssetSellerChargeUpsertDetailCharge) GetChargeType() string`

GetChargeType returns the ChargeType field if non-nil, zero value otherwise.

### GetChargeTypeOk

`func (o *AeAssetSellerChargeUpsertDetailCharge) GetChargeTypeOk() (*string, bool)`

GetChargeTypeOk returns a tuple with the ChargeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeType

`func (o *AeAssetSellerChargeUpsertDetailCharge) SetChargeType(v string)`

SetChargeType sets ChargeType field to given value.


### GetAmount

`func (o *AeAssetSellerChargeUpsertDetailCharge) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AeAssetSellerChargeUpsertDetailCharge) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AeAssetSellerChargeUpsertDetailCharge) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCost

`func (o *AeAssetSellerChargeUpsertDetailCharge) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *AeAssetSellerChargeUpsertDetailCharge) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *AeAssetSellerChargeUpsertDetailCharge) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *AeAssetSellerChargeUpsertDetailCharge) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetNote

`func (o *AeAssetSellerChargeUpsertDetailCharge) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *AeAssetSellerChargeUpsertDetailCharge) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *AeAssetSellerChargeUpsertDetailCharge) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *AeAssetSellerChargeUpsertDetailCharge) HasNote() bool`

HasNote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


