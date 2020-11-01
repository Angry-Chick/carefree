package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var (
	port = 3000
	root = Resolve("../")
)

func init() {
	flag.IntVar(&port, "port", port, "HTTP server port")
}

func main() {
	hd, err := initHandlers()
	if err != nil {
		log.Fatal(err)
	}
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("HTTP server listen on %v", ln.Addr())
	log.Fatal(http.Serve(ln, hd))
}

var buildPath = "carefree/project/door/frontend/build"

func initHandlers() (*gin.Engine, error) {
	r := gin.Default()
	r.LoadHTMLFiles(path.Join(root, buildPath, "index.html"))
	r.Static("/static", path.Join(root, buildPath, "static"))
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Door",
		})
	})
	return r, nil
}

func Resolve(p string) string {
	r, err := filepath.Abs(p)
	if err != nil {
		return p
	}
	return r
}
