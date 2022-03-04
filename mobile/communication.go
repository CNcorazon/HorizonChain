package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"horizon/server/model"
	"io/ioutil"
	"net/http"
)

const (
	url = "http://127.0.0.1:8080/witness/requestTransaction"
)

func main() {
	data := model.WitnessTransactionsRequest{
		Id: "TestID123",
	}
	payload, err := json.Marshal(&data)
	if err != nil {
		fmt.Println("Marshal err:", err)
		return
	}

	resp, err := http.Post(url, "application/x-www-form-urlencoded", bytes.NewBuffer(payload))

	if err != nil {
		fmt.Println("connection err:", err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
