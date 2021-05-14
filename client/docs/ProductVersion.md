# ProductVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**About** | **string** |  | 
**Version** | **string** |  | 

## Methods

### NewProductVersion

`func NewProductVersion(about string, version string, ) *ProductVersion`

NewProductVersion instantiates a new ProductVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductVersionWithDefaults

`func NewProductVersionWithDefaults() *ProductVersion`

NewProductVersionWithDefaults instantiates a new ProductVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbout

`func (o *ProductVersion) GetAbout() string`

GetAbout returns the About field if non-nil, zero value otherwise.

### GetAboutOk

`func (o *ProductVersion) GetAboutOk() (*string, bool)`

GetAboutOk returns a tuple with the About field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbout

`func (o *ProductVersion) SetAbout(v string)`

SetAbout sets About field to given value.


### GetVersion

`func (o *ProductVersion) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ProductVersion) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ProductVersion) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


