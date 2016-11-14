package polar

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

// ReadPolar reads polar information about a boat
func ReadPolar(csvFilename string) (sailChar SailCharacteristic, err error) {
	sailChar = SailCharacteristic{}
	sailChar.Polars = make([]Polar, 0)
	sailChar.Winds = make([]float64, 0)

	csvFile, err := os.Open(csvFilename)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.Comma = ';'
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
		return
	}

	sailChar.Winds = make([]float64, len(csvData[0]))
	for index, wind := range csvData[0][1:] {
		sailChar.Winds[index], _ = strconv.ParseFloat(wind, 32)
	}

	sailChar.Polars = make([]Polar, len(csvData)-1)
	for index, polarSample := range csvData {
		if index > 0 {
			newPolar := Polar{}
			newPolar.Angle, _ = strconv.ParseFloat(polarSample[0], 32)
			newPolar.speed = make([]float64, len(polarSample)-1)
			for i, speed := range polarSample {
				if i > 0 {
					newPolar.speed[i-1], _ = strconv.ParseFloat(speed, 32)
				}
			}
			sailChar.Polars[index-1] = newPolar
		}
	}

	return
}

// GetSpeed get an interpolation of the speed depending on the wind
// All speeds are in m/s
func (sailChar SailCharacteristic) GetSpeed(windAngle float64, windSpeed float64) (speed float64) {
	speed = 0
	// TODO interpolation
	// * on the wind
	// * on the angle
	return
}
