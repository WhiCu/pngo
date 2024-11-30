package chunk

func NewIDAT(data []byte) *Chunk {
	c := Chunk{
		Type: [4]byte{'I', 'D', 'A', 'T'},
		Data: data,
	}
	copy(c.Length[:], U32toBytes(uint32(len(data))))

	CRC32(&c)
	return &c
}
