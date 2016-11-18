package polar

import (
	"encoding/csv"
	"io"
	"log"
	"math"
	"sort"
	"strconv"
)

type byAngle []Polar

func (a byAngle) Len() int           { return len(a) }
func (a byAngle) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAngle) Less(i, j int) bool { return a[i].Angle < a[j].Angle }

// ReadCsvPolar reads polar information about a boat
func ReadCsvPolar(csvFile io.Reader) (*SailCharacteristic, error) {
	sailChar := &SailCharacteristic{}
	sailChar.Polars = make([]Polar, 0)
	sailChar.Winds = make([]float64, 0)

	reader := csv.NewReader(csvFile)
	reader.Comma = ';'
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	sailChar.Winds = make([]float64, len(csvData[0]))
	for index, wind := range csvData[0][1:] {
		sailChar.Winds[index], err = strconv.ParseFloat(wind, 32)
		if err != nil {
			log.Printf("Error parsing wind data %s\n", wind)
			continue
		}
		// knot to m/s conversion
		sailChar.Winds[index] = sailChar.Winds[index] * float64(0.514444)
	}

	sort.Float64s(sailChar.Winds)

	sailChar.Polars = make([]Polar, len(csvData)-1)
	for index, polarSample := range csvData {
		//skip the firt record
		if index == 0 {
			continue
		}
		newPolar := Polar{}
		newPolar.Angle, _ = strconv.ParseFloat(polarSample[0], 32)
		newPolar.Speed = make([]float64, len(polarSample)-1)
		for i, speed := range polarSample {
			if i > 0 {
				newPolar.Speed[i-1], err = strconv.ParseFloat(speed, 32)
				if err != nil {
					log.Printf("Error parsing wind speed %s\n", speed)
					continue
				}
				// knot to m/s conversion
				newPolar.Speed[i-1] = newPolar.Speed[i-1] * float64(0.514444)
			}
		}
		sailChar.Polars[index-1] = newPolar
	}

	sort.Sort(byAngle(sailChar.Polars))

	return sailChar, nil
}

func getCursorValue(first float64, second float64, cursor float64) float64 {
	if first > second {
		first, second = second, first
	}
	return first + (second-first)*cursor
}

// GetSpeed get an interpolation of the speed depending on the wind
// @param windAngle Angle between boat and wind (0 is front of the boat)
// @param windSpeed Speed of the wind in m/s
// @return boat speed in m/s
//
func (sailChar *SailCharacteristic) GetSpeed(windAngle float64, windSpeed float64) (speed float64) {
	speed = 0
	firstPolar := Polar{}
	secondPolar := Polar{}

	windAngle = math.Mod(windAngle, 180)

	// get polar before and after windAngle
	for _, polarSample := range sailChar.Polars {
		secondPolar = polarSample
		if polarSample.Angle > windAngle {
			break
		}
		firstPolar = polarSample
	}

	angleCrusor := (windAngle - float64(firstPolar.Angle)) / (float64(secondPolar.Angle) - float64(firstPolar.Angle))

	// get windIndex
	firstWindIndex := 0
	secondWindIndex := 0
	for index, wind := range sailChar.Winds {
		secondWindIndex = index
		if wind > windSpeed {
			break
		}
		firstWindIndex = index
	}

	windCursor := float64(0)
	if secondWindIndex != firstWindIndex {
		windCursor = (windSpeed - sailChar.Winds[firstWindIndex]) / (sailChar.Winds[secondWindIndex] - sailChar.Winds[firstWindIndex])
	}

	firstSpeed := getCursorValue(firstPolar.Speed[firstWindIndex], firstPolar.Speed[secondWindIndex], windCursor)
	secondSpeed := getCursorValue(secondPolar.Speed[firstWindIndex], secondPolar.Speed[secondWindIndex], windCursor)

	speed = getCursorValue(firstSpeed, secondSpeed, angleCrusor)

	return
}
