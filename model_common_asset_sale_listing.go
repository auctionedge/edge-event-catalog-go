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
	"fmt"
)

// checks if the CommonAssetSaleListing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonAssetSaleListing{}

// CommonAssetSaleListing Common object of a sale listing for an asset
type CommonAssetSaleListing struct {
	// AMS unique identifier for this sale listing of the asset
	Id string `json:"id"`
	// Status of the sale listing
	Status string `json:"status"`
	Venue CommonVenue `json:"venue"`
	// The time that this sale listing for the asset begins
	SaleStartAt time.Time `json:"sale-start-at"`
	// The time that this sale listing for the asset ends
	SaleEndAt *time.Time `json:"sale-end-at,omitempty"`
	// The representative id an AMS account
	SellerRep string `json:"seller-rep"`
	// The lane the asset was in when deal made.
	Lane *string `json:"lane,omitempty"`
	// The lot the asset was in when deal made.
	Lot *string `json:"lot,omitempty"`
	// Auction announcements on the asset
	AnnouncementsText []string `json:"announcements-text"`
	// The current lights on the asset when the sale listing was made.
	AnnouncementLights string `json:"announcement-lights"`
	FloorAmount *CommonCurrency `json:"floor-amount,omitempty"`
	StartBid *CommonCurrency `json:"start-bid,omitempty"`
	HighBid *CommonCurrency `json:"high-bid,omitempty"`
	BidIncrement *CommonCurrency `json:"bid-increment,omitempty"`
	// The URL for the main image of the asset to be posted on the listing
	MainImageUrl *string `json:"main-image-url,omitempty"`
	// total number of photos attached to the asset
	PhotoCount *string `json:"photo-count,omitempty"`
	// timestamp of when photo count was last changed
	PhotoCountUpdatedAt *time.Time `json:"photo-count-updated-at,omitempty"`
	BuyNow *CommonAssetBuyNow `json:"buy-now,omitempty"`
	// The time in ISO-8601 formatted at which the record was generated
	UpdatedAt time.Time `json:"updated-at"`
	AdditionalProperties map[string]interface{}
}

type _CommonAssetSaleListing CommonAssetSaleListing

// NewCommonAssetSaleListing instantiates a new CommonAssetSaleListing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonAssetSaleListing(id string, status string, venue CommonVenue, saleStartAt time.Time, sellerRep string, announcementsText []string, announcementLights string, updatedAt time.Time) *CommonAssetSaleListing {
	this := CommonAssetSaleListing{}
	this.Id = id
	this.Status = status
	this.Venue = venue
	this.SaleStartAt = saleStartAt
	this.SellerRep = sellerRep
	this.AnnouncementsText = announcementsText
	this.AnnouncementLights = announcementLights
	this.UpdatedAt = updatedAt
	return &this
}

// NewCommonAssetSaleListingWithDefaults instantiates a new CommonAssetSaleListing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonAssetSaleListingWithDefaults() *CommonAssetSaleListing {
	this := CommonAssetSaleListing{}
	return &this
}

// GetId returns the Id field value
func (o *CommonAssetSaleListing) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CommonAssetSaleListing) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CommonAssetSaleListing) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *CommonAssetSaleListing) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CommonAssetSaleListing) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CommonAssetSaleListing) SetStatus(v string) {
	o.Status = v
}

// GetVenue returns the Venue field value
func (o *CommonAssetSaleListing) GetVenue() CommonVenue {
	if o == nil {
		var ret CommonVenue
		return ret
	}

	return o.Venue
}

// GetVenueOk returns a tuple with the Venue field value
// and a boolean to check if the value has been set.
func (o *CommonAssetSaleListing) GetVenueOk() (*CommonVenue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Venue, true
}

// SetVenue sets field value
func (o *CommonAssetSaleListing) SetVenue(v CommonVenue) {
	o.Venue = v
}

// GetSaleStartAt returns the SaleStartAt field value
func (o *CommonAssetSaleListing) GetSaleStartAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.SaleStartAt
}

