package main

import (
	"fabric-gateway/config"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/joho/godotenv"
)

var WALLET_PATH = filepath.Clean("./wallet/test")
var MSP_ID = "Org1MSP"
var CRYPTO_PATH = filepath.Clean("../../test-network/organizations/peerOrganizations/")
var CERT_PATH = filepath.Join(CRYPTO_PATH, "org1.example.com/users/User1@org1.example.com/msp/signcerts/cert.pem")
var KEYSTORE_DIR_PATH = filepath.Join(CRYPTO_PATH, "org1.example.com/users/User1@org1.example.com/msp/keystore")

// Test identity creation from file, using mspid Org1MSP
func TestCreateIdentity(t *testing.T) {

	files, err := os.ReadDir(KEYSTORE_DIR_PATH)

	if err != nil {
		t.Fatalf("Error reading from %s\nErr: %s", KEYSTORE_DIR_PATH, err)
	}

	keyStorePath := KEYSTORE_DIR_PATH + files[0].Name()

	identity, err := NewIdentityFromFile(MSP_ID, CERT_PATH)

	if identity == nil {
		t.Fatalf("Identity failed to be created from cert path: %s\n keystore path: %s", CERT_PATH, keyStorePath)
	}
}

func TestConfigParser(t *testing.T) {
	err := godotenv.Load(".env.test")

	if err != nil {
		t.Fatalf(err.Error())
	}

	var cfg config.Config
	cfg.ParseEnv()

	expectedMspID := "Org1MSP"
	expectedLabel := "User1"
	expectedCertPath := "../../test-network/organizations/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/signcerts/cert.pem"
	expectedKeystorePath := "../../test-network/organizations/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/keystore"
	expectedCCPPath := "./client-config/config.yaml"

	var checkMatches = func(name string, expected any, got any) {
		if got != expected {
			t.Fatalf("%s does not match the expected value: (expected: %s, got: %s)", name, expected, got)
		}
	}

	checkMatches("MspID", cfg.MspID, expectedMspID)
	checkMatches("Label", cfg.Label, expectedLabel)
	checkMatches("CertPath", cfg.CertPath, expectedCertPath)
	checkMatches("KeystorePath", cfg.KeystorePath, expectedKeystorePath)
	checkMatches("CCPPath", cfg.CCPPath, expectedCCPPath)
}

func TestGetIdentity(t *testing.T) {
	err := godotenv.Load(".env.test")

	if err != nil {
		t.Fatalf(err.Error())
	}

	var cfg config.Config
	cfg.ParseEnv()

	x509cert, err := LoadCertificate(cfg.CertPath)
	if err != nil {
		t.Fatalf(err.Error())
	}

	fmt.Printf("%+v\n", x509cert)
}
