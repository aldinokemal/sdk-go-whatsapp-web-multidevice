# SdkWhatsappWebMultiDevice\NewsletterAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UnfollowNewsletter**](NewsletterAPI.md#UnfollowNewsletter) | **Post** /newsletter/unfollow | Unfollow newsletter



## UnfollowNewsletter

> GenericResponse UnfollowNewsletter(ctx).UnfollowNewsletterRequest(unfollowNewsletterRequest).Execute()

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
	unfollowNewsletterRequest := *openapiclient.NewUnfollowNewsletterRequest() // UnfollowNewsletterRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NewsletterAPI.UnfollowNewsletter(context.Background()).UnfollowNewsletterRequest(unfollowNewsletterRequest).Execute()
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
 **unfollowNewsletterRequest** | [**UnfollowNewsletterRequest**](UnfollowNewsletterRequest.md) |  | 

### Return type

[**GenericResponse**](GenericResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

