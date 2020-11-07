package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/carefree/project/common/db"
	"github.com/carefree/project/door/backend/server/door"
	"github.com/carefree/project/door/backend/server/namespace"
	"github.com/carefree/server"
	"github.com/carefree/net/rpc"

	"gopkg.in/yaml.v2"
)

var (
	hport = 8080
	rport = 9090
)

func init() {
	flag.IntVar(&hport, "hport", hport, "HTTP server port")
	flag.IntVar(&rport, "rport", rport, "RPC server port")
}

func main() {
	yamlFile, err := ioutil.ReadFile("project/door/cfg.yaml")
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
	rpc.Handle(namespace.NewServer(db))
	rpc.Handle(door.NewServer(db))
	log.Fatal(server.Serve(hport, rport))
}
