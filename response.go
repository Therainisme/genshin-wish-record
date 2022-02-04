package main

type Response struct {
	Retcode int          `json:"retcode"`
	Message string       `json:"message"`
	Data    ResponseData `json:"data"`
}

type ResponseData struct {
	Page   string    `json:"page"`
	Size   string    `json:"size"`
	Total  string    `json:"total"`
	List   GachaList `json:"list"`
	Region string    `json:"region"`
}

type GachaList []GachaListItem

type GachaListItem struct {
	Uid       string `json:"uid"`
	GachaType string `json:"gacha_type"`
	ItemId    string `json:"item_id"`
	Count     string `json:"count"`
	Time      string `json:"time"`
	Name      string `json:"name"`
	Lang      string `json:"lang"`
	ItemType  string `json:"item_type"`
	RankType  string `json:"rank_type"`
	Id        string `json:"id"`
}
