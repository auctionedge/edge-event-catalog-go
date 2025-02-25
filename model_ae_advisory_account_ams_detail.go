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

// checks if the AeAdvisoryAccountAmsDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AeAdvisoryAccountAmsDetail{}

// AeAdvisoryAccountAmsDetail Advisory Account event for AMS accounts
type AeAdvisoryAccountAmsDetail struct {
	// The account key of an AMS account (i.e. the account or dealer number)
	AccountKey string `json:"account-key"`
	// The Auction Access ID of the AMS account.
	AaId NullableString `json:"aa-id"`
	// Source's unique identifier for account
	AccountId *string `json:"account-id,omitempty"`
	// Auction Edge unique identifier for an auction.
	AuctionId string `json:"auction-id"`
	// Source's unique identifier for account
	Id *string `json:"id,omitempty"`
	// The date and time that the record was initially created in the AMS
	EstablishedAt *time.Time `json:"established-at,omitempty"`
	// The display name of the account
	DisplayName string `json:"display-name"`
	// The legal name of the account
	LegalName *string `json:"legal-name,omitempty"`
	// The 'doing business as' name of the account
	DbaName *string `json:"dba-name,omitempty"`
	// Primary email of dealership
	PrimaryEmail *string `json:"primary-email,omitempty"`
	// URI dealer website
	WebsiteUri *string `json:"website-uri,omitempty"`
	// Main phone number of account
	PrimaryPhoneNumber *string `json:"primary-phone-number,omitempty"`
	// Cell phone number on record at AMS for account
	CellPhoneNumber *string `json:"cell-phone-number,omitempty"`
	// Fax phone number
	PrimaryFaxNumber *string `json:"primary-fax-number,omitempty"`
	PhysicalLocation *CommonAmsPostalAddress `json:"physical-location,omitempty"`
	MailingLocation *CommonAmsPostalAddress `json:"mailing-location,omitempty"`
	// True if account is allowed to buy a vehicle, false if not
	CanBuy *bool `json:"can-buy,omitempty"`
	// True if account is allowed to sell a vehicle, false if not
	CanSell *bool `json:"can-sell,omitempty"`
	// True if account is active, false if not
	IsActive *bool `json:"is-active,omitempty"`
	// True if account is in violation of an auction rule such as expired license, false if not
	IsInViolation *bool `json:"is-in-violation,omitempty"`
	// True if account is tax exempt, false if not
	IsTaxExempt *bool `json:"is-tax-exempt,omitempty"`
	// Auction staff member that assists account selling vehicles
	AuctionSellerRep *string `json:"auction-seller-rep,omitempty"`
	// Auction staff member that assists account buying vehicles
	AuctionBuyerRep *string `json:"auction-buyer-rep,omitempty"`
	// AutoIMS id
	// Deprecated
	AutoimsId *float32 `json:"autoims-id,omitempty"`
	// AutoIMS id
	AutoimsIdStr *string `json:"autoims-id-str,omitempty"`
	// The payment sources for the account
	PaymentSources []CommonAmsPaymentSource `json:"payment-sources,omitempty"`
	// The time in ISO-8601 formatted at which the record was generated
	UpdatedAt time.Time `json:"updated-at"`
	Initiator *CommonInitiator `json:"initiator,omitempty"`
}

type _AeAdvisoryAccountAmsDetail AeAdvisoryAccountAmsDetail

// NewAeAdvisoryAccountAmsDetail instantiates a new AeAdvisoryAccountAmsDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAeAdvisoryAccountAmsDetail(accountKey string, aaId NullableString, auctionId string, displayName string, updatedAt time.Time) *AeAdvisoryAccountAmsDetail {
	this := AeAdvisoryAccountAmsDetail{}
	this.AccountKey = accountKey
	this.AaId = aaId
	this.AuctionId = auctionId
	this.DisplayName = displayName
	this.UpdatedAt = updatedAt
	return &this
}

