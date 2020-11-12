package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/carefree/net/rpc"
	"github.com/carefree/project/home/frontend/router"
)

var (
	port                   = 3000
	homeServiceEndpoint    = "127.0.0.1:9090"
	accountServiceEndpoint = "127.0.0.1:9091"
)

func init() {
	flag.IntVar(&port, "port", port, "HTTP server port")
}

func main() {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	dcli, err := rpc.Dial(ctx, doorServiceEndpoint)
	if err != nil {
		log.Fatal(err)
	}
	ucli, err := rpc.Dial(ctx, userServiceEndpoint)
	if err != nil {
		log.Fatal(err)
	}
	sc := router.DefaultServiceConn
	sc.RegisterService(router.DoorService, dcli)
	sc.RegisterService(router.UserService, ucli)
	r := router.New(sc)
	r.RegisterHandle(ctx)

	log.Printf("HTTP server listen on %v", ln.Addr())
	log.Fatal(http.Serve(ln, r))
}
