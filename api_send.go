/*
WhatsApp API MultiDevice

This API is used for sending whatsapp via API

API version: 3.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk_go_whatsapp_web_multidevice

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"os"
)


// SendAPIService SendAPI service
type SendAPIService service

type ApiSendAudioRequest struct {
	ctx context.Context
	ApiService *SendAPIService
	phone *string
	audio *os.File
}

// Phone number with country code
func (r ApiSendAudioRequest) Phone(phone string) ApiSendAudioRequest {
	r.phone = &phone
	return r
}

// Audio to send
func (r ApiSendAudioRequest) Audio(audio *os.File) ApiSendAudioRequest {
	r.audio = audio
	return r
}

func (r ApiSendAudioRequest) Execute() (*SendResponse, *http.Response, error) {
	return r.ApiService.SendAudioExecute(r)
}

/*
SendAudio Send Audio

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSendAudioRequest
*/
func (a *SendAPIService) SendAudio(ctx context.Context) ApiSendAudioRequest {
	return ApiSendAudioRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SendResponse
func (a *SendAPIService) SendAudioExecute(r ApiSendAudioRequest) (*SendResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SendResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SendAPIService.SendAudio")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/send/audio"

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
		parameterAddToHeaderOrQuery(localVarFormParams, "phone", r.phone, "")
	}
	var audioLocalVarFormFileName string
	var audioLocalVarFileName     string
	var audioLocalVarFileBytes    []byte

	audioLocalVarFormFileName = "audio"
	audioLocalVarFile := r.audio

	if audioLocalVarFile != nil {
		fbs, _ := io.ReadAll(audioLocalVarFile)

		audioLocalVarFileBytes = fbs
		audioLocalVarFileName = audioLocalVarFile.Name()
		audioLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: audioLocalVarFileBytes, fileName: audioLocalVarFileName, formFileName: audioLocalVarFormFileName})
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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
	ApiService *SendAPIService
	phone *string
	contactName *string
	contactPhone *string
}

// Phone number with country code
func (r ApiSendContactRequest) Phone(phone string) ApiSendContactRequest {
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
func (a *SendAPIService) SendContact(ctx context.Context) ApiSendContactRequest {
	return ApiSendContactRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SendResponse
func (a *SendAPIService) SendContactExecute(r ApiSendContactRequest) (*SendResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SendResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SendAPIService.SendContact")
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
		parameterAddToHeaderOrQuery(localVarFormParams, "phone", r.phone, "")
	}
	if r.contactName != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "contact_name", r.contactName, "")
	}
	if r.contactPhone != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "contact_phone", r.contactPhone, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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
	ApiService *SendAPIService
	phone *string
	caption *string
	file *os.File
}

// Phone number with country code
func (r ApiSendFileRequest) Phone(phone string) ApiSendFileRequest {
	r.phone = &phone
	return r
}

// Caption to send
func (r ApiSendFileRequest) Caption(caption string) ApiSendFileRequest {
	r.caption = &caption
	return r
}

// File to send
func (r ApiSendFileRequest) File(file *os.File) ApiSendFileRequest {
	r.file = file
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
func (a *SendAPIService) SendFile(ctx context.Context) ApiSendFileRequest {
	return ApiSendFileRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SendResponse
func (a *SendAPIService) SendFileExecute(r ApiSendFileRequest) (*SendResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SendResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SendAPIService.SendFile")
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
		parameterAddToHeaderOrQuery(localVarFormParams, "phone", r.phone, "")
	}
	if r.caption != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "caption", r.caption, "")
	}
	var fileLocalVarFormFileName string
	var fileLocalVarFileName     string
	var fileLocalVarFileBytes    []byte

	fileLocalVarFormFileName = "file"
	fileLocalVarFile := r.file

	if fileLocalVarFile != nil {
		fbs, _ := io.ReadAll(fileLocalVarFile)

		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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
	ApiService *SendAPIService
	phone *string
	caption *string
	viewOnce *bool
	image *os.File
	compress *bool
}

