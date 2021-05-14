# CertificateRegenerateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidBy** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCertificateRegenerateSpec

`func NewCertificateRegenerateSpec() *CertificateRegenerateSpec`

NewCertificateRegenerateSpec instantiates a new CertificateRegenerateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateRegenerateSpecWithDefaults

`func NewCertificateRegenerateSpecWithDefaults() *CertificateRegenerateSpec`

NewCertificateRegenerateSpecWithDefaults instantiates a new CertificateRegenerateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValidBy

`func (o *CertificateRegenerateSpec) GetValidBy() time.Time`

GetValidBy returns the ValidBy field if non-nil, zero value otherwise.

### GetValidByOk

`func (o *CertificateRegenerateSpec) GetValidByOk() (*time.Time, bool)`

GetValidByOk returns a tuple with the ValidBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidBy

`func (o *CertificateRegenerateSpec) SetValidBy(v time.Time)`

SetValidBy sets ValidBy field to given value.

### HasValidBy

`func (o *CertificateRegenerateSpec) HasValidBy() bool`

HasValidBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


