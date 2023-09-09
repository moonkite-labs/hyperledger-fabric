package service

import (
	"fmt"
	"testing"

	"gocert-gateway/utils"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

const (
	ORG_NAME = "Org1"
	CA_ID    = "ca-org1"
)

func TestCreateFabSDK(t *testing.T) {
	sdk, err := fabsdk.New(config.FromFile(CCP_CONFIG_FILE))
	if err != nil {
		t.Fatal("Fail to create sdk")
	}
	defer sdk.Close()
}

func TestCreateRegistrationService(t *testing.T) {
	cfg, err := utils.SetupEnv(ENV_FILE)

	if err != nil {
		t.Fatal("Failed to setup environment")
	}

	NewRegistrationService(*cfg, ORG_NAME, CA_ID)
}

func TestEnrollAndRegister(t *testing.T) {
	cfg, err := utils.SetupEnv(ENV_FILE)

	if err != nil {
		t.Fatal("Failed to setup environment")
	}

	r := NewRegistrationService(*cfg, ORG_NAME, CA_ID)

	if err != nil {
		t.Error(err)
		t.Fatal("Fail to create ca client")
	}

	username := "user2"
	idtype := "client"
	secret := "user2pw"

	identity, err := r.RegisterAndEnroll(username, idtype, secret)

	if err != nil {
		t.Error(err)
		t.Fatalf("Fail to register and enroll user %s ca client\n", username)
	}

	// for i := 0; i < len(resp); i++ {
	fmt.Printf("%+v\n\n", identity)
	// }
}
