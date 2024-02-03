# Go API client for sdk_go_whatsapp_web_multidevice

This API is used for sending whatsapp via API

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 3.7.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import sdk_go_whatsapp_web_multidevice "github.com/aldinokemal/sdk-go-whatsapp-web-multidevice"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sdk_go_whatsapp_web_multidevice.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), sdk_go_whatsapp_web_multidevice.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sdk_go_whatsapp_web_multidevice.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), sdk_go_whatsapp_web_multidevice.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sdk_go_whatsapp_web_multidevice.ContextOperationServerIndices` and `sdk_go_whatsapp_web_multidevice.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), sdk_go_whatsapp_web_multidevice.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sdk_go_whatsapp_web_multidevice.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:3000*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AppAPI* | [**AppDevices**](docs/AppAPI.md#appdevices) | **Get** /app/devices | Get list connected devices
*AppAPI* | [**AppLogin**](docs/AppAPI.md#applogin) | **Get** /app/login | Login to whatsapp server
*AppAPI* | [**AppLogout**](docs/AppAPI.md#applogout) | **Get** /app/logout | Remove database and logout
*AppAPI* | [**AppReconnect**](docs/AppAPI.md#appreconnect) | **Get** /app/reconnect | Reconnecting to whatsapp server
*GroupAPI* | [**JoinGroupWithLink**](docs/GroupAPI.md#joingroupwithlink) | **Post** /group/join-with-link | Join group with link
*GroupAPI* | [**LeaveGroup**](docs/GroupAPI.md#leavegroup) | **Post** /group/leave | Leave group
*MessageAPI* | [**ReactMessage**](docs/MessageAPI.md#reactmessage) | **Post** /message/{message_id}/reaction | Send reaction to message
*MessageAPI* | [**RevokeMessage**](docs/MessageAPI.md#revokemessage) | **Post** /message/{message_id}/revoke | Revoke Message
*SendAPI* | [**SendAudio**](docs/SendAPI.md#sendaudio) | **Post** /send/audio | Send Audio
*SendAPI* | [**SendContact**](docs/SendAPI.md#sendcontact) | **Post** /send/contact | Send Contact
*SendAPI* | [**SendFile**](docs/SendAPI.md#sendfile) | **Post** /send/file | Send File
*SendAPI* | [**SendImage**](docs/SendAPI.md#sendimage) | **Post** /send/image | Send Image
*SendAPI* | [**SendLink**](docs/SendAPI.md#sendlink) | **Post** /send/link | Send Link
*SendAPI* | [**SendLocation**](docs/SendAPI.md#sendlocation) | **Post** /send/location | Send Location
*SendAPI* | [**SendMessage**](docs/SendAPI.md#sendmessage) | **Post** /send/message | Send Message
*SendAPI* | [**SendVideo**](docs/SendAPI.md#sendvideo) | **Post** /send/video | Send Video
*UserAPI* | [**UserAvatar**](docs/UserAPI.md#useravatar) | **Get** /user/avatar | User Avatar
*UserAPI* | [**UserInfo**](docs/UserAPI.md#userinfo) | **Get** /user/info | User Info
*UserAPI* | [**UserMyGroups**](docs/UserAPI.md#usermygroups) | **Get** /user/my/groups | User My List Groups
*UserAPI* | [**UserMyPrivacy**](docs/UserAPI.md#usermyprivacy) | **Get** /user/my/privacy | User My Privacy Setting


## Documentation For Models

 - [DeviceResponse](docs/DeviceResponse.md)
 - [DeviceResponseResultsInner](docs/DeviceResponseResultsInner.md)
 - [ErrorBadRequest](docs/ErrorBadRequest.md)
 - [ErrorInternalServer](docs/ErrorInternalServer.md)
 - [GenericResponse](docs/GenericResponse.md)
 - [LoginResponse](docs/LoginResponse.md)
 - [LoginResponseResults](docs/LoginResponseResults.md)
 - [SendResponse](docs/SendResponse.md)
 - [SendResponseResults](docs/SendResponseResults.md)
 - [UserAvatarResponse](docs/UserAvatarResponse.md)
 - [UserAvatarResponseResults](docs/UserAvatarResponseResults.md)
 - [UserGroupResponse](docs/UserGroupResponse.md)
 - [UserGroupResponseResults](docs/UserGroupResponseResults.md)
 - [UserGroupResponseResultsDataInner](docs/UserGroupResponseResultsDataInner.md)
 - [UserGroupResponseResultsDataInnerParticipantsInner](docs/UserGroupResponseResultsDataInnerParticipantsInner.md)
 - [UserInfoResponse](docs/UserInfoResponse.md)
 - [UserInfoResponseResults](docs/UserInfoResponseResults.md)
 - [UserInfoResponseResultsDevicesInner](docs/UserInfoResponseResultsDevicesInner.md)
 - [UserPrivacyResponse](docs/UserPrivacyResponse.md)
 - [UserPrivacyResponseResults](docs/UserPrivacyResponseResults.md)


## Documentation For Authorization

Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



