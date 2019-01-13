package plantuml

import (
	"strings"
	"testing"
)

var (
	raw = "Alice -> Bob: Hello World!"
	enc = "Syp9J4vLqBLJSCfFibBmICt9oLS8po_AIL440000"
)

func TestEncodeSuccess(t *testing.T) {
	buf := strings.NewReader(raw)

	res, err := Encode(buf)
	if err != nil {
		t.Fatalf("exists error %#v", err)
	}
	if res != enc {
		t.Logf("success: %#v", enc)
		t.Logf("convert: %#v", res)
		t.Fatalf("converted strings do not match")
	}
}
