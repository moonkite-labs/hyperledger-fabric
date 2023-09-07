package main

import (
	"log"

	"certificates/contracts"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	contractChaincode, err := contractapi.NewChaincode(&contracts.CertificateContract{})
	if err != nil {
		log.Panicf("Error creating certificate-manager chaincode: %v", err)
	}

	if err := assetChaincode.Start(); err != nil {
		log.Panicf("Error starting certificate-manager chaincode: %v", err)
	}
}
