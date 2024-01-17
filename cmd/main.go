package main

import (
	"fmt"
	"log"
	"verilay/pkg/relays"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	conn, err := ethclient.Dial("https://rpc.holesky.ethpandaops.io")
	contractaddress := common.HexToAddress("0x018eb2438064E21B4074a9c411e64BA0316125eC")
	if err != nil {
		log.Fatalf("Failed to establish connection to Holesky-Testnet: %v", err)
	}

	contract, err := relays.NewRelaysCaller(contractaddress, conn)
	if err != nil {
		log.Fatalf("Failed to intantiate contract: %v", err)
	}

	opt := bind.CallOpts{Pending: false, Context: nil}
	valHead, _ := contract.GetFinalizedBlockRoot(&opt)

	fmt.Println(valHead)
}
