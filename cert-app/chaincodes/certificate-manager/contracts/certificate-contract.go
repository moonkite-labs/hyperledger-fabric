package contracts

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/pkg/errors"

	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// CertificateContract provides functions for managing a Certificate
type CertificateContract struct {
	contractapi.Contract
}

func (s *CertificateContract) GetName() string {
	return "CertificateContract"
}

func (s *CertificateContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	Certs := []Certificate{}

	for _, Cert := range Certs {
		CertJSON, err := json.Marshal(Cert)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(Cert.CertificateId, CertJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}

// id string, course string, recipientId string, eventName string, recipientEmail string, recipientName string, individualPersonId string, individualPublicKey []byte, certificateType CertificateType, templateName string, customTemplateUrl string, issuerId string, issuerName string, issuedDate time.Time, expiryDate time.Time, createdDate time.Time, issuedPersonId string, issuerPublicKey []byte, uid string, description string, certificateName string, ownerId string, ownerName string, ownerEmail string, sharedFrom string, masterId string, typeOfCopy CopyType, isValid bool, revokedReason string, badgeUrl string, templateType string, certificateDesc string, status string, templateJson string
func (s *CertificateContract) CreateCertificate(ctx contractapi.TransactionContextInterface, id string, recipientId string, recipientPublicKey []byte, issuerCertHash string, recipientCertHash string) error {
	exists, err := s.CertificateExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the certificate %s already exists", id)
	}

	orgid, err := ctx.GetClientIdentity().GetMSPID()

	if err != nil {
		return err
	}

	cert := Certificate{
		CertificateId:      id,
		OrganisationId:     orgid,
		RecipientId:        recipientId,
		RecipientPublicKey: recipientPublicKey,
		IssuerCertHash:     issuerCertHash,
		RecipientCertHash:  recipientCertHash,
	}
	certJSON, err := json.Marshal(cert)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, certJSON)
}

// IssueCertificate issue a certificate to a receiver
// func (s *CertificateContract) IssueCertificate(ctx contractapi.TransactionContextInterface, id string, receiverId string) error {
// 	// Check if the caller has the issuer role
// 	isIssuer := s.isIssuer(ctx)

// 	if !isIssuer {
// 		return errors.New("Caller is not authorised as an issuer!")
// 	}

// 	// Get the requester's instance
// 	requester, err := cid.New(ctx.GetStub())

// 	if err != nil {
// 		return errors.Wrap(err, "Failed to get the requester's instance")
// 	}

// 	// Get the requester's public key certificate
// 	pubKeyCert, err := requester.GetX509Certificate()

// 	if err != nil {
// 		return errors.Wrap(err, "Failed to get the requester's public key certificate")
// 	}

// 	// Get the requester's organisation
// 	requesterOrg := pubKeyCert.Subject.Organization

// }

// CertificateExists returns true when asset with given ID exists in world state
func (s *CertificateContract) CertificateExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	certificateJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, errors.Wrap(err, "failed to read from world state")
	}

	return certificateJSON != nil, nil
}

func (s *CertificateContract) isIssuer(ctx contractapi.TransactionContextInterface) bool {

	// cid is the invoker of the chaincode
	hasIssuerOU, err := cid.HasOUValue(ctx.GetStub(), "hf.CertIssuer")
	if err != nil {
		log.Fatalf("Failed to check for function invoker's issuer identity.\n")
	}
	hasIssuerAttr := ctx.GetClientIdentity().AssertAttributeValue("Issuer", "true") != nil
	return hasIssuerOU || hasIssuerAttr
}

func containsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
