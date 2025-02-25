# AeAdvisoryAccountAmsDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountKey** | **string** | The account key of an AMS account (i.e. the account or dealer number) | 
**AaId** | **NullableString** | The Auction Access ID of the AMS account. | 
**AccountId** | Pointer to **string** | Source&#39;s unique identifier for account | [optional] 
**AuctionId** | **string** | Auction Edge unique identifier for an auction. | 
**Id** | Pointer to **string** | Source&#39;s unique identifier for account | [optional] 
**EstablishedAt** | Pointer to **time.Time** | The date and time that the record was initially created in the AMS | [optional] 
**DisplayName** | **string** | The display name of the account | 
**LegalName** | Pointer to **string** | The legal name of the account | [optional] 
**DbaName** | Pointer to **string** | The &#39;doing business as&#39; name of the account | [optional] 
**PrimaryEmail** | Pointer to **string** | Primary email of dealership | [optional] 
**WebsiteUri** | Pointer to **string** | URI dealer website | [optional] 
**PrimaryPhoneNumber** | Pointer to **string** | Main phone number of account | [optional] 
**CellPhoneNumber** | Pointer to **string** | Cell phone number on record at AMS for account | [optional] 
**PrimaryFaxNumber** | Pointer to **string** | Fax phone number | [optional] 
**PhysicalLocation** | Pointer to [**CommonAmsPostalAddress**](CommonAmsPostalAddress.md) |  | [optional] 
**MailingLocation** | Pointer to [**CommonAmsPostalAddress**](CommonAmsPostalAddress.md) |  | [optional] 
**CanBuy** | Pointer to **bool** | True if account is allowed to buy a vehicle, false if not | [optional] 
**CanSell** | Pointer to **bool** | True if account is allowed to sell a vehicle, false if not | [optional] 
**IsActive** | Pointer to **bool** | True if account is active, false if not | [optional] 
**IsInViolation** | Pointer to **bool** | True if account is in violation of an auction rule such as expired license, false if not | [optional] 
**IsTaxExempt** | Pointer to **bool** | True if account is tax exempt, false if not | [optional] 
**AuctionSellerRep** | Pointer to **string** | Auction staff member that assists account selling vehicles | [optional] 
**AuctionBuyerRep** | Pointer to **string** | Auction staff member that assists account buying vehicles | [optional] 
**AutoimsId** | Pointer to **float32** | AutoIMS id | [optional] 
**AutoimsIdStr** | Pointer to **string** | AutoIMS id | [optional] 
**PaymentSources** | Pointer to [**[]CommonAmsPaymentSource**](CommonAmsPaymentSource.md) | The payment sources for the account | [optional] 
**UpdatedAt** | **time.Time** | The time in ISO-8601 formatted at which the record was generated | 
**Initiator** | Pointer to [**CommonInitiator**](CommonInitiator.md) |  | [optional] 

## Methods

### NewAeAdvisoryAccountAmsDetail

`func NewAeAdvisoryAccountAmsDetail(accountKey string, aaId NullableString, auctionId string, displayName string, updatedAt time.Time, ) *AeAdvisoryAccountAmsDetail`

NewAeAdvisoryAccountAmsDetail instantiates a new AeAdvisoryAccountAmsDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAdvisoryAccountAmsDetailWithDefaults

`func NewAeAdvisoryAccountAmsDetailWithDefaults() *AeAdvisoryAccountAmsDetail`

NewAeAdvisoryAccountAmsDetailWithDefaults instantiates a new AeAdvisoryAccountAmsDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountKey

`func (o *AeAdvisoryAccountAmsDetail) GetAccountKey() string`

GetAccountKey returns the AccountKey field if non-nil, zero value otherwise.

### GetAccountKeyOk

`func (o *AeAdvisoryAccountAmsDetail) GetAccountKeyOk() (*string, bool)`

GetAccountKeyOk returns a tuple with the AccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountKey

`func (o *AeAdvisoryAccountAmsDetail) SetAccountKey(v string)`

SetAccountKey sets AccountKey field to given value.


### GetAaId

`func (o *AeAdvisoryAccountAmsDetail) GetAaId() string`

GetAaId returns the AaId field if non-nil, zero value otherwise.

### GetAaIdOk

`func (o *AeAdvisoryAccountAmsDetail) GetAaIdOk() (*string, bool)`

