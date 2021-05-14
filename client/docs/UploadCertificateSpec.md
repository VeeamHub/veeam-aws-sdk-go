# UploadCertificateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateFile** | ***os.File** | The .PFX certificate to upload. | 
**CertificatePassword** | **string** | .PFX certificate password | 

## Methods

### NewUploadCertificateSpec

`func NewUploadCertificateSpec(certificateFile *os.File, certificatePassword string, ) *UploadCertificateSpec`

NewUploadCertificateSpec instantiates a new UploadCertificateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadCertificateSpecWithDefaults

`func NewUploadCertificateSpecWithDefaults() *UploadCertificateSpec`

NewUploadCertificateSpecWithDefaults instantiates a new UploadCertificateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateFile

`func (o *UploadCertificateSpec) GetCertificateFile() *os.File`

GetCertificateFile returns the CertificateFile field if non-nil, zero value otherwise.

### GetCertificateFileOk

`func (o *UploadCertificateSpec) GetCertificateFileOk() (**os.File, bool)`

GetCertificateFileOk returns a tuple with the CertificateFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateFile

`func (o *UploadCertificateSpec) SetCertificateFile(v *os.File)`

SetCertificateFile sets CertificateFile field to given value.


### GetCertificatePassword

`func (o *UploadCertificateSpec) GetCertificatePassword() string`

GetCertificatePassword returns the CertificatePassword field if non-nil, zero value otherwise.

### GetCertificatePasswordOk

`func (o *UploadCertificateSpec) GetCertificatePasswordOk() (*string, bool)`

GetCertificatePasswordOk returns a tuple with the CertificatePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatePassword

`func (o *UploadCertificateSpec) SetCertificatePassword(v string)`

SetCertificatePassword sets CertificatePassword field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


