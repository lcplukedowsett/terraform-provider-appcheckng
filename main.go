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
	fmt.Println("AppCheck client created:", appCheckClient)
	// Remove the Authenticate method call
	// token, err := appCheckClient.Authenticate()
	// if err != nil {
	// 	fmt.Println("Error authenticating:", err)
	// 	return
	// }

	// fmt.Println("Authentication token:", token)
	fmt.Println("Client initialized successfully")
}
