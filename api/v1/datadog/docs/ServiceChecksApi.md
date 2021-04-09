# \ServiceChecksApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubmitServiceCheck**](ServiceChecksApi.md#SubmitServiceCheck) | **Post** /api/v1/check_run | Submit a Service Check



## SubmitServiceCheck

> IntakePayloadAccepted SubmitServiceCheck(ctx).Body(body).Execute()

Submit a Service Check



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())

    body := []datadog.ServiceCheck{*datadog.NewServiceCheck("app.ok", "app.host1", datadog.ServiceCheckStatus(0), []string{"Tags_example"})} // []ServiceCheck | Service Check request body.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceChecksApi.SubmitServiceCheck(ctx).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceChecksApi.SubmitServiceCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitServiceCheck`: IntakePayloadAccepted
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from ServiceChecksApi.SubmitServiceCheck:\n%s\n", responseContent)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitServiceCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]ServiceCheck**](ServiceCheck.md) | Service Check request body. | 

### Return type

[**IntakePayloadAccepted**](IntakePayloadAccepted.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
