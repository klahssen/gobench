package payload

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"io"
	"log"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	//"github.com/stretchr/testify/assert"
)

var (
	p20auds Payload
)

func init() {
	p20auds = genPayload(20)
}

func TestProtoMarshaling(t *testing.T) {
	t.Log("Payload comparisons with 20 audiences and ids 14chars long")
	b, err := proto.Marshal(&p20auds)
	if err != nil {
		t.Errorf("failed to marshal: %v", err)
		return
	}
	log.Printf("%d bytes", len(b))
	dest := Payload{}
	err = proto.Unmarshal(b, &dest)
	if err != nil {
		t.Errorf("failed to unmarshal: %v", err)
		return
	}
	//assert.Equal(t, p10auds, dest)
	if err = payloadsAreEqual(p20auds, dest); err != nil {
		t.Errorf("%v", err)
	}
}

func TestCompareMarshaledSize(t *testing.T) {
	//protobuf
	t0 := time.Now()
	b1, err := proto.Marshal(&p20auds)
	if err != nil {
		t.Errorf("failed to proto.Marshal")
		return
	}
	t.Logf("proto.Marshal: %d bytes (in %s)", len(b1), time.Since(t0))
	var zbuf bytes.Buffer
	zw := zlib.NewWriter(&zbuf)
	t0 = time.Now()
	nzw, err := zw.Write(b1)
	if err != nil {
		t.Errorf("zlib.Write proto payload: %v", err)
		return
	}
	if nzw != len(b1) {
		t.Errorf("zlib: wrote %d bytes instead of %d", nzw, len(b1))
	}
	zw.Close()
	b3 := zbuf.Bytes()
	t.Logf("zlib.Compress(proto): %d bytes (in %s)", len(b3), time.Since(t0))
	//t.Logf("%q", zbuf.Bytes())
	zr, err := zlib.NewReader(&zbuf)
	if err != nil {
		t.Errorf("zlib.NewReader: %v", err)
	}
	b4 := make([]byte, len(b1))
	t0 = time.Now()
	nzr, err := zr.Read(b4)
	if err != nil && err != io.EOF {
		t.Errorf("zlib.Read proto payload: %v", err)
		return
	}
	t.Logf("zlib.Decompress(proto) in %s", time.Since(t0))
	if nzr != len(b1) {
		t.Errorf("zlib: read %d bytes instead of %d", nzr, len(b1))
		return
	}
	dest := Payload{}
	t0 = time.Now()
	if err = proto.Unmarshal(b4, &dest); err != nil {
		t.Errorf("json.Unmarshal after zlib decompress: %v", err)
		return
	}
	t.Logf("proto.Unmarshal in %s", time.Since(t0))
	if err = payloadsAreEqual(p20auds, dest); err != nil {
		t.Errorf("proto.Unmarshal of zlib.Decompressed: invalid result: %v", err)
		return
	}
	//t.Logf("%q", b1)
	//json ==========================================================================
	t0 = time.Now()
	b2, err := json.Marshal(&p20auds)
	if err != nil {
		t.Errorf("failed to json.Marshal")
		return
	}
	t.Logf("json.Marshal: %d bytes (in %s)", len(b2), time.Since(t0))
	//t.Logf("%q", b2)
	//var zbuf bytes.Buffer
	zw = zlib.NewWriter(&zbuf)
	t0 = time.Now()
	nzw, err = zw.Write(b2)
	if err != nil {
		t.Errorf("zlib.Write JSON payload: %v", err)
		return
	}
	if nzw != len(b2) {
		t.Errorf("zlib: wrote %d bytes instead of %d", nzw, len(b2))
	}
	zw.Close()
	b3 = zbuf.Bytes()
	t.Logf("zlib.Compress(json): %d bytes (in %s)", len(b3), time.Since(t0))
	//t.Logf("%q", zbuf.Bytes())
	zr, err = zlib.NewReader(&zbuf)
	if err != nil {
		t.Errorf("zlib.NewReader: %v", err)
	}
	b4 = make([]byte, len(b2))
	t0 = time.Now()
	nzr, err = zr.Read(b4)
	if err != nil && err != io.EOF {
		t.Errorf("zlib.Read JSON payload: %v", err)
		return
	}
	t.Logf("zlib.Decompress(json) in %s", time.Since(t0))
	if nzr != len(b2) {
		t.Errorf("zlib: read %d bytes instead of %d", nzr, len(b2))
		return
	}
	dest = Payload{}
	t0 = time.Now()
	if err = json.Unmarshal(b4, &dest); err != nil {
		t.Errorf("json.Unmarshal after zlib decompress: %v", err)
		return
	}
	t.Logf("json.Unmarshal in %s", time.Since(t0))
	if err = payloadsAreEqual(p20auds, dest); err != nil {
		t.Errorf("json.Unmarshal of zlib.Decompressed: invalid result: %v", err)
		return
	}
}
