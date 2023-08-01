package tasks

import (
	"log"
	"time"

	"example.com/refcode/v1/app/models"
	"example.com/refcode/v1/app/models/contract"
	"example.com/refcode/v1/pkg/configs"
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
			log.Println("WITHDRAW EVENTS")
			log.Printf("TO ADDRESS %s\n", vlog.Account.Hex())
			log.Printf("AMOUNT %f\n", float64(amount.Int64())/float64(blockchain.EtherConst.Int64()))
			log.Printf("DOMAIN %s\n", vlog.Domain)
			history := models.WithdrawHistory{
				Address: vlog.Account.Hex(),
				Amount:  float64(amount.Int64()) / float64(blockchain.EtherConst.Int64()),
				Created: time.Now().UTC(),
				Domain:  vlog.Domain,
			}
			err = history.Save()
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
