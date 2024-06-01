package packet

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"math/rand"
	"openra-server/pkg/protocol"

	"gopkg.in/yaml.v3"
)

type HandshakeRequestData struct {
	Mod       string `yaml:"Mod"`
	Version   string `yaml:"Version"`
	AuthToken string `yaml:"AuthToken"`
}

type HandshakeRequest struct {
	HandshakeRequest HandshakeRequestData `yaml:"HandshakeRequestHandshake"`
}

func (pk HandshakeRequest) Build() ([]byte, error) {
	// Make a random token for client to sign with their authid
	token := base64.StdEncoding.EncodeToString(protocol.MakeArray(256, func(_ int) byte {
		return byte(rand.Intn(256))
	}))

	pk.HandshakeRequest.AuthToken = token

	// Marshal the handshake request to YAML
	data, err := yaml.Marshal(pk)
	println(string(data))
	if err != nil {
		return nil, err
	}

	// Calculate the length of the packet (data length + 4 bytes for length field)
	//packetLength := int32(len(data))

	// Create a buffer to construct the packet
	buf := new(bytes.Buffer)

	// Write the length of the packet
	//err = binary.Write(buf, binary.LittleEndian, packetLength)
	if err != nil {
		return nil, err
	}
	// Write sender client id
	err = binary.Write(buf, binary.LittleEndian, int32(0))
	if err != nil {
		return nil, err
	}

	// Write the YAML data
	_, err = buf.Write(data)
	if err != nil {
		return nil, err
	}

	// Return the constructed packet as bytes
	return buf.Bytes(), nil
}
