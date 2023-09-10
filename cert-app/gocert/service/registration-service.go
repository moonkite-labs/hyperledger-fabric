package service

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/pkg/errors"

	mspctx "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

func printError(err error, msg string) error {
	if err != nil {
		if msg != "" {
			return fmt.Errorf(msg)
		} else {
			return fmt.Errorf(err.Error())
		}
	}
	return nil
}

type RegistrationService struct {
	sdk      *fabsdk.FabricSDK
	caclient *msp.Client
}

func NewRegistrationService(config_path string, orgname string, caid string) (*RegistrationService, error) {
	sdk, err := fabsdk.New(config.FromFile(config_path))
	if err != nil {
		return nil, errors.Wrap(err, "Fail to create sdk")
	}
	caclient, err := createCACalient(sdk, msp.WithCAInstance(caid), msp.WithOrg(orgname))
	if err != nil {
		return nil, errors.Wrap(err, "Fail to create sdk")
	}
	return &RegistrationService{sdk: sdk, caclient: caclient}, nil
}

func createCACalient(sdk *fabsdk.FabricSDK, options ...msp.ClientOption) (*msp.Client, error) {

	ctx := sdk.Context()

	c, err := msp.New(ctx, options...) // Create MSP client

	if err != nil {
		printError(err, "")
	}

	return c, nil
}

func (r RegistrationService) RegisterAndEnroll(username string, idtype string, secret string, attrs ...msp.Attribute) (mspctx.SigningIdentity, error) {

	CAInfo, err := r.caclient.GetCAInfo()

	if err != nil {
		printError(err, "")
	}

	reg := msp.RegistrationRequest{
		Name:       username,
		Type:       idtype,
		Secret:     secret,
		CAName:     CAInfo.CAName,
		Attributes: attrs,
	}

	secret, err = r.caclient.Register(&reg)

	if err != nil {
		return nil, err
	}

	err = r.caclient.Enroll(username, msp.WithSecret(secret))

	if err != nil {
		return nil, err
	}

	return r.caclient.GetSigningIdentity(username)
}
