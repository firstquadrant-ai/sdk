<a name="__pageTop"></a>
# openapi_client.apis.tags.incoming_webhooks_api.IncomingWebhooksApi

All URIs are relative to *https://api.firstquadrant.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**incoming_webhooks_mark_conversations_as_won_post**](#incoming_webhooks_mark_conversations_as_won_post) | **post** /incoming-webhooks/mark-conversations-as-won | Mark conversations as won

# **incoming_webhooks_mark_conversations_as_won_post**
<a name="incoming_webhooks_mark_conversations_as_won_post"></a>
> Success incoming_webhooks_mark_conversations_as_won_post(mark_conversation_as_won_body)

Mark conversations as won

Mark any open conversations with a particular lead as \"Won\" as a custom goal.

### Example

* Api Key Authentication (api_key):
```python
import openapi_client
from openapi_client.apis.tags import incoming_webhooks_api
from openapi_client.model.mark_conversation_as_won_body import MarkConversationAsWonBody
from openapi_client.model.success import Success
from pprint import pprint
# Defining the host is optional and defaults to https://api.firstquadrant.ai
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://api.firstquadrant.ai"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: api_key
configuration.api_key['api_key'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['api_key'] = 'Bearer'
# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = incoming_webhooks_api.IncomingWebhooksApi(api_client)

    # example passing only required values which don't have defaults set
    body = MarkConversationAsWonBody(None)
    try:
        # Mark conversations as won
        api_response = api_instance.incoming_webhooks_mark_conversations_as_won_post(
            body=body,
        )
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling IncomingWebhooksApi->incoming_webhooks_mark_conversations_as_won_post: %s\n" % e)
```
### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
body | typing.Union[SchemaForRequestBodyApplicationJson] | required |
content_type | str | optional, default is 'application/json' | Selects the schema and serialization of the request body
accept_content_types | typing.Tuple[str] | default is ('application/json', ) | Tells the server the content type(s) that are accepted by the client
stream | bool | default is False | if True then the response.content will be streamed and loaded from a file like object. When downloading a file, set this to True to force the code to deserialize the content to a FileSchema file
timeout | typing.Optional[typing.Union[int, typing.Tuple]] | default is None | the timeout used by the rest client
skip_deserialization | bool | default is False | when True, headers and body will be unset and an instance of api_client.ApiResponseWithoutDeserialization will be returned

### body

# SchemaForRequestBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**MarkConversationAsWonBody**](../../models/MarkConversationAsWonBody.md) |  | 


### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#incoming_webhooks_mark_conversations_as_won_post.ApiResponseFor200) | Successfully queued operation

#### incoming_webhooks_mark_conversations_as_won_post.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor200ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor200ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**Success**](../../models/Success.md) |  | 


### Authorization

[api_key](../../../README.md#api_key)

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

