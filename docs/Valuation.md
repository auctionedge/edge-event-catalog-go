# Valuation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to [**CommonCurrency**](CommonCurrency.md) |  | [optional] 
**Source** | Pointer to **string** | The source of the valuation | [optional] 

## Methods

### NewValuation

`func NewValuation() *Valuation`

NewValuation instantiates a new Valuation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValuationWithDefaults

`func NewValuationWithDefaults() *Valuation`

NewValuationWithDefaults instantiates a new Valuation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *Valuation) GetAmount() CommonCurrency`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Valuation) GetAmountOk() (*CommonCurrency, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Valuation) SetAmount(v CommonCurrency)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Valuation) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetSource

`func (o *Valuation) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Valuation) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Valuation) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *Valuation) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


