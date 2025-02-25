# CommonVehicleEngine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cylinders** | Pointer to **float32** | The number of cylinders the engine has. | [optional] 
**CylinderConfiguration** | Pointer to **string** | How the cylinders are arranged in an engine. | [optional] 
**Displacement** | Pointer to **string** | the total amount of air and fuel that can be drawn into the cylinders at one time. | [optional] 
**FuelType** | Pointer to **string** | The type of fuel the engine consumes to operate. | [optional] 

## Methods

### NewCommonVehicleEngine

`func NewCommonVehicleEngine() *CommonVehicleEngine`

NewCommonVehicleEngine instantiates a new CommonVehicleEngine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonVehicleEngineWithDefaults

`func NewCommonVehicleEngineWithDefaults() *CommonVehicleEngine`

NewCommonVehicleEngineWithDefaults instantiates a new CommonVehicleEngine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCylinders

`func (o *CommonVehicleEngine) GetCylinders() float32`

GetCylinders returns the Cylinders field if non-nil, zero value otherwise.

### GetCylindersOk

`func (o *CommonVehicleEngine) GetCylindersOk() (*float32, bool)`

GetCylindersOk returns a tuple with the Cylinders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCylinders

`func (o *CommonVehicleEngine) SetCylinders(v float32)`

SetCylinders sets Cylinders field to given value.

### HasCylinders

`func (o *CommonVehicleEngine) HasCylinders() bool`

HasCylinders returns a boolean if a field has been set.

### GetCylinderConfiguration

`func (o *CommonVehicleEngine) GetCylinderConfiguration() string`

GetCylinderConfiguration returns the CylinderConfiguration field if non-nil, zero value otherwise.

### GetCylinderConfigurationOk

`func (o *CommonVehicleEngine) GetCylinderConfigurationOk() (*string, bool)`

GetCylinderConfigurationOk returns a tuple with the CylinderConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCylinderConfiguration

`func (o *CommonVehicleEngine) SetCylinderConfiguration(v string)`

SetCylinderConfiguration sets CylinderConfiguration field to given value.

### HasCylinderConfiguration

`func (o *CommonVehicleEngine) HasCylinderConfiguration() bool`

HasCylinderConfiguration returns a boolean if a field has been set.

### GetDisplacement

`func (o *CommonVehicleEngine) GetDisplacement() string`

GetDisplacement returns the Displacement field if non-nil, zero value otherwise.

### GetDisplacementOk

`func (o *CommonVehicleEngine) GetDisplacementOk() (*string, bool)`

GetDisplacementOk returns a tuple with the Displacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplacement

`func (o *CommonVehicleEngine) SetDisplacement(v string)`

SetDisplacement sets Displacement field to given value.

### HasDisplacement

`func (o *CommonVehicleEngine) HasDisplacement() bool`

HasDisplacement returns a boolean if a field has been set.

### GetFuelType

`func (o *CommonVehicleEngine) GetFuelType() string`

GetFuelType returns the FuelType field if non-nil, zero value otherwise.

### GetFuelTypeOk

`func (o *CommonVehicleEngine) GetFuelTypeOk() (*string, bool)`

GetFuelTypeOk returns a tuple with the FuelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelType

`func (o *CommonVehicleEngine) SetFuelType(v string)`

SetFuelType sets FuelType field to given value.

### HasFuelType

`func (o *CommonVehicleEngine) HasFuelType() bool`

HasFuelType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


