# AeAssetGatepassBuyerReleasableDetailAsset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Year** | **int32** | The year the vehicle was manufactured. | 
**Make** | Pointer to **NullableString** | The manufacturer make of the asset | [optional] 
**Model** | Pointer to **NullableString** | The model of the asset | [optional] 
**Trimline** | Pointer to **NullableString** | The trim of the asset | [optional] 
**Exterior** | Pointer to [**CommonVehicleExterior**](CommonVehicleExterior.md) |  | [optional] 
**CurrentOdometerReading** | Pointer to **NullableFloat32** | The vehicle&#39;s odometer reading | [optional] 
**LotLocation** | Pointer to **NullableString** | The name of the auction lot assigned to the asset | [optional] 
**Id** | **string** | Source&#39;s unique identifier for asset | 
**Stock** | **string** | The stock number of the asset. | 
**Vin** | **string** | The Vehicle Identification Number(VIN) of the asset. | 

## Methods

### NewAeAssetGatepassBuyerReleasableDetailAsset

`func NewAeAssetGatepassBuyerReleasableDetailAsset(year int32, id string, stock string, vin string, ) *AeAssetGatepassBuyerReleasableDetailAsset`

NewAeAssetGatepassBuyerReleasableDetailAsset instantiates a new AeAssetGatepassBuyerReleasableDetailAsset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetGatepassBuyerReleasableDetailAssetWithDefaults

`func NewAeAssetGatepassBuyerReleasableDetailAssetWithDefaults() *AeAssetGatepassBuyerReleasableDetailAsset`

NewAeAssetGatepassBuyerReleasableDetailAssetWithDefaults instantiates a new AeAssetGatepassBuyerReleasableDetailAsset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetYear

`func (o *AeAssetGatepassBuyerReleasableDetailAsset) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *AeAssetGatepassBuyerReleasableDetailAsset) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *AeAssetGatepassBuyerReleasableDetailAsset) SetYear(v int32)`

SetYear sets Year field to given value.


### GetMake

`func (o *AeAssetGatepassBuyerReleasableDetailAsset) GetMake() string`

GetMake returns the Make field if non-nil, zero value otherwise.

### GetMakeOk

`func (o *AeAssetGatepassBuyerReleasableDetailAsset) GetMakeOk() (*string, bool)`

GetMakeOk returns a tuple with the Make field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMake

`func (o *AeAssetGatepassBuyerReleasableDetailAsset) SetMake(v string)`

SetMake sets Make field to given value.

### HasMake

`func (o *AeAssetGatepassBuyerReleasableDetailAsset) HasMake() bool`

HasMake returns a boolean if a field has been set.

### SetMakeNil

`func (o *AeAssetGatepassBuyerReleasableDetailAsset) SetMakeNil(b bool)`

 SetMakeNil sets the value for Make to be an explicit nil

### UnsetMake
`func (o *AeAssetGatepassBuyerReleasableDetailAsset) UnsetMake()`

UnsetMake ensures that no value is present for Make, not even an explicit nil
### GetModel

