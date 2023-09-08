package main

import (
	"crypto/x509"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"flag"

	db "fabric-gateway/service/db"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"fabric-gateway/config"
	"fabric-gateway/utils"
)

// Create a new X509Identity from associated mspid, issued certificate file and private key file
func NewIdentityFromFile(mspid string, certpath string) (*identity.X509Identity, error) {

	cert, err := LoadCertificate(certpath)

	utils.CheckError(err)

	return identity.NewX509Identity(mspid, cert)
}

// Return a gateway connected using the given identity
func GetGateway(cfg config.Config) *client.Gateway {
	wallet := db.PostgreWalletService{}
	err := wallet.Connect(cfg.DB_HOST, cfg.DB_USER, cfg.DB_PASS, cfg.DB_NAME, cfg.DB_PORT)

	if err != nil {
		log.Panic(err.Error())
	}

	if !wallet.Exists(cfg.Label) {
		log.Panicf("Identity %s not found in wallet!", cfg.Label)
	}

	clientConnection := newGrpcConnection(cfg)
	// Remember to close connection after use!

	id, err := wallet.Get(cfg.Label)

	if err != nil {
		log.Panic(err.Error())
	}

	x509id, err := id.ToX509Identity()

	if err != nil {
		log.Panic(err.Error())
	}

	sign, err := id.ToSign()

	if err != nil {
		log.Panic(err.Error())
	}

	// Create a Gateway connection for a specific client identity
	gw, err := client.Connect(
		x509id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		// Default timeouts for different gRPC calls
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)

	utils.CheckError(err)

	return gw
}

func LoadCertificate(path string) (*x509.Certificate, error) {
	certbytes, err := os.ReadFile(filepath.Clean(path))

	utils.CheckError(err)

	cert, err := identity.CertificateFromPEM(certbytes)

	utils.CheckError(err)

	return cert, err
}

// newGrpcConnection creates a gRPC connection to the Gateway server.
func newGrpcConnection(cfg config.Config) *grpc.ClientConn {
	certificate, err := LoadCertificate(cfg.CertPath)
	utils.CheckError(err)
	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	transportCredentials := credentials.NewClientTLSFromCert(certPool, cfg.GatewayPeer)

	connection, err := grpc.Dial(cfg.PeerEndpoint, grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		panic(fmt.Errorf("failed to create gRPC connection: %w", err))
	}

	return connection
}

// Get a channel instance from a connected gateway using channel name
func GetNetwork(gw client.Gateway, channelName string) *client.Network {
	network := gw.GetNetwork(channelName)
	return network
}

// Get a contract instance from a connected channel using chaincode name and contract name
func GetContract(network client.Network, chaincodeName string, contractName string) *client.Contract {
	contract := network.GetContractWithName(chaincodeName, contractName)
	return contract
}

var Usage = func() {
	var CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	fmt.Fprintf(CommandLine.Output(), "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}

func parseArgs(label *string, ccpPath *string, mspid *string, certPath *string, keystorePath *string) {
	flag.StringVar(mspid, "mspid", "", "The MSP ID managing the identity")
	flag.StringVar(label, "label", "", "A unique name to store and get the identity from the wallet")
	flag.StringVar(ccpPath, "ccpPath", "", "The path to the associated connection profile")
	flag.StringVar(certPath, "certPath", "", "The path to the signed identity certificate issued by the CA and MSP")
	flag.StringVar(keystorePath, "keystorePath", "", "The path to the identity's private key issued by the CA and MSP")
}

func main() {
	// Some required
	var label, ccpPath, mspid, certPath, keystorePath string
	var isEnv = flag.Bool("env", true, "Use environment variables to get the required values")
	var isNew = flag.Bool("new", false, "Put the new identity into the wallet by label, this argument will not be parsed from env vars")
	var isHelp = flag.Bool("h", false, "Print help message")

	parseArgs(&label, &ccpPath, &mspid, &certPath, &keystorePath)

	flag.Parse()

	if *isHelp {
		Usage()
		os.Exit(0)
	}

	var cfg config.Config
	if *isEnv {
		cfg.ParseEnv()
	} else {

		cfg = config.Config{
			MspID:        mspid,
			CertPath:     certPath,
			KeystorePath: keystorePath,
			CCPPath:      ccpPath,
			Label:        label,
		}
	}

	if *isNew {
		file, err := os.Lstat(keystorePath)
		utils.CheckError(err)

		// If the given path is a directory, take the first file
		// (The generated filename is random each time, so giving the directory is easier for testing)
		if file.IsDir() {
			files, err := os.ReadDir(cfg.KeystorePath)

			utils.CheckError(err)

			keystorePath = cfg.KeystorePath + files[0].Name()
		}

		_, err = NewIdentityFromFile(mspid, certPath)
	}

	gw := GetGateway(cfg)
	fmt.Printf("Connected to the gateway %+v\n", gw)
}
