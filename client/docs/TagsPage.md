# TagsPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]Tag**](Tag.md) |  | 
**TotalCount** | **int32** |  | 
**Links** | Pointer to [**[]Link**](Link.md) |  | [optional] 

## Methods

### NewTagsPage

`func NewTagsPage(results []Tag, totalCount int32, ) *TagsPage`

NewTagsPage instantiates a new TagsPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagsPageWithDefaults

`func NewTagsPageWithDefaults() *TagsPage`

NewTagsPageWithDefaults instantiates a new TagsPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *TagsPage) GetResults() []Tag`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *TagsPage) GetResultsOk() (*[]Tag, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *TagsPage) SetResults(v []Tag)`

SetResults sets Results field to given value.


### GetTotalCount

`func (o *TagsPage) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *TagsPage) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *TagsPage) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetLinks

`func (o *TagsPage) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TagsPage) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TagsPage) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TagsPage) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