// GetSaleStartAtOk returns a tuple with the SaleStartAt field value
// and a boolean to check if the value has been set.
func (o *CommonAssetSaleListing) GetSaleStartAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SaleStartAt, true
}

// SetSaleStartAt sets field value
func (o *CommonAssetSaleListing) SetSaleStartAt(v time.Time) {
	o.SaleStartAt = v
}

// GetSaleEndAt returns the SaleEndAt field value if set, zero value otherwise.
func (o *CommonAssetSaleListing) GetSaleEndAt() time.Time {
	if o == nil || IsNil(o.SaleEndAt) {
		var ret time.Time
		return ret
	}
	return *o.SaleEndAt
}

// GetSaleEndAtOk returns a tuple with the SaleEndAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAssetSaleListing) GetSaleEndAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.SaleEndAt) {
		return nil, false
	}
	return o.SaleEndAt, true
}

// HasSaleEndAt returns a boolean if a field has been set.
func (o *CommonAssetSaleListing) HasSaleEndAt() bool {
	if o != nil && !IsNil(o.SaleEndAt) {
		return true
	}

	return false
}

// SetSaleEndAt gets a reference to the given time.Time and assigns it to the SaleEndAt field.
func (o *CommonAssetSaleListing) SetSaleEndAt(v time.Time) {
	o.SaleEndAt = &v
}

// GetSellerRep returns the SellerRep field value
func (o *CommonAssetSaleListing) GetSellerRep() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SellerRep
}

// GetSellerRepOk returns a tuple with the SellerRep field value
// and a boolean to check if the value has been set.
func (o *CommonAssetSaleListing) GetSellerRepOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SellerRep, true
}

// SetSellerRep sets field value
func (o *CommonAssetSaleListing) SetSellerRep(v string) {
	o.SellerRep = v
}

// GetLane returns the Lane field value if set, zero value otherwise.
func (o *CommonAssetSaleListing) GetLane() string {
	if o == nil || IsNil(o.Lane) {
		var ret string
		return ret
	}
	return *o.Lane
}

// GetLaneOk returns a tuple with the Lane field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAssetSaleListing) GetLaneOk() (*string, bool) {
	if o == nil || IsNil(o.Lane) {
		return nil, false
	}
	return o.Lane, true
}

// HasLane returns a boolean if a field has been set.
func (o *CommonAssetSaleListing) HasLane() bool {
	if o != nil && !IsNil(o.Lane) {
		return true
	}

	return false
}

// SetLane gets a reference to the given string and assigns it to the Lane field.
func (o *CommonAssetSaleListing) SetLane(v string) {
	o.Lane = &v
}

// GetLot returns the Lot field value if set, zero value otherwise.
func (o *CommonAssetSaleListing) GetLot() string {
	if o == nil || IsNil(o.Lot) {
		var ret string
		return ret
	}
	return *o.Lot
}

// GetLotOk returns a tuple with the Lot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAssetSaleListing) GetLotOk() (*string, bool) {
	if o == nil || IsNil(o.Lot) {
		return nil, false
	}
	return o.Lot, true
}

// HasLot returns a boolean if a field has been set.
func (o *CommonAssetSaleListing) HasLot() bool {
	if o != nil && !IsNil(o.Lot) {
		return true
	}

	return false
}

// SetLot gets a reference to the given string and assigns it to the Lot field.
func (o *CommonAssetSaleListing) SetLot(v string) {
	o.Lot = &v
}

// GetAnnouncementsText returns the AnnouncementsText field value
func (o *CommonAssetSaleListing) GetAnnouncementsText() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AnnouncementsText
}

// GetAnnouncementsTextOk returns a tuple with the AnnouncementsText field value
// and a boolean to check if the value has been set.
func (o *CommonAssetSaleListing) GetAnnouncementsTextOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AnnouncementsText, true
}

// SetAnnouncementsText sets field value
func (o *CommonAssetSaleListing) SetAnnouncementsText(v []string) {
	o.AnnouncementsText = v
}

