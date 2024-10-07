package tests

import (
	"FuzzerBuzzer/internal/fuzz_logic"
	"FuzzerBuzzer/internal/generator"
	"FuzzerBuzzer/internal/http"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestHTTPClient_Post tests the HTTP client's POST request
func TestHTTPClient_Post(t *testing.T) {
	// Mock server to simulate an HTTP endpoint
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer mockServer.Close()

	// Set up the client
	client := http.NewClient(map[string]string{
		"Content-Type": "application/json",
	})

	// Test the Post function
	resp, err := client.Post(mockServer.URL, "application/json", bytes.NewBuffer([]byte(`{"test":"data"}`)))
	if err != nil {
		t.Fatalf("Failed to send POST request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}
}

// TestInputGenerator_GenerateRandomString tests the random string generation function
func TestInputGenerator_GenerateRandomString(t *testing.T) {
	ig := generator.NewInputGenerator(42)

	// Test random string generation
	randomStr := ig.GenerateRandomString(10)
	if len(randomStr) != 10 {
		t.Errorf("Expected random string length of 10, got %d", len(randomStr))
	}
}

// TestFuzzer_Start tests the fuzzing logic by mocking HTTP requests
func TestFuzzer_Start(t *testing.T) {
	// Mock server to simulate a target endpoint
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer mockServer.Close()

	// Create the HTTP client and fuzz logic
	client := http.NewClient(map[string]string{
		"Content-Type": "application/json",
	})
	fuzzer := fuzz_logic.NewFuzzer(client, mockServer.URL)

	// Simulate a few iterations of fuzzing
	for i := 0; i < 5; i++ {
		// Generate random input
		input := fuzzer.GenerateRandomInput()

		// Ensure the input is not empty
		if len(input) == 0 {
			t.Errorf("Generated empty input on iteration %d", i)
		}

		// Send the fuzzing request
		resp, err := client.Post(mockServer.URL, "application/json", bytes.NewReader(input))
		if err != nil {
			t.Fatalf("Fuzzing request failed: %v", err)
		}

		// Ensure response is successful
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code 200, got %d", resp.StatusCode)
		}
	}
}
