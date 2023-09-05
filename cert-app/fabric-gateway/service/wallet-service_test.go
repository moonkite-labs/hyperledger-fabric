package service

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

const (
	HOSTNAME            = "localhost"
	PORT                = "5432"
	USER                = "postgres"
	PASSWORD            = "postgres"
	DBNAME              = "postgres"
	TEST_DATA_ROOT_PATH = "./test_data"
)

var (
	PK_PATH = filepath.Join(TEST_DATA_ROOT_PATH, "msp", "signcerts", "User1@org1.example.com-cert.pem")
	SK_PATH = filepath.Join(TEST_DATA_ROOT_PATH, "msp", "keystore", "priv_sk")
)

func TestPostgreConnection(t *testing.T) {
	p := PostgreWalletService{}
	err := p.Connect(HOSTNAME, USER, PASSWORD, DBNAME, PORT)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPutIdentity(t *testing.T) {
	p := PostgreWalletService{}
	err := p.Connect(HOSTNAME, USER, PASSWORD, DBNAME, PORT)
	if err != nil {
		t.Error(err)
		t.Fatal("Fail to connect to the database")
	}

	label := "user1"
	eCert, privKey, err := getKeys(PK_PATH, SK_PATH)
	if err != nil {
		t.Error(err)
		t.Fatal("Fail to retrieve keys")
	}

	err = p.Put(label, eCert, privKey)
	if err != nil {
		t.Error(err)
		t.Fatal("Fail to put identity into wallet")
	}
}

func TestGetIdentity(t *testing.T) {
	p := PostgreWalletService{}
	err := p.Connect(HOSTNAME, USER, PASSWORD, DBNAME, PORT)
	if err != nil {
		t.Error(err)
		t.Fatal("Fail to connect to the database")
	}

	label := "user1"
	expectedECert, expectedPrivKey, err := getKeys(PK_PATH, SK_PATH)
	if err != nil {
		t.Error(err)
		t.Fatal("Fail to retrieve keys")
	}

	i := p.Get(label)

	if label != i.Label {
		t.Fatalf("Label does not match!\nExpected: %s\nReceived: %s", label, i.Label)
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
