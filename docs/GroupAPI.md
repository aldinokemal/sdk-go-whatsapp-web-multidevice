# \GroupAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddParticipantToGroup**](GroupAPI.md#AddParticipantToGroup) | **Post** /group/participants | Adding more participants to group
[**CreateGroup**](GroupAPI.md#CreateGroup) | **Post** /group | Create group and add participant
[**DemoteParticipantToMember**](GroupAPI.md#DemoteParticipantToMember) | **Post** /group/participants/demote | Demote participants to member
[**JoinGroupWithLink**](GroupAPI.md#JoinGroupWithLink) | **Post** /group/join-with-link | Join group with link
[**LeaveGroup**](GroupAPI.md#LeaveGroup) | **Post** /group/leave | Leave group
[**PromoteParticipantToAdmin**](GroupAPI.md#PromoteParticipantToAdmin) | **Post** /group/participants/promote | Promote participants to admin
[**RemoveParticipantFromGroup**](GroupAPI.md#RemoveParticipantFromGroup) | **Post** /group/participants/remove | Remove participants from group



## AddParticipantToGroup

> ManageParticipantResponse AddParticipantToGroup(ctx).ManageParticipantRequest(manageParticipantRequest).Execute()

Adding more participants to group

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
	manageParticipantRequest := *openapiclient.NewManageParticipantRequest() // ManageParticipantRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.AddParticipantToGroup(context.Background()).ManageParticipantRequest(manageParticipantRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.AddParticipantToGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddParticipantToGroup`: ManageParticipantResponse
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.AddParticipantToGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddParticipantToGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **manageParticipantRequest** | [**ManageParticipantRequest**](ManageParticipantRequest.md) |  | 

### Return type

[**ManageParticipantResponse**](ManageParticipantResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGroup

> CreateGroupResponse CreateGroup(ctx).CreateGroupRequest(createGroupRequest).Execute()

Create group and add participant

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
	createGroupRequest := *openapiclient.NewCreateGroupRequest() // CreateGroupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.CreateGroup(context.Background()).CreateGroupRequest(createGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.CreateGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGroup`: CreateGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.CreateGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createGroupRequest** | [**CreateGroupRequest**](CreateGroupRequest.md) |  | 

### Return type

[**CreateGroupResponse**](CreateGroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DemoteParticipantToMember

> ManageParticipantResponse DemoteParticipantToMember(ctx).ManageParticipantRequest(manageParticipantRequest).Execute()

Demote participants to member

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
	manageParticipantRequest := *openapiclient.NewManageParticipantRequest() // ManageParticipantRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.DemoteParticipantToMember(context.Background()).ManageParticipantRequest(manageParticipantRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.DemoteParticipantToMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DemoteParticipantToMember`: ManageParticipantResponse
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.DemoteParticipantToMember`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDemoteParticipantToMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **manageParticipantRequest** | [**ManageParticipantRequest**](ManageParticipantRequest.md) |  | 

### Return type

[**ManageParticipantResponse**](ManageParticipantResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JoinGroupWithLink

> GenericResponse JoinGroupWithLink(ctx).JoinGroupWithLinkRequest(joinGroupWithLinkRequest).Execute()

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
	joinGroupWithLinkRequest := *openapiclient.NewJoinGroupWithLinkRequest() // JoinGroupWithLinkRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.JoinGroupWithLink(context.Background()).JoinGroupWithLinkRequest(joinGroupWithLinkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.JoinGroupWithLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JoinGroupWithLink`: GenericResponse
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.JoinGroupWithLink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiJoinGroupWithLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **joinGroupWithLinkRequest** | [**JoinGroupWithLinkRequest**](JoinGroupWithLinkRequest.md) |  | 

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


## LeaveGroup

> GenericResponse LeaveGroup(ctx).LeaveGroupRequest(leaveGroupRequest).Execute()

Leave group

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
	leaveGroupRequest := *openapiclient.NewLeaveGroupRequest() // LeaveGroupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.LeaveGroup(context.Background()).LeaveGroupRequest(leaveGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.LeaveGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LeaveGroup`: GenericResponse
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.LeaveGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLeaveGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **leaveGroupRequest** | [**LeaveGroupRequest**](LeaveGroupRequest.md) |  | 

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


## PromoteParticipantToAdmin

> ManageParticipantResponse PromoteParticipantToAdmin(ctx).ManageParticipantRequest(manageParticipantRequest).Execute()

Promote participants to admin

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
	manageParticipantRequest := *openapiclient.NewManageParticipantRequest() // ManageParticipantRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.PromoteParticipantToAdmin(context.Background()).ManageParticipantRequest(manageParticipantRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.PromoteParticipantToAdmin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PromoteParticipantToAdmin`: ManageParticipantResponse
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.PromoteParticipantToAdmin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPromoteParticipantToAdminRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **manageParticipantRequest** | [**ManageParticipantRequest**](ManageParticipantRequest.md) |  | 

### Return type

[**ManageParticipantResponse**](ManageParticipantResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveParticipantFromGroup

> ManageParticipantResponse RemoveParticipantFromGroup(ctx).ManageParticipantRequest(manageParticipantRequest).Execute()

Remove participants from group

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
	manageParticipantRequest := *openapiclient.NewManageParticipantRequest() // ManageParticipantRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.RemoveParticipantFromGroup(context.Background()).ManageParticipantRequest(manageParticipantRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.RemoveParticipantFromGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveParticipantFromGroup`: ManageParticipantResponse
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.RemoveParticipantFromGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveParticipantFromGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **manageParticipantRequest** | [**ManageParticipantRequest**](ManageParticipantRequest.md) |  | 

### Return type

[**ManageParticipantResponse**](ManageParticipantResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

