# CloudEncryptionKeysPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** |  | 
**Results** | [**[]CloudEncryptionKey**](CloudEncryptionKey.md) |  | 
**Links** | Pointer to [**[]Link**](Link.md) |  | [optional] 

## Methods

### NewCloudEncryptionKeysPage

`func NewCloudEncryptionKeysPage(totalCount int32, results []CloudEncryptionKey, ) *CloudEncryptionKeysPage`

NewCloudEncryptionKeysPage instantiates a new CloudEncryptionKeysPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudEncryptionKeysPageWithDefaults

`func NewCloudEncryptionKeysPageWithDefaults() *CloudEncryptionKeysPage`

NewCloudEncryptionKeysPageWithDefaults instantiates a new CloudEncryptionKeysPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *CloudEncryptionKeysPage) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *CloudEncryptionKeysPage) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *CloudEncryptionKeysPage) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetResults

`func (o *CloudEncryptionKeysPage) GetResults() []CloudEncryptionKey`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CloudEncryptionKeysPage) GetResultsOk() (*[]CloudEncryptionKey, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CloudEncryptionKeysPage) SetResults(v []CloudEncryptionKey)`

SetResults sets Results field to given value.


### GetLinks

`func (o *CloudEncryptionKeysPage) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CloudEncryptionKeysPage) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CloudEncryptionKeysPage) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CloudEncryptionKeysPage) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