`func (o *AeAssetGatepassBuyerReleasableDetailAsset) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AeAssetGatepassBuyerReleasableDetailAsset) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AeAssetGatepassBuyerReleasableDetailAsset) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AeAssetGatepassBuyerReleasableDetailAsset) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *AeAssetGatepassBuyerReleasableDetailAsset) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *AeAssetGatepassBuyerReleasableDetailAsset) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetTrimline

`func (o *AeAssetGatepassBuyerReleasableDetailAsset) GetTrimline() string`

GetTrimline returns the Trimline field if non-nil, zero value otherwise.

### GetTrimlineOk

`func (o *AeAssetGatepassBuyerReleasableDetailAsset) GetTrimlineOk() (*string, bool)`

GetTrimlineOk returns a tuple with the Trimline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrimline

`func (o *AeAssetGatepassBuyerReleasableDetailAsset) SetTrimline(v string)`

SetTrimline sets Trimline field to given value.

### HasTrimline

`func (o *AeAssetGatepassBuyerReleasableDetailAsset) HasTrimline() bool`

HasTrimline returns a boolean if a field has been set.

### SetTrimlineNil

`func (o *AeAssetGatepassBuyerReleasableDetailAsset) SetTrimlineNil(b bool)`

 SetTrimlineNil sets the value for Trimline to be an explicit nil

### UnsetTrimline
`func (o *AeAssetGatepassBuyerReleasableDetailAsset) UnsetTrimline()`

UnsetTrimline ensures that no value is present for Trimline, not even an explicit nil
### GetExterior

`func (o *AeAssetGatepassBuyerReleasableDetailAsset) GetExterior() CommonVehicleExterior`

GetExterior returns the Exterior field if non-nil, zero value otherwise.

### GetExteriorOk

`func (o *AeAssetGatepassBuyerReleasableDetailAsset) GetExteriorOk() (*CommonVehicleExterior, bool)`

GetExteriorOk returns a tuple with the Exterior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExterior

`func (o *AeAssetGatepassBuyerReleasableDetailAsset) SetExterior(v CommonVehicleExterior)`

SetExterior sets Exterior field to given value.

### HasExterior

`func (o *AeAssetGatepassBuyerReleasableDetailAsset) HasExterior() bool`

HasExterior returns a boolean if a field has been set.

### GetCurrentOdometerReading

`func (o *AeAssetGatepassBuyerReleasableDetailAsset) GetCurrentOdometerReading() float32`

GetCurrentOdometerReading returns the CurrentOdometerReading field if non-nil, zero value otherwise.

### GetCurrentOdometerReadingOk

`func (o *AeAssetGatepassBuyerReleasableDetailAsset) GetCurrentOdometerReadingOk() (*float32, bool)`

GetCurrentOdometerReadingOk returns a tuple with the CurrentOdometerReading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentOdometerReading

`func (o *AeAssetGatepassBuyerReleasableDetailAsset) SetCurrentOdometerReading(v float32)`

SetCurrentOdometerReading sets CurrentOdometerReading field to given value.

### HasCurrentOdometerReading

`func (o *AeAssetGatepassBuyerReleasableDetailAsset) HasCurrentOdometerReading() bool`

HasCurrentOdometerReading returns a boolean if a field has been set.

### SetCurrentOdometerReadingNil

`func (o *AeAssetGatepassBuyerReleasableDetailAsset) SetCurrentOdometerReadingNil(b bool)`

 SetCurrentOdometerReadingNil sets the value for CurrentOdometerReading to be an explicit nil

### UnsetCurrentOdometerReading
`func (o *AeAssetGatepassBuyerReleasableDetailAsset) UnsetCurrentOdometerReading()`

UnsetCurrentOdometerReading ensures that no value is present for CurrentOdometerReading, not even an explicit nil
### GetLotLocation

`func (o *AeAssetGatepassBuyerReleasableDetailAsset) GetLotLocation() string`

GetLotLocation returns the LotLocation field if non-nil, zero value otherwise.

### GetLotLocationOk

`func (o *AeAssetGatepassBuyerReleasableDetailAsset) GetLotLocationOk() (*string, bool)`

GetLotLocationOk returns a tuple with the LotLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLotLocation

`func (o *AeAssetGatepassBuyerReleasableDetailAsset) SetLotLocation(v string)`

SetLotLocation sets LotLocation field to given value.

### HasLotLocation

`func (o *AeAssetGatepassBuyerReleasableDetailAsset) HasLotLocation() bool`

HasLotLocation returns a boolean if a field has been set.

### SetLotLocationNil

`func (o *AeAssetGatepassBuyerReleasableDetailAsset) SetLotLocationNil(b bool)`

 SetLotLocationNil sets the value for LotLocation to be an explicit nil

### UnsetLotLocation
`func (o *AeAssetGatepassBuyerReleasableDetailAsset) UnsetLotLocation()`

UnsetLotLocation ensures that no value is present for LotLocation, not even an explicit nil
### GetId

`func (o *AeAssetGatepassBuyerReleasableDetailAsset) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AeAssetGatepassBuyerReleasableDetailAsset) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AeAssetGatepassBuyerReleasableDetailAsset) SetId(v string)`

SetId sets Id field to given value.


### GetStock

`func (o *AeAssetGatepassBuyerReleasableDetailAsset) GetStock() string`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *AeAssetGatepassBuyerReleasableDetailAsset) GetStockOk() (*string, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *AeAssetGatepassBuyerReleasableDetailAsset) SetStock(v string)`

SetStock sets Stock field to given value.


### GetVin

`func (o *AeAssetGatepassBuyerReleasableDetailAsset) GetVin() string`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *AeAssetGatepassBuyerReleasableDetailAsset) GetVinOk() (*string, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *AeAssetGatepassBuyerReleasableDetailAsset) SetVin(v string)`

SetVin sets Vin field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


