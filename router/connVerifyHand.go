package router

import (
	"net/http"
)

func HandleConnVerification(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "405: METHOD NOT ALLOWED", http.StatusMethodNotAllowed)
	}

}
