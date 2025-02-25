# CommonAssetSaleListing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | AMS unique identifier for this sale listing of the asset | 
**Status** | **string** | Status of the sale listing | 
**Venue** | [**CommonVenue**](CommonVenue.md) |  | 
**SaleStartAt** | **time.Time** | The time that this sale listing for the asset begins | 
**SaleEndAt** | Pointer to **time.Time** | The time that this sale listing for the asset ends | [optional] 
**SellerRep** | **string** | The representative id an AMS account | 
**Lane** | Pointer to **string** | The lane the asset was in when deal made. | [optional] 
**Lot** | Pointer to **string** | The lot the asset was in when deal made. | [optional] 
**AnnouncementsText** | **[]string** | Auction announcements on the asset | 
**AnnouncementLights** | **string** | The current lights on the asset when the sale listing was made. | 
**FloorAmount** | Pointer to [**CommonCurrency**](CommonCurrency.md) |  | [optional] 
**StartBid** | Pointer to [**CommonCurrency**](CommonCurrency.md) |  | [optional] 
**HighBid** | Pointer to [**CommonCurrency**](CommonCurrency.md) |  | [optional] 
**BidIncrement** | Pointer to [**CommonCurrency**](CommonCurrency.md) |  | [optional] 
**MainImageUrl** | Pointer to **string** | The URL for the main image of the asset to be posted on the listing | [optional] 
**PhotoCount** | Pointer to **string** | total number of photos attached to the asset | [optional] 
**PhotoCountUpdatedAt** | Pointer to **time.Time** | timestamp of when photo count was last changed | [optional] 
**BuyNow** | Pointer to [**CommonAssetBuyNow**](CommonAssetBuyNow.md) |  | [optional] 
**UpdatedAt** | **time.Time** | The time in ISO-8601 formatted at which the record was generated | 

## Methods

### NewCommonAssetSaleListing

`func NewCommonAssetSaleListing(id string, status string, venue CommonVenue, saleStartAt time.Time, sellerRep string, announcementsText []string, announcementLights string, updatedAt time.Time, ) *CommonAssetSaleListing`

NewCommonAssetSaleListing instantiates a new CommonAssetSaleListing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonAssetSaleListingWithDefaults

`func NewCommonAssetSaleListingWithDefaults() *CommonAssetSaleListing`

NewCommonAssetSaleListingWithDefaults instantiates a new CommonAssetSaleListing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommonAssetSaleListing) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommonAssetSaleListing) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommonAssetSaleListing) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *CommonAssetSaleListing) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CommonAssetSaleListing) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CommonAssetSaleListing) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetVenue

