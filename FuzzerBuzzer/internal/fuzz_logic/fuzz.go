package fuzz_logic

import (
	"bytes"
	"math/rand"
	"net/http"
	"time"
)

// Fuzzer struct that holds the HTTP client and target URL
type Fuzzer struct {
	Client    *http.Client
	TargetURL string
}

// NewFuzzer creates a new instance of Fuzzer
func NewFuzzer(client *http.Client, targetURL string) *Fuzzer {
	return &Fuzzer{
		Client:    client,
		TargetURL: targetURL,
	}
}

// Start begins the fuzzing process
func (f *Fuzzer) Start() {
	for {
		// Generate random input
		input := f.GenerateRandomInput()

		// Send the fuzzing request
		resp, err := f.Client.Post(f.TargetURL, "application/json", bytes.NewReader(input))
		if err != nil {
			// Handle the error (e.g., log it)
			continue
		}
		if resp != nil {
			resp.Body.Close()
		}

		// You can add logic to analyze the response here
	}
}

// GenerateRandomInput generates random input for fuzzing
func (f *Fuzzer) GenerateRandomInput() []byte {
	rand.Seed(time.Now().UnixNano())
	length := rand.Intn(100) // Generate random length between 0 and 100
	data := make([]byte, length)

	// Fill data with random bytes
	for i := 0; i < length; i++ {
		data[i] = byte(rand.Intn(256)) // Random byte value
	}

	return data
}
