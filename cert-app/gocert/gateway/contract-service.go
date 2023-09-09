package gateway

import (
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"github.com/pkg/errors"
)

type CertContractService struct {
	gw       *client.Gateway
	network  *client.Network
	contract *client.Contract
}

func NewCertContractService(gw *client.Gateway, channelName string, chaincodeName string, contractName string) *CertContractService {
	network := gw.GetNetwork(channelName)
	return &CertContractService{gw: gw, network: network, contract: network.GetContractWithName(chaincodeName, contractName)}
}

// TODO: ClaimRequest?
func (cs *CertContractService) IssueCertificate(issuerPersonId string, issuerOrgSigner identity.Sign, masterCertificateId string, issuerHash string, recipientHash string, recipientId string, recipientPublicKey []byte, recipientCertId string) (string, error) {

	// CreateCertificate args: id string, recipientId string, recipientPublicKey []byte, issuerCertHash string, recipientCertHash string
	txnProposal, err := cs.contract.NewProposal("CreateCertificate", client.WithArguments([]string{masterCertificateId, recipientId}...), client.WithBytesArguments(recipientPublicKey), client.WithArguments(issuerHash, recipientHash))
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
	// return "Called", nil
}
