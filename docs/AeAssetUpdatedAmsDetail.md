# AeAssetUpdatedAmsDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Source&#39;s unique identifier for asset | 
**Stock** | **string** | The stock number of the asset. | 
**Vin** | **string** | The Vehicle Identification Number(VIN) of the asset. | 
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**Seller** | [**CommonAmsAccountPointer**](CommonAmsAccountPointer.md) |  | 
**CheckIn** | [**CommonAssetCheckIn**](CommonAssetCheckIn.md) |  | 
**Year** | Pointer to **int32** | The year the vehicle was manufactured. | [optional] 
**Make** | Pointer to **string** | The manufacturer of the vehicle. | [optional] 
**Model** | Pointer to **string** | The model of the vehicle. | [optional] 
**Trimline** | Pointer to **string** | The trim of the vehicle. | [optional] 
**BodyStyle** | Pointer to **string** | The body style of the vehicle | [optional] 
**FloorAmount** | Pointer to [**CommonCurrency**](CommonCurrency.md) |  | [optional] 
**FloorAmountUsd** | Pointer to **float32** | The lowest amount the seller is willing to take for this asset. | [optional] 
**Interior** | Pointer to [**CommonVehicleInterior**](CommonVehicleInterior.md) |  | [optional] 
**Exterior** | Pointer to [**CommonVehicleExterior**](CommonVehicleExterior.md) |  | [optional] 
**CurrentOdometerReading** | Pointer to **float32** | The vehicle&#39;s odometer reading | [optional] 
**Engine** | Pointer to [**CommonVehicleEngine**](CommonVehicleEngine.md) |  | [optional] 
**TransmissionType** | Pointer to **string** | The type of transmission the asset has. | [optional] 
**Drivetrain** | Pointer to [**CommonVehicleDrivetrain**](CommonVehicleDrivetrain.md) |  | [optional] 
**AuctionGrade** | Pointer to **float32** | A value rating the over all condition of the vehicle typically a number between 0-5 (5 being the best). | [optional] 
**UpdatedAt** | Pointer to **NullableTime** | The date and time at which the asset was last updated at. | [optional] 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 
**Valuations** | Pointer to [**[]Valuation**](Valuation.md) | The valuations of the asset. | [optional] 
**Title** | Pointer to [**CommonTitle**](CommonTitle.md) |  | [optional] 
**PresaleAnnouncements** | Pointer to **string** | The announcements on the asset. | [optional] 
**Location** | Pointer to **string** | The location of the asset. | [optional] 
**AutoGrade** | Pointer to **float32** | A value rating the over all condition of the vehicle typically a number between 0-5 (5 being the best). | [optional] 

## Methods

### NewAeAssetUpdatedAmsDetail

`func NewAeAssetUpdatedAmsDetail(id string, stock string, vin string, auctionId string, seller CommonAmsAccountPointer, checkIn CommonAssetCheckIn, ) *AeAssetUpdatedAmsDetail`

NewAeAssetUpdatedAmsDetail instantiates a new AeAssetUpdatedAmsDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetUpdatedAmsDetailWithDefaults

`func NewAeAssetUpdatedAmsDetailWithDefaults() *AeAssetUpdatedAmsDetail`

NewAeAssetUpdatedAmsDetailWithDefaults instantiates a new AeAssetUpdatedAmsDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AeAssetUpdatedAmsDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AeAssetUpdatedAmsDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AeAssetUpdatedAmsDetail) SetId(v string)`

SetId sets Id field to given value.


### GetStock

`func (o *AeAssetUpdatedAmsDetail) GetStock() string`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *AeAssetUpdatedAmsDetail) GetStockOk() (*string, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *AeAssetUpdatedAmsDetail) SetStock(v string)`

SetStock sets Stock field to given value.


### GetVin

