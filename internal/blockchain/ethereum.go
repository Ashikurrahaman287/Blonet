package blockchain

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

type Ethereum struct {
	client *ethclient.Client
}

func NewEthereum(url string) (*Ethereum, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	return &Ethereum{client: client}, nil
}

func (e *Ethereum) DeploySmartContract(auth *bind.TransactOpts, bytecode []byte) (common.Address, error) {
	address, tx, _, err := bind.DeployContract(auth, abi, bytecode, e.client)
	if err != nil {
		return common.Address{}, err
	}
	log.Printf("Contract deployed! Wait for tx: %s", tx.Hash().Hex())
	return address, nil
}

func (e *Ethereum) InteractWithContract(contractAddress common.Address, method string, params ...interface{}) (interface{}, error) {
	// Interaction logic here
	return nil, nil
}
