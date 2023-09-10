package service

import (
	"gocert-gateway/db"
	"gocert-gateway/utils"
	"testing"

	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
)

func TestCreateNewIdentityService(t *testing.T) {
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

	ws := db.NewWalletService(&b)
	indiS := db.NewIndividualService(&b)

	NewIdentityService(*ws, *indiS)
}

func TestSaveNewIdentityFromEnrollment(t *testing.T) {
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

	ws := db.NewWalletService(&b)
	indiS := db.NewIndividualService(&b)

	idS := NewIdentityService(*ws, *indiS)

	individuals := db.LoadIndividualsFromFile(MOCK_USER_PATH)

	// Mock SignIdentity created from Enrollment
	regServ, err := NewRegistrationService(CCP_CONFIG_FILE, ORG_NAME, CA_ID)

	if err != nil {
		t.Error(err)
		t.Fatal("Error creating registration service")
	}

	pbKey, pvKey := db.LoadKeys(cfg.CertPath, cfg.KeystorePath)

	signId, err := regServ.caclient.CreateSigningIdentity(msp.WithCert(pbKey), msp.WithPrivateKey(pvKey))

	if err != nil {
		t.Error(err)
		t.Fatal("Error creating signing identity")
	}

	err = idS.CreateAccount(&individuals[0], signId)

	if err != nil {
		t.Error(err)
		t.Fatal("Error creating new account with individual and identity")
	}
}

func TestSaveNewIdentityFromExistingKeys(t *testing.T) {
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

	ws := db.NewWalletService(&b)
	indiS := db.NewIndividualService(&b)

	is := NewIdentityService(*ws, *indiS)

	CERT_PEM := `-----BEGIN CERTIFICATE-----
MIICoDCCAkagAwIBAgIUWppxpretymrcBnFh8ylfvKqujZAwCgYIKoZIzj0EAwIw
cDELMAkGA1UEBhMCVVMxFzAVBgNVBAgTDk5vcnRoIENhcm9saW5hMQ8wDQYDVQQH
EwZEdXJoYW0xGTAXBgNVBAoTEG9yZzEuZXhhbXBsZS5jb20xHDAaBgNVBAMTE2Nh
Lm9yZzEuZXhhbXBsZS5jb20wHhcNMjMwOTA3MDk1MTAwWhcNMjQwOTA2MDk1NjAw
WjBdMQswCQYDVQQGEwJVUzEXMBUGA1UECBMOTm9ydGggQ2Fyb2xpbmExFDASBgNV
BAoTC0h5cGVybGVkZ2VyMQ8wDQYDVQQLEwZjbGllbnQxDjAMBgNVBAMTBXVzZXIx
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE6M9SsKn7UywAKkAoAw5M9um2Du/O
omt6DB/AKK8eTwkohGsiSnhG966OOcM2wH1jhbrAIVlP5pxAUDFCc1VYEqOB0DCB
zTAOBgNVHQ8BAf8EBAMCB4AwDAYDVR0TAQH/BAIwADAdBgNVHQ4EFgQU0STPiR5p
rgaLLZFhSxuGguTczmwwHwYDVR0jBBgwFoAUpLguGfyo3Q7+W7x5omF76SKMEaUw
EwYDVR0RBAwwCoIIU2xlZXB5VFQwWAYIKgMEBQYHCAEETHsiYXR0cnMiOnsiaGYu
QWZmaWxpYXRpb24iOiIiLCJoZi5FbnJvbGxtZW50SUQiOiJ1c2VyMSIsImhmLlR5
cGUiOiJjbGllbnQifX0wCgYIKoZIzj0EAwIDSAAwRQIhAMpvdEb4iUVs2+yLfsSk
NSA5QvyQlFLipWBiJwC54okDAiAktL2Mt18MgXkGpCrWNTskMlaUQfwNIyVI3s1h
ELxD+g==
-----END CERTIFICATE-----`

	KEY_PEM := `-----BEGIN PRIVATE KEY-----
MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgvxPVWTe7Fn3XwqZl
PGGhpqpNFQR/Zllm7AwvS4xrvkKhRANCAAToz1KwqftTLAAqQCgDDkz26bYO786i
a3oMH8Aorx5PCSiEayJKeEb3ro45wzbAfWOFusAhWU/mnEBQMUJzVVgS
-----END PRIVATE KEY-----`

	cert_bytes := []byte(CERT_PEM)
	key_bytes := []byte(KEY_PEM)
	mspid := "Org1MSP"
	username := "User1"

	err = is.CreateAccountFromExistingKeys(mspid, username, cert_bytes, key_bytes, nil)

	if err != nil {
		t.Fatal(err)
	}

}
