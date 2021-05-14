# SummaryStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlarmsCount** | **int32** |  | 
**AlarmsCountChangesTrend** | [**SummaryChangesTrends**](SummaryChangesTrends.md) |  | 
**InstancesCount** | **int32** |  | 
**InstancesCountChangesTrend** | [**SummaryChangesTrends**](SummaryChangesTrends.md) |  | 
**ProtectedInstancesCount** | **int32** |  | 
**ProtectedInstancesCountChangesTrend** | [**SummaryChangesTrends**](SummaryChangesTrends.md) |  | 
**PoliciesCount** | **int32** |  | 
**PoliciesCountChangesTrend** | [**SummaryChangesTrends**](SummaryChangesTrends.md) |  | 
**RepositoriesCount** | **int32** |  | 
**RepositoriesCountChangesTrend** | [**SummaryChangesTrends**](SummaryChangesTrends.md) |  | 

## Methods

### NewSummaryStatistics

`func NewSummaryStatistics(alarmsCount int32, alarmsCountChangesTrend SummaryChangesTrends, instancesCount int32, instancesCountChangesTrend SummaryChangesTrends, protectedInstancesCount int32, protectedInstancesCountChangesTrend SummaryChangesTrends, policiesCount int32, policiesCountChangesTrend SummaryChangesTrends, repositoriesCount int32, repositoriesCountChangesTrend SummaryChangesTrends, ) *SummaryStatistics`

NewSummaryStatistics instantiates a new SummaryStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummaryStatisticsWithDefaults

`func NewSummaryStatisticsWithDefaults() *SummaryStatistics`

NewSummaryStatisticsWithDefaults instantiates a new SummaryStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlarmsCount

`func (o *SummaryStatistics) GetAlarmsCount() int32`

GetAlarmsCount returns the AlarmsCount field if non-nil, zero value otherwise.

### GetAlarmsCountOk

`func (o *SummaryStatistics) GetAlarmsCountOk() (*int32, bool)`

GetAlarmsCountOk returns a tuple with the AlarmsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmsCount

`func (o *SummaryStatistics) SetAlarmsCount(v int32)`

SetAlarmsCount sets AlarmsCount field to given value.


### GetAlarmsCountChangesTrend

`func (o *SummaryStatistics) GetAlarmsCountChangesTrend() SummaryChangesTrends`

GetAlarmsCountChangesTrend returns the AlarmsCountChangesTrend field if non-nil, zero value otherwise.

### GetAlarmsCountChangesTrendOk

`func (o *SummaryStatistics) GetAlarmsCountChangesTrendOk() (*SummaryChangesTrends, bool)`

GetAlarmsCountChangesTrendOk returns a tuple with the AlarmsCountChangesTrend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmsCountChangesTrend

`func (o *SummaryStatistics) SetAlarmsCountChangesTrend(v SummaryChangesTrends)`

SetAlarmsCountChangesTrend sets AlarmsCountChangesTrend field to given value.


### GetInstancesCount

`func (o *SummaryStatistics) GetInstancesCount() int32`

GetInstancesCount returns the InstancesCount field if non-nil, zero value otherwise.

### GetInstancesCountOk

`func (o *SummaryStatistics) GetInstancesCountOk() (*int32, bool)`

GetInstancesCountOk returns a tuple with the InstancesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstancesCount

`func (o *SummaryStatistics) SetInstancesCount(v int32)`

SetInstancesCount sets InstancesCount field to given value.


### GetInstancesCountChangesTrend

`func (o *SummaryStatistics) GetInstancesCountChangesTrend() SummaryChangesTrends`

GetInstancesCountChangesTrend returns the InstancesCountChangesTrend field if non-nil, zero value otherwise.

### GetInstancesCountChangesTrendOk

`func (o *SummaryStatistics) GetInstancesCountChangesTrendOk() (*SummaryChangesTrends, bool)`

GetInstancesCountChangesTrendOk returns a tuple with the InstancesCountChangesTrend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstancesCountChangesTrend

`func (o *SummaryStatistics) SetInstancesCountChangesTrend(v SummaryChangesTrends)`

SetInstancesCountChangesTrend sets InstancesCountChangesTrend field to given value.


