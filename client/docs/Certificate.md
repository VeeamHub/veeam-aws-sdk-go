# Certificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Thumbprint** | **string** | certificate thumpbprint | 
**SerialNumber** | **string** |  | 
**KeyAlgorithm** | **string** | certificate key algorithm | 
**KeySize** | **string** | certificate key size | 
**Subject** | **string** | certificate subject | 
**IssuedTo** | **string** | certificate acquirer | 
**IssuedBy** | **string** | certificate issuer | 
**ValidFrom** | **time.Time** |  | 
**ValidBy** | **time.Time** |  | 
**AutomaticallyGenerated** | **bool** |  | 

## Methods

### NewCertificate

`func NewCertificate(thumbprint string, serialNumber string, keyAlgorithm string, keySize string, subject string, issuedTo string, issuedBy string, validFrom time.Time, validBy time.Time, automaticallyGenerated bool, ) *Certificate`

NewCertificate instantiates a new Certificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateWithDefaults

`func NewCertificateWithDefaults() *Certificate`

NewCertificateWithDefaults instantiates a new Certificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThumbprint

`func (o *Certificate) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *Certificate) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *Certificate) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.


### GetSerialNumber

`func (o *Certificate) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *Certificate) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *Certificate) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### GetKeyAlgorithm

`func (o *Certificate) GetKeyAlgorithm() string`

GetKeyAlgorithm returns the KeyAlgorithm field if non-nil, zero value otherwise.

### GetKeyAlgorithmOk

`func (o *Certificate) GetKeyAlgorithmOk() (*string, bool)`

GetKeyAlgorithmOk returns a tuple with the KeyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyAlgorithm

`func (o *Certificate) SetKeyAlgorithm(v string)`

SetKeyAlgorithm sets KeyAlgorithm field to given value.


### GetKeySize

`func (o *Certificate) GetKeySize() string`

GetKeySize returns the KeySize field if non-nil, zero value otherwise.

### GetKeySizeOk

`func (o *Certificate) GetKeySizeOk() (*string, bool)`

GetKeySizeOk returns a tuple with the KeySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySize

`func (o *Certificate) SetKeySize(v string)`

SetKeySize sets KeySize field to given value.


### GetSubject

`func (o *Certificate) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *Certificate) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *Certificate) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetIssuedTo

`func (o *Certificate) GetIssuedTo() string`

GetIssuedTo returns the IssuedTo field if non-nil, zero value otherwise.

### GetIssuedToOk

`func (o *Certificate) GetIssuedToOk() (*string, bool)`

GetIssuedToOk returns a tuple with the IssuedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedTo

`func (o *Certificate) SetIssuedTo(v string)`

SetIssuedTo sets IssuedTo field to given value.


### GetIssuedBy

`func (o *Certificate) GetIssuedBy() string`

GetIssuedBy returns the IssuedBy field if non-nil, zero value otherwise.

### GetIssuedByOk

`func (o *Certificate) GetIssuedByOk() (*string, bool)`

GetIssuedByOk returns a tuple with the IssuedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedBy

`func (o *Certificate) SetIssuedBy(v string)`

SetIssuedBy sets IssuedBy field to given value.


### GetValidFrom

`func (o *Certificate) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *Certificate) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *Certificate) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.


### GetValidBy

`func (o *Certificate) GetValidBy() time.Time`

GetValidBy returns the ValidBy field if non-nil, zero value otherwise.

### GetValidByOk

`func (o *Certificate) GetValidByOk() (*time.Time, bool)`

GetValidByOk returns a tuple with the ValidBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidBy

`func (o *Certificate) SetValidBy(v time.Time)`

SetValidBy sets ValidBy field to given value.


### GetAutomaticallyGenerated

`func (o *Certificate) GetAutomaticallyGenerated() bool`

GetAutomaticallyGenerated returns the AutomaticallyGenerated field if non-nil, zero value otherwise.

### GetAutomaticallyGeneratedOk

`func (o *Certificate) GetAutomaticallyGeneratedOk() (*bool, bool)`

GetAutomaticallyGeneratedOk returns a tuple with the AutomaticallyGenerated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticallyGenerated

`func (o *Certificate) SetAutomaticallyGenerated(v bool)`

SetAutomaticallyGenerated sets AutomaticallyGenerated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


