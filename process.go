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

func getGachaLogByType(args *RequestArgs, gachaType int) (*GachaList, error) {
	page := 1
	endId := "0"
	characterGachaLog := make(GachaList, 0)
	for {
		url, _ := getGachaLogUrl(args, gachaType, page, endId)
		// todo: error handle

		gachaLog, err := getGachaLogByUrl(url)
		if err != nil {
			fmt.Printf("get err: %v\n", err)
		}
		characterGachaLog = append(characterGachaLog, *gachaLog...)
		time.Sleep(time.Millisecond * 200)

		size := len(*gachaLog)

		if size < 20 {
			break
		}

		page++
		endId = (*gachaLog)[size-1].Id
	}

	ReverseSlice(characterGachaLog)

	return &characterGachaLog, nil
}

func getGachaLog(args *RequestArgs) (characterGachaLog *GachaList, weaponGachaLog *GachaList, ordinaryGachaLog *GachaList) {
	// 角色活动祈愿 301
	characterGachaLog, err := getGachaLogByType(args, 301)
	if err != nil {
		fmt.Printf("get character gacha log err: %v\n", err)
	}

	// 武器活动祈愿 302
	weaponGachaLog, err = getGachaLogByType(args, 302)
	if err != nil {
		fmt.Printf("get weapon gacha log err: %v\n", err)
	}

	// 常驻祈愿 200
	ordinaryGachaLog, err = getGachaLogByType(args, 200)
	if err != nil {
		fmt.Printf("get ordinary gacha log err: %v\n", err)
	}

	return characterGachaLog, weaponGachaLog, ordinaryGachaLog
}
