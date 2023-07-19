package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// variiables declaration
const (
	APIEndpoint       = "http://api.mathjs.org/v4/"
	MaxRequestsPerSec = 50
)

// struct for API request
type MathJSRequest struct {
	Expr      []string `json:"expr"`
	Precision int      `json:"precision"`
}

// struct for api response
type MathJSResponse struct {
	Result []string `json:"result"`
	Error  string   `json:"error"`
}

/*
Func to evaluate the espressions
1) Btach the bulk expressions
2) Hit the  api with bulk expressions
3) Get response and print it
*/
func evaluateExpressions(expressions []string) {
	results := make(chan string)

	// Divide expressions into batches
	numExpressions := len(expressions)
	numBatches := (numExpressions + MaxRequestsPerSec - 1) / MaxRequestsPerSec

	// Process each batch concurrently
	for i := 0; i < numBatches; i++ {
		startIndex := i * MaxRequestsPerSec
		endIndex := (i + 1) * MaxRequestsPerSec
		if endIndex > numExpressions {
			endIndex = numExpressions
		}

		exprBatch := expressions[startIndex:endIndex]
		
		// Execute each batch in a separate goroutine
		go func(batch []string) {
			sendRequest(batch, results)
		}(exprBatch)
	}

	// Collect results from the channel
	for i := 0; i < numExpressions; i++ {
		result := <-results
		fmt.Println(result)
	}
}

func sendRequest(expressions []string, results chan<- string) {
	reqBody, err := json.Marshal(MathJSRequest{
		Expr:      expressions,
		Precision: 14,
	})
	if err != nil {
		results <- fmt.Sprintf("Failed to create request body: %v", err)
		return
	}

	resp, err := http.Post(APIEndpoint, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		results <- fmt.Sprintf("Failed to send request to the API: %v", err)
		return
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		results <- fmt.Sprintf("Failed to read response body: %v", err)
		return
	}

	var apiResponse MathJSResponse
	err = json.Unmarshal(respBody, &apiResponse)
	if err != nil {
		results <- fmt.Sprintf("Failed to parse API response: %v", err)
		return
	}

	if apiResponse.Error != "" {
		results <- fmt.Sprintf("API Error: %s", apiResponse.Error)
		return
	}

	for i, result := range apiResponse.Result {
		results <- fmt.Sprintf("%s => %s", expressions[i], result)
	}
}

func main() {
	expressions := []string{
		"2 * 4 * 4",
		"5 / (7 - 5)",
		"sqrt(5^2 - 4^2)",
		"sqrt(-3^2 - 4^2)",
	}

	evaluateExpressions(expressions)
}
