# Development Notes

## ToDo list

- Backend
    - Network Address generation avoiding infinite relation loops
    - Sankey Address generation without loops
- Frontend
    - After Sankey visualization, go to Network Address visualization and render without Sankey dots
- Both
    - Tx/day
    - Tx volume
    - Block size
    - Blockchain size

other
- To get tx/hours of last 24 hours
    Search for TxModel with DateF > last24h
    Count for each hour
- To get tx/day of last month
    Search TxModel with DateF > last month
    Count each day
- Add counter with total blocks, total tx, total address

- store date hour, day, etc:
    ```Go
    type DateModel struct {
        Hour        string  `json:"hour"`
        Day         string  `json:"day"`
        Month       string  `json:"month"`
        Amount      float64 `json:"amount"`
        BlockHash   string  `json:"blockhash"`
        BlockHeight string  `json:"blockheight"`
    }
```
