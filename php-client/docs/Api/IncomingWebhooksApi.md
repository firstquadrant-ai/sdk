# OpenAPI\Client\IncomingWebhooksApi

All URIs are relative to https://api.firstquadrant.ai, except if the operation defines another base path.

| Method | HTTP request | Description |
| ------------- | ------------- | ------------- |
| [**incomingWebhooksMarkConversationsAsWonPost()**](IncomingWebhooksApi.md#incomingWebhooksMarkConversationsAsWonPost) | **POST** /incoming-webhooks/mark-conversations-as-won | Mark conversations as won |


## `incomingWebhooksMarkConversationsAsWonPost()`

```php
incomingWebhooksMarkConversationsAsWonPost($mark_conversation_as_won_body): \OpenAPI\Client\Model\Success
```

Mark conversations as won

Mark any open conversations with a particular lead as \"Won\" as a custom goal.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: api_key
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\IncomingWebhooksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$mark_conversation_as_won_body = new \OpenAPI\Client\Model\MarkConversationAsWonBody(); // \OpenAPI\Client\Model\MarkConversationAsWonBody | Find conversations using the given email address and create unique goal events based on the idempotency key.

try {
    $result = $apiInstance->incomingWebhooksMarkConversationsAsWonPost($mark_conversation_as_won_body);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling IncomingWebhooksApi->incomingWebhooksMarkConversationsAsWonPost: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **mark_conversation_as_won_body** | [**\OpenAPI\Client\Model\MarkConversationAsWonBody**](../Model/MarkConversationAsWonBody.md)| Find conversations using the given email address and create unique goal events based on the idempotency key. | |

### Return type

[**\OpenAPI\Client\Model\Success**](../Model/Success.md)

### Authorization

[api_key](../../README.md#api_key)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)
