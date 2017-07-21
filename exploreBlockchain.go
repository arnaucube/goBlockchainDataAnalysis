package main

import (
	"fmt"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcrpcclient"
)

func explore(client *btcrpcclient.Client, blockHash string) {
	var realBlocks int
	for blockHash != "" {
		//generate hash from string
		bh, err := chainhash.NewHashFromStr(blockHash)
		check(err)
		block, err := client.GetBlockVerbose(bh)
		check(err)

		th, err := chainhash.NewHashFromStr(block.Tx[0])
		check(err)
		tx, err := client.GetRawTransactionVerbose(th)
		check(err)

		var totalFee float64
		for _, Vo := range tx.Vout {
			totalFee = totalFee + Vo.Value
		}

		//for each Tx, get the Tx value
		var totalAmount float64
		for k, txHash := range block.Tx {
			//first Tx is the Fee
			//after first Tx is the Sent Amount
			if k > 0 {
				th, err := chainhash.NewHashFromStr(txHash)
				check(err)
				tx, err := client.GetRawTransactionVerbose(th)
				check(err)
				for _, Vo := range tx.Vout {
					totalAmount = totalAmount + Vo.Value
				}
			}
		}
		if totalAmount > 0 {
			var newBlock BlockModel
			newBlock.Hash = block.Hash
			newBlock.Height = block.Height
			newBlock.Confirmations = block.Confirmations
			newBlock.Amount = totalAmount
			newBlock.Fee = totalFee
			saveBlock(blockCollection, newBlock)
			fmt.Println(newBlock.Height)
			fmt.Println(newBlock.Amount)
			fmt.Println(newBlock.Fee)
			fmt.Println("-----")
			realBlocks++
		}

		//set the next block
		blockHash = block.NextHash
	}
	fmt.Print("realBlocks (blocks with Fee and Amount values): ")
	fmt.Println(realBlocks)
	fmt.Println("reached the end of blockchain")
}
