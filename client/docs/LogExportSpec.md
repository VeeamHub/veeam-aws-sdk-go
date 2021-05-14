# LogExportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExportLogsType** | Pointer to **string** |  | [optional] 
**Days** | Pointer to **int32** |  | [optional] 
**FromDateUtc** | Pointer to **time.Time** |  | [optional] 
**ToDateUtc** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewLogExportSpec

`func NewLogExportSpec() *LogExportSpec`

NewLogExportSpec instantiates a new LogExportSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogExportSpecWithDefaults

`func NewLogExportSpecWithDefaults() *LogExportSpec`

NewLogExportSpecWithDefaults instantiates a new LogExportSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExportLogsType

`func (o *LogExportSpec) GetExportLogsType() string`

GetExportLogsType returns the ExportLogsType field if non-nil, zero value otherwise.

### GetExportLogsTypeOk

`func (o *LogExportSpec) GetExportLogsTypeOk() (*string, bool)`

GetExportLogsTypeOk returns a tuple with the ExportLogsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportLogsType

`func (o *LogExportSpec) SetExportLogsType(v string)`

SetExportLogsType sets ExportLogsType field to given value.

### HasExportLogsType

`func (o *LogExportSpec) HasExportLogsType() bool`

HasExportLogsType returns a boolean if a field has been set.

### GetDays

`func (o *LogExportSpec) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *LogExportSpec) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *LogExportSpec) SetDays(v int32)`

SetDays sets Days field to given value.

### HasDays

`func (o *LogExportSpec) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetFromDateUtc

`func (o *LogExportSpec) GetFromDateUtc() time.Time`

GetFromDateUtc returns the FromDateUtc field if non-nil, zero value otherwise.

### GetFromDateUtcOk

`func (o *LogExportSpec) GetFromDateUtcOk() (*time.Time, bool)`

GetFromDateUtcOk returns a tuple with the FromDateUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDateUtc

`func (o *LogExportSpec) SetFromDateUtc(v time.Time)`

SetFromDateUtc sets FromDateUtc field to given value.

### HasFromDateUtc

`func (o *LogExportSpec) HasFromDateUtc() bool`

HasFromDateUtc returns a boolean if a field has been set.

### GetToDateUtc

`func (o *LogExportSpec) GetToDateUtc() time.Time`

GetToDateUtc returns the ToDateUtc field if non-nil, zero value otherwise.

### GetToDateUtcOk

`func (o *LogExportSpec) GetToDateUtcOk() (*time.Time, bool)`

GetToDateUtcOk returns a tuple with the ToDateUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDateUtc

`func (o *LogExportSpec) SetToDateUtc(v time.Time)`

SetToDateUtc sets ToDateUtc field to given value.

### HasToDateUtc

`func (o *LogExportSpec) HasToDateUtc() bool`

HasToDateUtc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


