package json

import (
	"bytes"
	"encoding/json"
)

type payload struct {
	ProfileID string   `json:"id"`
	List1     []string `json:"l1"`
	List2     []string `json:"l2"`
}

func generatePayloadSlice(n int) []payload {
	l := make([]payload, 0, n)
	p := payload{
		ProfileID: "abcdefghijkl",
		List1:     []string{"abcdefghijkl", "abcdefghijkl", "abcdefghijkl", "abcdefghijkl", "abcdefghijkl", "abcdefghijkl", "abcdefghijkl", "abcdefghijkl", "abcdefghijkl", "abcdefghijkl", "abcdefghijkl"},
		List2:     []string{"abcdefghijkl", "abcdefghijkl", "abcdefghijkl", "abcdefghijkl", "abcdefghijkl", "abcdefghijkl", "abcdefghijkl", "abcdefghijkl", "abcdefghijkl", "abcdefghijkl", "abcdefghijkl"},
	}
	for i := 1; i <= n; i++ {
		l = append(l, p)
	}
	return l
}

func manualDecodeLoop(b []byte, dest *[]payload) error {
	p := payload{}
	dec := json.NewDecoder(bytes.NewBuffer(b))
	dec.Token()
	for dec.More() {
		if err := dec.Decode(&p); err != nil {
			return err
		}
		*dest = append(*dest, p)
	}
	return nil
}

func stdUnmarshal(b []byte, dest *[]payload) error {
	return json.Unmarshal(b, dest)
}
