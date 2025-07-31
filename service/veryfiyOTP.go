package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Gustavo-DCosta/server/model"
)

func VerifyOTPSupabase(PhoneNumber string, Token string, UUID string) (string, error) {
	url := os.Getenv("PhoneOtpUrlVerify")
	if url == "" {
		return "", fmt.Errorf("couldn't get the URL from environment variables")
	}

	verifyPayload := model.VerifyOTPpayload{
		Type:  "sms",
		Phone: PhoneNumber,
		Token: Token,
	}

	jsonPayload, err := json.Marshal(verifyPayload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal JSON: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Add headers
	req.Header.Set("apikey", os.Getenv("SUPABASE_KEY")) // or however you store your API key
	req.Header.Set("Content-Type", "application/json")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	// Read the raw response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	// Save to file with timestamp
	var verifySupabaseOTPrep model.VerifySupabaseResponse

	err = json.Unmarshal(body, &verifySupabaseOTPrep)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return verifySupabaseOTPrep.AccesToken, nil
}
