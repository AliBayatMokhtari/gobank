package constants

type walletStatusRegistry struct {
	ACTIVE   string
	DEACTIVE string
	PENDING  string
}

func newWalletStatusRegistry() *walletStatusRegistry {
	return &walletStatusRegistry{
		ACTIVE:   "ACTIVE",
		PENDING:  "PENDING",
		DEACTIVE: "DEACTIVE",
	}
}

var WalletStatus = newWalletStatusRegistry()
