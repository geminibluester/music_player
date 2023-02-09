package main

import (
	"flag"
	"music_player/httpserver"
)

var mode string
var address string

func init() {
	flag.StringVar(&mode, "m", "release", "是否开启debug调试")
	flag.StringVar(&address, "s", "0.0.0.0:8080", "服务地址")
}
func main() {
	flag.Parse()
	httpserver.RunServer(address, mode)
}
