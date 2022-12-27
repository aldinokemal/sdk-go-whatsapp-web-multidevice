/*
WhatsApp API MultiDevice

This API is used for sending whatsapp via API

API version: 3.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)


// MessageApiService MessageApi service
type MessageApiService service

type ApiRevokeMessageRequest struct {
	ctx context.Context
	ApiService *MessageApiService
	phone *int32
	messageId *string
}

// Phone number with country code
func (r ApiRevokeMessageRequest) Phone(phone int32) ApiRevokeMessageRequest {
	r.phone = &phone
	return r
}

// Message ID
func (r ApiRevokeMessageRequest) MessageId(messageId string) ApiRevokeMessageRequest {
	r.messageId = &messageId
	return r
}

func (r ApiRevokeMessageRequest) Execute() (*SendResponse, *http.Response, error) {
	return r.ApiService.RevokeMessageExecute(r)
}

/*
RevokeMessage Send Link

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRevokeMessageRequest
*/
func (a *MessageApiService) RevokeMessage(ctx context.Context) ApiRevokeMessageRequest {
	return ApiRevokeMessageRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SendResponse
func (a *MessageApiService) RevokeMessageExecute(r ApiRevokeMessageRequest) (*SendResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SendResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MessageApiService.RevokeMessage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/message/:message_id/revoke"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.phone != nil {
		parameterAddToQuery(localVarFormParams, "phone", r.phone, "")
	}
	if r.messageId != nil {
		parameterAddToQuery(localVarFormParams, "message_id", r.messageId, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorBadRequest
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorInternalServer
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSendContactRequest struct {
	ctx context.Context
	ApiService *MessageApiService
	phone *int32
	contactName *string
	contactPhone *string
}

// Phone number with country code
func (r ApiSendContactRequest) Phone(phone int32) ApiSendContactRequest {
	r.phone = &phone
	return r
}

// Contact name
func (r ApiSendContactRequest) ContactName(contactName string) ApiSendContactRequest {
	r.contactName = &contactName
	return r
}

// Contact phone number
func (r ApiSendContactRequest) ContactPhone(contactPhone string) ApiSendContactRequest {
	r.contactPhone = &contactPhone
	return r
}

func (r ApiSendContactRequest) Execute() (*SendResponse, *http.Response, error) {
	return r.ApiService.SendContactExecute(r)
}

/*
SendContact Send Contact

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSendContactRequest
*/
func (a *MessageApiService) SendContact(ctx context.Context) ApiSendContactRequest {
	return ApiSendContactRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SendResponse
func (a *MessageApiService) SendContactExecute(r ApiSendContactRequest) (*SendResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SendResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MessageApiService.SendContact")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/send/contact"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.phone != nil {
		parameterAddToQuery(localVarFormParams, "phone", r.phone, "")
	}
	if r.contactName != nil {
		parameterAddToQuery(localVarFormParams, "contact_name", r.contactName, "")
	}
	if r.contactPhone != nil {
		parameterAddToQuery(localVarFormParams, "contact_phone", r.contactPhone, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorBadRequest
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorInternalServer
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSendFileRequest struct {
	ctx context.Context
	ApiService *MessageApiService
	phone *int32
	caption *string
	file *os.File
}

// Phone number with country code
func (r ApiSendFileRequest) Phone(phone int32) ApiSendFileRequest {
	r.phone = &phone
	return r
}

// Caption to send
func (r ApiSendFileRequest) Caption(caption string) ApiSendFileRequest {
	r.caption = &caption
	return r
}

// File to send
func (r ApiSendFileRequest) File(file os.File) ApiSendFileRequest {
	r.file = &file
	return r
}

func (r ApiSendFileRequest) Execute() (*SendResponse, *http.Response, error) {
	return r.ApiService.SendFileExecute(r)
}

/*
SendFile Send File

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSendFileRequest
*/
func (a *MessageApiService) SendFile(ctx context.Context) ApiSendFileRequest {
	return ApiSendFileRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SendResponse
func (a *MessageApiService) SendFileExecute(r ApiSendFileRequest) (*SendResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SendResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MessageApiService.SendFile")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/send/file"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.phone != nil {
		parameterAddToQuery(localVarFormParams, "phone", r.phone, "")
	}
	if r.caption != nil {
		parameterAddToQuery(localVarFormParams, "caption", r.caption, "")
	}
	var fileLocalVarFormFileName string
	var fileLocalVarFileName     string
	var fileLocalVarFileBytes    []byte

	fileLocalVarFormFileName = "file"

	var fileLocalVarFile *os.File
	if r.file != nil {
		fileLocalVarFile = r.file
	}
	if fileLocalVarFile != nil {
		fbs, _ := ioutil.ReadAll(fileLocalVarFile)
		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
	}
	formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorBadRequest
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorInternalServer
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSendImageRequest struct {
	ctx context.Context
	ApiService *MessageApiService
	phone *int32
	caption *string
	viewOnce *bool
	image *os.File
	compress *bool
}

// Phone number with country code
func (r ApiSendImageRequest) Phone(phone int32) ApiSendImageRequest {
	r.phone = &phone
	return r
}

// Caption to send
func (r ApiSendImageRequest) Caption(caption string) ApiSendImageRequest {
	r.caption = &caption
	return r
}

// View once
func (r ApiSendImageRequest) ViewOnce(viewOnce bool) ApiSendImageRequest {
	r.viewOnce = &viewOnce
	return r
}

// Image to send
func (r ApiSendImageRequest) Image(image os.File) ApiSendImageRequest {
	r.image = &image
	return r
}

// Compress image
func (r ApiSendImageRequest) Compress(compress bool) ApiSendImageRequest {
	r.compress = &compress
	return r
}

func (r ApiSendImageRequest) Execute() (*SendResponse, *http.Response, error) {
	return r.ApiService.SendImageExecute(r)
}

/*
SendImage Send Image

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSendImageRequest
*/
func (a *MessageApiService) SendImage(ctx context.Context) ApiSendImageRequest {
	return ApiSendImageRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SendResponse
func (a *MessageApiService) SendImageExecute(r ApiSendImageRequest) (*SendResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SendResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MessageApiService.SendImage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/send/image"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.phone != nil {
		parameterAddToQuery(localVarFormParams, "phone", r.phone, "")
	}
	if r.caption != nil {
		parameterAddToQuery(localVarFormParams, "caption", r.caption, "")
	}
	if r.viewOnce != nil {
		parameterAddToQuery(localVarFormParams, "view_once", r.viewOnce, "")
	}
	var imageLocalVarFormFileName string
	var imageLocalVarFileName     string
	var imageLocalVarFileBytes    []byte

	imageLocalVarFormFileName = "image"

	var imageLocalVarFile *os.File
	if r.image != nil {
		imageLocalVarFile = r.image
	}
	if imageLocalVarFile != nil {
		fbs, _ := ioutil.ReadAll(imageLocalVarFile)
		imageLocalVarFileBytes = fbs
		imageLocalVarFileName = imageLocalVarFile.Name()
		imageLocalVarFile.Close()
	}
	formFiles = append(formFiles, formFile{fileBytes: imageLocalVarFileBytes, fileName: imageLocalVarFileName, formFileName: imageLocalVarFormFileName})
	if r.compress != nil {
		parameterAddToQuery(localVarFormParams, "compress", r.compress, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorBadRequest
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorInternalServer
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSendLinkRequest struct {
	ctx context.Context
	ApiService *MessageApiService
	phone *int32
	link *string
	caption *string
}

// Phone number with country code
func (r ApiSendLinkRequest) Phone(phone int32) ApiSendLinkRequest {
	r.phone = &phone
	return r
}

// Link to send
func (r ApiSendLinkRequest) Link(link string) ApiSendLinkRequest {
	r.link = &link
	return r
}

// Caption to send
func (r ApiSendLinkRequest) Caption(caption string) ApiSendLinkRequest {
	r.caption = &caption
	return r
}

func (r ApiSendLinkRequest) Execute() (*SendResponse, *http.Response, error) {
	return r.ApiService.SendLinkExecute(r)
}

/*
SendLink Send Link

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSendLinkRequest
*/
func (a *MessageApiService) SendLink(ctx context.Context) ApiSendLinkRequest {
	return ApiSendLinkRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SendResponse
func (a *MessageApiService) SendLinkExecute(r ApiSendLinkRequest) (*SendResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SendResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MessageApiService.SendLink")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/send/link"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.phone != nil {
		parameterAddToQuery(localVarFormParams, "phone", r.phone, "")
	}
	if r.link != nil {
		parameterAddToQuery(localVarFormParams, "link", r.link, "")
	}
	if r.caption != nil {
		parameterAddToQuery(localVarFormParams, "caption", r.caption, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorBadRequest
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorInternalServer
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSendLocationRequest struct {
	ctx context.Context
	ApiService *MessageApiService
	phone *int32
	latitude *string
	longitude *string
}

// Phone number with country code
func (r ApiSendLocationRequest) Phone(phone int32) ApiSendLocationRequest {
	r.phone = &phone
	return r
}

// Latitude coordinate
func (r ApiSendLocationRequest) Latitude(latitude string) ApiSendLocationRequest {
	r.latitude = &latitude
	return r
}

// Longitude coordinate
func (r ApiSendLocationRequest) Longitude(longitude string) ApiSendLocationRequest {
	r.longitude = &longitude
	return r
}

func (r ApiSendLocationRequest) Execute() (*SendResponse, *http.Response, error) {
	return r.ApiService.SendLocationExecute(r)
}

/*
SendLocation Send Location

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSendLocationRequest
*/
func (a *MessageApiService) SendLocation(ctx context.Context) ApiSendLocationRequest {
	return ApiSendLocationRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SendResponse
func (a *MessageApiService) SendLocationExecute(r ApiSendLocationRequest) (*SendResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SendResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MessageApiService.SendLocation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/send/location"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.phone != nil {
		parameterAddToQuery(localVarFormParams, "phone", r.phone, "")
	}
	if r.latitude != nil {
		parameterAddToQuery(localVarFormParams, "latitude", r.latitude, "")
	}
	if r.longitude != nil {
		parameterAddToQuery(localVarFormParams, "longitude", r.longitude, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorBadRequest
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorInternalServer
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSendMessageRequest struct {
	ctx context.Context
	ApiService *MessageApiService
	phone *int32
	message *string
}

// Phone number with country code
func (r ApiSendMessageRequest) Phone(phone int32) ApiSendMessageRequest {
	r.phone = &phone
	return r
}

// Message to send
func (r ApiSendMessageRequest) Message(message string) ApiSendMessageRequest {
	r.message = &message
	return r
}

func (r ApiSendMessageRequest) Execute() (*SendResponse, *http.Response, error) {
	return r.ApiService.SendMessageExecute(r)
}

/*
SendMessage Send Message

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSendMessageRequest
*/
func (a *MessageApiService) SendMessage(ctx context.Context) ApiSendMessageRequest {
	return ApiSendMessageRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SendResponse
func (a *MessageApiService) SendMessageExecute(r ApiSendMessageRequest) (*SendResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SendResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MessageApiService.SendMessage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/send/message"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.phone != nil {
		parameterAddToQuery(localVarFormParams, "phone", r.phone, "")
	}
	if r.message != nil {
		parameterAddToQuery(localVarFormParams, "message", r.message, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorBadRequest
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorInternalServer
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSendVideoRequest struct {
	ctx context.Context
	ApiService *MessageApiService
	phone *int32
	caption *string
	viewOnce *bool
	video *os.File
	compress *bool
}

// Phone number with country code
func (r ApiSendVideoRequest) Phone(phone int32) ApiSendVideoRequest {
	r.phone = &phone
	return r
}

// Caption to send
func (r ApiSendVideoRequest) Caption(caption string) ApiSendVideoRequest {
	r.caption = &caption
	return r
}

// View once
func (r ApiSendVideoRequest) ViewOnce(viewOnce bool) ApiSendVideoRequest {
	r.viewOnce = &viewOnce
	return r
}

// Video to send
func (r ApiSendVideoRequest) Video(video os.File) ApiSendVideoRequest {
	r.video = &video
	return r
}

// Compress video
func (r ApiSendVideoRequest) Compress(compress bool) ApiSendVideoRequest {
	r.compress = &compress
	return r
}

func (r ApiSendVideoRequest) Execute() (*SendResponse, *http.Response, error) {
	return r.ApiService.SendVideoExecute(r)
}

/*
SendVideo Send Video

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSendVideoRequest
*/
func (a *MessageApiService) SendVideo(ctx context.Context) ApiSendVideoRequest {
	return ApiSendVideoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SendResponse
func (a *MessageApiService) SendVideoExecute(r ApiSendVideoRequest) (*SendResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SendResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MessageApiService.SendVideo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/send/video"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.phone != nil {
		parameterAddToQuery(localVarFormParams, "phone", r.phone, "")
	}
	if r.caption != nil {
		parameterAddToQuery(localVarFormParams, "caption", r.caption, "")
	}
	if r.viewOnce != nil {
		parameterAddToQuery(localVarFormParams, "view_once", r.viewOnce, "")
	}
	var videoLocalVarFormFileName string
	var videoLocalVarFileName     string
	var videoLocalVarFileBytes    []byte

	videoLocalVarFormFileName = "video"

	var videoLocalVarFile *os.File
	if r.video != nil {
		videoLocalVarFile = r.video
	}
	if videoLocalVarFile != nil {
		fbs, _ := ioutil.ReadAll(videoLocalVarFile)
		videoLocalVarFileBytes = fbs
		videoLocalVarFileName = videoLocalVarFile.Name()
		videoLocalVarFile.Close()
	}
	formFiles = append(formFiles, formFile{fileBytes: videoLocalVarFileBytes, fileName: videoLocalVarFileName, formFileName: videoLocalVarFormFileName})
	if r.compress != nil {
		parameterAddToQuery(localVarFormParams, "compress", r.compress, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorBadRequest
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorInternalServer
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
