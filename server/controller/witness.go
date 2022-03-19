package controller

import (
	"horizon/server/model"

	"github.com/gin-gonic/gin"
)

type BlockStorage struct {
	Shard_id      int
	Shard_length  int
	Previous_hash string
	Merkleroot    string
}

var MyChain BlockStorage

func WitnessGetTransaction(c *gin.Context) {
	//声明接收到的变量
	// var json model.WitnessTransactionsRequest

	// logrus.Info("收到了%s的请求", c.ClientIP())

	// //判断请求的结构体是否符合定义
	// if err := c.ShouldBindJSON(&json); err != nil {
	// gin.H封装了生成json数据的工具
	// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// return
	// }

	//如果符合请求的话就返回结果，这里只是模仿的结果
	res := model.WitnessTransactionResponse{
		TransactionList: []string{"Hello", "World", "Test"},
	}

	//c.JSON可以把Go语言结构体转成json格式序列化之后发送
	c.JSON(200, res)
}

func WitnessGetProposalMsg(c *gin.Context) {
	res := model.WitnessProposalMsg{
		Shard_id:     MyChain.Shard_id,
		Shard_length: MyChain.Shard_length,
	}
	c.JSON(200, res)
}
