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
	Txid        string  `json:"txid"`
	From        string  `json:"from"`
	To          string  `json:"to"`
	Label       float64 `json:"label"` //the value of transaction
	Arrows      string  `json:"arrows"`
	BlockHeight int64   `json:"blockheight"`
}
type NetworkModel struct {
	Nodes []NodeModel `json:"nodes"`
	Edges []EdgeModel `json:"edges"`
}

type SankeyNodeModel struct {
	//StringNode string `json:"stringnode"`
	Node int    `json:"node"`
	Name string `json:"name"`
}
type SankeyLinkModel struct {
	//StringSource string  `json:"stringsource"`
	Source int `json:"source"`
	//StringTarget string  `json:"stringtarget"`
	Target int     `json:"target"`
	Value  float64 `json:"value"`
}
type SankeyModel struct {
	Nodes []SankeyNodeModel `json:"nodes"`
	Links []SankeyLinkModel `json:"links"`
}

type HourCountModel struct {
	Hour  string `json:"hour"`
	Count int    `json:"count"`
}
type HourAnalysisResp struct {
	Labels []string `json:"labels"`
	Data   []int    `json:"data"`
}
type DateCountModel struct {
	Time  int64  `json:"time"`
	Date  string `json:"date"`
	Count int    `json:"count"`
}
