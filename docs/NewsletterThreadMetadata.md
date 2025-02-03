# NewsletterThreadMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationTime** | Pointer to **string** |  | [optional] 
**Invite** | Pointer to **string** |  | [optional] 
**Name** | Pointer to [**NewsletterThreadMetadataName**](NewsletterThreadMetadataName.md) |  | [optional] 
**Description** | Pointer to [**NewsletterThreadMetadataDescription**](NewsletterThreadMetadataDescription.md) |  | [optional] 
**SubscribersCount** | Pointer to **string** |  | [optional] 
**Verification** | Pointer to **string** |  | [optional] 
**Picture** | Pointer to [**NewsletterThreadMetadataPicture**](NewsletterThreadMetadataPicture.md) |  | [optional] 
**Preview** | Pointer to [**NewsletterThreadMetadataPreview**](NewsletterThreadMetadataPreview.md) |  | [optional] 
**Settings** | Pointer to [**NewsletterThreadMetadataSettings**](NewsletterThreadMetadataSettings.md) |  | [optional] 

## Methods

### NewNewsletterThreadMetadata

`func NewNewsletterThreadMetadata() *NewsletterThreadMetadata`

NewNewsletterThreadMetadata instantiates a new NewsletterThreadMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewsletterThreadMetadataWithDefaults

`func NewNewsletterThreadMetadataWithDefaults() *NewsletterThreadMetadata`

NewNewsletterThreadMetadataWithDefaults instantiates a new NewsletterThreadMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationTime

`func (o *NewsletterThreadMetadata) GetCreationTime() string`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *NewsletterThreadMetadata) GetCreationTimeOk() (*string, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *NewsletterThreadMetadata) SetCreationTime(v string)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *NewsletterThreadMetadata) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetInvite

`func (o *NewsletterThreadMetadata) GetInvite() string`

GetInvite returns the Invite field if non-nil, zero value otherwise.

### GetInviteOk

`func (o *NewsletterThreadMetadata) GetInviteOk() (*string, bool)`

GetInviteOk returns a tuple with the Invite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvite

`func (o *NewsletterThreadMetadata) SetInvite(v string)`

SetInvite sets Invite field to given value.

### HasInvite

`func (o *NewsletterThreadMetadata) HasInvite() bool`

HasInvite returns a boolean if a field has been set.

### GetName

`func (o *NewsletterThreadMetadata) GetName() NewsletterThreadMetadataName`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewsletterThreadMetadata) GetNameOk() (*NewsletterThreadMetadataName, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewsletterThreadMetadata) SetName(v NewsletterThreadMetadataName)`

SetName sets Name field to given value.

### HasName

`func (o *NewsletterThreadMetadata) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *NewsletterThreadMetadata) GetDescription() NewsletterThreadMetadataDescription`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewsletterThreadMetadata) GetDescriptionOk() (*NewsletterThreadMetadataDescription, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NewsletterThreadMetadata) SetDescription(v NewsletterThreadMetadataDescription)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NewsletterThreadMetadata) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSubscribersCount

`func (o *NewsletterThreadMetadata) GetSubscribersCount() string`

GetSubscribersCount returns the SubscribersCount field if non-nil, zero value otherwise.

### GetSubscribersCountOk

`func (o *NewsletterThreadMetadata) GetSubscribersCountOk() (*string, bool)`

GetSubscribersCountOk returns a tuple with the SubscribersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribersCount

`func (o *NewsletterThreadMetadata) SetSubscribersCount(v string)`

SetSubscribersCount sets SubscribersCount field to given value.

### HasSubscribersCount

`func (o *NewsletterThreadMetadata) HasSubscribersCount() bool`

HasSubscribersCount returns a boolean if a field has been set.

### GetVerification

`func (o *NewsletterThreadMetadata) GetVerification() string`

GetVerification returns the Verification field if non-nil, zero value otherwise.

### GetVerificationOk

`func (o *NewsletterThreadMetadata) GetVerificationOk() (*string, bool)`

GetVerificationOk returns a tuple with the Verification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerification

`func (o *NewsletterThreadMetadata) SetVerification(v string)`

SetVerification sets Verification field to given value.

### HasVerification

`func (o *NewsletterThreadMetadata) HasVerification() bool`

HasVerification returns a boolean if a field has been set.

### GetPicture

`func (o *NewsletterThreadMetadata) GetPicture() NewsletterThreadMetadataPicture`

GetPicture returns the Picture field if non-nil, zero value otherwise.

### GetPictureOk

`func (o *NewsletterThreadMetadata) GetPictureOk() (*NewsletterThreadMetadataPicture, bool)`

GetPictureOk returns a tuple with the Picture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicture

`func (o *NewsletterThreadMetadata) SetPicture(v NewsletterThreadMetadataPicture)`

SetPicture sets Picture field to given value.

### HasPicture

`func (o *NewsletterThreadMetadata) HasPicture() bool`

HasPicture returns a boolean if a field has been set.

### GetPreview

`func (o *NewsletterThreadMetadata) GetPreview() NewsletterThreadMetadataPreview`

GetPreview returns the Preview field if non-nil, zero value otherwise.

### GetPreviewOk

`func (o *NewsletterThreadMetadata) GetPreviewOk() (*NewsletterThreadMetadataPreview, bool)`

GetPreviewOk returns a tuple with the Preview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreview

`func (o *NewsletterThreadMetadata) SetPreview(v NewsletterThreadMetadataPreview)`

SetPreview sets Preview field to given value.

### HasPreview

`func (o *NewsletterThreadMetadata) HasPreview() bool`

HasPreview returns a boolean if a field has been set.

### GetSettings

`func (o *NewsletterThreadMetadata) GetSettings() NewsletterThreadMetadataSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *NewsletterThreadMetadata) GetSettingsOk() (*NewsletterThreadMetadataSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *NewsletterThreadMetadata) SetSettings(v NewsletterThreadMetadataSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *NewsletterThreadMetadata) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


