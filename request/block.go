package request

import (
	"bytes"
	"encoding/json"
	"horizon/model"
	"horizon/structure"
	"io/ioutil"
	"log"
	"net/http"
)

func WitnessTransaction(shardNum uint, url string, route string) model.TxWitnessResponse {
	URL := url + route
	data := model.TxWitnessRequest{
		Shard: shardNum,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	request, _ := http.NewRequest("POST", URL, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(response.Body)
	var res model.TxWitnessResponse
	json.Unmarshal(body, &res)
	return res
}

func WitnessTransaction_2(shardNum uint, url string, route string, txlist model.TxWitnessResponse) model.TxWitnessResponse_2 {
	URL := url + route
	data := model.TxWitnessRequest_2{
		Shard:          shardNum,
		Height:         txlist.Height,
		Num:            txlist.Num,
		InternalList:   txlist.InternalList,
		CrossShardList: txlist.CrossShardList,
		RelayList:      txlist.RelayList,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	request, _ := http.NewRequest("POST", URL, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(response.Body)
	var res model.TxWitnessResponse_2
	json.Unmarshal(body, &res)
	return res
}

func RequestTransaction(shardNum uint, url string, route string, id string) model.BlockTransactionResponse {
	URL := url + route
	data := model.BlockTransactionRequest{
		Shard: shardNum,
		Id:    id,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	request, _ := http.NewRequest("POST", URL, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(response.Body)
	var res model.BlockTransactionResponse
	json.Unmarshal(body, &res)
	return res
}

func RequestAccount(shardNum uint, url string, route string) model.BlockAccountResponse {
	URL := url + route
	data := model.BlockAccountRequest{
		Shard: shardNum,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	request, _ := http.NewRequest("POST", URL, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(response.Body)
	var res model.BlockAccountResponse
	json.Unmarshal(body, &res)
	return res
}

func UploadBlock(shardNum uint, block structure.Block, id string, url string, route string) model.BlockUploadResponse {
	URL := url + route
	data := model.BlockUploadRequest{
		Shard:     shardNum,
		Height:    block.Header.Height,
		Id:        id,
		Block:     block,
		ReLayList: block.Body.SuperTransaction.SuperList,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	request, _ := http.NewRequest("POST", URL, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(response.Body)
	var res model.BlockUploadResponse
	json.Unmarshal(body, &res)
	return res
}

// 请求已见证过的交易，进行验证
func RequestBlock(shardNum uint, url string, route string) model.BlockTransactionResponse {
	URL := url + route
	data := model.BlockTransactionRequest{
		Shard: shardNum,
		// Id:    id,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	request, _ := http.NewRequest("POST", URL, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(response.Body)
	var res model.BlockTransactionResponse
	json.Unmarshal(body, &res)
	return res
}

func UploadRoot(shardNum uint, id string, height uint, root string, url string, route string) model.RootUploadResponse {
	URL := url + route
	data := model.RootUploadRequest{
		Shard:  shardNum,
		Height: height,
		Id:     id,
		Root:   root,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	request, _ := http.NewRequest("POST", URL, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(response.Body)
	var res model.RootUploadResponse
	json.Unmarshal(body, &res)
	return res
}
