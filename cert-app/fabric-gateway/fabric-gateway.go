package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"flag"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	MspID        string `envconfig:"MSP_ID"`
	CertPath     string `envconfig:"CERT_PATH"`
	KeystorePath string `envconfig:"KEYSTORE_PATH"`
	WalletPath   string `envconfig:"WALLET_PATH"`
	CCPPath      string `envconfig:"CCP_PATH"`
	Label        string `envconfig:"LABEL"`
}

// Function to create or get a wallet from the given wallet path
func CreateWallet(walletPath string) *gateway.Wallet {
	walletPath = filepath.Clean(walletPath)

	// If wallet exist, the existing wallet will be returned instead
	wallet, err := gateway.NewFileSystemWallet(walletPath)
	CheckError(err)

	return wallet
}

// Create a new X509Identity from associated mspid, issued certificate file and private key file
func NewIdentityFromFile(mspid string, certpath string, keypath string) *gateway.X509Identity {
	certbytes, err := os.ReadFile(filepath.Clean(certpath))

	CheckError(err)

	keybytes, err := os.ReadFile(filepath.Clean(certpath))

	CheckError(err)

	return gateway.NewX509Identity(mspid, string(certbytes), string(keybytes))
}

// Parse environment variables into a Config struct
func ParseEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	CheckError(err)
}

// Return a gateway connected using the given identity
func GetGateway(cfg Config) *gateway.Gateway {
	wallet := CreateWallet(cfg.WalletPath)

	if !wallet.Exists(cfg.Label) {
		log.Panicf("Identity %s not found in wallet %s!", cfg.Label, cfg.WalletPath)
	}

	ccp := config.FromFile(cfg.CCPPath)

	gw, err := gateway.Connect(
		gateway.WithConfig(ccp),
		gateway.WithIdentity(wallet, cfg.Label),
	)

	CheckError(err)

	return gw
}

// Get a channel instance from a connected gateway using channel name
func GetNetwork(gw gateway.Gateway, channelName string) *gateway.Network {
	network, err := gw.GetNetwork(channelName)
	CheckError(err)
	return network
}

// Get a contract instance from a connected channel using chaincode name and contract name
func GetContract(network gateway.Network, chaincodeName string, contractName string) *gateway.Contract {
	contract := network.GetContractWithName(chaincodeName, contractName)
	return contract
}

var Usage = func() {
	var CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	fmt.Fprintf(CommandLine.Output(), "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}

func parseArgs(label *string, walletPath *string, ccpPath *string, mspid *string, certPath *string, keystorePath *string) {
	flag.StringVar(mspid, "mspid", "", "The MSP ID managing the identity")
	flag.StringVar(label, "label", "", "A unique name to store and get the identity from the wallet")
	flag.StringVar(walletPath, "walletPath", "", "The path to create or get a wallet")
	flag.StringVar(ccpPath, "ccpPath", "", "The path to the associated connection profile")
	flag.StringVar(certPath, "certPath", "", "The path to the signed identity certificate issued by the CA and MSP")
	flag.StringVar(keystorePath, "keystorePath", "", "The path to the identity's private key issued by the CA and MSP")
}

func main() {
	// Some required
	var label, walletPath, ccpPath, mspid, certPath, keystorePath string
	var isEnv = flag.Bool("env", true, "Use environment variables to get the required values")
	var isNew = flag.Bool("new", false, "Put the new identity into the wallet by label, this argument will not be parsed from env vars")
	var isHelp = flag.Bool("h", false, "Print help message")

	parseArgs(&label, &walletPath, &ccpPath, &mspid, &certPath, &keystorePath)

	flag.Parse()

	if *isHelp {
		Usage()
		os.Exit(0)
	}

	var cfg Config
	if *isEnv {
		ParseEnv(&cfg)
	} else {

		cfg = Config{
			MspID:        mspid,
			CertPath:     certPath,
			KeystorePath: keystorePath,
			WalletPath:   walletPath,
			CCPPath:      ccpPath,
			Label:        label,
		}
	}

	if *isNew {
		wallet := CreateWallet(walletPath)
		file, err := os.Lstat(keystorePath)
		CheckError(err)

		// If the given path is a directory, take the first file
		// (The generated filename is random each time, so giving the directory is easier for testing)
		if file.IsDir() {
			files, err := os.ReadDir(cfg.KeystorePath)

			CheckError(err)

			keystorePath = cfg.KeystorePath + files[0].Name()
		}

		identity := NewIdentityFromFile(mspid, certPath, keystorePath)
		wallet.Put(label, identity)
	}

	gw := GetGateway(cfg)
	fmt.Printf("Connected to the gateway %+v\n", gw)
}
