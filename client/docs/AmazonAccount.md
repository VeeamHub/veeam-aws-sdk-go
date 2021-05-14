# AmazonAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**Awsid** | **string** |  | 
**RegionType** | Pointer to [**RegionTypes**](RegionTypes.md) |  | [optional] 
**AccessKeys** | Pointer to [**AmazonAccountAccessKeys**](AmazonAccountAccessKeys.md) |  | [optional] 
**IAMRole** | Pointer to [**AmazonAccountIAMRole**](AmazonAccountIAMRole.md) |  | [optional] 
**IAMRoleFromAnotherAccount** | Pointer to [**AmazonAccountIAMRoleFromAnotherAccount**](AmazonAccountIAMRoleFromAnotherAccount.md) |  | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) |  | [optional] 

## Methods

### NewAmazonAccount

`func NewAmazonAccount(id string, name string, description string, awsid string, ) *AmazonAccount`

NewAmazonAccount instantiates a new AmazonAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonAccountWithDefaults

`func NewAmazonAccountWithDefaults() *AmazonAccount`

NewAmazonAccountWithDefaults instantiates a new AmazonAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AmazonAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AmazonAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AmazonAccount) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AmazonAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AmazonAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AmazonAccount) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AmazonAccount) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AmazonAccount) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AmazonAccount) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetAwsid

`func (o *AmazonAccount) GetAwsid() string`

GetAwsid returns the Awsid field if non-nil, zero value otherwise.

### GetAwsidOk

`func (o *AmazonAccount) GetAwsidOk() (*string, bool)`

GetAwsidOk returns a tuple with the Awsid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsid

`func (o *AmazonAccount) SetAwsid(v string)`

SetAwsid sets Awsid field to given value.


### GetRegionType

`func (o *AmazonAccount) GetRegionType() RegionTypes`

GetRegionType returns the RegionType field if non-nil, zero value otherwise.

### GetRegionTypeOk

`func (o *AmazonAccount) GetRegionTypeOk() (*RegionTypes, bool)`

GetRegionTypeOk returns a tuple with the RegionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionType

`func (o *AmazonAccount) SetRegionType(v RegionTypes)`

SetRegionType sets RegionType field to given value.

### HasRegionType

`func (o *AmazonAccount) HasRegionType() bool`

HasRegionType returns a boolean if a field has been set.

### GetAccessKeys

`func (o *AmazonAccount) GetAccessKeys() AmazonAccountAccessKeys`

GetAccessKeys returns the AccessKeys field if non-nil, zero value otherwise.

### GetAccessKeysOk

`func (o *AmazonAccount) GetAccessKeysOk() (*AmazonAccountAccessKeys, bool)`

GetAccessKeysOk returns a tuple with the AccessKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeys

`func (o *AmazonAccount) SetAccessKeys(v AmazonAccountAccessKeys)`

SetAccessKeys sets AccessKeys field to given value.

### HasAccessKeys

`func (o *AmazonAccount) HasAccessKeys() bool`

HasAccessKeys returns a boolean if a field has been set.

### GetIAMRole

`func (o *AmazonAccount) GetIAMRole() AmazonAccountIAMRole`

GetIAMRole returns the IAMRole field if non-nil, zero value otherwise.

### GetIAMRoleOk

`func (o *AmazonAccount) GetIAMRoleOk() (*AmazonAccountIAMRole, bool)`

GetIAMRoleOk returns a tuple with the IAMRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIAMRole

`func (o *AmazonAccount) SetIAMRole(v AmazonAccountIAMRole)`

SetIAMRole sets IAMRole field to given value.

### HasIAMRole

`func (o *AmazonAccount) HasIAMRole() bool`

HasIAMRole returns a boolean if a field has been set.

### GetIAMRoleFromAnotherAccount

`func (o *AmazonAccount) GetIAMRoleFromAnotherAccount() AmazonAccountIAMRoleFromAnotherAccount`

GetIAMRoleFromAnotherAccount returns the IAMRoleFromAnotherAccount field if non-nil, zero value otherwise.

### GetIAMRoleFromAnotherAccountOk

`func (o *AmazonAccount) GetIAMRoleFromAnotherAccountOk() (*AmazonAccountIAMRoleFromAnotherAccount, bool)`

GetIAMRoleFromAnotherAccountOk returns a tuple with the IAMRoleFromAnotherAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIAMRoleFromAnotherAccount

`func (o *AmazonAccount) SetIAMRoleFromAnotherAccount(v AmazonAccountIAMRoleFromAnotherAccount)`

SetIAMRoleFromAnotherAccount sets IAMRoleFromAnotherAccount field to given value.

### HasIAMRoleFromAnotherAccount

`func (o *AmazonAccount) HasIAMRoleFromAnotherAccount() bool`

HasIAMRoleFromAnotherAccount returns a boolean if a field has been set.

### GetLinks

`func (o *AmazonAccount) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AmazonAccount) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AmazonAccount) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AmazonAccount) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


