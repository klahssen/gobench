package items

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type person struct {
	Name    string    `json:"name"`
	Age     int       `json:"age"`
	Friends []*person `json:"friends"`
}

func encodeLinesAppend(nPerson, nFriendsPerPers int) ([]byte, error) {
	//l := make([]*person, 0, nPerson)
	f := make([]*person, 0, nFriendsPerPers)
	for i := 1; i <= nFriendsPerPers; i++ {
		f = append(f, &person{
			Name: fmt.Sprintf("Friend%02d", i),
			Age:  30,
		})
	}
	b := new(bytes.Buffer)
	enc := json.NewEncoder(b)
	var p person
	var err error
	b.Write([]byte("["))
	for i := 1; i <= nPerson; i++ {
		p = person{
			Name:    fmt.Sprintf("Person%02d", i),
			Age:     30,
			Friends: f,
		}
		if err = enc.Encode(p); err != nil {
			return nil, fmt.Errorf("encode: %v", err)
		}
		if i < nPerson {
			b.Write([]byte(","))
		}
	}
	b.Write([]byte("]"))
	return b.Bytes(), nil
}
func decodeLinesLoop(b []byte) ([]person, error) {
	dec := json.NewDecoder(bytes.NewBuffer(b))
	var err error
	//var token json.Token
	if _, err := dec.Token(); err != nil {
		return nil, fmt.Errorf("check for begining token: %v", err)
	}
	//log.Printf("token: '%v'", token)
	var p person
	res := make([]person, 0, 50)
	var n int64
	for dec.More() {
		if err = dec.Decode(&p); err != nil {
			return nil, fmt.Errorf("decode: iter [%d]: %v", n, err)
		}
		n++
		res = append(res, p)
	}
	return res, nil
}
