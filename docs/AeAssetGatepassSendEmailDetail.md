# AeAssetGatepassSendEmailDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**Asset** | Pointer to [**CommonAmsAssetPointer**](CommonAmsAssetPointer.md) |  | [optional] 
**Vin** | **string** | The Vehicle Identification Number(VIN) of the asset. | 
**Stock** | **string** | The stock number of the asset. | 
**InitiatorEmail** | Pointer to **string** | The email address of the initiator of the gatepass email request.  Errors will be sent here. | [optional] 
**SendToEmail** | **[]string** |  | 
**GatepassType** | Pointer to [**GatepassType**](GatepassType.md) |  | [optional] [default to G]
**SenderName** | Pointer to **string** | Optional name of the sender. | [optional] 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeAssetGatepassSendEmailDetail

`func NewAeAssetGatepassSendEmailDetail(auctionId string, vin string, stock string, sendToEmail []string, ) *AeAssetGatepassSendEmailDetail`

NewAeAssetGatepassSendEmailDetail instantiates a new AeAssetGatepassSendEmailDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetGatepassSendEmailDetailWithDefaults

`func NewAeAssetGatepassSendEmailDetailWithDefaults() *AeAssetGatepassSendEmailDetail`

NewAeAssetGatepassSendEmailDetailWithDefaults instantiates a new AeAssetGatepassSendEmailDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionId

`func (o *AeAssetGatepassSendEmailDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAssetGatepassSendEmailDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAssetGatepassSendEmailDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetAsset

`func (o *AeAssetGatepassSendEmailDetail) GetAsset() CommonAmsAssetPointer`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *AeAssetGatepassSendEmailDetail) GetAssetOk() (*CommonAmsAssetPointer, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *AeAssetGatepassSendEmailDetail) SetAsset(v CommonAmsAssetPointer)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *AeAssetGatepassSendEmailDetail) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetVin

`func (o *AeAssetGatepassSendEmailDetail) GetVin() string`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *AeAssetGatepassSendEmailDetail) GetVinOk() (*string, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *AeAssetGatepassSendEmailDetail) SetVin(v string)`

SetVin sets Vin field to given value.


### GetStock

`func (o *AeAssetGatepassSendEmailDetail) GetStock() string`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *AeAssetGatepassSendEmailDetail) GetStockOk() (*string, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *AeAssetGatepassSendEmailDetail) SetStock(v string)`

SetStock sets Stock field to given value.


### GetInitiatorEmail

`func (o *AeAssetGatepassSendEmailDetail) GetInitiatorEmail() string`

GetInitiatorEmail returns the InitiatorEmail field if non-nil, zero value otherwise.

### GetInitiatorEmailOk

`func (o *AeAssetGatepassSendEmailDetail) GetInitiatorEmailOk() (*string, bool)`

GetInitiatorEmailOk returns a tuple with the InitiatorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorEmail

`func (o *AeAssetGatepassSendEmailDetail) SetInitiatorEmail(v string)`

SetInitiatorEmail sets InitiatorEmail field to given value.

### HasInitiatorEmail

`func (o *AeAssetGatepassSendEmailDetail) HasInitiatorEmail() bool`

HasInitiatorEmail returns a boolean if a field has been set.

### GetSendToEmail

`func (o *AeAssetGatepassSendEmailDetail) GetSendToEmail() []string`

GetSendToEmail returns the SendToEmail field if non-nil, zero value otherwise.

### GetSendToEmailOk

`func (o *AeAssetGatepassSendEmailDetail) GetSendToEmailOk() (*[]string, bool)`

GetSendToEmailOk returns a tuple with the SendToEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendToEmail

`func (o *AeAssetGatepassSendEmailDetail) SetSendToEmail(v []string)`

SetSendToEmail sets SendToEmail field to given value.


### GetGatepassType

`func (o *AeAssetGatepassSendEmailDetail) GetGatepassType() GatepassType`

GetGatepassType returns the GatepassType field if non-nil, zero value otherwise.

### GetGatepassTypeOk

`func (o *AeAssetGatepassSendEmailDetail) GetGatepassTypeOk() (*GatepassType, bool)`

GetGatepassTypeOk returns a tuple with the GatepassType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatepassType

`func (o *AeAssetGatepassSendEmailDetail) SetGatepassType(v GatepassType)`

SetGatepassType sets GatepassType field to given value.

### HasGatepassType

`func (o *AeAssetGatepassSendEmailDetail) HasGatepassType() bool`

HasGatepassType returns a boolean if a field has been set.

### GetSenderName

`func (o *AeAssetGatepassSendEmailDetail) GetSenderName() string`

GetSenderName returns the SenderName field if non-nil, zero value otherwise.

### GetSenderNameOk

`func (o *AeAssetGatepassSendEmailDetail) GetSenderNameOk() (*string, bool)`

GetSenderNameOk returns a tuple with the SenderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderName

`func (o *AeAssetGatepassSendEmailDetail) SetSenderName(v string)`

SetSenderName sets SenderName field to given value.

### HasSenderName

`func (o *AeAssetGatepassSendEmailDetail) HasSenderName() bool`

HasSenderName returns a boolean if a field has been set.

### GetInitiator

`func (o *AeAssetGatepassSendEmailDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAssetGatepassSendEmailDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAssetGatepassSendEmailDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAssetGatepassSendEmailDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


