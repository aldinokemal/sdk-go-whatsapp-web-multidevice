# \SendAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendAudio**](SendAPI.md#SendAudio) | **Post** /send/audio | Send Audio
[**SendContact**](SendAPI.md#SendContact) | **Post** /send/contact | Send Contact
[**SendFile**](SendAPI.md#SendFile) | **Post** /send/file | Send File
[**SendImage**](SendAPI.md#SendImage) | **Post** /send/image | Send Image
[**SendLink**](SendAPI.md#SendLink) | **Post** /send/link | Send Link
[**SendLocation**](SendAPI.md#SendLocation) | **Post** /send/location | Send Location
[**SendMessage**](SendAPI.md#SendMessage) | **Post** /send/message | Send Message
[**SendPoll**](SendAPI.md#SendPoll) | **Post** /send/poll | Send Poll / Vote
[**SendVideo**](SendAPI.md#SendVideo) | **Post** /send/video | Send Video



## SendAudio

> SendResponse SendAudio(ctx).Phone(phone).Audio(audio).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SendAPI.SendAudio(context.Background()).Phone(phone).Audio(audio).Execute()
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

### Return type

[**SendResponse**](SendResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendContact

> SendResponse SendContact(ctx).Phone(phone).ContactName(contactName).ContactPhone(contactPhone).Execute()

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
	phone := "phone_example" // string | Phone number with country code (optional)
	contactName := "contactName_example" // string | Contact name (optional)
	contactPhone := "contactPhone_example" // string | Contact phone number (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SendAPI.SendContact(context.Background()).Phone(phone).ContactName(contactName).ContactPhone(contactPhone).Execute()
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
 **phone** | **string** | Phone number with country code | 
 **contactName** | **string** | Contact name | 
 **contactPhone** | **string** | Contact phone number | 

### Return type

[**SendResponse**](SendResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendFile

> SendResponse SendFile(ctx).Phone(phone).Caption(caption).File(file).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SendAPI.SendFile(context.Background()).Phone(phone).Caption(caption).File(file).Execute()
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

### Return type

[**SendResponse**](SendResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendImage

> SendResponse SendImage(ctx).Phone(phone).Caption(caption).ViewOnce(viewOnce).Image(image).Compress(compress).Execute()

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
	compress := true // bool | Compress image (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SendAPI.SendImage(context.Background()).Phone(phone).Caption(caption).ViewOnce(viewOnce).Image(image).Compress(compress).Execute()
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
 **compress** | **bool** | Compress image | 

### Return type

[**SendResponse**](SendResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendLink

> SendResponse SendLink(ctx).Phone(phone).Link(link).Caption(caption).Execute()

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
	phone := "phone_example" // string | Phone number with country code (optional)
	link := "link_example" // string | Link to send (optional)
	caption := "caption_example" // string | Caption to send (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SendAPI.SendLink(context.Background()).Phone(phone).Link(link).Caption(caption).Execute()
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
 **phone** | **string** | Phone number with country code | 
 **link** | **string** | Link to send | 
 **caption** | **string** | Caption to send | 

### Return type

[**SendResponse**](SendResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendLocation

> SendResponse SendLocation(ctx).Phone(phone).Latitude(latitude).Longitude(longitude).Execute()

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
	phone := "phone_example" // string | Phone number with country code (optional)
	latitude := "latitude_example" // string | Latitude coordinate (optional)
	longitude := "longitude_example" // string | Longitude coordinate (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SendAPI.SendLocation(context.Background()).Phone(phone).Latitude(latitude).Longitude(longitude).Execute()
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
 **phone** | **string** | Phone number with country code | 
 **latitude** | **string** | Latitude coordinate | 
 **longitude** | **string** | Longitude coordinate | 

### Return type

[**SendResponse**](SendResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendMessage

> SendResponse SendMessage(ctx).Phone(phone).Message(message).ReplyMessageId(replyMessageId).Execute()

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
	phone := "phone_example" // string | Phone number with country code (optional)
	message := "message_example" // string | Message to send (optional)
	replyMessageId := "replyMessageId_example" // string | Message ID that you want reply (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SendAPI.SendMessage(context.Background()).Phone(phone).Message(message).ReplyMessageId(replyMessageId).Execute()
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
 **phone** | **string** | Phone number with country code | 
 **message** | **string** | Message to send | 
 **replyMessageId** | **string** | Message ID that you want reply | 

### Return type

[**SendResponse**](SendResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendVideo

> SendResponse SendVideo(ctx).Phone(phone).Caption(caption).ViewOnce(viewOnce).Video(video).Compress(compress).Execute()

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
	compress := true // bool | Compress video (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SendAPI.SendVideo(context.Background()).Phone(phone).Caption(caption).ViewOnce(viewOnce).Video(video).Compress(compress).Execute()
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
 **compress** | **bool** | Compress video | 

### Return type

[**SendResponse**](SendResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

