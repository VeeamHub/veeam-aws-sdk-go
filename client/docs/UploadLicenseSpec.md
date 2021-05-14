# UploadLicenseSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LicenseFile** | ***os.File** | The .lic license to upload. | 

## Methods

### NewUploadLicenseSpec

`func NewUploadLicenseSpec(licenseFile *os.File, ) *UploadLicenseSpec`

NewUploadLicenseSpec instantiates a new UploadLicenseSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadLicenseSpecWithDefaults

`func NewUploadLicenseSpecWithDefaults() *UploadLicenseSpec`

NewUploadLicenseSpecWithDefaults instantiates a new UploadLicenseSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicenseFile

`func (o *UploadLicenseSpec) GetLicenseFile() *os.File`

GetLicenseFile returns the LicenseFile field if non-nil, zero value otherwise.

### GetLicenseFileOk

`func (o *UploadLicenseSpec) GetLicenseFileOk() (**os.File, bool)`

GetLicenseFileOk returns a tuple with the LicenseFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseFile

`func (o *UploadLicenseSpec) SetLicenseFile(v *os.File)`

SetLicenseFile sets LicenseFile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


