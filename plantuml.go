package plantuml

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"errors"
	"io"
	"io/ioutil"
	"unicode/utf8"
)

func Encode(r io.Reader) (string, error) {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	if !utf8.Valid(buf) {
		return "", errors.New("invalid utf8 string")
	}

	fpr, fpw := io.Pipe()
	fenc, err := flate.NewWriter(fpw, flate.BestCompression)
	if err != nil {
		return "", err
	}
	go func() {
		defer fenc.Close()
		_, err = io.Copy(fenc, r)
		if err != nil {
			fpw.CloseWithError(err)
		} else {
			fpw.Close()
		}
	}()

	bpr, bpw := io.Pipe()
	benc := base64.NewEncoder(base64.URLEncoding.WithPadding('\x00'), bpw)
	go func() {
		defer benc.Close()
		_, err = io.Copy(benc, fpr)
		if err != nil {
			bpw.CloseWithError(err)
		} else {
			bpw.Close()
		}
	}()

	res := new(bytes.Buffer)
	res.ReadFrom(bpr)
	return res.String(), nil
}
