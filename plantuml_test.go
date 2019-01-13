package plantuml

import (
	"bytes"
	"testing"
)

var (
	raw = "Alice -> Bob: Hello World!"
	enc = "Syp9J4vLqBLJSCfFibBmICt9oLS8po_AIL440000"
)

func TestEncodeSuccess(t *testing.T) {
	buf := bytes.NewBufferString(raw)

	res, err := Encode(buf)
	if err != nil {
		t.Fatalf("exists error %#v", err)
	}
	if res != enc {
		t.Fatalf("converted strings do not match")
	}
}
