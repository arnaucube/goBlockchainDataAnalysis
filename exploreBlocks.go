package main

import (
	"fmt"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcrpcclient"
)

func explore(client *btcrpcclient.Client, blockHash string) {
	for blockHash != "" {
		//generate hash from string
		bh, err := chainhash.NewHashFromStr(blockHash)
		check(err)
		fmt.Print("blockHash: ")
		fmt.Println(bh)
		block, err := client.GetBlockVerbose(bh)
		check(err)
		fmt.Print("height: ")
		fmt.Println(block.Height)
		fmt.Print("rawTx: ")
		fmt.Println(block.RawTx)
		fmt.Print("Tx: ")
		fmt.Println(block.Tx)
		fmt.Print("Time: ")
		fmt.Println(block.Time)
		fmt.Print("Confirmations: ")
		fmt.Println(block.Confirmations)

		fmt.Print("Fee: ")
		th, err := chainhash.NewHashFromStr(block.Tx[0])
		check(err)
		tx, err := client.GetRawTransactionVerbose(th)
		check(err)

		var totalFee float64
		for _, Vo := range tx.Vout {
			totalFee = totalFee + Vo.Value
		}
		fmt.Print("totalFee: ")
		fmt.Print(totalFee)
		fmt.Println(" FAIR")

		//for each Tx, get the Tx value
		var totalAmount float64
		for k, txHash := range block.Tx {
			//first Tx is the Fee
			//after first Tx is the Sent Amount
			if k > 0 {
				th, err := chainhash.NewHashFromStr(txHash)
				check(err)
				fmt.Print("tx hash: ")
				fmt.Println(th)
				tx, err := client.GetRawTransactionVerbose(th)
				check(err)
				for _, Vo := range tx.Vout {
					totalAmount = totalAmount + Vo.Value
				}
				fmt.Print("totalAmount: ")
				fmt.Print(totalAmount)
				fmt.Println(" FAIR")
			}
		}
		fmt.Println("-----")

		//set the next block
		blockHash = block.NextHash
	}
	fmt.Println("reached the end of blockchain")
}
