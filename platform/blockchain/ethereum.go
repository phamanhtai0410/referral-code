package blockchain

import (
	"log"
	"math/big"
	"sync"

	"example.com/refcode/v1/pkg/configs"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var cli *ethclient.Client = nil
var lock sync.Mutex

var EtherConst *big.Int

func init() {
	EtherConst = big.NewInt(1e18)
}

func NewTransactOpts(config *configs.TransactOptsConfig) *bind.TransactOpts {
	key, err := crypto.HexToECDSA(config.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(config.ChainID))
	if err != nil {
		log.Fatal(err)
	}
	return auth
}

func GetClient() *ethclient.Client {
	if cli == nil {
		lock.Lock()
		if cli == nil {
			var err error
			cli, err = ethclient.Dial(configs.BlockChainRPC)
			if err != nil {
				log.Fatal("CLIENT ERROR: ", err)
			}
			lock.Unlock()
			return cli
		}
		lock.Unlock()
	}
	return cli
}