// Phone number with country code
func (r ApiSendImageRequest) Phone(phone string) ApiSendImageRequest {
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
func (r ApiSendImageRequest) Image(image *os.File) ApiSendImageRequest {
	r.image = image
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
func (a *SendAPIService) SendImage(ctx context.Context) ApiSendImageRequest {
	return ApiSendImageRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SendResponse
func (a *SendAPIService) SendImageExecute(r ApiSendImageRequest) (*SendResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SendResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SendAPIService.SendImage")
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
		parameterAddToHeaderOrQuery(localVarFormParams, "phone", r.phone, "")
	}
	if r.caption != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "caption", r.caption, "")
	}
	if r.viewOnce != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "view_once", r.viewOnce, "")
	}
	var imageLocalVarFormFileName string
	var imageLocalVarFileName     string
	var imageLocalVarFileBytes    []byte

	imageLocalVarFormFileName = "image"
	imageLocalVarFile := r.image

	if imageLocalVarFile != nil {
		fbs, _ := io.ReadAll(imageLocalVarFile)

		imageLocalVarFileBytes = fbs
		imageLocalVarFileName = imageLocalVarFile.Name()
		imageLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: imageLocalVarFileBytes, fileName: imageLocalVarFileName, formFileName: imageLocalVarFormFileName})
	}
	if r.compress != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "compress", r.compress, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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
	ApiService *SendAPIService
	phone *string
	link *string
	caption *string
}

// Phone number with country code
func (r ApiSendLinkRequest) Phone(phone string) ApiSendLinkRequest {
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
func (a *SendAPIService) SendLink(ctx context.Context) ApiSendLinkRequest {
	return ApiSendLinkRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SendResponse
func (a *SendAPIService) SendLinkExecute(r ApiSendLinkRequest) (*SendResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SendResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SendAPIService.SendLink")
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
		parameterAddToHeaderOrQuery(localVarFormParams, "phone", r.phone, "")
	}
	if r.link != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "link", r.link, "")
	}
	if r.caption != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "caption", r.caption, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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
	ApiService *SendAPIService
	phone *string
	latitude *string
	longitude *string
}

// Phone number with country code
func (r ApiSendLocationRequest) Phone(phone string) ApiSendLocationRequest {
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
func (a *SendAPIService) SendLocation(ctx context.Context) ApiSendLocationRequest {
	return ApiSendLocationRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SendResponse
func (a *SendAPIService) SendLocationExecute(r ApiSendLocationRequest) (*SendResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SendResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SendAPIService.SendLocation")
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
		parameterAddToHeaderOrQuery(localVarFormParams, "phone", r.phone, "")
	}
	if r.latitude != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "latitude", r.latitude, "")
	}
	if r.longitude != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "longitude", r.longitude, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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
	ApiService *SendAPIService
	phone *string
	message *string
	replyMessageId *string
}

// Phone number with country code
func (r ApiSendMessageRequest) Phone(phone string) ApiSendMessageRequest {
	r.phone = &phone
	return r
}

// Message to send
func (r ApiSendMessageRequest) Message(message string) ApiSendMessageRequest {
	r.message = &message
	return r
}

// Message ID that you want reply
func (r ApiSendMessageRequest) ReplyMessageId(replyMessageId string) ApiSendMessageRequest {
	r.replyMessageId = &replyMessageId
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
func (a *SendAPIService) SendMessage(ctx context.Context) ApiSendMessageRequest {
	return ApiSendMessageRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SendResponse
func (a *SendAPIService) SendMessageExecute(r ApiSendMessageRequest) (*SendResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SendResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SendAPIService.SendMessage")
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
		parameterAddToHeaderOrQuery(localVarFormParams, "phone", r.phone, "")
	}
	if r.message != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "message", r.message, "")
	}
	if r.replyMessageId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "reply_message_id", r.replyMessageId, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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
	ApiService *SendAPIService
	phone *string
	caption *string
	viewOnce *bool
	video *os.File
	compress *bool
}

// Phone number with country code
func (r ApiSendVideoRequest) Phone(phone string) ApiSendVideoRequest {
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
func (r ApiSendVideoRequest) Video(video *os.File) ApiSendVideoRequest {
	r.video = video
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
func (a *SendAPIService) SendVideo(ctx context.Context) ApiSendVideoRequest {
	return ApiSendVideoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SendResponse
func (a *SendAPIService) SendVideoExecute(r ApiSendVideoRequest) (*SendResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SendResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SendAPIService.SendVideo")
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
		parameterAddToHeaderOrQuery(localVarFormParams, "phone", r.phone, "")
	}
	if r.caption != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "caption", r.caption, "")
	}
	if r.viewOnce != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "view_once", r.viewOnce, "")
	}
	var videoLocalVarFormFileName string
	var videoLocalVarFileName     string
	var videoLocalVarFileBytes    []byte

	videoLocalVarFormFileName = "video"
	videoLocalVarFile := r.video

	if videoLocalVarFile != nil {
		fbs, _ := io.ReadAll(videoLocalVarFile)

		videoLocalVarFileBytes = fbs
		videoLocalVarFileName = videoLocalVarFile.Name()
		videoLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: videoLocalVarFileBytes, fileName: videoLocalVarFileName, formFileName: videoLocalVarFormFileName})
	}
	if r.compress != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "compress", r.compress, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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
