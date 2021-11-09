package markup

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"io"
)

func Encode(markup []byte) string {
	out := &bytes.Buffer{}
	b64 := base64.NewEncoder(base64.StdEncoding, out)
	writer := zlib.NewWriter(b64)
	_, _ = writer.Write(markup)
	_ = writer.Flush()
	_ = writer.Close()
	return out.String()
}

func Decode(encoded string) []byte {
	outBuf := &bytes.Buffer{}
	in := bytes.NewBufferString(encoded)
	b64 := base64.NewDecoder(base64.StdEncoding, in)
	reader, _ := zlib.NewReader(b64)
	_, _ = io.Copy(outBuf, reader)
	return outBuf.Bytes()
}