`func (o *AeAssetUpdatedAmsDetail) GetVin() string`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *AeAssetUpdatedAmsDetail) GetVinOk() (*string, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *AeAssetUpdatedAmsDetail) SetVin(v string)`

SetVin sets Vin field to given value.


### GetAuctionId

`func (o *AeAssetUpdatedAmsDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAssetUpdatedAmsDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAssetUpdatedAmsDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetSeller

`func (o *AeAssetUpdatedAmsDetail) GetSeller() CommonAmsAccountPointer`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *AeAssetUpdatedAmsDetail) GetSellerOk() (*CommonAmsAccountPointer, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *AeAssetUpdatedAmsDetail) SetSeller(v CommonAmsAccountPointer)`

SetSeller sets Seller field to given value.


### GetCheckIn

`func (o *AeAssetUpdatedAmsDetail) GetCheckIn() CommonAssetCheckIn`

GetCheckIn returns the CheckIn field if non-nil, zero value otherwise.

### GetCheckInOk

`func (o *AeAssetUpdatedAmsDetail) GetCheckInOk() (*CommonAssetCheckIn, bool)`

GetCheckInOk returns a tuple with the CheckIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckIn

`func (o *AeAssetUpdatedAmsDetail) SetCheckIn(v CommonAssetCheckIn)`

SetCheckIn sets CheckIn field to given value.


### GetYear

`func (o *AeAssetUpdatedAmsDetail) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *AeAssetUpdatedAmsDetail) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *AeAssetUpdatedAmsDetail) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *AeAssetUpdatedAmsDetail) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetMake

`func (o *AeAssetUpdatedAmsDetail) GetMake() string`

GetMake returns the Make field if non-nil, zero value otherwise.

### GetMakeOk

`func (o *AeAssetUpdatedAmsDetail) GetMakeOk() (*string, bool)`

GetMakeOk returns a tuple with the Make field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMake

`func (o *AeAssetUpdatedAmsDetail) SetMake(v string)`

SetMake sets Make field to given value.

### HasMake

`func (o *AeAssetUpdatedAmsDetail) HasMake() bool`

HasMake returns a boolean if a field has been set.

### GetModel

