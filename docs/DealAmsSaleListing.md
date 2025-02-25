# DealAmsSaleListing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SaleDate** | **string** | Date on which the asset ran that the deal was created for. | 
**Lane** | **string** | The lane that the asset ran in. | 
**Lot** | **string** | The lot that the asset was attached to. | 

## Methods

### NewDealAmsSaleListing

`func NewDealAmsSaleListing(saleDate string, lane string, lot string, ) *DealAmsSaleListing`

NewDealAmsSaleListing instantiates a new DealAmsSaleListing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDealAmsSaleListingWithDefaults

`func NewDealAmsSaleListingWithDefaults() *DealAmsSaleListing`

NewDealAmsSaleListingWithDefaults instantiates a new DealAmsSaleListing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSaleDate

`func (o *DealAmsSaleListing) GetSaleDate() string`

GetSaleDate returns the SaleDate field if non-nil, zero value otherwise.

### GetSaleDateOk

`func (o *DealAmsSaleListing) GetSaleDateOk() (*string, bool)`

GetSaleDateOk returns a tuple with the SaleDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleDate

`func (o *DealAmsSaleListing) SetSaleDate(v string)`

SetSaleDate sets SaleDate field to given value.


### GetLane

`func (o *DealAmsSaleListing) GetLane() string`

GetLane returns the Lane field if non-nil, zero value otherwise.

### GetLaneOk

`func (o *DealAmsSaleListing) GetLaneOk() (*string, bool)`

GetLaneOk returns a tuple with the Lane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLane

`func (o *DealAmsSaleListing) SetLane(v string)`

SetLane sets Lane field to given value.


### GetLot

`func (o *DealAmsSaleListing) GetLot() string`

GetLot returns the Lot field if non-nil, zero value otherwise.

### GetLotOk

`func (o *DealAmsSaleListing) GetLotOk() (*string, bool)`

GetLotOk returns a tuple with the Lot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLot

`func (o *DealAmsSaleListing) SetLot(v string)`

SetLot sets Lot field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


