package main

func main() {
	r := parseArgs(testURL)
	characterGachaLog, weaponGachaLog, ordinaryGachaLog := getGachaLog(r)

	storeData := mergeLocalGachaLog(characterGachaLog, weaponGachaLog, ordinaryGachaLog)

	finalResult := analysisStoreData(storeData)
	finalResult.OutputHTML()
}