`func (o *CommonAssetSaleListing) GetVenue() CommonVenue`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *CommonAssetSaleListing) GetVenueOk() (*CommonVenue, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *CommonAssetSaleListing) SetVenue(v CommonVenue)`

SetVenue sets Venue field to given value.


### GetSaleStartAt

`func (o *CommonAssetSaleListing) GetSaleStartAt() time.Time`

GetSaleStartAt returns the SaleStartAt field if non-nil, zero value otherwise.

### GetSaleStartAtOk

`func (o *CommonAssetSaleListing) GetSaleStartAtOk() (*time.Time, bool)`

GetSaleStartAtOk returns a tuple with the SaleStartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleStartAt

`func (o *CommonAssetSaleListing) SetSaleStartAt(v time.Time)`

SetSaleStartAt sets SaleStartAt field to given value.


### GetSaleEndAt

`func (o *CommonAssetSaleListing) GetSaleEndAt() time.Time`

GetSaleEndAt returns the SaleEndAt field if non-nil, zero value otherwise.

### GetSaleEndAtOk

`func (o *CommonAssetSaleListing) GetSaleEndAtOk() (*time.Time, bool)`

GetSaleEndAtOk returns a tuple with the SaleEndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleEndAt

`func (o *CommonAssetSaleListing) SetSaleEndAt(v time.Time)`

SetSaleEndAt sets SaleEndAt field to given value.

### HasSaleEndAt

`func (o *CommonAssetSaleListing) HasSaleEndAt() bool`

HasSaleEndAt returns a boolean if a field has been set.

### GetSellerRep

`func (o *CommonAssetSaleListing) GetSellerRep() string`

GetSellerRep returns the SellerRep field if non-nil, zero value otherwise.

### GetSellerRepOk

`func (o *CommonAssetSaleListing) GetSellerRepOk() (*string, bool)`

GetSellerRepOk returns a tuple with the SellerRep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerRep

`func (o *CommonAssetSaleListing) SetSellerRep(v string)`

SetSellerRep sets SellerRep field to given value.


### GetLane

`func (o *CommonAssetSaleListing) GetLane() string`

GetLane returns the Lane field if non-nil, zero value otherwise.

### GetLaneOk

`func (o *CommonAssetSaleListing) GetLaneOk() (*string, bool)`

GetLaneOk returns a tuple with the Lane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLane

`func (o *CommonAssetSaleListing) SetLane(v string)`

SetLane sets Lane field to given value.

### HasLane

`func (o *CommonAssetSaleListing) HasLane() bool`

HasLane returns a boolean if a field has been set.

### GetLot

`func (o *CommonAssetSaleListing) GetLot() string`

GetLot returns the Lot field if non-nil, zero value otherwise.

### GetLotOk

`func (o *CommonAssetSaleListing) GetLotOk() (*string, bool)`

GetLotOk returns a tuple with the Lot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLot

`func (o *CommonAssetSaleListing) SetLot(v string)`

SetLot sets Lot field to given value.

### HasLot

`func (o *CommonAssetSaleListing) HasLot() bool`

HasLot returns a boolean if a field has been set.

### GetAnnouncementsText

`func (o *CommonAssetSaleListing) GetAnnouncementsText() []string`

GetAnnouncementsText returns the AnnouncementsText field if non-nil, zero value otherwise.

### GetAnnouncementsTextOk

`func (o *CommonAssetSaleListing) GetAnnouncementsTextOk() (*[]string, bool)`

GetAnnouncementsTextOk returns a tuple with the AnnouncementsText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementsText

`func (o *CommonAssetSaleListing) SetAnnouncementsText(v []string)`

SetAnnouncementsText sets AnnouncementsText field to given value.


### GetAnnouncementLights

`func (o *CommonAssetSaleListing) GetAnnouncementLights() string`

GetAnnouncementLights returns the AnnouncementLights field if non-nil, zero value otherwise.

### GetAnnouncementLightsOk

`func (o *CommonAssetSaleListing) GetAnnouncementLightsOk() (*string, bool)`

GetAnnouncementLightsOk returns a tuple with the AnnouncementLights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementLights

`func (o *CommonAssetSaleListing) SetAnnouncementLights(v string)`

SetAnnouncementLights sets AnnouncementLights field to given value.


### GetFloorAmount

`func (o *CommonAssetSaleListing) GetFloorAmount() CommonCurrency`

GetFloorAmount returns the FloorAmount field if non-nil, zero value otherwise.

### GetFloorAmountOk

`func (o *CommonAssetSaleListing) GetFloorAmountOk() (*CommonCurrency, bool)`

GetFloorAmountOk returns a tuple with the FloorAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorAmount

`func (o *CommonAssetSaleListing) SetFloorAmount(v CommonCurrency)`

SetFloorAmount sets FloorAmount field to given value.

### HasFloorAmount

`func (o *CommonAssetSaleListing) HasFloorAmount() bool`

HasFloorAmount returns a boolean if a field has been set.

### GetStartBid

`func (o *CommonAssetSaleListing) GetStartBid() CommonCurrency`

GetStartBid returns the StartBid field if non-nil, zero value otherwise.

### GetStartBidOk

`func (o *CommonAssetSaleListing) GetStartBidOk() (*CommonCurrency, bool)`

GetStartBidOk returns a tuple with the StartBid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartBid

`func (o *CommonAssetSaleListing) SetStartBid(v CommonCurrency)`

SetStartBid sets StartBid field to given value.

### HasStartBid

`func (o *CommonAssetSaleListing) HasStartBid() bool`

HasStartBid returns a boolean if a field has been set.

### GetHighBid

`func (o *CommonAssetSaleListing) GetHighBid() CommonCurrency`

GetHighBid returns the HighBid field if non-nil, zero value otherwise.

### GetHighBidOk

`func (o *CommonAssetSaleListing) GetHighBidOk() (*CommonCurrency, bool)`

GetHighBidOk returns a tuple with the HighBid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighBid

`func (o *CommonAssetSaleListing) SetHighBid(v CommonCurrency)`

SetHighBid sets HighBid field to given value.

### HasHighBid

`func (o *CommonAssetSaleListing) HasHighBid() bool`

HasHighBid returns a boolean if a field has been set.

### GetBidIncrement

`func (o *CommonAssetSaleListing) GetBidIncrement() CommonCurrency`

GetBidIncrement returns the BidIncrement field if non-nil, zero value otherwise.

### GetBidIncrementOk

`func (o *CommonAssetSaleListing) GetBidIncrementOk() (*CommonCurrency, bool)`

GetBidIncrementOk returns a tuple with the BidIncrement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidIncrement

`func (o *CommonAssetSaleListing) SetBidIncrement(v CommonCurrency)`

SetBidIncrement sets BidIncrement field to given value.

### HasBidIncrement

`func (o *CommonAssetSaleListing) HasBidIncrement() bool`

HasBidIncrement returns a boolean if a field has been set.

### GetMainImageUrl

`func (o *CommonAssetSaleListing) GetMainImageUrl() string`

GetMainImageUrl returns the MainImageUrl field if non-nil, zero value otherwise.

### GetMainImageUrlOk

`func (o *CommonAssetSaleListing) GetMainImageUrlOk() (*string, bool)`

GetMainImageUrlOk returns a tuple with the MainImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainImageUrl

`func (o *CommonAssetSaleListing) SetMainImageUrl(v string)`

SetMainImageUrl sets MainImageUrl field to given value.

### HasMainImageUrl

`func (o *CommonAssetSaleListing) HasMainImageUrl() bool`

HasMainImageUrl returns a boolean if a field has been set.

### GetPhotoCount

`func (o *CommonAssetSaleListing) GetPhotoCount() string`

GetPhotoCount returns the PhotoCount field if non-nil, zero value otherwise.

### GetPhotoCountOk

`func (o *CommonAssetSaleListing) GetPhotoCountOk() (*string, bool)`

GetPhotoCountOk returns a tuple with the PhotoCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoCount

`func (o *CommonAssetSaleListing) SetPhotoCount(v string)`

SetPhotoCount sets PhotoCount field to given value.

### HasPhotoCount

`func (o *CommonAssetSaleListing) HasPhotoCount() bool`

HasPhotoCount returns a boolean if a field has been set.

### GetPhotoCountUpdatedAt

`func (o *CommonAssetSaleListing) GetPhotoCountUpdatedAt() time.Time`

GetPhotoCountUpdatedAt returns the PhotoCountUpdatedAt field if non-nil, zero value otherwise.

### GetPhotoCountUpdatedAtOk

`func (o *CommonAssetSaleListing) GetPhotoCountUpdatedAtOk() (*time.Time, bool)`

GetPhotoCountUpdatedAtOk returns a tuple with the PhotoCountUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoCountUpdatedAt

`func (o *CommonAssetSaleListing) SetPhotoCountUpdatedAt(v time.Time)`

SetPhotoCountUpdatedAt sets PhotoCountUpdatedAt field to given value.

### HasPhotoCountUpdatedAt

`func (o *CommonAssetSaleListing) HasPhotoCountUpdatedAt() bool`

HasPhotoCountUpdatedAt returns a boolean if a field has been set.

### GetBuyNow

`func (o *CommonAssetSaleListing) GetBuyNow() CommonAssetBuyNow`

GetBuyNow returns the BuyNow field if non-nil, zero value otherwise.

### GetBuyNowOk

`func (o *CommonAssetSaleListing) GetBuyNowOk() (*CommonAssetBuyNow, bool)`

GetBuyNowOk returns a tuple with the BuyNow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyNow

`func (o *CommonAssetSaleListing) SetBuyNow(v CommonAssetBuyNow)`

SetBuyNow sets BuyNow field to given value.

### HasBuyNow

`func (o *CommonAssetSaleListing) HasBuyNow() bool`

HasBuyNow returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CommonAssetSaleListing) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CommonAssetSaleListing) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CommonAssetSaleListing) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


