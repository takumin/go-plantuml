package plantuml

import (
	"testing"
)

var (
	raw = "Alice -> Bob: Hello World!"
	enc = "UDfopCbCJbNGjLDmoa-oKl18pSd9LmZFByf9KGG4003__sl223G"
)

func TestEncodeSuccess(t *testing.T) {
	res, err := Encode([]byte(raw))
	if err != nil {
		t.Fatalf("exists error %#v", err)
	}
	if res != enc {
		t.Logf("success: %#v", enc)
		t.Logf("convert: %#v", res)
		t.Fatalf("converted strings do not match")
	}
}
