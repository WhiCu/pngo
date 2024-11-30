package header

import "fmt"

type Header struct {
	NonASCII  byte
	PNG       []byte
	NewLine   []byte
	EndOfFile byte
	LF        byte
}

func New() *Header {
	return &Header{
		NonASCII:  0x89,
		PNG:       []byte{0x50, 0x4E, 0x47}, // Changing to byte values for PNG
		NewLine:   []byte{0x0D, 0x0A},       // Changing to byte values for NewLine
		EndOfFile: 0x1A,                     // Changing EndOfFile to 0x1A
		LF:        0x0A,
	}
}

func (c *Header) Bytes() []byte {
	var data []byte

	data = append(data, c.NonASCII)
	data = append(data, c.PNG...)
	data = append(data, c.NewLine...)
	data = append(data, c.EndOfFile, c.LF)

	return data

}

func (c *Header) StandardBytes() []byte {
	return []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}
}

// String returns a string representation of the Header fields in hexadecimal format.
// Each byte of the NonASCII, PNG, NewLine, EndOfFile, and LF fields are converted to
// a two-digit hex value with a space separator.
func (h Header) String() string {
	var result string

	// Convert each field to hex format
	result += fmt.Sprintf("%02X ", h.NonASCII)

	for _, value := range h.PNG {
		result += fmt.Sprintf("%02X ", value)
	}

	for _, value := range h.NewLine {
		result += fmt.Sprintf("%02X ", value)
	}

	// Include EndOfFile and LF
	result += fmt.Sprintf("%02X %02X", h.EndOfFile, h.LF)

	return result
}
