package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func SendCredentialsSupabase(phone_number string, email_adress string) error {
	url := os.Getenv("PhoneOtpUrl")
	if url == "" {
		fmt.Println("Error grabbing register URL endpoint")
	}

	supabaseData := map[string]interface{}{
		"phone": phone_number,
		"data": map[string]interface{}{ //<- this is the key part
			"email": email_adress, // use correct field name    // add anything else you want
		},
	}

	jsonData, err := json.Marshal(supabaseData)
	if err != nil {
		fmt.Println("Error marsheling data", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating new request", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("apikey", os.Getenv("Supabase_Anon_Key"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error....", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("supabase returned status: %d", resp.StatusCode)
	}

	return nil
}
