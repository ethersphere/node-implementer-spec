package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethersphere/swarm/network"
)

func main() {
	addr := make([]byte, 32)
	key := make([]byte, 65)

	c, err := rand.Read(addr)
	if err != nil {
		panic(err)
	} else if c != 32 {
		panic("short read")
	}

	key[0] = 0x04
	c, err = rand.Read(key[1:])
	if err != nil {
		panic(err)
	} else if c != 64 {
		panic("short read")
	}

	keyHex := make([]byte, len(key)*2)
	hex.Encode(keyHex, key)
	enodeBytes := append([]byte("enode://"), keyHex...)
	bzzAddr := &network.BzzAddr{
		UAddr: enodeBytes,
		OAddr: addr,
	}

	addrHex := make([]byte, len(addr)*2)
	hex.Encode(addrHex, addr)
	fmt.Fprintf(os.Stderr, string(addrHex)+"\n")
	fmt.Fprintf(os.Stderr, string(enodeBytes)+"\n")

	msg := network.HandshakeMsg{
		Version:   42,
		NetworkID: 622,
		Addr:      bzzAddr,
		LightNode: true,
	}

	serializedMsg, err := rlp.EncodeToBytes(msg)

	os.Stdout.Write(serializedMsg)
}
