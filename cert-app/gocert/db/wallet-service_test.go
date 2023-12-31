package db

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"gocert-gateway/utils"
)

func TestPostgreConnection(t *testing.T) {
	cfg, err := utils.SetupEnv(ENV_FILE)
	if err != nil {
		t.Fatal(err)
	}

	b := BaseDBService{}
	err = b.Connect(cfg.DB_HOST, cfg.DB_USER, cfg.DB_PASS, cfg.DB_NAME, cfg.DB_PORT)
	if err != nil {
		t.Error(err)
		t.Fatal("Fail to connect to the database")
	}
}

func TestPutIdentity(t *testing.T) {
	cfg, err := utils.SetupEnv(ENV_FILE)
	if err != nil {
		t.Fatal(err)
	}

	b := BaseDBService{}
	err = b.Connect(cfg.DB_HOST, cfg.DB_USER, cfg.DB_PASS, cfg.DB_NAME, cfg.DB_PORT)

	if err != nil {
		t.Error(err)
		t.Fatal("Fail to connect to the database")
	}

	p := NewWalletService(&b)

	label := "user1"
	mspid := "Org1MSP"
	eCert, privKey, err := getKeys(PK_PATH, SK_PATH)
	if err != nil {
		t.Error(err)
		t.Fatal("Fail to retrieve keys")
	}

	err = p.Put(label, mspid, eCert, privKey)
	if err != nil {
		t.Error(err)
		t.Fatal("Fail to put identity into wallet")
	}
}

func TestGetIdentity(t *testing.T) {
	cfg, err := utils.SetupEnv(ENV_FILE)
	if err != nil {
		t.Fatal(err)
	}

	b := BaseDBService{}
	err = b.Connect(cfg.DB_HOST, cfg.DB_USER, cfg.DB_PASS, cfg.DB_NAME, cfg.DB_PORT)

	if err != nil {
		t.Error(err)
		t.Fatal("Fail to connect to the database")
	}

	p := NewWalletService(&b)

	label := "user1"
	expectedMSPID := "Org1MSP"
	expectedECert, expectedPrivKey, err := getKeys(PK_PATH, SK_PATH)
	if err != nil {
		t.Error(err)
		t.Fatal("Fail to retrieve keys")
	}

	i, err := p.Get(label)

	if err != nil {
		t.Error(err)
		t.Fatal("Fail to get users")
	}

	if label != i.Label {
		t.Fatalf("Label does not match!\nExpected: %s\nReceived: %s", label, i.Label)
	}

	if expectedMSPID != i.MSPID {
		t.Fatalf("MSPID does not match!\nExpected: %s\nReceived: %s", expectedMSPID, i.MSPID)
	}

	if !bytes.Equal(expectedECert, i.PublicKey) {
		t.Fatalf("Public key does not match!\nExpected: %s\nReceived: %s", expectedECert, i.PublicKey)
	}

	if !bytes.Equal(expectedPrivKey, i.PrivateKey) {
		t.Fatalf("Private key does not match!\nExpected: %s\nReceived: %s", &expectedPrivKey, i.PrivateKey)
	}
}

func getKeys(pkPath string, skPath string) ([]byte, []byte, error) {
	certbytes, err := os.ReadFile(filepath.Clean(pkPath))

	if err != nil {
		return nil, nil, err
	}

	keybytes, err := os.ReadFile(filepath.Clean(skPath))

	if err != nil {
		return nil, nil, err
	}

	return certbytes, keybytes, nil
}
