package smartcontract

import (
	"log"

	"github.com/Ashikurrahaman287/Blonet/internal/blockchain"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Deployer struct {
	ethereum *blockchain.Ethereum
	bsc      *blockchain.BSC
}

func NewDeployer(ethereum *blockchain.Ethereum, bsc *blockchain.BSC) *Deployer {
	return &Deployer{
		ethereum: ethereum,
		bsc:      bsc,
	}
}

func (d *Deployer) DeployOnEthereum(auth *bind.TransactOpts, bytecode []byte) (common.Address, *types.Transaction, error) {
	address, tx, err := d.ethereum.DeploySmartContract(auth, bytecode)
	if err != nil {
		log.Printf("Failed to deploy contract on Ethereum: %v", err)
		return common.Address{}, nil, err
	}
	log.Printf("Contract deployed on Ethereum at address: %s", address.Hex())
	return address, tx, nil
}

func (d *Deployer) DeployOnBSC(auth *bind.TransactOpts, bytecode []byte) (common.Address, *types.Transaction, error) {
	address, tx, err := d.bsc.DeploySmartContract(auth, bytecode)
	if err != nil {
		log.Printf("Failed to deploy contract on BSC: %v", err)
		return common.Address{}, nil, err
	}
	log.Printf("Contract deployed on BSC at address: %s", address.Hex())
	return address, tx, nil
}
