# \TeamsApi

All URIs are relative to *https://api.firstquadrant.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamsMeGet**](TeamsApi.md#TeamsMeGet) | **Get** /teams/me | Get team



## TeamsMeGet

> Team TeamsMeGet(ctx).Execute()

Get team



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsMeGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsMeGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsMeGet`: Team
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsMeGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsMeGetRequest struct via the builder pattern


### Return type

[**Team**](Team.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

