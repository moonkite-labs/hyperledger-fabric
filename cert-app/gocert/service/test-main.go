package service

import (
	"path/filepath"
)

const (
	ENV_FILE            = "../.env.test"
	TEST_DATA_ROOT_PATH = "../test_data"
	CCP_CONFIG_FILE     = "../client-config/config.yaml"

	CHANNEL_NAME   = "certapp"
	CHAINCODE_NAME = "certificate-manager"
	CONTRACT_NAME  = "CertificateContract"
)

var (
	PK_PATH        = filepath.Join(TEST_DATA_ROOT_PATH, "msp", "signcerts", "User1@org1.example.com-cert.pem")
	SK_PATH        = filepath.Join(TEST_DATA_ROOT_PATH, "msp", "keystore", "priv_sk")
	MOCK_DATA_PATH = filepath.Join(TEST_DATA_ROOT_PATH, "mock_data")
	MOCK_USER_PATH = filepath.Join(MOCK_DATA_PATH, "mock-user.json")
	MOCK_CERT_PATH = filepath.Join(MOCK_DATA_PATH, "mock-certificate.json")
)
