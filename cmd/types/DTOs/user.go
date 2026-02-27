package dtos

type User struct {
	UserID       string `json:"id"`
	LastName     string `json:"last_name"`
	WalletID     string `json:"wallet_id"`
	FirstName    string `json:"first_name"`
	PhoneNumber  string `json:"phone_number"`
	NationalCode string `json:"national_code"`
}
