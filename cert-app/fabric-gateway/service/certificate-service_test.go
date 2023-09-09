package service

import (
	"encoding/json"
	"fmt"
	"gocert-gateway/db"
	"gocert-gateway/models"
	"gocert-gateway/utils"
	"os"
	"path/filepath"
	"testing"
)

func TestIssueCertificate(t *testing.T) {
	issuerId := "c8303ccd-6138-4588-b7c9-d0db89028da1"
	recipientId := "52ef8e0b-cef9-4045-8f58-9361b82ef576"
	certificates := loadCertificateFromFile(MOCK_CERT_PATH)

	cfg, err := utils.SetupEnv(ENV_FILE)
	if err != nil {
		t.Fatal(err)
	}

	b := db.BaseDBService{}
	err = b.Connect(cfg.DB_HOST, cfg.DB_USER, cfg.DB_PASS, cfg.DB_NAME, cfg.DB_PORT)

	if err != nil {
		t.Error(err)
		t.Fatal("Fail to connect to the database")
	}

	indiService := db.NewIndividualService(&b)
	walletService := db.NewWalletService(&b)
	contractService := ContractService{}

	cs := NewCertificateService(&b, *indiService, *walletService, contractService)
	err = cs.IssueCertificate(issuerId, recipientId, &certificates[0], &certificates[1])
	if err != nil {
		t.Error(err)
		t.Fatal("Failed to issue certificate")
	}
}

func loadCertificateFromFile(path string) []models.Certificate {
	certificates := &[]models.Certificate{}

	data, err := os.ReadFile(filepath.Clean(path))

	if err != nil {
		panic("Failed to read mock data file")
	}

	err = json.Unmarshal(data, certificates)

	if err != nil {
		fmt.Errorf("Failed to parse json!")
		panic(err)
	}

	return *certificates
}
