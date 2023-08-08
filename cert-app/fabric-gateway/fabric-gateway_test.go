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

	if len(cfg.Identity.MspID) == 0 {
		t.Fatalf("MSPID is empty!")
	}

	if len(cfg.Identity.CertPath) == 0 {
		t.Fatalf("CertPath is empty!")
	}

	if len(cfg.Identity.KeystorePath) == 0 {
		t.Fatalf("KeystorePath is empty!")
	}

	if len(cfg.Identity.WalletPath) == 0 {
		t.Fatalf("WalletPath is empty!")
	}

	if len(cfg.Identity.CCPPath) == 0 {
		t.Fatalf("CCPPath is empty!")
	}
}

func TestPutIdentity(t *testing.T) {

	err := godotenv.Load(".env.test")

	if err != nil {
		t.Fatalf(err.Error())
	}

	var cfg Config
	ParseEnv(&cfg)

	wallet := CreateWallet(cfg.Identity.WalletPath)

	if wallet == nil {
		t.Fatalf("Failed to create wallet at %s", cfg.Identity.WalletPath)
	}

	files, err := os.ReadDir(cfg.Identity.KeystorePath)

	if err != nil {
		t.Fatalf("Error reading from %s\nErr: %s", cfg.Identity.KeystorePath, err)
	}

	keyStorePath := cfg.Identity.KeystorePath + files[0].Name()

	identity := NewIdentityFromFile(MSP_ID, CERT_PATH, keyStorePath)

	if identity == nil {
		t.Fatalf("Identity failed to be created from cert path: %s\n keystore path: %s", CERT_PATH, keyStorePath)
	}

	label := "User1@org1"

	id, err := wallet.Get(label)

	if id != nil {
		t.Fatalf("Identity %s already exists in wallet!", label)
	}

	err = wallet.Put(label, identity)

	if err != nil {
		t.Fatalf("Fail to put identity %s to wallet!", label)
	}

	t.Cleanup(func() { os.RemoveAll(WALLET_PATH) })
}
