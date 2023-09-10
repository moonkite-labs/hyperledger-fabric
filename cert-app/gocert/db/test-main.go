package db

import (
	"encoding/json"
	"fmt"
	"gocert-gateway/models"
	"gocert-gateway/utils"
	"os"
	"path/filepath"
)

const (
	ENV_FILE            = "../.env.test"
	TEST_DATA_ROOT_PATH = "../test_data"
)

var (
	PK_PATH        = filepath.Join(TEST_DATA_ROOT_PATH, "msp", "signcerts", "User1@org1.example.com-cert.pem")
	SK_PATH        = filepath.Join(TEST_DATA_ROOT_PATH, "msp", "keystore", "priv_sk")
	MOCK_DATA_PATH = filepath.Join(TEST_DATA_ROOT_PATH, "mock_data")
	MOCK_USER_PATH = filepath.Join(MOCK_DATA_PATH, "mock-user.json")
	MOCK_CERT_PATH = filepath.Join(MOCK_DATA_PATH, "mock-certificate.json")
)

func LoadIndividualsFromFile(path string) []models.Individual {
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

	cert, pvKey := LoadKeys(cfg.CertPath, cfg.KeystorePath)
	CRYPTO_PATH := "../../../test-network/organizations/peerOrganizations/org2.example.com"
	CERT_PATH2 := filepath.Join(CRYPTO_PATH, "/users/User1@org2.example.com/msp/signcerts/cert.pem")
	KEYSTORE_PATH2 := filepath.Join(CRYPTO_PATH, "/users/User1@org2.example.com/msp/keystore")
	cert2, pvKey2 := LoadKeys(CERT_PATH2, KEYSTORE_PATH2)
	(&(*individuals)[0]).Wallet.PublicKey = cert
	(&(*individuals)[0]).Wallet.PrivateKey = pvKey
	(&(*individuals)[1]).Wallet.PublicKey = cert2
	(&(*individuals)[1]).Wallet.PrivateKey = pvKey2

	return *individuals
}

func LoadKeys(pbPath string, pvPath string) ([]byte, []byte) {
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
