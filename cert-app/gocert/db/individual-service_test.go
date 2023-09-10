package db

import (
	"fmt"
	"testing"
	"time"

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

	individuals := LoadIndividualsFromFile(MOCK_USER_PATH)
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

	individuals := LoadIndividualsFromFile(MOCK_USER_PATH)
	for _, i := range individuals {
		err = p.CreateIndividual(i.OrgId, i.CimPersonId, i.CimOrganisationId, i.Country, i.FirstName, i.LastName, i.PictureUrl, time.Time(i.DateOfBirth), i.OrgName, &i.Wallet)

		if err != nil {
			t.Error(err)
			t.Fatal("Fail to create individual!")
		}
	}
}

func TestMarshalByte(t *testing.T) {
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

	individual1, err := p.FindIndividualById(uint64(1))

	if err != nil {
		t.Error(err)
		t.Fatal("Fail to find user with id 1")
	}

	individual2, err := p.FindIndividualById(uint64(2))

	if err != nil {
		t.Error(err)
		t.Fatal("Fail to find user with id 2")
	}

	addr1, err := individual1.Wallet.ToPublicAddress()

	if err != nil {
		t.Error(err)
		t.Fatal("Fail to parse individual 1's address")
	}

	addr2, err := individual2.Wallet.ToPublicAddress()

	if err != nil {
		t.Error(err)
		t.Fatal("Fail to parse individual 2's address")
	}

	fmt.Printf("Public address of individual 1: %s\n", addr1)

	fmt.Printf("Public address of individual 2: %s\n", addr2)
}
