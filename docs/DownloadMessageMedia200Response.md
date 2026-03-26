# DownloadMessageMedia200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Results** | Pointer to [**DownloadMessageMedia200ResponseResults**](DownloadMessageMedia200ResponseResults.md) |  | [optional] 

## Methods

### NewDownloadMessageMedia200Response

`func NewDownloadMessageMedia200Response() *DownloadMessageMedia200Response`

NewDownloadMessageMedia200Response instantiates a new DownloadMessageMedia200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDownloadMessageMedia200ResponseWithDefaults

`func NewDownloadMessageMedia200ResponseWithDefaults() *DownloadMessageMedia200Response`

NewDownloadMessageMedia200ResponseWithDefaults instantiates a new DownloadMessageMedia200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *DownloadMessageMedia200Response) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DownloadMessageMedia200Response) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DownloadMessageMedia200Response) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DownloadMessageMedia200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCode

`func (o *DownloadMessageMedia200Response) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DownloadMessageMedia200Response) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DownloadMessageMedia200Response) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *DownloadMessageMedia200Response) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *DownloadMessageMedia200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DownloadMessageMedia200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DownloadMessageMedia200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DownloadMessageMedia200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResults

`func (o *DownloadMessageMedia200Response) GetResults() DownloadMessageMedia200ResponseResults`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *DownloadMessageMedia200Response) GetResultsOk() (*DownloadMessageMedia200ResponseResults, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *DownloadMessageMedia200Response) SetResults(v DownloadMessageMedia200ResponseResults)`

SetResults sets Results field to given value.

### HasResults

`func (o *DownloadMessageMedia200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


