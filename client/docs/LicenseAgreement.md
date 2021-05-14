# LicenseAgreement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checksum** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**Type** | [**LicenseAgreementTypes**](LicenseAgreementTypes.md) |  | 

## Methods

### NewLicenseAgreement

`func NewLicenseAgreement(type_ LicenseAgreementTypes, ) *LicenseAgreement`

NewLicenseAgreement instantiates a new LicenseAgreement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseAgreementWithDefaults

`func NewLicenseAgreementWithDefaults() *LicenseAgreement`

NewLicenseAgreementWithDefaults instantiates a new LicenseAgreement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChecksum

`func (o *LicenseAgreement) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *LicenseAgreement) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *LicenseAgreement) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *LicenseAgreement) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### GetContent

`func (o *LicenseAgreement) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *LicenseAgreement) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *LicenseAgreement) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *LicenseAgreement) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetType

`func (o *LicenseAgreement) GetType() LicenseAgreementTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LicenseAgreement) GetTypeOk() (*LicenseAgreementTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LicenseAgreement) SetType(v LicenseAgreementTypes)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