// NewAeAdvisoryAccountAmsDetailWithDefaults instantiates a new AeAdvisoryAccountAmsDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAeAdvisoryAccountAmsDetailWithDefaults() *AeAdvisoryAccountAmsDetail {
	this := AeAdvisoryAccountAmsDetail{}
	return &this
}

// GetAccountKey returns the AccountKey field value
func (o *AeAdvisoryAccountAmsDetail) GetAccountKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountKey
}

// GetAccountKeyOk returns a tuple with the AccountKey field value
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAccountAmsDetail) GetAccountKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountKey, true
}

// SetAccountKey sets field value
func (o *AeAdvisoryAccountAmsDetail) SetAccountKey(v string) {
	o.AccountKey = v
}

// GetAaId returns the AaId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AeAdvisoryAccountAmsDetail) GetAaId() string {
	if o == nil || o.AaId.Get() == nil {
		var ret string
		return ret
	}

	return *o.AaId.Get()
}

// GetAaIdOk returns a tuple with the AaId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AeAdvisoryAccountAmsDetail) GetAaIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AaId.Get(), o.AaId.IsSet()
}

// SetAaId sets field value
func (o *AeAdvisoryAccountAmsDetail) SetAaId(v string) {
	o.AaId.Set(&v)
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AeAdvisoryAccountAmsDetail) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAccountAmsDetail) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AeAdvisoryAccountAmsDetail) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AeAdvisoryAccountAmsDetail) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAuctionId returns the AuctionId field value
func (o *AeAdvisoryAccountAmsDetail) GetAuctionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuctionId
}

// GetAuctionIdOk returns a tuple with the AuctionId field value
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAccountAmsDetail) GetAuctionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuctionId, true
}

// SetAuctionId sets field value
func (o *AeAdvisoryAccountAmsDetail) SetAuctionId(v string) {
	o.AuctionId = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AeAdvisoryAccountAmsDetail) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAccountAmsDetail) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AeAdvisoryAccountAmsDetail) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AeAdvisoryAccountAmsDetail) SetId(v string) {
	o.Id = &v
}

// GetEstablishedAt returns the EstablishedAt field value if set, zero value otherwise.
func (o *AeAdvisoryAccountAmsDetail) GetEstablishedAt() time.Time {
	if o == nil || IsNil(o.EstablishedAt) {
		var ret time.Time
		return ret
	}
	return *o.EstablishedAt
}

// GetEstablishedAtOk returns a tuple with the EstablishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAccountAmsDetail) GetEstablishedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EstablishedAt) {
		return nil, false
	}
	return o.EstablishedAt, true
}

// HasEstablishedAt returns a boolean if a field has been set.
func (o *AeAdvisoryAccountAmsDetail) HasEstablishedAt() bool {
	if o != nil && !IsNil(o.EstablishedAt) {
		return true
	}

	return false
}

// SetEstablishedAt gets a reference to the given time.Time and assigns it to the EstablishedAt field.
func (o *AeAdvisoryAccountAmsDetail) SetEstablishedAt(v time.Time) {
	o.EstablishedAt = &v
}

// GetDisplayName returns the DisplayName field value
func (o *AeAdvisoryAccountAmsDetail) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAccountAmsDetail) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *AeAdvisoryAccountAmsDetail) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetLegalName returns the LegalName field value if set, zero value otherwise.
func (o *AeAdvisoryAccountAmsDetail) GetLegalName() string {
	if o == nil || IsNil(o.LegalName) {
		var ret string
		return ret
	}
	return *o.LegalName
}

// GetLegalNameOk returns a tuple with the LegalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAccountAmsDetail) GetLegalNameOk() (*string, bool) {
	if o == nil || IsNil(o.LegalName) {
		return nil, false
	}
	return o.LegalName, true
}

// HasLegalName returns a boolean if a field has been set.
func (o *AeAdvisoryAccountAmsDetail) HasLegalName() bool {
	if o != nil && !IsNil(o.LegalName) {
		return true
	}

	return false
}

// SetLegalName gets a reference to the given string and assigns it to the LegalName field.
func (o *AeAdvisoryAccountAmsDetail) SetLegalName(v string) {
	o.LegalName = &v
}