// GetAnnouncementLights returns the AnnouncementLights field value
func (o *CommonAssetSaleListing) GetAnnouncementLights() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AnnouncementLights
}

// GetAnnouncementLightsOk returns a tuple with the AnnouncementLights field value
// and a boolean to check if the value has been set.
func (o *CommonAssetSaleListing) GetAnnouncementLightsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnnouncementLights, true
}

// SetAnnouncementLights sets field value
func (o *CommonAssetSaleListing) SetAnnouncementLights(v string) {
	o.AnnouncementLights = v
}

// GetFloorAmount returns the FloorAmount field value if set, zero value otherwise.
func (o *CommonAssetSaleListing) GetFloorAmount() CommonCurrency {
	if o == nil || IsNil(o.FloorAmount) {
		var ret CommonCurrency
		return ret
	}
	return *o.FloorAmount
}

// GetFloorAmountOk returns a tuple with the FloorAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAssetSaleListing) GetFloorAmountOk() (*CommonCurrency, bool) {
	if o == nil || IsNil(o.FloorAmount) {
		return nil, false
	}
	return o.FloorAmount, true
}

// HasFloorAmount returns a boolean if a field has been set.
func (o *CommonAssetSaleListing) HasFloorAmount() bool {
	if o != nil && !IsNil(o.FloorAmount) {
		return true
	}

	return false
}

// SetFloorAmount gets a reference to the given CommonCurrency and assigns it to the FloorAmount field.
func (o *CommonAssetSaleListing) SetFloorAmount(v CommonCurrency) {
	o.FloorAmount = &v
}

// GetStartBid returns the StartBid field value if set, zero value otherwise.
func (o *CommonAssetSaleListing) GetStartBid() CommonCurrency {
	if o == nil || IsNil(o.StartBid) {
		var ret CommonCurrency
		return ret
	}
	return *o.StartBid
}

// GetStartBidOk returns a tuple with the StartBid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAssetSaleListing) GetStartBidOk() (*CommonCurrency, bool) {
	if o == nil || IsNil(o.StartBid) {
		return nil, false
	}
	return o.StartBid, true
}

// HasStartBid returns a boolean if a field has been set.
func (o *CommonAssetSaleListing) HasStartBid() bool {
	if o != nil && !IsNil(o.StartBid) {
		return true
	}

	return false
}

// SetStartBid gets a reference to the given CommonCurrency and assigns it to the StartBid field.
func (o *CommonAssetSaleListing) SetStartBid(v CommonCurrency) {
	o.StartBid = &v
}

// GetHighBid returns the HighBid field value if set, zero value otherwise.
func (o *CommonAssetSaleListing) GetHighBid() CommonCurrency {
	if o == nil || IsNil(o.HighBid) {
		var ret CommonCurrency
		return ret
	}
	return *o.HighBid
}

// GetHighBidOk returns a tuple with the HighBid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAssetSaleListing) GetHighBidOk() (*CommonCurrency, bool) {
	if o == nil || IsNil(o.HighBid) {
		return nil, false
	}
	return o.HighBid, true
}

// HasHighBid returns a boolean if a field has been set.
func (o *CommonAssetSaleListing) HasHighBid() bool {
	if o != nil && !IsNil(o.HighBid) {
		return true
	}

	return false
}

// SetHighBid gets a reference to the given CommonCurrency and assigns it to the HighBid field.
func (o *CommonAssetSaleListing) SetHighBid(v CommonCurrency) {
	o.HighBid = &v
}

// GetBidIncrement returns the BidIncrement field value if set, zero value otherwise.
func (o *CommonAssetSaleListing) GetBidIncrement() CommonCurrency {
	if o == nil || IsNil(o.BidIncrement) {
		var ret CommonCurrency
		return ret
	}
	return *o.BidIncrement
}

// GetBidIncrementOk returns a tuple with the BidIncrement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAssetSaleListing) GetBidIncrementOk() (*CommonCurrency, bool) {
	if o == nil || IsNil(o.BidIncrement) {
		return nil, false
	}
	return o.BidIncrement, true
}

