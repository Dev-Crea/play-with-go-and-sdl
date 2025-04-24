# \DataAPI

All URIs are relative to *https://api.spacetraders.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSupplyChain**](DataAPI.md#GetSupplyChain) | **Get** /market/supply-chain | Get Supply Chain



## GetSupplyChain

> GetSupplyChain200Response GetSupplyChain(ctx).Execute()

Get Supply Chain



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
	resp, r, err := apiClient.DataAPI.GetSupplyChain(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataAPI.GetSupplyChain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSupplyChain`: GetSupplyChain200Response
	fmt.Fprintf(os.Stdout, "Response from `DataAPI.GetSupplyChain`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSupplyChainRequest struct via the builder pattern


### Return type

[**GetSupplyChain200Response**](GetSupplyChain200Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

