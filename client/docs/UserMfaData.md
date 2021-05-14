# UserMfaData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserName** | Pointer to **string** |  | [optional] 
**QrString** | Pointer to **string** |  | [optional] 
**SecretKey** | Pointer to **string** |  | [optional] 
**ScratchCodes** | Pointer to **[]string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 

## Methods

### NewUserMfaData

`func NewUserMfaData() *UserMfaData`

NewUserMfaData instantiates a new UserMfaData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserMfaDataWithDefaults

`func NewUserMfaDataWithDefaults() *UserMfaData`

NewUserMfaDataWithDefaults instantiates a new UserMfaData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserName

`func (o *UserMfaData) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *UserMfaData) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *UserMfaData) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *UserMfaData) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetQrString

`func (o *UserMfaData) GetQrString() string`

GetQrString returns the QrString field if non-nil, zero value otherwise.

### GetQrStringOk

`func (o *UserMfaData) GetQrStringOk() (*string, bool)`

GetQrStringOk returns a tuple with the QrString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrString

`func (o *UserMfaData) SetQrString(v string)`

SetQrString sets QrString field to given value.

### HasQrString

`func (o *UserMfaData) HasQrString() bool`

HasQrString returns a boolean if a field has been set.

### GetSecretKey

`func (o *UserMfaData) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *UserMfaData) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *UserMfaData) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *UserMfaData) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetScratchCodes

`func (o *UserMfaData) GetScratchCodes() []string`

GetScratchCodes returns the ScratchCodes field if non-nil, zero value otherwise.

### GetScratchCodesOk

`func (o *UserMfaData) GetScratchCodesOk() (*[]string, bool)`

GetScratchCodesOk returns a tuple with the ScratchCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScratchCodes

`func (o *UserMfaData) SetScratchCodes(v []string)`

SetScratchCodes sets ScratchCodes field to given value.

### HasScratchCodes

`func (o *UserMfaData) HasScratchCodes() bool`

HasScratchCodes returns a boolean if a field has been set.

### GetToken

`func (o *UserMfaData) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UserMfaData) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UserMfaData) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UserMfaData) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


