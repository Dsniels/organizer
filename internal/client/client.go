package client

import (
	"os"
	"strings"

	auth "github.com/microsoft/kiota-authentication-azure-go"
	graph "github.com/microsoftgraph/msgraph-beta-sdk-go"
)

func GetMSGraphClient() *graph.GraphServiceClient {
	credentials, err := getMSCredentials()
	if err != nil {
		panic(err)
	}
	provider, err := auth.NewAzureIdentityAuthenticationProviderWithScopes(credentials, strings.Split(os.Getenv("SCOPES"), ","))
	if err != nil {
		panic(err)
	}
	adapter, err := graph.NewGraphRequestAdapter(provider)
	if err != nil {
		panic(err)
	}

	client := graph.NewGraphServiceClient(adapter)

	return client

}