// HasBidIncrement returns a boolean if a field has been set.
func (o *CommonAssetSaleListing) HasBidIncrement() bool {
	if o != nil && !IsNil(o.BidIncrement) {
		return true
	}

	return false
}

// SetBidIncrement gets a reference to the given CommonCurrency and assigns it to the BidIncrement field.
func (o *CommonAssetSaleListing) SetBidIncrement(v CommonCurrency) {
	o.BidIncrement = &v
}

// GetMainImageUrl returns the MainImageUrl field value if set, zero value otherwise.
func (o *CommonAssetSaleListing) GetMainImageUrl() string {
	if o == nil || IsNil(o.MainImageUrl) {
		var ret string
		return ret
	}
	return *o.MainImageUrl
}

// GetMainImageUrlOk returns a tuple with the MainImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAssetSaleListing) GetMainImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.MainImageUrl) {
		return nil, false
	}
	return o.MainImageUrl, true
}

// HasMainImageUrl returns a boolean if a field has been set.
func (o *CommonAssetSaleListing) HasMainImageUrl() bool {
	if o != nil && !IsNil(o.MainImageUrl) {
		return true
	}

	return false
}

// SetMainImageUrl gets a reference to the given string and assigns it to the MainImageUrl field.
func (o *CommonAssetSaleListing) SetMainImageUrl(v string) {
	o.MainImageUrl = &v
}

// GetPhotoCount returns the PhotoCount field value if set, zero value otherwise.
func (o *CommonAssetSaleListing) GetPhotoCount() string {
	if o == nil || IsNil(o.PhotoCount) {
		var ret string
		return ret
	}
	return *o.PhotoCount
}

// GetPhotoCountOk returns a tuple with the PhotoCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAssetSaleListing) GetPhotoCountOk() (*string, bool) {
	if o == nil || IsNil(o.PhotoCount) {
		return nil, false
	}
	return o.PhotoCount, true
}

// HasPhotoCount returns a boolean if a field has been set.
func (o *CommonAssetSaleListing) HasPhotoCount() bool {
	if o != nil && !IsNil(o.PhotoCount) {
		return true
	}

	return false
}

// SetPhotoCount gets a reference to the given string and assigns it to the PhotoCount field.
func (o *CommonAssetSaleListing) SetPhotoCount(v string) {
	o.PhotoCount = &v
}

// GetPhotoCountUpdatedAt returns the PhotoCountUpdatedAt field value if set, zero value otherwise.
func (o *CommonAssetSaleListing) GetPhotoCountUpdatedAt() time.Time {
	if o == nil || IsNil(o.PhotoCountUpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.PhotoCountUpdatedAt
}

// GetPhotoCountUpdatedAtOk returns a tuple with the PhotoCountUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAssetSaleListing) GetPhotoCountUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PhotoCountUpdatedAt) {
		return nil, false
	}
	return o.PhotoCountUpdatedAt, true
}

// HasPhotoCountUpdatedAt returns a boolean if a field has been set.
func (o *CommonAssetSaleListing) HasPhotoCountUpdatedAt() bool {
	if o != nil && !IsNil(o.PhotoCountUpdatedAt) {
		return true
	}

	return false
}

// SetPhotoCountUpdatedAt gets a reference to the given time.Time and assigns it to the PhotoCountUpdatedAt field.
func (o *CommonAssetSaleListing) SetPhotoCountUpdatedAt(v time.Time) {
	o.PhotoCountUpdatedAt = &v
}

// GetBuyNow returns the BuyNow field value if set, zero value otherwise.
func (o *CommonAssetSaleListing) GetBuyNow() CommonAssetBuyNow {
	if o == nil || IsNil(o.BuyNow) {
		var ret CommonAssetBuyNow
		return ret
	}
	return *o.BuyNow
}

// GetBuyNowOk returns a tuple with the BuyNow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAssetSaleListing) GetBuyNowOk() (*CommonAssetBuyNow, bool) {
	if o == nil || IsNil(o.BuyNow) {
		return nil, false
	}
	return o.BuyNow, true
}

