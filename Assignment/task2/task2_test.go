package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestEvaluateExpressions(t *testing.T) {
	// Mock API endpoint
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate API response with dummy results
		response := MathJSResponse{
			Result: []string{
				"2 * 4 * 4",
				"5 / (7 - 5)",
				"sqrt(5^2 - 4^2)",
				"sqrt(-3^2 - 4^2)",
			},
			Error: "",
		}
		jsonResponse, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}))
	defer mockAPI.Close()

	// Set the API endpoint to the mock server URL

	expressions := []string{
		"2 * 4 * 4",
		"5 / (7 - 5)",
		"sqrt(5^2 - 4^2)",
		"sqrt(-3^2 - 4^2)",
	}

	// Redirect stdout to capture the output
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run the evaluation
	evaluateExpressions(expressions)

	// Capture the output
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = oldStdout

	// Expected output
	expectedOutput := `2 * 4 * 4 => 32
5 / (7 - 5) => 2.5
sqrt(5^2 - 4^2) => 3
sqrt(-3^2 - 4^2) => 5i
`

	// Compare the expected output with the captured output
	if string(out) != expectedOutput {
		t.Errorf("Unexpected output. Expected:\n%s\n\nActual:\n%s", expectedOutput, string(out))
	}
}

func TestSendRequest(t *testing.T) {
	// Mock API endpoint
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate API response with dummy results
		response := MathJSResponse{
			Result: []string{
				"2 * 4 * 4",
				"5 / (7 - 5)",
				"sqrt(5^2 - 4^2)",
				"sqrt(-3^2 - 4^2)",
			},
			Error: "",
		}
		jsonResponse, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}))
	defer mockAPI.Close()

	expressions := []string{
		"2 * 4 * 4",
		"5 / (7 - 5)",
		"sqrt(5^2 - 4^2)",
		"sqrt(-3^2 - 4^2)",
	}

	results := make(chan string, len(expressions))

	sendRequest(expressions, results)

	// Expected results
	expectedResults := []string{
		"2 * 4 * 4 => 32",
		"5 / (7 - 5) => 2.5",
		"sqrt(5^2 - 4^2) => 3",
		"sqrt(-3^2 - 4^2) => 5i",
	}

	// Check if the results received match the expected results
	for i := 0; i < len(expressions); i++ {
		result := <-results
		if result != expectedResults[i] {
			t.Errorf("Unexpected result. Expected: %s, Actual: %s", expectedResults[i], result)
		}
	}
}