// GetDbaName returns the DbaName field value if set, zero value otherwise.
func (o *AeAdvisoryAccountAmsDetail) GetDbaName() string {
	if o == nil || IsNil(o.DbaName) {
		var ret string
		return ret
	}
	return *o.DbaName
}

// GetDbaNameOk returns a tuple with the DbaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAccountAmsDetail) GetDbaNameOk() (*string, bool) {
	if o == nil || IsNil(o.DbaName) {
		return nil, false
	}
	return o.DbaName, true
}

// HasDbaName returns a boolean if a field has been set.
func (o *AeAdvisoryAccountAmsDetail) HasDbaName() bool {
	if o != nil && !IsNil(o.DbaName) {
		return true
	}

	return false
}

// SetDbaName gets a reference to the given string and assigns it to the DbaName field.
func (o *AeAdvisoryAccountAmsDetail) SetDbaName(v string) {
	o.DbaName = &v
}

// GetPrimaryEmail returns the PrimaryEmail field value if set, zero value otherwise.
func (o *AeAdvisoryAccountAmsDetail) GetPrimaryEmail() string {
	if o == nil || IsNil(o.PrimaryEmail) {
		var ret string
		return ret
	}
	return *o.PrimaryEmail
}

// GetPrimaryEmailOk returns a tuple with the PrimaryEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAccountAmsDetail) GetPrimaryEmailOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryEmail) {
		return nil, false
	}
	return o.PrimaryEmail, true
}

// HasPrimaryEmail returns a boolean if a field has been set.
func (o *AeAdvisoryAccountAmsDetail) HasPrimaryEmail() bool {
	if o != nil && !IsNil(o.PrimaryEmail) {
		return true
	}

	return false
}

// SetPrimaryEmail gets a reference to the given string and assigns it to the PrimaryEmail field.
func (o *AeAdvisoryAccountAmsDetail) SetPrimaryEmail(v string) {
	o.PrimaryEmail = &v
}

// GetWebsiteUri returns the WebsiteUri field value if set, zero value otherwise.
func (o *AeAdvisoryAccountAmsDetail) GetWebsiteUri() string {
	if o == nil || IsNil(o.WebsiteUri) {
		var ret string
		return ret
	}
	return *o.WebsiteUri
}

// GetWebsiteUriOk returns a tuple with the WebsiteUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAccountAmsDetail) GetWebsiteUriOk() (*string, bool) {
	if o == nil || IsNil(o.WebsiteUri) {
		return nil, false
	}
	return o.WebsiteUri, true
}

// HasWebsiteUri returns a boolean if a field has been set.
func (o *AeAdvisoryAccountAmsDetail) HasWebsiteUri() bool {
	if o != nil && !IsNil(o.WebsiteUri) {
		return true
	}

	return false
}

// SetWebsiteUri gets a reference to the given string and assigns it to the WebsiteUri field.
func (o *AeAdvisoryAccountAmsDetail) SetWebsiteUri(v string) {
	o.WebsiteUri = &v
}

// GetPrimaryPhoneNumber returns the PrimaryPhoneNumber field value if set, zero value otherwise.
func (o *AeAdvisoryAccountAmsDetail) GetPrimaryPhoneNumber() string {
	if o == nil || IsNil(o.PrimaryPhoneNumber) {
		var ret string
		return ret
	}
	return *o.PrimaryPhoneNumber
}

// GetPrimaryPhoneNumberOk returns a tuple with the PrimaryPhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAccountAmsDetail) GetPrimaryPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryPhoneNumber) {
		return nil, false
	}
	return o.PrimaryPhoneNumber, true
}

// HasPrimaryPhoneNumber returns a boolean if a field has been set.
func (o *AeAdvisoryAccountAmsDetail) HasPrimaryPhoneNumber() bool {
	if o != nil && !IsNil(o.PrimaryPhoneNumber) {
		return true
	}

	return false
}

