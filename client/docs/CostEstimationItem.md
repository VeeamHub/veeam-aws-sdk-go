# CostEstimationItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VirtualMachineId** | **string** |  | 
**VirtualMachineName** | **string** |  | 
**SnapshotCost** | Pointer to **float64** |  | [optional] 
**ReplicaCost** | Pointer to **float64** |  | [optional] 
**BackupCost** | Pointer to **float64** |  | [optional] 
**TrafficCost** | **float64** |  | 
**TransactionCost** | **float64** |  | 
**TotalCost** | **float64** |  | 

## Methods

### NewCostEstimationItem

`func NewCostEstimationItem(virtualMachineId string, virtualMachineName string, trafficCost float64, transactionCost float64, totalCost float64, ) *CostEstimationItem`

NewCostEstimationItem instantiates a new CostEstimationItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostEstimationItemWithDefaults

`func NewCostEstimationItemWithDefaults() *CostEstimationItem`

NewCostEstimationItemWithDefaults instantiates a new CostEstimationItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVirtualMachineId

`func (o *CostEstimationItem) GetVirtualMachineId() string`

GetVirtualMachineId returns the VirtualMachineId field if non-nil, zero value otherwise.

### GetVirtualMachineIdOk

`func (o *CostEstimationItem) GetVirtualMachineIdOk() (*string, bool)`

GetVirtualMachineIdOk returns a tuple with the VirtualMachineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachineId

`func (o *CostEstimationItem) SetVirtualMachineId(v string)`

SetVirtualMachineId sets VirtualMachineId field to given value.


### GetVirtualMachineName

`func (o *CostEstimationItem) GetVirtualMachineName() string`

GetVirtualMachineName returns the VirtualMachineName field if non-nil, zero value otherwise.

### GetVirtualMachineNameOk

`func (o *CostEstimationItem) GetVirtualMachineNameOk() (*string, bool)`

GetVirtualMachineNameOk returns a tuple with the VirtualMachineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachineName

`func (o *CostEstimationItem) SetVirtualMachineName(v string)`

SetVirtualMachineName sets VirtualMachineName field to given value.


### GetSnapshotCost

`func (o *CostEstimationItem) GetSnapshotCost() float64`

GetSnapshotCost returns the SnapshotCost field if non-nil, zero value otherwise.

### GetSnapshotCostOk

`func (o *CostEstimationItem) GetSnapshotCostOk() (*float64, bool)`

GetSnapshotCostOk returns a tuple with the SnapshotCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotCost

`func (o *CostEstimationItem) SetSnapshotCost(v float64)`

SetSnapshotCost sets SnapshotCost field to given value.

### HasSnapshotCost

`func (o *CostEstimationItem) HasSnapshotCost() bool`

HasSnapshotCost returns a boolean if a field has been set.

### GetReplicaCost

`func (o *CostEstimationItem) GetReplicaCost() float64`

GetReplicaCost returns the ReplicaCost field if non-nil, zero value otherwise.

### GetReplicaCostOk

`func (o *CostEstimationItem) GetReplicaCostOk() (*float64, bool)`

GetReplicaCostOk returns a tuple with the ReplicaCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaCost

`func (o *CostEstimationItem) SetReplicaCost(v float64)`

SetReplicaCost sets ReplicaCost field to given value.

### HasReplicaCost

`func (o *CostEstimationItem) HasReplicaCost() bool`

HasReplicaCost returns a boolean if a field has been set.

### GetBackupCost

`func (o *CostEstimationItem) GetBackupCost() float64`

GetBackupCost returns the BackupCost field if non-nil, zero value otherwise.

### GetBackupCostOk

`func (o *CostEstimationItem) GetBackupCostOk() (*float64, bool)`

GetBackupCostOk returns a tuple with the BackupCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupCost

`func (o *CostEstimationItem) SetBackupCost(v float64)`

SetBackupCost sets BackupCost field to given value.

### HasBackupCost

`func (o *CostEstimationItem) HasBackupCost() bool`

HasBackupCost returns a boolean if a field has been set.

### GetTrafficCost

`func (o *CostEstimationItem) GetTrafficCost() float64`

GetTrafficCost returns the TrafficCost field if non-nil, zero value otherwise.

### GetTrafficCostOk

`func (o *CostEstimationItem) GetTrafficCostOk() (*float64, bool)`

GetTrafficCostOk returns a tuple with the TrafficCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficCost

`func (o *CostEstimationItem) SetTrafficCost(v float64)`

SetTrafficCost sets TrafficCost field to given value.


### GetTransactionCost

`func (o *CostEstimationItem) GetTransactionCost() float64`

GetTransactionCost returns the TransactionCost field if non-nil, zero value otherwise.

### GetTransactionCostOk

`func (o *CostEstimationItem) GetTransactionCostOk() (*float64, bool)`

GetTransactionCostOk returns a tuple with the TransactionCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionCost

`func (o *CostEstimationItem) SetTransactionCost(v float64)`

SetTransactionCost sets TransactionCost field to given value.


### GetTotalCost

`func (o *CostEstimationItem) GetTotalCost() float64`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *CostEstimationItem) GetTotalCostOk() (*float64, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *CostEstimationItem) SetTotalCost(v float64)`

SetTotalCost sets TotalCost field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


