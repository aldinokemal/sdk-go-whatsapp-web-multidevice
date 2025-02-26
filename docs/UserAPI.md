# SdkWhatsappWebMultiDevice\UserAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserAvatar**](UserAPI.md#UserAvatar) | **Get** /user/avatar | User Avatar
[**UserChangeAvatar**](UserAPI.md#UserChangeAvatar) | **Post** /user/avatar | User Change Avatar
[**UserInfo**](UserAPI.md#UserInfo) | **Get** /user/info | User Info
[**UserMyContacts**](UserAPI.md#UserMyContacts) | **Get** /user/my/contacts | Get list of user contacts
[**UserMyGroups**](UserAPI.md#UserMyGroups) | **Get** /user/my/groups | User My List Groups
[**UserMyNewsletter**](UserAPI.md#UserMyNewsletter) | **Get** /user/my/newsletters | User My List Groups
[**UserMyPrivacy**](UserAPI.md#UserMyPrivacy) | **Get** /user/my/privacy | User My Privacy Setting



## UserAvatar

> UserAvatarResponse UserAvatar(ctx).Phone(phone).IsPreview(isPreview).Execute()

User Avatar

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
	phone := "6289685028129@s.whatsapp.net" // string | Phone number with country code (optional)
	isPreview := true // bool | Whether to fetch a preview of the avatar (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserAvatar(context.Background()).Phone(phone).IsPreview(isPreview).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserAvatar`: UserAvatarResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserAvatar`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phone** | **string** | Phone number with country code | 
 **isPreview** | **bool** | Whether to fetch a preview of the avatar | 

### Return type

[**UserAvatarResponse**](UserAvatarResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserChangeAvatar

> GenericResponse UserChangeAvatar(ctx).Avatar(avatar).Execute()

User Change Avatar

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
	avatar := os.NewFile(1234, "some_file") // *os.File | Avatar to send (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserChangeAvatar(context.Background()).Avatar(avatar).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserChangeAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserChangeAvatar`: GenericResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserChangeAvatar`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserChangeAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **avatar** | ***os.File** | Avatar to send | 

### Return type

[**GenericResponse**](GenericResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserInfo

> UserInfoResponse UserInfo(ctx).Phone(phone).Execute()

User Info

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
	phone := "6289685028129@s.whatsapp.net" // string | Phone number with country code (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserInfo(context.Background()).Phone(phone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserInfo`: UserInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phone** | **string** | Phone number with country code | 

### Return type

[**UserInfoResponse**](UserInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserMyContacts

> MyListContactsResponse UserMyContacts(ctx).Execute()

Get list of user contacts

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserMyContacts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserMyContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserMyContacts`: MyListContactsResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserMyContacts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserMyContactsRequest struct via the builder pattern


### Return type

[**MyListContactsResponse**](MyListContactsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserMyGroups

> UserGroupResponse UserMyGroups(ctx).Execute()

User My List Groups

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserMyGroups(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserMyGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserMyGroups`: UserGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserMyGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserMyGroupsRequest struct via the builder pattern


### Return type

[**UserGroupResponse**](UserGroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserMyNewsletter

> NewsletterResponse UserMyNewsletter(ctx).Execute()

User My List Groups

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserMyNewsletter(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserMyNewsletter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserMyNewsletter`: NewsletterResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserMyNewsletter`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserMyNewsletterRequest struct via the builder pattern


### Return type

[**NewsletterResponse**](NewsletterResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserMyPrivacy

> UserPrivacyResponse UserMyPrivacy(ctx).Execute()

User My Privacy Setting

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserMyPrivacy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserMyPrivacy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserMyPrivacy`: UserPrivacyResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserMyPrivacy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserMyPrivacyRequest struct via the builder pattern


### Return type

[**UserPrivacyResponse**](UserPrivacyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

