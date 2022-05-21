package structure

import (
	"encoding/json"
	"log"
	"time"
)

type (
	Block struct {
		Header BlockHeader
		Body   BlockBody
	}

	Root        string
	BlockHeader struct {
		// Shard                uint   //表示是第几号分片中的区块
		Height               uint   //当前区块的高度
		Time                 int64  //区块产生的时候的Unix时间戳
		Vote                 uint   //本区块收到的移动节点的票数
		TransactionRoot      string //修改了本分片状态的交易区块的SHA256值
		SuperTransactionRoot string //产生的超级交易区块的SHA256值
		StateRoot            GSRoot //当前执行完本交易之后，当前区块链账本的世界状态
	}

	BlockBody struct {
		// Shard            uint
		Height           uint
		Transaction      TransactionBlock
		SuperTransaction SuperTransactionBlock
	}

	GSRoot struct {
		StateRoot string
		Vote      map[uint]map[string]int //记录每个执行分片计算出的subTreeRoot以及对应的票数
	}

	TransactionBlock struct {
		InternalList   map[uint][]InternalTransaction
		CrossShardList map[uint][]CrossShardTransaction
		SuperList      map[uint][]SuperTransaction //需要被打包进这个区块内部的SueprList
	}

	SuperTransactionBlock struct {
		SuperList map[uint][]SuperTransaction //执行完成TransactionList之后生成的一个SuperList
	}
)

func (t *TransactionBlock) CalculateRoot() string {
	jsonString, err := json.Marshal(t)
	if err != nil {
		log.Fatalln("计算交易区块Root失败")
	}
	return string(jsonString)
}

func (r *SuperTransactionBlock) CalculateRoot() string {
	jsonString, err := json.Marshal(r)
	if err != nil {
		log.Fatalln("计算接力交易区块Root失败")
	}
	return string(jsonString)
}

func MakeTransactionBlock(IntTraList map[uint][]InternalTransaction, CroList map[uint][]CrossShardTransaction, SuList map[uint][]SuperTransaction) TransactionBlock {
	res := TransactionBlock{
		InternalList:   IntTraList,
		CrossShardList: CroList,
		SuperList:      SuList,
	}
	return res
}

func MakeBlock(IntTraList map[uint][]InternalTransaction, CroList map[uint][]CrossShardTransaction, SuList map[uint][]SuperTransaction, s *State, height uint, gsroot GSRoot) Block {
	//首先打包生成本区快的交易区块
	transBlock := MakeTransactionBlock(IntTraList, CroList, SuList)
	//根据打包好的交易区块，记录生成的接力交易区块
	//只执行跨分片交易且不修改状态
	SuList_new := make(map[uint][]SuperTransaction, ShardNum)
	for i := 1; i <= ShardNum; i++ {
		for _, tran := range CroList[uint(i)] {
			res := ExcuteCross(tran, height, s, i)
			SuList_new[uint(i)] = append(SuList_new[uint(i)], *res)
		}
	}

	SuperBlock := SuperTransactionBlock{
		SuperList: SuList_new,
	}

	body := BlockBody{
		// Shard:            s.Shard,
		Height:           height,
		Transaction:      transBlock,
		SuperTransaction: SuperBlock,
	}

	header := BlockHeader{
		// Shard:                s.Shard,
		Height:               height,
		Time:                 time.Now().Unix(),
		Vote:                 1,
		TransactionRoot:      transBlock.CalculateRoot(),
		SuperTransactionRoot: SuperBlock.CalculateRoot(),
		StateRoot:            gsroot,
	}

	block := Block{
		Header: header,
		Body:   body,
	}

	return block
}

func CompareBlocks(b1 Block, b2 Block) bool {
	//只考虑BlockHeader中的三个root,time无所谓
	if b1.Header.Height != b2.Header.Height {
		log.Printf("两个区块的高度不同")
		return false
	} else if b1.Header.TransactionRoot != b2.Header.TransactionRoot {
		log.Printf("两个区块的交易列表不同")
		return false
	} else if b1.Header.SuperTransactionRoot != b2.Header.SuperTransactionRoot {
		log.Printf("两个区块生成的接力交易列表不同")
		return false
	} else if b1.Header.StateRoot.StateRoot != b2.Header.StateRoot.StateRoot {
		log.Printf("两个区块的state root不同")
		return false
	}
	log.Printf("高度为%v的区块经过检验", b1.Header.Height)
	return true
}
