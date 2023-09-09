package service

import (
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
)

type ContractService struct {
	gw      *client.Gateway
	network *client.Network
}

// TODO: ClaimRequest?
func (cs *ContractService) IssueCertificate(chaincodeName string, contractName string, issuerPersonId string, issuerOrgSigner identity.Sign, masterCertificateId string, issuerHash string, recipientHash string, recipientId string, recipientPublicKey []byte, recipientCertId string) (string, error) {

	// contract := cs.network.GetContractWithName(chaincodeName, contractName)

	// CreateCertificate args: id string, recipientId string, recipientPublicKey []byte, issuerCertHash string, recipientCertHash string
	// txnProposal, err := contract.NewProposal("CreateCertificate", client.WithArguments([]string{masterCertificateId, recipientId}...), client.WithBytesArguments(recipientPublicKey), client.WithArguments(issuerHash, recipientHash))
	// if err != nil {
	// 	errors.Wrap(err, "Error creating new proposal to create certificate.")
	// 	return "", err
	// }
	// txnEndorsed, err := txnProposal.Endorse()
	// if err != nil {
	// 	errors.Wrap(err, "Error endorsing txn.")
	// 	return "", err
	// }
	// txnCommitted, err := txnEndorsed.Submit()
	// if err != nil {
	// 	errors.Wrap(err, "Error submitting transaction:")
	// 	return "", err
	// }

	// return txnCommitted.TransactionID(), nil
	return "Called", nil
}
