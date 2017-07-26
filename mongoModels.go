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
	Id    string `json:"id"`
	Label string `json:"label"`
	Title string `json:"title"`
	Group string `json:"group"`
	Value int    `json:"value"`
	Shape string `json:"shape"`
}

type EdgeModel struct {
	Txid   string  `json:"txid"`
	From   string  `json:"from"`
	To     string  `json:"to"`
	Label  float64 `json:"label"`
	Arrows string  `json:"arrows"`
}
type NetworkModel struct {
	Nodes []NodeModel `json:"nodes"`
	Edges []EdgeModel `json:"edges"`
}
