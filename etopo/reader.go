package etopo

import (
	"bytes"
	"encoding/binary"
	"image"
	"log"
	"os"

	"github.com/regattebzh/trajectory/mapper"
)

const width = 21601
const height = 10801
const dataSize = 2 // sizeof int16

//ReadRectangle reads a rectangle
func ReadRectangle(file os.File, r image.Rectangle) (mapper.Map, error) {

	r = r.Add(image.Point{180, 90})

	buffer := mapper.New(r, 1, 1)
	var err error

	fileLine := make([]byte, r.Dx()*dataSize)
	bufferLine := 0
	for line := int64(r.Min.Y + 90); line < int64(r.Max.Y+90); line++ {
		// read line in file
		offset := (line*int64(width) + int64(r.Min.X+180)) * dataSize
		file.Seek(offset, 0)
		if _, err = file.Read(fileLine); err != nil {
			log.Fatal("etopo: file.Read failed (ReadRectangle)\n", err)
		}
		// convert in an array of int16
		data := make([]int16, r.Dx())
		dataBuf := bytes.NewReader(fileLine)
		if err := binary.Read(dataBuf, binary.LittleEndian, data); err != nil {
			log.Fatal("Byte to int16 failed\n", err)
		}
		// store data
		for index, value := range data {
			buffer.Set(image.Point{index, bufferLine}, Altitude(value))
		}
		bufferLine++
	}

	return buffer, err
}
