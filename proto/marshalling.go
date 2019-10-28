package proto

import (
	"log"
	"time"

	"github.com/golang/protobuf/proto"
)

var (
	helloBytes   []byte
	profileBytes []byte
)

func init() {
	helloBytes = MarshalHello()
	profileBytes = MarshalProfile()
}

func UnmarshalHello() {
	h := &Hello{}
	proto.Unmarshal(helloBytes, h)
}

func MarshalHello() []byte {
	h := &Hello{Name: "Superman"}
	b, err := proto.Marshal(h)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func UnmarshalProfile() {
	p := &Profile{}
	proto.Unmarshal(profileBytes, p)
}

func MarshalProfile() []byte {
	ts := time.Now().Unix()
	p := &Profile{Crea: ts, Upd: ts, Fname: "Clark", Lname: "Kent", Age: 25, Married: false, City: "metropolis", Country: "United-States", Hobbies: []string{"flying", "swagging", "saving people", "being a hero", "journalism"}}
	b, err := proto.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	return b
}
