# CommonPhysicalLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | Name to be displayed by human | [optional] 
**Address1** | **string** | First line of address | 
**Address2** | Pointer to **string** | Second line of address | [optional] 
**City** | **string** | The city that the location is in | 
**State** | **string** | The state or province that the location is in | 
**PostalCode** | **string** | The postal code of this location | 
**Zip4** | Pointer to **string** | The 4-number suffix of the postal-code | [optional] 
**CountryIsoCode** | **string** | The ISO code for the country that this location is in | 
**RegionCode** | Pointer to **string** | Auction Edge region that this location is in | [optional] 
**TimeZone** | Pointer to **string** | The time zone that the location is in | [optional] 

## Methods

### NewCommonPhysicalLocation

`func NewCommonPhysicalLocation(address1 string, city string, state string, postalCode string, countryIsoCode string, ) *CommonPhysicalLocation`

NewCommonPhysicalLocation instantiates a new CommonPhysicalLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonPhysicalLocationWithDefaults

`func NewCommonPhysicalLocationWithDefaults() *CommonPhysicalLocation`

NewCommonPhysicalLocationWithDefaults instantiates a new CommonPhysicalLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *CommonPhysicalLocation) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CommonPhysicalLocation) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CommonPhysicalLocation) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CommonPhysicalLocation) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetAddress1

`func (o *CommonPhysicalLocation) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *CommonPhysicalLocation) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *CommonPhysicalLocation) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.


### GetAddress2

`func (o *CommonPhysicalLocation) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *CommonPhysicalLocation) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *CommonPhysicalLocation) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *CommonPhysicalLocation) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetCity

`func (o *CommonPhysicalLocation) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CommonPhysicalLocation) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CommonPhysicalLocation) SetCity(v string)`

SetCity sets City field to given value.


### GetState

`func (o *CommonPhysicalLocation) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CommonPhysicalLocation) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CommonPhysicalLocation) SetState(v string)`

SetState sets State field to given value.


### GetPostalCode

`func (o *CommonPhysicalLocation) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *CommonPhysicalLocation) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *CommonPhysicalLocation) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### GetZip4

`func (o *CommonPhysicalLocation) GetZip4() string`

GetZip4 returns the Zip4 field if non-nil, zero value otherwise.

### GetZip4Ok

`func (o *CommonPhysicalLocation) GetZip4Ok() (*string, bool)`

GetZip4Ok returns a tuple with the Zip4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip4

`func (o *CommonPhysicalLocation) SetZip4(v string)`

SetZip4 sets Zip4 field to given value.

### HasZip4

`func (o *CommonPhysicalLocation) HasZip4() bool`

HasZip4 returns a boolean if a field has been set.

### GetCountryIsoCode

`func (o *CommonPhysicalLocation) GetCountryIsoCode() string`

GetCountryIsoCode returns the CountryIsoCode field if non-nil, zero value otherwise.

### GetCountryIsoCodeOk

`func (o *CommonPhysicalLocation) GetCountryIsoCodeOk() (*string, bool)`

GetCountryIsoCodeOk returns a tuple with the CountryIsoCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIsoCode

`func (o *CommonPhysicalLocation) SetCountryIsoCode(v string)`

SetCountryIsoCode sets CountryIsoCode field to given value.


### GetRegionCode

`func (o *CommonPhysicalLocation) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *CommonPhysicalLocation) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *CommonPhysicalLocation) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *CommonPhysicalLocation) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetTimeZone

`func (o *CommonPhysicalLocation) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *CommonPhysicalLocation) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *CommonPhysicalLocation) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *CommonPhysicalLocation) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


