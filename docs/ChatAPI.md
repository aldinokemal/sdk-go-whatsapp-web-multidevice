# SdkWhatsappWebMultiDevice\ChatAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetChatMessages**](ChatAPI.md#GetChatMessages) | **Get** /chat/{chat_jid}/messages | Get messages from a specific chat
[**LabelChat**](ChatAPI.md#LabelChat) | **Post** /chat/{chat_jid}/label | Label or unlabel a chat
[**ListChats**](ChatAPI.md#ListChats) | **Get** /chats | Get list of chats
[**PinChat**](ChatAPI.md#PinChat) | **Post** /chat/{chat_jid}/pin | Pin or unpin a chat



## GetChatMessages

> ChatMessagesResponse GetChatMessages(ctx, chatJid).Limit(limit).Offset(offset).StartTime(startTime).EndTime(endTime).MediaOnly(mediaOnly).IsFromMe(isFromMe).Search(search).Execute()

Get messages from a specific chat



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/aldinokemal/sdk-go-whatsapp-web-multidevice"
)

func main() {
	chatJid := "6289685028129@s.whatsapp.net" // string | Chat JID (e.g., phone@s.whatsapp.net for individual or groupid@g.us for group)
	limit := int32(56) // int32 | Maximum number of messages to return (optional) (default to 50)
	offset := int32(56) // int32 | Number of messages to skip (for pagination) (optional) (default to 0)
	startTime := time.Now() // time.Time | Filter messages from this timestamp (ISO 8601 format) (optional)
	endTime := time.Now() // time.Time | Filter messages until this timestamp (ISO 8601 format) (optional)
	mediaOnly := true // bool | Only return messages with media content (optional) (default to false)
	isFromMe := true // bool | Filter messages by sender (true for messages sent by you, false for received messages). When both media_only=true and isFromMe=false are provided, media_only takes precedence and will return all media messages regardless of sender. (optional)
	search := "search_example" // string | Search messages by content text (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAPI.GetChatMessages(context.Background(), chatJid).Limit(limit).Offset(offset).StartTime(startTime).EndTime(endTime).MediaOnly(mediaOnly).IsFromMe(isFromMe).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAPI.GetChatMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChatMessages`: ChatMessagesResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatAPI.GetChatMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatJid** | **string** | Chat JID (e.g., phone@s.whatsapp.net for individual or groupid@g.us for group) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChatMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Maximum number of messages to return | [default to 50]
 **offset** | **int32** | Number of messages to skip (for pagination) | [default to 0]
 **startTime** | **time.Time** | Filter messages from this timestamp (ISO 8601 format) | 
 **endTime** | **time.Time** | Filter messages until this timestamp (ISO 8601 format) | 
 **mediaOnly** | **bool** | Only return messages with media content | [default to false]
 **isFromMe** | **bool** | Filter messages by sender (true for messages sent by you, false for received messages). When both media_only&#x3D;true and isFromMe&#x3D;false are provided, media_only takes precedence and will return all media messages regardless of sender. | 
 **search** | **string** | Search messages by content text | 

### Return type

[**ChatMessagesResponse**](ChatMessagesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelChat

> LabelChatResponse LabelChat(ctx, chatJid).LabelChatRequest(labelChatRequest).Execute()

Label or unlabel a chat



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
	chatJid := "6289685028129@s.whatsapp.net" // string | Chat JID (e.g., phone@s.whatsapp.net for individual or groupid@g.us for group)
	labelChatRequest := *openapiclient.NewLabelChatRequest("label_123", "Important", true) // LabelChatRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAPI.LabelChat(context.Background(), chatJid).LabelChatRequest(labelChatRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAPI.LabelChat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LabelChat`: LabelChatResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatAPI.LabelChat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatJid** | **string** | Chat JID (e.g., phone@s.whatsapp.net for individual or groupid@g.us for group) | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **labelChatRequest** | [**LabelChatRequest**](LabelChatRequest.md) |  | 

### Return type

[**LabelChatResponse**](LabelChatResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListChats

> ChatListResponse ListChats(ctx).Limit(limit).Offset(offset).Search(search).HasMedia(hasMedia).Execute()

Get list of chats



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
	limit := int32(56) // int32 | Maximum number of chats to return (optional) (default to 25)
	offset := int32(56) // int32 | Number of chats to skip (for pagination) (optional) (default to 0)
	search := "search_example" // string | Search chats by name (optional)
	hasMedia := true // bool | Filter chats that contain media messages (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAPI.ListChats(context.Background()).Limit(limit).Offset(offset).Search(search).HasMedia(hasMedia).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAPI.ListChats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListChats`: ChatListResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatAPI.ListChats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListChatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of chats to return | [default to 25]
 **offset** | **int32** | Number of chats to skip (for pagination) | [default to 0]
 **search** | **string** | Search chats by name | 
 **hasMedia** | **bool** | Filter chats that contain media messages | [default to false]

### Return type

[**ChatListResponse**](ChatListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PinChat

> PinChatResponse PinChat(ctx, chatJid).PinChatRequest(pinChatRequest).Execute()

Pin or unpin a chat



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
	chatJid := "6289685028129@s.whatsapp.net" // string | Chat JID (e.g., phone@s.whatsapp.net for individual or groupid@g.us for group)
	pinChatRequest := *openapiclient.NewPinChatRequest(true) // PinChatRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAPI.PinChat(context.Background(), chatJid).PinChatRequest(pinChatRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAPI.PinChat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PinChat`: PinChatResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatAPI.PinChat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatJid** | **string** | Chat JID (e.g., phone@s.whatsapp.net for individual or groupid@g.us for group) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPinChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pinChatRequest** | [**PinChatRequest**](PinChatRequest.md) |  | 

### Return type

[**PinChatResponse**](PinChatResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

