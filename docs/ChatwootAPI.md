# SdkWhatsappWebMultiDevice\ChatwootAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatwootSyncHistory**](ChatwootAPI.md#ChatwootSyncHistory) | **Post** /chatwoot/sync | Sync message history to Chatwoot
[**ChatwootSyncStatus**](ChatwootAPI.md#ChatwootSyncStatus) | **Get** /chatwoot/sync/status | Get Chatwoot sync progress
[**ChatwootWebhook**](ChatwootAPI.md#ChatwootWebhook) | **Post** /chatwoot/webhook | Chatwoot webhook endpoint



## ChatwootSyncHistory

> ChatwootSyncResponse ChatwootSyncHistory(ctx).ChatwootSyncHistoryRequest(chatwootSyncHistoryRequest).Execute()

Sync message history to Chatwoot



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
	chatwootSyncHistoryRequest := *openapiclient.NewChatwootSyncHistoryRequest() // ChatwootSyncHistoryRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatwootAPI.ChatwootSyncHistory(context.Background()).ChatwootSyncHistoryRequest(chatwootSyncHistoryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatwootAPI.ChatwootSyncHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatwootSyncHistory`: ChatwootSyncResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatwootAPI.ChatwootSyncHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatwootSyncHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatwootSyncHistoryRequest** | [**ChatwootSyncHistoryRequest**](ChatwootSyncHistoryRequest.md) |  | 

### Return type

[**ChatwootSyncResponse**](ChatwootSyncResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatwootSyncStatus

> ChatwootSyncStatusResponse ChatwootSyncStatus(ctx).DeviceId(deviceId).Execute()

Get Chatwoot sync progress



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
	deviceId := "deviceId_example" // string | Device ID to check sync status for (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatwootAPI.ChatwootSyncStatus(context.Background()).DeviceId(deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatwootAPI.ChatwootSyncStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatwootSyncStatus`: ChatwootSyncStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatwootAPI.ChatwootSyncStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatwootSyncStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | Device ID to check sync status for | 

### Return type

[**ChatwootSyncStatusResponse**](ChatwootSyncStatusResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatwootWebhook

> ChatwootWebhook(ctx).Body(body).Execute()

Chatwoot webhook endpoint



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
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChatwootAPI.ChatwootWebhook(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatwootAPI.ChatwootWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatwootWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

