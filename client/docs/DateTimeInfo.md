# DateTimeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UtcDateTime** | **time.Time** |  | 

## Methods

### NewDateTimeInfo

`func NewDateTimeInfo(utcDateTime time.Time, ) *DateTimeInfo`

NewDateTimeInfo instantiates a new DateTimeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDateTimeInfoWithDefaults

`func NewDateTimeInfoWithDefaults() *DateTimeInfo`

NewDateTimeInfoWithDefaults instantiates a new DateTimeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUtcDateTime

`func (o *DateTimeInfo) GetUtcDateTime() time.Time`

GetUtcDateTime returns the UtcDateTime field if non-nil, zero value otherwise.

### GetUtcDateTimeOk

`func (o *DateTimeInfo) GetUtcDateTimeOk() (*time.Time, bool)`

GetUtcDateTimeOk returns a tuple with the UtcDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtcDateTime

`func (o *DateTimeInfo) SetUtcDateTime(v time.Time)`

SetUtcDateTime sets UtcDateTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


