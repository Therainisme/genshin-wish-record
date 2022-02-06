package main

import (
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
)

func getUserHomeDir() string {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}

	return u.HomeDir
}

func getGenshinLogFilePath(userHomeDir string) string {
	return filepath.Join(userHomeDir, `AppData`, `LocalLow`, `miHoYo`, `原神`, `output_log.txt`)
}

func getUrlFromGenshinLogFile(path string) string {
	logFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	dataByte, err := ioutil.ReadAll(logFile)
	if err != nil {
		panic(err)
	}

	dataString := string(dataByte)

	compileRegex := regexp.MustCompile(`OnGetWebViewPageFinish:https://[\S]*log`)
	matchArr := compileRegex.FindAllString(dataString, -1)
	if len(matchArr) == 0 {
		panic("读取原神日志文件失败")
	}
	return matchArr[len(matchArr)-1][len("OnGetWebViewPageFinish:"):]
}
