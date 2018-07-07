package main

import (
	"flag"
	"fmt"

	"github.com/ComputePractice2018/businesscard/backend/utils"
)

func main() {
	var name = flag.String("name", "Ilay", "имя для приветствия")
	flag.Parse()
	fmt.Println(utils.GetHelloWorldString(*name))
}
