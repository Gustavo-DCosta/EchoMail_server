package model

// The Struct prefix on the structure fields helps disambiguate
type ServerConnHandlerResponse struct {
	StructUUID string `json:"uuid"`
}

type ConnRequestBody struct {
	StructPhone     string `json:"phone"`
	StructEmaill    string `json:"email"`
	StructAccStatus bool   `json:"status"`
}

type VerifyRequestBody struct {
	StructUuid  string `json:"uuid"`
	StructToken string `json:"token"`
}

type VerifyServerRequest struct {
	StructType        string `json:"type"`
	StructPhoneNumber string `json:"phone"`
	StructToken       string `json:"token"`
}

type SupabaseAuthResponse struct {
	AccessToken string `json:"access_token"`
	User        struct {
		Phone string `json:"phone"`
	} `json:"user"`
}

type ServerJWTresponse struct {
	StructAcessToke string `json:"acess_token"`
}
