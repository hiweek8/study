package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://api.xuebaclass.com/xuebaapi/v1/provinces"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	type Province struct {
		Id       int    `json:"id"`
		Province string `json:"province"`
	}
	provinces := make([]Province, 0)
	err := json.Unmarshal([]byte(body), &provinces)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(provinces)

}
