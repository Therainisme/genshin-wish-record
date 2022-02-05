package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type StoreData struct {
	CharacterGachaLog *GachaList `json:"character_gacha_log"`
	WeaponGachaLog    *GachaList `json:"weapon_gacha_log"`
	OrdinaryGachaLog  *GachaList `json:"ordinary_gacha_log"`
}

func mergeLocalGachaLog(gachaLogList ...*GachaList) {

	// 尝试通过新记录获取 uid
	var uid string
	for _, gachaLog := range gachaLogList {
		if len(*gachaLog) == 0 {
			continue
		}

		if uid == "" {
			uid = (*gachaLog)[0].Uid
		}
	}

	// 查询本地对应 uid 的记录
	var storeData StoreData = *readFromFile(uid)

	// 合并
	for i, gachaLog := range gachaLogList {
		switch i {
		case 0:
			storeData.CharacterGachaLog = merge(storeData.CharacterGachaLog, gachaLog)
		case 1:
			storeData.WeaponGachaLog = merge(storeData.WeaponGachaLog, gachaLog)
		case 2:
			storeData.OrdinaryGachaLog = merge(storeData.OrdinaryGachaLog, gachaLog)
		}
	}

	// 保存
	saveToFile(&storeData, uid)
}

func merge(old, src *GachaList) (new *GachaList) {
	new = old

	for _, v := range *src {
		if len(*old) == 0 || v.Id > (*old)[len(*old)-1].Id {
			*new = append(*new, v)
		}
	}

	return new
}

func readFromFile(uid string) *StoreData {
	// 无法从新记录中找到 uid
	if uid == "" {
		return &StoreData{
			CharacterGachaLog: &GachaList{},
			WeaponGachaLog:    &GachaList{},
			OrdinaryGachaLog:  &GachaList{},
		}
	}

	f, err := os.Open(filepath.Join(".", uid+".json"))
	if err != nil {
		// 本地没有 uid 对应的记录
		if os.IsNotExist(err) {
			return &StoreData{
				CharacterGachaLog: &GachaList{},
				WeaponGachaLog:    &GachaList{},
				OrdinaryGachaLog:  &GachaList{},
			}
		}
	}
	defer f.Close()

	// 从本地对应 uid 记录中读取数据
	dataByte, _ := ioutil.ReadAll(f)
	storeData := &StoreData{}
	json.Unmarshal(dataByte, storeData)

	return storeData
}

func saveToFile(storeData *StoreData, uid string) {
	f, err := os.Open(filepath.Join(".", uid+".json"))
	if err != nil {
		if os.IsNotExist(err) {
			f, _ = os.Create(filepath.Join(".", uid+".json"))
		}
	}
	defer f.Close()

	dataByte, _ := json.Marshal(storeData)
	f.Write(dataByte)
}
