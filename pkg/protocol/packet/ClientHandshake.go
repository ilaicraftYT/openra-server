package packet

import "gopkg.in/yaml.v3"

type HandshakeResponseClientData struct {
	Color          string `yaml:"Color"`
	PreferredColor string `yaml:"PreferredColor"`
	Password       string `yaml:"Password"`
}

type HandshakeResponseData struct {
	Mod            string                      `yaml:"Mod"`
	Version        string                      `yaml:"Version"`
	Client         HandshakeResponseClientData `yaml:"Client"`
	Fingerprint    string                      `yaml:"Fingerprint"`
	AuthSignature  string                      `yaml:"AuthSignature"`
	OrdersProtocol int32                       `yaml:"OrdersProtocol"`
}

type HandshakeResponse struct {
	HandshakeResponse HandshakeResponseData `yaml:"HandshakeResponse"`
}

func (res *HandshakeResponse) From(encoded []byte) {
	yaml.Unmarshal(encoded, res)
}
