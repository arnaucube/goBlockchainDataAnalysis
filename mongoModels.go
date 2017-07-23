package main

type TxModel struct {
	Txid   string
	From   string
	To     string
	Amount float64
}
type BlockModel struct {
	Hash          string
	Height        int64
	Confirmations uint64
	Amount        float64
	Fee           float64
	Tx            []TxModel
}

type NodeModel struct {
	Id    string
	Label string
	Title string
	Group string
	Value int
	Shape string
}

type EdgeModel struct {
	From   string
	To     string
	Label  string
	Arrows string
}
