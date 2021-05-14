# DefaultUserSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instance** | [**CheckInstanceIdSpec**](CheckInstanceIdSpec.md) |  | 
**UserSpec** | [**UserCreateSpec**](UserCreateSpec.md) |  | 

## Methods

### NewDefaultUserSpec

`func NewDefaultUserSpec(instance CheckInstanceIdSpec, userSpec UserCreateSpec, ) *DefaultUserSpec`

NewDefaultUserSpec instantiates a new DefaultUserSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultUserSpecWithDefaults

`func NewDefaultUserSpecWithDefaults() *DefaultUserSpec`

NewDefaultUserSpecWithDefaults instantiates a new DefaultUserSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstance

`func (o *DefaultUserSpec) GetInstance() CheckInstanceIdSpec`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *DefaultUserSpec) GetInstanceOk() (*CheckInstanceIdSpec, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *DefaultUserSpec) SetInstance(v CheckInstanceIdSpec)`

SetInstance sets Instance field to given value.


### GetUserSpec

`func (o *DefaultUserSpec) GetUserSpec() UserCreateSpec`

GetUserSpec returns the UserSpec field if non-nil, zero value otherwise.

### GetUserSpecOk

`func (o *DefaultUserSpec) GetUserSpecOk() (*UserCreateSpec, bool)`

GetUserSpecOk returns a tuple with the UserSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSpec

`func (o *DefaultUserSpec) SetUserSpec(v UserCreateSpec)`

SetUserSpec sets UserSpec field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


