package service

import (
	"fmt"
	"gocert-gateway/db"

	mspctx "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
)

// Manages storage of a user's cryptographic identity
type IdentityService struct {
	wallet db.PostgreWalletService
}

func NewIdentityService(wallet db.PostgreWalletService) *IdentityService {
	return &IdentityService{wallet: wallet}
}

func (is IdentityService) CreateAccount(signId mspctx.SigningIdentity) error {

	label := createLabel(signId.Identifier().MSPID, signId.Identifier().ID)

	pvkey, err := signId.PrivateKey().Bytes()

	if err != nil {
		return err
	}

	return is.wallet.Put(label, signId.Identifier().MSPID, signId.EnrollmentCertificate(), pvkey)
}

func (is IdentityService) CreateAccountFromExistingKeys(mspid string, username string, pubKey []byte, pvKey []byte) error {
	label := createLabel(mspid, username)
	return is.wallet.Put(label, mspid, pubKey, pvKey)
}

func createLabel(mspid string, username string) string {
	return fmt.Sprintf("%s::%s", mspid, username)
}
