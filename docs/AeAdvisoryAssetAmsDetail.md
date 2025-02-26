# AeAdvisoryAssetAmsDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**Id** | **string** | Source&#39;s unique identifier for asset | 
**Vin** | **string** | The Vehicle Identification Number(VIN) of the asset. | 
**Stock** | **string** | The stock number of the asset. | 
**Seller** | [**CommonAmsAccountPointer**](CommonAmsAccountPointer.md) |  | 
**CheckIn** | [**CommonAssetCheckIn**](CommonAssetCheckIn.md) |  | 
**Year** | Pointer to **int32** | The year the vehicle was manufactured. | [optional] 
**Make** | Pointer to **string** | The manufacturer of the vehicle. | [optional] 
**Model** | Pointer to **string** | The model of the vehicle. | [optional] 
**Trimline** | Pointer to **string** | The trim of the vehicle. | [optional] 
**BodyStyle** | Pointer to **string** | The body style of the vehicle | [optional] 
**Deal** | Pointer to [**Deal**](Deal.md) |  | [optional] 
**FloorAmountUsd** | Pointer to **float32** | The lowest amount the seller is willing to take for this asset. | [optional] 
**Interior** | Pointer to [**CommonVehicleInterior**](CommonVehicleInterior.md) |  | [optional] 
**Exterior** | Pointer to [**CommonVehicleExterior**](CommonVehicleExterior.md) |  | [optional] 
**CurrentOdometerReading** | Pointer to **float32** | The vehicle&#39;s odometer reading | [optional] 
**Engine** | Pointer to [**CommonVehicleEngine**](CommonVehicleEngine.md) |  | [optional] 
**TransmissionType** | Pointer to **string** | The type of transmission the asset has. | [optional] 
**Drivetrain** | Pointer to [**CommonVehicleDrivetrain**](CommonVehicleDrivetrain.md) |  | [optional] 
**Title** | Pointer to [**CommonTitle**](CommonTitle.md) |  | [optional] 
**AuctionGrade** | Pointer to **float32** | A value rating the over all condition of the vehicle typically a number between 0-5 (5 being the best). | [optional] 
**UpdatedAt** | Pointer to **NullableTime** | The date and time at which the asset was last updated at. | [optional] 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeAdvisoryAssetAmsDetail

`func NewAeAdvisoryAssetAmsDetail(auctionId string, id string, vin string, stock string, seller CommonAmsAccountPointer, checkIn CommonAssetCheckIn, ) *AeAdvisoryAssetAmsDetail`

NewAeAdvisoryAssetAmsDetail instantiates a new AeAdvisoryAssetAmsDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAdvisoryAssetAmsDetailWithDefaults

`func NewAeAdvisoryAssetAmsDetailWithDefaults() *AeAdvisoryAssetAmsDetail`

NewAeAdvisoryAssetAmsDetailWithDefaults instantiates a new AeAdvisoryAssetAmsDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeAdvisoryAssetAmsDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAdvisoryAssetAmsDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAdvisoryAssetAmsDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetId

`func (o *AeAdvisoryAssetAmsDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AeAdvisoryAssetAmsDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AeAdvisoryAssetAmsDetail) SetId(v string)`

SetId sets Id field to given value.


### GetVin

`func (o *AeAdvisoryAssetAmsDetail) GetVin() string`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *AeAdvisoryAssetAmsDetail) GetVinOk() (*string, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *AeAdvisoryAssetAmsDetail) SetVin(v string)`

SetVin sets Vin field to given value.


### GetStock

`func (o *AeAdvisoryAssetAmsDetail) GetStock() string`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *AeAdvisoryAssetAmsDetail) GetStockOk() (*string, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *AeAdvisoryAssetAmsDetail) SetStock(v string)`

SetStock sets Stock field to given value.


### GetSeller

`func (o *AeAdvisoryAssetAmsDetail) GetSeller() CommonAmsAccountPointer`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *AeAdvisoryAssetAmsDetail) GetSellerOk() (*CommonAmsAccountPointer, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *AeAdvisoryAssetAmsDetail) SetSeller(v CommonAmsAccountPointer)`

SetSeller sets Seller field to given value.


### GetCheckIn

`func (o *AeAdvisoryAssetAmsDetail) GetCheckIn() CommonAssetCheckIn`

GetCheckIn returns the CheckIn field if non-nil, zero value otherwise.

### GetCheckInOk

`func (o *AeAdvisoryAssetAmsDetail) GetCheckInOk() (*CommonAssetCheckIn, bool)`

GetCheckInOk returns a tuple with the CheckIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckIn

`func (o *AeAdvisoryAssetAmsDetail) SetCheckIn(v CommonAssetCheckIn)`

SetCheckIn sets CheckIn field to given value.


### GetYear

`func (o *AeAdvisoryAssetAmsDetail) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *AeAdvisoryAssetAmsDetail) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *AeAdvisoryAssetAmsDetail) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *AeAdvisoryAssetAmsDetail) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetMake

