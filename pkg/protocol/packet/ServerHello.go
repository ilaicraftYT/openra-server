package packet

import (
	"bytes"
	"encoding/binary"
)

type ServerHello struct {
	Handshake  int32
	OrderTypes int32
}

func (s ServerHello) Build(clientCount int32) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, s.Handshake)
	binary.Write(buf, binary.LittleEndian, clientCount)
	return buf.Bytes()
}
