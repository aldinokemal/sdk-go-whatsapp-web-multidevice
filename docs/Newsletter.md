# Newsletter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**State** | Pointer to [**NewsletterState**](NewsletterState.md) |  | [optional] 
**ThreadMetadata** | Pointer to [**NewsletterThreadMetadata**](NewsletterThreadMetadata.md) |  | [optional] 
**ViewerMetadata** | Pointer to [**NewsletterViewerMetadata**](NewsletterViewerMetadata.md) |  | [optional] 

## Methods

### NewNewsletter

`func NewNewsletter() *Newsletter`

NewNewsletter instantiates a new Newsletter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewsletterWithDefaults

`func NewNewsletterWithDefaults() *Newsletter`

NewNewsletterWithDefaults instantiates a new Newsletter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Newsletter) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Newsletter) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Newsletter) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Newsletter) HasId() bool`

HasId returns a boolean if a field has been set.

### GetState

`func (o *Newsletter) GetState() NewsletterState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Newsletter) GetStateOk() (*NewsletterState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Newsletter) SetState(v NewsletterState)`

SetState sets State field to given value.

### HasState

`func (o *Newsletter) HasState() bool`

HasState returns a boolean if a field has been set.

### GetThreadMetadata

`func (o *Newsletter) GetThreadMetadata() NewsletterThreadMetadata`

GetThreadMetadata returns the ThreadMetadata field if non-nil, zero value otherwise.

### GetThreadMetadataOk

`func (o *Newsletter) GetThreadMetadataOk() (*NewsletterThreadMetadata, bool)`

GetThreadMetadataOk returns a tuple with the ThreadMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadMetadata

`func (o *Newsletter) SetThreadMetadata(v NewsletterThreadMetadata)`

SetThreadMetadata sets ThreadMetadata field to given value.

### HasThreadMetadata

`func (o *Newsletter) HasThreadMetadata() bool`

HasThreadMetadata returns a boolean if a field has been set.

### GetViewerMetadata

`func (o *Newsletter) GetViewerMetadata() NewsletterViewerMetadata`

GetViewerMetadata returns the ViewerMetadata field if non-nil, zero value otherwise.

### GetViewerMetadataOk

`func (o *Newsletter) GetViewerMetadataOk() (*NewsletterViewerMetadata, bool)`

GetViewerMetadataOk returns a tuple with the ViewerMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerMetadata

`func (o *Newsletter) SetViewerMetadata(v NewsletterViewerMetadata)`

SetViewerMetadata sets ViewerMetadata field to given value.

### HasViewerMetadata

`func (o *Newsletter) HasViewerMetadata() bool`

HasViewerMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


