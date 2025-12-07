# SdkWhatsappWebMultiDevice\SendAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendAudio**](SendAPI.md#SendAudio) | **Post** /send/audio | Send Audio
[**SendChatPresence**](SendAPI.md#SendChatPresence) | **Post** /send/chat-presence | Send chat presence (typing indicator)
[**SendContact**](SendAPI.md#SendContact) | **Post** /send/contact | Send Contact
[**SendFile**](SendAPI.md#SendFile) | **Post** /send/file | Send File
[**SendImage**](SendAPI.md#SendImage) | **Post** /send/image | Send Image
[**SendLink**](SendAPI.md#SendLink) | **Post** /send/link | Send Link
[**SendLocation**](SendAPI.md#SendLocation) | **Post** /send/location | Send Location
[**SendMessage**](SendAPI.md#SendMessage) | **Post** /send/message | Send Message
[**SendPoll**](SendAPI.md#SendPoll) | **Post** /send/poll | Send Poll / Vote
[**SendPresence**](SendAPI.md#SendPresence) | **Post** /send/presence | Send presence status
[**SendSticker**](SendAPI.md#SendSticker) | **Post** /send/sticker | Send Sticker
[**SendVideo**](SendAPI.md#SendVideo) | **Post** /send/video | Send Video



## SendAudio

> SendResponse SendAudio(ctx).Phone(phone).Audio(audio).AudioUrl(audioUrl).IsForwarded(isForwarded).Duration(duration).Execute()

Send Audio

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aldinokemal/sdk-go-whatsapp-web-multidevice"
)

func main() {
	phone := "phone_example" // string | Phone number with country code (optional)
	audio := os.NewFile(1234, "some_file") // *os.File | Audio to send (optional)
	audioUrl := "audioUrl_example" // string | Audio URL to send (optional)
	isForwarded := true // bool | Whether this is a forwarded message (optional)
	duration := int32(56) // int32 | Disappearing message duration in seconds (optional) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SendAPI.SendAudio(context.Background()).Phone(phone).Audio(audio).AudioUrl(audioUrl).IsForwarded(isForwarded).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.SendAudio``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendAudio`: SendResponse
	fmt.Fprintf(os.Stdout, "Response from `SendAPI.SendAudio`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendAudioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phone** | **string** | Phone number with country code | 
 **audio** | ***os.File** | Audio to send | 
 **audioUrl** | **string** | Audio URL to send | 
 **isForwarded** | **bool** | Whether this is a forwarded message | 
 **duration** | **int32** | Disappearing message duration in seconds (optional) | 

### Return type

[**SendResponse**](SendResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendChatPresence

> SendResponse SendChatPresence(ctx).SendChatPresenceRequest(sendChatPresenceRequest).Execute()

Send chat presence (typing indicator)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aldinokemal/sdk-go-whatsapp-web-multidevice"
)

func main() {
	sendChatPresenceRequest := *openapiclient.NewSendChatPresenceRequest("6289685024051@s.whatsapp.net", "start") // SendChatPresenceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SendAPI.SendChatPresence(context.Background()).SendChatPresenceRequest(sendChatPresenceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.SendChatPresence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendChatPresence`: SendResponse
	fmt.Fprintf(os.Stdout, "Response from `SendAPI.SendChatPresence`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendChatPresenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendChatPresenceRequest** | [**SendChatPresenceRequest**](SendChatPresenceRequest.md) |  | 

### Return type

[**SendResponse**](SendResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendContact

> SendResponse SendContact(ctx).SendContactRequest(sendContactRequest).Execute()

Send Contact

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aldinokemal/sdk-go-whatsapp-web-multidevice"
)

func main() {
	sendContactRequest := *openapiclient.NewSendContactRequest() // SendContactRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SendAPI.SendContact(context.Background()).SendContactRequest(sendContactRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.SendContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendContact`: SendResponse
	fmt.Fprintf(os.Stdout, "Response from `SendAPI.SendContact`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendContactRequest** | [**SendContactRequest**](SendContactRequest.md) |  | 

### Return type

[**SendResponse**](SendResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendFile

> SendResponse SendFile(ctx).Phone(phone).Caption(caption).File(file).IsForwarded(isForwarded).Duration(duration).Execute()

Send File

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aldinokemal/sdk-go-whatsapp-web-multidevice"
)

func main() {
	phone := "phone_example" // string | Phone number with country code (optional)
	caption := "caption_example" // string | Caption to send (optional)
	file := os.NewFile(1234, "some_file") // *os.File | File to send (optional)
	isForwarded := true // bool | Whether this is a forwarded message (optional)
	duration := int32(56) // int32 | Disappearing message duration in seconds (optional) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SendAPI.SendFile(context.Background()).Phone(phone).Caption(caption).File(file).IsForwarded(isForwarded).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.SendFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendFile`: SendResponse
	fmt.Fprintf(os.Stdout, "Response from `SendAPI.SendFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phone** | **string** | Phone number with country code | 
 **caption** | **string** | Caption to send | 
 **file** | ***os.File** | File to send | 
 **isForwarded** | **bool** | Whether this is a forwarded message | 
 **duration** | **int32** | Disappearing message duration in seconds (optional) | 

### Return type

[**SendResponse**](SendResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendImage

> SendResponse SendImage(ctx).Phone(phone).Caption(caption).ViewOnce(viewOnce).Image(image).ImageUrl(imageUrl).Compress(compress).Duration(duration).IsForwarded(isForwarded).Execute()

Send Image

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aldinokemal/sdk-go-whatsapp-web-multidevice"
)

func main() {
	phone := "phone_example" // string | Phone number with country code (optional)
	caption := "caption_example" // string | Caption to send (optional)
	viewOnce := true // bool | View once (optional)
	image := os.NewFile(1234, "some_file") // *os.File | Image to send (optional)
	imageUrl := "imageUrl_example" // string | Image URL to send (optional)
	compress := true // bool | Compress image (optional)
	duration := int32(56) // int32 | Disappearing message duration in seconds (optional) (optional)
	isForwarded := true // bool | Whether this is a forwarded message (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SendAPI.SendImage(context.Background()).Phone(phone).Caption(caption).ViewOnce(viewOnce).Image(image).ImageUrl(imageUrl).Compress(compress).Duration(duration).IsForwarded(isForwarded).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.SendImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendImage`: SendResponse
	fmt.Fprintf(os.Stdout, "Response from `SendAPI.SendImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phone** | **string** | Phone number with country code | 
 **caption** | **string** | Caption to send | 
 **viewOnce** | **bool** | View once | 
 **image** | ***os.File** | Image to send | 
 **imageUrl** | **string** | Image URL to send | 
 **compress** | **bool** | Compress image | 
 **duration** | **int32** | Disappearing message duration in seconds (optional) | 
 **isForwarded** | **bool** | Whether this is a forwarded message | 

### Return type

[**SendResponse**](SendResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendLink

> SendResponse SendLink(ctx).SendLinkRequest(sendLinkRequest).Execute()

Send Link

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aldinokemal/sdk-go-whatsapp-web-multidevice"
)

func main() {
	sendLinkRequest := *openapiclient.NewSendLinkRequest() // SendLinkRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SendAPI.SendLink(context.Background()).SendLinkRequest(sendLinkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.SendLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendLink`: SendResponse
	fmt.Fprintf(os.Stdout, "Response from `SendAPI.SendLink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendLinkRequest** | [**SendLinkRequest**](SendLinkRequest.md) |  | 

### Return type

[**SendResponse**](SendResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendLocation

> SendResponse SendLocation(ctx).SendLocationRequest(sendLocationRequest).Execute()

Send Location

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aldinokemal/sdk-go-whatsapp-web-multidevice"
)

func main() {
	sendLocationRequest := *openapiclient.NewSendLocationRequest() // SendLocationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SendAPI.SendLocation(context.Background()).SendLocationRequest(sendLocationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.SendLocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendLocation`: SendResponse
	fmt.Fprintf(os.Stdout, "Response from `SendAPI.SendLocation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendLocationRequest** | [**SendLocationRequest**](SendLocationRequest.md) |  | 

### Return type

[**SendResponse**](SendResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendMessage

> SendResponse SendMessage(ctx).SendMessageRequest(sendMessageRequest).Execute()

Send Message

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aldinokemal/sdk-go-whatsapp-web-multidevice"
)

func main() {
	sendMessageRequest := *openapiclient.NewSendMessageRequest() // SendMessageRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SendAPI.SendMessage(context.Background()).SendMessageRequest(sendMessageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.SendMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendMessage`: SendResponse
	fmt.Fprintf(os.Stdout, "Response from `SendAPI.SendMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendMessageRequest** | [**SendMessageRequest**](SendMessageRequest.md) |  | 

### Return type

[**SendResponse**](SendResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendPoll

> SendResponse SendPoll(ctx).SendPollRequest(sendPollRequest).Execute()

Send Poll / Vote

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aldinokemal/sdk-go-whatsapp-web-multidevice"
)

func main() {
	sendPollRequest := *openapiclient.NewSendPollRequest("6289685024421@s.whatsapp.net", "Siapa Nama Avatar The Last Air Bender?", []string{"Options_example"}, int32(2)) // SendPollRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SendAPI.SendPoll(context.Background()).SendPollRequest(sendPollRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.SendPoll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendPoll`: SendResponse
	fmt.Fprintf(os.Stdout, "Response from `SendAPI.SendPoll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendPollRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendPollRequest** | [**SendPollRequest**](SendPollRequest.md) |  | 

### Return type

[**SendResponse**](SendResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendPresence

> SendResponse SendPresence(ctx).SendPresenceRequest(sendPresenceRequest).Execute()

Send presence status

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aldinokemal/sdk-go-whatsapp-web-multidevice"
)

func main() {
	sendPresenceRequest := *openapiclient.NewSendPresenceRequest("available") // SendPresenceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SendAPI.SendPresence(context.Background()).SendPresenceRequest(sendPresenceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.SendPresence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendPresence`: SendResponse
	fmt.Fprintf(os.Stdout, "Response from `SendAPI.SendPresence`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendPresenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendPresenceRequest** | [**SendPresenceRequest**](SendPresenceRequest.md) |  | 

### Return type

[**SendResponse**](SendResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendSticker

> SendResponse SendSticker(ctx).Phone(phone).Sticker(sticker).StickerUrl(stickerUrl).Duration(duration).IsForwarded(isForwarded).Execute()

Send Sticker



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aldinokemal/sdk-go-whatsapp-web-multidevice"
)

func main() {
	phone := "phone_example" // string | Phone number with country code (optional)
	sticker := os.NewFile(1234, "some_file") // *os.File | Sticker image file (jpg/jpeg/png/webp/gif) (optional)
	stickerUrl := "stickerUrl_example" // string | URL of sticker image to send (optional)
	duration := int32(56) // int32 | Disappearing message duration in seconds (optional) (optional)
	isForwarded := true // bool | Whether this is a forwarded sticker (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SendAPI.SendSticker(context.Background()).Phone(phone).Sticker(sticker).StickerUrl(stickerUrl).Duration(duration).IsForwarded(isForwarded).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.SendSticker``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendSticker`: SendResponse
	fmt.Fprintf(os.Stdout, "Response from `SendAPI.SendSticker`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendStickerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phone** | **string** | Phone number with country code | 
 **sticker** | ***os.File** | Sticker image file (jpg/jpeg/png/webp/gif) | 
 **stickerUrl** | **string** | URL of sticker image to send | 
 **duration** | **int32** | Disappearing message duration in seconds (optional) | 
 **isForwarded** | **bool** | Whether this is a forwarded sticker | 

### Return type

[**SendResponse**](SendResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendVideo

> SendResponse SendVideo(ctx).Phone(phone).Caption(caption).ViewOnce(viewOnce).Video(video).VideoUrl(videoUrl).Compress(compress).Duration(duration).IsForwarded(isForwarded).Execute()

Send Video

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/aldinokemal/sdk-go-whatsapp-web-multidevice"
)

func main() {
	phone := "phone_example" // string | Phone number with country code (optional)
	caption := "caption_example" // string | Caption to send (optional)
	viewOnce := true // bool | View once (optional)
	video := os.NewFile(1234, "some_file") // *os.File | Video to send (optional)
	videoUrl := "videoUrl_example" // string | Video URL to send (optional)
	compress := true // bool | Compress video (optional)
	duration := int32(56) // int32 | Disappearing message duration in seconds (optional) (optional)
	isForwarded := true // bool | Whether this is a forwarded message (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SendAPI.SendVideo(context.Background()).Phone(phone).Caption(caption).ViewOnce(viewOnce).Video(video).VideoUrl(videoUrl).Compress(compress).Duration(duration).IsForwarded(isForwarded).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.SendVideo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendVideo`: SendResponse
	fmt.Fprintf(os.Stdout, "Response from `SendAPI.SendVideo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendVideoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phone** | **string** | Phone number with country code | 
 **caption** | **string** | Caption to send | 
 **viewOnce** | **bool** | View once | 
 **video** | ***os.File** | Video to send | 
 **videoUrl** | **string** | Video URL to send | 
 **compress** | **bool** | Compress video | 
 **duration** | **int32** | Disappearing message duration in seconds (optional) | 
 **isForwarded** | **bool** | Whether this is a forwarded message | 

### Return type

[**SendResponse**](SendResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

