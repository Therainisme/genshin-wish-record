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
		return nil, fmt.Errorf("请求祈愿接口失败，错误信息：%v", err)
	}

	jsonByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("解析响应体失败, 错误信息：%v", err)
	}

	r := Response{}
	err = json.Unmarshal(jsonByte, &r)
	if err != nil {
		return nil, fmt.Errorf("解析响应体失败, 错误信息：%v", err)
	}

	if r.Retcode < 0 {
		return nil, fmt.Errorf("URL 过期或错误, 请重新按照 README 文档获取")
	}

	return &r.Data.List, nil
}

func getGachaLogByType(args *RequestArgs, gachaType int) *GachaList {
	page := 1
	endId := "0"
	characterGachaLog := make(GachaList, 0)
	for {
		url := getGachaLogUrl(args, gachaType, page, endId)

		gachaLog, err := getGachaLogByUrl(url)
		if err != nil {
			panic(fmt.Sprintf("错误信息：%v。卡池代码：%d", err, gachaType))
		}

		characterGachaLog = append(characterGachaLog, *gachaLog...)
		time.Sleep(time.Millisecond * 300)

		size := len(*gachaLog)

		if size < 20 {
			break
		}

		page++
		endId = (*gachaLog)[size-1].Id
	}

	ReverseSlice(characterGachaLog)

	return &characterGachaLog
}

func getGachaLog(args *RequestArgs) (characterGachaLog *GachaList, weaponGachaLog *GachaList, ordinaryGachaLog *GachaList) {
	// 角色活动祈愿 301
	characterGachaLog = getGachaLogByType(args, 301)

	// 武器活动祈愿 302
	weaponGachaLog = getGachaLogByType(args, 302)

	// 常驻祈愿 200
	ordinaryGachaLog = getGachaLogByType(args, 200)

	return characterGachaLog, weaponGachaLog, ordinaryGachaLog
}
