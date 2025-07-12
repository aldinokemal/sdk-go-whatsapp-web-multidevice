# SdkWhatsappWebMultiDevice\GroupAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddParticipantToGroup**](GroupAPI.md#AddParticipantToGroup) | **Post** /group/participants | Adding more participants to group
[**ApproveGroupParticipantRequest**](GroupAPI.md#ApproveGroupParticipantRequest) | **Post** /group/participant-requests/approve | Approve participant request to join group
[**CreateGroup**](GroupAPI.md#CreateGroup) | **Post** /group | Create group and add participant
[**DemoteParticipantToMember**](GroupAPI.md#DemoteParticipantToMember) | **Post** /group/participants/demote | Demote participants to member
[**GetGroupInfoFromLink**](GroupAPI.md#GetGroupInfoFromLink) | **Get** /group/info-from-link | Get group information from invitation link
[**GetGroupParticipantRequests**](GroupAPI.md#GetGroupParticipantRequests) | **Get** /group/participant-requests | Get list of participant requests to join group
[**GroupInfo**](GroupAPI.md#GroupInfo) | **Get** /group/info | Group Info
[**JoinGroupWithLink**](GroupAPI.md#JoinGroupWithLink) | **Post** /group/join-with-link | Join group with link
[**LeaveGroup**](GroupAPI.md#LeaveGroup) | **Post** /group/leave | Leave group
[**PromoteParticipantToAdmin**](GroupAPI.md#PromoteParticipantToAdmin) | **Post** /group/participants/promote | Promote participants to admin
[**RejectGroupParticipantRequest**](GroupAPI.md#RejectGroupParticipantRequest) | **Post** /group/participant-requests/reject | Reject participant request to join group
[**RemoveParticipantFromGroup**](GroupAPI.md#RemoveParticipantFromGroup) | **Post** /group/participants/remove | Remove participants from group
[**SetGroupAnnounce**](GroupAPI.md#SetGroupAnnounce) | **Post** /group/announce | Set group announce mode
[**SetGroupLocked**](GroupAPI.md#SetGroupLocked) | **Post** /group/locked | Set group locked status
[**SetGroupName**](GroupAPI.md#SetGroupName) | **Post** /group/name | Set group name
[**SetGroupPhoto**](GroupAPI.md#SetGroupPhoto) | **Post** /group/photo | Set group photo
[**SetGroupTopic**](GroupAPI.md#SetGroupTopic) | **Post** /group/topic | Set group topic



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

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApproveGroupParticipantRequest

> GenericResponse ApproveGroupParticipantRequest(ctx).ApproveGroupParticipantRequestRequest(approveGroupParticipantRequestRequest).Execute()

Approve participant request to join group

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
	approveGroupParticipantRequestRequest := *openapiclient.NewApproveGroupParticipantRequestRequest("120363024512399999@g.us", []string{"Participants_example"}) // ApproveGroupParticipantRequestRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.ApproveGroupParticipantRequest(context.Background()).ApproveGroupParticipantRequestRequest(approveGroupParticipantRequestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.ApproveGroupParticipantRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApproveGroupParticipantRequest`: GenericResponse
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.ApproveGroupParticipantRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApproveGroupParticipantRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **approveGroupParticipantRequestRequest** | [**ApproveGroupParticipantRequestRequest**](ApproveGroupParticipantRequestRequest.md) |  | 

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

[basicAuth](../README.md#basicAuth)

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

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupInfoFromLink

> GroupInfoFromLinkResponse GetGroupInfoFromLink(ctx).Link(link).Execute()

Get group information from invitation link



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
	link := "https://chat.whatsapp.com/whatsappKeyJoinGroup" // string | WhatsApp group invitation link

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.GetGroupInfoFromLink(context.Background()).Link(link).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.GetGroupInfoFromLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupInfoFromLink`: GroupInfoFromLinkResponse
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.GetGroupInfoFromLink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupInfoFromLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **link** | **string** | WhatsApp group invitation link | 

### Return type

[**GroupInfoFromLinkResponse**](GroupInfoFromLinkResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupParticipantRequests

> GroupParticipantRequestListResponse GetGroupParticipantRequests(ctx).GroupId(groupId).Execute()

Get list of participant requests to join group

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
	groupId := "120363024512399999@g.us" // string | The group ID to get participant requests for

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.GetGroupParticipantRequests(context.Background()).GroupId(groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.GetGroupParticipantRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupParticipantRequests`: GroupParticipantRequestListResponse
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.GetGroupParticipantRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupParticipantRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **string** | The group ID to get participant requests for | 

### Return type

[**GroupParticipantRequestListResponse**](GroupParticipantRequestListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupInfo

> GroupInfoResponse GroupInfo(ctx).GroupId(groupId).Execute()

Group Info

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
	groupId := "120363025982934543@g.us" // string | WhatsApp Group ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.GroupInfo(context.Background()).GroupId(groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.GroupInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupInfo`: GroupInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.GroupInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **string** | WhatsApp Group ID | 

### Return type

[**GroupInfoResponse**](GroupInfoResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
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

[basicAuth](../README.md#basicAuth)

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

[basicAuth](../README.md#basicAuth)

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

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RejectGroupParticipantRequest

> GenericResponse RejectGroupParticipantRequest(ctx).RejectGroupParticipantRequestRequest(rejectGroupParticipantRequestRequest).Execute()

Reject participant request to join group

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
	rejectGroupParticipantRequestRequest := *openapiclient.NewRejectGroupParticipantRequestRequest("120363024512399999@g.us", []string{"Participants_example"}) // RejectGroupParticipantRequestRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.RejectGroupParticipantRequest(context.Background()).RejectGroupParticipantRequestRequest(rejectGroupParticipantRequestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.RejectGroupParticipantRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RejectGroupParticipantRequest`: GenericResponse
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.RejectGroupParticipantRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRejectGroupParticipantRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rejectGroupParticipantRequestRequest** | [**RejectGroupParticipantRequestRequest**](RejectGroupParticipantRequestRequest.md) |  | 

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

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetGroupAnnounce

> GenericResponse SetGroupAnnounce(ctx).SetGroupAnnounceRequest(setGroupAnnounceRequest).Execute()

Set group announce mode



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
	setGroupAnnounceRequest := *openapiclient.NewSetGroupAnnounceRequest("120363024512399999@g.us", true) // SetGroupAnnounceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.SetGroupAnnounce(context.Background()).SetGroupAnnounceRequest(setGroupAnnounceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.SetGroupAnnounce``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetGroupAnnounce`: GenericResponse
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.SetGroupAnnounce`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetGroupAnnounceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setGroupAnnounceRequest** | [**SetGroupAnnounceRequest**](SetGroupAnnounceRequest.md) |  | 

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


## SetGroupLocked

> GenericResponse SetGroupLocked(ctx).SetGroupLockedRequest(setGroupLockedRequest).Execute()

Set group locked status



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
	setGroupLockedRequest := *openapiclient.NewSetGroupLockedRequest("120363024512399999@g.us", true) // SetGroupLockedRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.SetGroupLocked(context.Background()).SetGroupLockedRequest(setGroupLockedRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.SetGroupLocked``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetGroupLocked`: GenericResponse
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.SetGroupLocked`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetGroupLockedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setGroupLockedRequest** | [**SetGroupLockedRequest**](SetGroupLockedRequest.md) |  | 

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


## SetGroupName

> GenericResponse SetGroupName(ctx).SetGroupNameRequest(setGroupNameRequest).Execute()

Set group name

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
	setGroupNameRequest := *openapiclient.NewSetGroupNameRequest("120363024512399999@g.us", "New Group Name") // SetGroupNameRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.SetGroupName(context.Background()).SetGroupNameRequest(setGroupNameRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.SetGroupName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetGroupName`: GenericResponse
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.SetGroupName`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetGroupNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setGroupNameRequest** | [**SetGroupNameRequest**](SetGroupNameRequest.md) |  | 

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


## SetGroupPhoto

> SetGroupPhotoResponse SetGroupPhoto(ctx).GroupId(groupId).Photo(photo).Execute()

Set group photo

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
	groupId := "groupId_example" // string | The group ID
	photo := os.NewFile(1234, "some_file") // *os.File | Group photo to upload (JPEG format recommended). Leave empty to remove photo. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.SetGroupPhoto(context.Background()).GroupId(groupId).Photo(photo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.SetGroupPhoto``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetGroupPhoto`: SetGroupPhotoResponse
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.SetGroupPhoto`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetGroupPhotoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **string** | The group ID | 
 **photo** | ***os.File** | Group photo to upload (JPEG format recommended). Leave empty to remove photo. | 

### Return type

[**SetGroupPhotoResponse**](SetGroupPhotoResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetGroupTopic

> GenericResponse SetGroupTopic(ctx).SetGroupTopicRequest(setGroupTopicRequest).Execute()

Set group topic



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
	setGroupTopicRequest := *openapiclient.NewSetGroupTopicRequest("120363024512399999@g.us") // SetGroupTopicRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.SetGroupTopic(context.Background()).SetGroupTopicRequest(setGroupTopicRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.SetGroupTopic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetGroupTopic`: GenericResponse
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.SetGroupTopic`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetGroupTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setGroupTopicRequest** | [**SetGroupTopicRequest**](SetGroupTopicRequest.md) |  | 

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

