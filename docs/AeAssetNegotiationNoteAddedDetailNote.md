# AeAssetNegotiationNoteAddedDetailNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | unique identifier of the note | 
**NotedBy** | **string** | User that added the note | 
**NotedAt** | **time.Time** | the time the note was taken | 
**Source** | **string** | the source system that created the note | 
**Subject** | **string** |  | 
**Body** | **string** | The body of the note | 

## Methods

### NewAeAssetNegotiationNoteAddedDetailNote

`func NewAeAssetNegotiationNoteAddedDetailNote(id string, notedBy string, notedAt time.Time, source string, subject string, body string, ) *AeAssetNegotiationNoteAddedDetailNote`

NewAeAssetNegotiationNoteAddedDetailNote instantiates a new AeAssetNegotiationNoteAddedDetailNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAeAssetNegotiationNoteAddedDetailNoteWithDefaults

`func NewAeAssetNegotiationNoteAddedDetailNoteWithDefaults() *AeAssetNegotiationNoteAddedDetailNote`

NewAeAssetNegotiationNoteAddedDetailNoteWithDefaults instantiates a new AeAssetNegotiationNoteAddedDetailNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AeAssetNegotiationNoteAddedDetailNote) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AeAssetNegotiationNoteAddedDetailNote) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AeAssetNegotiationNoteAddedDetailNote) SetId(v string)`

SetId sets Id field to given value.


### GetNotedBy

`func (o *AeAssetNegotiationNoteAddedDetailNote) GetNotedBy() string`

GetNotedBy returns the NotedBy field if non-nil, zero value otherwise.

### GetNotedByOk

`func (o *AeAssetNegotiationNoteAddedDetailNote) GetNotedByOk() (*string, bool)`

GetNotedByOk returns a tuple with the NotedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotedBy

`func (o *AeAssetNegotiationNoteAddedDetailNote) SetNotedBy(v string)`

SetNotedBy sets NotedBy field to given value.


### GetNotedAt

`func (o *AeAssetNegotiationNoteAddedDetailNote) GetNotedAt() time.Time`

GetNotedAt returns the NotedAt field if non-nil, zero value otherwise.

### GetNotedAtOk

`func (o *AeAssetNegotiationNoteAddedDetailNote) GetNotedAtOk() (*time.Time, bool)`

GetNotedAtOk returns a tuple with the NotedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotedAt

`func (o *AeAssetNegotiationNoteAddedDetailNote) SetNotedAt(v time.Time)`

SetNotedAt sets NotedAt field to given value.


### GetSource

`func (o *AeAssetNegotiationNoteAddedDetailNote) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AeAssetNegotiationNoteAddedDetailNote) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AeAssetNegotiationNoteAddedDetailNote) SetSource(v string)`

SetSource sets Source field to given value.


### GetSubject

`func (o *AeAssetNegotiationNoteAddedDetailNote) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *AeAssetNegotiationNoteAddedDetailNote) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *AeAssetNegotiationNoteAddedDetailNote) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetBody

`func (o *AeAssetNegotiationNoteAddedDetailNote) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *AeAssetNegotiationNoteAddedDetailNote) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *AeAssetNegotiationNoteAddedDetailNote) SetBody(v string)`

SetBody sets Body field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


