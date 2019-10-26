package proto

import (
	"log"
	"time"
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
	h.XXX_Unmarshal(helloBytes)
}

func MarshalHello() []byte {
	h := &Hello{Name: "Superman"}
	b, err := h.XXX_Marshal(nil, true)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func UnmarshalProfile() {
	h := &Hello{}
	h.XXX_Unmarshal(helloBytes)
}

func MarshalProfile() []byte {
	ts := time.Now().Unix()
	h := &Profile{Crea: ts, Upd: ts, Fname: "Clark", Lname: "Kent", Age: 25, Married: false, City: "metropolis", Country: "United-States", Hobbies: []string{"flying", "swagging", "saving people", "being a hero", "journalism"}}
	b, err := h.XXX_Marshal(nil, true)
	if err != nil {
		log.Fatal(err)
	}
	return b
}
