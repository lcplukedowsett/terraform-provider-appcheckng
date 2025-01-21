package main

import (
	"fmt"
	"os"
	"terraform-provider-appcheckng/client"
)

func main() {
	apiKey := os.Getenv("APPCHECK_API_KEY")
	endpoint := os.Getenv("APPCHECK_ENDPOINT")

	if apiKey == "" || endpoint == "" {
		fmt.Println("API key or Endpoint not set in environment variables")
		return
	}

	appCheckClient := client.NewAppCheckClient(apiKey, endpoint)
	token, err := appCheckClient.Authenticate()
	if err != nil {
		fmt.Println("Error authenticating:", err)
		return
	}

	fmt.Println("Authentication token:", token)
}
