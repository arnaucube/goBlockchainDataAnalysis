package main

type BlockModel struct {
	Hash          string
	Height        int64
	Confirmations uint64
	Amount        float64
	Fee           float64
}