// SetPrimaryPhoneNumber gets a reference to the given string and assigns it to the PrimaryPhoneNumber field.
func (o *AeAdvisoryAccountAmsDetail) SetPrimaryPhoneNumber(v string) {
	o.PrimaryPhoneNumber = &v
}

// GetCellPhoneNumber returns the CellPhoneNumber field value if set, zero value otherwise.
func (o *AeAdvisoryAccountAmsDetail) GetCellPhoneNumber() string {
	if o == nil || IsNil(o.CellPhoneNumber) {
		var ret string
		return ret
	}
	return *o.CellPhoneNumber
}

// GetCellPhoneNumberOk returns a tuple with the CellPhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAccountAmsDetail) GetCellPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.CellPhoneNumber) {
		return nil, false
	}
	return o.CellPhoneNumber, true
}

// HasCellPhoneNumber returns a boolean if a field has been set.
func (o *AeAdvisoryAccountAmsDetail) HasCellPhoneNumber() bool {
	if o != nil && !IsNil(o.CellPhoneNumber) {
		return true
	}

	return false
}

// SetCellPhoneNumber gets a reference to the given string and assigns it to the CellPhoneNumber field.
func (o *AeAdvisoryAccountAmsDetail) SetCellPhoneNumber(v string) {
	o.CellPhoneNumber = &v
}

// GetPrimaryFaxNumber returns the PrimaryFaxNumber field value if set, zero value otherwise.
func (o *AeAdvisoryAccountAmsDetail) GetPrimaryFaxNumber() string {
	if o == nil || IsNil(o.PrimaryFaxNumber) {
		var ret string
		return ret
	}
	return *o.PrimaryFaxNumber
}

// GetPrimaryFaxNumberOk returns a tuple with the PrimaryFaxNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAccountAmsDetail) GetPrimaryFaxNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryFaxNumber) {
		return nil, false
	}
	return o.PrimaryFaxNumber, true
}

// HasPrimaryFaxNumber returns a boolean if a field has been set.
func (o *AeAdvisoryAccountAmsDetail) HasPrimaryFaxNumber() bool {
	if o != nil && !IsNil(o.PrimaryFaxNumber) {
		return true
	}

	return false
}

// SetPrimaryFaxNumber gets a reference to the given string and assigns it to the PrimaryFaxNumber field.
func (o *AeAdvisoryAccountAmsDetail) SetPrimaryFaxNumber(v string) {
	o.PrimaryFaxNumber = &v
}

// GetPhysicalLocation returns the PhysicalLocation field value if set, zero value otherwise.
func (o *AeAdvisoryAccountAmsDetail) GetPhysicalLocation() CommonAmsPostalAddress {
	if o == nil || IsNil(o.PhysicalLocation) {
		var ret CommonAmsPostalAddress
		return ret
	}
	return *o.PhysicalLocation
}

// GetPhysicalLocationOk returns a tuple with the PhysicalLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAccountAmsDetail) GetPhysicalLocationOk() (*CommonAmsPostalAddress, bool) {
	if o == nil || IsNil(o.PhysicalLocation) {
		return nil, false
	}
	return o.PhysicalLocation, true
}

// HasPhysicalLocation returns a boolean if a field has been set.
func (o *AeAdvisoryAccountAmsDetail) HasPhysicalLocation() bool {
	if o != nil && !IsNil(o.PhysicalLocation) {
		return true
	}

	return false
}

// SetPhysicalLocation gets a reference to the given CommonAmsPostalAddress and assigns it to the PhysicalLocation field.
func (o *AeAdvisoryAccountAmsDetail) SetPhysicalLocation(v CommonAmsPostalAddress) {
	o.PhysicalLocation = &v
}

// GetMailingLocation returns the MailingLocation field value if set, zero value otherwise.
func (o *AeAdvisoryAccountAmsDetail) GetMailingLocation() CommonAmsPostalAddress {
	if o == nil || IsNil(o.MailingLocation) {
		var ret CommonAmsPostalAddress
		return ret
	}
	return *o.MailingLocation
}

