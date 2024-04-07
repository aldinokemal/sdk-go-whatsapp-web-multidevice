# \MessageAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReactMessage**](MessageAPI.md#ReactMessage) | **Post** /message/{message_id}/reaction | Send reaction to message
[**RevokeMessage**](MessageAPI.md#RevokeMessage) | **Post** /message/{message_id}/revoke | Revoke Message
[**UpdateMessage**](MessageAPI.md#UpdateMessage) | **Post** /message/{message_id}/update | Edit message by message ID before 15 minutes



## ReactMessage

> SendResponse ReactMessage(ctx, messageId).Phone(phone).Emoji(emoji).Execute()

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
	phone := "phone_example" // string | Phone number with country code (optional)
	emoji := "emoji_example" // string | Emoji to react (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageAPI.ReactMessage(context.Background(), messageId).Phone(phone).Emoji(emoji).Execute()
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

 **phone** | **string** | Phone number with country code | 
 **emoji** | **string** | Emoji to react | 

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


## RevokeMessage

> SendResponse RevokeMessage(ctx, messageId).Phone(phone).Execute()

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
	phone := "phone_example" // string | Phone number with country code (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageAPI.RevokeMessage(context.Background(), messageId).Phone(phone).Execute()
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

 **phone** | **string** | Phone number with country code | 

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


## UpdateMessage

> SendResponse UpdateMessage(ctx, messageId).UpdateMessageRequest(updateMessageRequest).Execute()

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
	updateMessageRequest := *openapiclient.NewUpdateMessageRequest("62819273192397132@s.whatsapp.net", "Hello World") // UpdateMessageRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageAPI.UpdateMessage(context.Background(), messageId).UpdateMessageRequest(updateMessageRequest).Execute()
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

 **updateMessageRequest** | [**UpdateMessageRequest**](UpdateMessageRequest.md) |  | 

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

