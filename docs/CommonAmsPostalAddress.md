# CommonAmsPostalAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address1** | **string** | Address line 1 of location | 
**Address2** | Pointer to **string** | Address line 2 of location | [optional] 
**Address3** | Pointer to **string** | Address line 2 of location | [optional] 
**City** | **string** | City of location | 
**State** | **string** | State of location | 
**PostalCode** | **string** | Postal code of location | 

## Methods

### NewCommonAmsPostalAddress

`func NewCommonAmsPostalAddress(address1 string, city string, state string, postalCode string, ) *CommonAmsPostalAddress`

NewCommonAmsPostalAddress instantiates a new CommonAmsPostalAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonAmsPostalAddressWithDefaults

`func NewCommonAmsPostalAddressWithDefaults() *CommonAmsPostalAddress`

NewCommonAmsPostalAddressWithDefaults instantiates a new CommonAmsPostalAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress1

`func (o *CommonAmsPostalAddress) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *CommonAmsPostalAddress) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *CommonAmsPostalAddress) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.


### GetAddress2

`func (o *CommonAmsPostalAddress) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *CommonAmsPostalAddress) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *CommonAmsPostalAddress) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *CommonAmsPostalAddress) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetAddress3

`func (o *CommonAmsPostalAddress) GetAddress3() string`

GetAddress3 returns the Address3 field if non-nil, zero value otherwise.

### GetAddress3Ok

`func (o *CommonAmsPostalAddress) GetAddress3Ok() (*string, bool)`

GetAddress3Ok returns a tuple with the Address3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress3

`func (o *CommonAmsPostalAddress) SetAddress3(v string)`

SetAddress3 sets Address3 field to given value.

### HasAddress3

`func (o *CommonAmsPostalAddress) HasAddress3() bool`

HasAddress3 returns a boolean if a field has been set.

### GetCity

`func (o *CommonAmsPostalAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CommonAmsPostalAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CommonAmsPostalAddress) SetCity(v string)`

SetCity sets City field to given value.


### GetState

`func (o *CommonAmsPostalAddress) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CommonAmsPostalAddress) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CommonAmsPostalAddress) SetState(v string)`

SetState sets State field to given value.


### GetPostalCode

`func (o *CommonAmsPostalAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *CommonAmsPostalAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *CommonAmsPostalAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