// GetMailingLocationOk returns a tuple with the MailingLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAccountAmsDetail) GetMailingLocationOk() (*CommonAmsPostalAddress, bool) {
	if o == nil || IsNil(o.MailingLocation) {
		return nil, false
	}
	return o.MailingLocation, true
}

// HasMailingLocation returns a boolean if a field has been set.
func (o *AeAdvisoryAccountAmsDetail) HasMailingLocation() bool {
	if o != nil && !IsNil(o.MailingLocation) {
		return true
	}

	return false
}

// SetMailingLocation gets a reference to the given CommonAmsPostalAddress and assigns it to the MailingLocation field.
func (o *AeAdvisoryAccountAmsDetail) SetMailingLocation(v CommonAmsPostalAddress) {
	o.MailingLocation = &v
}

// GetCanBuy returns the CanBuy field value if set, zero value otherwise.
func (o *AeAdvisoryAccountAmsDetail) GetCanBuy() bool {
	if o == nil || IsNil(o.CanBuy) {
		var ret bool
		return ret
	}
	return *o.CanBuy
}

// GetCanBuyOk returns a tuple with the CanBuy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAccountAmsDetail) GetCanBuyOk() (*bool, bool) {
	if o == nil || IsNil(o.CanBuy) {
		return nil, false
	}
	return o.CanBuy, true
}

// HasCanBuy returns a boolean if a field has been set.
func (o *AeAdvisoryAccountAmsDetail) HasCanBuy() bool {
	if o != nil && !IsNil(o.CanBuy) {
		return true
	}

	return false
}

// SetCanBuy gets a reference to the given bool and assigns it to the CanBuy field.
func (o *AeAdvisoryAccountAmsDetail) SetCanBuy(v bool) {
	o.CanBuy = &v
}

// GetCanSell returns the CanSell field value if set, zero value otherwise.
func (o *AeAdvisoryAccountAmsDetail) GetCanSell() bool {
	if o == nil || IsNil(o.CanSell) {
		var ret bool
		return ret
	}
	return *o.CanSell
}

// GetCanSellOk returns a tuple with the CanSell field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAccountAmsDetail) GetCanSellOk() (*bool, bool) {
	if o == nil || IsNil(o.CanSell) {
		return nil, false
	}
	return o.CanSell, true
}

// HasCanSell returns a boolean if a field has been set.
func (o *AeAdvisoryAccountAmsDetail) HasCanSell() bool {
	if o != nil && !IsNil(o.CanSell) {
		return true
	}

	return false
}

// SetCanSell gets a reference to the given bool and assigns it to the CanSell field.
func (o *AeAdvisoryAccountAmsDetail) SetCanSell(v bool) {
	o.CanSell = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *AeAdvisoryAccountAmsDetail) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAccountAmsDetail) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *AeAdvisoryAccountAmsDetail) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *AeAdvisoryAccountAmsDetail) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetIsInViolation returns the IsInViolation field value if set, zero value otherwise.
func (o *AeAdvisoryAccountAmsDetail) GetIsInViolation() bool {
	if o == nil || IsNil(o.IsInViolation) {
		var ret bool
		return ret
	}
	return *o.IsInViolation
}

// GetIsInViolationOk returns a tuple with the IsInViolation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAccountAmsDetail) GetIsInViolationOk() (*bool, bool) {
	if o == nil || IsNil(o.IsInViolation) {
		return nil, false
	}
	return o.IsInViolation, true
}

// HasIsInViolation returns a boolean if a field has been set.
func (o *AeAdvisoryAccountAmsDetail) HasIsInViolation() bool {
	if o != nil && !IsNil(o.IsInViolation) {
		return true
	}

	return false
}

// SetIsInViolation gets a reference to the given bool and assigns it to the IsInViolation field.
func (o *AeAdvisoryAccountAmsDetail) SetIsInViolation(v bool) {
	o.IsInViolation = &v
}

// GetIsTaxExempt returns the IsTaxExempt field value if set, zero value otherwise.
func (o *AeAdvisoryAccountAmsDetail) GetIsTaxExempt() bool {
	if o == nil || IsNil(o.IsTaxExempt) {
		var ret bool
		return ret
	}
	return *o.IsTaxExempt
}

