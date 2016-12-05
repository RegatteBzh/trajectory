package main

import (
	"fmt"
	"image"
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

	// etopoFile, err := os.Open("./data/etopo/etopo1_ice_g_i2.bin")
	// if err != nil {
	// 	log.Fatal(err) //log.Fatal run an os.Exit
	// }
	// defer etopoFile.Close()

	sail, err := polar.ReadCsvPolar(csvFile)
	if err != nil { //do not skip err checking
		log.Fatal(err) //log.Fatal run an os.Exit
	}
	winds, err := wind.Read(binFile)
	if err != nil {
		log.Fatal(err) //log.Fatal run an os.Exit
	}

	// etopoData, err := etopo.Read(etopoFile)
	// if err != nil {
	// 	log.Fatal(err) //log.Fatal run an os.Exit
	// }

	myWind, _ := wind.GetWind(winds, image.Point{00, 0})
	fmt.Printf("Wind: %+v\n", myWind)

	speed := sail.GetSpeed(60, 25)
	fmt.Printf("Polar: %f\n", speed)

	// topo, _ := etopo.GetAltitude(etopoData, image.Point{00, 0})
	// fmt.Printf("Altitude: %d\n", topo)

}
