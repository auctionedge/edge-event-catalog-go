# AeAssetGatepassBuyerReleasableDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auction** | [**AeAssetGatepassBuyerReleasableDetailAuction**](AeAssetGatepassBuyerReleasableDetailAuction.md) |  | 
**Account** | [**AeAssetGatepassBuyerReleasableDetailAccount**](AeAssetGatepassBuyerReleasableDetailAccount.md) |  | 
**Asset** | [**AeAssetGatepassBuyerReleasableDetailAsset**](AeAssetGatepassBuyerReleasableDetailAsset.md) |  | 
**AmsSaleListing** | [**AeAssetGatepassBuyerReleasableDetailAmsSaleListing**](AeAssetGatepassBuyerReleasableDetailAmsSaleListing.md) |  | 
**BarcodePrefix** | Pointer to **string** | The prefix of the string to generate the barcode from | [optional] 
**ShowVinBarcode** | Pointer to **bool** | Flag to show asset VIN on gate pass | [optional] [default to false]
**Notes** | Pointer to **string** | Text customized by auction to print on gate pass | [optional] 
**PrintSignatureLine** | Pointer to **bool** | Flag to print a line of signature on gate pass | [optional] [default to false]
**PrintDriverLicenseLine** | Pointer to **bool** | Flag to print line on gate pass so write driver license number | [optional] [default to false]
**UpdatedAt** | Pointer to **time.Time** | The updated date time of the gatepass status | [optional] 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeAssetGatepassBuyerReleasableDetail

`func NewAeAssetGatepassBuyerReleasableDetail(auction AeAssetGatepassBuyerReleasableDetailAuction, account AeAssetGatepassBuyerReleasableDetailAccount, asset AeAssetGatepassBuyerReleasableDetailAsset, amsSaleListing AeAssetGatepassBuyerReleasableDetailAmsSaleListing, ) *AeAssetGatepassBuyerReleasableDetail`

NewAeAssetGatepassBuyerReleasableDetail instantiates a new AeAssetGatepassBuyerReleasableDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetGatepassBuyerReleasableDetailWithDefaults

`func NewAeAssetGatepassBuyerReleasableDetailWithDefaults() *AeAssetGatepassBuyerReleasableDetail`

NewAeAssetGatepassBuyerReleasableDetailWithDefaults instantiates a new AeAssetGatepassBuyerReleasableDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuction

`func (o *AeAssetGatepassBuyerReleasableDetail) GetAuction() AeAssetGatepassBuyerReleasableDetailAuction`

GetAuction returns the Auction field if non-nil, zero value otherwise.

### GetAuctionOk

`func (o *AeAssetGatepassBuyerReleasableDetail) GetAuctionOk() (*AeAssetGatepassBuyerReleasableDetailAuction, bool)`

GetAuctionOk returns a tuple with the Auction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuction

`func (o *AeAssetGatepassBuyerReleasableDetail) SetAuction(v AeAssetGatepassBuyerReleasableDetailAuction)`

SetAuction sets Auction field to given value.


### GetAccount

`func (o *AeAssetGatepassBuyerReleasableDetail) GetAccount() AeAssetGatepassBuyerReleasableDetailAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AeAssetGatepassBuyerReleasableDetail) GetAccountOk() (*AeAssetGatepassBuyerReleasableDetailAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AeAssetGatepassBuyerReleasableDetail) SetAccount(v AeAssetGatepassBuyerReleasableDetailAccount)`

SetAccount sets Account field to given value.


### GetAsset

