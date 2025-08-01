package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func CreateAccSupabase(phoneNumber, email string) (string, error) {
	uuid := GenerateUUID()
	err := WriteRedis2ways(uuid, phoneNumber)
	if err != nil {
		fmt.Println("Error writing data to redis | ERROR CODE: ", err)
		return " ", err
	}

	url := os.Getenv("PhoneOtpUrl")
	supabaseBody := map[string]interface{}{
		"phone":       phoneNumber,
		"create_user": true,
		"data": map[string]interface{}{ //<- this is the key part
			"email": email, // use correct field name    // add anything else you want
		},
	}
	requestPayload, err := json.Marshal(supabaseBody)
	if err != nil {
		fmt.Println("Error marsheling the request | ERROR CODE: ", err)
		return " ", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestPayload))
	if err != nil {
		fmt.Println("Error creating a new request | ERROR CODE: ", err)
		return " ", err
	}
	AnonKey := os.Getenv("Supabase_Anon_Key")

	req.Header.Set("apikey", AnonKey)
	req.Header.Set("Content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending the request | Error code: ", err)
		return " ", err
	}

	resp.Body.Close()
	return uuid, nil
}

func LoginAccSupabase(phoneNumber, email string) (string, error) {
	uuid := GenerateUUID()
	err := WriteRedis2ways(uuid, phoneNumber)
	if err != nil {
		fmt.Println("Error writing data to redis | ERROR CODE: ", err)
		return " ", err
	}

	url := os.Getenv("PhoneOtpUrl")
	supabaseBody := map[string]interface{}{
		"phone":       phoneNumber,
		"create_user": false,
		"data": map[string]interface{}{ //<- this is the key part
			"email": email, // use correct field name    // add anything else you want
		},
	}
	requestPayload, err := json.Marshal(supabaseBody)
	if err != nil {
		fmt.Println("Error marsheling the request | ERROR CODE: ", err)
		return " ", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestPayload))
	if err != nil {
		fmt.Println("Error creating a new request | ERROR CODE: ", err)
		return " ", err
	}
	AnonKey := os.Getenv("Supabase_Anon_Key")

	req.Header.Set("apikey", AnonKey)
	req.Header.Set("Content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending the request | Error code: ", err)
		return " ", err
	}

	resp.Body.Close()

	return uuid, nil
}
