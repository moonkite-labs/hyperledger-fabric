package identity

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	mspctx "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
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

func CreateCACalient(sdk fabsdk.FabricSDK, options ...msp.ClientOption) (*msp.Client, error) {

	ctx := sdk.Context()

	c, err := msp.New(ctx, options...) // Create MSP client

	if err != nil {
		printError(err, "")
	}

	return c, nil
}

func RegisterAndEnroll(caclient *msp.Client, username string, idtype string, secret string, attrs ...msp.Attribute) (mspctx.SigningIdentity, error) {

	CAInfo, err := caclient.GetCAInfo()

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

	secret, err = caclient.Register(&reg)

	if err != nil {
		return nil, err
	}

	err = caclient.Enroll(username, msp.WithSecret(secret))

	if err != nil {
		return nil, err
	}

	return caclient.GetSigningIdentity(username)

}
