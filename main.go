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
		log.Fatal(err) //log.Fatal run an os.Exit
	}
	defer csvFile.Close()

	binFile, err := os.Open("./data/weather/gfs.t00z.pgrb2.1p00.f000.bin")
	if err != nil {
		log.Fatal(err)
	}
	defer binFile.Close()

	sail, err := polar.ReadCsvPolar(csvFile)
	if err != nil { //do not skip err checking
		log.Fatal(err) //log.Fatal run an os.Exit
	}
	winds, err := wind.ReadWind(binFile)
	if err != nil {
		log.Fatal(err) //log.Fatal run an os.Exit
	}

	fmt.Printf("%+v\n", winds[0])

	speed := sail.GetSpeed(60, 25)
	fmt.Printf("%f\n", speed)

}