GetAaIdOk returns a tuple with the AaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAaId

`func (o *AeAdvisoryAccountAmsDetail) SetAaId(v string)`

SetAaId sets AaId field to given value.


### SetAaIdNil

`func (o *AeAdvisoryAccountAmsDetail) SetAaIdNil(b bool)`

 SetAaIdNil sets the value for AaId to be an explicit nil

### UnsetAaId
`func (o *AeAdvisoryAccountAmsDetail) UnsetAaId()`

UnsetAaId ensures that no value is present for AaId, not even an explicit nil
### GetAccountId

`func (o *AeAdvisoryAccountAmsDetail) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AeAdvisoryAccountAmsDetail) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AeAdvisoryAccountAmsDetail) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AeAdvisoryAccountAmsDetail) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAuctionId

`func (o *AeAdvisoryAccountAmsDetail) GetAuctionId() string`

GetAuctionId returns the AuctionId field if non-nil, zero value otherwise.

### GetAuctionIdOk

`func (o *AeAdvisoryAccountAmsDetail) GetAuctionIdOk() (*string, bool)`

GetAuctionIdOk returns a tuple with the AuctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionId

`func (o *AeAdvisoryAccountAmsDetail) SetAuctionId(v string)`

SetAuctionId sets AuctionId field to given value.


### GetId

`func (o *AeAdvisoryAccountAmsDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AeAdvisoryAccountAmsDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AeAdvisoryAccountAmsDetail) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AeAdvisoryAccountAmsDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEstablishedAt

`func (o *AeAdvisoryAccountAmsDetail) GetEstablishedAt() time.Time`

GetEstablishedAt returns the EstablishedAt field if non-nil, zero value otherwise.

### GetEstablishedAtOk

`func (o *AeAdvisoryAccountAmsDetail) GetEstablishedAtOk() (*time.Time, bool)`

GetEstablishedAtOk returns a tuple with the EstablishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstablishedAt

`func (o *AeAdvisoryAccountAmsDetail) SetEstablishedAt(v time.Time)`

SetEstablishedAt sets EstablishedAt field to given value.

### HasEstablishedAt

`func (o *AeAdvisoryAccountAmsDetail) HasEstablishedAt() bool`

HasEstablishedAt returns a boolean if a field has been set.

### GetDisplayName

`func (o *AeAdvisoryAccountAmsDetail) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AeAdvisoryAccountAmsDetail) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AeAdvisoryAccountAmsDetail) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetLegalName

`func (o *AeAdvisoryAccountAmsDetail) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *AeAdvisoryAccountAmsDetail) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *AeAdvisoryAccountAmsDetail) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.

### HasLegalName

`func (o *AeAdvisoryAccountAmsDetail) HasLegalName() bool`

HasLegalName returns a boolean if a field has been set.

### GetDbaName

`func (o *AeAdvisoryAccountAmsDetail) GetDbaName() string`

GetDbaName returns the DbaName field if non-nil, zero value otherwise.

### GetDbaNameOk

`func (o *AeAdvisoryAccountAmsDetail) GetDbaNameOk() (*string, bool)`

GetDbaNameOk returns a tuple with the DbaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbaName

`func (o *AeAdvisoryAccountAmsDetail) SetDbaName(v string)`

SetDbaName sets DbaName field to given value.

### HasDbaName

`func (o *AeAdvisoryAccountAmsDetail) HasDbaName() bool`

HasDbaName returns a boolean if a field has been set.

### GetPrimaryEmail

`func (o *AeAdvisoryAccountAmsDetail) GetPrimaryEmail() string`

GetPrimaryEmail returns the PrimaryEmail field if non-nil, zero value otherwise.

### GetPrimaryEmailOk

`func (o *AeAdvisoryAccountAmsDetail) GetPrimaryEmailOk() (*string, bool)`

GetPrimaryEmailOk returns a tuple with the PrimaryEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryEmail

`func (o *AeAdvisoryAccountAmsDetail) SetPrimaryEmail(v string)`

SetPrimaryEmail sets PrimaryEmail field to given value.

### HasPrimaryEmail

`func (o *AeAdvisoryAccountAmsDetail) HasPrimaryEmail() bool`

HasPrimaryEmail returns a boolean if a field has been set.

