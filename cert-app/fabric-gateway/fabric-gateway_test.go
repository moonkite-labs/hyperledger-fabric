package main

import (
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

// Test wallet creation at the specified path
func TestCreateWallet(t *testing.T) {
	wallet := CreateWallet(WALLET_PATH)
	if wallet == nil {
		t.Fatalf("Failed to create wallet at path %s", WALLET_PATH)
	}
	// Clean up
	t.Cleanup(func() { os.Remove(WALLET_PATH) })
}

// Test identity creation from file, using mspid Org1MSP
func TestCreateIdentity(t *testing.T) {

	files, err := os.ReadDir(KEYSTORE_DIR_PATH)

	if err != nil {
		t.Fatalf("Error reading from %s\nErr: %s", KEYSTORE_DIR_PATH, err)
	}

	keyStorePath := KEYSTORE_DIR_PATH + files[0].Name()

	identity := NewIdentityFromFile(MSP_ID, CERT_PATH, keyStorePath)

	if identity == nil {
		t.Fatalf("Identity failed to be created from cert path: %s\n keystore path: %s", CERT_PATH, keyStorePath)
	}
}

func TestConfigParser(t *testing.T) {
	err := godotenv.Load(".env.test")

	if err != nil {
		t.Fatalf(err.Error())
	}

	var cfg Config
	ParseEnv(&cfg)

	expectedMspID := "Org1MSP"
	expectedLabel := "User1"
	expectedCertPath := "../../test-network/organizations/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/signcerts/cert.pem"
	expectedKeystorePath := "../../test-network/organizations/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/keystore"
	expectedWalletPath := "./wallet/test"
	expectedCCPPath := "../../test-network/organizations/peerOrganizations/org1.example.com/connection-org1.json"

	var checkMatches = func(name string, expected any, got any) {
		if got != expected {
			t.Fatalf("%s does not match the expected value: (expected: %s, got: %s)", name, expected, got)
		}
	}

	checkMatches("MspID", cfg.MspID, expectedMspID)
	checkMatches("Label", cfg.Label, expectedLabel)
	checkMatches("CertPath", cfg.CertPath, expectedCertPath)
	checkMatches("KeystorePath", cfg.KeystorePath, expectedKeystorePath)
	checkMatches("WalletPath", cfg.WalletPath, expectedWalletPath)
	checkMatches("CCPPath", cfg.CCPPath, expectedCCPPath)
}

func TestPutIdentity(t *testing.T) {

	err := godotenv.Load(".env.test")

	if err != nil {
		t.Fatalf(err.Error())
	}

	var cfg Config
	ParseEnv(&cfg)

	wallet := CreateWallet(cfg.WalletPath)

	if wallet == nil {
		t.Fatalf("Failed to create wallet at %s", cfg.WalletPath)
	}

	files, err := os.ReadDir(cfg.KeystorePath)

	if err != nil {
		t.Fatalf("Error reading from %s\nErr: %s", cfg.KeystorePath, err)
	}

	keyStorePath := cfg.KeystorePath + files[0].Name()

	identity := NewIdentityFromFile(MSP_ID, CERT_PATH, keyStorePath)

	if identity == nil {
		t.Fatalf("Identity failed to be created from cert path: %s\n keystore path: %s", CERT_PATH, keyStorePath)
	}

	id, _ := wallet.Get(cfg.Label)

	if id != nil {
		t.Fatalf("Identity %s already exists in wallet!", cfg.Label)
	}

	err = wallet.Put(cfg.Label, identity)

	if err != nil {
		t.Fatalf("Fail to put identity %s to wallet!", cfg.Label)
	}

	t.Cleanup(func() { os.RemoveAll(WALLET_PATH) })
}
