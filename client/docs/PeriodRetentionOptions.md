# PeriodRetentionOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**RetentionTypes**](RetentionTypes.md) |  | 
**Count** | **int32** |  | 

## Methods

### NewPeriodRetentionOptions

`func NewPeriodRetentionOptions(type_ RetentionTypes, count int32, ) *PeriodRetentionOptions`

NewPeriodRetentionOptions instantiates a new PeriodRetentionOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeriodRetentionOptionsWithDefaults

`func NewPeriodRetentionOptionsWithDefaults() *PeriodRetentionOptions`

NewPeriodRetentionOptionsWithDefaults instantiates a new PeriodRetentionOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PeriodRetentionOptions) GetType() RetentionTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PeriodRetentionOptions) GetTypeOk() (*RetentionTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PeriodRetentionOptions) SetType(v RetentionTypes)`

SetType sets Type field to given value.


### GetCount

`func (o *PeriodRetentionOptions) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PeriodRetentionOptions) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PeriodRetentionOptions) SetCount(v int32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


