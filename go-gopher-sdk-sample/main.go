package main

import (
	"context"
	"fmt"
	"os"

	openapiclient "github.com/scraly/learning-go-by-examples/go-gopher-sdk"
)

func main() {

	//client := gophersapi.NewAPIClient(gophersapi.NewConfiguration())
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GophersApi.GophersGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GophersApi.GophersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GophersGet`: []Gopher
	fmt.Fprintf(os.Stdout, "Response from `GophersApi.GophersGet`: %v\n", resp)
}