`func (o *AeAssetGatepassBuyerReleasableDetail) GetAsset() AeAssetGatepassBuyerReleasableDetailAsset`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *AeAssetGatepassBuyerReleasableDetail) GetAssetOk() (*AeAssetGatepassBuyerReleasableDetailAsset, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *AeAssetGatepassBuyerReleasableDetail) SetAsset(v AeAssetGatepassBuyerReleasableDetailAsset)`

SetAsset sets Asset field to given value.


### GetAmsSaleListing

`func (o *AeAssetGatepassBuyerReleasableDetail) GetAmsSaleListing() AeAssetGatepassBuyerReleasableDetailAmsSaleListing`

GetAmsSaleListing returns the AmsSaleListing field if non-nil, zero value otherwise.

### GetAmsSaleListingOk

`func (o *AeAssetGatepassBuyerReleasableDetail) GetAmsSaleListingOk() (*AeAssetGatepassBuyerReleasableDetailAmsSaleListing, bool)`

GetAmsSaleListingOk returns a tuple with the AmsSaleListing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmsSaleListing

`func (o *AeAssetGatepassBuyerReleasableDetail) SetAmsSaleListing(v AeAssetGatepassBuyerReleasableDetailAmsSaleListing)`

SetAmsSaleListing sets AmsSaleListing field to given value.


### GetBarcodePrefix

`func (o *AeAssetGatepassBuyerReleasableDetail) GetBarcodePrefix() string`

GetBarcodePrefix returns the BarcodePrefix field if non-nil, zero value otherwise.

### GetBarcodePrefixOk

`func (o *AeAssetGatepassBuyerReleasableDetail) GetBarcodePrefixOk() (*string, bool)`

GetBarcodePrefixOk returns a tuple with the BarcodePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcodePrefix

`func (o *AeAssetGatepassBuyerReleasableDetail) SetBarcodePrefix(v string)`

SetBarcodePrefix sets BarcodePrefix field to given value.

### HasBarcodePrefix

`func (o *AeAssetGatepassBuyerReleasableDetail) HasBarcodePrefix() bool`

HasBarcodePrefix returns a boolean if a field has been set.

### GetShowVinBarcode

`func (o *AeAssetGatepassBuyerReleasableDetail) GetShowVinBarcode() bool`

GetShowVinBarcode returns the ShowVinBarcode field if non-nil, zero value otherwise.

### GetShowVinBarcodeOk

`func (o *AeAssetGatepassBuyerReleasableDetail) GetShowVinBarcodeOk() (*bool, bool)`

GetShowVinBarcodeOk returns a tuple with the ShowVinBarcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowVinBarcode

`func (o *AeAssetGatepassBuyerReleasableDetail) SetShowVinBarcode(v bool)`

SetShowVinBarcode sets ShowVinBarcode field to given value.

### HasShowVinBarcode

`func (o *AeAssetGatepassBuyerReleasableDetail) HasShowVinBarcode() bool`

HasShowVinBarcode returns a boolean if a field has been set.

### GetNotes

`func (o *AeAssetGatepassBuyerReleasableDetail) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AeAssetGatepassBuyerReleasableDetail) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AeAssetGatepassBuyerReleasableDetail) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *AeAssetGatepassBuyerReleasableDetail) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetPrintSignatureLine

`func (o *AeAssetGatepassBuyerReleasableDetail) GetPrintSignatureLine() bool`

GetPrintSignatureLine returns the PrintSignatureLine field if non-nil, zero value otherwise.

### GetPrintSignatureLineOk

`func (o *AeAssetGatepassBuyerReleasableDetail) GetPrintSignatureLineOk() (*bool, bool)`

GetPrintSignatureLineOk returns a tuple with the PrintSignatureLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintSignatureLine

`func (o *AeAssetGatepassBuyerReleasableDetail) SetPrintSignatureLine(v bool)`

SetPrintSignatureLine sets PrintSignatureLine field to given value.

### HasPrintSignatureLine

`func (o *AeAssetGatepassBuyerReleasableDetail) HasPrintSignatureLine() bool`

HasPrintSignatureLine returns a boolean if a field has been set.

### GetPrintDriverLicenseLine

`func (o *AeAssetGatepassBuyerReleasableDetail) GetPrintDriverLicenseLine() bool`

GetPrintDriverLicenseLine returns the PrintDriverLicenseLine field if non-nil, zero value otherwise.

### GetPrintDriverLicenseLineOk

`func (o *AeAssetGatepassBuyerReleasableDetail) GetPrintDriverLicenseLineOk() (*bool, bool)`

GetPrintDriverLicenseLineOk returns a tuple with the PrintDriverLicenseLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintDriverLicenseLine

`func (o *AeAssetGatepassBuyerReleasableDetail) SetPrintDriverLicenseLine(v bool)`

SetPrintDriverLicenseLine sets PrintDriverLicenseLine field to given value.

### HasPrintDriverLicenseLine

`func (o *AeAssetGatepassBuyerReleasableDetail) HasPrintDriverLicenseLine() bool`

HasPrintDriverLicenseLine returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AeAssetGatepassBuyerReleasableDetail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AeAssetGatepassBuyerReleasableDetail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AeAssetGatepassBuyerReleasableDetail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AeAssetGatepassBuyerReleasableDetail) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetInitiator

`func (o *AeAssetGatepassBuyerReleasableDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAssetGatepassBuyerReleasableDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAssetGatepassBuyerReleasableDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAssetGatepassBuyerReleasableDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