### GetWebsiteUri

`func (o *AeAdvisoryAccountAmsDetail) GetWebsiteUri() string`

GetWebsiteUri returns the WebsiteUri field if non-nil, zero value otherwise.

### GetWebsiteUriOk

`func (o *AeAdvisoryAccountAmsDetail) GetWebsiteUriOk() (*string, bool)`

GetWebsiteUriOk returns a tuple with the WebsiteUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUri

`func (o *AeAdvisoryAccountAmsDetail) SetWebsiteUri(v string)`

SetWebsiteUri sets WebsiteUri field to given value.

### HasWebsiteUri

`func (o *AeAdvisoryAccountAmsDetail) HasWebsiteUri() bool`

HasWebsiteUri returns a boolean if a field has been set.

### GetPrimaryPhoneNumber

`func (o *AeAdvisoryAccountAmsDetail) GetPrimaryPhoneNumber() string`

GetPrimaryPhoneNumber returns the PrimaryPhoneNumber field if non-nil, zero value otherwise.

### GetPrimaryPhoneNumberOk

`func (o *AeAdvisoryAccountAmsDetail) GetPrimaryPhoneNumberOk() (*string, bool)`

GetPrimaryPhoneNumberOk returns a tuple with the PrimaryPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryPhoneNumber

`func (o *AeAdvisoryAccountAmsDetail) SetPrimaryPhoneNumber(v string)`

SetPrimaryPhoneNumber sets PrimaryPhoneNumber field to given value.

### HasPrimaryPhoneNumber

`func (o *AeAdvisoryAccountAmsDetail) HasPrimaryPhoneNumber() bool`

HasPrimaryPhoneNumber returns a boolean if a field has been set.

### GetCellPhoneNumber

`func (o *AeAdvisoryAccountAmsDetail) GetCellPhoneNumber() string`

GetCellPhoneNumber returns the CellPhoneNumber field if non-nil, zero value otherwise.

### GetCellPhoneNumberOk

`func (o *AeAdvisoryAccountAmsDetail) GetCellPhoneNumberOk() (*string, bool)`

GetCellPhoneNumberOk returns a tuple with the CellPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellPhoneNumber

`func (o *AeAdvisoryAccountAmsDetail) SetCellPhoneNumber(v string)`

SetCellPhoneNumber sets CellPhoneNumber field to given value.

### HasCellPhoneNumber

`func (o *AeAdvisoryAccountAmsDetail) HasCellPhoneNumber() bool`

HasCellPhoneNumber returns a boolean if a field has been set.

### GetPrimaryFaxNumber

`func (o *AeAdvisoryAccountAmsDetail) GetPrimaryFaxNumber() string`

GetPrimaryFaxNumber returns the PrimaryFaxNumber field if non-nil, zero value otherwise.

### GetPrimaryFaxNumberOk

`func (o *AeAdvisoryAccountAmsDetail) GetPrimaryFaxNumberOk() (*string, bool)`

GetPrimaryFaxNumberOk returns a tuple with the PrimaryFaxNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryFaxNumber

`func (o *AeAdvisoryAccountAmsDetail) SetPrimaryFaxNumber(v string)`

SetPrimaryFaxNumber sets PrimaryFaxNumber field to given value.

### HasPrimaryFaxNumber

`func (o *AeAdvisoryAccountAmsDetail) HasPrimaryFaxNumber() bool`

HasPrimaryFaxNumber returns a boolean if a field has been set.

### GetPhysicalLocation

`func (o *AeAdvisoryAccountAmsDetail) GetPhysicalLocation() CommonAmsPostalAddress`

GetPhysicalLocation returns the PhysicalLocation field if non-nil, zero value otherwise.

### GetPhysicalLocationOk

`func (o *AeAdvisoryAccountAmsDetail) GetPhysicalLocationOk() (*CommonAmsPostalAddress, bool)`

GetPhysicalLocationOk returns a tuple with the PhysicalLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalLocation

`func (o *AeAdvisoryAccountAmsDetail) SetPhysicalLocation(v CommonAmsPostalAddress)`

SetPhysicalLocation sets PhysicalLocation field to given value.

### HasPhysicalLocation

`func (o *AeAdvisoryAccountAmsDetail) HasPhysicalLocation() bool`

HasPhysicalLocation returns a boolean if a field has been set.

