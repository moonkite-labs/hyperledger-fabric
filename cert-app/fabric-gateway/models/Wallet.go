package models

import (
	"time"

	"github.com/pkg/errors"

	"github.com/hyperledger/fabric-gateway/pkg/identity"
)

type Tabler interface {
	TableName() string
}

type Wallet struct {
	Id         uint64 `gorm:"primaryKey;auto_increment"`
	Label      string `gorm:"column:cim_person_id"`
	MSPID      string
	PublicKey  []byte
	PrivateKey []byte
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}

// TableName overrides the table name used by Wallet to `wallet`
func (Wallet) TableName() string {
	return "wallet"
}

// Create an X509 certificate object using public key
func (i *Wallet) ToX509Identity() (*identity.X509Identity, error) {
	cert, err := identity.CertificateFromPEM(i.PublicKey)

	if err != nil {
		return nil, errors.Wrap(err, "Fail to parse certificate from wallet's identity")
	}

	return identity.NewX509Identity(i.MSPID, cert)
}

// Create a signer object using private key
func (i *Wallet) ToSign() (identity.Sign, error) {
	privateKey, err := identity.PrivateKeyFromPEM(i.PrivateKey)
	if err != nil {
		return nil, errors.Wrap(err, "Fail to parse private key from wallet's identity")
	}

	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		return nil, errors.Wrap(err, "Fail to create private key sign")
	}

	return sign, nil
}