`func (o *AeAssetUpdatedAmsDetail) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AeAssetUpdatedAmsDetail) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AeAssetUpdatedAmsDetail) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AeAssetUpdatedAmsDetail) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetTrimline

`func (o *AeAssetUpdatedAmsDetail) GetTrimline() string`

GetTrimline returns the Trimline field if non-nil, zero value otherwise.

### GetTrimlineOk

`func (o *AeAssetUpdatedAmsDetail) GetTrimlineOk() (*string, bool)`

GetTrimlineOk returns a tuple with the Trimline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrimline

`func (o *AeAssetUpdatedAmsDetail) SetTrimline(v string)`

SetTrimline sets Trimline field to given value.

### HasTrimline

`func (o *AeAssetUpdatedAmsDetail) HasTrimline() bool`

HasTrimline returns a boolean if a field has been set.

### GetBodyStyle

`func (o *AeAssetUpdatedAmsDetail) GetBodyStyle() string`

GetBodyStyle returns the BodyStyle field if non-nil, zero value otherwise.

### GetBodyStyleOk

`func (o *AeAssetUpdatedAmsDetail) GetBodyStyleOk() (*string, bool)`

GetBodyStyleOk returns a tuple with the BodyStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyStyle

`func (o *AeAssetUpdatedAmsDetail) SetBodyStyle(v string)`

SetBodyStyle sets BodyStyle field to given value.

### HasBodyStyle

`func (o *AeAssetUpdatedAmsDetail) HasBodyStyle() bool`

HasBodyStyle returns a boolean if a field has been set.

### GetFloorAmount

`func (o *AeAssetUpdatedAmsDetail) GetFloorAmount() CommonCurrency`

GetFloorAmount returns the FloorAmount field if non-nil, zero value otherwise.

### GetFloorAmountOk

`func (o *AeAssetUpdatedAmsDetail) GetFloorAmountOk() (*CommonCurrency, bool)`

GetFloorAmountOk returns a tuple with the FloorAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorAmount

`func (o *AeAssetUpdatedAmsDetail) SetFloorAmount(v CommonCurrency)`

SetFloorAmount sets FloorAmount field to given value.

### HasFloorAmount

`func (o *AeAssetUpdatedAmsDetail) HasFloorAmount() bool`

HasFloorAmount returns a boolean if a field has been set.

### GetFloorAmountUsd

`func (o *AeAssetUpdatedAmsDetail) GetFloorAmountUsd() float32`

GetFloorAmountUsd returns the FloorAmountUsd field if non-nil, zero value otherwise.

### GetFloorAmountUsdOk

`func (o *AeAssetUpdatedAmsDetail) GetFloorAmountUsdOk() (*float32, bool)`

GetFloorAmountUsdOk returns a tuple with the FloorAmountUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorAmountUsd

`func (o *AeAssetUpdatedAmsDetail) SetFloorAmountUsd(v float32)`

SetFloorAmountUsd sets FloorAmountUsd field to given value.

### HasFloorAmountUsd

`func (o *AeAssetUpdatedAmsDetail) HasFloorAmountUsd() bool`

HasFloorAmountUsd returns a boolean if a field has been set.

### GetInterior

`func (o *AeAssetUpdatedAmsDetail) GetInterior() CommonVehicleInterior`

GetInterior returns the Interior field if non-nil, zero value otherwise.

### GetInteriorOk

`func (o *AeAssetUpdatedAmsDetail) GetInteriorOk() (*CommonVehicleInterior, bool)`

GetInteriorOk returns a tuple with the Interior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterior

`func (o *AeAssetUpdatedAmsDetail) SetInterior(v CommonVehicleInterior)`

SetInterior sets Interior field to given value.

### HasInterior

`func (o *AeAssetUpdatedAmsDetail) HasInterior() bool`

HasInterior returns a boolean if a field has been set.

### GetExterior

`func (o *AeAssetUpdatedAmsDetail) GetExterior() CommonVehicleExterior`

GetExterior returns the Exterior field if non-nil, zero value otherwise.

### GetExteriorOk

`func (o *AeAssetUpdatedAmsDetail) GetExteriorOk() (*CommonVehicleExterior, bool)`

GetExteriorOk returns a tuple with the Exterior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExterior

`func (o *AeAssetUpdatedAmsDetail) SetExterior(v CommonVehicleExterior)`

SetExterior sets Exterior field to given value.

### HasExterior

`func (o *AeAssetUpdatedAmsDetail) HasExterior() bool`

HasExterior returns a boolean if a field has been set.

### GetCurrentOdometerReading

`func (o *AeAssetUpdatedAmsDetail) GetCurrentOdometerReading() float32`

GetCurrentOdometerReading returns the CurrentOdometerReading field if non-nil, zero value otherwise.

### GetCurrentOdometerReadingOk

`func (o *AeAssetUpdatedAmsDetail) GetCurrentOdometerReadingOk() (*float32, bool)`

GetCurrentOdometerReadingOk returns a tuple with the CurrentOdometerReading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentOdometerReading

`func (o *AeAssetUpdatedAmsDetail) SetCurrentOdometerReading(v float32)`

SetCurrentOdometerReading sets CurrentOdometerReading field to given value.

### HasCurrentOdometerReading

`func (o *AeAssetUpdatedAmsDetail) HasCurrentOdometerReading() bool`

HasCurrentOdometerReading returns a boolean if a field has been set.

### GetEngine

`func (o *AeAssetUpdatedAmsDetail) GetEngine() CommonVehicleEngine`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *AeAssetUpdatedAmsDetail) GetEngineOk() (*CommonVehicleEngine, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *AeAssetUpdatedAmsDetail) SetEngine(v CommonVehicleEngine)`

SetEngine sets Engine field to given value.

### HasEngine

`func (o *AeAssetUpdatedAmsDetail) HasEngine() bool`

HasEngine returns a boolean if a field has been set.

### GetTransmissionType

`func (o *AeAssetUpdatedAmsDetail) GetTransmissionType() string`

GetTransmissionType returns the TransmissionType field if non-nil, zero value otherwise.

### GetTransmissionTypeOk

`func (o *AeAssetUpdatedAmsDetail) GetTransmissionTypeOk() (*string, bool)`

GetTransmissionTypeOk returns a tuple with the TransmissionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmissionType

`func (o *AeAssetUpdatedAmsDetail) SetTransmissionType(v string)`

SetTransmissionType sets TransmissionType field to given value.

### HasTransmissionType

`func (o *AeAssetUpdatedAmsDetail) HasTransmissionType() bool`

HasTransmissionType returns a boolean if a field has been set.

### GetDrivetrain

`func (o *AeAssetUpdatedAmsDetail) GetDrivetrain() CommonVehicleDrivetrain`

GetDrivetrain returns the Drivetrain field if non-nil, zero value otherwise.

### GetDrivetrainOk

`func (o *AeAssetUpdatedAmsDetail) GetDrivetrainOk() (*CommonVehicleDrivetrain, bool)`

