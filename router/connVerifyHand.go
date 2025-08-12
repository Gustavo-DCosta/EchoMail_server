package router

import (
	"encoding/json"
	"net/http"

	"github.com/Gustavo-DCosta/server/model"
	"github.com/Gustavo-DCosta/server/service"
)

func HandleConnVerification(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "405: METHOD NOT ALLOWED", http.StatusMethodNotAllowed)
	}

	var httpPayload model.VerifyRequestBody
	err := json.NewDecoder(r.Body).Decode(&httpPayload)
	if err != nil {
		http.Error(w, "Json file invalid", http.StatusBadRequest)
		return
	}

	phoneNumber, err := service.CrossUuidToPhone(httpPayload.StructUuid)
	if err != nil {
		http.Error(w, "Couldn't request redis", http.StatusConflict)
		return
	}

	authResponse, err := service.SendotpSupabase(phoneNumber, httpPayload.StructToken)
	if err != nil {
		http.Error(w, "Coudln't request Supabase", http.StatusConflict)
	}

	value, err := service.CrossPhonetoUuid(authResponse.User.Phone)
	if err != nil {
		http.Error(w, "conflict", http.StatusConflict)
	}

	if value == httpPayload.StructUuid {
		response := model.ServerJWTresponse{
			StructAcessToke: authResponse.AccessToken,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Error something went wrong try again")
	}
}
