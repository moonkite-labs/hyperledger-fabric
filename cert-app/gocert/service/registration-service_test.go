package service

import (
	"fmt"
	"path/filepath"
	"testing"

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
		t.Error(err)
		t.Fatal("Fail to create sdk")
	}
	defer sdk.Close()
}

func TestCreateRegistrationService(t *testing.T) {
	ccp_path := filepath.Join("../", filepath.Clean(CCP_CONFIG_FILE))

	_, err := NewRegistrationService(ccp_path, ORG_NAME, CA_ID)

	if err != nil {
		t.Error(err)
		t.Fatal("Failed to create a registration service")
	}
}

func TestEnrollAndRegister(t *testing.T) {
	ccp_path := filepath.Join("../", filepath.Clean(CCP_CONFIG_FILE))

	r, err := NewRegistrationService(ccp_path, ORG_NAME, CA_ID)

	if err != nil {
		t.Error(err)
		t.Fatal("Failed to create a registration service")
	}

	username := "user2"
	idtype := "client"
	secret := "user2pw"

	identity, err := r.RegisterAndEnroll(username, idtype, secret)

	if err != nil {
		t.Error(err)
		t.Fatalf("Failed to register and enroll user %s ca client\n", username)
	}

	// for i := 0; i < len(resp); i++ {
	fmt.Printf("%+v\n\n", identity)
	// }
}
