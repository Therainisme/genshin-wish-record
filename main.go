package main

func main() {
	userHomeDir := getUserHomeDir()
	genshinLogFilePath := getGenshinLogFilePath(userHomeDir)
	url := getUrlFromGenshinLogFile(genshinLogFilePath)

	args := parseArgs(url)
	characterGachaLog, weaponGachaLog, ordinaryGachaLog := getGachaLog(args)

	storeData := mergeLocalGachaLog(characterGachaLog, weaponGachaLog, ordinaryGachaLog)

	finalResult := analysisStoreData(storeData)
	finalResult.OutputHTML()
}
