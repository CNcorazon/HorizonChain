package main

import (
	"encoding/json"
	"fmt"
	"horizon/model"
	"horizon/request"
	"horizon/structure"
	"log"
	"time"
)

const (
	HTTPURL = "http://127.0.0.1:8080"
	WSURL   = "ws://127.0.0.1:8080"
	// HTTPURL           = "http://172.18.166.60:8800"
	// WSURL             = "ws://http://172.18.166.60:8800"
	blockTransaction  = "/block/transaction"
	blockAccount      = "/block/account"
	blockUpload       = "/block/upload"
	blockWitness      = "/block/witness"
	blockWitness_2    = "/block/witness_2"
	blockTxValidation = "/block/validate"
	blockUploadRoot   = "/block/uploadroot"

	shardNum       = "/shard/shardNum"
	phaseNum       = "/shard/phase"
	register       = "/shard/register"
	multicastblock = "/shard/block"
	sendtvote      = "/shard/vote"
	heightNum      = "/shard/height"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	//首先发起请求获取当前可以加入的shard
	for {
		shard, flag1, flag2 := request.ShardRequest(HTTPURL, shardNum)
		//首先判断flag2
		if !flag2 {
			log.Printf("服务器尚未开启")
			time.Sleep(5 * time.Second)
		}
		if !flag1 {
			log.Printf("当前没有分片需要节点")
			time.Sleep(5 * time.Second)
			continue
		}
		log.Printf("被分配到了%v分片作为移动节点执行", shard)

		conn := request.RegisterWSRequest(shard, WSURL, register)

		var id, winid string
		var idlist []string
		var winflag bool
		var consensusflag bool
		var metamessage model.MessageMetaData
		//根据传来的消息获取本客户端的Id和本轮的胜者id
		conn.ReadJSON(&metamessage)

		if metamessage.MessageType == 1 {
			var iswin model.MessageIsWin
			err := json.Unmarshal(metamessage.Message, &iswin)
			if err != nil {
				log.Printf("err")
				return
			}
			consensusflag = iswin.IsConsensus
			winflag = iswin.IsWin
			id = iswin.PersonalID
			winid = iswin.WinID
			for i := 0; i < len(iswin.IdList); i++ {
				idlist = append(idlist, iswin.IdList[i])
			}
		}

		//进行区块的共识
		if consensusflag && winflag {
			shard = uint(0)
			//如果获得了提出区块的权利，就获得了打包区块的权利
			//首先是获取交易列表，获取账户列表，生成状态，从而得到新的区块
			log.Println("进入委员会，获得了区块的提出权")
			TransactionList := request.RequestTransaction(shard, HTTPURL, blockTransaction, id)
			//验证交易的区块见证，即其他移动节点的签名
			//这边假设有1000个节点
			time.Sleep(1000 * structure.SIGN_VERIFY_TIME * time.Millisecond)
			//请求账户的状态
			//和各分片树根签名 accList.GSRoot
			accList := request.RequestAccount(shard, HTTPURL, blockAccount)
			log.Println(accList.GSRoot)
			// log.Println(TransactionList.RelayList)
			state := structure.MakeStateWithAccount(shard, accList.AccountList, accList.GSRoot)
			// 验证树根签名
			time.Sleep(structure.ShardNum * 1000 * structure.SIGN_VERIFY_TIME * time.Millisecond)
			newBlock := structure.MakeBlock(TransactionList.InternalList, TransactionList.CrossShardList, TransactionList.RelayList, state, TransactionList.Height, accList.GSRoot)

			blockPointer := &newBlock
			//需要将该区块发送给分片内部的成员共识一下
			resp := request.MulticastBlock(shard, HTTPURL, multicastblock, newBlock, id)
			log.Println(resp.Message)
			//等待一段时间获取投票结果
			voteMap := make(map[string]bool) //记录最终的投票结果
			//等待所有的投票结果
			for i := 0; i < len(idlist)-1; i++ {
				var metaMessage model.MessageMetaData
				conn.ReadJSON(&metaMessage)
				// log.Println(metaMessage.MessageType)
				if metaMessage.MessageType == 3 {
					var vote model.SendVoteRequest
					err := json.Unmarshal(metaMessage.Message, &vote)
					if err != nil {
						log.Println(err)
						return
					}
					voteMap[vote.PersonalID] = vote.Agree
					if vote.Agree {
						log.Printf("区块获得了来自%v节点的投票", vote.PersonalID)
					}

				}
			}
			//最终根据投票结果更新区块接收到的票数目
			for _, value := range voteMap {
				if value {
					blockPointer.Header.Vote++
				}
			}
			finalBlock := *blockPointer
			//最后提交区块
			res := request.UploadBlock(shard, finalBlock, winid, HTTPURL, blockUpload)
			log.Printf("分片%v%v,当前链的高度为%v", res.Shard, res.Message, res.Height)
			// time.Sleep(time.Second)
		} else if consensusflag && !winflag {
			//如果没有获得出块权利，就只能等待Leader生成的区块
			log.Printf("进入委员会，等待其他节点提出区块")
			var metaMessage model.MessageMetaData
			conn.ReadJSON(&metaMessage)
			if metaMessage.MessageType == 2 {
				var blockMessage model.MultiCastBlockRequest
				json.Unmarshal(metaMessage.Message, &blockMessage)
				//验证收到的Leader执行生成的区块是否正确
				transactionList := blockMessage.Block.Body.Transaction
				//验证交易的区块见证，即其他移动节点的签名
				time.Sleep(1000 * structure.SIGN_VERIFY_TIME * time.Millisecond)
				// log.Println(transactionList.SuperList)
				accList := request.RequestAccount(shard, HTTPURL, blockAccount)
				// log.Println(accList.Shard)
				state := structure.MakeStateWithAccount(shard, accList.AccountList, accList.GSRoot)
				time.Sleep(structure.ShardNum * 1000 * structure.SIGN_VERIFY_TIME * time.Millisecond)
				newBlock2 := structure.MakeBlock(transactionList.InternalList, transactionList.CrossShardList, transactionList.SuperList, state, accList.Height+1, accList.GSRoot)
				//检验区块
				log.Printf("下面检查区块是否正确")
				// log.Println(blockMessage.Block.Header.StateRoot.StateRoot)
				// log.Println(accList)
				flag := structure.CompareBlocks(blockMessage.Block, newBlock2)

				//进行投票
				resp := request.SendVote(shard, int(newBlock2.Header.Height), winid, id, flag, HTTPURL, sendtvote)
				log.Println(resp.Message)
				// time.Sleep(time.Second)
			}
		} else {
			log.Printf("未进入委员会, 进行区块见证")
			// 如果是执行节点，进行区块见证/交易验证
			// 当共识完成之后就停止区块见证，进行交易验证

			height_old := request.HeightRequest(HTTPURL, heightNum)
			log.Printf("区块见证前的区块高度为%v", height_old)
			for {
				height_new := request.HeightRequest(HTTPURL, heightNum)
				// log.Printf("此时的区块高度为%v", height_new)
				if height_new != height_old {
					log.Printf("共识完成，区块见证结束")
					break
				}
				time.Sleep(structure.SIGN_VERIFY_TIME * time.Millisecond)
				txlist := request.WitnessTransaction(shard, HTTPURL, blockWitness)
				// log.Println(txlist.InternalList)
				request.WitnessTransaction_2(shard, HTTPURL, blockWitness_2, txlist)
				// fmt.Println(res.Message)
			}
			log.Printf("开始交易验证")
			//交易验证哈希
			time.Sleep(structure.TX_NUM * 3 * structure.SIGN_VERIFY_TIME * time.Millisecond)
			transaction := request.RequestBlock(shard, HTTPURL, blockTxValidation)
			accList := request.RequestAccount(shard, HTTPURL, blockAccount)
			log.Printf("gsroot:%v", accList.GSRoot)
			state := structure.MakeStateWithAccount(shard, accList.AccountList, accList.GSRoot)
			txlist := structure.TransactionBlock{
				InternalList:   transaction.InternalList,
				CrossShardList: transaction.CrossShardList,
				SuperList:      transaction.RelayList,
			}
			root := structure.UpdateStateWithTxBlock(txlist, transaction.Height, state, shard)
			res := request.UploadRoot(shard, id, transaction.Height, root, HTTPURL, blockUploadRoot)
			fmt.Println(res.Message)
			// r := rand.New(rand.NewSource(time.Now().UnixNano()))
			// num := r.Intn(5)
			// time.Sleep(5 * time.Second)
		}
	}
}
