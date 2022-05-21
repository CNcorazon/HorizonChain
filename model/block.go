package model

import (
	"horizon/structure"
)

type (
	BlockTransactionRequest struct {
		Shard uint
		Id    string
	}
	BlockTransactionResponse struct {
		Shard          uint
		Height         uint //需要生成的区块的高度，是当前区块链的高度+1
		Num            int
		InternalList   map[uint][]structure.InternalTransaction
		CrossShardList map[uint][]structure.CrossShardTransaction
		RelayList      map[uint][]structure.SuperTransaction
	}

	BlockRequest struct {
		Shard  uint
		Height uint
	}
	BlockAccountRequest struct {
		Shard uint
	}

	BlockAccountResponse struct {
		Shard       uint
		Height      uint //当前区块链的高度
		AccountList []structure.Account
		GSRoot      structure.GSRoot
	}

	BlockUploadRequest struct {
		Shard     uint
		Height    uint
		Id        string
		Block     structure.Block
		ReLayList map[uint][]structure.SuperTransaction
	}

	BlockUploadResponse struct {
		Shard   uint
		Height  uint
		Message string
	}

	TxWitnessRequest struct {
		Shard uint
	}

	TxWitnessResponse struct {
		Shard          uint
		Height         uint
		Num            int
		InternalList   map[uint][]structure.InternalTransaction
		CrossShardList map[uint][]structure.CrossShardTransaction
		RelayList      map[uint][]structure.SuperTransaction
	}

	TxWitnessRequest_2 struct {
		Shard          uint
		Height         uint
		Num            int
		InternalList   map[uint][]structure.InternalTransaction
		CrossShardList map[uint][]structure.CrossShardTransaction
		RelayList      map[uint][]structure.SuperTransaction
	}

	TxWitnessResponse_2 struct {
		Message string
	}

	RootUploadRequest struct {
		Shard  uint
		Height uint
		Id     string
		Root   string
	}

	RootUploadResponse struct {
		// Shard   uint
		Height  uint
		Message string
	}
)
