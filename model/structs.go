package model

type SignupRequest struct {
	Phone_Number  string `json:"phone_number"`
	Email_Address string `json:"email_address"`
}

type VerifyOTPRequest struct {
	UUID  string `json:"uuid"`
	Token string `json:"token"`
}

type UUIDResponse struct {
	UUID string `json:"uuid"`
}

type JwtStruct struct {
}

type VerifyOTPpayload struct {
	Type  string `json:"type"`
	Phone string `json:"phone"`
	Token string `json:"token"`
}

type OTPResponse struct {
	Success bool   `json:"success"`
	JWT     string `json:"jwt"`
}

type VerifySupabaseResponse struct {
	AccesToken  string `json:"access_token"`
	PhoneNumber string `json:"phone"`
}

type ConnRequestBody struct {
	StructPhone     string `json:"phone"`
	StructEmaill    string `json:"email"`
	StructAccStatus bool   `json:"status"`
}
