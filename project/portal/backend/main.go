package main

import (
	"context"
	"flag"
	"log"

	"github.com/carefree/net/rpc"
	"github.com/carefree/project/common/config"
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

	serviceName     = "portal"
	accountEndpoint = "127.0.0.1:9090"
	configEndpoint  = "http://127.0.0.1:8848"
	configUsername  = "nacos"
	configPassword  = "nacos"
)

func init() {
	flag.IntVar(&hport, "hport", hport, "HTTP server port")
	flag.IntVar(&rport, "rport", rport, "RPC server port")
	flag.StringVar(&accountEndpoint, "account_endpoint", accountEndpoint, "account endpoint addr")
	flag.StringVar(&configEndpoint, "config_endpoint", configEndpoint, "config endpoint addr")
	flag.StringVar(&configUsername, "config_username", configUsername, "config server username")
	flag.StringVar(&configPassword, "config_password", configPassword, "config server password")
}

func main() {
	flag.Parse()
	cg, err := config.DefaultConfig(configEndpoint, configUsername, configPassword)
	if err != nil {
		log.Fatal(err)
	}
	ncli, err := config.NewClient(cg)
	if err != nil {
		log.Fatal(err)
	}
	cfgYaml, err := ncli.GetConfig("database", serviceName)
	if err != nil {
		log.Fatal(err)
	}
	var cfg db.Config
	err = yaml.Unmarshal([]byte(cfgYaml), &cfg)
	if err != nil {
		log.Fatal(err)
	}
	db, err := db.New(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	ac, err := rpc.Dial(ctx, accountEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	rpc.Handle(space.NewServer(db))
	rpc.Handle(user.NewServer(db))
	rpc.Handle(slice.NewServer(db))
	rpc.Handle(account.NewServer(db, iac.NewUserClient(ac)))
	log.Fatal(server.Serve(hport, rport))
}
