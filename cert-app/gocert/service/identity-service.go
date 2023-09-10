package service

import (
	"gocert-gateway/db"
	"gocert-gateway/models"
	"time"

	mspctx "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
)

// Manages storage of a user's cryptographic identity
type IdentityService struct {
	wallet      db.PostgreWalletService
	indiService db.PostgreIndividualService
}

func NewIdentityService(wallet db.PostgreWalletService, indiService db.PostgreIndividualService) *IdentityService {
	return &IdentityService{wallet: wallet, indiService: indiService}
}

func (is IdentityService) CreateAccount(individual *models.Individual, signId mspctx.SigningIdentity) error {

	pvkey, err := signId.PrivateKey().Bytes()

	if err != nil {
		return err
	}

	wallet := &models.Wallet{
		Label:      individual.CimPersonId,
		MSPID:      signId.Identifier().MSPID,
		PublicKey:  signId.EnrollmentCertificate(),
		PrivateKey: pvkey,
	}

	return is.indiService.CreateIndividual(individual.OrgId, individual.CimPersonId, individual.CimOrganisationId, individual.Country, individual.FirstName, individual.LastName, individual.PictureUrl, time.Time(individual.DateOfBirth), individual.OrgName, wallet)
}

func (is IdentityService) CreateAccountFromExistingKeys(mspid string, username string, pubKey []byte, pvKey []byte, individual *models.Individual) error {
	label := individual.CimPersonId
	err := is.wallet.Put(label, mspid, pubKey, pvKey)
	if err != nil {
		return err
	}

	return is.indiService.CreateIndividual(individual.OrgId, individual.CimPersonId, individual.CimOrganisationId, individual.Country, individual.FirstName, individual.LastName, individual.PictureUrl, time.Time(individual.DateOfBirth), individual.OrgName, nil)
}

// func createLabel(mspid string, username string) string {
// 	return fmt.Sprintf("%s::%s", mspid, username)
// }
