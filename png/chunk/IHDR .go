package chunk

import (
	"bytes"
	"fmt"
)

func NewIHDR(width, height uint32, depth uint8, colorType uint8, compressionMethod uint8, filterMethod uint8, interlaceMethod uint8) *Chunk {
	c := Chunk{
		Length: [4]byte{0x00, 0x00, 0x00, 0x0D},
		Type:   [4]byte{'I', 'H', 'D', 'R'},
		Data:   nil,
		CRC:    [4]byte{0, 0, 0, 0},
	}
	c.Data = append(c.Data, U32toBytes(width)...)
	c.Data = append(c.Data, U32toBytes(height)...)
	c.Data = append(c.Data, depth, colorType, compressionMethod, filterMethod, interlaceMethod)

	if l, d := BytesToU32(c.Length[:]), uint32(len(c.Data)); l != d {
		panic(fmt.Sprintf("IHDR length is not correct: Length = %d, len(c.Data) = %d", l, d))
	}

	CRC32(&c)
	return &c
}

func IHDRFromBytes(data []byte) *Chunk {
	if len(data) != 25 {
		panic(fmt.Sprintf("IHDR length is not correct: len(data) = %d", len(data)))
	}
	return &Chunk{
		Length: [4]byte(bytes.Clone(data[0:4])),
		Type:   [4]byte(bytes.Clone(data[4:8])),
		Data:   bytes.Clone(data[8:21]),
		CRC:    [4]byte(bytes.Clone(data[21:25])),
	}
}
