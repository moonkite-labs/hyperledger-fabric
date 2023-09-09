package gateway

import (
	"crypto/x509"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	db "gocert-gateway/db"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"gocert-gateway/config"
	"gocert-gateway/utils"
)

// Create a new X509Identity from associated mspid, issued certificate file and private key file
func NewIdentityFromFile(mspid string, certpath string) (*identity.X509Identity, error) {

	cert, err := LoadCertificate(certpath)

	utils.CheckError(err)

	return identity.NewX509Identity(mspid, cert)
}

// Return a gateway connected using the given identity
func GetGateway(cfg config.Config) *client.Gateway {
	baseDb := db.BaseDBService{}
	err := baseDb.Connect(cfg.DB_HOST, cfg.DB_USER, cfg.DB_PASS, cfg.DB_NAME, cfg.DB_PORT)

	wallet := db.NewWalletService(&baseDb)

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
	certificate, err := LoadCertificate(cfg.TlsCertPath)
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
