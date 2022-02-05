package main

import "fmt"

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

func Analysis(gachaLog *GachaList, gachaName string) *Result {
	luckList := make(LuckList, 0)
	idx := 0
	for _, v := range *gachaLog {
		idx++

		if v.RankType == "5" {
			luckList = append(luckList, Luck{
				Index: idx,
				Name:  v.Name,
				Type:  "5",
			})

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
		AverageLuck: 0, // todo
		Total:       len(*gachaLog),
		Unluck:      idx,
		LuckList:    luckList,
	}
}
