package individual

import (
	mspctx "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
	"github.com/pkg/errors"
)

type Individual struct {
	avatar          string   // An URL to the user's avatar
	name            string   // Username
	org             string   // Organisation of the user
	mspID           string   // MSPID of the user
	roles           []string // The roles
	signingIdentity mspctx.SigningIdentity
}

func NewIndividual(signingIdentity mspctx.SigningIdentity, opt ...IndividualOption) (*Individual, error) {
	i, err := initIndividualFromOptions(signingIdentity, opt...)
	if err != nil {
		return nil, errors.Wrap(err, "Fail to create user")
	}

	uerr := errors.New("Fail to create user")

	if signingIdentity == nil {
		return nil, errors.Wrap(uerr, "a SigningIdentity is required")
	}

	if i.name == "" {
		i.name = signingIdentity.Identifier().ID
	}

	if i.name == "" {
		return nil, errors.Wrap(uerr, "name is required!")
	}

	if i.org == "" {
		return nil, errors.Wrap(uerr, "org is required!")
	}

	if i.mspID == "" {
		i.mspID = signingIdentity.Identifier().MSPID
	}

	if i.mspID == "" {
		return nil, errors.Wrap(uerr, "MSPID is required!")
	}

	if i.roles == nil {
		i.roles = []string{}
	}

	return i, nil
}

func initIndividualFromOptions(signingIdentity mspctx.SigningIdentity, opts ...IndividualOption) (*Individual, error) {
	o := individualOptions{}
	for _, param := range opts {
		err := param(&o)
		if err != nil {
			return nil, errors.WithMessage(err, "failed to create Individual")
		}
	}
	c := Individual{
		avatar:          o.avatar,
		name:            o.name,
		org:             o.org,
		mspID:           o.mspID,
		roles:           o.roles,
		signingIdentity: signingIdentity,
	}

	return &c, nil
}

func (i Individual) GetName() string {
	return i.name
}

func (i Individual) GetAvatar() string {
	return i.avatar
}

func (i Individual) GetOrg() string {
	return i.org
}

func (i Individual) GetMspID() string {
	return i.mspID
}

func (i Individual) GetRoles() []string {
	return i.roles
}

func (i Individual) GetSigningIdentity() mspctx.SigningIdentity {
	return i.signingIdentity
}
