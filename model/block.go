package model

type (
	ProBlock struct {
		Shard_id           int    `json:"shard_id"`
		Shard_length       int    `json:"shard_length"`
		Previous_hash      string `json:"previous_hash"`
		Txblock_list       string `json:"txblock_list"`
		Globalstate_minus2 string `json:"globalstate_minus2"`

		Time_stamp int64   `json:"time_stamp"`
		Vrf        float64 `json:"vrf"`
		Public_key string  `json:"public_key"`
	}

	TxBlock struct {
		Shard_id     int      `json:"shard_id"`
		Shard_length int      `json:"shard_length"`
		Tx_list      []string `json:"tx_list"`

		Time_stamp int64  `json:"time_stamp"`
		Sig        string `json:"sig"`
		Public_key string `json:"public_key"`
	}

	WitnessTransactionsRequest struct {
		Id string `json:"id"`
	}
)
