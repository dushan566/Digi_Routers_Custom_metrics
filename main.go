package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"router/api"
	"router/configs"
	"time"
)

func init() {
	// Load environment variables
	api.GetEnvVars()
	drmUserName := os.Getenv("DRMUSER")
	drmPassword := os.Getenv("DRMPASS")
	drmApiEndpoint := os.Getenv("DEVICE_ENDPOINT")

	apiUrl := "https://" + drmUserName + ":" + drmPassword + "@" + drmApiEndpoint

	// Retry logic in case of network errors
	var resp *http.Response
	var err error
	for retries := 0; retries < 3; retries++ { // Try up to 3 times
		resp, err = http.Get(apiUrl)
		if err == nil && resp.StatusCode == http.StatusOK {
			break
		}
		fmt.Printf("Attempt %d: No response from request or error occurred\n", retries+1)
		time.Sleep(2 * time.Second) // Wait before retrying
	}
	if err != nil {
		fmt.Println("Failed to get response after retries:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var responseMap map[string]interface{}
	err = json.Unmarshal(body, &responseMap)
	if err != nil {
		fmt.Println("Error parsing JSON response:", err)
		return
	}

	// Generate dynamic functions/vars based on HTTP response
	configs.TemplateGenerator("templates/vars.txt", responseMap, "configs/generatedVars.go")
	configs.TemplateGenerator("templates/funcs.txt", responseMap, "configs/generatedFuncs.go")
}

func main() {
	// Run Prometheus handler and HTTP server
	configs.PrometheusHandler()
}
