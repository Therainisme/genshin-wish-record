package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func getGachaLogByUrl(url string) (gachaLog *GachaList, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	jsonByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	r := Response{}
	err = json.Unmarshal(jsonByte, &r)
	if err != nil {
		return nil, err
	}

	return &r.Data.List, nil
}

func getCharacterGachaLog(args *RequestArgs) {
	gachaType := 301
	page := 1
	endId := "0"
	times := 0
	for {
		url, _ := getGachaLogUrl(args, gachaType, page, endId)
		gachaLog, err := getGachaLogByUrl(url)
		time.Sleep(time.Millisecond * 500)
		if err != nil {
			fmt.Printf("get err: %v\n", err)
		}

		size := len(*gachaLog)

		for _, v := range *gachaLog {
			times++
			if v.RankType == "4" || v.RankType == "5" {
				println(times, v.Name, v.Time)
			}
		}

		if size < 20 {
			break
		}

		page++
		endId = (*gachaLog)[size-1].Id

	}
}
