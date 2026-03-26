# SdkWhatsappWebMultiDevice\MessageAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMessage**](MessageAPI.md#DeleteMessage) | **Post** /message/{message_id}/delete | Delete Message
[**DownloadMessageMedia**](MessageAPI.md#DownloadMessageMedia) | **Get** /message/{message_id}/download | Download media from message
[**ReactMessage**](MessageAPI.md#ReactMessage) | **Post** /message/{message_id}/reaction | Send reaction to message
[**ReadMessage**](MessageAPI.md#ReadMessage) | **Post** /message/{message_id}/read | Mark as read message
[**RevokeMessage**](MessageAPI.md#RevokeMessage) | **Post** /message/{message_id}/revoke | Revoke Message
[**StarMessage**](MessageAPI.md#StarMessage) | **Post** /message/{message_id}/star | Star message
[**UnstarMessage**](MessageAPI.md#UnstarMessage) | **Post** /message/{message_id}/unstar | Unstar message
[**UpdateMessage**](MessageAPI.md#UpdateMessage) | **Post** /message/{message_id}/update | Edit message by message ID before 15 minutes



## DeleteMessage

> SendResponse DeleteMessage(ctx, messageId).XDeviceId(xDeviceId).RevokeMessageRequest(revokeMessageRequest).Execute()

Delete Message

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
	messageId := "messageId_example" // string | Message ID
	xDeviceId := "my-device-id" // string | Device identifier for multi-device support. Required when multiple devices are registered. If only one device is registered, it will be used as the default. Can also be provided as `device_id` query parameter.  (optional)
	revokeMessageRequest := *openapiclient.NewRevokeMessageRequest() // RevokeMessageRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageAPI.DeleteMessage(context.Background(), messageId).XDeviceId(xDeviceId).RevokeMessageRequest(revokeMessageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageAPI.DeleteMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMessage`: SendResponse
	fmt.Fprintf(os.Stdout, "Response from `MessageAPI.DeleteMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | Message ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xDeviceId** | **string** | Device identifier for multi-device support. Required when multiple devices are registered. If only one device is registered, it will be used as the default. Can also be provided as &#x60;device_id&#x60; query parameter.  | 
 **revokeMessageRequest** | [**RevokeMessageRequest**](RevokeMessageRequest.md) |  | 

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


## DownloadMessageMedia

> DownloadMessageMedia200Response DownloadMessageMedia(ctx, messageId).Phone(phone).XDeviceId(xDeviceId).Execute()

Download media from message



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
	messageId := "3EB0123456789ABCDEF" // string | Message ID
	phone := "6289685028129@s.whatsapp.net" // string | Phone number with country code
	xDeviceId := "my-device-id" // string | Device identifier for multi-device support. Required when multiple devices are registered. If only one device is registered, it will be used as the default. Can also be provided as `device_id` query parameter.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageAPI.DownloadMessageMedia(context.Background(), messageId).Phone(phone).XDeviceId(xDeviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageAPI.DownloadMessageMedia``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadMessageMedia`: DownloadMessageMedia200Response
	fmt.Fprintf(os.Stdout, "Response from `MessageAPI.DownloadMessageMedia`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | Message ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadMessageMediaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **phone** | **string** | Phone number with country code | 
 **xDeviceId** | **string** | Device identifier for multi-device support. Required when multiple devices are registered. If only one device is registered, it will be used as the default. Can also be provided as &#x60;device_id&#x60; query parameter.  | 

### Return type

[**DownloadMessageMedia200Response**](DownloadMessageMedia200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReactMessage

> SendResponse ReactMessage(ctx, messageId).XDeviceId(xDeviceId).ReactMessageRequest(reactMessageRequest).Execute()

Send reaction to message

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
	messageId := "messageId_example" // string | Message ID
	xDeviceId := "my-device-id" // string | Device identifier for multi-device support. Required when multiple devices are registered. If only one device is registered, it will be used as the default. Can also be provided as `device_id` query parameter.  (optional)
	reactMessageRequest := *openapiclient.NewReactMessageRequest() // ReactMessageRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageAPI.ReactMessage(context.Background(), messageId).XDeviceId(xDeviceId).ReactMessageRequest(reactMessageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageAPI.ReactMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReactMessage`: SendResponse
	fmt.Fprintf(os.Stdout, "Response from `MessageAPI.ReactMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | Message ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiReactMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xDeviceId** | **string** | Device identifier for multi-device support. Required when multiple devices are registered. If only one device is registered, it will be used as the default. Can also be provided as &#x60;device_id&#x60; query parameter.  | 
 **reactMessageRequest** | [**ReactMessageRequest**](ReactMessageRequest.md) |  | 

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


## ReadMessage

> SendResponse ReadMessage(ctx, messageId).XDeviceId(xDeviceId).ReadMessageRequest(readMessageRequest).Execute()

Mark as read message

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
	messageId := "messageId_example" // string | Message ID
	xDeviceId := "my-device-id" // string | Device identifier for multi-device support. Required when multiple devices are registered. If only one device is registered, it will be used as the default. Can also be provided as `device_id` query parameter.  (optional)
	readMessageRequest := *openapiclient.NewReadMessageRequest("62819273192397132@s.whatsapp.net") // ReadMessageRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageAPI.ReadMessage(context.Background(), messageId).XDeviceId(xDeviceId).ReadMessageRequest(readMessageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageAPI.ReadMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadMessage`: SendResponse
	fmt.Fprintf(os.Stdout, "Response from `MessageAPI.ReadMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | Message ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xDeviceId** | **string** | Device identifier for multi-device support. Required when multiple devices are registered. If only one device is registered, it will be used as the default. Can also be provided as &#x60;device_id&#x60; query parameter.  | 
 **readMessageRequest** | [**ReadMessageRequest**](ReadMessageRequest.md) |  | 

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


## RevokeMessage

> SendResponse RevokeMessage(ctx, messageId).XDeviceId(xDeviceId).RevokeMessageRequest(revokeMessageRequest).Execute()

Revoke Message

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
	messageId := "messageId_example" // string | Message ID
	xDeviceId := "my-device-id" // string | Device identifier for multi-device support. Required when multiple devices are registered. If only one device is registered, it will be used as the default. Can also be provided as `device_id` query parameter.  (optional)
	revokeMessageRequest := *openapiclient.NewRevokeMessageRequest() // RevokeMessageRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageAPI.RevokeMessage(context.Background(), messageId).XDeviceId(xDeviceId).RevokeMessageRequest(revokeMessageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageAPI.RevokeMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RevokeMessage`: SendResponse
	fmt.Fprintf(os.Stdout, "Response from `MessageAPI.RevokeMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | Message ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xDeviceId** | **string** | Device identifier for multi-device support. Required when multiple devices are registered. If only one device is registered, it will be used as the default. Can also be provided as &#x60;device_id&#x60; query parameter.  | 
 **revokeMessageRequest** | [**RevokeMessageRequest**](RevokeMessageRequest.md) |  | 

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


## StarMessage

> GenericResponse StarMessage(ctx, messageId).XDeviceId(xDeviceId).ReadMessageRequest(readMessageRequest).Execute()

Star message

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
	messageId := "messageId_example" // string | Message ID
	xDeviceId := "my-device-id" // string | Device identifier for multi-device support. Required when multiple devices are registered. If only one device is registered, it will be used as the default. Can also be provided as `device_id` query parameter.  (optional)
	readMessageRequest := *openapiclient.NewReadMessageRequest("62819273192397132@s.whatsapp.net") // ReadMessageRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageAPI.StarMessage(context.Background(), messageId).XDeviceId(xDeviceId).ReadMessageRequest(readMessageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageAPI.StarMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StarMessage`: GenericResponse
	fmt.Fprintf(os.Stdout, "Response from `MessageAPI.StarMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | Message ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStarMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xDeviceId** | **string** | Device identifier for multi-device support. Required when multiple devices are registered. If only one device is registered, it will be used as the default. Can also be provided as &#x60;device_id&#x60; query parameter.  | 
 **readMessageRequest** | [**ReadMessageRequest**](ReadMessageRequest.md) |  | 

### Return type

[**GenericResponse**](GenericResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnstarMessage

> GenericResponse UnstarMessage(ctx, messageId).XDeviceId(xDeviceId).ReadMessageRequest(readMessageRequest).Execute()

Unstar message

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
	messageId := "messageId_example" // string | Message ID
	xDeviceId := "my-device-id" // string | Device identifier for multi-device support. Required when multiple devices are registered. If only one device is registered, it will be used as the default. Can also be provided as `device_id` query parameter.  (optional)
	readMessageRequest := *openapiclient.NewReadMessageRequest("62819273192397132@s.whatsapp.net") // ReadMessageRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageAPI.UnstarMessage(context.Background(), messageId).XDeviceId(xDeviceId).ReadMessageRequest(readMessageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageAPI.UnstarMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnstarMessage`: GenericResponse
	fmt.Fprintf(os.Stdout, "Response from `MessageAPI.UnstarMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | Message ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnstarMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xDeviceId** | **string** | Device identifier for multi-device support. Required when multiple devices are registered. If only one device is registered, it will be used as the default. Can also be provided as &#x60;device_id&#x60; query parameter.  | 
 **readMessageRequest** | [**ReadMessageRequest**](ReadMessageRequest.md) |  | 

### Return type

[**GenericResponse**](GenericResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMessage

> SendResponse UpdateMessage(ctx, messageId).XDeviceId(xDeviceId).UpdateMessageRequest(updateMessageRequest).Execute()

Edit message by message ID before 15 minutes

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
	messageId := "messageId_example" // string | Message ID
	xDeviceId := "my-device-id" // string | Device identifier for multi-device support. Required when multiple devices are registered. If only one device is registered, it will be used as the default. Can also be provided as `device_id` query parameter.  (optional)
	updateMessageRequest := *openapiclient.NewUpdateMessageRequest("62819273192397132@s.whatsapp.net", "Hello World") // UpdateMessageRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageAPI.UpdateMessage(context.Background(), messageId).XDeviceId(xDeviceId).UpdateMessageRequest(updateMessageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageAPI.UpdateMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMessage`: SendResponse
	fmt.Fprintf(os.Stdout, "Response from `MessageAPI.UpdateMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | Message ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xDeviceId** | **string** | Device identifier for multi-device support. Required when multiple devices are registered. If only one device is registered, it will be used as the default. Can also be provided as &#x60;device_id&#x60; query parameter.  | 
 **updateMessageRequest** | [**UpdateMessageRequest**](UpdateMessageRequest.md) |  | 

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

