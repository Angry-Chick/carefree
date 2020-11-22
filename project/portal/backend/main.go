package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/carefree/net/rpc"
	"github.com/carefree/project/common/db"
	"github.com/carefree/project/portal/backend/server/account"
	"github.com/carefree/project/portal/backend/server/slice"
	"github.com/carefree/project/portal/backend/server/space"
	"github.com/carefree/project/portal/backend/server/user"
	iac "github.com/carefree/project/portal/integration/account"
	"github.com/carefree/server"
	"gopkg.in/yaml.v2"
)

var (
	hport = 8080
	rport = 9090

	accountEndpoint = "localhost:9091"
)

func init() {
	flag.IntVar(&hport, "hport", hport, "HTTP server port")
	flag.IntVar(&rport, "rport", rport, "RPC server port")
}

func main() {
	yamlFile, err := ioutil.ReadFile("project/portal/cfg.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	var cfg db.Config
	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		fmt.Println(err.Error())
	}
	db, err := db.New(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	ac, err := rpc.Dial(ctx, accountEndpoint)

	rpc.Handle(space.NewServer(db))
	rpc.Handle(user.NewServer(db))
	rpc.Handle(slice.NewServer(db))
	rpc.Handle(account.NewServer(db, iac.NewUserClient(ac)))
	log.Fatal(server.Serve(hport, rport))
}