// HasBuyNow returns a boolean if a field has been set.
func (o *CommonAssetSaleListing) HasBuyNow() bool {
	if o != nil && !IsNil(o.BuyNow) {
		return true
	}

	return false
}

// SetBuyNow gets a reference to the given CommonAssetBuyNow and assigns it to the BuyNow field.
func (o *CommonAssetSaleListing) SetBuyNow(v CommonAssetBuyNow) {
	o.BuyNow = &v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *CommonAssetSaleListing) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *CommonAssetSaleListing) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *CommonAssetSaleListing) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o CommonAssetSaleListing) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonAssetSaleListing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["status"] = o.Status
	toSerialize["venue"] = o.Venue
	toSerialize["sale-start-at"] = o.SaleStartAt
	if !IsNil(o.SaleEndAt) {
		toSerialize["sale-end-at"] = o.SaleEndAt
	}
	toSerialize["seller-rep"] = o.SellerRep
	if !IsNil(o.Lane) {
		toSerialize["lane"] = o.Lane
	}
	if !IsNil(o.Lot) {
		toSerialize["lot"] = o.Lot
	}
	toSerialize["announcements-text"] = o.AnnouncementsText
	toSerialize["announcement-lights"] = o.AnnouncementLights
	if !IsNil(o.FloorAmount) {
		toSerialize["floor-amount"] = o.FloorAmount
	}
	if !IsNil(o.StartBid) {
		toSerialize["start-bid"] = o.StartBid
	}
	if !IsNil(o.HighBid) {
		toSerialize["high-bid"] = o.HighBid
	}
	if !IsNil(o.BidIncrement) {
		toSerialize["bid-increment"] = o.BidIncrement
	}
	if !IsNil(o.MainImageUrl) {
		toSerialize["main-image-url"] = o.MainImageUrl
	}
	if !IsNil(o.PhotoCount) {
		toSerialize["photo-count"] = o.PhotoCount
	}
	if !IsNil(o.PhotoCountUpdatedAt) {
		toSerialize["photo-count-updated-at"] = o.PhotoCountUpdatedAt
	}
	if !IsNil(o.BuyNow) {
		toSerialize["buy-now"] = o.BuyNow
	}
	toSerialize["updated-at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CommonAssetSaleListing) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"status",
		"venue",
		"sale-start-at",
		"seller-rep",
		"announcements-text",
		"announcement-lights",
		"updated-at",
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

	varCommonAssetSaleListing := _CommonAssetSaleListing{}

	err = json.Unmarshal(data, &varCommonAssetSaleListing)

	if err != nil {
		return err
	}

	*o = CommonAssetSaleListing(varCommonAssetSaleListing)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "venue")
		delete(additionalProperties, "sale-start-at")
		delete(additionalProperties, "sale-end-at")
		delete(additionalProperties, "seller-rep")
		delete(additionalProperties, "lane")
		delete(additionalProperties, "lot")
		delete(additionalProperties, "announcements-text")
		delete(additionalProperties, "announcement-lights")
		delete(additionalProperties, "floor-amount")
		delete(additionalProperties, "start-bid")
		delete(additionalProperties, "high-bid")
		delete(additionalProperties, "bid-increment")
		delete(additionalProperties, "main-image-url")
		delete(additionalProperties, "photo-count")
		delete(additionalProperties, "photo-count-updated-at")
		delete(additionalProperties, "buy-now")
		delete(additionalProperties, "updated-at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCommonAssetSaleListing struct {
	value *CommonAssetSaleListing
	isSet bool
}

func (v NullableCommonAssetSaleListing) Get() *CommonAssetSaleListing {
	return v.value
}

func (v *NullableCommonAssetSaleListing) Set(val *CommonAssetSaleListing) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonAssetSaleListing) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonAssetSaleListing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonAssetSaleListing(val *CommonAssetSaleListing) *NullableCommonAssetSaleListing {
	return &NullableCommonAssetSaleListing{value: val, isSet: true}
}

func (v NullableCommonAssetSaleListing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonAssetSaleListing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


