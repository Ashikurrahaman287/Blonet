package blockchain

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type BSC struct {
	client *ethclient.Client
}

func NewBSC(url string) (*BSC, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	return &BSC{client: client}, nil
}

func (b *BSC) DeploySmartContract(auth *bind.TransactOpts, bytecode []byte) (common.Address, *types.Transaction, error) {
	address, tx, _, err := bind.DeployContract(auth, bind.ContractABI{}, bytecode, b.client)
	if err != nil {
		return common.Address{}, nil, err
	}
	log.Printf("Contract deployed! Wait for tx: %s", tx.Hash().Hex())
	return address, tx, nil
}

func (b *BSC) InteractWithContract(contractAddress common.Address, method string, params ...interface{}) (interface{}, error) {
	// Implementation for interacting with a deployed contract
	// This will vary based on the contract's ABI and methods
	return nil, nil
}

func (b *BSC) GetBalance(address common.Address) (*big.Int, error) {
	balance, err := b.client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		return nil, err
	}
	return balance, nil
}
