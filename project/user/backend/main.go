package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/carefree/project/common/db"
	"github.com/carefree/project/user/backend/server/user"
	"github.com/carefree/server"
	"github.com/carefree/server/rpc"

	"gopkg.in/yaml.v2"
)

var (
	dbpath = "project/user/cfg.yaml"

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
	rpc.Handle(user.NewServer(db))
	log.Fatal(server.Serve(hport, rport))
}
