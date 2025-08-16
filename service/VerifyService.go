package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Gustavo-DCosta/server/model"
)

func SendotpSupabase(phoneNumber, token string) (*model.SupabaseAuthResponse, error) {
	url := os.Getenv("PhoneOtpUrlVerify")
	if url == "" {
		fmt.Println("Couldn't get the url")
		return nil, errors.New("missing PhoneOtpUrlVerify env var")
	}

	verifyOtp := model.VerifyServerRequest{
		StructType:        "sms",
		StructPhoneNumber: phoneNumber,
		StructToken:       token,
	}

	verifyPayload, err := json.Marshal(verifyOtp)
	if err != nil {
		fmt.Println("Error marshaling request | ERROR:", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(verifyPayload))
	if err != nil {
		fmt.Println("Couldn't start HTTP request | ERROR:", err)
		return nil, err
	}

	anonKey := os.Getenv("Supabase_Anon_Key")
	req.Header.Set("apikey", anonKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request failed | ERROR:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		fmt.Printf("Supabase error: HTTP %d\nResponse body: %s\n", resp.StatusCode, string(bodyBytes))
		return nil, fmt.Errorf("supabase returned status %d", resp.StatusCode)
	}
	// decode the body
	var authResponse model.SupabaseAuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResponse); err != nil {
		fmt.Println("Failed to decode response | ERROR:", err)
		return nil, err
	}

	return &authResponse, nil
}