### GetMailingLocation

`func (o *AeAdvisoryAccountAmsDetail) GetMailingLocation() CommonAmsPostalAddress`

GetMailingLocation returns the MailingLocation field if non-nil, zero value otherwise.

### GetMailingLocationOk

`func (o *AeAdvisoryAccountAmsDetail) GetMailingLocationOk() (*CommonAmsPostalAddress, bool)`

GetMailingLocationOk returns a tuple with the MailingLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailingLocation

`func (o *AeAdvisoryAccountAmsDetail) SetMailingLocation(v CommonAmsPostalAddress)`

SetMailingLocation sets MailingLocation field to given value.

### HasMailingLocation

`func (o *AeAdvisoryAccountAmsDetail) HasMailingLocation() bool`

HasMailingLocation returns a boolean if a field has been set.

### GetCanBuy

`func (o *AeAdvisoryAccountAmsDetail) GetCanBuy() bool`

GetCanBuy returns the CanBuy field if non-nil, zero value otherwise.

### GetCanBuyOk

`func (o *AeAdvisoryAccountAmsDetail) GetCanBuyOk() (*bool, bool)`

GetCanBuyOk returns a tuple with the CanBuy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanBuy

`func (o *AeAdvisoryAccountAmsDetail) SetCanBuy(v bool)`

SetCanBuy sets CanBuy field to given value.

### HasCanBuy

`func (o *AeAdvisoryAccountAmsDetail) HasCanBuy() bool`

HasCanBuy returns a boolean if a field has been set.

### GetCanSell

`func (o *AeAdvisoryAccountAmsDetail) GetCanSell() bool`

GetCanSell returns the CanSell field if non-nil, zero value otherwise.

### GetCanSellOk

`func (o *AeAdvisoryAccountAmsDetail) GetCanSellOk() (*bool, bool)`

GetCanSellOk returns a tuple with the CanSell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSell

`func (o *AeAdvisoryAccountAmsDetail) SetCanSell(v bool)`

SetCanSell sets CanSell field to given value.

### HasCanSell

`func (o *AeAdvisoryAccountAmsDetail) HasCanSell() bool`

HasCanSell returns a boolean if a field has been set.

### GetIsActive

