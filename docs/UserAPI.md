# SdkWhatsappWebMultiDevice\UserAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserAvatar**](UserAPI.md#UserAvatar) | **Get** /user/avatar | User Avatar
[**UserBusinessProfile**](UserAPI.md#UserBusinessProfile) | **Get** /user/business-profile | Get Business Profile Information
[**UserChangeAvatar**](UserAPI.md#UserChangeAvatar) | **Post** /user/avatar | User Change Avatar
[**UserChangePushName**](UserAPI.md#UserChangePushName) | **Post** /user/pushname | User Change Push Name
[**UserCheck**](UserAPI.md#UserCheck) | **Get** /user/check | Check if user is on WhatsApp
[**UserInfo**](UserAPI.md#UserInfo) | **Get** /user/info | User Info
[**UserMyContacts**](UserAPI.md#UserMyContacts) | **Get** /user/my/contacts | Get list of user contacts
[**UserMyGroups**](UserAPI.md#UserMyGroups) | **Get** /user/my/groups | User My List Groups
[**UserMyNewsletter**](UserAPI.md#UserMyNewsletter) | **Get** /user/my/newsletters | User My List Groups
[**UserMyPrivacy**](UserAPI.md#UserMyPrivacy) | **Get** /user/my/privacy | User My Privacy Setting



## UserAvatar

> UserAvatarResponse UserAvatar(ctx).XDeviceId(xDeviceId).Phone(phone).IsPreview(isPreview).IsCommunity(isCommunity).Execute()

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
	xDeviceId := "my-device-id" // string | Device identifier for multi-device support. Required when multiple devices are registered. If only one device is registered, it will be used as the default. Can also be provided as `device_id` query parameter.  (optional)
	phone := "6289685028129@s.whatsapp.net" // string | Phone number with country code (optional)
	isPreview := true // bool | Whether to fetch a preview of the avatar (optional)
	isCommunity := false // bool | Whether to fetch a community avatar (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserAvatar(context.Background()).XDeviceId(xDeviceId).Phone(phone).IsPreview(isPreview).IsCommunity(isCommunity).Execute()
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
 **xDeviceId** | **string** | Device identifier for multi-device support. Required when multiple devices are registered. If only one device is registered, it will be used as the default. Can also be provided as &#x60;device_id&#x60; query parameter.  | 
 **phone** | **string** | Phone number with country code | 
 **isPreview** | **bool** | Whether to fetch a preview of the avatar | 
 **isCommunity** | **bool** | Whether to fetch a community avatar | 

### Return type

[**UserAvatarResponse**](UserAvatarResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserBusinessProfile

> BusinessProfileResponse UserBusinessProfile(ctx).Phone(phone).XDeviceId(xDeviceId).Execute()

Get Business Profile Information



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
	phone := "6289685028129@s.whatsapp.net" // string | Phone number with country code of the business account
	xDeviceId := "my-device-id" // string | Device identifier for multi-device support. Required when multiple devices are registered. If only one device is registered, it will be used as the default. Can also be provided as `device_id` query parameter.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserBusinessProfile(context.Background()).Phone(phone).XDeviceId(xDeviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserBusinessProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserBusinessProfile`: BusinessProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserBusinessProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserBusinessProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phone** | **string** | Phone number with country code of the business account | 
 **xDeviceId** | **string** | Device identifier for multi-device support. Required when multiple devices are registered. If only one device is registered, it will be used as the default. Can also be provided as &#x60;device_id&#x60; query parameter.  | 

### Return type

[**BusinessProfileResponse**](BusinessProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserChangeAvatar

> GenericResponse UserChangeAvatar(ctx).XDeviceId(xDeviceId).Avatar(avatar).Execute()

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
	xDeviceId := "my-device-id" // string | Device identifier for multi-device support. Required when multiple devices are registered. If only one device is registered, it will be used as the default. Can also be provided as `device_id` query parameter.  (optional)
	avatar := os.NewFile(1234, "some_file") // *os.File | Avatar to send (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserChangeAvatar(context.Background()).XDeviceId(xDeviceId).Avatar(avatar).Execute()
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
 **xDeviceId** | **string** | Device identifier for multi-device support. Required when multiple devices are registered. If only one device is registered, it will be used as the default. Can also be provided as &#x60;device_id&#x60; query parameter.  | 
 **avatar** | ***os.File** | Avatar to send | 

### Return type

[**GenericResponse**](GenericResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserChangePushName

> GenericResponse UserChangePushName(ctx).XDeviceId(xDeviceId).UserChangePushNameRequest(userChangePushNameRequest).Execute()

User Change Push Name



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
	userChangePushNameRequest := *openapiclient.NewUserChangePushNameRequest("John Doe") // UserChangePushNameRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserChangePushName(context.Background()).XDeviceId(xDeviceId).UserChangePushNameRequest(userChangePushNameRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserChangePushName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserChangePushName`: GenericResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserChangePushName`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserChangePushNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xDeviceId** | **string** | Device identifier for multi-device support. Required when multiple devices are registered. If only one device is registered, it will be used as the default. Can also be provided as &#x60;device_id&#x60; query parameter.  | 
 **userChangePushNameRequest** | [**UserChangePushNameRequest**](UserChangePushNameRequest.md) |  | 

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


## UserCheck

> UserCheckResponse UserCheck(ctx).XDeviceId(xDeviceId).Phone(phone).Execute()

Check if user is on WhatsApp

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
	phone := "628912344551" // string | Phone number with country code (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserCheck(context.Background()).XDeviceId(xDeviceId).Phone(phone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserCheck`: UserCheckResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xDeviceId** | **string** | Device identifier for multi-device support. Required when multiple devices are registered. If only one device is registered, it will be used as the default. Can also be provided as &#x60;device_id&#x60; query parameter.  | 
 **phone** | **string** | Phone number with country code | 

### Return type

[**UserCheckResponse**](UserCheckResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserInfo

> UserInfoResponse UserInfo(ctx).XDeviceId(xDeviceId).Phone(phone).Execute()

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
	xDeviceId := "my-device-id" // string | Device identifier for multi-device support. Required when multiple devices are registered. If only one device is registered, it will be used as the default. Can also be provided as `device_id` query parameter.  (optional)
	phone := "6289685028129@s.whatsapp.net" // string | Phone number with country code (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserInfo(context.Background()).XDeviceId(xDeviceId).Phone(phone).Execute()
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
 **xDeviceId** | **string** | Device identifier for multi-device support. Required when multiple devices are registered. If only one device is registered, it will be used as the default. Can also be provided as &#x60;device_id&#x60; query parameter.  | 
 **phone** | **string** | Phone number with country code | 

### Return type

[**UserInfoResponse**](UserInfoResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserMyContacts

> MyListContactsResponse UserMyContacts(ctx).XDeviceId(xDeviceId).Execute()

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
	xDeviceId := "my-device-id" // string | Device identifier for multi-device support. Required when multiple devices are registered. If only one device is registered, it will be used as the default. Can also be provided as `device_id` query parameter.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserMyContacts(context.Background()).XDeviceId(xDeviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserMyContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserMyContacts`: MyListContactsResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserMyContacts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserMyContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xDeviceId** | **string** | Device identifier for multi-device support. Required when multiple devices are registered. If only one device is registered, it will be used as the default. Can also be provided as &#x60;device_id&#x60; query parameter.  | 

### Return type

[**MyListContactsResponse**](MyListContactsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserMyGroups

> UserGroupResponse UserMyGroups(ctx).XDeviceId(xDeviceId).Execute()

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
	xDeviceId := "my-device-id" // string | Device identifier for multi-device support. Required when multiple devices are registered. If only one device is registered, it will be used as the default. Can also be provided as `device_id` query parameter.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserMyGroups(context.Background()).XDeviceId(xDeviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserMyGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserMyGroups`: UserGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserMyGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserMyGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xDeviceId** | **string** | Device identifier for multi-device support. Required when multiple devices are registered. If only one device is registered, it will be used as the default. Can also be provided as &#x60;device_id&#x60; query parameter.  | 

### Return type

[**UserGroupResponse**](UserGroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserMyNewsletter

> NewsletterResponse UserMyNewsletter(ctx).XDeviceId(xDeviceId).Execute()

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
	xDeviceId := "my-device-id" // string | Device identifier for multi-device support. Required when multiple devices are registered. If only one device is registered, it will be used as the default. Can also be provided as `device_id` query parameter.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserMyNewsletter(context.Background()).XDeviceId(xDeviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserMyNewsletter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserMyNewsletter`: NewsletterResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserMyNewsletter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserMyNewsletterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xDeviceId** | **string** | Device identifier for multi-device support. Required when multiple devices are registered. If only one device is registered, it will be used as the default. Can also be provided as &#x60;device_id&#x60; query parameter.  | 

### Return type

[**NewsletterResponse**](NewsletterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserMyPrivacy

> UserPrivacyResponse UserMyPrivacy(ctx).XDeviceId(xDeviceId).Execute()

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
	xDeviceId := "my-device-id" // string | Device identifier for multi-device support. Required when multiple devices are registered. If only one device is registered, it will be used as the default. Can also be provided as `device_id` query parameter.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserMyPrivacy(context.Background()).XDeviceId(xDeviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserMyPrivacy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserMyPrivacy`: UserPrivacyResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserMyPrivacy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserMyPrivacyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xDeviceId** | **string** | Device identifier for multi-device support. Required when multiple devices are registered. If only one device is registered, it will be used as the default. Can also be provided as &#x60;device_id&#x60; query parameter.  | 

### Return type

[**UserPrivacyResponse**](UserPrivacyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

