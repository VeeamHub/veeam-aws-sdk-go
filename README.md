# Veeam Backup for AWS SDK for Go

veeam-aws-sdk-go is the unofficial Veeam Backup for AWS SDK for the Go programming language.

This SDK was generated using the [OpenAPI Generator](https://openapi-generator.tech/). As such, any API call in the [Veeam Backup for AWS v3.0 REST API](https://helpcenter.veeam.com/docs/vbaws/rest/overview.html?ver=30) is present in this SDK.

## 📗 Documentation

### Installation

Use `go get` to retrieve the SDK to add it to your `GOPATH` workspace, or
project's Go module dependencies.

```bash
go get github.com/veeamhub/veeam-aws-sdk-go
```

To update the SDK use `go get -u` to retrieve the latest version of the SDK.

```bash
go get -u github.com/veeamhub/veeam-aws-sdk-go
```

#### Dependencies

The SDK includes a `vendor` folder containing the runtime dependencies of the SDK. The metadata of the SDK's dependencies can be found in the Go module file `go.mod`.

### Example

This example shows a complete working Go file which will list the names of all Backup Policies and demonstrate how to configure the timeout logic that will cancel the request if it takes too long. This example highlights how to login to a VBA appliance with trusted or self-signed certs, make a request, handle the error, process the response, and then logout.

```go
package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/veeamhub/veeam-aws-sdk-go/client"
)

// Retrieves and prints out all Backup Job names of the specified VBA server.
// NOTE: API port access is restricted by default. Make sure your EC2 instance Security Group allows the API port. 
//
// The Variables below need to be set according to your environment.
func main() {
	// Setting variables
	host := "ec2-53-90-267-234.compute-1.amazonaws.com:11005" // default API port 11005
	username := "jsmith"                                      // VBA username
	password := "password"                                    // VBA password

	// Setting constants
	const (
		apiVersion = "1.0-rev0"       // default API version (1.0-rev0)
		timeout    = 30 * time.Second // 30 seconds
		skipTls    = true             // skip TLS certificate verification
	)

	// Generating API client
	tlsClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: skipTls,
			},
		},
	}
	config := client.NewConfiguration()
	config.HTTPClient = tlsClient
	config.Host = host
	config.HTTPClient.Timeout = timeout
	config.Scheme = "https"
	veeam := client.NewAPIClient(config)

	// Authenticating to VBA API
	login, r, err := veeam.TokenApi.Authenticate(context.Background()).XApiVersion(apiVersion).GrantType("password").Username(username).Password(password).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		panic(err)
	}

	// Setting authorization
	// From this point, all calls will 'auth' as the context
	auth := context.WithValue(context.Background(), client.ContextAccessToken,
		*login.AccessToken)

	// Retrieving all backup policies
	policies, r, err := veeam.PoliciesApi.PoliciesGetAll(auth).XApiVersion(apiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		panic(err)
	}

	// Printing out policy names
	// Working with the response payload
	fmt.Printf("Policy Names:\n\n")
	for _, policy := range policies.Results {
		fmt.Println(policy.Name)
	}

	// Logging out session
	_, err = veeam.TokenApi.SignOut(auth).XApiVersion(apiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		panic(err)
	}
}
```

## ✍ Contributions

We welcome contributions from the community! We encourage you to create [issues](https://github.com/VeeamHub/veeam-aws-sdk-go/issues/new/choose) for Bugs & Feature Requests and submit Pull Requests. For more detailed information, refer to our [Contributing Guide](CONTRIBUTING.md).

## 🤝🏾 License

* [MIT License](LICENSE)

## 🤔 Questions

If you have any questions or something is unclear, please don't hesitate to [create an issue](https://github.com/VeeamHub/veeam-aws-sdk-go/issues/new/choose) and let us know!
