package contracts

import (
	"encoding/json"
	"fmt"
	"crypto/x509"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// CertificateContract provides functions for managing a Certificate
type CertificateContract struct {
	contractapi.Contract
}

func (s *CertificateContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	Certs := []Certificate{}

	for _, Cert := range Certs {
		CertJSON, err := json.Marshal(Cert)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(Cert.Id, CertJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}

func (s *CertificateContract) IssueCertificate(ctx contractapi.TransactionContextInterface) error {

}

func (s *CertificateContract) isIssuer(ctx contractapi.TransactionContextInterface) (bool, error) {
	idcert, err := ctx.GetClientIdentity().GetX509Certificate()
	if err != nil {
		return false, err
	}
	ous := idcert.Issuer.OrganizationalUnit
	if containsString(ous, "Issuer")

}

func containsString(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}
