# SendPollRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | **string** | The WhatsApp phone number to send the poll to, including the &#39;@s.whatsapp.net&#39; suffix. | 
**Question** | **string** | The question for the poll. | 
**Options** | **[]string** | The options for the poll. | 
**MaxAnswer** | **int32** | The maximum number of answers allowed for the poll. | 
**Duration** | Pointer to **int32** | Disappearing message duration in seconds (optional) | [optional] 

## Methods

### NewSendPollRequest

`func NewSendPollRequest(phone string, question string, options []string, maxAnswer int32, ) *SendPollRequest`

NewSendPollRequest instantiates a new SendPollRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendPollRequestWithDefaults

`func NewSendPollRequestWithDefaults() *SendPollRequest`

NewSendPollRequestWithDefaults instantiates a new SendPollRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *SendPollRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SendPollRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SendPollRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetQuestion

`func (o *SendPollRequest) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *SendPollRequest) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *SendPollRequest) SetQuestion(v string)`

SetQuestion sets Question field to given value.


### GetOptions

`func (o *SendPollRequest) GetOptions() []string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SendPollRequest) GetOptionsOk() (*[]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SendPollRequest) SetOptions(v []string)`

SetOptions sets Options field to given value.


### GetMaxAnswer

`func (o *SendPollRequest) GetMaxAnswer() int32`

GetMaxAnswer returns the MaxAnswer field if non-nil, zero value otherwise.

### GetMaxAnswerOk

`func (o *SendPollRequest) GetMaxAnswerOk() (*int32, bool)`

GetMaxAnswerOk returns a tuple with the MaxAnswer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAnswer

`func (o *SendPollRequest) SetMaxAnswer(v int32)`

SetMaxAnswer sets MaxAnswer field to given value.


### GetDuration

`func (o *SendPollRequest) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *SendPollRequest) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *SendPollRequest) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *SendPollRequest) HasDuration() bool`

HasDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


