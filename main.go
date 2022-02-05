package main

func main() {
	r := parseArgs(testURL)
	characterGachaLog, weaponGachaLog, ordinaryGachaLog := getGachaLog(r)

	characterResult := Analysis(characterGachaLog, "角色活动祈愿")
	characterResult.Print()
	weaponResult := Analysis(weaponGachaLog, "武器活动祈愿")
	weaponResult.Print()
	ordinaryResult := Analysis(ordinaryGachaLog, "普通活动祈愿")
	ordinaryResult.Print()

	mergeLocalGachaLog(characterGachaLog, weaponGachaLog, ordinaryGachaLog)
}
