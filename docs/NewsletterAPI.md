# SdkWhatsappWebMultiDevice\NewsletterAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UnfollowNewsletter**](NewsletterAPI.md#UnfollowNewsletter) | **Post** /newsletter/unfollow | Unfollow newsletter



## UnfollowNewsletter

> GenericResponse UnfollowNewsletter(ctx).XDeviceId(xDeviceId).UnfollowNewsletterRequest(unfollowNewsletterRequest).Execute()

Unfollow newsletter

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
	xDeviceId := "my-device-id" // string | Device identifier for multi-device support. Required when multiple devices are registered. If only one device is registered, it will be used as the default. Can also be provided as `device_id` query parameter.  (optional)
	unfollowNewsletterRequest := *openapiclient.NewUnfollowNewsletterRequest() // UnfollowNewsletterRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NewsletterAPI.UnfollowNewsletter(context.Background()).XDeviceId(xDeviceId).UnfollowNewsletterRequest(unfollowNewsletterRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NewsletterAPI.UnfollowNewsletter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnfollowNewsletter`: GenericResponse
	fmt.Fprintf(os.Stdout, "Response from `NewsletterAPI.UnfollowNewsletter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnfollowNewsletterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xDeviceId** | **string** | Device identifier for multi-device support. Required when multiple devices are registered. If only one device is registered, it will be used as the default. Can also be provided as &#x60;device_id&#x60; query parameter.  | 
 **unfollowNewsletterRequest** | [**UnfollowNewsletterRequest**](UnfollowNewsletterRequest.md) |  | 

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

