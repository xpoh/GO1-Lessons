package main

import (
	"./conf"
	"fmt"
)

func main() {
	var c, err = conf.ReadConfig()
	if err != nil {
		panic(err)
	}
	fmt.Print(c, err)
}
