package common

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"io"
	"strings"
)

func Decompress(compressed []byte) ([]byte, error) {
	r, err := zlib.NewReader(bytes.NewReader(compressed))
	if err != nil {
		return nil, err
	}
	return io.ReadAll(r)
}

// lzib and base64 and replace '+','/'
func PobEncode(data []byte) (string, error) {
	buf := new(bytes.Buffer)
	//error never happen because of flate.DefaultCompression
	w := zlib.NewWriter(buf)
	_, err := w.Write(data)
	if err != nil {
		return "", err
	}
	w.Close()

	encoded := base64.StdEncoding.EncodeToString(buf.Bytes())
	encoded = strings.ReplaceAll(encoded, "+", "-")
	return strings.ReplaceAll(encoded, "/", "_"), nil
}

func PobDecode(code string) (string, error) {
	code = strings.ReplaceAll(code, "-", "+")
	code = strings.ReplaceAll(code, "_", "/")

	compressed, err := base64.StdEncoding.DecodeString(code)
	if err != nil {
		return "", err
	}

	r, err := zlib.NewReader(bytes.NewReader(compressed))
	if err != nil {
		return "", err
	}
	data, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
