package main

import (
	"context"
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"
)

type FinalResult struct {
	CharacterGachaResult *Result
	WeaponGachaResult    *Result
	OrdinaryGachaResult  *Result
}

func (r *FinalResult) OutputHTML() {
	tmpl := template.Must(template.New("result").Parse(templateHTML))
	f, _ := os.Create("output.html")
	tmpl.Execute(f, r)

	// automatically open browser
	url := "file://" + filepath.Join(getCurrentAbPath(), "output.html")

	ctx, cancel := context.WithCancel(context.Background())
	cmd := exec.CommandContext(ctx, `cmd`, `/c`, `start`, url)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Start()

	fmt.Printf("Automatically open browser......\n")
	time.Sleep(10 * time.Second)

	cancel()

	cmd.Wait()
}

type Result struct {
	GachaName   string
	AverageLuck float64
	Total       int
	Unluck      int
	LuckList    LuckList
}

type LuckList []Luck

type Luck struct {
	Index int
	Name  string
	Type  string
}

func (r *Result) Print() {
	fmt.Printf("GachaName: %s\n", r.GachaName)
	fmt.Printf("Total: %d\n", r.Total)

	for _, v := range r.LuckList {
		switch v.Type {
		case "5":
			fmt.Printf("\033[1;33;40m[%d]%s\033[0m ", v.Index, v.Name)
		case "4":
			fmt.Printf("\033[1;35;40m[%d]%s\033[0m ", v.Index, v.Name)
		}
	}

	fmt.Printf("\n已经有%d未出金\n", r.Unluck)
	fmt.Printf("\n")
}

func analysisGachaLog(gachaLog *GachaList, gachaName string) *Result {
	luckList := make(LuckList, 0)
	idx := 0
	totalFiveStar := 0
	for _, v := range *gachaLog {
		idx++

		if v.RankType == "5" {
			luckList = append(luckList, Luck{
				Index: idx,
				Name:  v.Name,
				Type:  "5",
			})

			totalFiveStar++
			idx = 0
		} else if v.RankType == "4" {
			luckList = append(luckList, Luck{
				Index: idx,
				Name:  v.Name,
				Type:  "4",
			})
		}
	}

	return &Result{
		GachaName:   gachaName,
		AverageLuck: Decimal(float64(len(*gachaLog)-idx) / float64(totalFiveStar)),
		Total:       len(*gachaLog),
		Unluck:      idx,
		LuckList:    luckList,
	}
}

func analysisStoreData(storeData *StoreData) (finalResult *FinalResult) {
	return &FinalResult{
		CharacterGachaResult: analysisGachaLog(storeData.CharacterGachaLog, "角色活动祈愿"),
		WeaponGachaResult:    analysisGachaLog(storeData.WeaponGachaLog, "武器活动祈愿"),
		OrdinaryGachaResult:  analysisGachaLog(storeData.OrdinaryGachaLog, "常驻祈愿"),
	}
}
