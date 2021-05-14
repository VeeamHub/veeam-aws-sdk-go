# StartFlrResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionId** | **string** |  | 
**Url** | **string** |  | 
**Links** | Pointer to [**[]Link**](Link.md) |  | [optional] 

## Methods

### NewStartFlrResponse

`func NewStartFlrResponse(sessionId string, url string, ) *StartFlrResponse`

NewStartFlrResponse instantiates a new StartFlrResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartFlrResponseWithDefaults

`func NewStartFlrResponseWithDefaults() *StartFlrResponse`

NewStartFlrResponseWithDefaults instantiates a new StartFlrResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionId

`func (o *StartFlrResponse) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *StartFlrResponse) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *StartFlrResponse) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetUrl

`func (o *StartFlrResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *StartFlrResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *StartFlrResponse) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetLinks

`func (o *StartFlrResponse) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StartFlrResponse) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StartFlrResponse) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StartFlrResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


