package main

import (
	"fmt"

	"github.com/regattebzh/trajectory/polar"
)

func main() {
	sail, _ := polar.ReadPolar("./data/vpp_1_1.csv")
	fmt.Printf("%+v\n", sail)
}
