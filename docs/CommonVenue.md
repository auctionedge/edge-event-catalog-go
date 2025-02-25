# CommonVenue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | The name of this specific venue for the organization | 
**Organization** | **string** | Auction Edge unique identifier for an auction. | 
**Location** | [**CommonPhysicalLocation**](CommonPhysicalLocation.md) |  | 
**VenueType** | **string** | The type of the venue | 

## Methods

### NewCommonVenue

`func NewCommonVenue(displayName string, organization string, location CommonPhysicalLocation, venueType string, ) *CommonVenue`

NewCommonVenue instantiates a new CommonVenue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonVenueWithDefaults

`func NewCommonVenueWithDefaults() *CommonVenue`

NewCommonVenueWithDefaults instantiates a new CommonVenue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *CommonVenue) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CommonVenue) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CommonVenue) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetOrganization

`func (o *CommonVenue) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *CommonVenue) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *CommonVenue) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### GetLocation

`func (o *CommonVenue) GetLocation() CommonPhysicalLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CommonVenue) GetLocationOk() (*CommonPhysicalLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CommonVenue) SetLocation(v CommonPhysicalLocation)`

SetLocation sets Location field to given value.


### GetVenueType

`func (o *CommonVenue) GetVenueType() string`

GetVenueType returns the VenueType field if non-nil, zero value otherwise.

### GetVenueTypeOk

`func (o *CommonVenue) GetVenueTypeOk() (*string, bool)`

GetVenueTypeOk returns a tuple with the VenueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenueType

`func (o *CommonVenue) SetVenueType(v string)`

SetVenueType sets VenueType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


