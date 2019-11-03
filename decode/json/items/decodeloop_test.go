package items

import (
	"log"
	"testing"
)

var (
	linesAppend10x10Bytes []byte
)

func init() {
	var err error
	linesAppend10x10Bytes, err = encodeLinesAppend(10, 10)
	if err != nil {
		log.Fatal(err)
	}
	//log.Printf("%q\n", linesAppend10x10Bytes)
}

func BenchmarkEncodeLines10x10(b *testing.B) {
	var err error
	for i := 1; i <= b.N; i++ {
		_, err = encodeLinesAppend(10, 10)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkDecodeLinesLoop10x10(b *testing.B) {
	var err error
	//var res []person
	for i := 1; i <= b.N; i++ {
		_, err = decodeLinesLoop(linesAppend10x10Bytes)
		if err != nil {
			log.Fatal(err)
		}
		//log.Printf("%#v\n", res)
	}
}
