package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/carefree/net/rpc"
	"github.com/carefree/project/common/db"
	"github.com/carefree/project/home/backend/server/admin/home"
	"github.com/carefree/project/home/backend/server/admin/user"
	"github.com/carefree/project/home/backend/server/room"

	"github.com/carefree/server"

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
	yamlFile, err := ioutil.ReadFile("project/home/cfg.yaml")
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
	rpc.Handle(home.NewServer(db))
	rpc.Handle(user.NewServer(db))
	rpc.Handle(room.NewServer(db))
	log.Fatal(server.Serve(hport, rport))
}
