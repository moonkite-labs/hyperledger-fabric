package service

import (
	"crypto/x509"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"github.com/pkg/errors"
)

type ContractService struct {
	gw *client.Gateway
}

// TODO: ClaimRequest?
func (cs ContractService) IssueCertificate(contract *client.Contract, issuerPersonId string, issuerIdentity identity.X509Identity, issuerOrgSigner identity.Sign, masterCertificateId string, issuerHash string, recipientHash string, recipientPublicKey *x509.Certificate, recipientCertId string) (string, error) {

	recipientId := recipientPublicKey.Subject.SerialNumber

	rPubBytes, err := identity.CertificateToPEM(recipientPublicKey)

	// CreateCertificate args: id string, recipientId string, recipientPublicKey []byte, issuerCertHash string, recipientCertHash string
	txnProposal, err := contract.NewProposal("CreateCertificate", client.WithArguments([]string{masterCertificateId, recipientId}...), client.WithBytesArguments(rPubBytes), client.WithArguments(issuerHash, recipientHash))
	if err != nil {
		errors.Wrap(err, "Error creating new proposal to create certificate.")
		return "", err
	}
	txnEndorsed, err := txnProposal.Endorse()
	if err != nil {
		errors.Wrap(err, "Error endorsing txn.")
		return "", err
	}
	txnCommitted, err := txnEndorsed.Submit()
	if err != nil {
		errors.Wrap(err, "Error submitting transaction:")
		return "", err
	}

	return txnCommitted.TransactionID(), nil
}
