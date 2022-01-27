package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/athega/haberdasher/rpc"
)

func main() {
	var (
		baseURL string
		inches  uint64
	)

	flag.StringVar(&baseURL, "base", "http://localhost:8080", "Base URL to the Twirp server")
	flag.Uint64Var(&inches, "inches", 22, "Size of the hat in inches")
	flag.Parse()

	var (
		httpClient = &http.Client{}
		client     = rpc.NewHaberdasherProtobufClient(baseURL, httpClient)
		ctx        = context.Background()
		size       = &rpc.Size{Inches: inches}
	)

	hat, err := client.MakeHat(ctx, size)
	if err != nil {
		fmt.Printf("🔥 oh no: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("🎩 I have a nice new hat: %+v\n", hat)
}
