# \IncomingWebhooksApi

All URIs are relative to *https://api.firstquadrant.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IncomingWebhooksMarkConversationsAsWonPost**](IncomingWebhooksApi.md#IncomingWebhooksMarkConversationsAsWonPost) | **Post** /incoming-webhooks/mark-conversations-as-won | Mark conversations as won



## IncomingWebhooksMarkConversationsAsWonPost

> Success IncomingWebhooksMarkConversationsAsWonPost(ctx).MarkConversationAsWonBody(markConversationAsWonBody).Execute()

Mark conversations as won



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    markConversationAsWonBody := *openapiclient.NewMarkConversationAsWonBody(interface{}(user@example.com), interface{}(aac843f5-69ab-4f88-9afb-0ed33a383ee4)) // MarkConversationAsWonBody | Find conversations using the given email address and create unique goal events based on the idempotency key.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IncomingWebhooksApi.IncomingWebhooksMarkConversationsAsWonPost(context.Background()).MarkConversationAsWonBody(markConversationAsWonBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncomingWebhooksApi.IncomingWebhooksMarkConversationsAsWonPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IncomingWebhooksMarkConversationsAsWonPost`: Success
    fmt.Fprintf(os.Stdout, "Response from `IncomingWebhooksApi.IncomingWebhooksMarkConversationsAsWonPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIncomingWebhooksMarkConversationsAsWonPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **markConversationAsWonBody** | [**MarkConversationAsWonBody**](MarkConversationAsWonBody.md) | Find conversations using the given email address and create unique goal events based on the idempotency key. | 

### Return type

[**Success**](Success.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

