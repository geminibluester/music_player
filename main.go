package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	gogpt "github.com/sashabaranov/go-gpt3"
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
	chatgpt()
	// httpserver.RunServer(address, mode)
}
func chatgpt() {
	c := gogpt.NewClient("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")
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model:     gogpt.GPT3Ada,
		MaxTokens: 5,
		Prompt:    "Lorem ipsum",
		Stream:    true,
	}
	stream, err := c.CreateCompletionStream(ctx, req)
	if err != nil {
		return
	}
	defer stream.Close()

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("Stream finished")
			return
		}

		if err != nil {
			fmt.Printf("Stream error: %v\n", err)
			return
		}

		fmt.Printf("Stream response: %v\n", response)
	}
}
