# AmazonAccountAccessKeysCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | **string** |  | 
**SecretKey** | **string** |  | 
**DefaultRegionType** | [**RegionTypes**](RegionTypes.md) |  | 

## Methods

### NewAmazonAccountAccessKeysCreateSpec

`func NewAmazonAccountAccessKeysCreateSpec(accessKey string, secretKey string, defaultRegionType RegionTypes, ) *AmazonAccountAccessKeysCreateSpec`

NewAmazonAccountAccessKeysCreateSpec instantiates a new AmazonAccountAccessKeysCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonAccountAccessKeysCreateSpecWithDefaults

`func NewAmazonAccountAccessKeysCreateSpecWithDefaults() *AmazonAccountAccessKeysCreateSpec`

NewAmazonAccountAccessKeysCreateSpecWithDefaults instantiates a new AmazonAccountAccessKeysCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *AmazonAccountAccessKeysCreateSpec) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *AmazonAccountAccessKeysCreateSpec) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *AmazonAccountAccessKeysCreateSpec) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### GetSecretKey

`func (o *AmazonAccountAccessKeysCreateSpec) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *AmazonAccountAccessKeysCreateSpec) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *AmazonAccountAccessKeysCreateSpec) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.


### GetDefaultRegionType

`func (o *AmazonAccountAccessKeysCreateSpec) GetDefaultRegionType() RegionTypes`

GetDefaultRegionType returns the DefaultRegionType field if non-nil, zero value otherwise.

### GetDefaultRegionTypeOk

`func (o *AmazonAccountAccessKeysCreateSpec) GetDefaultRegionTypeOk() (*RegionTypes, bool)`

GetDefaultRegionTypeOk returns a tuple with the DefaultRegionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRegionType

`func (o *AmazonAccountAccessKeysCreateSpec) SetDefaultRegionType(v RegionTypes)`

SetDefaultRegionType sets DefaultRegionType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


