package identity

import (
	"fmt"
	"testing"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

var TEST_FILE = "../test-data/config.yaml"

func TestCreateFabSDK(t *testing.T) {
	sdk, err := fabsdk.New(config.FromFile(TEST_FILE))
	if err != nil {
		t.Fatal("Fail to create sdk")
	}
	defer sdk.Close()
}

func TestCreateCAClient(t *testing.T) {
	sdk, err := fabsdk.New(config.FromFile(TEST_FILE))
	if err != nil {
		t.Fatal("Fail to create sdk")
	}
	defer sdk.Close()

	caclient, err := CreateCACalient(*sdk, msp.WithCAInstance("ca-org1"))

	if err != nil {
		t.Error(err)
		t.Fatal("Fail to create ca client")
	}

	creds, err := caclient.GetSigningIdentity("user1")

	if err != nil {
		t.Error(err)
		t.Fatal("Fail to get signing identity")
	}

	fmt.Printf("%+v\n", creds)

}

func TestEnrollAndRegister(t *testing.T) {
	sdk, err := fabsdk.New(config.FromFile(TEST_FILE))
	if err != nil {
		t.Error(err)
		t.Fatal("Fail to create sdk")
	}
	defer sdk.Close()

	caclient, err := CreateCACalient(*sdk)

	if err != nil {
		t.Error(err)
		t.Fatal("Fail to create ca client")
	}

	username := "user1"
	idtype := "client"
	secret := "user1pw"

	identity, err := RegisterAndEnroll(caclient, username, idtype, secret)

	if err != nil {
		t.Error(err)
		t.Fatalf("Fail to register and enroll user %s ca client\n", username)
	}

	// for i := 0; i < len(resp); i++ {
	fmt.Printf("%+v\n\n", identity)
	// }
}
