# \GroupApi

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**JoinGroupWithLink**](GroupApi.md#JoinGroupWithLink) | **Post** /group/join-with-link | Join group with link



## JoinGroupWithLink

> SendResponse JoinGroupWithLink(ctx).Link(link).Execute()

Join group with link

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
    link := "link_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupApi.JoinGroupWithLink(context.Background()).Link(link).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.JoinGroupWithLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `JoinGroupWithLink`: SendResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.JoinGroupWithLink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiJoinGroupWithLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **link** | **string** |  | 

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

