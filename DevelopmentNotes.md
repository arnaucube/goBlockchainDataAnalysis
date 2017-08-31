# Development Notes

## ToDo list

- Backend
    - Sankey Address generation without loops
- Frontend
    - After Sankey visualization, go to Network Address visualization and render without Sankey dots
- Both
    - Tx volume
    - Block size
    - Blockchain size

other

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

- mantain connection with wallet using websockets

- num address evolution throught time

- fix error in exploreBlockchain.go, when getting the tx.Vin
tx 8f04960da36beaa928b9693f7dca4afae5a6122bb6874d409a1156e4c6c55024 has 4 vin, but exploreBlockchain is only getting the first

- pagination in address network generation

- stop rendering dots of sankey, when view change

- sidebar pages:
    list of addresses in fairmarket (addresses of shops), to view statistics in time of the inputs and outputs in a timeline

- refresh blockchain database every minute
