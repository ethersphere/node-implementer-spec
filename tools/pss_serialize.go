package main

import (
	"crypto/rand"
	"encoding/hex"
	"flag"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/rlp"
	whisper "github.com/ethereum/go-ethereum/whisper/whisperv6"
	"github.com/ethersphere/swarm/pss"
)

func main() {

	var addrString *string
	var addrLength *int
	var addr []byte
	var expireInt *int
	var expire uint32
	var topicString *string
	var topic whisper.TopicType

	nowTime := int(time.Now().Unix())

	addrLength = flag.Int("l", 32, "address length")
	addrString = flag.String("a", "", "literal address bytes in hex (overrides -l)")
	expireInt = flag.Int("e", nowTime, "set expire timestamp (default: now)")
	topicString = flag.String("t", "00000000", "set topic (default: 0x00000000")
	flag.Parse()

	// set the data for the payload
	data, err := hex.DecodeString(flag.Arg(0))
	if err != nil {
		data = []byte(flag.Arg(0))
	}

	// set random or explicit address, depending on input
	if addrString == nil {
		if *addrLength > 32 || *addrLength < 0 {
			panic("Address must be 0 <= n <= 32")
		}
		addr = make([]byte, *addrLength)
		rand.Read(addr)
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

	// cast expire to correct uint
	expire = uint32(*expireInt)

	env := &whisper.Envelope{
		Expiry: 0,
		TTL:    0,
		Topic:  topic,
		Data:   data,
		Nonce:  0,
	}
	msg := pss.PssMsg{
		To:      addr,
		Control: []byte{0x02}, // forces raw message mode
		Expire:  expire,
		Payload: env,
	}

	serializedMsg, err := rlp.EncodeToBytes(msg)

	os.Stdout.Write(serializedMsg)
}
