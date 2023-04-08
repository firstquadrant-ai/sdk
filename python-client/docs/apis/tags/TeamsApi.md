<a name="__pageTop"></a>
# openapi_client.apis.tags.teams_api.TeamsApi

All URIs are relative to *https://api.firstquadrant.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**teams_me_get**](#teams_me_get) | **get** /teams/me | Get team

# **teams_me_get**
<a name="teams_me_get"></a>
> Team teams_me_get()

Get team

Get the team associated with the API key.

### Example

* Api Key Authentication (api_key):
```python
import openapi_client
from openapi_client.apis.tags import teams_api
from openapi_client.model.team import Team
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
    api_instance = teams_api.TeamsApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        # Get team
        api_response = api_instance.teams_me_get()
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling TeamsApi->teams_me_get: %s\n" % e)
```
### Parameters
This endpoint does not need any parameter.

### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
200 | [ApiResponseFor200](#teams_me_get.ApiResponseFor200) | Team object

#### teams_me_get.ApiResponseFor200
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor200ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor200ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**Team**](../../models/Team.md) |  | 


### Authorization

[api_key](../../../README.md#api_key)

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

