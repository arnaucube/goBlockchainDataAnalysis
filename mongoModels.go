package main

type AddressModel struct {
	Hash      string  `json:"hash"`
	Amount    float64 `json:"amount"`
	InBittrex bool    `json:"inbittrex"`
}
type DateModel struct {
	Hour        string  `json:"hour"`
	Day         string  `json:"day"`
	Month       string  `json:"month"`
	Amount      float64 `json:"amount"`
	BlockHash   string  `json:"blockhash"`
	BlockHeight string  `json:"blockheight"`
}
type TxModel struct {
	Hex         string  `json:"hex"`
	Txid        string  `json:"txid"`
	Hash        string  `json:"hash"`
	From        string  `json:"from"` //hash of address
	To          string  `json:"to"`   //hash of address
	Amount      float64 `json:"amount"`
	BlockHash   string  `json:"blockhash"`
	BlockHeight string  `json:"blockheight"`
	Time        int64   `json:"time"`
	DateF       string  `json:"datef"` //date formated
	Date        DateModel
}
type BlockModel struct {
	Hash          string `json:"hash"`
	Confirmations uint64 `json:"confirmations"`
	Size          int32  `json:"size"`
	Height        int64  `json:"height"`
	//Amount        float64  `json:"amount"`
	//Fee           float64  `json:"fee"`
	Tx           []string `json:"txid"` //txid of the TxModel
	PreviousHash string   `json:"previoushash"`
	NextHash     string   `json:"nexthash"`
	Time         int64    `json:"time"`
	DateF        string   `json:"datef"` //date formated
	Date         DateModel
}

type NodeModel struct {
	Id    string `json:"id"`
	Label string `json:"label"`
	Title string `json:"title"`
	Group string `json:"group"`
	Value int    `json:"value"`
	Shape string `json:"shape"`
	Type  string `json:"type"`
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
