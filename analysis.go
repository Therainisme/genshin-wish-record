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
		fmt.Printf("%d: %s\n", v.Index, v.Name)
	}

	fmt.Printf("已经有%d未出金\n", r.Unluck)
	fmt.Printf("-------------------\n")
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
				Type:  "", // todo
			})

			idx = 0
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