### GetProtectedInstancesCount

`func (o *SummaryStatistics) GetProtectedInstancesCount() int32`

GetProtectedInstancesCount returns the ProtectedInstancesCount field if non-nil, zero value otherwise.

### GetProtectedInstancesCountOk

`func (o *SummaryStatistics) GetProtectedInstancesCountOk() (*int32, bool)`

GetProtectedInstancesCountOk returns a tuple with the ProtectedInstancesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedInstancesCount

`func (o *SummaryStatistics) SetProtectedInstancesCount(v int32)`

SetProtectedInstancesCount sets ProtectedInstancesCount field to given value.


### GetProtectedInstancesCountChangesTrend

`func (o *SummaryStatistics) GetProtectedInstancesCountChangesTrend() SummaryChangesTrends`

GetProtectedInstancesCountChangesTrend returns the ProtectedInstancesCountChangesTrend field if non-nil, zero value otherwise.

### GetProtectedInstancesCountChangesTrendOk

`func (o *SummaryStatistics) GetProtectedInstancesCountChangesTrendOk() (*SummaryChangesTrends, bool)`

GetProtectedInstancesCountChangesTrendOk returns a tuple with the ProtectedInstancesCountChangesTrend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedInstancesCountChangesTrend

`func (o *SummaryStatistics) SetProtectedInstancesCountChangesTrend(v SummaryChangesTrends)`

SetProtectedInstancesCountChangesTrend sets ProtectedInstancesCountChangesTrend field to given value.


### GetPoliciesCount

`func (o *SummaryStatistics) GetPoliciesCount() int32`

GetPoliciesCount returns the PoliciesCount field if non-nil, zero value otherwise.

### GetPoliciesCountOk

`func (o *SummaryStatistics) GetPoliciesCountOk() (*int32, bool)`

GetPoliciesCountOk returns a tuple with the PoliciesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoliciesCount

`func (o *SummaryStatistics) SetPoliciesCount(v int32)`

SetPoliciesCount sets PoliciesCount field to given value.


### GetPoliciesCountChangesTrend

`func (o *SummaryStatistics) GetPoliciesCountChangesTrend() SummaryChangesTrends`

GetPoliciesCountChangesTrend returns the PoliciesCountChangesTrend field if non-nil, zero value otherwise.

### GetPoliciesCountChangesTrendOk

`func (o *SummaryStatistics) GetPoliciesCountChangesTrendOk() (*SummaryChangesTrends, bool)`

GetPoliciesCountChangesTrendOk returns a tuple with the PoliciesCountChangesTrend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoliciesCountChangesTrend

`func (o *SummaryStatistics) SetPoliciesCountChangesTrend(v SummaryChangesTrends)`

SetPoliciesCountChangesTrend sets PoliciesCountChangesTrend field to given value.


### GetRepositoriesCount

`func (o *SummaryStatistics) GetRepositoriesCount() int32`

GetRepositoriesCount returns the RepositoriesCount field if non-nil, zero value otherwise.

### GetRepositoriesCountOk

`func (o *SummaryStatistics) GetRepositoriesCountOk() (*int32, bool)`

GetRepositoriesCountOk returns a tuple with the RepositoriesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoriesCount

`func (o *SummaryStatistics) SetRepositoriesCount(v int32)`

SetRepositoriesCount sets RepositoriesCount field to given value.


### GetRepositoriesCountChangesTrend

`func (o *SummaryStatistics) GetRepositoriesCountChangesTrend() SummaryChangesTrends`

GetRepositoriesCountChangesTrend returns the RepositoriesCountChangesTrend field if non-nil, zero value otherwise.

### GetRepositoriesCountChangesTrendOk

`func (o *SummaryStatistics) GetRepositoriesCountChangesTrendOk() (*SummaryChangesTrends, bool)`

GetRepositoriesCountChangesTrendOk returns a tuple with the RepositoriesCountChangesTrend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoriesCountChangesTrend

`func (o *SummaryStatistics) SetRepositoriesCountChangesTrend(v SummaryChangesTrends)`

SetRepositoriesCountChangesTrend sets RepositoriesCountChangesTrend field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


