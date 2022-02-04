package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
)

const RequestUrl = "https://hk4e-api.mihoyo.com/event/gacha_info/api/getGachaLog"

type RequestArgs struct {
	AuthkeyVer  int
	SignType    int
	AuthAppid   string
	InitType    int
	GachaId     string
	Timestamp   int
	Lang        string
	DeviceType  string
	Ext         string
	GameVersion string
	PlatType    string
	Region      string
	Authkey     string
	GameBiz     string
	EndId       string
}

func (r *RequestArgs) ToUrlValues() *url.Values {
	v := url.Values{}
	v.Set("authkey_ver", strconv.Itoa(r.AuthkeyVer))
	v.Set("sign_type", strconv.Itoa(r.SignType))
	v.Set("auth_appid", r.AuthAppid)
	v.Set("init_type", strconv.Itoa(r.InitType))
	v.Set("gacha_id", r.GachaId)
	v.Set("timestamp", strconv.Itoa(r.Timestamp))
	v.Set("lang", r.Lang)
	v.Set("device_type", r.DeviceType)
	v.Set("ext", r.Ext)
	v.Set("game_version", r.GameVersion)
	v.Set("plat_type", r.PlatType)
	v.Set("region", r.Region)
	v.Set("authkey", r.Authkey)
	v.Set("game_biz", r.GameBiz)
	v.Set("end_id", r.EndId)

	return &v
}

func parseArgs(str string) *RequestArgs {
	u, err := url.Parse(str)
	values := u.Query()
	if err != nil {
		fmt.Fprintf(os.Stderr, "url parse err: %v\n", err)
	}

	r := RequestArgs{}
	r.AuthkeyVer, _ = strconv.Atoi(values.Get("authkey_ver"))
	r.SignType, _ = strconv.Atoi(values.Get("sign_type"))
	r.AuthAppid = values.Get("auth_appid")
	r.InitType, _ = strconv.Atoi(values.Get("init_type"))
	r.GachaId = values.Get("gacha_id")
	r.Timestamp, _ = strconv.Atoi(values.Get("timestamp"))
	r.Lang = values.Get("lang")
	r.DeviceType = values.Get("device_type")
	r.Ext = values.Get("ext")
	r.GameVersion = values.Get("game_version")
	r.PlatType = values.Get("plat_type")
	r.Region = values.Get("region")
	r.Authkey = values.Get("authkey")
	r.GameBiz = values.Get("game_biz")

	return &r
}

func getGachaLogUrl(args *RequestArgs, gachaType int, page int, endId string) (string, error) {
	args.EndId = endId
	str := fmt.Sprintf("%s?%s&page=%d&gacha_type=%d&size=20", RequestUrl, url.Values(*args.ToUrlValues()).Encode(), page, gachaType)
	return str, nil
}
