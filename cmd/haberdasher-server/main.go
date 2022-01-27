package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/athega/haberdasher"
	"github.com/athega/haberdasher/rpc"
)

func main() {
	var (
		logger  = log.New(os.Stdout, "", 0)
		service = haberdasher.NewService(logger)
		server  = rpc.NewHaberdasherServer(service)

		//serverHooks = &twirp.ServerHooks{
		//	RequestRouted: func(ctx context.Context) (context.Context, error) {
		//		method, _ := twirp.MethodName(ctx)
		//		logger.Println("Method: " + method)
		//		return ctx, nil
		//	},
		//	Error: func(ctx context.Context, twerr twirp.Error) context.Context {
		//		logger.Println("Error: " + string(twerr.Code()))
		//		return ctx
		//	},
		//	ResponseSent: func(ctx context.Context) {
		//		logger.Println("Response Sent (error or success)")
		//	},
		//}
		//server = rpc.NewHaberdasherServer(service, twirp.WithServerHooks(serverHooks))

		hs = &http.Server{Addr: ":8080", Handler: server}
	)

	logger.Printf("Serving %s\n", serviceLink(hs.Addr))

	if err := hs.ListenAndServe(); err != nil {
		panic(err)
	}
}

func serviceLink(addr string) string {
	return fmt.Sprintf("http://localhost%s/twirp/athega.Haberdasher", addr)
}