`func (o *AeAdvisoryAssetAmsDetail) GetMake() string`

GetMake returns the Make field if non-nil, zero value otherwise.

### GetMakeOk

`func (o *AeAdvisoryAssetAmsDetail) GetMakeOk() (*string, bool)`

GetMakeOk returns a tuple with the Make field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMake

`func (o *AeAdvisoryAssetAmsDetail) SetMake(v string)`

SetMake sets Make field to given value.

### HasMake

`func (o *AeAdvisoryAssetAmsDetail) HasMake() bool`

HasMake returns a boolean if a field has been set.

### GetModel

`func (o *AeAdvisoryAssetAmsDetail) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AeAdvisoryAssetAmsDetail) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AeAdvisoryAssetAmsDetail) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AeAdvisoryAssetAmsDetail) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetTrimline

`func (o *AeAdvisoryAssetAmsDetail) GetTrimline() string`

GetTrimline returns the Trimline field if non-nil, zero value otherwise.

### GetTrimlineOk

`func (o *AeAdvisoryAssetAmsDetail) GetTrimlineOk() (*string, bool)`

GetTrimlineOk returns a tuple with the Trimline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrimline

`func (o *AeAdvisoryAssetAmsDetail) SetTrimline(v string)`

SetTrimline sets Trimline field to given value.

### HasTrimline

`func (o *AeAdvisoryAssetAmsDetail) HasTrimline() bool`

HasTrimline returns a boolean if a field has been set.

### GetBodyStyle

`func (o *AeAdvisoryAssetAmsDetail) GetBodyStyle() string`

GetBodyStyle returns the BodyStyle field if non-nil, zero value otherwise.

### GetBodyStyleOk

`func (o *AeAdvisoryAssetAmsDetail) GetBodyStyleOk() (*string, bool)`

GetBodyStyleOk returns a tuple with the BodyStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyStyle

`func (o *AeAdvisoryAssetAmsDetail) SetBodyStyle(v string)`

SetBodyStyle sets BodyStyle field to given value.

### HasBodyStyle

`func (o *AeAdvisoryAssetAmsDetail) HasBodyStyle() bool`

HasBodyStyle returns a boolean if a field has been set.

### GetDeal

`func (o *AeAdvisoryAssetAmsDetail) GetDeal() Deal`

GetDeal returns the Deal field if non-nil, zero value otherwise.

### GetDealOk

`func (o *AeAdvisoryAssetAmsDetail) GetDealOk() (*Deal, bool)`

GetDealOk returns a tuple with the Deal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeal

`func (o *AeAdvisoryAssetAmsDetail) SetDeal(v Deal)`

SetDeal sets Deal field to given value.

### HasDeal

`func (o *AeAdvisoryAssetAmsDetail) HasDeal() bool`

HasDeal returns a boolean if a field has been set.

### GetFloorAmountUsd

`func (o *AeAdvisoryAssetAmsDetail) GetFloorAmountUsd() float32`

GetFloorAmountUsd returns the FloorAmountUsd field if non-nil, zero value otherwise.

### GetFloorAmountUsdOk

`func (o *AeAdvisoryAssetAmsDetail) GetFloorAmountUsdOk() (*float32, bool)`

GetFloorAmountUsdOk returns a tuple with the FloorAmountUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorAmountUsd

`func (o *AeAdvisoryAssetAmsDetail) SetFloorAmountUsd(v float32)`

SetFloorAmountUsd sets FloorAmountUsd field to given value.

### HasFloorAmountUsd

`func (o *AeAdvisoryAssetAmsDetail) HasFloorAmountUsd() bool`

HasFloorAmountUsd returns a boolean if a field has been set.

### GetInterior

`func (o *AeAdvisoryAssetAmsDetail) GetInterior() CommonVehicleInterior`

GetInterior returns the Interior field if non-nil, zero value otherwise.

### GetInteriorOk

`func (o *AeAdvisoryAssetAmsDetail) GetInteriorOk() (*CommonVehicleInterior, bool)`

GetInteriorOk returns a tuple with the Interior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterior

`func (o *AeAdvisoryAssetAmsDetail) SetInterior(v CommonVehicleInterior)`

SetInterior sets Interior field to given value.

### HasInterior

`func (o *AeAdvisoryAssetAmsDetail) HasInterior() bool`

HasInterior returns a boolean if a field has been set.

### GetExterior

`func (o *AeAdvisoryAssetAmsDetail) GetExterior() CommonVehicleExterior`

GetExterior returns the Exterior field if non-nil, zero value otherwise.

### GetExteriorOk

`func (o *AeAdvisoryAssetAmsDetail) GetExteriorOk() (*CommonVehicleExterior, bool)`

GetExteriorOk returns a tuple with the Exterior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExterior

`func (o *AeAdvisoryAssetAmsDetail) SetExterior(v CommonVehicleExterior)`

SetExterior sets Exterior field to given value.

### HasExterior

`func (o *AeAdvisoryAssetAmsDetail) HasExterior() bool`

HasExterior returns a boolean if a field has been set.

### GetCurrentOdometerReading

`func (o *AeAdvisoryAssetAmsDetail) GetCurrentOdometerReading() float32`

GetCurrentOdometerReading returns the CurrentOdometerReading field if non-nil, zero value otherwise.

### GetCurrentOdometerReadingOk

`func (o *AeAdvisoryAssetAmsDetail) GetCurrentOdometerReadingOk() (*float32, bool)`

GetCurrentOdometerReadingOk returns a tuple with the CurrentOdometerReading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentOdometerReading

`func (o *AeAdvisoryAssetAmsDetail) SetCurrentOdometerReading(v float32)`

SetCurrentOdometerReading sets CurrentOdometerReading field to given value.

### HasCurrentOdometerReading

`func (o *AeAdvisoryAssetAmsDetail) HasCurrentOdometerReading() bool`

HasCurrentOdometerReading returns a boolean if a field has been set.

### GetEngine

`func (o *AeAdvisoryAssetAmsDetail) GetEngine() CommonVehicleEngine`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *AeAdvisoryAssetAmsDetail) GetEngineOk() (*CommonVehicleEngine, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *AeAdvisoryAssetAmsDetail) SetEngine(v CommonVehicleEngine)`

