package wind

import (
	"bytes"
	"encoding/binary"
	"image"
	"io"
	"log"

	"github.com/regattebzh/trajectory/mapper"
)

func readBlock(file io.Reader) (count uint32, winds []float32, err error) {
	size := make([]byte, 4)

	// read data block
	if _, err = file.Read(size); err != nil {
		log.Fatal("file.Read failed (ReadBlock)\n", err)
		return
	}
	countBuf := bytes.NewReader(size)
	if err = binary.Read(countBuf, binary.LittleEndian, &count); err != nil {
		log.Fatal("Byte to uint32 failed\n", err)
		return
	}

	// read data
	windb := make([]byte, count)
	if _, err = file.Read(windb); err != nil {
		log.Fatal("file.Read failed (ReadBlock)\n", err)
		return
	}

	// read the size again
	if _, err = file.Read(size); err != nil {
		log.Fatal("file.Read failed (ReadBlock)\n", err)
		return
	}

	count = count / 4

	winds = make([]float32, count)
	windsBuf := bytes.NewReader(windb)
	if err = binary.Read(windsBuf, binary.LittleEndian, &winds); err != nil {
		log.Fatal("Byte to float32 failed\n", err)
		return
	}

	return
}

// Read reads wind from binary file
func Read(file io.Reader) (mapper.Map, error) {
	countU, windsU, err := readBlock(file)
	countV, windsV, err := readBlock(file)

	count := countU
	if count > countV {
		count = countV
	}

	windMap := mapper.Map{
		Height: 181,
		Width:  360,
		CellH:  60,
		CellW:  60,
		Data:   make([]mapper.Element, 181*360),
	}

	phi := -90
	lambda := 0
	for i := uint32(0); i < count; i++ {
		if lambda >= 180 {
			// lambda[180 - 359] => x=[0-179]
			SetWind(
				windMap,
				image.Point{
					lambda - 180,
					phi + 90,
				},
				Speed{
					SpeedU: windsU[i],
					SpeedV: windsV[i],
				},
			)
		} else {
			// lambda[0 - 179] => x=[180-359]
			SetWind(
				windMap,
				image.Point{
					lambda + 180,
					phi + 90,
				},
				Speed{
					SpeedU: windsU[i],
					SpeedV: windsV[i],
				},
			)
		}

		lambda = lambda + 1

		if lambda >= 360 {
			phi++
			lambda = 0
		}
	}

	return windMap, err
}
