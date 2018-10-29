package client

import (
	"encoding/binary"
)

// Packet class.
type Packet struct {
	Request string
	Data    []byte
}

// DataLen Packet method, return 8 bytes with packet length in little endian order.
func (p *Packet) DataLen() []byte {
	dataLen := make([]byte, 8)
	// JSONData, _ := json.Marshal(p)
	binary.LittleEndian.PutUint32(dataLen, uint32(len(p.Data)))
	return dataLen
}