`func (o *AeAdvisoryAccountAmsDetail) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *AeAdvisoryAccountAmsDetail) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *AeAdvisoryAccountAmsDetail) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *AeAdvisoryAccountAmsDetail) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsInViolation

`func (o *AeAdvisoryAccountAmsDetail) GetIsInViolation() bool`

GetIsInViolation returns the IsInViolation field if non-nil, zero value otherwise.

### GetIsInViolationOk

`func (o *AeAdvisoryAccountAmsDetail) GetIsInViolationOk() (*bool, bool)`

GetIsInViolationOk returns a tuple with the IsInViolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInViolation

`func (o *AeAdvisoryAccountAmsDetail) SetIsInViolation(v bool)`

SetIsInViolation sets IsInViolation field to given value.

### HasIsInViolation

`func (o *AeAdvisoryAccountAmsDetail) HasIsInViolation() bool`

HasIsInViolation returns a boolean if a field has been set.

### GetIsTaxExempt

`func (o *AeAdvisoryAccountAmsDetail) GetIsTaxExempt() bool`

GetIsTaxExempt returns the IsTaxExempt field if non-nil, zero value otherwise.

### GetIsTaxExemptOk

`func (o *AeAdvisoryAccountAmsDetail) GetIsTaxExemptOk() (*bool, bool)`

GetIsTaxExemptOk returns a tuple with the IsTaxExempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTaxExempt

`func (o *AeAdvisoryAccountAmsDetail) SetIsTaxExempt(v bool)`

SetIsTaxExempt sets IsTaxExempt field to given value.

### HasIsTaxExempt

`func (o *AeAdvisoryAccountAmsDetail) HasIsTaxExempt() bool`

HasIsTaxExempt returns a boolean if a field has been set.

### GetAuctionSellerRep

`func (o *AeAdvisoryAccountAmsDetail) GetAuctionSellerRep() string`

GetAuctionSellerRep returns the AuctionSellerRep field if non-nil, zero value otherwise.

### GetAuctionSellerRepOk

`func (o *AeAdvisoryAccountAmsDetail) GetAuctionSellerRepOk() (*string, bool)`

GetAuctionSellerRepOk returns a tuple with the AuctionSellerRep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionSellerRep

`func (o *AeAdvisoryAccountAmsDetail) SetAuctionSellerRep(v string)`

SetAuctionSellerRep sets AuctionSellerRep field to given value.

### HasAuctionSellerRep

`func (o *AeAdvisoryAccountAmsDetail) HasAuctionSellerRep() bool`

HasAuctionSellerRep returns a boolean if a field has been set.

### GetAuctionBuyerRep

`func (o *AeAdvisoryAccountAmsDetail) GetAuctionBuyerRep() string`

GetAuctionBuyerRep returns the AuctionBuyerRep field if non-nil, zero value otherwise.

### GetAuctionBuyerRepOk

`func (o *AeAdvisoryAccountAmsDetail) GetAuctionBuyerRepOk() (*string, bool)`

GetAuctionBuyerRepOk returns a tuple with the AuctionBuyerRep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionBuyerRep

`func (o *AeAdvisoryAccountAmsDetail) SetAuctionBuyerRep(v string)`

SetAuctionBuyerRep sets AuctionBuyerRep field to given value.

### HasAuctionBuyerRep

`func (o *AeAdvisoryAccountAmsDetail) HasAuctionBuyerRep() bool`

HasAuctionBuyerRep returns a boolean if a field has been set.

### GetAutoimsId

`func (o *AeAdvisoryAccountAmsDetail) GetAutoimsId() float32`

GetAutoimsId returns the AutoimsId field if non-nil, zero value otherwise.

### GetAutoimsIdOk

`func (o *AeAdvisoryAccountAmsDetail) GetAutoimsIdOk() (*float32, bool)`

GetAutoimsIdOk returns a tuple with the AutoimsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoimsId

`func (o *AeAdvisoryAccountAmsDetail) SetAutoimsId(v float32)`

SetAutoimsId sets AutoimsId field to given value.

### HasAutoimsId

`func (o *AeAdvisoryAccountAmsDetail) HasAutoimsId() bool`

HasAutoimsId returns a boolean if a field has been set.

### GetAutoimsIdStr

`func (o *AeAdvisoryAccountAmsDetail) GetAutoimsIdStr() string`

GetAutoimsIdStr returns the AutoimsIdStr field if non-nil, zero value otherwise.

### GetAutoimsIdStrOk

`func (o *AeAdvisoryAccountAmsDetail) GetAutoimsIdStrOk() (*string, bool)`

GetAutoimsIdStrOk returns a tuple with the AutoimsIdStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoimsIdStr

`func (o *AeAdvisoryAccountAmsDetail) SetAutoimsIdStr(v string)`

SetAutoimsIdStr sets AutoimsIdStr field to given value.

### HasAutoimsIdStr

`func (o *AeAdvisoryAccountAmsDetail) HasAutoimsIdStr() bool`

HasAutoimsIdStr returns a boolean if a field has been set.

### GetPaymentSources

`func (o *AeAdvisoryAccountAmsDetail) GetPaymentSources() []CommonAmsPaymentSource`

GetPaymentSources returns the PaymentSources field if non-nil, zero value otherwise.

### GetPaymentSourcesOk

`func (o *AeAdvisoryAccountAmsDetail) GetPaymentSourcesOk() (*[]CommonAmsPaymentSource, bool)`

GetPaymentSourcesOk returns a tuple with the PaymentSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSources

`func (o *AeAdvisoryAccountAmsDetail) SetPaymentSources(v []CommonAmsPaymentSource)`

SetPaymentSources sets PaymentSources field to given value.

### HasPaymentSources

`func (o *AeAdvisoryAccountAmsDetail) HasPaymentSources() bool`

HasPaymentSources returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AeAdvisoryAccountAmsDetail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AeAdvisoryAccountAmsDetail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AeAdvisoryAccountAmsDetail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetInitiator

`func (o *AeAdvisoryAccountAmsDetail) GetInitiator() CommonInitiator`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *AeAdvisoryAccountAmsDetail) GetInitiatorOk() (*CommonInitiator, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *AeAdvisoryAccountAmsDetail) SetInitiator(v CommonInitiator)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *AeAdvisoryAccountAmsDetail) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


