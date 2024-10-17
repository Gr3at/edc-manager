package utils

import (
	"testing"
)

func TestEncryptDecryptKey(t *testing.T) {
	// Define test cases as a slice of structs
	testCases := []struct {
		name            string
		credentials     string
		serverSecretKey string
	}{
		{
			name:            "Case 1: Normal API key and secret",
			credentials:     "dummy-api-secret-kesggegry",
			serverSecretKey: "}JL|K=1RYhm#h]@m(OqX(@$x[~^uWpN{",
		},
		{
			name:            "Case 2: API key with special characters",
			credentials:     "!@#dummy-API$%^&*-secret",
			serverSecretKey: "pI5tK2fVpzL7H5$wX#iR3a5cZ",
		},
		{
			name:            "Case 3: Short API key",
			credentials:     "key",
			serverSecretKey: "lG29KpwJf4@6LoL&$z91O7W#12",
		},
		{
			name:            "Case 4: Long API key",
			credentials:     "super-long-api-key-with-extra-characters-1234567890-!@#$%^&*()_+",
			serverSecretKey: "Fj*23#7KLMl8PQ!qUvCz78!@&(FG123L",
		},
		{
			name:            "Case 5: Long Secret Key",
			credentials:     "dummy-api-secret-ABCD1234",
			serverSecretKey: "super-long-secret-key-that-is-overly-complex-and-needs-to-be-tested-for-boundaries-!@#",
		},
	}

	// Iterate through each test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Encrypt the Credentials
			encryptedCredentials, err := EncryptKey(tc.credentials, tc.serverSecretKey)
			if err != nil {
				t.Fatalf("Expected no error during encryption, got %v", err)
			}

			// Decrypt the encrypted Credentials
			decryptedCredentials, err := DecryptKey(encryptedCredentials, tc.serverSecretKey)
			if err != nil {
				t.Fatalf("Expected no error during decryption, got %v", err)
			}

			// Check if the decrypted Credentials matches the original one
			if decryptedCredentials != tc.credentials {
				t.Fatalf("Expected decrypted Credentials %q, got %q", tc.credentials, decryptedCredentials)
			}

			t.Logf("Test %s passed. Decrypted Credentials: %+v", tc.name, decryptedCredentials)
		})
	}
}
