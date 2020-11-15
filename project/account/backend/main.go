package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/carefree/net/rpc"
	"github.com/carefree/project/account/backend/server/admin/namespace"
	"github.com/carefree/project/account/backend/server/admin/user"
	"github.com/carefree/project/account/backend/server/account"
	"github.com/carefree/project/common/db"
	"github.com/carefree/server"

	"gopkg.in/yaml.v2"
)

var (
	dbpath = "project/account/cfg.yaml"

	hport = 8081
	rport = 9091
)

func init() {
	flag.IntVar(&hport, "hport", hport, "HTTP server port")
	flag.IntVar(&rport, "rport", rport, "RPC server port")
}

func main() {
	yamlFile, err := ioutil.ReadFile(dbpath)
	if err != nil {
		log.Fatal(err)
	}
	var cfg db.Config
	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		log.Fatal(err)
	}
	db, err := db.New(&cfg)
	if err != nil {
		log.Fatalf("cannot connect database, err:%v", err)
	}
	rpc.Handle(account.NewServer(db))
	rpc.Handle(namespace.NewServer(db))
	rpc.Handle(user.NewServer(db))
	log.Fatal(server.Serve(hport, rport))
}
