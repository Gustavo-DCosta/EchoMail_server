package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Gustavo-DCosta/server/model"
	"github.com/Gustavo-DCosta/server/service"
)

func handleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not accepted", http.StatusMethodNotAllowed)
		return
	}

	var data model.SignupRequest
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Json file invalid", http.StatusBadRequest)
		return
	}

	if data.Email_Address == "" || data.Phone_Number == "" {
		http.Error(w, "Email and phone number are required", http.StatusBadRequest)
		return
	}

	uuid := service.GenerateUUID()

	err := service.SaveDataToRedis1way(uuid, data.Phone_Number)
	if err != nil {
		fmt.Println("Error writing to redis", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	err = service.SaveDataToRedis2ways(uuid, data.Phone_Number)
	if err != nil {
		fmt.Println("Error writing to redis", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Send OTP to Supabase BEFORE responding
	if err := service.SendCredentialsSupabase(data.Phone_Number, data.Email_Address); err != nil {
		http.Error(w, "Failed to send OTP", http.StatusInternalServerError)
		return
	}

	// Send response only once, after everything succeeds
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"message": "OTP sent successfully",
		"uuid":    uuid,
	}
	json.NewEncoder(w).Encode(response)
}
