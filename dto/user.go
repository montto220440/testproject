package dto

import "time"

type ListRequest struct {
	Id       int    `json:"id" binding:"required"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Age      string `json:"age"`
	Gender   string `json:"gender"`
}

type ListResponse struct {
	ID struct {
		Oid string `json:"$oid"`
	} `json:"_id"`
	Username       string `json:"username"`
	UsernamePrefix string `json:"username_prefix"`
	DateRegis      struct {
		Date time.Time `json:"$date"`
	} `json:"date_regis"`
	UsernameOld string `json:"username_old"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	FullName    string `json:"full_name"`
	BankList    []struct {
		BankCode   string `json:"bank_code"`
		BankName   string `json:"bank_name"`
		BankNumber string `json:"bank_number"`
		Status     int    `json:"status"`
		Active     int    `json:"active"`
	} `json:"bank_list"`
	Point         int    `json:"point"`
	PhoneVerify   string `json:"phone_verify"`
	OtpCode       string `json:"otp_code"`
	Token         string `json:"token"`
	IP            string `json:"ip"`
	Active        int    `json:"active"`
	SessionID     string `json:"session_id"`
	UserType      string `json:"user_type"`
	LineUserID    string `json:"line_userId"`
	LineUserLogin string `json:"line_user_login"`
	IsBonus       int    `json:"is_bonus"`
	Affiliate     string `json:"affiliate"`
	Clickid       string `json:"clickid"`
	Rank          int    `json:"rank"`
	ExpDeposit    int    `json:"exp_deposit"`
	BankLayer     struct {
		NumberLong string `json:"$numberLong"`
	} `json:"bank_layer"`
	IsFollow   int    `json:"is_follow"`
	HydraID    string `json:"hydra_id"`
	StartRegis struct {
		Date struct {
			NumberLong string `json:"$numberLong"`
		} `json:"$date"`
	} `json:"start_regis"`
	DeviceType    string `json:"device_type"`
	WithdrawSport bool   `json:"withdraw_sport"`
	TopupStatus   int    `json:"topup_status"`
	TopupBonus    int    `json:"topup_bonus"`
	TopupNobonus  int    `json:"topup_nobonus"`
	ProfileURL    string `json:"profile_url"`
}
