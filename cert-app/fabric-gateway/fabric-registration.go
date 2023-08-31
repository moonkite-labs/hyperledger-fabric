package main

import (
	"fmt"
	"path"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

func GetCAClientFromOrg(tlsConfigPath string) (*msp.Client, error) {
	tlsConfigPath = path.Clean(tlsConfigPath)
	sdk, err := fabsdk.New(config.FromFile(tlsConfigPath))

	CheckError(err)

	ctx := sdk.Context()

	c, err := msp.New(ctx) // Create MSP client

	CheckError(err)

	return c, nil
}

func Register(tlsConfigPath string, mspid string, org string, username string, idtype string, secret string) (string, error) {
	CAClient, err := GetCAClientFromOrg(tlsConfigPath)

	if err != nil {
		return secret, fmt.Errorf("%s", err)
	}

	CAInfo, err := CAClient.GetCAInfo()

	CheckError(err)

	reg := msp.RegistrationRequest{
		Name:   username,
		Type:   idtype,
		Secret: secret,
		CAName: CAInfo.CAName,
	}

	return CAClient.Register(&reg)
}
