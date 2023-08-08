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

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func createWallet(walletPath string) *gateway.Wallet {
	walletPath = filepath.Clean(walletPath)

	// If wallet exist, the existing wallet will be returned instead
	wallet, err := gateway.NewFileSystemWallet(walletPath)
	checkError(err)

	return wallet
}

func newIdentityFromFile(mspid string, certpath string, keypath string) *gateway.X509Identity {
	certbytes, err := os.ReadFile(filepath.Clean(certpath))

	checkError(err)

	keybytes, err := os.ReadFile(filepath.Clean(certpath))

	checkError(err)

	return gateway.NewX509Identity(mspid, string(certbytes), string(keybytes))
}

func parseEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	checkError(err)
}

func getGateway(identityLabel string, cfg *Config) *gateway.Gateway {
	wallet := createWallet(cfg.Identity.WalletPath)

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

func getNetwork(gw gateway.Gateway, channelName string) *gateway.Network {
	network, err := gw.GetNetwork(channelName)
	checkError(err)
	return network
}

func getContract(network gateway.Network, chaincodeName string, contractName string) *gateway.Contract {
	contract := network.GetContractWithName(chaincodeName, contractName)
	return contract
}
