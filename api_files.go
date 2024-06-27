/*
OpenAI API

The OpenAI REST API. Please see https://platform.openai.com/docs/api-reference for more details.

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openai

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"os"
)


// FilesAPIService FilesAPI service
type FilesAPIService service

type ApiCreateFileRequest struct {
	ctx context.Context
	ApiService *FilesAPIService
	file *os.File
	purpose *string
}

// The File object (not file name) to be uploaded. 
func (r ApiCreateFileRequest) File(file *os.File) ApiCreateFileRequest {
	r.file = file
	return r
}

// The intended purpose of the uploaded file.  Use \\\&quot;assistants\\\&quot; for [Assistants](/docs/api-reference/assistants) and [Message](/docs/api-reference/messages) files, \\\&quot;vision\\\&quot; for Assistants image file inputs, \\\&quot;batch\\\&quot; for [Batch API](/docs/guides/batch), and \\\&quot;fine-tune\\\&quot; for [Fine-tuning](/docs/api-reference/fine-tuning). 
func (r ApiCreateFileRequest) Purpose(purpose string) ApiCreateFileRequest {
	r.purpose = &purpose
	return r
}

func (r ApiCreateFileRequest) Execute() (*OpenAIFile, *http.Response, error) {
	return r.ApiService.CreateFileExecute(r)
}

/*
CreateFile Upload a file that can be used across various endpoints. Individual files can be up to 512 MB, and the size of all files uploaded by one organization can be up to 100 GB.  The Assistants API supports files up to 2 million tokens and of specific file types. See the [Assistants Tools guide](/docs/assistants/tools) for details.  The Fine-tuning API only supports `.jsonl` files. The input also has certain required formats for fine-tuning [chat](/docs/api-reference/fine-tuning/chat-input) or [completions](/docs/api-reference/fine-tuning/completions-input) models.  The Batch API only supports `.jsonl` files up to 100 MB in size. The input also has a specific required [format](/docs/api-reference/batch/request-input).  Please [contact us](https://help.openai.com/) if you need to increase these storage limits. 

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateFileRequest
*/
func (a *FilesAPIService) CreateFile(ctx context.Context) ApiCreateFileRequest {
	return ApiCreateFileRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return OpenAIFile
func (a *FilesAPIService) CreateFileExecute(r ApiCreateFileRequest) (*OpenAIFile, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OpenAIFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesAPIService.CreateFile")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/files"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.file == nil {
		return localVarReturnValue, nil, reportError("file is required and must be specified")
	}
	if r.purpose == nil {
		return localVarReturnValue, nil, reportError("purpose is required and must be specified")
	}

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
	parameterAddToHeaderOrQuery(localVarFormParams, "purpose", r.purpose, "")
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

type ApiDeleteFileRequest struct {
	ctx context.Context
	ApiService *FilesAPIService
	fileId string
}

func (r ApiDeleteFileRequest) Execute() (*DeleteFileResponse, *http.Response, error) {
	return r.ApiService.DeleteFileExecute(r)
}

/*
DeleteFile Delete a file.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fileId The ID of the file to use for this request.
 @return ApiDeleteFileRequest
*/
func (a *FilesAPIService) DeleteFile(ctx context.Context, fileId string) ApiDeleteFileRequest {
	return ApiDeleteFileRequest{
		ApiService: a,
		ctx: ctx,
		fileId: fileId,
	}
}

// Execute executes the request
//  @return DeleteFileResponse
func (a *FilesAPIService) DeleteFileExecute(r ApiDeleteFileRequest) (*DeleteFileResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DeleteFileResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesAPIService.DeleteFile")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/files/{file_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"file_id"+"}", url.PathEscape(parameterValueToString(r.fileId, "fileId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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

type ApiDownloadFileRequest struct {
	ctx context.Context
	ApiService *FilesAPIService
	fileId string
}

func (r ApiDownloadFileRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.DownloadFileExecute(r)
}

/*
DownloadFile Returns the contents of the specified file.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fileId The ID of the file to use for this request.
 @return ApiDownloadFileRequest
*/
func (a *FilesAPIService) DownloadFile(ctx context.Context, fileId string) ApiDownloadFileRequest {
	return ApiDownloadFileRequest{
		ApiService: a,
		ctx: ctx,
		fileId: fileId,
	}
}

// Execute executes the request
//  @return string
func (a *FilesAPIService) DownloadFileExecute(r ApiDownloadFileRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesAPIService.DownloadFile")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/files/{file_id}/content"
	localVarPath = strings.Replace(localVarPath, "{"+"file_id"+"}", url.PathEscape(parameterValueToString(r.fileId, "fileId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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

type ApiListFilesRequest struct {
	ctx context.Context
	ApiService *FilesAPIService
	purpose *string
}

// Only return files with the given purpose.
func (r ApiListFilesRequest) Purpose(purpose string) ApiListFilesRequest {
	r.purpose = &purpose
	return r
}

func (r ApiListFilesRequest) Execute() (*ListFilesResponse, *http.Response, error) {
	return r.ApiService.ListFilesExecute(r)
}

/*
ListFiles Returns a list of files that belong to the user's organization.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListFilesRequest
*/
func (a *FilesAPIService) ListFiles(ctx context.Context) ApiListFilesRequest {
	return ApiListFilesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListFilesResponse
func (a *FilesAPIService) ListFilesExecute(r ApiListFilesRequest) (*ListFilesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListFilesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesAPIService.ListFiles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/files"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.purpose != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "purpose", r.purpose, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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

type ApiRetrieveFileRequest struct {
	ctx context.Context
	ApiService *FilesAPIService
	fileId string
}

func (r ApiRetrieveFileRequest) Execute() (*OpenAIFile, *http.Response, error) {
	return r.ApiService.RetrieveFileExecute(r)
}

/*
RetrieveFile Returns information about a specific file.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fileId The ID of the file to use for this request.
 @return ApiRetrieveFileRequest
*/
func (a *FilesAPIService) RetrieveFile(ctx context.Context, fileId string) ApiRetrieveFileRequest {
	return ApiRetrieveFileRequest{
		ApiService: a,
		ctx: ctx,
		fileId: fileId,
	}
}

// Execute executes the request
//  @return OpenAIFile
func (a *FilesAPIService) RetrieveFileExecute(r ApiRetrieveFileRequest) (*OpenAIFile, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OpenAIFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesAPIService.RetrieveFile")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/files/{file_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"file_id"+"}", url.PathEscape(parameterValueToString(r.fileId, "fileId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