GetDrivetrainOk returns a tuple with the Drivetrain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivetrain

`func (o *AeAssetUpdatedAmsDetail) SetDrivetrain(v CommonVehicleDrivetrain)`

SetDrivetrain sets Drivetrain field to given value.

### HasDrivetrain

`func (o *AeAssetUpdatedAmsDetail) HasDrivetrain() bool`

HasDrivetrain returns a boolean if a field has been set.

### GetAuctionGrade

`func (o *AeAssetUpdatedAmsDetail) GetAuctionGrade() float32`

GetAuctionGrade returns the AuctionGrade field if non-nil, zero value otherwise.

### GetAuctionGradeOk

`func (o *AeAssetUpdatedAmsDetail) GetAuctionGradeOk() (*float32, bool)`

GetAuctionGradeOk returns a tuple with the AuctionGrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionGrade

`func (o *AeAssetUpdatedAmsDetail) SetAuctionGrade(v float32)`

SetAuctionGrade sets AuctionGrade field to given value.

### HasAuctionGrade

`func (o *AeAssetUpdatedAmsDetail) HasAuctionGrade() bool`

HasAuctionGrade returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AeAssetUpdatedAmsDetail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AeAssetUpdatedAmsDetail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AeAssetUpdatedAmsDetail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AeAssetUpdatedAmsDetail) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *AeAssetUpdatedAmsDetail) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *AeAssetUpdatedAmsDetail) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetInitiator

`func (o *AeAssetUpdatedAmsDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAssetUpdatedAmsDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAssetUpdatedAmsDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAssetUpdatedAmsDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.

### GetValuations

`func (o *AeAssetUpdatedAmsDetail) GetValuations() []Valuation`

GetValuations returns the Valuations field if non-nil, zero value otherwise.

### GetValuationsOk

`func (o *AeAssetUpdatedAmsDetail) GetValuationsOk() (*[]Valuation, bool)`

GetValuationsOk returns a tuple with the Valuations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuations

`func (o *AeAssetUpdatedAmsDetail) SetValuations(v []Valuation)`

SetValuations sets Valuations field to given value.

### HasValuations

`func (o *AeAssetUpdatedAmsDetail) HasValuations() bool`

HasValuations returns a boolean if a field has been set.

### GetTitle

`func (o *AeAssetUpdatedAmsDetail) GetTitle() CommonTitle`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AeAssetUpdatedAmsDetail) GetTitleOk() (*CommonTitle, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AeAssetUpdatedAmsDetail) SetTitle(v CommonTitle)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AeAssetUpdatedAmsDetail) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetPresaleAnnouncements

`func (o *AeAssetUpdatedAmsDetail) GetPresaleAnnouncements() string`

GetPresaleAnnouncements returns the PresaleAnnouncements field if non-nil, zero value otherwise.

### GetPresaleAnnouncementsOk

`func (o *AeAssetUpdatedAmsDetail) GetPresaleAnnouncementsOk() (*string, bool)`

GetPresaleAnnouncementsOk returns a tuple with the PresaleAnnouncements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresaleAnnouncements

`func (o *AeAssetUpdatedAmsDetail) SetPresaleAnnouncements(v string)`

SetPresaleAnnouncements sets PresaleAnnouncements field to given value.

### HasPresaleAnnouncements

`func (o *AeAssetUpdatedAmsDetail) HasPresaleAnnouncements() bool`

HasPresaleAnnouncements returns a boolean if a field has been set.

### GetLocation

`func (o *AeAssetUpdatedAmsDetail) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AeAssetUpdatedAmsDetail) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AeAssetUpdatedAmsDetail) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *AeAssetUpdatedAmsDetail) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetAutoGrade

`func (o *AeAssetUpdatedAmsDetail) GetAutoGrade() float32`

GetAutoGrade returns the AutoGrade field if non-nil, zero value otherwise.

### GetAutoGradeOk

`func (o *AeAssetUpdatedAmsDetail) GetAutoGradeOk() (*float32, bool)`

GetAutoGradeOk returns a tuple with the AutoGrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoGrade

`func (o *AeAssetUpdatedAmsDetail) SetAutoGrade(v float32)`

SetAutoGrade sets AutoGrade field to given value.

### HasAutoGrade

`func (o *AeAssetUpdatedAmsDetail) HasAutoGrade() bool`

HasAutoGrade returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


