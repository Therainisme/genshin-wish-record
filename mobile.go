package main

import "flag"

var mobile = flag.String("i", "", "手动输入url")

func init() {
	flag.Parse()
}
