package project

type WalletProvider string

const (
	WalletProviderMetamask      = "metamask"
	WalletProviderWalletConnect = "walletconnect"
)

type Wallet struct {
	Provider WalletProvider `json:"provider"`
	Address  string         `json:"address"`
}

// ValidWalletProvider
func ValidWalletProvider(wp WalletProvider) bool {
	walletProviders := []WalletProvider{
		WalletProviderMetamask,
		WalletProviderWalletConnect,
	}

	for _, walletProvider := range walletProviders {
		if walletProvider == wp {
			return true
		}
	}

	return false
}
