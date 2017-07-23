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

		var newBlock BlockModel
		newBlock.Hash = block.Hash
		newBlock.Height = block.Height
		newBlock.Confirmations = block.Confirmations

		//get Fee value
		th, err := chainhash.NewHashFromStr(block.Tx[0])
		check(err)
		tx, err := client.GetRawTransactionVerbose(th)
		check(err)
		var totalFee float64
		for _, Vo := range tx.Vout {
			totalFee = totalFee + Vo.Value
		}
		newBlock.Fee = totalFee

		//for each Tx, get the Tx value
		var totalAmount float64
		/*inside each block, there are []Tx
		inside each Tx, if is the Tx[0], is the Fees
		in the Tx[n] where n>0, each Tx is independent,
		and inside each Tx there are []Vout.
		Usually Vout[0] is the real Tx value
		and the Vout[1] is the rest of the amount in the original wallet.
		Each Tx moves all the wallet amount, and the realTx amount is sent to the destination
		and the rest of the wallet amount, is send to another owned wallet
		*/
		for k, txHash := range block.Tx {
			//first Tx is the Fee
			//after first Tx is the Sent Amount
			if k > 0 {
				th, err := chainhash.NewHashFromStr(txHash)
				check(err)
				tx, err := client.GetRawTransactionVerbose(th)
				check(err)
				var originAddress string
				for _, Vi := range tx.Vin {
					th, err := chainhash.NewHashFromStr(Vi.Txid)
					check(err)
					txVi, err := client.GetRawTransactionVerbose(th)
					check(err)
					if len(txVi.Vout[0].ScriptPubKey.Addresses) > 0 {
						originAddress = txVi.Vout[0].ScriptPubKey.Addresses[0]
					}
				}
				for _, Vo := range tx.Vout {
					totalAmount = totalAmount + Vo.Value

					var blockTx TxModel
					blockTx.Txid = tx.Txid
					blockTx.Amount = Vo.Value
					blockTx.From = originAddress
					blockTx.To = Vo.ScriptPubKey.Addresses[0]
					newBlock.Tx = append(newBlock.Tx, blockTx)
				}
			}
		}

		if totalAmount > 0 {
			newBlock.Amount = totalAmount
			saveBlock(blockCollection, newBlock)
			fmt.Print("Height: ")
			fmt.Println(newBlock.Height)
			fmt.Print("Amount: ")
			fmt.Println(newBlock.Amount)
			fmt.Print("Fee: ")
			fmt.Println(newBlock.Fee)
			fmt.Println("-----")
			realBlocks++
		}

		//set the next block
		blockHash = block.NextHash

		//analyze block creator

	}
	fmt.Print("realBlocks (blocks with Fee and Amount values): ")
	fmt.Println(realBlocks)
	fmt.Println("reached the end of blockchain")
}
