package png

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"pngo/png/chunk"
	"pngo/png/header"
)

type PNG struct {
	Header *header.Header
	IHDR   *chunk.Chunk
	//PLTE Palette
	IDAT []*chunk.Chunk
	IEND *chunk.Chunk
}

func NewPNG() *PNG {
	return &PNG{
		Header: header.New(),
		IHDR:   chunk.NewIHDR(1, 1, 8, 2, 0, 0, 0),
		IDAT: []*chunk.Chunk{
			chunk.NewIDAT([]byte{0x78, 0x01, 0x63, 0x60, 0x00, 0x00, 0x00, 0x02, 0x00, 0x01}),
		},
		IEND: chunk.NewIEND(),
	}
}

func (p *PNG) Bytes() []byte {
	var data []byte

	data = append(data, p.Header.Bytes()...)
	data = append(data, p.IHDR.Bytes()...)

	for _, chunk := range p.IDAT {
		data = append(data, chunk.Bytes()...)
	}

	data = append(data, p.IEND.Bytes()...)

	return data
}

func (p *PNG) String() string {
	var IDATs string

	for _, chunk := range p.IDAT {
		IDATs += chunk.String()
	}

	return p.Header.String() + " " + p.IHDR.String() + " " + IDATs + " " + p.IEND.String()
}

func (p *PNG) UserFormatString() string {
	var IDATs string

	for _, chunk := range p.IDAT {
		IDATs += "(\n" + chunk.UserFormatString() + ")\n"
	}

	return "PNG signature: " + p.Header.String() + "\n" + "Image header: [\n" + p.IHDR.UserFormatString() + "]\n" + "Image data: [\n" + IDATs + "]\n" + "End of image: [\n" + p.IEND.UserFormatString() + "]"
}

func PNGFromFile(file *os.File) *PNG {
	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var png PNG

	if bytes.Equal(data[:8], header.New().StandardBytes()) && bytes.Equal(data[len(data)-12:], chunk.StandardIEND()) {
		png = PNG{
			Header: header.New(),
			IHDR:   chunk.IHDRFromBytes(data[8:33]),
			IDAT:   chunk.ChunksFromBytes(data[33 : len(data)-12]),
			IEND:   chunk.NewIEND(),
		}
	} else {
		panic(fmt.Sprintf("File %s is not a PNG file", file.Name()))
	}

	return &png
}
