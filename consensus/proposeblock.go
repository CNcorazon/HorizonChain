package consensus

import (
	"horizon/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//输入请求发送至comm，服务器返回相应信息
func collectMsg(msg string) []string {
	var candidateproblocks []string

	return candidateproblocks
}

//根据收到的信息和自己的vrf，生成proposalblock
func Proposeblock(vrf float64, public_key string, c *gin.Context) {
	var proposemsg string
	var msg []string = collectMsg(proposemsg)
	shard_id, err := strconv.Atoi(msg[0])
	if err != nil {
		println("load shard_id failed")
	}
	shard_length, err := strconv.Atoi(msg[1])
	if err != nil {
		println("load shard_length failed")
	}
	previous_hash := msg[2]
	txblock_list := msg[3]
	globalstate_minus2 := msg[4]
	time_stamp := time.Now().UnixNano() / 1e6 //毫秒级别的时间戳
	proposal_block := model.ProBlock{
		Shard_id:           shard_id,
		Shard_length:       shard_length,
		Previous_hash:      previous_hash,
		Globalstate_minus2: globalstate_minus2,
		Txblock_list:       txblock_list,
		Time_stamp:         time_stamp,
		Vrf:                vrf,
		Public_key:         public_key,
	}
	//msg := model.ProBlock{Shard_id: shard_id, Shard_length: shard_length, Previous_hash: previous_hash, Txblock_list: txblock_list, Globalstate_minus2: globalstate_minus2, Time_stamp: time_stamp, VRF: vrf, Public_key: public_key}
	//不会通信
	c.JSON(200, proposal_block)
}