// GetIsTaxExemptOk returns a tuple with the IsTaxExempt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAccountAmsDetail) GetIsTaxExemptOk() (*bool, bool) {
	if o == nil || IsNil(o.IsTaxExempt) {
		return nil, false
	}
	return o.IsTaxExempt, true
}

// HasIsTaxExempt returns a boolean if a field has been set.
func (o *AeAdvisoryAccountAmsDetail) HasIsTaxExempt() bool {
	if o != nil && !IsNil(o.IsTaxExempt) {
		return true
	}

	return false
}

// SetIsTaxExempt gets a reference to the given bool and assigns it to the IsTaxExempt field.
func (o *AeAdvisoryAccountAmsDetail) SetIsTaxExempt(v bool) {
	o.IsTaxExempt = &v
}

// GetAuctionSellerRep returns the AuctionSellerRep field value if set, zero value otherwise.
func (o *AeAdvisoryAccountAmsDetail) GetAuctionSellerRep() string {
	if o == nil || IsNil(o.AuctionSellerRep) {
		var ret string
		return ret
	}
	return *o.AuctionSellerRep
}

// GetAuctionSellerRepOk returns a tuple with the AuctionSellerRep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAccountAmsDetail) GetAuctionSellerRepOk() (*string, bool) {
	if o == nil || IsNil(o.AuctionSellerRep) {
		return nil, false
	}
	return o.AuctionSellerRep, true
}

// HasAuctionSellerRep returns a boolean if a field has been set.
func (o *AeAdvisoryAccountAmsDetail) HasAuctionSellerRep() bool {
	if o != nil && !IsNil(o.AuctionSellerRep) {
		return true
	}

	return false
}

// SetAuctionSellerRep gets a reference to the given string and assigns it to the AuctionSellerRep field.
func (o *AeAdvisoryAccountAmsDetail) SetAuctionSellerRep(v string) {
	o.AuctionSellerRep = &v
}

// GetAuctionBuyerRep returns the AuctionBuyerRep field value if set, zero value otherwise.
func (o *AeAdvisoryAccountAmsDetail) GetAuctionBuyerRep() string {
	if o == nil || IsNil(o.AuctionBuyerRep) {
		var ret string
		return ret
	}
	return *o.AuctionBuyerRep
}

// GetAuctionBuyerRepOk returns a tuple with the AuctionBuyerRep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAccountAmsDetail) GetAuctionBuyerRepOk() (*string, bool) {
	if o == nil || IsNil(o.AuctionBuyerRep) {
		return nil, false
	}
	return o.AuctionBuyerRep, true
}

// HasAuctionBuyerRep returns a boolean if a field has been set.
func (o *AeAdvisoryAccountAmsDetail) HasAuctionBuyerRep() bool {
	if o != nil && !IsNil(o.AuctionBuyerRep) {
		return true
	}

	return false
}

// SetAuctionBuyerRep gets a reference to the given string and assigns it to the AuctionBuyerRep field.
func (o *AeAdvisoryAccountAmsDetail) SetAuctionBuyerRep(v string) {
	o.AuctionBuyerRep = &v
}

// GetAutoimsId returns the AutoimsId field value if set, zero value otherwise.
// Deprecated
func (o *AeAdvisoryAccountAmsDetail) GetAutoimsId() float32 {
	if o == nil || IsNil(o.AutoimsId) {
		var ret float32
		return ret
	}
	return *o.AutoimsId
}

// GetAutoimsIdOk returns a tuple with the AutoimsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AeAdvisoryAccountAmsDetail) GetAutoimsIdOk() (*float32, bool) {
	if o == nil || IsNil(o.AutoimsId) {
		return nil, false
	}
	return o.AutoimsId, true
}

// HasAutoimsId returns a boolean if a field has been set.
func (o *AeAdvisoryAccountAmsDetail) HasAutoimsId() bool {
	if o != nil && !IsNil(o.AutoimsId) {
		return true
	}

	return false
}

