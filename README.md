# Go API client for openapi

This API is used for sending whatsapp via API

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 3.3.1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/aldinokemal/sdk-go-whatsapp-web-multidevice"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:3000*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AppApi* | [**AppLogin**](docs/AppApi.md#applogin) | **Get** /app/login | Login to whatsapp server
*AppApi* | [**AppLogout**](docs/AppApi.md#applogout) | **Get** /app/logout | Remove database and logout
*AppApi* | [**AppReconnect**](docs/AppApi.md#appreconnect) | **Get** /app/reconnect | Reconnecting to whatsapp server
*MessageApi* | [**RevokeMessage**](docs/MessageApi.md#revokemessage) | **Post** /message/:message_id/revoke | Send Link
*MessageApi* | [**SendContact**](docs/MessageApi.md#sendcontact) | **Post** /send/contact | Send Contact
*MessageApi* | [**SendFile**](docs/MessageApi.md#sendfile) | **Post** /send/file | Send File
*MessageApi* | [**SendImage**](docs/MessageApi.md#sendimage) | **Post** /send/image | Send Image
*MessageApi* | [**SendLink**](docs/MessageApi.md#sendlink) | **Post** /send/link | Send Link
*MessageApi* | [**SendLocation**](docs/MessageApi.md#sendlocation) | **Post** /send/location | Send Location
*MessageApi* | [**SendMessage**](docs/MessageApi.md#sendmessage) | **Post** /send/message | Send Message
*MessageApi* | [**SendVideo**](docs/MessageApi.md#sendvideo) | **Post** /send/video | Send Video
*UserApi* | [**UserAvatar**](docs/UserApi.md#useravatar) | **Get** /user/avatar | User Avatar
*UserApi* | [**UserInfo**](docs/UserApi.md#userinfo) | **Get** /user/info | User Info
*UserApi* | [**UserMyGroups**](docs/UserApi.md#usermygroups) | **Get** /user/my/groups | User My List Groups
*UserApi* | [**UserMyPrivacy**](docs/UserApi.md#usermyprivacy) | **Get** /user/my/privacy | User My Privacy Setting


## Documentation For Models

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



