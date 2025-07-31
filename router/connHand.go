package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Gustavo-DCosta/server/model"
	"github.com/Gustavo-DCosta/server/service"
)

func HandleConn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "405: METHOD NOT ALLOWED", http.StatusMethodNotAllowed)
	}

	var httpPayload model.ConnRequestBody
	err := json.NewDecoder(r.Body).Decode(&httpPayload)
	if err != nil {
		http.Error(w, "Json file invalid", http.StatusBadRequest)
		return
	}

	if httpPayload.StructPhone == "" || httpPayload.StructEmaill == "" {
		http.Error(w, "Email and phone number are required", http.StatusBadRequest)
		return
	}

	if httpPayload.StructAccStatus == false {
		err := service.LoginAccSupabase(httpPayload.StructPhone, httpPayload.StructEmaill)
		if err != nil {
			fmt.Println("Problem sending reques to supabase: | ERROR CODE: ", err)
		}
	} else {
		err := service.CreateAccSupabase(httpPayload.StructPhone, httpPayload.StructEmaill)
		if err != nil {
			fmt.Println("Problem sending reques to supabase: | ERROR CODE: ", err)
		}
	}
}
