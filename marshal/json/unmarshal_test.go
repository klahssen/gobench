package json

import (
	"encoding/json"
	"log"
	"testing"
)

var (
	payloadBytesX100 []byte
	payloadX100      []payload
)

func init() {
	var err error
	payloadX100 = generatePayloadSlice(100)
	payloadBytesX100, err = json.Marshal(payloadX100)
	if err != nil {
		log.Fatal(err)
	}
}

func BenchmarkDecodeLoop100items(b *testing.B) {
	for i := 1; i <= b.N; i++ {
		dest := []payload{}
		if err := manualDecodeLoop(payloadBytesX100, &dest); err != nil {
			log.Fatal(err)
		}
		//b.Log(len(dest))
	}
}

func BenchmarkStdUnmarshal100items(b *testing.B) {
	for i := 1; i <= b.N; i++ {
		dest := []payload{}
		if err := stdUnmarshal(payloadBytesX100, &dest); err != nil {
			log.Fatal(err)
		}
		//b.Log(len(dest))
	}
}
