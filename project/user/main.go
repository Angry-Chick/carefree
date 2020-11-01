package main

import (
	"fmt"

	"github.com/carefree/project/common/jwt"
)

func main() {
	j := jwt.New()
	t, err := j.GenerateToken(&jwt.Claims{
		Name:  "ljy",
		Email: "www.123.com",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)
	p, err := j.ParseToken(t)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
}
