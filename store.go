package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type StoreData struct {
	CharacterGachaLog *GachaList `json:"character_gacha_log"`
	WeaponGachaLog    *GachaList `json:"weapon_gacha_log"`
	OrdinaryGachaLog  *GachaList `json:"ordinary_gacha_log"`
}

func mergeLocalGachaLog(gachaLogList ...*GachaList) (storeData *StoreData) {

	var uid = getUidByNewGachaLogList(gachaLogList...)

	storeData = readFromFile(uid)

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

	saveToFile(storeData, uid)
	return storeData
}

func getUidByNewGachaLogList(gachaLogList ...*GachaList) (uid string) {
	for _, gachaLog := range gachaLogList {
		if len(*gachaLog) == 0 {
			continue
		}

		if uid == "" {
			return (*gachaLog)[0].Uid
		}
	}

	return ""
}

func merge(old, src *GachaList) *GachaList {
	new := *old

	for _, v := range *src {
		if len(*old) == 0 || v.Id > (*old)[len(*old)-1].Id {
			new = append(new, v)
		}
	}

	return &new
}

func readFromFile(uid string) *StoreData {
	// new gacha data is completely empty
	// have not uid
	if uid == "" {
		return &StoreData{
			CharacterGachaLog: &GachaList{},
			WeaponGachaLog:    &GachaList{},
			OrdinaryGachaLog:  &GachaList{},
		}
	}

	f, err := os.Open(filepath.Join(".", uid+".json"))
	if err != nil {
		// the uid have no local gacha log
		if os.IsNotExist(err) {
			return &StoreData{
				CharacterGachaLog: &GachaList{},
				WeaponGachaLog:    &GachaList{},
				OrdinaryGachaLog:  &GachaList{},
			}
		}
	}

	// read data from local file by uid
	dataByte, _ := ioutil.ReadAll(f)
	storeData := &StoreData{}
	json.Unmarshal(dataByte, storeData)

	// rename old data file
	f.Close()
	err = os.Rename(filepath.Join(".", uid+".json"), filepath.Join(".", uid+".bak"))
	if err != nil {
		fmt.Printf("rename old file err: %v\n", err)
	}

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

	// delete bak file
	err = os.Remove(filepath.Join(".", uid+".bak"))
	if err != nil && !os.IsNotExist(err) {
		fmt.Printf("delete bak file err: %v\n", err)
	}
}
