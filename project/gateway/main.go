package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/carefree/net/rpc"
	"github.com/carefree/project/gateway/auth"
	"github.com/carefree/project/gateway/router"
)

var (
	port                   = 3001
	portalServiceEndpoint  = "http://localhost:9003"
	accountServiceEndpoint = "http://localhost:9004"
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
	hcli, err := rpc.Dial(ctx, portalServiceEndpoint)
	if err != nil {
		log.Fatal(err)
	}
	acli, err := rpc.Dial(ctx, accountServiceEndpoint)
	if err != nil {
		log.Fatal(err)
	}
	sc := router.DefaultServiceConn
	sc.RegisterService(router.AccountService, acli)
	sc.RegisterService(router.PortalService, hcli)
	r := router.New(sc)
	r.Use(auth.JwtAuthentication)
	r.RegisterHandle()

	log.Printf("HTTP server listen on %v", ln.Addr())
	log.Fatal(http.Serve(ln, r))
}
