package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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
		fmt.Println("Json file invalid")
		return
	}

	// added input protection on client side
	phoneNumber, err := service.CrossUuidToPhone(httpPayload.StructUuid)
	if err != nil {
		fmt.Println("Couldn't cross check UUID -> PHONE |ERROR|	", err)
		http.Error(w, "Couldn't cross check UUID -> PHONE |ERROR|	", http.StatusConflict)
		return
	}

	authResponse, err := service.SendotpSupabase(phoneNumber, httpPayload.StructToken)
	if err != nil {
		fmt.Println("Couldn't request supabase |ERROR|	", err)
		http.Error(w, "Coudln't request Supabase", http.StatusConflict)
		return
	}

	phoneNumberOff := authResponse.User.Phone
	if !strings.HasPrefix(phoneNumberOff, "+") {
		phoneNumberOff = "+" + phoneNumberOff
	}

	value, err := service.CrossPhonetoUuid(phoneNumberOff)
	if err != nil {
		fmt.Println("Couldn't cross check Phone -> UUID on redis |ERROR|	", err)
		http.Error(w, "Couldn't cross check Phone -> UUID on redis", http.StatusConflict)
		return
	}

	if value == httpPayload.StructUuid {
		response := model.ServerJWTresponse{
			StructAccessToken: authResponse.AccessToken,
		}
		if response.StructAccessToken == "" {
			http.Error(w, "JWT is empty", http.StatusPreconditionRequired)
		}
		fmt.Println("token:	", authResponse.AccessToken)

		fmt.Println("Struct token: ", response)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Error something went wrong try again")
	}
}