// SetAutoimsId gets a reference to the given float32 and assigns it to the AutoimsId field.
// Deprecated
func (o *AeAdvisoryAccountAmsDetail) SetAutoimsId(v float32) {
	o.AutoimsId = &v
}

// GetAutoimsIdStr returns the AutoimsIdStr field value if set, zero value otherwise.
func (o *AeAdvisoryAccountAmsDetail) GetAutoimsIdStr() string {
	if o == nil || IsNil(o.AutoimsIdStr) {
		var ret string
		return ret
	}
	return *o.AutoimsIdStr
}

// GetAutoimsIdStrOk returns a tuple with the AutoimsIdStr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAccountAmsDetail) GetAutoimsIdStrOk() (*string, bool) {
	if o == nil || IsNil(o.AutoimsIdStr) {
		return nil, false
	}
	return o.AutoimsIdStr, true
}

// HasAutoimsIdStr returns a boolean if a field has been set.
func (o *AeAdvisoryAccountAmsDetail) HasAutoimsIdStr() bool {
	if o != nil && !IsNil(o.AutoimsIdStr) {
		return true
	}

	return false
}

// SetAutoimsIdStr gets a reference to the given string and assigns it to the AutoimsIdStr field.
func (o *AeAdvisoryAccountAmsDetail) SetAutoimsIdStr(v string) {
	o.AutoimsIdStr = &v
}

// GetPaymentSources returns the PaymentSources field value if set, zero value otherwise.
func (o *AeAdvisoryAccountAmsDetail) GetPaymentSources() []CommonAmsPaymentSource {
	if o == nil || IsNil(o.PaymentSources) {
		var ret []CommonAmsPaymentSource
		return ret
	}
	return o.PaymentSources
}

// GetPaymentSourcesOk returns a tuple with the PaymentSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAccountAmsDetail) GetPaymentSourcesOk() ([]CommonAmsPaymentSource, bool) {
	if o == nil || IsNil(o.PaymentSources) {
		return nil, false
	}
	return o.PaymentSources, true
}

// HasPaymentSources returns a boolean if a field has been set.
func (o *AeAdvisoryAccountAmsDetail) HasPaymentSources() bool {
	if o != nil && !IsNil(o.PaymentSources) {
		return true
	}

	return false
}

// SetPaymentSources gets a reference to the given []CommonAmsPaymentSource and assigns it to the PaymentSources field.
func (o *AeAdvisoryAccountAmsDetail) SetPaymentSources(v []CommonAmsPaymentSource) {
	o.PaymentSources = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *AeAdvisoryAccountAmsDetail) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAccountAmsDetail) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *AeAdvisoryAccountAmsDetail) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetInitiator returns the Initiator field value if set, zero value otherwise.
func (o *AeAdvisoryAccountAmsDetail) GetInitiator() CommonInitiator {
	if o == nil || IsNil(o.Initiator) {
		var ret CommonInitiator
		return ret
	}
	return *o.Initiator
}

// GetInitiatorOk returns a tuple with the Initiator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AeAdvisoryAccountAmsDetail) GetInitiatorOk() (*CommonInitiator, bool) {
	if o == nil || IsNil(o.Initiator) {
		return nil, false
	}
	return o.Initiator, true
}

// HasInitiator returns a boolean if a field has been set.
func (o *AeAdvisoryAccountAmsDetail) HasInitiator() bool {
	if o != nil && !IsNil(o.Initiator) {
		return true
	}

	return false
}

// SetInitiator gets a reference to the given CommonInitiator and assigns it to the Initiator field.
func (o *AeAdvisoryAccountAmsDetail) SetInitiator(v CommonInitiator) {
	o.Initiator = &v
}

func (o AeAdvisoryAccountAmsDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AeAdvisoryAccountAmsDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account-key"] = o.AccountKey
	toSerialize["aa-id"] = o.AaId.Get()
	if !IsNil(o.AccountId) {
		toSerialize["account-id"] = o.AccountId
	}
	toSerialize["auction-id"] = o.AuctionId
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.EstablishedAt) {
		toSerialize["established-at"] = o.EstablishedAt
	}
	toSerialize["display-name"] = o.DisplayName
	if !IsNil(o.LegalName) {
		toSerialize["legal-name"] = o.LegalName
	}
	if !IsNil(o.DbaName) {
		toSerialize["dba-name"] = o.DbaName
	}
	if !IsNil(o.PrimaryEmail) {
		toSerialize["primary-email"] = o.PrimaryEmail
	}
	if !IsNil(o.WebsiteUri) {
		toSerialize["website-uri"] = o.WebsiteUri
	}
	if !IsNil(o.PrimaryPhoneNumber) {
		toSerialize["primary-phone-number"] = o.PrimaryPhoneNumber
	}
	if !IsNil(o.CellPhoneNumber) {
		toSerialize["cell-phone-number"] = o.CellPhoneNumber
	}
	if !IsNil(o.PrimaryFaxNumber) {
		toSerialize["primary-fax-number"] = o.PrimaryFaxNumber
	}
	if !IsNil(o.PhysicalLocation) {
		toSerialize["physical-location"] = o.PhysicalLocation
	}
	if !IsNil(o.MailingLocation) {
		toSerialize["mailing-location"] = o.MailingLocation
	}
	if !IsNil(o.CanBuy) {
		toSerialize["can-buy"] = o.CanBuy
	}
	if !IsNil(o.CanSell) {
		toSerialize["can-sell"] = o.CanSell
	}
	if !IsNil(o.IsActive) {
		toSerialize["is-active"] = o.IsActive
	}
	if !IsNil(o.IsInViolation) {
		toSerialize["is-in-violation"] = o.IsInViolation
	}
	if !IsNil(o.IsTaxExempt) {
		toSerialize["is-tax-exempt"] = o.IsTaxExempt
	}
	if !IsNil(o.AuctionSellerRep) {
		toSerialize["auction-seller-rep"] = o.AuctionSellerRep
	}
	if !IsNil(o.AuctionBuyerRep) {
		toSerialize["auction-buyer-rep"] = o.AuctionBuyerRep
	}
	if !IsNil(o.AutoimsId) {
		toSerialize["autoims-id"] = o.AutoimsId
	}
	if !IsNil(o.AutoimsIdStr) {
		toSerialize["autoims-id-str"] = o.AutoimsIdStr
	}
	if !IsNil(o.PaymentSources) {
		toSerialize["payment-sources"] = o.PaymentSources
	}
	toSerialize["updated-at"] = o.UpdatedAt
	if !IsNil(o.Initiator) {
		toSerialize["initiator"] = o.Initiator
	}
	return toSerialize, nil
}

func (o *AeAdvisoryAccountAmsDetail) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account-key",
		"aa-id",
		"auction-id",
		"display-name",
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

	varAeAdvisoryAccountAmsDetail := _AeAdvisoryAccountAmsDetail{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAeAdvisoryAccountAmsDetail)

	if err != nil {
		return err
	}

	*o = AeAdvisoryAccountAmsDetail(varAeAdvisoryAccountAmsDetail)

	return err
}

type NullableAeAdvisoryAccountAmsDetail struct {
	value *AeAdvisoryAccountAmsDetail
	isSet bool
}

func (v NullableAeAdvisoryAccountAmsDetail) Get() *AeAdvisoryAccountAmsDetail {
	return v.value
}

func (v *NullableAeAdvisoryAccountAmsDetail) Set(val *AeAdvisoryAccountAmsDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableAeAdvisoryAccountAmsDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableAeAdvisoryAccountAmsDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAeAdvisoryAccountAmsDetail(val *AeAdvisoryAccountAmsDetail) *NullableAeAdvisoryAccountAmsDetail {
	return &NullableAeAdvisoryAccountAmsDetail{value: val, isSet: true}
}

func (v NullableAeAdvisoryAccountAmsDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAeAdvisoryAccountAmsDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


