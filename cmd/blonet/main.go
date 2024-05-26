package main

import (
	"log"

	"github.com/Ashikurrahaman287/Blonet/internal/blockchain"
	"github.com/Ashikurrahaman287/Blonet/internal/smartcontract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	ethClient, err := blockchain.NewEthereum("https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID")
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum: %v", err)
	}

	bscClient, err := blockchain.NewBSC("https://bsc-dataseed.binance.org/")
	if err != nil {
		log.Fatalf("Failed to connect to BSC: %v", err)
	}

	deployer := smartcontract.NewDeployer(ethClient, bscClient)

	privateKey, err := crypto.HexToECDSA("your-private-key-hex")
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	bytecode := []byte("your-smart-contract-bytecode")

	// Deploy on Ethereum
	ethAddress, ethTx, err := deployer.DeployOnEthereum(auth, bytecode)
	if err != nil {
		log.Fatalf("Failed to deploy contract on Ethereum: %v", err)
	}
	log.Printf("Contract deployed on Ethereum at address: %s", ethAddress.Hex())
	log.Printf("Ethereum transaction hash: %s", ethTx.Hash().Hex())

	// Deploy on BSC
	bscAddress, bscTx, err := deployer.DeployOnBSC(auth, bytecode)
	if err != nil {
		log.Fatalf("Failed to deploy contract on BSC: %v", err)
	}
	log.Printf("Contract deployed on BSC at address: %s", bscAddress.Hex())
	log.Printf("BSC transaction hash: %s", bscTx.Hash().Hex())
}