SetEngine sets Engine field to given value.

### HasEngine

`func (o *AeAdvisoryAssetAmsDetail) HasEngine() bool`

HasEngine returns a boolean if a field has been set.

### GetTransmissionType

`func (o *AeAdvisoryAssetAmsDetail) GetTransmissionType() string`

GetTransmissionType returns the TransmissionType field if non-nil, zero value otherwise.

### GetTransmissionTypeOk

`func (o *AeAdvisoryAssetAmsDetail) GetTransmissionTypeOk() (*string, bool)`

GetTransmissionTypeOk returns a tuple with the TransmissionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmissionType

`func (o *AeAdvisoryAssetAmsDetail) SetTransmissionType(v string)`

SetTransmissionType sets TransmissionType field to given value.

### HasTransmissionType

`func (o *AeAdvisoryAssetAmsDetail) HasTransmissionType() bool`

HasTransmissionType returns a boolean if a field has been set.

### GetDrivetrain

`func (o *AeAdvisoryAssetAmsDetail) GetDrivetrain() CommonVehicleDrivetrain`

GetDrivetrain returns the Drivetrain field if non-nil, zero value otherwise.

### GetDrivetrainOk

`func (o *AeAdvisoryAssetAmsDetail) GetDrivetrainOk() (*CommonVehicleDrivetrain, bool)`

GetDrivetrainOk returns a tuple with the Drivetrain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivetrain

`func (o *AeAdvisoryAssetAmsDetail) SetDrivetrain(v CommonVehicleDrivetrain)`

SetDrivetrain sets Drivetrain field to given value.

### HasDrivetrain

`func (o *AeAdvisoryAssetAmsDetail) HasDrivetrain() bool`

HasDrivetrain returns a boolean if a field has been set.

### GetTitle

`func (o *AeAdvisoryAssetAmsDetail) GetTitle() CommonTitle`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AeAdvisoryAssetAmsDetail) GetTitleOk() (*CommonTitle, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AeAdvisoryAssetAmsDetail) SetTitle(v CommonTitle)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AeAdvisoryAssetAmsDetail) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetAuctionGrade

`func (o *AeAdvisoryAssetAmsDetail) GetAuctionGrade() float32`

GetAuctionGrade returns the AuctionGrade field if non-nil, zero value otherwise.

### GetAuctionGradeOk

`func (o *AeAdvisoryAssetAmsDetail) GetAuctionGradeOk() (*float32, bool)`

GetAuctionGradeOk returns a tuple with the AuctionGrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionGrade

`func (o *AeAdvisoryAssetAmsDetail) SetAuctionGrade(v float32)`

SetAuctionGrade sets AuctionGrade field to given value.

### HasAuctionGrade

`func (o *AeAdvisoryAssetAmsDetail) HasAuctionGrade() bool`

HasAuctionGrade returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AeAdvisoryAssetAmsDetail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AeAdvisoryAssetAmsDetail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AeAdvisoryAssetAmsDetail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AeAdvisoryAssetAmsDetail) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *AeAdvisoryAssetAmsDetail) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *AeAdvisoryAssetAmsDetail) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetInitiator

`func (o *AeAdvisoryAssetAmsDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAdvisoryAssetAmsDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAdvisoryAssetAmsDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAdvisoryAssetAmsDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


