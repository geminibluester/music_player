package main

import (
	"flag"
	"fmt"
	"music_player/httpserver"
	"os"
)

var mode string
var address string
var h bool

func init() {
	flag.BoolVar(&h, "h", false, "查看本帮助")
	flag.StringVar(&mode, "m", "release", "是否开启debug调试")
	flag.StringVar(&address, "s", "0.0.0.0:8080", "服务地址")
	flag.Usage = usage
}
func usage() {
	fmt.Fprintf(os.Stderr, `oliyo music server
`)
	flag.PrintDefaults()
}
func main() {
	flag.Parse()
	if h {
		flag.Usage()
		os.Exit(0)
	}
	httpserver.RunServer(address, mode)
}
