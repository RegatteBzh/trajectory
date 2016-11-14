package main

import (
	"fmt"
	"log"
	"os"

	"github.com/regattebzh/trajectory/polar"
	"github.com/regattebzh/trajectory/wind"
)

func main() {
	csvFile, err := os.Open("./data/polar/vpp_1_1.csv")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer csvFile.Close()

	binFile, err := os.Open("./data/weather/gfs.t00z.pgrb2.1p00.f000.bin")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer binFile.Close()

	sail, _ := polar.ReadCsvPolar(csvFile)
	winds, _ := wind.ReadWind(binFile)

	fmt.Printf("%+v\n", winds[0])

	speed := sail.GetSpeed(6, 25)
	fmt.Printf("%f\n", speed)

}
