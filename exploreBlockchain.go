package main

import (
	"fmt"
	"strconv"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcrpcclient"
)

func explore(client *btcrpcclient.Client, blockHash string) {
	var realBlocks int
	var nOrigin NodeModel
	nOrigin.Id = "origin"
	nOrigin.Label = "origin"
	nOrigin.Title = "origin"
	nOrigin.Group = "origin"
	nOrigin.Value = 1
	nOrigin.Shape = "dot"
	saveNode(nodeCollection, nOrigin)

	for blockHash != "" {
		//generate hash from string
		bh, err := chainhash.NewHashFromStr(blockHash)
		check(err)
		block, err := client.GetBlockVerbose(bh)
		check(err)

		if block.Height > config.StartFromBlock {
			for k, txHash := range block.Tx {
				if k > 0 {
					realBlocks++
					fmt.Print("Block Height: ")
					fmt.Print(block.Height)
					fmt.Print(", num of Tx: ")
					fmt.Print(k)
					fmt.Print("/")
					fmt.Println(len(block.Tx) - 1)

					th, err := chainhash.NewHashFromStr(txHash)
					check(err)
					tx, err := client.GetRawTransactionVerbose(th)
					check(err)

					//save Tx
					var nTx NodeModel
					nTx.Id = txHash
					nTx.Label = txHash
					nTx.Title = txHash
					nTx.Group = strconv.FormatInt(block.Height, 10)
					nTx.Value = 1
					nTx.Shape = "square"
					nTx.Type = "tx"
					saveNode(nodeCollection, nTx)

					var originAddresses []string
					var outputAddresses []string
					var outputAmount []float64
					for _, Vo := range tx.Vout {
						//if Vo.Value > 0 {
						for _, outputAddr := range Vo.ScriptPubKey.Addresses {
							outputAddresses = append(outputAddresses, outputAddr)
							outputAmount = append(outputAmount, Vo.Value)
							var n2 NodeModel
							n2.Id = outputAddr
							n2.Label = outputAddr
							n2.Title = outputAddr
							n2.Group = strconv.FormatInt(block.Height, 10)
							n2.Value = 1
							n2.Shape = "dot"
							n2.Type = "address"
							saveNode(nodeCollection, n2)
						}
						//}
					}
					for _, Vi := range tx.Vin {
						th, err := chainhash.NewHashFromStr(Vi.Txid)
						check(err)
						txVi, err := client.GetRawTransactionVerbose(th)
						check(err)
						if len(txVi.Vout[Vi.Vout].ScriptPubKey.Addresses) > 0 {
							for _, originAddr := range txVi.Vout[Vi.Vout].ScriptPubKey.Addresses {
								originAddresses = append(originAddresses, originAddr)
								var n1 NodeModel
								n1.Id = originAddr
								n1.Label = originAddr
								n1.Title = originAddr
								n1.Group = strconv.FormatInt(block.Height, 10)
								n1.Value = 1
								n1.Shape = "dot"
								n1.Type = "address"
								saveNode(nodeCollection, n1)

								for k, outputAddr := range outputAddresses {
									var eIn EdgeModel
									eIn.From = originAddr
									eIn.To = txHash
									eIn.Label = txVi.Vout[Vi.Vout].Value
									eIn.Txid = tx.Txid
									eIn.Arrows = "to"
									eIn.BlockHeight = block.Height
									saveEdge(edgeCollection, eIn)

									var eOut EdgeModel
									eOut.From = txHash
									eOut.To = outputAddr
									eOut.Label = outputAmount[k]
									eOut.Txid = tx.Txid
									eOut.Arrows = "to"
									eOut.BlockHeight = block.Height
									saveEdge(edgeCollection, eOut)

									//date analysis
									//dateAnalysis(e, tx.Time)
									//hour analysis
									hourAnalysis(eIn, tx.Time)
								}
							}
						} else {
							originAddresses = append(originAddresses, "origin")
						}

					}
					fmt.Print("originAddresses: ")
					fmt.Println(len(originAddresses))
					fmt.Print("outputAddresses: ")
					fmt.Println(len(outputAddresses))
				}
			}
		}
		//set the next block
		blockHash = block.NextHash
	}

	fmt.Print("realBlocks (blocks with Fee and Amount values): ")
	fmt.Println(realBlocks)
	fmt.Println("reached the end of blockchain")
}
