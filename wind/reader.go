package wind

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
)

// ReadBlock reads a block of binary
func ReadBlock(file io.Reader) (count uint32, winds []float32, err error) {
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

// ReadWind reads wind from binary file
func ReadWind(file io.Reader) (winds []Wind, err error) {
	countU, windsU, err := ReadBlock(file)
	countV, windsV, err := ReadBlock(file)

	count := countU
	if count > countV {
		count = countV
	}

	winds = make([]Wind, count)

	phi := float64(0)
	lambda := float64(-90)
	for i := uint32(0); i < count; i++ {
		winds[i] = Wind{
			lambda,
			phi,
			float64(windsU[i]),
			float64(windsV[i]),
		}
		if lambda > 359 {
			phi++
			lambda = 0
		}
	}
	return
}
