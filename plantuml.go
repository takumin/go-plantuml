package plantuml

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"errors"
	"io"
	"io/ioutil"
	"unicode/utf8"
)

const mapper = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_"

func Encode(p []byte) (string, error) {
	if !utf8.Valid(p) {
		return "", errors.New("invalid utf8 string")
	}

	r := bytes.NewReader(p)

	zpr, zpw := io.Pipe()
	zenc, err := zlib.NewWriterLevel(zpw, zlib.BestCompression)
	if err != nil {
		return "", err
	}
	go func() {
		_, err = io.Copy(zenc, r)
		zenc.Close()
		if err != nil {
			zpw.CloseWithError(err)
		} else {
			zpw.Close()
		}
	}()

	penc := base64.NewEncoding(mapper).WithPadding(base64.NoPadding)

	bpr, bpw := io.Pipe()
	benc := base64.NewEncoder(penc, bpw)
	go func() {
		_, err = io.Copy(benc, zpr)
		benc.Close()
		if err != nil {
			bpw.CloseWithError(err)
		} else {
			bpw.Close()
		}
	}()

	res, err := ioutil.ReadAll(bpr)
	return string(res), err
}
