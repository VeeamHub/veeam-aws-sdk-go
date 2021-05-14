# LicensedWorkload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**LastBackupTime** | Pointer to **time.Time** |  | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) |  | [optional] 

## Methods

### NewLicensedWorkload

`func NewLicensedWorkload() *LicensedWorkload`

NewLicensedWorkload instantiates a new LicensedWorkload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicensedWorkloadWithDefaults

`func NewLicensedWorkloadWithDefaults() *LicensedWorkload`

NewLicensedWorkloadWithDefaults instantiates a new LicensedWorkload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LicensedWorkload) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LicensedWorkload) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LicensedWorkload) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LicensedWorkload) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *LicensedWorkload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LicensedWorkload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LicensedWorkload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LicensedWorkload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLastBackupTime

`func (o *LicensedWorkload) GetLastBackupTime() time.Time`

GetLastBackupTime returns the LastBackupTime field if non-nil, zero value otherwise.

### GetLastBackupTimeOk

`func (o *LicensedWorkload) GetLastBackupTimeOk() (*time.Time, bool)`

GetLastBackupTimeOk returns a tuple with the LastBackupTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBackupTime

`func (o *LicensedWorkload) SetLastBackupTime(v time.Time)`

SetLastBackupTime sets LastBackupTime field to given value.

### HasLastBackupTime

`func (o *LicensedWorkload) HasLastBackupTime() bool`

HasLastBackupTime returns a boolean if a field has been set.

### GetLinks

`func (o *LicensedWorkload) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *LicensedWorkload) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *LicensedWorkload) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *LicensedWorkload) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


