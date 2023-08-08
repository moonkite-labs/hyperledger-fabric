package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Identity struct {
		MspID        string `envconfig:"MSP_ID"`
		CertPath     string `envconfig:"CERT_PATH"`
		KeystorePath string `envconfig:"KEYSTORE_PATH"`
		WalletPath   string `envconfig:"WALLET_PATH"`
		CCPPath      string `envconfig:"CCP_PATH"`
	}
}

// Stops execution if an error is found
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// Function to create or get a wallet from the given wallet path
func CreateWallet(walletPath string) *gateway.Wallet {
	walletPath = filepath.Clean(walletPath)

	// If wallet exist, the existing wallet will be returned instead
	wallet, err := gateway.NewFileSystemWallet(walletPath)
	checkError(err)

	return wallet
}

// Create a new X509Identity from associated mspid, issued certificate file and private key file
func NewIdentityFromFile(mspid string, certpath string, keypath string) *gateway.X509Identity {
	certbytes, err := os.ReadFile(filepath.Clean(certpath))

	checkError(err)

	keybytes, err := os.ReadFile(filepath.Clean(certpath))

	checkError(err)

	return gateway.NewX509Identity(mspid, string(certbytes), string(keybytes))
}

// Parse environment variables into a Config struct
func ParseEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	checkError(err)
}

// Return a gateway connected using the given identity
func GetGateway(identityLabel string, cfg *Config) *gateway.Gateway {
	wallet := CreateWallet(cfg.Identity.WalletPath)

	if !wallet.Exists(identityLabel) {
		log.Panicf("Identity %s not found in wallet %s!", identityLabel, cfg.Identity.WalletPath)
	}

	ccp := config.FromFile(cfg.Identity.CCPPath)

	gw, err := gateway.Connect(
		gateway.WithConfig(ccp),
		gateway.WithIdentity(wallet, identityLabel),
	)

	checkError(err)

	return gw
}

// Get a channel instance from a connected gateway using channel name
func GetNetwork(gw gateway.Gateway, channelName string) *gateway.Network {
	network, err := gw.GetNetwork(channelName)
	checkError(err)
	return network
}

// Get a contract instance from a connected channel using chaincode name and contract name
func GetContract(network gateway.Network, chaincodeName string, contractName string) *gateway.Contract {
	contract := network.GetContractWithName(chaincodeName, contractName)
	return contract
}
