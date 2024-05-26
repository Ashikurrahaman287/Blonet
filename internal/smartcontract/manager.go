package smartcontract

import (
	"github.com/Ashikurrahaman287/Blonet/internal/blockchain"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Manager struct {
	ethereum *blockchain.Ethereum
	bsc      *blockchain.BSC
}

func NewManager(ethereum *blockchain.Ethereum, bsc *blockchain.BSC) *Manager {
	return &Manager{
		ethereum: ethereum,
		bsc:      bsc,
	}
}

func (m *Manager) DeployOnEthereum(auth *bind.TransactOpts, bytecode []byte) (common.Address, *types.Transaction, error) {
	return m.ethereum.DeploySmartContract(auth, bytecode)
}

func (m *Manager) DeployOnBSC(auth *bind.TransactOpts, bytecode []byte) (common.Address, *types.Transaction, error) {
	return m.bsc.DeploySmartContract(auth, bytecode)
}

func (m *Manager) InteractWithEthereumContract(contractAddress common.Address, method string, params ...interface{}) (interface{}, error) {
	return m.ethereum.InteractWithContract(contractAddress, method, params...)
}

func (m *Manager) InteractWithBSCContract(contractAddress common.Address, method string, params ...interface{}) (interface{}, error) {
	return m.bsc.InteractWithContract(contractAddress, method, params...)
}
