package main

import (
	"fmt"

	"github.com/regattebzh/trajectory/polar"
)

func main() {
	sail, _ := polar.ReadCsvPolar("./data/vpp_1_1.csv")

	speed := sail.GetSpeed(6, 25)

	fmt.Printf("%f\n", speed)
}
