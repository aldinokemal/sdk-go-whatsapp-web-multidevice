# LabelChatRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LabelId** | **string** | Unique identifier for the label | 
**LabelName** | **string** | Display name for the label | 
**Labeled** | **bool** | Whether to apply (true) or remove (false) the label | 

## Methods

### NewLabelChatRequest

`func NewLabelChatRequest(labelId string, labelName string, labeled bool, ) *LabelChatRequest`

NewLabelChatRequest instantiates a new LabelChatRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelChatRequestWithDefaults

`func NewLabelChatRequestWithDefaults() *LabelChatRequest`

NewLabelChatRequestWithDefaults instantiates a new LabelChatRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabelId

`func (o *LabelChatRequest) GetLabelId() string`

GetLabelId returns the LabelId field if non-nil, zero value otherwise.

### GetLabelIdOk

`func (o *LabelChatRequest) GetLabelIdOk() (*string, bool)`

GetLabelIdOk returns a tuple with the LabelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelId

`func (o *LabelChatRequest) SetLabelId(v string)`

SetLabelId sets LabelId field to given value.


### GetLabelName

`func (o *LabelChatRequest) GetLabelName() string`

GetLabelName returns the LabelName field if non-nil, zero value otherwise.

### GetLabelNameOk

`func (o *LabelChatRequest) GetLabelNameOk() (*string, bool)`

GetLabelNameOk returns a tuple with the LabelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelName

`func (o *LabelChatRequest) SetLabelName(v string)`

SetLabelName sets LabelName field to given value.


### GetLabeled

`func (o *LabelChatRequest) GetLabeled() bool`

GetLabeled returns the Labeled field if non-nil, zero value otherwise.

### GetLabeledOk

`func (o *LabelChatRequest) GetLabeledOk() (*bool, bool)`

GetLabeledOk returns a tuple with the Labeled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabeled

`func (o *LabelChatRequest) SetLabeled(v bool)`

SetLabeled sets Labeled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


