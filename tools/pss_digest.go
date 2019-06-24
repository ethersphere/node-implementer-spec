package main

import (
	"crypto/rand"
	"encoding/hex"
	"flag"
	"os"

	"github.com/ethereum/go-ethereum/rlp"
	whisper "github.com/ethereum/go-ethereum/whisper/whisperv6"
	"github.com/ethersphere/swarm/storage"
)

type digestStruct struct {
	To      []byte
	Payload *whisper.Envelope
}

func main() {

	var addrString *string
	var addrLength *int
	var addr []byte
	var topicString *string
	var topic whisper.TopicType
	var nodigest *bool

	addrLength = flag.Int("l", 32, "address length")
	addrString = flag.String("a", "", "literal address bytes in hex (overrides -l)")
	topicString = flag.String("t", "00000000", "set topic (default: 0x00000000")
	nodigest = flag.Bool("r", false, "set to get rlp encoding used for digest (default: false")
	flag.Parse()

	// set the data for the payload
	data, err := hex.DecodeString(flag.Arg(0))
	if err != nil {
		data = []byte(flag.Arg(0))
	}

	// set random or explicit address, depending on input
	if *addrString == "" {
		if *addrLength > 32 || *addrLength < 0 {
			panic("Address must be 0 <= n <= 32")
		}
		addr = make([]byte, *addrLength)
		c, err := rand.Read(addr)
		if err != nil {
			panic(err)
		} else if c != *addrLength {
			panic("short read")
		}
	} else {
		addr, err = hex.DecodeString(*addrString)
		if err != nil {
			panic(err)
		}
	}

	// parse the topic string
	topicBytes, err := hex.DecodeString(*topicString)
	if err != nil {
		panic(err)
	}
	copy(topic[:], topicBytes)

	env := &whisper.Envelope{
		Expiry: 0,
		TTL:    0,
		Topic:  topic,
		Data:   data,
		Nonce:  0,
	}
	msg := digestStruct{
		To:      addr,
		Payload: env,
	}

	serializedMsg, err := rlp.EncodeToBytes(msg)
	if err != nil {
		panic(err)
	}

	if *nodigest {
		os.Stdout.Write(serializedMsg)
		return
	}

	hasher := storage.MakeHashFunc(storage.DefaultHash)()
	hasher.Write(serializedMsg)
	digest := hasher.Sum(nil)
	os.Stdout.Write(digest)
}
