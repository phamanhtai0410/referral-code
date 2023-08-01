package tasks

import (
	"log"

	"example.com/refcode/v1/app/models"
	"example.com/refcode/v1/app/models/contract"
	"example.com/refcode/v1/pkg/configs"
	"example.com/refcode/v1/pkg/workers"
	"example.com/refcode/v1/platform/blockchain"
	"github.com/ethereum/go-ethereum/common"
)

func ListenChainEvent() {
	address := common.HexToAddress(configs.ContractAddr)
	refCode, err := contract.NewRefCode(address, blockchain.GetClient())
	if err != nil {
		log.Fatal("WATCH WITHDRAW ERROR #1 ", err)
	}
	logs := make(chan *contract.RefCodeWithdrawn)
	sub, _ := refCode.WatchWithdrawn(nil, logs)
	for {
		select {
		case err = <-sub.Err():
			log.Fatal("WATCH WITHDRAW ERROR #2: ", err)
		case vlog := <-logs:
			amount := vlog.Amount
			err = workers.Delay(
				"Worker.SaveHistory",
				SaveHistory,
				vlog.Account.Hex(),
				vlog.Domain,
				float64(amount.Int64())/float64(blockchain.EtherConst.Int64()),
			)
			if err != nil {
				log.Fatal("WATCH WITHDRAW ERROR #3: ", err)
			}
			var user models.User
			user.Withdraw(vlog.Domain)
			if err != nil {
				log.Fatal("WATCH WITHDRAW ERROR #3: ", err)
			}
		}
	}
}
