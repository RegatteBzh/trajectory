package etopo

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"

	"github.com/regattebzh/trajectory/mapper"
)

//Read read ETOPO binary
func Read(file io.Reader) (mapper.Map, error) {
	width := 21601
	height := 10801
	buffer := mapper.Map{
		Width:  width,
		Height: height,
		Data:   make([]mapper.Element, width*height),
		CellH:  1,
		CellW:  1,
	}
	preData := make([]byte, width*height*2)
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
