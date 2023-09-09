package db

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"gocert-gateway/models"
	"gocert-gateway/utils"
)

func TestIndividualServiceConnection(t *testing.T) {
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

	NewIndividualService(&b)
}

func TestCreateIndividual(t *testing.T) {
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

	p := NewIndividualService(&b)

	individuals := loadIndividualFromFile(MOCK_USER_PATH)
	i := individuals[0]
	err = p.CreateIndividual(i.OrgId, i.CimPersonId, i.CimOrganisationId, i.Country, i.FirstName, i.LastName, i.PictureUrl, time.Time(i.DateOfBirth), i.OrgName, &i.Wallet)

	if err != nil {
		t.Error(err)
		t.Fatal("Fail to create individual!")
	}
}

func TestFindIndividualById(t *testing.T) {
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

	p := NewIndividualService(&b)

	expectedId := uint64(1)
	expectedWalletId := uint64(1)
	expectedCimPersonId := "c8303ccd-6138-4588-b7c9-d0db89028da1"
	if err != nil {
		t.Error(err)
		t.Fatal("Fail to retrieve keys")
	}

	i, err := p.FindIndividualById(expectedId)

	if err != nil {
		t.Error(err)
		t.Fatal("Fail to get users")
	}

	if expectedId != i.Id {
		t.Fatalf("Id does not match!\nExpected: %d\nReceived: %d", expectedId, i.Id)
	}

	if expectedCimPersonId != i.CimPersonId {
		t.Fatalf("CimPersonId does not match!\nExpected: %s\nReceived: %s", expectedCimPersonId, i.CimPersonId)
	}

	if expectedWalletId != i.Wallet.Id {
		t.Fatalf("WalletId does not match!\nExpected: %d\nReceived: %d", expectedWalletId, i.Wallet.Id)
	}
}

func TestInit(t *testing.T) {
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

	p := NewIndividualService(&b)

	individuals := loadIndividualFromFile(MOCK_USER_PATH)
	for _, i := range individuals {
		err = p.CreateIndividual(i.OrgId, i.CimPersonId, i.CimOrganisationId, i.Country, i.FirstName, i.LastName, i.PictureUrl, time.Time(i.DateOfBirth), i.OrgName, &i.Wallet)

		if err != nil {
			t.Error(err)
			t.Fatal("Fail to create individual!")
		}
	}
}

func loadIndividualFromFile(path string) []models.Individual {
	individuals := &[]models.Individual{}

	data, err := os.ReadFile(filepath.Clean(path))

	if err != nil {
		panic("Failed to read mock data file")
	}

	err = json.Unmarshal(data, individuals)

	if err != nil {
		fmt.Errorf("Failed to parse json!")
		panic(err)
	}

	cfg, err := utils.SetupEnv(ENV_FILE)

	if err != nil {
		fmt.Errorf("Error setting up environment variables")
		panic(err)
	}

	cert, pvKey := loadKeys(cfg.CertPath, cfg.KeystorePath)
	CRYPTO_PATH := "../../../test-network/organizations/peerOrganizations/org2.example.com"
	CERT_PATH2 := filepath.Join(CRYPTO_PATH, "/users/User1@org2.example.com/msp/signcerts/cert.pem")
	KEYSTORE_PATH2 := filepath.Join(CRYPTO_PATH, "/users/User1@org2.example.com/msp/keystore")
	cert2, pvKey2 := loadKeys(CERT_PATH2, KEYSTORE_PATH2)
	(&(*individuals)[0]).Wallet.PublicKey = cert
	(&(*individuals)[0]).Wallet.PrivateKey = pvKey
	(&(*individuals)[1]).Wallet.PublicKey = cert2
	(&(*individuals)[1]).Wallet.PrivateKey = pvKey2

	return *individuals
}

func loadKeys(pbPath string, pvPath string) ([]byte, []byte) {
	cert, err := os.ReadFile(filepath.Join(filepath.Clean(pbPath)))

	if err != nil {
		fmt.Errorf("Error reading cert %s", pbPath)
		panic(err)
	}

	pvKeyDir, err := os.ReadDir(filepath.Join(filepath.Clean(pvPath)))

	if err != nil {
		fmt.Errorf("Error reading keystore directory %s", pvPath)
		panic(err)
	}

	file_name := pvKeyDir[0].Name()
	pvKey, err := os.ReadFile(filepath.Join(filepath.Clean(pvPath), file_name))

	if err != nil {
		fmt.Errorf("Error reading keystore file %s", file_name)
		panic(err)
	}

	return cert, pvKey
}

// func TestMarshalByte(t *testing.T) {
// 	cfg, _ := utils.SetupEnv(ENV_FILE)

// 	cert, pvKey := loadKeys(cfg.CertPath, cfg.KeystorePath)

// 	cert2json, _ := json.Marshal(cert)
// 	pvKey2json, _ := json.Marshal(pvKey)

// 	os.WriteFile("certbytes", cert2json, 0644)
// 	os.WriteFile("pvKeybytes", pvKey2json, 0644)

// }
