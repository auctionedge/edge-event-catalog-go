/*
Edge Event Schemas

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package events

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the AeAdvisoryAssetAmsDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AeAdvisoryAssetAmsDetail{}

// AeAdvisoryAssetAmsDetail Advisory Asset event for AMS assets
type AeAdvisoryAssetAmsDetail struct {
	// Auction Edge unique identifier for an auction.
	AuctionId string `json:"auction-id"`
	// Source's unique identifier for asset
	Id string `json:"id"`
	// The Vehicle Identification Number(VIN) of the asset.
	Vin string `json:"vin"`
	// The stock number of the asset.
	Stock string `json:"stock"`
	Seller CommonAmsAccountPointer `json:"seller"`
	CheckIn CommonAssetCheckIn `json:"check-in"`
	// The year the vehicle was manufactured.
	Year *int32 `json:"year,omitempty"`
	// The manufacturer of the vehicle.
	Make *string `json:"make,omitempty"`
	// The model of the vehicle.
	Model *string `json:"model,omitempty"`
	// The trim of the vehicle.
	Trimline *string `json:"trimline,omitempty"`
	// The body style of the vehicle
	BodyStyle *string `json:"body-style,omitempty"`
	Deal *Deal `json:"deal,omitempty"`
	// The lowest amount the seller is willing to take for this asset.
	FloorAmountUsd *float32 `json:"floor-amount-usd,omitempty"`
	Interior *CommonVehicleInterior `json:"interior,omitempty"`
	Exterior *CommonVehicleExterior `json:"exterior,omitempty"`
	// The vehicle's odometer reading
	CurrentOdometerReading *float32 `json:"current-odometer-reading,omitempty"`
	Engine *CommonVehicleEngine `json:"engine,omitempty"`
	// The type of transmission the asset has.
	TransmissionType *string `json:"transmission-type,omitempty"`
	Drivetrain *CommonVehicleDrivetrain `json:"drivetrain,omitempty"`
	Title *CommonTitle `json:"title,omitempty"`
	// A value rating the over all condition of the vehicle typically a number between 0-5 (5 being the best).
	AuctionGrade *float32 `json:"auction-grade,omitempty"`
	// The date and time at which the asset was last updated at.
	UpdatedAt NullableTime `json:"updated-at,omitempty"`
	Initiator *CommonInitiator `json:"initiator,omitempty"`
}

type _AeAdvisoryAssetAmsDetail AeAdvisoryAssetAmsDetail

// NewAeAdvisoryAssetAmsDetail instantiates a new AeAdvisoryAssetAmsDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAeAdvisoryAssetAmsDetail(auctionId string, id string, vin string, stock string, seller CommonAmsAccountPointer, checkIn CommonAssetCheckIn) *AeAdvisoryAssetAmsDetail {
	this := AeAdvisoryAssetAmsDetail{}
	this.AuctionId = auctionId
	this.Id = id
	this.Vin = vin
	this.Stock = stock
	this.Seller = seller
	this.CheckIn = checkIn
	return &this
}

// NewAeAdvisoryAssetAmsDetailWithDefaults instantiates a new AeAdvisoryAssetAmsDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAeAdvisoryAssetAmsDetailWithDefaults() *AeAdvisoryAssetAmsDetail {
	this := AeAdvisoryAssetAmsDetail{}
	return &this
}

// GetAuctionId returns the AuctionId field value
func (o *AeAdvisoryAssetAmsDetail) GetAuctionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuctionId
}

// GetAuctionIdOk returns a tuple with the AuctionId field value
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAssetAmsDetail) GetAuctionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuctionId, true
}

// SetAuctionId sets field value
func (o *AeAdvisoryAssetAmsDetail) SetAuctionId(v string) {
	o.AuctionId = v
}

// GetId returns the Id field value
func (o *AeAdvisoryAssetAmsDetail) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAssetAmsDetail) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AeAdvisoryAssetAmsDetail) SetId(v string) {
	o.Id = v
}

// GetVin returns the Vin field value
func (o *AeAdvisoryAssetAmsDetail) GetVin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vin
}

// GetVinOk returns a tuple with the Vin field value
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAssetAmsDetail) GetVinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vin, true
}

// SetVin sets field value
func (o *AeAdvisoryAssetAmsDetail) SetVin(v string) {
	o.Vin = v
}

// GetStock returns the Stock field value
func (o *AeAdvisoryAssetAmsDetail) GetStock() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Stock
}

// GetStockOk returns a tuple with the Stock field value
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAssetAmsDetail) GetStockOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stock, true
}

// SetStock sets field value
func (o *AeAdvisoryAssetAmsDetail) SetStock(v string) {
	o.Stock = v
}

// GetSeller returns the Seller field value
func (o *AeAdvisoryAssetAmsDetail) GetSeller() CommonAmsAccountPointer {
	if o == nil {
		var ret CommonAmsAccountPointer
		return ret
	}

	return o.Seller
}

// GetSellerOk returns a tuple with the Seller field value
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAssetAmsDetail) GetSellerOk() (*CommonAmsAccountPointer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Seller, true
}

// SetSeller sets field value
func (o *AeAdvisoryAssetAmsDetail) SetSeller(v CommonAmsAccountPointer) {
	o.Seller = v
}

// GetCheckIn returns the CheckIn field value
func (o *AeAdvisoryAssetAmsDetail) GetCheckIn() CommonAssetCheckIn {
	if o == nil {
		var ret CommonAssetCheckIn
		return ret
	}

	return o.CheckIn
}

// GetCheckInOk returns a tuple with the CheckIn field value
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAssetAmsDetail) GetCheckInOk() (*CommonAssetCheckIn, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CheckIn, true
}

// SetCheckIn sets field value
func (o *AeAdvisoryAssetAmsDetail) SetCheckIn(v CommonAssetCheckIn) {
	o.CheckIn = v
}

// GetYear returns the Year field value if set, zero value otherwise.
func (o *AeAdvisoryAssetAmsDetail) GetYear() int32 {
	if o == nil || IsNil(o.Year) {
		var ret int32
		return ret
	}
	return *o.Year
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAssetAmsDetail) GetYearOk() (*int32, bool) {
	if o == nil || IsNil(o.Year) {
		return nil, false
	}
	return o.Year, true
}

// HasYear returns a boolean if a field has been set.
func (o *AeAdvisoryAssetAmsDetail) HasYear() bool {
	if o != nil && !IsNil(o.Year) {
		return true
	}

	return false
}

// SetYear gets a reference to the given int32 and assigns it to the Year field.
func (o *AeAdvisoryAssetAmsDetail) SetYear(v int32) {
	o.Year = &v
}

// GetMake returns the Make field value if set, zero value otherwise.
func (o *AeAdvisoryAssetAmsDetail) GetMake() string {
	if o == nil || IsNil(o.Make) {
		var ret string
		return ret
	}
	return *o.Make
}

// GetMakeOk returns a tuple with the Make field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAssetAmsDetail) GetMakeOk() (*string, bool) {
	if o == nil || IsNil(o.Make) {
		return nil, false
	}
	return o.Make, true
}

// HasMake returns a boolean if a field has been set.
func (o *AeAdvisoryAssetAmsDetail) HasMake() bool {
	if o != nil && !IsNil(o.Make) {
		return true
	}

	return false
}

// SetMake gets a reference to the given string and assigns it to the Make field.
func (o *AeAdvisoryAssetAmsDetail) SetMake(v string) {
	o.Make = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *AeAdvisoryAssetAmsDetail) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAssetAmsDetail) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *AeAdvisoryAssetAmsDetail) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *AeAdvisoryAssetAmsDetail) SetModel(v string) {
	o.Model = &v
}

// GetTrimline returns the Trimline field value if set, zero value otherwise.
func (o *AeAdvisoryAssetAmsDetail) GetTrimline() string {
	if o == nil || IsNil(o.Trimline) {
		var ret string
		return ret
	}
	return *o.Trimline
}

// GetTrimlineOk returns a tuple with the Trimline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAssetAmsDetail) GetTrimlineOk() (*string, bool) {
	if o == nil || IsNil(o.Trimline) {
		return nil, false
	}
	return o.Trimline, true
}

// HasTrimline returns a boolean if a field has been set.
func (o *AeAdvisoryAssetAmsDetail) HasTrimline() bool {
	if o != nil && !IsNil(o.Trimline) {
		return true
	}

	return false
}

// SetTrimline gets a reference to the given string and assigns it to the Trimline field.
func (o *AeAdvisoryAssetAmsDetail) SetTrimline(v string) {
	o.Trimline = &v
}

// GetBodyStyle returns the BodyStyle field value if set, zero value otherwise.
func (o *AeAdvisoryAssetAmsDetail) GetBodyStyle() string {
	if o == nil || IsNil(o.BodyStyle) {
		var ret string
		return ret
	}
	return *o.BodyStyle
}

// GetBodyStyleOk returns a tuple with the BodyStyle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAssetAmsDetail) GetBodyStyleOk() (*string, bool) {
	if o == nil || IsNil(o.BodyStyle) {
		return nil, false
	}
	return o.BodyStyle, true
}

// HasBodyStyle returns a boolean if a field has been set.
func (o *AeAdvisoryAssetAmsDetail) HasBodyStyle() bool {
	if o != nil && !IsNil(o.BodyStyle) {
		return true
	}

	return false
}

// SetBodyStyle gets a reference to the given string and assigns it to the BodyStyle field.
func (o *AeAdvisoryAssetAmsDetail) SetBodyStyle(v string) {
	o.BodyStyle = &v
}

// GetDeal returns the Deal field value if set, zero value otherwise.
func (o *AeAdvisoryAssetAmsDetail) GetDeal() Deal {
	if o == nil || IsNil(o.Deal) {
		var ret Deal
		return ret
	}
	return *o.Deal
}

// GetDealOk returns a tuple with the Deal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAssetAmsDetail) GetDealOk() (*Deal, bool) {
	if o == nil || IsNil(o.Deal) {
		return nil, false
	}
	return o.Deal, true
}

// HasDeal returns a boolean if a field has been set.
func (o *AeAdvisoryAssetAmsDetail) HasDeal() bool {
	if o != nil && !IsNil(o.Deal) {
		return true
	}

	return false
}

// SetDeal gets a reference to the given Deal and assigns it to the Deal field.
func (o *AeAdvisoryAssetAmsDetail) SetDeal(v Deal) {
	o.Deal = &v
}

// GetFloorAmountUsd returns the FloorAmountUsd field value if set, zero value otherwise.
func (o *AeAdvisoryAssetAmsDetail) GetFloorAmountUsd() float32 {
	if o == nil || IsNil(o.FloorAmountUsd) {
		var ret float32
		return ret
	}
	return *o.FloorAmountUsd
}

// GetFloorAmountUsdOk returns a tuple with the FloorAmountUsd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAssetAmsDetail) GetFloorAmountUsdOk() (*float32, bool) {
	if o == nil || IsNil(o.FloorAmountUsd) {
		return nil, false
	}
	return o.FloorAmountUsd, true
}

// HasFloorAmountUsd returns a boolean if a field has been set.
func (o *AeAdvisoryAssetAmsDetail) HasFloorAmountUsd() bool {
	if o != nil && !IsNil(o.FloorAmountUsd) {
		return true
	}

	return false
}

// SetFloorAmountUsd gets a reference to the given float32 and assigns it to the FloorAmountUsd field.
func (o *AeAdvisoryAssetAmsDetail) SetFloorAmountUsd(v float32) {
	o.FloorAmountUsd = &v
}

// GetInterior returns the Interior field value if set, zero value otherwise.
func (o *AeAdvisoryAssetAmsDetail) GetInterior() CommonVehicleInterior {
	if o == nil || IsNil(o.Interior) {
		var ret CommonVehicleInterior
		return ret
	}
	return *o.Interior
}

// GetInteriorOk returns a tuple with the Interior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAssetAmsDetail) GetInteriorOk() (*CommonVehicleInterior, bool) {
	if o == nil || IsNil(o.Interior) {
		return nil, false
	}
	return o.Interior, true
}

// HasInterior returns a boolean if a field has been set.
func (o *AeAdvisoryAssetAmsDetail) HasInterior() bool {
	if o != nil && !IsNil(o.Interior) {
		return true
	}

	return false
}

// SetInterior gets a reference to the given CommonVehicleInterior and assigns it to the Interior field.
func (o *AeAdvisoryAssetAmsDetail) SetInterior(v CommonVehicleInterior) {
	o.Interior = &v
}

// GetExterior returns the Exterior field value if set, zero value otherwise.
func (o *AeAdvisoryAssetAmsDetail) GetExterior() CommonVehicleExterior {
	if o == nil || IsNil(o.Exterior) {
		var ret CommonVehicleExterior
		return ret
	}
	return *o.Exterior
}

// GetExteriorOk returns a tuple with the Exterior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAssetAmsDetail) GetExteriorOk() (*CommonVehicleExterior, bool) {
	if o == nil || IsNil(o.Exterior) {
		return nil, false
	}
	return o.Exterior, true
}

// HasExterior returns a boolean if a field has been set.
func (o *AeAdvisoryAssetAmsDetail) HasExterior() bool {
	if o != nil && !IsNil(o.Exterior) {
		return true
	}

	return false
}

// SetExterior gets a reference to the given CommonVehicleExterior and assigns it to the Exterior field.
func (o *AeAdvisoryAssetAmsDetail) SetExterior(v CommonVehicleExterior) {
	o.Exterior = &v
}

// GetCurrentOdometerReading returns the CurrentOdometerReading field value if set, zero value otherwise.
func (o *AeAdvisoryAssetAmsDetail) GetCurrentOdometerReading() float32 {
	if o == nil || IsNil(o.CurrentOdometerReading) {
		var ret float32
		return ret
	}
	return *o.CurrentOdometerReading
}

// GetCurrentOdometerReadingOk returns a tuple with the CurrentOdometerReading field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAssetAmsDetail) GetCurrentOdometerReadingOk() (*float32, bool) {
	if o == nil || IsNil(o.CurrentOdometerReading) {
		return nil, false
	}
	return o.CurrentOdometerReading, true
}

// HasCurrentOdometerReading returns a boolean if a field has been set.
func (o *AeAdvisoryAssetAmsDetail) HasCurrentOdometerReading() bool {
	if o != nil && !IsNil(o.CurrentOdometerReading) {
		return true
	}

	return false
}

// SetCurrentOdometerReading gets a reference to the given float32 and assigns it to the CurrentOdometerReading field.
func (o *AeAdvisoryAssetAmsDetail) SetCurrentOdometerReading(v float32) {
	o.CurrentOdometerReading = &v
}

// GetEngine returns the Engine field value if set, zero value otherwise.
func (o *AeAdvisoryAssetAmsDetail) GetEngine() CommonVehicleEngine {
	if o == nil || IsNil(o.Engine) {
		var ret CommonVehicleEngine
		return ret
	}
	return *o.Engine
}

// GetEngineOk returns a tuple with the Engine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAssetAmsDetail) GetEngineOk() (*CommonVehicleEngine, bool) {
	if o == nil || IsNil(o.Engine) {
		return nil, false
	}
	return o.Engine, true
}

// HasEngine returns a boolean if a field has been set.
func (o *AeAdvisoryAssetAmsDetail) HasEngine() bool {
	if o != nil && !IsNil(o.Engine) {
		return true
	}

	return false
}

// SetEngine gets a reference to the given CommonVehicleEngine and assigns it to the Engine field.
func (o *AeAdvisoryAssetAmsDetail) SetEngine(v CommonVehicleEngine) {
	o.Engine = &v
}

// GetTransmissionType returns the TransmissionType field value if set, zero value otherwise.
func (o *AeAdvisoryAssetAmsDetail) GetTransmissionType() string {
	if o == nil || IsNil(o.TransmissionType) {
		var ret string
		return ret
	}
	return *o.TransmissionType
}

// GetTransmissionTypeOk returns a tuple with the TransmissionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAssetAmsDetail) GetTransmissionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TransmissionType) {
		return nil, false
	}
	return o.TransmissionType, true
}

// HasTransmissionType returns a boolean if a field has been set.
func (o *AeAdvisoryAssetAmsDetail) HasTransmissionType() bool {
	if o != nil && !IsNil(o.TransmissionType) {
		return true
	}

	return false
}

// SetTransmissionType gets a reference to the given string and assigns it to the TransmissionType field.
func (o *AeAdvisoryAssetAmsDetail) SetTransmissionType(v string) {
	o.TransmissionType = &v
}

// GetDrivetrain returns the Drivetrain field value if set, zero value otherwise.
func (o *AeAdvisoryAssetAmsDetail) GetDrivetrain() CommonVehicleDrivetrain {
	if o == nil || IsNil(o.Drivetrain) {
		var ret CommonVehicleDrivetrain
		return ret
	}
	return *o.Drivetrain
}

// GetDrivetrainOk returns a tuple with the Drivetrain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAssetAmsDetail) GetDrivetrainOk() (*CommonVehicleDrivetrain, bool) {
	if o == nil || IsNil(o.Drivetrain) {
		return nil, false
	}
	return o.Drivetrain, true
}

// HasDrivetrain returns a boolean if a field has been set.
func (o *AeAdvisoryAssetAmsDetail) HasDrivetrain() bool {
	if o != nil && !IsNil(o.Drivetrain) {
		return true
	}

	return false
}

// SetDrivetrain gets a reference to the given CommonVehicleDrivetrain and assigns it to the Drivetrain field.
func (o *AeAdvisoryAssetAmsDetail) SetDrivetrain(v CommonVehicleDrivetrain) {
	o.Drivetrain = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AeAdvisoryAssetAmsDetail) GetTitle() CommonTitle {
	if o == nil || IsNil(o.Title) {
		var ret CommonTitle
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAssetAmsDetail) GetTitleOk() (*CommonTitle, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AeAdvisoryAssetAmsDetail) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given CommonTitle and assigns it to the Title field.
func (o *AeAdvisoryAssetAmsDetail) SetTitle(v CommonTitle) {
	o.Title = &v
}

// GetAuctionGrade returns the AuctionGrade field value if set, zero value otherwise.
func (o *AeAdvisoryAssetAmsDetail) GetAuctionGrade() float32 {
	if o == nil || IsNil(o.AuctionGrade) {
		var ret float32
		return ret
	}
	return *o.AuctionGrade
}

// GetAuctionGradeOk returns a tuple with the AuctionGrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAssetAmsDetail) GetAuctionGradeOk() (*float32, bool) {
	if o == nil || IsNil(o.AuctionGrade) {
		return nil, false
	}
	return o.AuctionGrade, true
}

// HasAuctionGrade returns a boolean if a field has been set.
func (o *AeAdvisoryAssetAmsDetail) HasAuctionGrade() bool {
	if o != nil && !IsNil(o.AuctionGrade) {
		return true
	}

	return false
}

// SetAuctionGrade gets a reference to the given float32 and assigns it to the AuctionGrade field.
func (o *AeAdvisoryAssetAmsDetail) SetAuctionGrade(v float32) {
	o.AuctionGrade = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AeAdvisoryAssetAmsDetail) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AeAdvisoryAssetAmsDetail) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AeAdvisoryAssetAmsDetail) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *AeAdvisoryAssetAmsDetail) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *AeAdvisoryAssetAmsDetail) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *AeAdvisoryAssetAmsDetail) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetInitiator returns the Initiator field value if set, zero value otherwise.
func (o *AeAdvisoryAssetAmsDetail) GetInitiator() CommonInitiator {
	if o == nil || IsNil(o.Initiator) {
		var ret CommonInitiator
		return ret
	}
	return *o.Initiator
}

// GetInitiatorOk returns a tuple with the Initiator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAssetAmsDetail) GetInitiatorOk() (*CommonInitiator, bool) {
	if o == nil || IsNil(o.Initiator) {
		return nil, false
	}
	return o.Initiator, true
}

// HasInitiator returns a boolean if a field has been set.
func (o *AeAdvisoryAssetAmsDetail) HasInitiator() bool {
	if o != nil && !IsNil(o.Initiator) {
		return true
	}

	return false
}

// SetInitiator gets a reference to the given CommonInitiator and assigns it to the Initiator field.
func (o *AeAdvisoryAssetAmsDetail) SetInitiator(v CommonInitiator) {
	o.Initiator = &v
}

func (o AeAdvisoryAssetAmsDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AeAdvisoryAssetAmsDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["auction-id"] = o.AuctionId
	toSerialize["id"] = o.Id
	toSerialize["vin"] = o.Vin
	toSerialize["stock"] = o.Stock
	toSerialize["seller"] = o.Seller
	toSerialize["check-in"] = o.CheckIn
	if !IsNil(o.Year) {
		toSerialize["year"] = o.Year
	}
	if !IsNil(o.Make) {
		toSerialize["make"] = o.Make
	}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !IsNil(o.Trimline) {
		toSerialize["trimline"] = o.Trimline
	}
	if !IsNil(o.BodyStyle) {
		toSerialize["body-style"] = o.BodyStyle
	}
	if !IsNil(o.Deal) {
		toSerialize["deal"] = o.Deal
	}
	if !IsNil(o.FloorAmountUsd) {
		toSerialize["floor-amount-usd"] = o.FloorAmountUsd
	}
	if !IsNil(o.Interior) {
		toSerialize["interior"] = o.Interior
	}
	if !IsNil(o.Exterior) {
		toSerialize["exterior"] = o.Exterior
	}
	if !IsNil(o.CurrentOdometerReading) {
		toSerialize["current-odometer-reading"] = o.CurrentOdometerReading
	}
	if !IsNil(o.Engine) {
		toSerialize["engine"] = o.Engine
	}
	if !IsNil(o.TransmissionType) {
		toSerialize["transmission-type"] = o.TransmissionType
	}
	if !IsNil(o.Drivetrain) {
		toSerialize["drivetrain"] = o.Drivetrain
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.AuctionGrade) {
		toSerialize["auction-grade"] = o.AuctionGrade
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated-at"] = o.UpdatedAt.Get()
	}
	if !IsNil(o.Initiator) {
		toSerialize["initiator"] = o.Initiator
	}
	return toSerialize, nil
}

func (o *AeAdvisoryAssetAmsDetail) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"auction-id",
		"id",
		"vin",
		"stock",
		"seller",
		"check-in",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAeAdvisoryAssetAmsDetail := _AeAdvisoryAssetAmsDetail{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAeAdvisoryAssetAmsDetail)

	if err != nil {
		return err
	}

	*o = AeAdvisoryAssetAmsDetail(varAeAdvisoryAssetAmsDetail)

	return err
}

type NullableAeAdvisoryAssetAmsDetail struct {
	value *AeAdvisoryAssetAmsDetail
	isSet bool
}

func (v NullableAeAdvisoryAssetAmsDetail) Get() *AeAdvisoryAssetAmsDetail {
	return v.value
}

func (v *NullableAeAdvisoryAssetAmsDetail) Set(val *AeAdvisoryAssetAmsDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableAeAdvisoryAssetAmsDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableAeAdvisoryAssetAmsDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAeAdvisoryAssetAmsDetail(val *AeAdvisoryAssetAmsDetail) *NullableAeAdvisoryAssetAmsDetail {
	return &NullableAeAdvisoryAssetAmsDetail{value: val, isSet: true}
}

func (v NullableAeAdvisoryAssetAmsDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAeAdvisoryAssetAmsDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


