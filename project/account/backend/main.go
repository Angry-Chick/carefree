package main

import (
	"flag"
	"log"

	"github.com/carefree/net/rpc"
	"github.com/carefree/project/account/backend/server/account"
	"github.com/carefree/project/account/backend/server/user"
	"github.com/carefree/project/common/config"
	"github.com/carefree/project/common/db"
	"github.com/carefree/server"

	"gopkg.in/yaml.v2"
)

var (
	hport = 8080
	rport = 9090

	serviceName    = "account"
	configEndpoint = "http://127.0.0.1:8848"
	configUsername = "nacos"
	configPassword = "nacos"
)

func init() {
	flag.IntVar(&hport, "hport", hport, "HTTP server port")
	flag.IntVar(&rport, "rport", rport, "RPC server port")
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
		log.Fatalf("cannot connect database, err:%v", err)
	}
	rpc.Handle(account.NewServer(db))
	rpc.Handle(user.NewServer(db))
	log.Fatal(server.Serve(hport, rport))
}
