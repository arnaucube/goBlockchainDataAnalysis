package main

import (
	"fmt"
	"strconv"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
)

func explore(client *rpcclient.Client, blockHash string) {
	var realBlocks int
	/*var nOrigin NodeModel
	nOrigin.Id = "origin"
	nOrigin.Label = "origin"
	nOrigin.Title = "origin"
	nOrigin.Group = "origin"
	nOrigin.Value = 1
	nOrigin.Shape = "dot"
	saveNode(nodeCollection, nOrigin)*/

	for blockHash != "" {
		//generate hash from string
		bh, err := chainhash.NewHashFromStr(blockHash)
		check(err)
		block, err := client.GetBlockVerbose(bh)
		check(err)
		if block.Height > config.StartFromBlock {
			var newBlock BlockModel
			newBlock.Hash = block.Hash
			newBlock.Confirmations = block.Confirmations
			newBlock.Size = block.Size
			newBlock.Height = block.Height
			//newBlock.Amount = block.Amount
			//newBlock.Fee = block.Fee
			newBlock.PreviousHash = block.PreviousHash
			newBlock.NextHash = block.NextHash
			newBlock.Time = block.Time
			newBlock.DateT = unixTimeToTime(block.Time)
			newBlock.Date.Year, newBlock.Date.Month, newBlock.Date.Day, newBlock.Date.Hour = decomposeDate(block.Time)

			//stats blocks
			stats := getStats()
			stats.BlockCount++
			if len(block.Tx) > 1 {
				stats.RealBlockCount++
			}
			updateStats(stats)

			for k, txHash := range block.Tx {
				if k > 0 {
					realBlocks++

					//stats txs
					stats := getStats()
					stats.TxCount++
					updateStats(stats)

					fmt.Print("Block Height: ")
					fmt.Print(block.Height)
					fmt.Print(", num of Tx: ")
					fmt.Print(k)
					fmt.Print("/")
					fmt.Println(len(block.Tx) - 1)

					th, err := chainhash.NewHashFromStr(txHash)
					check(err)
					blockTx, err := client.GetRawTransactionVerbose(th)
					check(err)

					//save Tx Node
					var nodeTx NodeModel
					nodeTx.Id = txHash
					nodeTx.Label = txHash
					nodeTx.Title = txHash
					nodeTx.Group = strconv.FormatInt(block.Height, 10)
					nodeTx.Value = 1
					nodeTx.Shape = "square"
					nodeTx.Type = "tx"
					saveNode(nodeCollection, nodeTx)

					var originAddresses []string
					var outputAddresses []string
					var outputAmount []float64
					for _, Vo := range blockTx.Vout {
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

							//Address
							var addr AddressModel
							addr.Hash = outputAddr
							addr.InBittrex = false
							saveAddress(addr)
						}
					}
					for _, Vi := range blockTx.Vin {
						th, err := chainhash.NewHashFromStr(Vi.Txid)
						check(err)
						txVi, err := client.GetRawTransactionVerbose(th)
						check(err)
						if len(txVi.Vout[Vi.Vout].ScriptPubKey.Addresses) > 0 {
							//add tx to newBlock
							newBlock.Tx = append(newBlock.Tx, blockTx.Txid)

							//Tx save
							var newTx TxModel
							newTx.Hex = blockTx.Hex
							newTx.Txid = blockTx.Txid
							newTx.Hash = blockTx.Hash
							newTx.BlockHash = block.Hash
							newTx.BlockHeight = strconv.FormatInt(block.Height, 10)
							newTx.Time = blockTx.Time
							newTx.DateT = unixTimeToTime(block.Time)
							newTx.Date.Year, newTx.Date.Month, newTx.Date.Day, newTx.Date.Hour = decomposeDate(block.Time)
							for _, originAddr := range txVi.Vout[Vi.Vout].ScriptPubKey.Addresses {
								originAddresses = append(originAddresses, originAddr)

								newTx.From = originAddr

								var n1 NodeModel
								n1.Id = originAddr
								n1.Label = originAddr
								n1.Title = originAddr
								n1.Group = strconv.FormatInt(block.Height, 10)
								n1.Value = 1
								n1.Shape = "dot"
								n1.Type = "address"
								saveNode(nodeCollection, n1)

								//Address
								var addr AddressModel
								addr.Hash = originAddr
								addr.InBittrex = false
								saveAddress(addr)

								for k, outputAddr := range outputAddresses {
									var eIn EdgeModel
									eIn.From = originAddr
									eIn.To = txHash
									eIn.Label = txVi.Vout[Vi.Vout].Value
									eIn.Txid = blockTx.Txid
									eIn.Arrows = "to"
									eIn.BlockHeight = block.Height
									saveEdge(edgeCollection, eIn)

									var eOut EdgeModel
									eOut.From = txHash
									eOut.To = outputAddr
									eOut.Label = outputAmount[k]
									eOut.Txid = blockTx.Txid
									eOut.Arrows = "to"
									eOut.BlockHeight = block.Height
									saveEdge(edgeCollection, eOut)

									//date analysis
									//dateAnalysis(e, tx.Time)
									//hour analysis
									hourAnalysis(eIn, blockTx.Time)

									newTx.To = outputAddr

								}
							}
							saveTx(newTx)
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
			saveBlock(newBlock)
		}
		//set the next block
		blockHash = block.NextHash
	}

	fmt.Print("realBlocks added (blocks with Fee and Amount values): ")
	fmt.Println(realBlocks)
	fmt.Println("reached the end of blockchain")
}
