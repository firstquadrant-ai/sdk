/*
FirstQuadrant API

The FirstQuadrant API is used to interact with FirstQuadrant programmatically. We also have SDKs available (coming soon).

API version: 0.0.0
Contact: inquiry@firstquadrant.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// IncomingWebhooksApiService IncomingWebhooksApi service
type IncomingWebhooksApiService service

type ApiIncomingWebhooksMarkConversationsAsWonPostRequest struct {
	ctx context.Context
	ApiService *IncomingWebhooksApiService
	markConversationAsWonBody *MarkConversationAsWonBody
}

// Find conversations using the given email address and create unique goal events based on the idempotency key.
func (r ApiIncomingWebhooksMarkConversationsAsWonPostRequest) MarkConversationAsWonBody(markConversationAsWonBody MarkConversationAsWonBody) ApiIncomingWebhooksMarkConversationsAsWonPostRequest {
	r.markConversationAsWonBody = &markConversationAsWonBody
	return r
}

func (r ApiIncomingWebhooksMarkConversationsAsWonPostRequest) Execute() (*Success, *http.Response, error) {
	return r.ApiService.IncomingWebhooksMarkConversationsAsWonPostExecute(r)
}

/*
IncomingWebhooksMarkConversationsAsWonPost Mark conversations as won

Mark any open conversations with a particular lead as "Won" as a custom goal.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiIncomingWebhooksMarkConversationsAsWonPostRequest
*/
func (a *IncomingWebhooksApiService) IncomingWebhooksMarkConversationsAsWonPost(ctx context.Context) ApiIncomingWebhooksMarkConversationsAsWonPostRequest {
	return ApiIncomingWebhooksMarkConversationsAsWonPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Success
func (a *IncomingWebhooksApiService) IncomingWebhooksMarkConversationsAsWonPostExecute(r ApiIncomingWebhooksMarkConversationsAsWonPostRequest) (*Success, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Success
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IncomingWebhooksApiService.IncomingWebhooksMarkConversationsAsWonPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/incoming-webhooks/mark-conversations-as-won"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.markConversationAsWonBody == nil {
		return localVarReturnValue, nil, reportError("markConversationAsWonBody is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.markConversationAsWonBody
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
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
