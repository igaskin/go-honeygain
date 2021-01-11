# Auth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 

## Methods

### NewAuth

`func NewAuth() *Auth`

NewAuth instantiates a new Auth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthWithDefaults

`func NewAuthWithDefaults() *Auth`

NewAuthWithDefaults instantiates a new Auth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *Auth) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Auth) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Auth) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Auth) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPassword

`func (o *Auth) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Auth) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Auth) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *Auth) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


