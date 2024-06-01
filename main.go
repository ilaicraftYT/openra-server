package main

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"net"
	"openra-server/pkg/protocol"
	"openra-server/pkg/protocol/packet"
	"time"
)

var clientCount int32 = 0

func main() {
	rand.Seed(time.Now().UnixNano())
	listener, err := net.Listen("tcp", "0.0.0.0:1235")

	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	fmt.Println("holas")
	clientCount++
	hello := packet.ServerHello{Handshake: protocol.ProtocolVersion.Handshake, OrderTypes: protocol.ProtocolVersion.Orders}.Build(clientCount)
	_, err := conn.Write(hello)
	if err != nil {
		panic(err)
	}

	serverHandshake, err := packet.HandshakeRequest{HandshakeRequest: packet.HandshakeRequestData{Mod: "ra", Version: "release-20231010"}}.Build()
	if err != nil {
		panic(err)
	}
	conn.Write(serverHandshake)

	for {
		// Read the first 4 bytes to determine the packet length
		header := make([]byte, 4)
		_, err := conn.Read(header)
		if err != nil {
			panic(err)
		}

		// Convert the header to an integer (assuming big-endian)
		packetLength := binary.LittleEndian.Uint32(header)
		if packetLength > 128*1024 {
			panic(packetLength)
		}

		// Allocate buffer based on the packet length
		buffer := make([]byte, packetLength)

		// Read the rest of the packet
		_, err = conn.Read(buffer)
		if err != nil {
			panic(err)
		}

		parsePacket(buffer)
	}
}

func parsePacket(data []byte) {
	if len(data) < 12 {
		fmt.Println("Packet too short")
		return
	}

	clientID := binary.LittleEndian.Uint32(data[0:4])
	gameTick := binary.LittleEndian.Uint32(data[4:8])
	ordersData := data[8:]

	fmt.Printf("Client ID: %d, Game Tick: %d\n", clientID, gameTick)

	// Process orders
	offset := 0
	for offset < len(ordersData) {
		if offset+1 > len(ordersData) {
			fmt.Println("Incomplete order found")
			return
		}
		orderType := ordersData[offset]
		offset++
		fmt.Printf("Order Type: %d\n", orderType)
		orderData := ordersData[offset:]

		switch orderType {
		case 0xFE:
			var handshakeRes packet.HandshakeResponse
			handshakeRes.From(orderData)
			fmt.Println(handshakeRes)
		}
	}
}
