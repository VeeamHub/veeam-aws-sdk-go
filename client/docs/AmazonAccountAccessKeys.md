# AmazonAccountAccessKeys

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | **string** |  | 
**DefaultRegion** | Pointer to [**RegionTypes**](RegionTypes.md) |  | [optional] 

## Methods

### NewAmazonAccountAccessKeys

`func NewAmazonAccountAccessKeys(accessKey string, ) *AmazonAccountAccessKeys`

NewAmazonAccountAccessKeys instantiates a new AmazonAccountAccessKeys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonAccountAccessKeysWithDefaults

`func NewAmazonAccountAccessKeysWithDefaults() *AmazonAccountAccessKeys`

NewAmazonAccountAccessKeysWithDefaults instantiates a new AmazonAccountAccessKeys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *AmazonAccountAccessKeys) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *AmazonAccountAccessKeys) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *AmazonAccountAccessKeys) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### GetDefaultRegion

`func (o *AmazonAccountAccessKeys) GetDefaultRegion() RegionTypes`

GetDefaultRegion returns the DefaultRegion field if non-nil, zero value otherwise.

### GetDefaultRegionOk

`func (o *AmazonAccountAccessKeys) GetDefaultRegionOk() (*RegionTypes, bool)`

GetDefaultRegionOk returns a tuple with the DefaultRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRegion

`func (o *AmazonAccountAccessKeys) SetDefaultRegion(v RegionTypes)`

SetDefaultRegion sets DefaultRegion field to given value.

### HasDefaultRegion

`func (o *AmazonAccountAccessKeys) HasDefaultRegion() bool`

HasDefaultRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


