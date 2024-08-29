package main

import (
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	// apiKey              = "" // api key to etherscan
	numberOfDaysToCheck = 4
	JsonFilename        = "abi's.json"

	rpcAddress      = "wss://ethereum-rpc.publicnode.com"          //rpc address (mainnet)
	etherscanAPIURL = "api.etherscan.io/api"                       //etherscan api
	stringAddress   = "0xdAC17F958D2ee523a2206206994597C13D831ec7" // address of smart contract as string (tether)
	eventName       = "Transfer"
	nodeOperatorId  = -1

	// rpcAddress = "wss://ethereum-holesky-rpc.publicnode.com" //rpc address (holesky)
	// etherscanAPIURL     = "api-holesky.etherscan.io/api" //etherscan api for holesky
	// eventName           = "ValidatorExitRequest"
	// stringAddress       = "0xffddf7025410412deaa05e3e1ce68fe53208afcb" // address of smart contract as string (ExitRequests Lido in Holesky)
	// nodeOperatorId      = 22

	// rpcAddress = "wss://ethereum-holesky-rpc.publicnode.com" //rpc address (holesky)
	// etherscanAPIURL     = "api-holesky.etherscan.io/api" //etherscan api for holesky
	// eventName      = "ELRewardsStealingPenaltyReported"
	// stringAddress  = "0x4562c3e63c2e586cD1651B958C22F88135aCAd4f" // address of smart contract as string (ELRewardsStealingPenaltyReported  Lido in Holesky)
	// nodeOperatorId = 37
)

func main() {
	client, err := ethclient.Dial(rpcAddress)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	go monitorEvents(client, stringAddress)

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":2112", nil))
}
