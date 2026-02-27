package dtos

type Wallet struct {
	Status   string  `json:"status"`
	Balance  float64 `json:"balance"`
	WalletID string  `json:"wallet_id"`
}
