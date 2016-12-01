package etopo

import (
	"bytes"
	"encoding/binary"
	"image"
	"io"
	"log"
	"os"

	"github.com/regattebzh/trajectory/mapper"
)

const width = 21601
const height = 10801
const dataSize = 2 // sizeof int16

//Read read ETOPO binary
func Read(file io.Reader) (mapper.Map, error) {

	buffer := mapper.New(image.Rect(0, 0, width, height), 1, 1)
	preData := make([]byte, width*height*dataSize)
	if _, err := file.Read(preData); err != nil {
		log.Fatal("file.Read failed (ReadEtopo)\n", err)
	}

	data := make([]int16, width*height)
	dataBuf := bytes.NewReader(preData)
	if err := binary.Read(dataBuf, binary.LittleEndian, data); err != nil {
		log.Fatal("Byte to int16 failed\n", err)
	}

	for index, value := range data {
		buffer.Data[index] = Altitude(value)
	}

	buffer.ComputeParameters()

	return buffer, nil
}

//ReadRectangle reads a rectangle
func ReadRectangle(file os.File, r image.Rectangle) (mapper.Map, error) {

	buffer := mapper.New(r, 1, 1)
	var err error

	fileLine := make([]byte, r.Dx()*dataSize)
	bufferLine := 0
	for line := int64(r.Min.Y); line < int64(r.Max.Y); line++ {
		// read line in file
		offset := (line*int64(width) + int64(r.Min.X)) * dataSize
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
			buffer.Data[index+r.Dx()*bufferLine] = Altitude(value)
		}
		bufferLine++
	}

	return buffer, err
}
