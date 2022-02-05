package main

func main() {
	r := parseArgs(testURL)
	characterGachaLog, weaponGachaLog, ordinaryGachaLog := getGachaLog(r)

	mergeLocalGachaLog(characterGachaLog, weaponGachaLog, ordinaryGachaLog)
}
