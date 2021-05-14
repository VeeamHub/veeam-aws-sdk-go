# CostEstimationsPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** |  | 
**Items** | [**[]CostEstimationItem**](CostEstimationItem.md) |  | 
**IsEmpty** | **bool** |  | 
**TotalSnapshotCost** | Pointer to **float64** |  | [optional] 
**TotalReplicaCost** | Pointer to **float64** |  | [optional] 
**TotalBackupCost** | Pointer to **float64** |  | [optional] 
**TotalTrafficCost** | **float64** |  | 
**TotalTransactionCost** | **float64** |  | 
**TotalCost** | **float64** |  | 
**Links** | Pointer to [**[]Link**](Link.md) |  | [optional] 

## Methods

### NewCostEstimationsPage

`func NewCostEstimationsPage(totalCount int32, items []CostEstimationItem, isEmpty bool, totalTrafficCost float64, totalTransactionCost float64, totalCost float64, ) *CostEstimationsPage`

NewCostEstimationsPage instantiates a new CostEstimationsPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostEstimationsPageWithDefaults

`func NewCostEstimationsPageWithDefaults() *CostEstimationsPage`

NewCostEstimationsPageWithDefaults instantiates a new CostEstimationsPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *CostEstimationsPage) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *CostEstimationsPage) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *CostEstimationsPage) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetItems

`func (o *CostEstimationsPage) GetItems() []CostEstimationItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CostEstimationsPage) GetItemsOk() (*[]CostEstimationItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CostEstimationsPage) SetItems(v []CostEstimationItem)`

SetItems sets Items field to given value.


### GetIsEmpty

`func (o *CostEstimationsPage) GetIsEmpty() bool`

GetIsEmpty returns the IsEmpty field if non-nil, zero value otherwise.

### GetIsEmptyOk

`func (o *CostEstimationsPage) GetIsEmptyOk() (*bool, bool)`

GetIsEmptyOk returns a tuple with the IsEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEmpty

`func (o *CostEstimationsPage) SetIsEmpty(v bool)`

SetIsEmpty sets IsEmpty field to given value.


### GetTotalSnapshotCost

`func (o *CostEstimationsPage) GetTotalSnapshotCost() float64`

GetTotalSnapshotCost returns the TotalSnapshotCost field if non-nil, zero value otherwise.

### GetTotalSnapshotCostOk

`func (o *CostEstimationsPage) GetTotalSnapshotCostOk() (*float64, bool)`

GetTotalSnapshotCostOk returns a tuple with the TotalSnapshotCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSnapshotCost

`func (o *CostEstimationsPage) SetTotalSnapshotCost(v float64)`

SetTotalSnapshotCost sets TotalSnapshotCost field to given value.

### HasTotalSnapshotCost

`func (o *CostEstimationsPage) HasTotalSnapshotCost() bool`

HasTotalSnapshotCost returns a boolean if a field has been set.

### GetTotalReplicaCost

`func (o *CostEstimationsPage) GetTotalReplicaCost() float64`

GetTotalReplicaCost returns the TotalReplicaCost field if non-nil, zero value otherwise.

### GetTotalReplicaCostOk

`func (o *CostEstimationsPage) GetTotalReplicaCostOk() (*float64, bool)`

GetTotalReplicaCostOk returns a tuple with the TotalReplicaCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReplicaCost

`func (o *CostEstimationsPage) SetTotalReplicaCost(v float64)`

SetTotalReplicaCost sets TotalReplicaCost field to given value.

### HasTotalReplicaCost

`func (o *CostEstimationsPage) HasTotalReplicaCost() bool`

HasTotalReplicaCost returns a boolean if a field has been set.

### GetTotalBackupCost

`func (o *CostEstimationsPage) GetTotalBackupCost() float64`

GetTotalBackupCost returns the TotalBackupCost field if non-nil, zero value otherwise.

### GetTotalBackupCostOk

`func (o *CostEstimationsPage) GetTotalBackupCostOk() (*float64, bool)`

GetTotalBackupCostOk returns a tuple with the TotalBackupCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBackupCost

`func (o *CostEstimationsPage) SetTotalBackupCost(v float64)`

SetTotalBackupCost sets TotalBackupCost field to given value.

### HasTotalBackupCost

`func (o *CostEstimationsPage) HasTotalBackupCost() bool`

HasTotalBackupCost returns a boolean if a field has been set.

### GetTotalTrafficCost

`func (o *CostEstimationsPage) GetTotalTrafficCost() float64`

GetTotalTrafficCost returns the TotalTrafficCost field if non-nil, zero value otherwise.

### GetTotalTrafficCostOk

`func (o *CostEstimationsPage) GetTotalTrafficCostOk() (*float64, bool)`

GetTotalTrafficCostOk returns a tuple with the TotalTrafficCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTrafficCost

`func (o *CostEstimationsPage) SetTotalTrafficCost(v float64)`

SetTotalTrafficCost sets TotalTrafficCost field to given value.


### GetTotalTransactionCost

`func (o *CostEstimationsPage) GetTotalTransactionCost() float64`

GetTotalTransactionCost returns the TotalTransactionCost field if non-nil, zero value otherwise.

### GetTotalTransactionCostOk

`func (o *CostEstimationsPage) GetTotalTransactionCostOk() (*float64, bool)`

GetTotalTransactionCostOk returns a tuple with the TotalTransactionCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTransactionCost

`func (o *CostEstimationsPage) SetTotalTransactionCost(v float64)`

SetTotalTransactionCost sets TotalTransactionCost field to given value.


### GetTotalCost

`func (o *CostEstimationsPage) GetTotalCost() float64`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *CostEstimationsPage) GetTotalCostOk() (*float64, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *CostEstimationsPage) SetTotalCost(v float64)`

SetTotalCost sets TotalCost field to given value.


### GetLinks

`func (o *CostEstimationsPage) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CostEstimationsPage) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CostEstimationsPage) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CostEstimationsPage) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


