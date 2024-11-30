package chunk

import (
	"fmt"
	"hash/crc32"
)

type Chunk struct {
	Length [4]byte
	Type   [4]byte
	Data   []byte
	CRC    [4]byte
}

func (c *Chunk) Bytes() []byte {
	var data []byte

	data = append(data, c.Length[:]...)
	data = append(data, c.Type[:]...)
	data = append(data, c.Data...)
	data = append(data, c.CRC[:]...)

	return data

}

// CRC32 computes the CRC32 checksum of the chunk type and data, and updates the
// chunk's CRC field with the result. The updated chunk is returned.
func CRC32(c *Chunk) [4]byte {
	h := crc32.NewIEEE()

	h.Write(c.Type[:])

	if c.Data != nil {
		h.Write(c.Data)
	}

	checksum := h.Sum32()

	var crcBytes [4]byte
	crcBytes[0] = byte(checksum >> 24)
	crcBytes[1] = byte(checksum >> 16)
	crcBytes[2] = byte(checksum >> 8)
	crcBytes[3] = byte(checksum)

	c.CRC = crcBytes

	return crcBytes
}

func (c *Chunk) String() string {
	if c == nil {
		return ""
	}

	var result string

	for _, value := range c.Length {
		result += fmt.Sprintf("%02X ", value)
	}

	for _, value := range c.Type {
		result += fmt.Sprintf("%02X ", value)
	}

	for _, value := range c.Data {
		result += fmt.Sprintf("%02X ", value)
	}

	for _, value := range c.CRC {
		result += fmt.Sprintf("%02X ", value)
	}

	return result[:len(result)-1]
}

func (c *Chunk) UserFormatString() string {
	if c == nil {
		return ""
	}

	var result string

	result += fmt.Sprintf("length: %d\n", BytesToU32(c.Length[:]))

	result += fmt.Sprintf("type: %s\n", string(c.Type[:]))

	result += "data: "
	for _, value := range c.Data {
		result += fmt.Sprintf("%02X ", value)
	}
	result += "\n"

	result += "crc: "
	for _, value := range c.CRC {
		result += fmt.Sprintf("%02X ", value)
	}
	result += "\n"

	return result
}

func U32toBytes(u uint32) []byte {
	var b [4]byte
	b[0] = byte(u >> 24)
	b[1] = byte(u >> 16)
	b[2] = byte(u >> 8)
	b[3] = byte(u)
	return b[:]
}

func BytesToU32(b []byte) uint32 {
	return uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])
}

func ChunksFromBytes(data []byte) []*Chunk {
	var chunks []*Chunk

	for count := 0; count < len(data); {
		chunk := Chunk{
			Length: [4]byte(data[count : count+4]),
			Type:   [4]byte(data[count+4 : count+8]),
		}

		length := int(BytesToU32(chunk.Length[:]))

		chunk.Data = data[count+8 : count+8+length]
		chunk.CRC = [4]byte(data[count+8+length : count+8+length+4])
		chunks = append(chunks, &chunk)

		count += length + 12
	}
	return chunks
}
